package domain

import (
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
)

const (
	workspace ID          = "https://trail.test/workspaces/alpha"
	steward   PrincipalID = "https://trail.test/principals/steward"
	provider  PrincipalID = "https://trail.test/principals/provider"
	operator  PrincipalID = "https://trail.test/principals/operator"
	agent     PrincipalID = "https://trail.test/principals/agent"
	workload  PrincipalID = "https://trail.test/principals/workload"
	verifier  PrincipalID = "https://trail.test/principals/verifier"
)

type manualClock struct{ now time.Time }

func (c *manualClock) Now() time.Time          { return c.now }
func (c *manualClock) Advance(d time.Duration) { c.now = c.now.Add(d) }

type testRig struct {
	t      *testing.T
	clock  *manualClock
	policy *RulePolicy
	kernel *Kernel
	seq    int
}

func newTestRig(t *testing.T) *testRig {
	t.Helper()
	clock := &manualClock{now: time.Date(2026, 9, 28, 10, 0, 0, 0, time.UTC)}
	policy := NewRulePolicy("https://trail.test/policies/reference", 1)
	policy.Allow(workspace, steward,
		ActionOutcomeCreate, ActionOutcomeRevise, ActionOutcomeAccept,
		ActionCommitmentPropose, ActionCommitmentAccept, ActionWorkCreate,
		ActionGrantIssue, ActionGrantRevoke, ActionLeaseAcquire,
	)
	policy.Allow(workspace, provider, ActionCommitmentAccept)
	policy.Allow(workspace, agent, ActionRunStart, ActionRunComplete)
	policy.Allow(workspace, operator, ActionRunStart, ActionRunComplete)
	policy.Allow(workspace, verifier, ActionEvidenceSubmit, ActionOutcomeAccept)
	return &testRig{
		t:      t,
		clock:  clock,
		policy: policy,
		kernel: NewKernel("https://trail.test", clock, NewSequenceIDs("urn:trail:test"), policy),
	}
}

func (r *testRig) meta(commandType string, subject ID, version uint64, actor ActorContext) CommandMeta {
	r.t.Helper()
	r.seq++
	return CommandMeta{
		SpecVersion:     SpecVersion,
		ID:              ID("urn:trail:test:command:" + formatSequence(uint64(r.seq))),
		Type:            commandType,
		Subject:         subject,
		ExpectedVersion: version,
		Actor:           actor,
		Workspace:       workspace,
		IdempotencyKey:  "test-key-" + formatSequence(uint64(r.seq)),
		SubmittedAt:     r.clock.Now(),
	}
}

func direct(principal PrincipalID) ActorContext {
	return ActorContext{Actor: principal, Accountable: principal}
}

func delegated(from, to PrincipalID) ActorContext {
	return ActorContext{
		Actor:       to,
		Accountable: from,
		Delegation:  []DelegationStep{{From: from, To: to, GrantID: "urn:trail:test:delegation"}},
	}
}

var deterministicCriterion = Criterion{
	ID:        "criterion:no-duplicate-effect",
	Statement: "Recovery creates no duplicate protected side effect.",
	Requirement: EvidenceRequirement{
		Deterministic: true,
		Independent:   true,
		Evaluators:    []PrincipalID{verifier},
	},
}

func (r *testRig) createOutcome(id ID) Outcome {
	r.t.Helper()
	result := r.kernel.CreateOutcome(CreateOutcomeCommand{
		Meta:             r.meta(CommandCreateOutcome, id, 0, direct(steward)),
		DesiredCondition: "Interrupted work can resume safely.",
		Steward:          steward,
		Scope:            "reference simulator",
		Horizon:          "before pilot",
		Criteria:         []Criterion{deterministicCriterion},
	})
	requireAccepted(r.t, result)
	value, ok := r.kernel.Outcome(id)
	if !ok {
		r.t.Fatal("created outcome was not projected")
	}
	return value
}

