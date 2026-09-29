package domain

import (
	"encoding/json"
	"sync"
	"time"
)

const SpecVersion = "graycode.trail/0.1-draft"

type ID string
type PrincipalID string
type Action string

type DelegationStep struct {
	From    PrincipalID `json:"from"`
	To      PrincipalID `json:"to"`
	GrantID ID          `json:"grantId"`
}

type ActorContext struct {
	Actor       PrincipalID      `json:"actor"`
	Accountable PrincipalID      `json:"accountable"`
	Delegation  []DelegationStep `json:"delegation,omitempty"`
}

type CommandMeta struct {
	SpecVersion     string       `json:"specVersion"`
	ID              ID           `json:"id"`
	Type            string       `json:"type"`
	Subject         ID           `json:"subject"`
	ExpectedVersion uint64       `json:"expectedVersion"`
	Actor           ActorContext `json:"actor"`
	Workspace       ID           `json:"workspace"`
	IdempotencyKey  string       `json:"idempotencyKey"`
	SubmittedAt     time.Time    `json:"submittedAt"`
	CausedBy        ID           `json:"causedBy,omitempty"`
	Correlation     ID           `json:"correlation,omitempty"`
}

type PolicyEffect string

const (
	PolicyAllow PolicyEffect = "allow"
	PolicyDeny  PolicyEffect = "deny"
	PolicyDefer PolicyEffect = "defer"
)

type PolicyDecision struct {
	ID         ID           `json:"id"`
	PolicyID   ID           `json:"policyId"`
	Version    uint64       `json:"version"`
	Effect     PolicyEffect `json:"effect"`
	ReasonCode string       `json:"reasonCode"`
}

type ValidTime struct {
	From time.Time  `json:"from"`
	To   *time.Time `json:"to"`
}

type Event struct {
	SpecVersion        string                     `json:"specVersion"`
	ID                 ID                         `json:"id"`
	Type               string                     `json:"type"`
	Subject            ID                         `json:"subject"`
	SubjectVersion     uint64                     `json:"subjectVersion"`
	Home               string                     `json:"home"`
	Workspace          ID                         `json:"workspace"`
	Actor              PrincipalID                `json:"actor"`
	Accountable        PrincipalID                `json:"accountable"`
	Delegation         []DelegationStep           `json:"delegation,omitempty"`
	RecordedAt         time.Time                  `json:"recordedAt"`
	ValidTime          ValidTime                  `json:"validTime"`
	CausedBy           ID                         `json:"causedBy,omitempty"`
	Correlation        ID                         `json:"correlation,omitempty"`
	CommandID          ID                         `json:"commandId"`
	IdempotencyKey     string                     `json:"idempotencyKey"`
	CommandDigest      string                     `json:"commandDigest"`
	Policy             PolicyDecision             `json:"policy"`
	Extensions         map[string]json.RawMessage `json:"extensions,omitempty"`
	RequiredExtensions []ID                       `json:"requiredExtensions,omitempty"`
	Payload            json.RawMessage            `json:"payload"`
	PayloadDigest      string                     `json:"payloadDigest"`
}

type ResultClass string

const (
	ResultAccepted    ResultClass = "accepted"
	ResultConflict    ResultClass = "conflict"
	ResultDenied      ResultClass = "denied"
	ResultInvalid     ResultClass = "invalid"
	ResultDeferred    ResultClass = "deferred"
	ResultUnavailable ResultClass = "unavailable"
)

type Result struct {
	Class           ResultClass     `json:"class"`
	Event           *Event          `json:"event,omitempty"`
	Code            string          `json:"code,omitempty"`
	SafeReason      string          `json:"safeReason,omitempty"`
	ExpectedVersion uint64          `json:"expectedVersion,omitempty"`
	CurrentVersion  uint64          `json:"currentVersion,omitempty"`
	Policy          *PolicyDecision `json:"policy,omitempty"`
}

func (r Result) Accepted() bool { return r.Class == ResultAccepted }

