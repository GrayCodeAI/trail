package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/graycodeai/trail/internal/domain"
)

const (
	workspace domain.ID          = "https://demo.trail.graycodeai.com/workspaces/recovery"
	steward   domain.PrincipalID = "https://demo.trail.graycodeai.com/principals/steward"
	provider  domain.PrincipalID = "https://demo.trail.graycodeai.com/principals/provider"
	operator  domain.PrincipalID = "https://demo.trail.graycodeai.com/principals/operator"
	agent     domain.PrincipalID = "https://demo.trail.graycodeai.com/principals/agent"
	workload  domain.PrincipalID = "https://demo.trail.graycodeai.com/principals/workload"
	verifier  domain.PrincipalID = "https://demo.trail.graycodeai.com/principals/verifier"
)

type fixedClock struct{ now time.Time }

func (c fixedClock) Now() time.Time { return c.now }

type summary struct {
	SpecVersion domain.ID         `json:"specVersion"`
	Outcome     domain.Outcome    `json:"outcome"`
	Commitment  domain.Commitment `json:"commitment"`
	Run         domain.Run        `json:"run"`
	Acceptance  domain.Acceptance `json:"acceptance"`
	EventCount  int               `json:"eventCount"`
	Invariants  map[string]bool   `json:"invariants"`
}

type scenarioPermission struct {
	Workspace domain.ID          `json:"workspace"`
	Principal domain.PrincipalID `json:"principal"`
	Actions   []domain.Action    `json:"actions"`
}

type scenarioPolicy struct {
	ID      domain.ID            `json:"id"`
	Version uint64               `json:"version"`
	Allow   []scenarioPermission `json:"allow"`
}

type scenarioExpected struct {
	Class     domain.ResultClass `json:"class"`
	Code      string             `json:"code,omitempty"`
	EventType string             `json:"eventType,omitempty"`
}

type scenarioStep struct {
	Command  domain.WireCommand `json:"command"`
	Expected scenarioExpected   `json:"expected"`
}

type scenarioCase struct {
	Name  string         `json:"name"`
	Steps []scenarioStep `json:"steps"`
}

type commandScenario struct {
	SpecVersion   string         `json:"specVersion"`
	Home          string         `json:"home"`
	EventIDPrefix string         `json:"eventIdPrefix"`
	Clock         time.Time      `json:"clock"`
	Policy        scenarioPolicy `json:"policy"`
	Cases         []scenarioCase `json:"cases"`
}

func wireCommand(meta domain.CommandMeta, typedCommand any) (domain.WireCommand, error) {
	encoded, err := json.Marshal(typedCommand)
	if err != nil {
		return domain.WireCommand{}, fmt.Errorf("encode typed command: %w", err)
	}
	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(encoded, &fields); err != nil {
		return domain.WireCommand{}, fmt.Errorf("split typed command: %w", err)
	}
	if _, exists := fields["meta"]; !exists {
		return domain.WireCommand{}, fmt.Errorf("typed command has no meta field")
	}
	delete(fields, "meta")
	payload, err := json.Marshal(fields)
	if err != nil {
		return domain.WireCommand{}, fmt.Errorf("encode command payload: %w", err)
	}
	return domain.WireCommand{
		SpecVersion:     meta.SpecVersion,
		ID:              meta.ID,
		Type:            meta.Type,
		Subject:         meta.Subject,
		ExpectedVersion: meta.ExpectedVersion,
		Actor:           meta.Actor.Actor,
		Accountable:     meta.Actor.Accountable,
		Delegation:      meta.Actor.Delegation,
		Workspace:       meta.Workspace,
		IdempotencyKey:  meta.IdempotencyKey,
		SubmittedAt:     meta.SubmittedAt,
		CausedBy:        meta.CausedBy,
		Correlation:     meta.Correlation,
		Payload:         payload,
	}, nil
}

