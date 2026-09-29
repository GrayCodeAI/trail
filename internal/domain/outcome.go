package domain

import "strings"

const (
	CommandCreateOutcome = "trail.outcome.create"
	CommandReviseOutcome = "trail.outcome.amend"
	CommandAcceptOutcome = "trail.acceptance.accept"
)

type CreateOutcomeCommand struct {
	Meta             CommandMeta `json:"meta"`
	DesiredCondition string      `json:"desiredCondition"`
	Steward          PrincipalID `json:"steward"`
	Scope            string      `json:"scope"`
	Horizon          string      `json:"horizon"`
	Criteria         []Criterion `json:"criteria"`
}

func (k *Kernel) CreateOutcome(command CreateOutcomeCommand) Result {
	k.mu.Lock()
	defer k.mu.Unlock()
	digest, prior := k.begin(command.Meta, CommandCreateOutcome, command)
	if prior != nil {
		return *prior
	}
	decision, rejected := k.authorize(command.Meta, ActionOutcomeCreate)
	if rejected != nil {
		return k.reject(command.Meta, digest, *rejected)
	}
	if conflict := k.ensureNewSubject(command.Meta); conflict != nil {
		return k.reject(command.Meta, digest, *conflict)
	}
	if code, reason := validateOutcomeFields(command.DesiredCondition, command.Steward, command.Scope, command.Horizon, command.Criteria); code != "" {
		return k.reject(command.Meta, digest, invalid(code, reason))
	}
	outcome := Outcome{
		ID:               command.Meta.Subject,
		Workspace:        command.Meta.Workspace,
		Version:          1,
		DesiredCondition: strings.TrimSpace(command.DesiredCondition),
		Steward:          command.Steward,
		Scope:            strings.TrimSpace(command.Scope),
		Horizon:          strings.TrimSpace(command.Horizon),
		Criteria:         cloneJSON(command.Criteria),
		Status:           OutcomeActive,
	}
	return k.accepted(command.Meta, digest, "trail.outcome.created", outcome.Version, decision, struct {
		Outcome Outcome `json:"outcome"`
	}{outcome})
}

type ReviseOutcomeCommand struct {
	Meta             CommandMeta `json:"meta"`
	DesiredCondition string      `json:"desiredCondition"`
	Steward          PrincipalID `json:"steward"`
	Scope            string      `json:"scope"`
	Horizon          string      `json:"horizon"`
	Criteria         []Criterion `json:"criteria"`
}

func (k *Kernel) ReviseOutcome(command ReviseOutcomeCommand) Result {
	k.mu.Lock()
	defer k.mu.Unlock()
	digest, prior := k.begin(command.Meta, CommandReviseOutcome, command)
	if prior != nil {
		return *prior
	}
	decision, rejected := k.authorize(command.Meta, ActionOutcomeRevise)
	if rejected != nil {
		return k.reject(command.Meta, digest, *rejected)
	}
	current, exists := k.outcomes[command.Meta.Subject]
	if !exists || current.Workspace != command.Meta.Workspace {
		return k.reject(command.Meta, digest, invalid("OUTCOME_NOT_FOUND", "outcome does not exist"))
	}
	if conflict := k.checkExpected(command.Meta, current.Version); conflict != nil {
		return k.reject(command.Meta, digest, *conflict)
	}
	if current.Status == OutcomeAccepted || current.Status == OutcomeCancelled || current.Status == OutcomeSuperseded {
		return k.reject(command.Meta, digest, invalid("OUTCOME_TERMINAL", "terminal outcome cannot be revised"))
	}
	if code, reason := validateOutcomeFields(command.DesiredCondition, command.Steward, command.Scope, command.Horizon, command.Criteria); code != "" {
		return k.reject(command.Meta, digest, invalid(code, reason))
	}
	current.Version++
	current.DesiredCondition = strings.TrimSpace(command.DesiredCondition)
	current.Steward = command.Steward
	current.Scope = strings.TrimSpace(command.Scope)
	current.Horizon = strings.TrimSpace(command.Horizon)
	current.Criteria = cloneJSON(command.Criteria)
	return k.accepted(command.Meta, digest, "trail.outcome.amended", current.Version, decision, struct {
		Outcome Outcome `json:"outcome"`
	}{current})
}

func validateOutcomeFields(desired string, steward PrincipalID, scope, horizon string, criteria []Criterion) (string, string) {
	if strings.TrimSpace(desired) == "" || steward == "" || strings.TrimSpace(scope) == "" || strings.TrimSpace(horizon) == "" {
		return "INCOMPLETE_OUTCOME", "outcome requires desired condition, steward, scope, and horizon"
	}
	if len(criteria) == 0 {
		return "OUTCOME_CRITERIA_REQUIRED", "outcome requires at least one acceptance criterion"
	}
	seen := make(map[ID]struct{}, len(criteria))
	for _, criterion := range criteria {
		if criterion.ID == "" || strings.TrimSpace(criterion.Statement) == "" {
			return "INVALID_CRITERION", "criterion requires an id and statement"
		}
		if _, duplicate := seen[criterion.ID]; duplicate {
			return "DUPLICATE_CRITERION", "criterion ids must be unique within an outcome"
		}
		seen[criterion.ID] = struct{}{}
	}
	return "", ""
}

func criterionByID(outcome Outcome, id ID) (Criterion, bool) {
	for _, criterion := range outcome.Criteria {
		if criterion.ID == id {
			return criterion, true
		}
	}
	return Criterion{}, false
}