func (r *testRig) createActiveCommitment(id, outcomeID ID) Commitment {
	r.t.Helper()
	requireAccepted(r.t, r.kernel.ProposeCommitment(ProposeCommitmentCommand{
		Meta:             r.meta(CommandProposeCommitment, id, 0, direct(steward)),
		Debtor:           provider,
		Creditor:         steward,
		Antecedent:       "outcome is active",
		Consequent:       "deliver safe recovery",
		OutcomeID:        outcomeID,
		Deadline:         r.clock.Now().Add(24 * time.Hour),
		AcceptancePolicy: "https://trail.test/policies/acceptance",
	}))
	requireAccepted(r.t, r.kernel.AcceptCommitment(AcceptCommitmentCommand{
		Meta:         r.meta(CommandAcceptCommitment, id, 1, direct(steward)),
		TermsVersion: 1,
	}))
	requireAccepted(r.t, r.kernel.AcceptCommitment(AcceptCommitmentCommand{
		Meta:         r.meta(CommandAcceptCommitment, id, 2, direct(provider)),
		TermsVersion: 1,
	}))
	value, _ := r.kernel.Commitment(id)
	return value
}

func (r *testRig) createWork(id, outcomeID, commitmentID ID) WorkItem {
	r.t.Helper()
	requireAccepted(r.t, r.kernel.CreateWork(CreateWorkCommand{
		Meta:           r.meta(CommandCreateWork, id, 0, direct(steward)),
		OutcomeID:      outcomeID,
		CommitmentID:   commitmentID,
		Plan:           "inject a crash and verify one protected effect",
		Exclusive:      true,
		Purpose:        "safe-recovery-test",
		EvidencePolicy: "https://trail.test/evidence-policies/deterministic",
	}))
	value, _ := r.kernel.Work(id)
	return value
}

func (r *testRig) issueGrant(id, workID ID, expires time.Time) CapabilityGrant {
	r.t.Helper()
	requireAccepted(r.t, r.kernel.IssueGrant(IssueGrantCommand{
		Meta:      r.meta(CommandIssueGrant, id, 0, direct(steward)),
		Subject:   workload,
		Actions:   []Action{ActionExecuteWork},
		Resource:  workID,
		Purpose:   "safe-recovery-test",
		NotBefore: r.clock.Now(),
		ExpiresAt: expires,
	}))
	value, _ := r.kernel.Grant(id)
	return value
}

func (r *testRig) acquireLease(id, workID ID, expires time.Time) ClaimLease {
	r.t.Helper()
	requireAccepted(r.t, r.kernel.AcquireLease(AcquireLeaseCommand{
		Meta:      r.meta(CommandAcquireLease, id, 0, direct(steward)),
		WorkID:    workID,
		Holder:    workload,
		ExpiresAt: expires,
	}))
	value, _ := r.kernel.Lease(id)
	return value
}

func (r *testRig) startRun(id ID, work WorkItem, lease ClaimLease, grant CapabilityGrant) Result {
	r.t.Helper()
	return r.kernel.StartRun(StartRunCommand{
		Meta:               r.meta(CommandStartRun, id, 0, delegated(operator, agent)),
		WorkID:             work.ID,
		WorkVersion:        work.Version,
		Executor:           agent,
		Operator:           operator,
		Workload:           workload,
		Delegation:         []PrincipalID{operator, agent, workload},
		LeaseID:            lease.ID,
		FencingToken:       lease.FencingToken,
		GrantIDs:           []ID{grant.ID},
		Budget:             BudgetLimit{Unit: "tool-call", Limit: 20},
		Deadline:           r.clock.Now().Add(time.Hour),
		CheckpointPolicy:   "after-each-protected-effect",
		Environment:        "sandbox:reference",
		EvidencePolicy:     work.EvidencePolicy,
		SideEffectStrategy: "idempotency-key-and-reconcile",
	})
}

func (r *testRig) completeRun(id ID) Run {
	r.t.Helper()
	requireAccepted(r.t, r.kernel.CompleteRun(CompleteRunCommand{
		Meta:    r.meta(CommandCompleteRun, id, 1, delegated(operator, agent)),
		Summary: "fault injection passed with one protected effect",
	}))
	value, _ := r.kernel.Run(id)
	return value
}

