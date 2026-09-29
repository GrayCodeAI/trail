package domain

import (
	"slices"
	"strings"
	"time"
)

const (
	CommandIssueGrant   = "trail.grant.issue"
	CommandRevokeGrant  = "trail.grant.revoke"
	CommandAcquireLease = "trail.claim.acquire"
	CommandReleaseLease = "trail.claim.release"
	CommandStartRun     = "trail.run.request"
	CommandCompleteRun  = "trail.run.report-result"
)

type IssueGrantCommand struct {
	Meta      CommandMeta `json:"meta"`
	Subject   PrincipalID `json:"subject"`
	Actions   []Action    `json:"actions"`
	Resource  ID          `json:"resource"`
	Purpose   string      `json:"purpose"`
	NotBefore time.Time   `json:"notBefore"`
	ExpiresAt time.Time   `json:"expiresAt"`
}

func (k *Kernel) IssueGrant(command IssueGrantCommand) Result {
	k.mu.Lock()
	defer k.mu.Unlock()
	digest, prior := k.begin(command.Meta, CommandIssueGrant, command)
	if prior != nil {
		return *prior
	}
	decision, rejected := k.authorize(command.Meta, ActionGrantIssue)
	if rejected != nil {
		return k.reject(command.Meta, digest, *rejected)
	}
	if conflict := k.ensureNewSubject(command.Meta); conflict != nil {
		return k.reject(command.Meta, digest, *conflict)
	}
	if command.Subject == "" || len(command.Actions) == 0 || command.Resource == "" || strings.TrimSpace(command.Purpose) == "" || command.NotBefore.IsZero() || !command.ExpiresAt.After(command.NotBefore) || !command.ExpiresAt.After(k.clock.Now()) {
		return k.reject(command.Meta, digest, invalid("INVALID_GRANT_SCOPE", "grant requires subject, actions, exact resource, purpose, and a current bounded validity interval"))
	}
	resource, exists := k.work[command.Resource]
	if !exists || resource.Workspace != command.Meta.Workspace {
		return k.reject(command.Meta, digest, invalid("GRANT_RESOURCE_NOT_FOUND", "grant resource does not exist"))
	}
	seen := make(map[Action]struct{}, len(command.Actions))
	for _, action := range command.Actions {
		if action == "" {
			return k.reject(command.Meta, digest, invalid("INVALID_GRANT_ACTION", "grant actions must be explicit"))
		}
		if _, duplicate := seen[action]; duplicate {
			return k.reject(command.Meta, digest, invalid("DUPLICATE_GRANT_ACTION", "grant actions must be unique"))
		}
		seen[action] = struct{}{}
	}
	grant := CapabilityGrant{
		ID:        command.Meta.Subject,
		Workspace: command.Meta.Workspace,
		Version:   1,
		Issuer:    command.Meta.Actor.Actor,
		Subject:   command.Subject,
		Actions:   slices.Clone(command.Actions),
		Resource:  command.Resource,
		Purpose:   strings.TrimSpace(command.Purpose),
		NotBefore: command.NotBefore.UTC(),
		ExpiresAt: command.ExpiresAt.UTC(),
		Status:    GrantActive,
	}
	return k.accepted(command.Meta, digest, "trail.grant.issued", grant.Version, decision, struct {
		Grant CapabilityGrant `json:"grant"`
	}{grant})
}

type RevokeGrantCommand struct {
	Meta   CommandMeta `json:"meta"`
	Reason string      `json:"reason"`
}

func (k *Kernel) RevokeGrant(command RevokeGrantCommand) Result {
	k.mu.Lock()
	defer k.mu.Unlock()
	digest, prior := k.begin(command.Meta, CommandRevokeGrant, command)
	if prior != nil {
		return *prior
	}
	decision, rejected := k.authorize(command.Meta, ActionGrantRevoke)
	if rejected != nil {
		return k.reject(command.Meta, digest, *rejected)
	}
	grant, exists := k.grants[command.Meta.Subject]
	if !exists || grant.Workspace != command.Meta.Workspace {
		return k.reject(command.Meta, digest, invalid("GRANT_NOT_FOUND", "grant does not exist"))
	}
	if conflict := k.checkExpected(command.Meta, grant.Version); conflict != nil {
		return k.reject(command.Meta, digest, *conflict)
	}
	if command.Meta.Actor.Actor != grant.Issuer {
		return k.reject(command.Meta, digest, invalid("GRANT_ISSUER_REQUIRED", "reference profile permits only the issuer to revoke a grant"))
	}
	if grant.Status != GrantActive {
		return k.reject(command.Meta, digest, invalid("GRANT_NOT_ACTIVE", "grant is not active"))
	}
	if strings.TrimSpace(command.Reason) == "" {
		return k.reject(command.Meta, digest, invalid("REVOCATION_REASON_REQUIRED", "grant revocation requires a reason"))
	}
	grant.Version++
	grant.Status = GrantRevoked
	return k.accepted(command.Meta, digest, "trail.grant.revoked", grant.Version, decision, struct {
		Grant  CapabilityGrant `json:"grant"`
		Reason string          `json:"reason"`
	}{grant, strings.TrimSpace(command.Reason)})
}

