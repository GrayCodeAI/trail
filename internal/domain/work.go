package domain

import "strings"

const CommandCreateWork = "trail.work.create"

type CreateWorkCommand struct {
	Meta           CommandMeta `json:"meta"`
	OutcomeID      ID          `json:"outcomeId,omitempty"`
	CommitmentID   ID          `json:"commitmentId,omitempty"`
	Plan           string      `json:"plan"`
	Exclusive      bool        `json:"exclusive"`
	Purpose        string      `json:"purpose"`
	EvidencePolicy ID          `json:"evidencePolicy"`
}

func (k *Kernel) CreateWork(command CreateWorkCommand) Result {
	k.mu.Lock()
	defer k.mu.Unlock()
	digest, prior := k.begin(command.Meta, CommandCreateWork, command)
	if prior != nil {
		return *prior
	}
	decision, rejected := k.authorize(command.Meta, ActionWorkCreate)
	if rejected != nil {
		return k.reject(command.Meta, digest, *rejected)
	}
	if conflict := k.ensureNewSubject(command.Meta); conflict != nil {
		return k.reject(command.Meta, digest, *conflict)
	}
	if command.OutcomeID == "" && command.CommitmentID == "" {
		return k.reject(command.Meta, digest, invalid("WORK_TRACE_REQUIRED", "work must trace to an outcome or commitment"))
	}
	if strings.TrimSpace(command.Plan) == "" || strings.TrimSpace(command.Purpose) == "" || command.EvidencePolicy == "" {
		return k.reject(command.Meta, digest, invalid("INCOMPLETE_WORK", "work requires a plan, purpose, and evidence policy"))
	}
	if command.OutcomeID != "" {
		outcome, exists := k.outcomes[command.OutcomeID]
		if !exists || outcome.Workspace != command.Meta.Workspace || outcome.Status != OutcomeActive {
			return k.reject(command.Meta, digest, invalid("OUTCOME_NOT_ACTIVE", "work outcome must be active in the same workspace"))
		}
	}
	if command.CommitmentID != "" {
		commitment, exists := k.commitments[command.CommitmentID]
		if !exists || commitment.Workspace != command.Meta.Workspace || commitment.Status != CommitmentActive {
			return k.reject(command.Meta, digest, invalid("COMMITMENT_NOT_ACTIVE", "work commitment must be active in the same workspace"))
		}
		if command.OutcomeID != "" && commitment.OutcomeID != command.OutcomeID {
			return k.reject(command.Meta, digest, invalid("WORK_TRACE_MISMATCH", "work outcome and commitment refer to different outcomes"))
		}
	}
	work := WorkItem{
		ID:             command.Meta.Subject,
		Workspace:      command.Meta.Workspace,
		Version:        1,
		OutcomeID:      command.OutcomeID,
		CommitmentID:   command.CommitmentID,
		Plan:           strings.TrimSpace(command.Plan),
		Exclusive:      command.Exclusive,
		Purpose:        strings.TrimSpace(command.Purpose),
		EvidencePolicy: command.EvidencePolicy,
		Status:         WorkReady,
	}
	return k.accepted(command.Meta, digest, "trail.work.ready", work.Version, decision, struct {
		Work WorkItem `json:"work"`
	}{work})
}
