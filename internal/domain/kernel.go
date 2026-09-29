package domain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"slices"
	"sync"
)

type idempotencyRecord struct {
	digest string
	result Result
}

type Kernel struct {
	mu     sync.Mutex
	home   string
	clock  Clock
	ids    IDGenerator
	policy PolicyEngine

	outcomes      map[ID]Outcome
	commitments   map[ID]Commitment
	work          map[ID]WorkItem
	grants        map[ID]CapabilityGrant
	leases        map[ID]ClaimLease
	runs          map[ID]Run
	evidence      map[ID]Evidence
	acceptances   map[ID]Acceptance
	activeLease   map[ID]ID
	fencingByWork map[ID]uint64
	journal       []Event
	eventIDs      map[ID]struct{}
	idempotency   map[string]idempotencyRecord
}

func NewKernel(home string, clock Clock, ids IDGenerator, policy PolicyEngine) *Kernel {
	return &Kernel{
		home:          home,
		clock:         clock,
		ids:           ids,
		policy:        policy,
		outcomes:      make(map[ID]Outcome),
		commitments:   make(map[ID]Commitment),
		work:          make(map[ID]WorkItem),
		grants:        make(map[ID]CapabilityGrant),
		leases:        make(map[ID]ClaimLease),
		runs:          make(map[ID]Run),
		evidence:      make(map[ID]Evidence),
		acceptances:   make(map[ID]Acceptance),
		activeLease:   make(map[ID]ID),
		fencingByWork: make(map[ID]uint64),
		eventIDs:      make(map[ID]struct{}),
		idempotency:   make(map[string]idempotencyRecord),
	}
}

func (k *Kernel) begin(meta CommandMeta, expectedType string, command any) (string, *Result) {
	digest, err := digestJSON(command)
	if err != nil {
		result := invalid("COMMAND_NOT_SERIALIZABLE", "command could not be serialized")
		return "", &result
	}
	if meta.Workspace != "" && meta.IdempotencyKey != "" {
		if previous, ok := k.idempotency[idempotencyScope(meta)]; ok {
			if previous.digest != digest {
				result := invalid("IDEMPOTENCY_KEY_REUSED", "idempotency key was already used for different command bytes")
				return digest, &result
			}
			result := cloneResult(previous.result)
			return digest, &result
		}
	}
	if code, reason := validateMeta(meta, expectedType); code != "" {
		result := invalid(code, reason)
		k.remember(meta, digest, result)
		return digest, &result
	}
	return digest, nil
}

func validateMeta(meta CommandMeta, expectedType string) (string, string) {
	if meta.SpecVersion != SpecVersion {
		return "UNSUPPORTED_SPEC_VERSION", "command uses an unsupported protocol version"
	}
	if meta.ID == "" || meta.Subject == "" || meta.Workspace == "" || meta.IdempotencyKey == "" || meta.SubmittedAt.IsZero() {
		return "INCOMPLETE_COMMAND_ENVELOPE", "command envelope is missing a required field"
	}
	if meta.Type != expectedType {
		return "COMMAND_TYPE_MISMATCH", "command type does not match the invoked handler"
	}
	if code, reason := validateActorContext(meta.Actor); code != "" {
		return code, reason
	}
	return "", ""
}

func validateActorContext(actor ActorContext) (string, string) {
	if actor.Actor == "" || actor.Accountable == "" {
		return "INVALID_ACTOR", "actor and accountable principal are required"
	}
	if actor.Actor == actor.Accountable {
		if len(actor.Delegation) != 0 {
			return "INVALID_DELEGATION", "a direct actor must not provide a delegation chain"
		}
		return "", ""
	}
	if len(actor.Delegation) == 0 || actor.Delegation[0].From != actor.Accountable || actor.Delegation[len(actor.Delegation)-1].To != actor.Actor {
		return "INVALID_DELEGATION", "delegation must connect the accountable principal to the actor"
	}
	for i, step := range actor.Delegation {
		if step.From == "" || step.To == "" || step.GrantID == "" || step.From == step.To {
			return "INVALID_DELEGATION", "delegation contains an invalid step"
		}
		if i > 0 && actor.Delegation[i-1].To != step.From {
			return "INVALID_DELEGATION", "delegation steps are not contiguous"
		}
	}
	return "", ""
}