type AcquireLeaseCommand struct {
	Meta      CommandMeta `json:"meta"`
	WorkID    ID          `json:"workId"`
	Holder    PrincipalID `json:"holder"`
	ExpiresAt time.Time   `json:"expiresAt"`
}

func (k *Kernel) AcquireLease(command AcquireLeaseCommand) Result {
	k.mu.Lock()
	defer k.mu.Unlock()
	digest, prior := k.begin(command.Meta, CommandAcquireLease, command)
	if prior != nil {
		return *prior
	}
	decision, rejected := k.authorize(command.Meta, ActionLeaseAcquire)
	if rejected != nil {
		return k.reject(command.Meta, digest, *rejected)
	}
	if conflict := k.ensureNewSubject(command.Meta); conflict != nil {
		return k.reject(command.Meta, digest, *conflict)
	}
	work, exists := k.work[command.WorkID]
	if !exists || work.Workspace != command.Meta.Workspace {
		return k.reject(command.Meta, digest, invalid("WORK_NOT_FOUND", "lease work does not exist in the workspace"))
	}
	if !work.Exclusive {
		return k.reject(command.Meta, digest, invalid("LEASE_NOT_REQUIRED", "only exclusive work uses a claim lease in this reference profile"))
	}
	if command.Holder == "" || !command.ExpiresAt.After(k.clock.Now()) {
		return k.reject(command.Meta, digest, invalid("INVALID_LEASE", "lease requires a holder and future expiry"))
	}
	if activeID := k.activeLease[command.WorkID]; activeID != "" {
		active := k.leases[activeID]
		if active.ActiveAt(k.clock.Now()) {
			return k.reject(command.Meta, digest, invalid("LEASE_HELD", "work already has an active exclusive lease"))
		}
	}
	token := k.fencingByWork[command.WorkID] + 1
	lease := ClaimLease{
		ID:           command.Meta.Subject,
		Workspace:    command.Meta.Workspace,
		Version:      1,
		WorkID:       command.WorkID,
		Holder:       command.Holder,
		FencingToken: token,
		IssuedAt:     k.clock.Now().UTC(),
		ExpiresAt:    command.ExpiresAt.UTC(),
		Status:       LeaseActive,
	}
	return k.accepted(command.Meta, digest, "trail.claim.issued", lease.Version, decision, struct {
		Lease ClaimLease `json:"lease"`
	}{lease})
}

type ReleaseLeaseCommand struct {
	Meta   CommandMeta `json:"meta"`
	Reason string      `json:"reason"`
}

func (k *Kernel) ReleaseLease(command ReleaseLeaseCommand) Result {
	k.mu.Lock()
	defer k.mu.Unlock()
	digest, prior := k.begin(command.Meta, CommandReleaseLease, command)
	if prior != nil {
		return *prior
	}
	decision, rejected := k.authorize(command.Meta, ActionLeaseRelease)
	if rejected != nil {
		return k.reject(command.Meta, digest, *rejected)
	}
	lease, exists := k.leases[command.Meta.Subject]
	if !exists || lease.Workspace != command.Meta.Workspace {
		return k.reject(command.Meta, digest, invalid("LEASE_NOT_FOUND", "lease does not exist"))
	}
	if conflict := k.checkExpected(command.Meta, lease.Version); conflict != nil {
		return k.reject(command.Meta, digest, *conflict)
	}
	if command.Meta.Actor.Actor != lease.Holder {
		return k.reject(command.Meta, digest, invalid("LEASE_HOLDER_REQUIRED", "reference profile permits only the holder to release a lease"))
	}
	if lease.Status != LeaseActive {
		return k.reject(command.Meta, digest, invalid("LEASE_NOT_ACTIVE", "lease is not active"))
	}
	if strings.TrimSpace(command.Reason) == "" {
		return k.reject(command.Meta, digest, invalid("RELEASE_REASON_REQUIRED", "lease release requires a reason"))
	}
	lease.Version++
	lease.Status = LeaseReleased
	return k.accepted(command.Meta, digest, "trail.claim.released", lease.Version, decision, struct {
		Lease  ClaimLease `json:"lease"`
		Reason string     `json:"reason"`
	}{lease, strings.TrimSpace(command.Reason)})
}