func (r *testRig) submitEvidence(id, outcomeID ID, outcomeVersion uint64, run Run, evaluator ActorContext, method string) Result {
	r.t.Helper()
	return r.kernel.SubmitEvidence(SubmitEvidenceCommand{
		Meta:                 r.meta(CommandSubmitEvidence, id, 0, evaluator),
		OutcomeID:            outcomeID,
		OutcomeVersion:       outcomeVersion,
		CriterionID:          deterministicCriterion.ID,
		SourceRunID:          run.ID,
		SourceRunVersion:     run.Version,
		Method:               method,
		Result:               EvidencePass,
		Evaluator:            evaluator.Actor,
		EvaluatorAccountable: evaluator.Accountable,
		Artifacts: []Artifact{{
			Digest:     "sha256:" + strings.Repeat("a", 64),
			MediaType:  "application/json",
			Size:       128,
			Locator:    "artifact://fault-report",
			Producer:   workload,
			Provenance: string(run.ID),
		}},
		ValidFrom:   r.clock.Now(),
		ValidUntil:  r.clock.Now().Add(time.Hour),
		Limitations: "reference simulator only",
	})
}

func requireAccepted(t *testing.T, result Result) Event {
	t.Helper()
	if !result.Accepted() || result.Event == nil {
		t.Fatalf("expected accepted result, got class=%s code=%s reason=%s", result.Class, result.Code, result.SafeReason)
	}
	return *result.Event
}

func TestCompleteLifecycleSeparatesExecutionEvidenceAndAcceptance(t *testing.T) {
	rig := newTestRig(t)
	outcome := rig.createOutcome("outcome:recovery")

	requireAccepted(t, rig.kernel.ProposeCommitment(ProposeCommitmentCommand{
		Meta:             rig.meta(CommandProposeCommitment, "commitment:recovery", 0, direct(steward)),
		Debtor:           provider,
		Creditor:         steward,
		Antecedent:       "outcome is active",
		Consequent:       "deliver safe recovery",
		OutcomeID:        outcome.ID,
		Deadline:         rig.clock.Now().Add(24 * time.Hour),
		AcceptancePolicy: "policy:acceptance",
	}))

	tooEarly := rig.kernel.CreateWork(CreateWorkCommand{
		Meta:           rig.meta(CommandCreateWork, "work:too-early", 0, direct(steward)),
		OutcomeID:      outcome.ID,
		CommitmentID:   "commitment:recovery",
		Plan:           "must not start from assignment alone",
		Exclusive:      true,
		Purpose:        "safe-recovery-test",
		EvidencePolicy: "policy:evidence",
	})
	if tooEarly.Class != ResultInvalid || tooEarly.Code != "COMMITMENT_NOT_ACTIVE" {
		t.Fatalf("work before explicit party acceptance: got %#v", tooEarly)
	}

	requireAccepted(t, rig.kernel.AcceptCommitment(AcceptCommitmentCommand{
		Meta:         rig.meta(CommandAcceptCommitment, "commitment:recovery", 1, direct(steward)),
		TermsVersion: 1,
	}))
	requireAccepted(t, rig.kernel.AcceptCommitment(AcceptCommitmentCommand{
		Meta:         rig.meta(CommandAcceptCommitment, "commitment:recovery", 2, direct(provider)),
		TermsVersion: 1,
	}))
	commitment, _ := rig.kernel.Commitment("commitment:recovery")
	if commitment.Status != CommitmentActive {
		t.Fatalf("commitment status = %s, want active", commitment.Status)
	}

	work := rig.createWork("work:recovery", outcome.ID, commitment.ID)
	grant := rig.issueGrant("grant:execute", work.ID, rig.clock.Now().Add(2*time.Hour))
	lease := rig.acquireLease("lease:first", work.ID, rig.clock.Now().Add(time.Hour))
	requireAccepted(t, rig.startRun("run:recovery", work, lease, grant))
	if check := rig.kernel.CheckSideEffect("run:recovery", lease.FencingToken, grant.ID); !check.Allowed {
		t.Fatalf("protected effect should be authorized, got %s", check.Code)
	}
	run := rig.completeRun("run:recovery")
	if check := rig.kernel.CheckSideEffect("run:recovery", lease.FencingToken, grant.ID); check.Allowed || check.Code != "RUN_NOT_ACTIVE" {
		t.Fatalf("completed run authorized a new side effect: %#v", check)
	}
	projectedOutcome, _ := rig.kernel.Outcome(outcome.ID)
	if projectedOutcome.Status != OutcomeActive {
		t.Fatalf("run completion changed outcome acceptance: status=%s", projectedOutcome.Status)
	}

	requireAccepted(t, rig.submitEvidence("evidence:fault", outcome.ID, outcome.Version, run, direct(verifier), "deterministic-test"))
	requireAccepted(t, rig.kernel.AcceptOutcome(AcceptOutcomeCommand{
		Meta:         rig.meta(CommandAcceptOutcome, outcome.ID, 1, direct(steward)),
		AcceptanceID: "acceptance:recovery-v1",
		EvidenceIDs:  []ID{"evidence:fault"},
	}))
	projectedOutcome, _ = rig.kernel.Outcome(outcome.ID)
	if projectedOutcome.Status != OutcomeAccepted || projectedOutcome.Version != 2 {
		t.Fatalf("accepted outcome = %#v", projectedOutcome)
	}
	acceptance, ok := rig.kernel.Acceptance("acceptance:recovery-v1")
	if !ok || acceptance.OutcomeVersion != 1 || acceptance.Evidence[deterministicCriterion.ID] != "evidence:fault" {
		t.Fatalf("acceptance does not bind exact subject and evidence: %#v", acceptance)
	}

	events := rig.kernel.Events()
	rebuilt, err := Rebuild("https://trail.test", rig.clock, NewSequenceIDs("urn:trail:rebuild"), rig.policy, events)
	if err != nil {
		t.Fatalf("rebuild: %v", err)
	}
	rebuiltOutcome, _ := rebuilt.Outcome(outcome.ID)
	if !reflect.DeepEqual(projectedOutcome, rebuiltOutcome) {
		t.Fatalf("rebuilt outcome differs\noriginal: %#v\nrebuilt:  %#v", projectedOutcome, rebuiltOutcome)
	}
}