func (k *Kernel) authorize(meta CommandMeta, action Action) (PolicyDecision, *Result) {
	if k.policy == nil {
		decision := PolicyDecision{Effect: PolicyDeny, ReasonCode: "POLICY_UNAVAILABLE"}
		result := denied(decision, "POLICY_UNAVAILABLE")
		return decision, &result
	}
	decision := k.policy.Decide(AuthorizationRequest{
		Actor:       meta.Actor,
		Action:      action,
		Resource:    meta.Subject,
		Workspace:   meta.Workspace,
		AuthorityAt: k.clock.Now(),
	})
	switch decision.Effect {
	case PolicyAllow:
		if decision.ID == "" || decision.PolicyID == "" || decision.Version == 0 {
			result := denied(decision, "MALFORMED_POLICY_DECISION")
			return decision, &result
		}
		return decision, nil
	case PolicyDefer:
		result := Result{Class: ResultDeferred, Code: decision.ReasonCode, SafeReason: "authorization requires another decision", Policy: &decision}
		return decision, &result
	default:
		result := denied(decision, decision.ReasonCode)
		return decision, &result
	}
}

func (k *Kernel) checkExpected(meta CommandMeta, current uint64) *Result {
	if meta.ExpectedVersion == current {
		return nil
	}
	result := Result{
		Class:           ResultConflict,
		Code:            "VERSION_CONFLICT",
		SafeReason:      "subject changed after the command's expected version",
		ExpectedVersion: meta.ExpectedVersion,
		CurrentVersion:  current,
	}
	return &result
}

func (k *Kernel) accepted(meta CommandMeta, commandDigest string, eventType string, version uint64, decision PolicyDecision, payload any) Result {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return invalid("EVENT_PAYLOAD_INVALID", "accepted event payload could not be serialized")
	}
	payloadBytes, err = normalizeJSON(payloadBytes)
	if err != nil {
		return invalid("EVENT_PAYLOAD_INVALID", "accepted event payload could not be normalized")
	}
	now := k.clock.Now().UTC()
	event := Event{
		SpecVersion:    SpecVersion,
		ID:             k.ids.Next("event"),
		Type:           eventType,
		Subject:        meta.Subject,
		SubjectVersion: version,
		Home:           k.home,
		Workspace:      meta.Workspace,
		Actor:          meta.Actor.Actor,
		Accountable:    meta.Actor.Accountable,
		Delegation:     slices.Clone(meta.Actor.Delegation),
		RecordedAt:     now,
		ValidTime:      ValidTime{From: now},
		CausedBy:       meta.CausedBy,
		Correlation:    meta.Correlation,
		CommandID:      meta.ID,
		IdempotencyKey: meta.IdempotencyKey,
		CommandDigest:  commandDigest,
		Policy:         decision,
		Payload:        payloadBytes,
		PayloadDigest:  digestBytes(payloadBytes),
	}
	if err := k.apply(event); err != nil {
		return invalid("EVENT_APPLICATION_FAILED", err.Error())
	}
	k.journal = append(k.journal, cloneEvent(event))
	k.eventIDs[event.ID] = struct{}{}
	result := Result{Class: ResultAccepted, Event: ptr(cloneEvent(event))}
	k.remember(meta, commandDigest, result)
	return result
}

func (k *Kernel) remember(meta CommandMeta, digest string, result Result) {
	if meta.Workspace == "" || meta.IdempotencyKey == "" || digest == "" {
		return
	}
	k.idempotency[idempotencyScope(meta)] = idempotencyRecord{digest: digest, result: cloneResult(result)}
}

func (k *Kernel) reject(meta CommandMeta, digest string, result Result) Result {
	k.remember(meta, digest, result)
	return result
}

func idempotencyScope(meta CommandMeta) string {
	return string(meta.Workspace) + "\x00" + meta.IdempotencyKey
}

func invalid(code, reason string) Result {
	return Result{Class: ResultInvalid, Code: code, SafeReason: reason}
}