type StartRunCommand struct {
	Meta               CommandMeta   `json:"meta"`
	WorkID             ID            `json:"workId"`
	WorkVersion        uint64        `json:"workVersion"`
	Executor           PrincipalID   `json:"executor"`
	Operator           PrincipalID   `json:"operator"`
	Workload           PrincipalID   `json:"workload"`
	Delegation         []PrincipalID `json:"delegation"`
	LeaseID            ID            `json:"leaseId"`
	FencingToken       uint64        `json:"fencingToken"`
	GrantIDs           []ID          `json:"grantIds"`
	Budget             BudgetLimit   `json:"budget"`
	Deadline           time.Time     `json:"deadline"`
	CheckpointPolicy   string        `json:"checkpointPolicy"`
	Environment        string        `json:"environment"`
	EvidencePolicy     ID            `json:"evidencePolicy"`
	SideEffectStrategy string        `json:"sideEffectStrategy"`
}

func (k *Kernel) StartRun(command StartRunCommand) Result {
	k.mu.Lock()
	defer k.mu.Unlock()
	digest, prior := k.begin(command.Meta, CommandStartRun, command)
	if prior != nil {
		return *prior
	}
	decision, rejected := k.authorize(command.Meta, ActionRunStart)
	if rejected != nil {
		return k.reject(command.Meta, digest, *rejected)
	}
	if conflict := k.ensureNewSubject(command.Meta); conflict != nil {
		return k.reject(command.Meta, digest, *conflict)
	}
	work, exists := k.work[command.WorkID]
	if !exists || work.Workspace != command.Meta.Workspace {
		return k.reject(command.Meta, digest, invalid("WORK_NOT_FOUND", "run work does not exist in the workspace"))
	}
	if command.WorkVersion != work.Version {
		return k.reject(command.Meta, digest, invalid("STALE_WORK_VERSION", "run must bind the current immutable work version"))
	}
	if command.Executor == "" || command.Operator == "" || command.Workload == "" || command.Meta.Actor.Actor != command.Executor || command.Meta.Actor.Accountable != command.Operator {
		return k.reject(command.Meta, digest, invalid("INVALID_RUN_IDENTITY", "run actor, executor, operator, and workload must be explicit and consistent"))
	}
	if len(command.Delegation) != 3 || command.Delegation[0] != command.Operator || command.Delegation[1] != command.Executor || command.Delegation[2] != command.Workload {
		return k.reject(command.Meta, digest, invalid("INVALID_RUN_DELEGATION", "run delegation must bind operator, executor, and workload in order"))
	}
	if command.Budget.Unit == "" || command.Budget.Limit <= 0 || !command.Deadline.After(k.clock.Now()) || strings.TrimSpace(command.CheckpointPolicy) == "" || strings.TrimSpace(command.Environment) == "" || command.EvidencePolicy == "" || strings.TrimSpace(command.SideEffectStrategy) == "" {
		return k.reject(command.Meta, digest, invalid("INCOMPLETE_RUN_ENVELOPE", "run requires a positive budget, deadline, checkpoints, environment, evidence policy, and side-effect strategy"))
	}
	lease, exists := k.leases[command.LeaseID]
	if !exists || lease.WorkID != work.ID || lease.Holder != command.Workload {
		return k.reject(command.Meta, digest, invalid("LEASE_SCOPE_MISMATCH", "lease does not bind this work and workload"))
	}
	if k.activeLease[work.ID] != lease.ID || k.fencingByWork[work.ID] != command.FencingToken || lease.FencingToken != command.FencingToken {
		return k.reject(command.Meta, digest, invalid("STALE_FENCING_TOKEN", "lease fencing token is no longer current"))
	}
	if !lease.ActiveAt(k.clock.Now()) {
		return k.reject(command.Meta, digest, invalid("LEASE_EXPIRED", "lease is not active at authority time"))
	}
	if len(command.GrantIDs) == 0 {
		return k.reject(command.Meta, digest, invalid("RUN_GRANT_REQUIRED", "run requires at least one scoped capability grant"))
	}
	seenGrants := make(map[ID]struct{}, len(command.GrantIDs))
	for _, grantID := range command.GrantIDs {
		if _, duplicate := seenGrants[grantID]; duplicate {
			return k.reject(command.Meta, digest, invalid("DUPLICATE_RUN_GRANT", "run grant ids must be unique"))
		}
		seenGrants[grantID] = struct{}{}
		grant, exists := k.grants[grantID]
		if !exists || grant.Subject != command.Workload || grant.Workspace != command.Meta.Workspace || !grant.Allows(ActionExecuteWork, work.ID, work.Purpose, k.clock.Now()) {
			return k.reject(command.Meta, digest, invalid("GRANT_NOT_EFFECTIVE", "run grant is missing, expired, revoked, or outside the required subject, action, resource, or purpose scope"))
		}
	}
	run := Run{
		ID:                 command.Meta.Subject,
		Workspace:          command.Meta.Workspace,
		Version:            1,
		WorkID:             work.ID,
		WorkVersion:        work.Version,
		Executor:           command.Executor,
		Operator:           command.Operator,
		Workload:           command.Workload,
		Delegation:         slices.Clone(command.Delegation),
		LeaseID:            lease.ID,
		FencingToken:       command.FencingToken,
		GrantIDs:           slices.Clone(command.GrantIDs),
		Budget:             command.Budget,
		Deadline:           command.Deadline.UTC(),
		CheckpointPolicy:   strings.TrimSpace(command.CheckpointPolicy),
		Environment:        strings.TrimSpace(command.Environment),
		EvidencePolicy:     command.EvidencePolicy,
		SideEffectStrategy: strings.TrimSpace(command.SideEffectStrategy),
		Status:             RunActive,
	}
	return k.accepted(command.Meta, digest, "trail.run.started", run.Version, decision, struct {
		Run Run `json:"run"`
	}{run})
}