func TestIdempotencyAndOptimisticConcurrency(t *testing.T) {
	rig := newTestRig(t)
	command := CreateOutcomeCommand{
		Meta:             rig.meta(CommandCreateOutcome, "outcome:idempotent", 0, direct(steward)),
		DesiredCondition: "one accepted event",
		Steward:          steward,
		Scope:            "reference",
		Horizon:          "now",
		Criteria:         []Criterion{deterministicCriterion},
	}
	first := rig.kernel.CreateOutcome(command)
	second := rig.kernel.CreateOutcome(command)
	if !first.Accepted() || !second.Accepted() || first.Event.ID != second.Event.ID || len(rig.kernel.Events()) != 1 {
		t.Fatalf("same command did not replay one accepted result: first=%#v second=%#v events=%d", first, second, len(rig.kernel.Events()))
	}
	changed := command
	changed.DesiredCondition = "different bytes under same key"
	reused := rig.kernel.CreateOutcome(changed)
	if reused.Class != ResultInvalid || reused.Code != "IDEMPOTENCY_KEY_REUSED" {
		t.Fatalf("changed idempotent command = %#v", reused)
	}

	stale := rig.kernel.ReviseOutcome(ReviseOutcomeCommand{
		Meta:             rig.meta(CommandReviseOutcome, command.Meta.Subject, 0, direct(steward)),
		DesiredCondition: "revised",
		Steward:          steward,
		Scope:            "reference",
		Horizon:          "now",
		Criteria:         []Criterion{deterministicCriterion},
	})
	if stale.Class != ResultConflict || stale.CurrentVersion != 1 {
		t.Fatalf("stale write = %#v", stale)
	}
}

func TestConcurrentDuplicateCommandAppendsOneEvent(t *testing.T) {
	rig := newTestRig(t)
	command := CreateOutcomeCommand{
		Meta:             rig.meta(CommandCreateOutcome, "outcome:concurrent-idempotency", 0, direct(steward)),
		DesiredCondition: "concurrent retries converge on one accepted event",
		Steward:          steward,
		Scope:            "reference",
		Horizon:          "now",
		Criteria:         []Criterion{deterministicCriterion},
	}
	const callers = 32
	results := make(chan Result, callers)
	var wait sync.WaitGroup
	for range callers {
		wait.Add(1)
		go func() {
			defer wait.Done()
			results <- rig.kernel.CreateOutcome(command)
		}()
	}
	wait.Wait()
	close(results)
	var acceptedEvent ID
	for result := range results {
		event := requireAccepted(t, result)
		if acceptedEvent == "" {
			acceptedEvent = event.ID
		}
		if event.ID != acceptedEvent {
			t.Fatalf("concurrent retry returned event %s, want %s", event.ID, acceptedEvent)
		}
	}
	if count := len(rig.kernel.Events()); count != 1 {
		t.Fatalf("concurrent retries appended %d events, want 1", count)
	}
}