func denied(decision PolicyDecision, code string) Result {
	if code == "" {
		code = "POLICY_DENIED"
	}
	return Result{Class: ResultDenied, Code: code, SafeReason: "policy did not authorize this command", Policy: &decision}
}

func digestJSON(value any) (string, error) {
	encoded, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return digestBytes(encoded), nil
}

func digestBytes(value []byte) string {
	hash := sha256.Sum256(value)
	return "sha256:" + hex.EncodeToString(hash[:])
}

func formatSequence(value uint64) string {
	return fmt.Sprintf("%012d", value)
}

func ptr[T any](value T) *T { return &value }

func cloneEvent(event Event) Event {
	event.Payload = slices.Clone(event.Payload)
	event.Delegation = slices.Clone(event.Delegation)
	event.RequiredExtensions = slices.Clone(event.RequiredExtensions)
	if event.Extensions != nil {
		extensions := event.Extensions
		event.Extensions = make(map[string]json.RawMessage, len(extensions))
		for key, value := range extensions {
			event.Extensions[key] = slices.Clone(value)
		}
	}
	return event
}

func cloneResult(result Result) Result {
	if result.Event != nil {
		result.Event = ptr(cloneEvent(*result.Event))
	}
	if result.Policy != nil {
		copy := *result.Policy
		result.Policy = &copy
	}
	return result
}

func (k *Kernel) objectVersion(id ID) (uint64, bool) {
	if value, ok := k.outcomes[id]; ok {
		return value.Version, true
	}
	if value, ok := k.commitments[id]; ok {
		return value.Version, true
	}
	if value, ok := k.work[id]; ok {
		return value.Version, true
	}
	if value, ok := k.grants[id]; ok {
		return value.Version, true
	}
	if value, ok := k.leases[id]; ok {
		return value.Version, true
	}
	if value, ok := k.runs[id]; ok {
		return value.Version, true
	}
	if value, ok := k.evidence[id]; ok {
		return value.Version, true
	}
	if _, ok := k.acceptances[id]; ok {
		return 1, true
	}
	return 0, false
}

func (k *Kernel) ensureNewSubject(meta CommandMeta) *Result {
	if meta.ExpectedVersion != 0 {
		return k.checkExpected(meta, 0)
	}
	if version, exists := k.objectVersion(meta.Subject); exists {
		return k.checkExpected(meta, version)
	}
	return nil
}

func (k *Kernel) Events() []Event {
	k.mu.Lock()
	defer k.mu.Unlock()
	events := make([]Event, len(k.journal))
	for i, event := range k.journal {
		events[i] = cloneEvent(event)
	}
	return events
}

func cloneJSON[T any](value T) T {
	encoded, _ := json.Marshal(value)
	var clone T
	_ = json.Unmarshal(encoded, &clone)
	return clone
}

func (k *Kernel) Outcome(id ID) (Outcome, bool) {
	k.mu.Lock()
	defer k.mu.Unlock()
	value, ok := k.outcomes[id]
	return cloneJSON(value), ok
}

func (k *Kernel) Commitment(id ID) (Commitment, bool) {
	k.mu.Lock()
	defer k.mu.Unlock()
	value, ok := k.commitments[id]
	return cloneJSON(value), ok
}

func (k *Kernel) Work(id ID) (WorkItem, bool) {
	k.mu.Lock()
	defer k.mu.Unlock()
	value, ok := k.work[id]
	return cloneJSON(value), ok
}

func (k *Kernel) Grant(id ID) (CapabilityGrant, bool) {
	k.mu.Lock()
	defer k.mu.Unlock()
	value, ok := k.grants[id]
	return cloneJSON(value), ok
}

func (k *Kernel) Lease(id ID) (ClaimLease, bool) {
	k.mu.Lock()
	defer k.mu.Unlock()
	value, ok := k.leases[id]
	return cloneJSON(value), ok
}

func (k *Kernel) Run(id ID) (Run, bool) {
	k.mu.Lock()
	defer k.mu.Unlock()
	value, ok := k.runs[id]
	return cloneJSON(value), ok
}

