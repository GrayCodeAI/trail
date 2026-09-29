package domain

import (
	"strings"
	"time"
)

const (
	CommandProposeCommitment = "trail.commitment.propose"
	CommandAcceptCommitment  = "trail.commitment.accept"
)

type ProposeCommitmentCommand struct {
	Meta             CommandMeta `json:"meta"`
	Debtor           PrincipalID `json:"debtor"`
	Creditor         PrincipalID `json:"creditor"`
	Antecedent       string      `json:"antecedent"`
	Consequent       string      `json:"consequent"`
	OutcomeID        ID          `json:"outcomeId"`
	Deadline         time.Time   `json:"deadline"`
	AcceptancePolicy ID          `json:"acceptancePolicy"`
}

func (k *Kernel) ProposeCommitment(command ProposeCommitmentCommand) Result {
	k.mu.Lock()
	defer k.mu.Unlock()
	digest, prior := k.begin(command.Meta, CommandProposeCommitment, command)
	if prior != nil {
		return *prior
	}
	decision, rejected := k.authorize(command.Meta, ActionCommitmentPropose)
	if rejected != nil {
		return k.reject(command.Meta, digest, *rejected)
	}
	if conflict := k.ensureNewSubject(command.Meta); conflict != nil {
		return k.reject(command.Meta, digest, *conflict)
	}
	if command.Debtor == "" || command.Creditor == "" || command.Debtor == command.Creditor || strings.TrimSpace(command.Antecedent) == "" || strings.TrimSpace(command.Consequent) == "" || command.OutcomeID == "" || command.AcceptancePolicy == "" || !command.Deadline.After(k.clock.Now()) {
		return k.reject(command.Meta, digest, invalid("INVALID_COMMITMENT_TERMS", "commitment requires distinct parties, explicit terms, an outcome, a future deadline, and an acceptance policy"))
	}
	if command.Meta.Actor.Actor != command.Debtor && command.Meta.Actor.Actor != command.Creditor {
		return k.reject(command.Meta, digest, invalid("PROPOSER_NOT_PARTY", "commitment proposer must be a debtor or creditor"))
	}
	outcome, exists := k.outcomes[command.OutcomeID]
	if !exists || outcome.Workspace != command.Meta.Workspace || outcome.Status != OutcomeActive {
		return k.reject(command.Meta, digest, invalid("OUTCOME_NOT_ACTIVE", "commitment must reference an active outcome in the same workspace"))
	}
	commitment := Commitment{
		ID:               command.Meta.Subject,
		Workspace:        command.Meta.Workspace,
		Version:          1,
		TermsVersion:     1,
		Debtor:           command.Debtor,
		Creditor:         command.Creditor,
		Antecedent:       strings.TrimSpace(command.Antecedent),
		Consequent:       strings.TrimSpace(command.Consequent),
		OutcomeID:        command.OutcomeID,
		Deadline:         command.Deadline.UTC(),
		AcceptancePolicy: command.AcceptancePolicy,
		AcceptedTerms:    make(map[PrincipalID]uint64),
		Status:           CommitmentProposed,
	}
	return k.accepted(command.Meta, digest, "trail.commitment.proposed", commitment.Version, decision, struct {
		Commitment Commitment `json:"commitment"`
	}{commitment})
}

type AcceptCommitmentCommand struct {
	Meta         CommandMeta `json:"meta"`
	TermsVersion uint64      `json:"termsVersion"`
}

func (k *Kernel) AcceptCommitment(command AcceptCommitmentCommand) Result {
	k.mu.Lock()
	defer k.mu.Unlock()
	digest, prior := k.begin(command.Meta, CommandAcceptCommitment, command)
	if prior != nil {
		return *prior
	}
	decision, rejected := k.authorize(command.Meta, ActionCommitmentAccept)
	if rejected != nil {
		return k.reject(command.Meta, digest, *rejected)
	}
	commitment, exists := k.commitments[command.Meta.Subject]
	if !exists || commitment.Workspace != command.Meta.Workspace {
		return k.reject(command.Meta, digest, invalid("COMMITMENT_NOT_FOUND", "commitment does not exist"))
	}
	if conflict := k.checkExpected(command.Meta, commitment.Version); conflict != nil {
		return k.reject(command.Meta, digest, *conflict)
	}
	actor := command.Meta.Actor.Actor
	if actor != commitment.Debtor && actor != commitment.Creditor {
		return k.reject(command.Meta, digest, invalid("ACCEPTOR_NOT_PARTY", "only a commitment party may accept its terms"))
	}
	if commitment.Status != CommitmentProposed {
		return k.reject(command.Meta, digest, invalid("COMMITMENT_NOT_PROPOSED", "only a proposed commitment can collect party acceptance"))
	}
	if command.TermsVersion != commitment.TermsVersion {
		return k.reject(command.Meta, digest, invalid("STALE_COMMITMENT_TERMS", "party acceptance must bind the current terms version"))
	}
	if commitment.AcceptedTerms[actor] == command.TermsVersion {
		return k.reject(command.Meta, digest, invalid("TERMS_ALREADY_ACCEPTED", "party already accepted this terms version"))
	}
	commitment.AcceptedTerms[actor] = command.TermsVersion
	commitment.Version++
	eventType := "trail.commitment.party-accepted"
	if commitment.AcceptedTerms[commitment.Debtor] == commitment.TermsVersion && commitment.AcceptedTerms[commitment.Creditor] == commitment.TermsVersion {
		commitment.Status = CommitmentActive
		eventType = "trail.commitment.activated"
	}
	return k.accepted(command.Meta, digest, eventType, commitment.Version, decision, struct {
		Commitment Commitment `json:"commitment"`
	}{commitment})
}
