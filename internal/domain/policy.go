package domain

import (
	"crypto/sha256"
	"encoding/hex"
	"sync"
	"time"
)

const (
	ActionOutcomeCreate     Action = "outcome.create"
	ActionOutcomeRevise     Action = "outcome.revise"
	ActionOutcomeAccept     Action = "outcome.accept"
	ActionCommitmentPropose Action = "commitment.propose"
	ActionCommitmentAccept  Action = "commitment.accept"
	ActionWorkCreate        Action = "work.create"
	ActionGrantIssue        Action = "grant.issue"
	ActionGrantRevoke       Action = "grant.revoke"
	ActionLeaseAcquire      Action = "lease.acquire"
	ActionLeaseRelease      Action = "lease.release"
	ActionRunStart          Action = "run.start"
	ActionRunComplete       Action = "run.complete"
	ActionEvidenceSubmit    Action = "evidence.submit"
	ActionExecuteWork       Action = "work.execute"
)

type AuthorizationRequest struct {
	Actor       ActorContext
	Action      Action
	Resource    ID
	Workspace   ID
	AuthorityAt time.Time
}

type PolicyEngine interface {
	Decide(AuthorizationRequest) PolicyDecision
}

// RulePolicy is a deterministic workspace-scoped, deny-by-default policy for the reference model.
// Delegated commands require both the acting and accountable principals to hold
// the action so an agent cannot inherit an operator's authority implicitly.
type RulePolicy struct {
	mu      sync.RWMutex
	id      ID
	version uint64
	allowed map[ID]map[PrincipalID]map[Action]struct{}
}

func NewRulePolicy(id ID, version uint64) *RulePolicy {
	return &RulePolicy{id: id, version: version, allowed: make(map[ID]map[PrincipalID]map[Action]struct{})}
}

func (p *RulePolicy) Allow(workspace ID, principal PrincipalID, actions ...Action) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.allowed[workspace] == nil {
		p.allowed[workspace] = make(map[PrincipalID]map[Action]struct{})
	}
	if p.allowed[workspace][principal] == nil {
		p.allowed[workspace][principal] = make(map[Action]struct{})
	}
	for _, action := range actions {
		p.allowed[workspace][principal][action] = struct{}{}
	}
}

func (p *RulePolicy) Decide(req AuthorizationRequest) PolicyDecision {
	p.mu.RLock()
	defer p.mu.RUnlock()
	allowed := p.has(req.Workspace, req.Actor.Actor, req.Action)
	if req.Actor.Accountable != req.Actor.Actor {
		allowed = allowed && p.has(req.Workspace, req.Actor.Accountable, req.Action)
	}
	effect := PolicyDeny
	reason := "POLICY_DENIED"
	if allowed {
		effect = PolicyAllow
		reason = "POLICY_ALLOWED"
	}
	hash := sha256.Sum256([]byte(string(req.Workspace) + "\x00" + string(req.Actor.Actor) + "\x00" + string(req.Action) + "\x00" + string(req.Resource)))
	return PolicyDecision{
		ID:         ID(string(p.id) + "/decisions/" + hex.EncodeToString(hash[:8])),
		PolicyID:   p.id,
		Version:    p.version,
		Effect:     effect,
		ReasonCode: reason,
	}
}

func (p *RulePolicy) has(workspace ID, principal PrincipalID, action Action) bool {
	_, ok := p.allowed[workspace][principal][action]
	return ok
}