func (k *Kernel) Evidence(id ID) (Evidence, bool) {
	k.mu.Lock()
	defer k.mu.Unlock()
	value, ok := k.evidence[id]
	return cloneJSON(value), ok
}

func (k *Kernel) Acceptance(id ID) (Acceptance, bool) {
	k.mu.Lock()
	defer k.mu.Unlock()
	value, ok := k.acceptances[id]
	return cloneJSON(value), ok
}

func Rebuild(home string, clock Clock, ids IDGenerator, policy PolicyEngine, events []Event) (*Kernel, error) {
	kernel := NewKernel(home, clock, ids, policy)
	for _, event := range events {
		if event.SpecVersion != SpecVersion {
			return nil, fmt.Errorf("event %s: unsupported spec version", event.ID)
		}
		if event.ID == "" || event.Subject == "" || event.Workspace == "" || event.CommandID == "" || event.IdempotencyKey == "" || event.CommandDigest == "" {
			return nil, fmt.Errorf("event %s: incomplete envelope", event.ID)
		}
		if event.Home != home {
			return nil, fmt.Errorf("event %s: home authority mismatch", event.ID)
		}
		if event.Policy.Effect != PolicyAllow {
			return nil, fmt.Errorf("event %s: accepted event lacks allow decision", event.ID)
		}
		if code, _ := validateActorContext(ActorContext{Actor: event.Actor, Accountable: event.Accountable, Delegation: event.Delegation}); code != "" {
			return nil, fmt.Errorf("event %s: %s", event.ID, code)
		}
		if len(event.RequiredExtensions) > 0 {
			return nil, fmt.Errorf("event %s: unsupported required extension %s", event.ID, event.RequiredExtensions[0])
		}
		normalizedPayload, err := normalizeJSON(event.Payload)
		if err != nil {
			return nil, fmt.Errorf("event %s: payload is not valid JSON", event.ID)
		}
		if event.PayloadDigest != digestBytes(normalizedPayload) {
			return nil, fmt.Errorf("event %s: payload digest mismatch", event.ID)
		}
		event.Payload = normalizedPayload
		if _, duplicate := kernel.eventIDs[event.ID]; duplicate {
			return nil, fmt.Errorf("event %s: duplicate event id", event.ID)
		}
		key := string(event.Workspace) + "\x00" + event.IdempotencyKey
		if prior, exists := kernel.idempotency[key]; exists && prior.digest != event.CommandDigest {
			return nil, fmt.Errorf("event %s: idempotency key has conflicting command digest", event.ID)
		}
		if err := kernel.apply(event); err != nil {
			return nil, fmt.Errorf("event %s: %w", event.ID, err)
		}
		copy := cloneEvent(event)
		kernel.journal = append(kernel.journal, copy)
		kernel.eventIDs[event.ID] = struct{}{}
		kernel.idempotency[key] = idempotencyRecord{digest: event.CommandDigest, result: Result{Class: ResultAccepted, Event: ptr(copy)}}
	}
	return kernel, nil
}

func normalizeJSON(value []byte) ([]byte, error) {
	decoder := json.NewDecoder(bytes.NewReader(value))
	decoder.UseNumber()
	var decoded any
	if err := decoder.Decode(&decoded); err != nil {
		return nil, err
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		if err == nil {
			return nil, errors.New("multiple JSON values")
		}
		return nil, err
	}
	return json.Marshal(decoded)
}