func TestFencingTokenAndAuthorityTimeStopStaleWorker(t *testing.T) {
	rig := newTestRig(t)
	outcome := rig.createOutcome("outcome:fencing")
	commitment := rig.createActiveCommitment("commitment:fencing", outcome.ID)
	work := rig.createWork("work:fencing", outcome.ID, commitment.ID)
	grant := rig.issueGrant("grant:fencing", work.ID, rig.clock.Now().Add(3*time.Hour))
	first := rig.acquireLease("lease:old", work.ID, rig.clock.Now().Add(time.Minute))
	rig.clock.Advance(2 * time.Minute)
	second := rig.acquireLease("lease:new", work.ID, rig.clock.Now().Add(time.Hour))
	if second.FencingToken != first.FencingToken+1 {
		t.Fatalf("fencing tokens are not monotonic: first=%d second=%d", first.FencingToken, second.FencingToken)
	}

	stale := rig.startRun("run:stale", work, first, grant)
	if stale.Class != ResultInvalid || stale.Code != "STALE_FENCING_TOKEN" {
		t.Fatalf("stale lease started run: %#v", stale)
	}
	requireAccepted(t, rig.startRun("run:current", work, second, grant))
	if check := rig.kernel.CheckSideEffect("run:current", first.FencingToken, grant.ID); check.Allowed || check.Code != "STALE_FENCING_TOKEN" {
		t.Fatalf("old token authorized side effect: %#v", check)
	}
	rig.clock.Advance(61 * time.Minute)
	if check := rig.kernel.CheckSideEffect("run:current", second.FencingToken, grant.ID); check.Allowed || check.Code != "LEASE_EXPIRED_OR_RELEASED" {
		t.Fatalf("expired authority authorized side effect: %#v", check)
	}
}

func TestRevokedGrantFailsClosedAtToolBoundary(t *testing.T) {
	rig := newTestRig(t)
	outcome := rig.createOutcome("outcome:revoke")
	commitment := rig.createActiveCommitment("commitment:revoke", outcome.ID)
	work := rig.createWork("work:revoke", outcome.ID, commitment.ID)
	grant := rig.issueGrant("grant:revoke", work.ID, rig.clock.Now().Add(2*time.Hour))
	lease := rig.acquireLease("lease:revoke", work.ID, rig.clock.Now().Add(time.Hour))
	requireAccepted(t, rig.startRun("run:revoke", work, lease, grant))
	requireAccepted(t, rig.kernel.RevokeGrant(RevokeGrantCommand{
		Meta:   rig.meta(CommandRevokeGrant, grant.ID, 1, direct(steward)),
		Reason: "operator stopped the capability",
	}))
	if check := rig.kernel.CheckSideEffect("run:revoke", lease.FencingToken, grant.ID); check.Allowed || check.Code != "GRANT_NOT_EFFECTIVE" {
		t.Fatalf("revoked grant authorized side effect: %#v", check)
	}
}