func main() {
	mode := flag.String("mode", "summary", "output mode: summary, events, or scenario")
	output := flag.String("output", "", "optional output file")
	flag.Parse()

	clock := fixedClock{now: time.Date(2026, 9, 28, 10, 0, 0, 0, time.UTC)}
	const home = "https://demo.trail.graycodeai.com"
	const eventIDPrefix = "urn:graycode:trail"
	policyDefinition := scenarioPolicy{
		ID:      "https://demo.trail.graycodeai.com/policies/reference-v1",
		Version: 1,
		Allow: []scenarioPermission{
			{
				Workspace: workspace,
				Principal: steward,
				Actions: []domain.Action{
					domain.ActionOutcomeCreate, domain.ActionOutcomeAccept,
					domain.ActionCommitmentPropose, domain.ActionCommitmentAccept,
					domain.ActionWorkCreate, domain.ActionGrantIssue, domain.ActionLeaseAcquire,
				},
			},
			{Workspace: workspace, Principal: provider, Actions: []domain.Action{domain.ActionCommitmentAccept}},
			{Workspace: workspace, Principal: operator, Actions: []domain.Action{domain.ActionRunStart, domain.ActionRunComplete}},
			{Workspace: workspace, Principal: agent, Actions: []domain.Action{domain.ActionRunStart, domain.ActionRunComplete}},
			{Workspace: workspace, Principal: verifier, Actions: []domain.Action{domain.ActionEvidenceSubmit}},
		},
	}
	policy := domain.NewRulePolicy(policyDefinition.ID, policyDefinition.Version)
	for _, permission := range policyDefinition.Allow {
		policy.Allow(permission.Workspace, permission.Principal, permission.Actions...)
	}

	kernel := domain.NewKernel(home, clock, domain.NewSequenceIDs(eventIDPrefix), policy)
	sequence := uint64(0)
	scenarioSteps := make([]scenarioStep, 0, 11)
	meta := func(commandType string, subject domain.ID, version uint64, actor domain.ActorContext) domain.CommandMeta {
		sequence++
		return domain.CommandMeta{
			SpecVersion:     domain.SpecVersion,
			ID:              domain.ID(fmt.Sprintf("urn:graycode:trail:command:%012d", sequence)),
			Type:            commandType,
			Subject:         subject,
			ExpectedVersion: version,
			Actor:           actor,
			Workspace:       workspace,
			IdempotencyKey:  fmt.Sprintf("demo-%012d", sequence),
			SubmittedAt:     clock.Now(),
			Correlation:     "urn:graycode:trail:scenario:safe-recovery",
		}
	}
	direct := func(principal domain.PrincipalID) domain.ActorContext {
		return domain.ActorContext{Actor: principal, Accountable: principal}
	}
	delegated := func(from, to domain.PrincipalID) domain.ActorContext {
		return domain.ActorContext{
			Actor:       to,
			Accountable: from,
			Delegation: []domain.DelegationStep{{
				From: from, To: to, GrantID: "urn:graycode:trail:delegation:operator-agent",
			}},
		}
	}
	must := func(result domain.Result) domain.Event {
		if !result.Accepted() || result.Event == nil {
			fmt.Fprintf(os.Stderr, "scenario rejected: class=%s code=%s reason=%s\n", result.Class, result.Code, result.SafeReason)
			os.Exit(1)
		}
		return *result.Event
	}
	execute := func(meta domain.CommandMeta, typedCommand any) domain.Result {
		wire, err := wireCommand(meta, typedCommand)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		encoded, err := json.Marshal(wire)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		result := kernel.ExecuteCommandJSON(encoded)
		expected := scenarioExpected{Class: result.Class, Code: result.Code}
		if result.Event != nil {
			expected.EventType = result.Event.Type
		}
		scenarioSteps = append(scenarioSteps, scenarioStep{Command: wire, Expected: expected})
		return result
	}

	outcomeID := domain.ID("https://demo.trail.graycodeai.com/outcomes/safe-recovery")
	criterionID := domain.ID("criterion:resume-without-duplicate-effect")
	createOutcome := domain.CreateOutcomeCommand{
		Meta:             meta(domain.CommandCreateOutcome, outcomeID, 0, direct(steward)),
		DesiredCondition: "An interrupted agent run resumes without a duplicate protected side effect.",
		Steward:          steward,
		Scope:            "reference fault-injection scenario",
		Horizon:          "before the first bounded-execution pilot",
		Criteria: []domain.Criterion{{
			ID:        criterionID,
			Statement: "Fault injection observes exactly one protected side effect.",
			Requirement: domain.EvidenceRequirement{
				Deterministic: true,
				Independent:   true,
				Evaluators:    []domain.PrincipalID{verifier},
			},
		}},
	}
	must(execute(createOutcome.Meta, createOutcome))

	commitmentID := domain.ID("https://demo.trail.graycodeai.com/commitments/safe-recovery")
	proposeCommitment := domain.ProposeCommitmentCommand{
		Meta:             meta(domain.CommandProposeCommitment, commitmentID, 0, direct(steward)),
		Debtor:           provider,
		Creditor:         steward,
		Antecedent:       "the safe-recovery outcome remains active",
		Consequent:       "provide a verified recovery implementation",
		OutcomeID:        outcomeID,
		Deadline:         clock.Now().Add(24 * time.Hour),
		AcceptancePolicy: "https://demo.trail.graycodeai.com/policies/recovery-acceptance-v1",
	}
	must(execute(proposeCommitment.Meta, proposeCommitment))
	acceptCommitmentAsSteward := domain.AcceptCommitmentCommand{
		Meta:         meta(domain.CommandAcceptCommitment, commitmentID, 1, direct(steward)),
		TermsVersion: 1,
	}
	must(execute(acceptCommitmentAsSteward.Meta, acceptCommitmentAsSteward))
	acceptCommitmentAsProvider := domain.AcceptCommitmentCommand{
		Meta:         meta(domain.CommandAcceptCommitment, commitmentID, 2, direct(provider)),
		TermsVersion: 1,
	}
	must(execute(acceptCommitmentAsProvider.Meta, acceptCommitmentAsProvider))

	workID := domain.ID("https://demo.trail.graycodeai.com/work/safe-recovery-test")
	createWork := domain.CreateWorkCommand{
		Meta:           meta(domain.CommandCreateWork, workID, 0, direct(steward)),
		OutcomeID:      outcomeID,
		CommitmentID:   commitmentID,
		Plan:           "inject interruption and reconcile the protected effect",
		Exclusive:      true,
		Purpose:        "safe-recovery-test",
		EvidencePolicy: "https://demo.trail.graycodeai.com/evidence-policies/fault-injection-v1",
	}
	must(execute(createWork.Meta, createWork))

	grantID := domain.ID("https://demo.trail.graycodeai.com/grants/safe-recovery")
	issueGrant := domain.IssueGrantCommand{
		Meta:      meta(domain.CommandIssueGrant, grantID, 0, direct(steward)),
		Subject:   workload,
		Actions:   []domain.Action{domain.ActionExecuteWork},
		Resource:  workID,
		Purpose:   "safe-recovery-test",
		NotBefore: clock.Now(),
		ExpiresAt: clock.Now().Add(2 * time.Hour),
	}
	must(execute(issueGrant.Meta, issueGrant))

	leaseID := domain.ID("https://demo.trail.graycodeai.com/leases/safe-recovery")
	acquireLease := domain.AcquireLeaseCommand{
		Meta:      meta(domain.CommandAcquireLease, leaseID, 0, direct(steward)),
		WorkID:    workID,
		Holder:    workload,
		ExpiresAt: clock.Now().Add(time.Hour),
	}
	must(execute(acquireLease.Meta, acquireLease))
	lease, _ := kernel.Lease(leaseID)
	work, _ := kernel.Work(workID)

	runID := domain.ID("https://demo.trail.graycodeai.com/runs/safe-recovery")
	startRun := domain.StartRunCommand{
		Meta:               meta(domain.CommandStartRun, runID, 0, delegated(operator, agent)),
		WorkID:             workID,
		WorkVersion:        work.Version,
		Executor:           agent,
		Operator:           operator,
		Workload:           workload,
		Delegation:         []domain.PrincipalID{operator, agent, workload},
		LeaseID:            leaseID,
		FencingToken:       lease.FencingToken,
		GrantIDs:           []domain.ID{grantID},
		Budget:             domain.BudgetLimit{Unit: "tool-call", Limit: 20},
		Deadline:           clock.Now().Add(45 * time.Minute),
		CheckpointPolicy:   "after-each-protected-effect",
		Environment:        "rover-sandbox:reference",
		EvidencePolicy:     work.EvidencePolicy,
		SideEffectStrategy: "idempotency-key-and-reconcile",
	}
	must(execute(startRun.Meta, startRun))
	toolCheck := kernel.CheckSideEffect(runID, lease.FencingToken, grantID)
	if !toolCheck.Allowed {
		fmt.Fprintf(os.Stderr, "tool boundary denied: %s\n", toolCheck.Code)
		os.Exit(1)
	}
	completeRun := domain.CompleteRunCommand{
		Meta:    meta(domain.CommandCompleteRun, runID, 1, delegated(operator, agent)),
		Summary: "interruption recovered with exactly one protected side effect",
	}
	must(execute(completeRun.Meta, completeRun))
	run, _ := kernel.Run(runID)

	evidenceID := domain.ID("https://demo.trail.graycodeai.com/evidence/safe-recovery")
	submitEvidence := domain.SubmitEvidenceCommand{
		Meta:                 meta(domain.CommandSubmitEvidence, evidenceID, 0, direct(verifier)),
		OutcomeID:            outcomeID,
		OutcomeVersion:       1,
		CriterionID:          criterionID,
		SourceRunID:          runID,
		SourceRunVersion:     run.Version,
		Method:               "deterministic-test",
		Result:               domain.EvidencePass,
		Evaluator:            verifier,
		EvaluatorAccountable: verifier,
		Artifacts: []domain.Artifact{{
			Digest:     "sha256:" + strings.Repeat("a", 64),
			MediaType:  "application/json",
			Size:       128,
			Locator:    "artifact://safe-recovery/fault-report.json",
			Producer:   workload,
			Provenance: string(runID),
		}},
		ValidFrom:   clock.Now(),
		ValidUntil:  clock.Now().Add(time.Hour),
		Limitations: "deterministic reference simulation; no production runtime",
	}
	must(execute(submitEvidence.Meta, submitEvidence))

	acceptanceID := domain.ID("https://demo.trail.graycodeai.com/acceptances/safe-recovery-v1")
	acceptOutcome := domain.AcceptOutcomeCommand{
		Meta:         meta(domain.CommandAcceptOutcome, outcomeID, 1, direct(steward)),
		AcceptanceID: acceptanceID,
		EvidenceIDs:  []domain.ID{evidenceID},
	}
	must(execute(acceptOutcome.Meta, acceptOutcome))

	events := kernel.Events()
	var value any
	switch *mode {
	case "events":
		value = events
	case "scenario":
		value = commandScenario{
			SpecVersion:   domain.SpecVersion,
			Home:          home,
			EventIDPrefix: eventIDPrefix,
			Clock:         clock.Now(),
			Policy:        policyDefinition,
			Cases: []scenarioCase{{
				Name:  "safe-recovery",
				Steps: scenarioSteps,
			}},
		}
	case "summary":
		outcome, _ := kernel.Outcome(outcomeID)
		commitment, _ := kernel.Commitment(commitmentID)
		acceptance, _ := kernel.Acceptance(acceptanceID)
		_, replayError := domain.Rebuild(home, clock, domain.NewSequenceIDs("urn:graycode:trail:replay"), policy, events)
		value = summary{
			SpecVersion: domain.ID(domain.SpecVersion),
			Outcome:     outcome,
			Commitment:  commitment,
			Run:         run,
			Acceptance:  acceptance,
			EventCount:  len(events),
			Invariants: map[string]bool{
				"commitmentExplicitlyAcceptedByBothParties": len(commitment.AcceptedTerms) == 2,
				"completionDidNotSelfAccept":                run.Status == domain.RunCompleted && acceptance.AcceptedBy != run.Executor,
				"evidenceBoundExactOutcomeVersion":          acceptance.OutcomeVersion == 1,
				"protectedEffectCheckedAtToolBoundary":      toolCheck.Allowed,
				"journalRebuildSucceeded":                   replayError == nil,
			},
		}
	default:
		fmt.Fprintf(os.Stderr, "unknown mode %q\n", *mode)
		os.Exit(2)
	}

	encoded, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	encoded = append(encoded, '\n')
	if *output == "" {
		_, _ = os.Stdout.Write(encoded)
		return
	}
	if err := os.WriteFile(*output, encoded, 0o644); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