func (k *Kernel) apply(event Event) error {
	if event.SubjectVersion == 0 {
		return errors.New("subject version must be positive")
	}
	current, exists := k.objectVersion(event.Subject)
	if exists && event.SubjectVersion != current+1 {
		return fmt.Errorf("non-sequential subject version: got %d after %d", event.SubjectVersion, current)
	}
	if !exists && event.SubjectVersion != 1 {
		return fmt.Errorf("first subject version must be 1, got %d", event.SubjectVersion)
	}

	switch event.Type {
	case "trail.outcome.created", "trail.outcome.amended":
		var payload struct {
			Outcome Outcome `json:"outcome"`
		}
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return err
		}
		if payload.Outcome.ID != event.Subject || payload.Outcome.Version != event.SubjectVersion || payload.Outcome.Workspace != event.Workspace {
			return errors.New("outcome snapshot does not match event subject")
		}
		k.outcomes[event.Subject] = payload.Outcome
	case "trail.acceptance.recorded":
		var payload struct {
			Outcome    Outcome    `json:"outcome"`
			Acceptance Acceptance `json:"acceptance"`
		}
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return err
		}
		if payload.Outcome.ID != event.Subject || payload.Outcome.Version != event.SubjectVersion || payload.Outcome.Workspace != event.Workspace || payload.Acceptance.ID == "" || payload.Acceptance.Workspace != event.Workspace {
			return errors.New("acceptance snapshot does not match event subject")
		}
		if _, collision := k.acceptances[payload.Acceptance.ID]; collision {
			return errors.New("acceptance id already exists")
		}
		k.outcomes[event.Subject] = payload.Outcome
		k.acceptances[payload.Acceptance.ID] = payload.Acceptance
	case "trail.commitment.proposed", "trail.commitment.party-accepted", "trail.commitment.activated":
		var payload struct {
			Commitment Commitment `json:"commitment"`
		}
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return err
		}
		if payload.Commitment.ID != event.Subject || payload.Commitment.Version != event.SubjectVersion || payload.Commitment.Workspace != event.Workspace {
			return errors.New("commitment snapshot does not match event subject")
		}
		k.commitments[event.Subject] = payload.Commitment
	case "trail.work.ready":
		var payload struct {
			Work WorkItem `json:"work"`
		}
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return err
		}
		if payload.Work.ID != event.Subject || payload.Work.Version != event.SubjectVersion || payload.Work.Workspace != event.Workspace {
			return errors.New("work snapshot does not match event subject")
		}
		k.work[event.Subject] = payload.Work
	case "trail.grant.issued", "trail.grant.revoked":
		var payload struct {
			Grant CapabilityGrant `json:"grant"`
		}
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return err
		}
		if payload.Grant.ID != event.Subject || payload.Grant.Version != event.SubjectVersion || payload.Grant.Workspace != event.Workspace {
			return errors.New("grant snapshot does not match event subject")
		}
		k.grants[event.Subject] = payload.Grant
	case "trail.claim.issued", "trail.claim.released":
		var payload struct {
			Lease ClaimLease `json:"lease"`
		}
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return err
		}
		if payload.Lease.ID != event.Subject || payload.Lease.Version != event.SubjectVersion || payload.Lease.Workspace != event.Workspace {
			return errors.New("lease snapshot does not match event subject")
		}
		if event.Type == "trail.claim.issued" {
			if payload.Lease.FencingToken <= k.fencingByWork[payload.Lease.WorkID] {
				return errors.New("fencing token is not monotonic")
			}
			k.fencingByWork[payload.Lease.WorkID] = payload.Lease.FencingToken
			k.activeLease[payload.Lease.WorkID] = payload.Lease.ID
		} else if k.activeLease[payload.Lease.WorkID] == payload.Lease.ID {
			delete(k.activeLease, payload.Lease.WorkID)
		}
		k.leases[event.Subject] = payload.Lease
	case "trail.run.started", "trail.run.completed":
		var payload struct {
			Run Run `json:"run"`
		}
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return err
		}
		if payload.Run.ID != event.Subject || payload.Run.Version != event.SubjectVersion || payload.Run.Workspace != event.Workspace {
			return errors.New("run snapshot does not match event subject")
		}
		k.runs[event.Subject] = payload.Run
	case "trail.evidence.submitted":
		var payload struct {
			Evidence Evidence `json:"evidence"`
		}
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return err
		}
		if payload.Evidence.ID != event.Subject || payload.Evidence.Version != event.SubjectVersion || payload.Evidence.Workspace != event.Workspace {
			return errors.New("evidence snapshot does not match event subject")
		}
		k.evidence[event.Subject] = payload.Evidence
	default:
		return fmt.Errorf("unknown event type %q", event.Type)
	}
	return nil
}