type CompleteRunCommand struct {
	Meta    CommandMeta `json:"meta"`
	Summary string      `json:"summary"`
}

func (k *Kernel) CompleteRun(command CompleteRunCommand) Result {
	k.mu.Lock()
	defer k.mu.Unlock()
	digest, prior := k.begin(command.Meta, CommandCompleteRun, command)
	if prior != nil {
		return *prior
	}
	decision, rejected := k.authorize(command.Meta, ActionRunComplete)
	if rejected != nil {
		return k.reject(command.Meta, digest, *rejected)
	}
	run, exists := k.runs[command.Meta.Subject]
	if !exists || run.Workspace != command.Meta.Workspace {
		return k.reject(command.Meta, digest, invalid("RUN_NOT_FOUND", "run does not exist"))
	}
	if conflict := k.checkExpected(command.Meta, run.Version); conflict != nil {
		return k.reject(command.Meta, digest, *conflict)
	}
	if command.Meta.Actor.Actor != run.Executor || command.Meta.Actor.Accountable != run.Operator {
		return k.reject(command.Meta, digest, invalid("RUN_EXECUTOR_REQUIRED", "only the recorded executor may complete this run"))
	}
	if run.Status != RunActive {
		return k.reject(command.Meta, digest, invalid("RUN_NOT_ACTIVE", "only an active run may complete"))
	}
	if strings.TrimSpace(command.Summary) == "" {
		return k.reject(command.Meta, digest, invalid("RUN_SUMMARY_REQUIRED", "run completion requires a summary"))
	}
	run.Version++
	run.Status = RunCompleted
	run.Summary = strings.TrimSpace(command.Summary)
	return k.accepted(command.Meta, digest, "trail.run.completed", run.Version, decision, struct {
		Run Run `json:"run"`
	}{run})
}

type SideEffectCheck struct {
	Allowed bool   `json:"allowed"`
	Code    string `json:"code"`
}

// CheckSideEffect must be called at the protected tool boundary immediately
// before a new external effect. It intentionally uses authority time rather
// than a command's submitted time or a worker's wall clock.
func (k *Kernel) CheckSideEffect(runID ID, fencingToken uint64, grantID ID) SideEffectCheck {
	k.mu.Lock()
	defer k.mu.Unlock()
	run, exists := k.runs[runID]
	if !exists || run.Status != RunActive {
		return SideEffectCheck{Code: "RUN_NOT_ACTIVE"}
	}
	lease, exists := k.leases[run.LeaseID]
	if !exists || !lease.ActiveAt(k.clock.Now()) {
		return SideEffectCheck{Code: "LEASE_EXPIRED_OR_RELEASED"}
	}
	if !k.clock.Now().Before(run.Deadline) {
		return SideEffectCheck{Code: "RUN_DEADLINE_EXCEEDED"}
	}
	if fencingToken != run.FencingToken || fencingToken != lease.FencingToken || fencingToken != k.fencingByWork[run.WorkID] || k.activeLease[run.WorkID] != lease.ID {
		return SideEffectCheck{Code: "STALE_FENCING_TOKEN"}
	}
	grant, exists := k.grants[grantID]
	work := k.work[run.WorkID]
	if !exists || !slices.Contains(run.GrantIDs, grantID) || grant.Subject != run.Workload || !grant.Allows(ActionExecuteWork, run.WorkID, work.Purpose, k.clock.Now()) {
		return SideEffectCheck{Code: "GRANT_NOT_EFFECTIVE"}
	}
	return SideEffectCheck{Allowed: true, Code: "AUTHORIZED"}
}