type EvidenceRequirement struct {
	Deterministic bool          `json:"deterministic"`
	Independent   bool          `json:"independent"`
	Evaluators    []PrincipalID `json:"evaluators,omitempty"`
}

type Criterion struct {
	ID          ID                  `json:"id"`
	Statement   string              `json:"statement"`
	Requirement EvidenceRequirement `json:"requirement"`
}

type OutcomeStatus string

const (
	OutcomeDraft      OutcomeStatus = "draft"
	OutcomeActive     OutcomeStatus = "active"
	OutcomeAccepted   OutcomeStatus = "accepted"
	OutcomeCancelled  OutcomeStatus = "cancelled"
	OutcomeSuperseded OutcomeStatus = "superseded"
)

type Outcome struct {
	ID               ID            `json:"id"`
	Workspace        ID            `json:"workspace"`
	Version          uint64        `json:"version"`
	DesiredCondition string        `json:"desiredCondition"`
	Steward          PrincipalID   `json:"steward"`
	Scope            string        `json:"scope"`
	Horizon          string        `json:"horizon"`
	Criteria         []Criterion   `json:"criteria"`
	Status           OutcomeStatus `json:"status"`
}

type CommitmentStatus string

const (
	CommitmentProposed  CommitmentStatus = "proposed"
	CommitmentActive    CommitmentStatus = "active"
	CommitmentSatisfied CommitmentStatus = "satisfied"
	CommitmentReleased  CommitmentStatus = "released"
	CommitmentViolated  CommitmentStatus = "violated"
)

type Commitment struct {
	ID               ID                     `json:"id"`
	Workspace        ID                     `json:"workspace"`
	Version          uint64                 `json:"version"`
	TermsVersion     uint64                 `json:"termsVersion"`
	Debtor           PrincipalID            `json:"debtor"`
	Creditor         PrincipalID            `json:"creditor"`
	Antecedent       string                 `json:"antecedent"`
	Consequent       string                 `json:"consequent"`
	OutcomeID        ID                     `json:"outcomeId"`
	Deadline         time.Time              `json:"deadline"`
	AcceptancePolicy ID                     `json:"acceptancePolicy"`
	AcceptedTerms    map[PrincipalID]uint64 `json:"acceptedTerms"`
	Status           CommitmentStatus       `json:"status"`
}

type WorkStatus string

const (
	WorkReady WorkStatus = "ready"
)

type WorkItem struct {
	ID             ID         `json:"id"`
	Workspace      ID         `json:"workspace"`
	Version        uint64     `json:"version"`
	OutcomeID      ID         `json:"outcomeId,omitempty"`
	CommitmentID   ID         `json:"commitmentId,omitempty"`
	Plan           string     `json:"plan"`
	Exclusive      bool       `json:"exclusive"`
	Purpose        string     `json:"purpose"`
	EvidencePolicy ID         `json:"evidencePolicy"`
	Status         WorkStatus `json:"status"`
}

type GrantStatus string

const (
	GrantActive  GrantStatus = "active"
	GrantRevoked GrantStatus = "revoked"
)

type CapabilityGrant struct {
	ID        ID          `json:"id"`
	Workspace ID          `json:"workspace"`
	Version   uint64      `json:"version"`
	Issuer    PrincipalID `json:"issuer"`
	Subject   PrincipalID `json:"subject"`
	Actions   []Action    `json:"actions"`
	Resource  ID          `json:"resource"`
	Purpose   string      `json:"purpose"`
	NotBefore time.Time   `json:"notBefore"`
	ExpiresAt time.Time   `json:"expiresAt"`
	Status    GrantStatus `json:"status"`
}

func (g CapabilityGrant) Allows(action Action, resource ID, purpose string, at time.Time) bool {
	if g.Status != GrantActive || at.Before(g.NotBefore) || !at.Before(g.ExpiresAt) || g.Resource != resource || g.Purpose != purpose {
		return false
	}
	for _, candidate := range g.Actions {
		if candidate == action {
			return true
		}
	}
	return false
}