func TestEvidenceMustBeIndependentAndCurrent(t *testing.T) {
	rig := newTestRig(t)
	rig.policy.Allow(workspace, operator, ActionEvidenceSubmit)
	rig.policy.Allow(workspace, verifier, ActionEvidenceSubmit)
	outcome := rig.createOutcome("outcome:evidence")
	commitment := rig.createActiveCommitment("commitment:evidence", outcome.ID)
	work := rig.createWork("work:evidence", outcome.ID, commitment.ID)
	grant := rig.issueGrant("grant:evidence", work.ID, rig.clock.Now().Add(2*time.Hour))
	lease := rig.acquireLease("lease:evidence", work.ID, rig.clock.Now().Add(time.Hour))
	requireAccepted(t, rig.startRun("run:evidence", work, lease, grant))
	run := rig.completeRun("run:evidence")

	notIndependent := rig.submitEvidence("evidence:not-independent", outcome.ID, 1, run, delegated(operator, verifier), "deterministic-test")
	if notIndependent.Class != ResultInvalid || notIndependent.Code != "INDEPENDENT_EVALUATOR_REQUIRED" {
		t.Fatalf("executor-accountable evidence = %#v", notIndependent)
	}
	requireAccepted(t, rig.submitEvidence("evidence:current", outcome.ID, 1, run, direct(verifier), "deterministic-test"))
	requireAccepted(t, rig.kernel.ReviseOutcome(ReviseOutcomeCommand{
		Meta:             rig.meta(CommandReviseOutcome, outcome.ID, 1, direct(steward)),
		DesiredCondition: "Interrupted work can resume safely under revised scope.",
		Steward:          steward,
		Scope:            "revised reference simulator",
		Horizon:          "before pilot",
		Criteria:         []Criterion{deterministicCriterion},
	}))
	stale := rig.kernel.AcceptOutcome(AcceptOutcomeCommand{
		Meta:         rig.meta(CommandAcceptOutcome, outcome.ID, 2, direct(steward)),
		AcceptanceID: "acceptance:stale",
		EvidenceIDs:  []ID{"evidence:current"},
	})
	if stale.Class != ResultInvalid || stale.Code != "STALE_ACCEPTANCE_EVIDENCE" {
		t.Fatalf("stale evidence accepted revised outcome: %#v", stale)
	}
}

func TestPolicyDeniesBeforeObjectDetailsAndReplayDetectsTampering(t *testing.T) {
	rig := newTestRig(t)
	outsider := PrincipalID("https://trail.test/principals/outsider")
	deniedResult := rig.kernel.ReviseOutcome(ReviseOutcomeCommand{
		Meta:             rig.meta(CommandReviseOutcome, "outcome:hidden", 7, direct(outsider)),
		DesiredCondition: "probe",
		Steward:          outsider,
		Scope:            "probe",
		Horizon:          "probe",
		Criteria:         []Criterion{deterministicCriterion},
	})
	if deniedResult.Class != ResultDenied || deniedResult.Code != "POLICY_DENIED" {
		t.Fatalf("unauthorized probe leaked object result: %#v", deniedResult)
	}

	rig.createOutcome("outcome:tamper")
	events := rig.kernel.Events()
	events[0].Payload[0] ^= 1
	if _, err := Rebuild("https://trail.test", rig.clock, NewSequenceIDs("urn:trail:rebuild"), rig.policy, events); err == nil {
		t.Fatalf("tampered journal replay error = %v", err)
	}
}

func TestPolicyAndAggregateLookupAreWorkspaceScoped(t *testing.T) {
	rig := newTestRig(t)
	outcome := rig.createOutcome("outcome:workspace-bound")
	otherWorkspace := ID("https://trail.test/workspaces/other")

	unauthorizedMeta := rig.meta(CommandReviseOutcome, outcome.ID, 1, direct(steward))
	unauthorizedMeta.Workspace = otherWorkspace
	deniedResult := rig.kernel.ReviseOutcome(ReviseOutcomeCommand{
		Meta:             unauthorizedMeta,
		DesiredCondition: "cross-workspace probe",
		Steward:          steward,
		Scope:            "other",
		Horizon:          "now",
		Criteria:         []Criterion{deterministicCriterion},
	})
	if deniedResult.Class != ResultDenied || deniedResult.Code != "POLICY_DENIED" {
		t.Fatalf("workspace-scoped policy result = %#v", deniedResult)
	}

	rig.policy.Allow(otherWorkspace, steward, ActionOutcomeRevise)
	authorizedMeta := rig.meta(CommandReviseOutcome, outcome.ID, 1, direct(steward))
	authorizedMeta.Workspace = otherWorkspace
	notFound := rig.kernel.ReviseOutcome(ReviseOutcomeCommand{
		Meta:             authorizedMeta,
		DesiredCondition: "cross-workspace probe",
		Steward:          steward,
		Scope:            "other",
		Horizon:          "now",
		Criteria:         []Criterion{deterministicCriterion},
	})
	if notFound.Class != ResultInvalid || notFound.Code != "OUTCOME_NOT_FOUND" {
		t.Fatalf("cross-workspace aggregate lookup = %#v", notFound)
	}
	original, _ := rig.kernel.Outcome(outcome.ID)
	if original.Version != 1 || original.DesiredCondition != outcome.DesiredCondition {
		t.Fatalf("cross-workspace command mutated outcome: %#v", original)
	}
}
