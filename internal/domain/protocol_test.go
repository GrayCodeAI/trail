package domain

import (
	"encoding/json"
	"testing"
)

func TestWireCommandDispatchesThroughTheSameKernelBoundary(t *testing.T) {
	rig := newTestRig(t)
	payload, err := json.Marshal(struct {
		DesiredCondition string      `json:"desiredCondition"`
		Steward          PrincipalID `json:"steward"`
		Scope            string      `json:"scope"`
		Horizon          string      `json:"horizon"`
		Criteria         []Criterion `json:"criteria"`
	}{
		DesiredCondition: "wire command reaches the domain kernel",
		Steward:          steward,
		Scope:            "protocol test",
		Horizon:          "now",
		Criteria:         []Criterion{deterministicCriterion},
	})
	if err != nil {
		t.Fatal(err)
	}
	wire := WireCommand{
		SpecVersion:     SpecVersion,
		ID:              "urn:trail:test:wire-command",
		Type:            CommandCreateOutcome,
		Subject:         "outcome:wire",
		ExpectedVersion: 0,
		Actor:           steward,
		Accountable:     steward,
		Workspace:       workspace,
		IdempotencyKey:  "wire-command-key",
		SubmittedAt:     rig.clock.Now(),
		Payload:         payload,
	}
	encoded, err := json.Marshal(wire)
	if err != nil {
		t.Fatal(err)
	}
	first := rig.kernel.ExecuteCommandJSON(encoded)
	second := rig.kernel.ExecuteCommandJSON(encoded)
	if !first.Accepted() || !second.Accepted() || first.Event.ID != second.Event.ID {
		t.Fatalf("wire command did not preserve accepted idempotency: first=%#v second=%#v", first, second)
	}
	outcome, ok := rig.kernel.Outcome(wire.Subject)
	if !ok || outcome.DesiredCondition != "wire command reaches the domain kernel" {
		t.Fatalf("wire command projection = %#v", outcome)
	}
}

func TestWireCommandRejectsUnsupportedShapesBeforeDispatch(t *testing.T) {
	rig := newTestRig(t)
	base := WireCommand{
		SpecVersion:     SpecVersion,
		ID:              "urn:trail:test:wire-invalid",
		Type:            CommandCreateOutcome,
		Subject:         "outcome:wire-invalid",
		ExpectedVersion: 0,
		Actor:           steward,
		Accountable:     steward,
		Workspace:       workspace,
		IdempotencyKey:  "wire-invalid-key",
		SubmittedAt:     rig.clock.Now(),
		Payload:         json.RawMessage(`{}`),
	}

	requiredExtension := base
	requiredExtension.RequiredExtensions = []ID{"urn:trail:extension:unknown"}
	encoded, _ := json.Marshal(requiredExtension)
	result := rig.kernel.ExecuteCommandJSON(encoded)
	if result.Class != ResultInvalid || result.Code != "UNSUPPORTED_REQUIRED_EXTENSION" {
		t.Fatalf("required extension result = %#v", result)
	}

	unknown := base
	unknown.Type = "trail.future.command"
	encoded, _ = json.Marshal(unknown)
	result = rig.kernel.ExecuteCommandJSON(encoded)
	if result.Class != ResultInvalid || result.Code != "UNKNOWN_COMMAND_TYPE" {
		t.Fatalf("unknown command result = %#v", result)
	}

	arrayPayload := base
	arrayPayload.Payload = json.RawMessage(`[]`)
	encoded, _ = json.Marshal(arrayPayload)
	result = rig.kernel.ExecuteCommandJSON(encoded)
	if result.Class != ResultInvalid || result.Code != "INVALID_COMMAND_PAYLOAD" {
		t.Fatalf("array payload result = %#v", result)
	}
}