type LeaseStatus string

const (
	LeaseActive   LeaseStatus = "active"
	LeaseReleased LeaseStatus = "released"
)

type ClaimLease struct {
	ID           ID          `json:"id"`
	Workspace    ID          `json:"workspace"`
	Version      uint64      `json:"version"`
	WorkID       ID          `json:"workId"`
	Holder       PrincipalID `json:"holder"`
	FencingToken uint64      `json:"fencingToken"`
	IssuedAt     time.Time   `json:"issuedAt"`
	ExpiresAt    time.Time   `json:"expiresAt"`
	Status       LeaseStatus `json:"status"`
}

func (l ClaimLease) ActiveAt(at time.Time) bool {
	return l.Status == LeaseActive && !at.Before(l.IssuedAt) && at.Before(l.ExpiresAt)
}

type BudgetLimit struct {
	Unit  string `json:"unit"`
	Limit int64  `json:"limit"`
}

type RunStatus string

const (
	RunActive    RunStatus = "active"
	RunCompleted RunStatus = "completed"
	RunFailed    RunStatus = "failed"
)

type Run struct {
	ID                 ID            `json:"id"`
	Workspace          ID            `json:"workspace"`
	Version            uint64        `json:"version"`
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
	Status             RunStatus     `json:"status"`
	Summary            string        `json:"summary,omitempty"`
}

type Artifact struct {
	Digest     string      `json:"digest"`
	MediaType  string      `json:"mediaType"`
	Size       int64       `json:"size"`
	Locator    string      `json:"locator"`
	Producer   PrincipalID `json:"producer"`
	Provenance string      `json:"provenance"`
}

type EvidenceResult string

const (
	EvidencePass         EvidenceResult = "pass"
	EvidenceFail         EvidenceResult = "fail"
	EvidenceInconclusive EvidenceResult = "inconclusive"
)

type Evidence struct {
	ID                   ID             `json:"id"`
	Workspace            ID             `json:"workspace"`
	Version              uint64         `json:"version"`
	OutcomeID            ID             `json:"outcomeId"`
	OutcomeVersion       uint64         `json:"outcomeVersion"`
	CriterionID          ID             `json:"criterionId"`
	SourceRunID          ID             `json:"sourceRunId"`
	SourceRunVersion     uint64         `json:"sourceRunVersion"`
	Method               string         `json:"method"`
	Result               EvidenceResult `json:"result"`
	Evaluator            PrincipalID    `json:"evaluator"`
	EvaluatorAccountable PrincipalID    `json:"evaluatorAccountable"`
	Artifacts            []Artifact     `json:"artifacts"`
	ValidFrom            time.Time      `json:"validFrom"`
	ValidUntil           time.Time      `json:"validUntil"`
	Limitations          string         `json:"limitations"`
}

type Acceptance struct {
	ID             ID          `json:"id"`
	Workspace      ID          `json:"workspace"`
	OutcomeID      ID          `json:"outcomeId"`
	OutcomeVersion uint64      `json:"outcomeVersion"`
	Evidence       map[ID]ID   `json:"evidence"`
	PolicyID       ID          `json:"policyId"`
	PolicyVersion  uint64      `json:"policyVersion"`
	AcceptedBy     PrincipalID `json:"acceptedBy"`
	AcceptedAt     time.Time   `json:"acceptedAt"`
}

type Clock interface {
	Now() time.Time
}

type ClockFunc func() time.Time

func (f ClockFunc) Now() time.Time { return f() }

type IDGenerator interface {
	Next(kind string) ID
}

type SequenceIDs struct {
	mu     sync.Mutex
	prefix string
	next   uint64
}

func NewSequenceIDs(prefix string) *SequenceIDs {
	return &SequenceIDs{prefix: prefix}
}

func (g *SequenceIDs) Next(kind string) ID {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.next++
	return ID(g.prefix + ":" + kind + ":" + formatSequence(g.next))
}
