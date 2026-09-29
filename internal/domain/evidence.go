package domain

import (
	"encoding/hex"
	"slices"
	"strings"
	"time"
)

const CommandSubmitEvidence = "trail.evidence.submit"

type SubmitEvidenceCommand struct {
	Meta                 CommandMeta    `json:"meta"`
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

func (k *Kernel) SubmitEvidence(command SubmitEvidenceCommand) Result {
	k.mu.Lock()
	defer k.mu.Unlock()
	digest, prior := k.begin(command.Meta, CommandSubmitEvidence, command)
	if prior != nil {
		return *prior
	}
	decision, rejected := k.authorize(command.Meta, ActionEvidenceSubmit)
	if rejected != nil {
		return k.reject(command.Meta, digest, *rejected)
	}
	if conflict := k.ensureNewSubject(command.Meta); conflict != nil {
		return k.reject(command.Meta, digest, *conflict)
	}
	outcome, exists := k.outcomes[command.OutcomeID]
	if !exists || outcome.Workspace != command.Meta.Workspace || outcome.Status != OutcomeActive {
		return k.reject(command.Meta, digest, invalid("OUTCOME_NOT_ACTIVE", "evidence must reference an active outcome in the same workspace"))
	}
	if command.OutcomeVersion != outcome.Version {
		return k.reject(command.Meta, digest, invalid("STALE_EVIDENCE_SUBJECT", "evidence must bind the current outcome version"))
	}
	criterion, exists := criterionByID(outcome, command.CriterionID)
	if !exists {
		return k.reject(command.Meta, digest, invalid("CRITERION_NOT_FOUND", "evidence criterion does not exist on this outcome version"))
	}
	run, exists := k.runs[command.SourceRunID]
	if !exists || run.Workspace != command.Meta.Workspace || run.Version != command.SourceRunVersion || run.Status != RunCompleted {
		return k.reject(command.Meta, digest, invalid("RUN_NOT_COMPLETED", "evidence must bind an exact completed run version"))
	}
	work := k.work[run.WorkID]
	if work.OutcomeID != outcome.ID {
		return k.reject(command.Meta, digest, invalid("RUN_OUTCOME_MISMATCH", "source run work does not trace to this outcome"))
	}
	if command.Evaluator == "" || command.EvaluatorAccountable == "" || command.Meta.Actor.Actor != command.Evaluator || command.Meta.Actor.Accountable != command.EvaluatorAccountable {
		return k.reject(command.Meta, digest, invalid("INVALID_EVALUATOR_IDENTITY", "evidence actor and accountable evaluator must be explicit and consistent"))
	}
	if len(criterion.Requirement.Evaluators) > 0 && !slices.Contains(criterion.Requirement.Evaluators, command.Evaluator) {
		return k.reject(command.Meta, digest, invalid("EVALUATOR_NOT_AUTHORIZED", "criterion does not authorize this evaluator"))
	}
	method := strings.TrimSpace(command.Method)
	if method == "" || (command.Result != EvidencePass && command.Result != EvidenceFail && command.Result != EvidenceInconclusive) {
		return k.reject(command.Meta, digest, invalid("INVALID_EVIDENCE_RESULT", "evidence requires a method and recognized result"))
	}
	if criterion.Requirement.Deterministic && method != "deterministic-test" {
		return k.reject(command.Meta, digest, invalid("DETERMINISTIC_EVIDENCE_REQUIRED", "criterion requires deterministic test evidence"))
	}
	if criterion.Requirement.Independent && command.EvaluatorAccountable == run.Operator {
		return k.reject(command.Meta, digest, invalid("INDEPENDENT_EVALUATOR_REQUIRED", "criterion requires an evaluator accountable to a different principal than the executor"))
	}
	if (criterion.Requirement.Deterministic || criterion.Requirement.Independent) && method == "agent-self-report" {
		return k.reject(command.Meta, digest, invalid("SELF_REPORT_INSUFFICIENT", "agent self-report cannot satisfy deterministic or independent verification"))
	}
	if command.ValidFrom.IsZero() || command.ValidUntil.IsZero() || command.ValidFrom.After(k.clock.Now()) || !command.ValidUntil.After(k.clock.Now()) || !command.ValidUntil.After(command.ValidFrom) {
		return k.reject(command.Meta, digest, invalid("INVALID_EVIDENCE_VALIDITY", "evidence validity must include current authority time"))
	}
	if len(command.Artifacts) == 0 {
		return k.reject(command.Meta, digest, invalid("EVIDENCE_ARTIFACT_REQUIRED", "evidence requires at least one immutable artifact manifest"))
	}
	for _, artifact := range command.Artifacts {
		if !validSHA256Digest(artifact.Digest) || strings.TrimSpace(artifact.MediaType) == "" || artifact.Size < 0 || strings.TrimSpace(artifact.Locator) == "" || artifact.Producer == "" || strings.TrimSpace(artifact.Provenance) == "" {
			return k.reject(command.Meta, digest, invalid("INVALID_ARTIFACT", "artifact requires sha256 digest, media type, non-negative size, locator, producer, and provenance"))
		}
	}
	evidence := Evidence{
		ID:                   command.Meta.Subject,
		Workspace:            command.Meta.Workspace,
		Version:              1,
		OutcomeID:            outcome.ID,
		OutcomeVersion:       outcome.Version,
		CriterionID:          criterion.ID,
		SourceRunID:          run.ID,
		SourceRunVersion:     run.Version,
		Method:               method,
		Result:               command.Result,
		Evaluator:            command.Evaluator,
		EvaluatorAccountable: command.EvaluatorAccountable,
		Artifacts:            cloneJSON(command.Artifacts),
		ValidFrom:            command.ValidFrom.UTC(),
		ValidUntil:           command.ValidUntil.UTC(),
		Limitations:          strings.TrimSpace(command.Limitations),
	}
	return k.accepted(command.Meta, digest, "trail.evidence.submitted", evidence.Version, decision, struct {
		Evidence Evidence `json:"evidence"`
	}{evidence})
}

type AcceptOutcomeCommand struct {
	Meta         CommandMeta `json:"meta"`
	AcceptanceID ID          `json:"acceptanceId"`
	EvidenceIDs  []ID        `json:"evidenceIds"`
}

func (k *Kernel) AcceptOutcome(command AcceptOutcomeCommand) Result {
	k.mu.Lock()
	defer k.mu.Unlock()
	digest, prior := k.begin(command.Meta, CommandAcceptOutcome, command)
	if prior != nil {
		return *prior
	}
	decision, rejected := k.authorize(command.Meta, ActionOutcomeAccept)
	if rejected != nil {
		return k.reject(command.Meta, digest, *rejected)
	}
	outcome, exists := k.outcomes[command.Meta.Subject]
	if !exists || outcome.Workspace != command.Meta.Workspace {
		return k.reject(command.Meta, digest, invalid("OUTCOME_NOT_FOUND", "outcome does not exist"))
	}
	if conflict := k.checkExpected(command.Meta, outcome.Version); conflict != nil {
		return k.reject(command.Meta, digest, *conflict)
	}
	if outcome.Status != OutcomeActive {
		return k.reject(command.Meta, digest, invalid("OUTCOME_NOT_ACTIVE", "only an active outcome can be accepted"))
	}
	if command.Meta.Actor.Accountable != outcome.Steward {
		return k.reject(command.Meta, digest, invalid("ACCEPTANCE_AUTHORITY_REQUIRED", "outcome steward must be the accountable acceptance authority"))
	}
	if command.AcceptanceID == "" {
		return k.reject(command.Meta, digest, invalid("ACCEPTANCE_ID_REQUIRED", "acceptance requires a stable id"))
	}
	if _, collision := k.objectVersion(command.AcceptanceID); collision {
		return k.reject(command.Meta, digest, invalid("ACCEPTANCE_ID_EXISTS", "acceptance id is already in use"))
	}
	if len(command.EvidenceIDs) == 0 {
		return k.reject(command.Meta, digest, invalid("ACCEPTANCE_EVIDENCE_REQUIRED", "acceptance requires criterion evidence"))
	}
	acceptedEvidence := make(map[ID]ID, len(outcome.Criteria))
	for _, evidenceID := range command.EvidenceIDs {
		evidence, exists := k.evidence[evidenceID]
		if !exists || evidence.Workspace != command.Meta.Workspace {
			return k.reject(command.Meta, digest, invalid("EVIDENCE_NOT_FOUND", "acceptance evidence does not exist in the workspace"))
		}
		if evidence.OutcomeID != outcome.ID || evidence.OutcomeVersion != outcome.Version {
			return k.reject(command.Meta, digest, invalid("STALE_ACCEPTANCE_EVIDENCE", "acceptance evidence does not bind the exact current outcome version"))
		}
		if evidence.Result != EvidencePass || k.clock.Now().Before(evidence.ValidFrom) || !k.clock.Now().Before(evidence.ValidUntil) {
			return k.reject(command.Meta, digest, invalid("EVIDENCE_NOT_ACCEPTABLE", "acceptance evidence must pass and be valid at authority time"))
		}
		if _, duplicate := acceptedEvidence[evidence.CriterionID]; duplicate {
			return k.reject(command.Meta, digest, invalid("DUPLICATE_CRITERION_EVIDENCE", "reference profile accepts one evidence packet per criterion"))
		}
		acceptedEvidence[evidence.CriterionID] = evidence.ID
	}
	for _, criterion := range outcome.Criteria {
		if acceptedEvidence[criterion.ID] == "" {
			return k.reject(command.Meta, digest, invalid("CRITERION_EVIDENCE_MISSING", "every outcome criterion requires acceptable evidence"))
		}
	}
	acceptedVersion := outcome.Version
	acceptance := Acceptance{
		ID:             command.AcceptanceID,
		Workspace:      command.Meta.Workspace,
		OutcomeID:      outcome.ID,
		OutcomeVersion: acceptedVersion,
		Evidence:       acceptedEvidence,
		PolicyID:       decision.PolicyID,
		PolicyVersion:  decision.Version,
		AcceptedBy:     command.Meta.Actor.Actor,
		AcceptedAt:     k.clock.Now().UTC(),
	}
	outcome.Version++
	outcome.Status = OutcomeAccepted
	return k.accepted(command.Meta, digest, "trail.acceptance.recorded", outcome.Version, decision, struct {
		Outcome    Outcome    `json:"outcome"`
		Acceptance Acceptance `json:"acceptance"`
	}{outcome, acceptance})
}

func validSHA256Digest(value string) bool {
	const prefix = "sha256:"
	if !strings.HasPrefix(value, prefix) || len(value) != len(prefix)+64 {
		return false
	}
	_, err := hex.DecodeString(strings.TrimPrefix(value, prefix))
	return err == nil
}
