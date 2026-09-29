package domain

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCanonicalSafeRecoveryFixtureReplays(t *testing.T) {
	path := filepath.Join("..", "..", "spec", "fixtures", "v0", "valid", "safe-recovery-events.json")
	encoded, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var events []Event
	if err := json.Unmarshal(encoded, &events); err != nil {
		t.Fatalf("decode fixture: %v", err)
	}
	if len(events) != 11 {
		t.Fatalf("event count = %d, want 11", len(events))
	}
	clock := ClockFunc(func() time.Time {
		return time.Date(2026, 9, 28, 10, 0, 0, 0, time.UTC)
	})
	rebuilt, err := Rebuild("https://demo.trail.graycodeai.com", clock, NewSequenceIDs("urn:fixture"), nil, events)
	if err != nil {
		t.Fatalf("replay fixture: %v", err)
	}
	outcome, ok := rebuilt.Outcome("https://demo.trail.graycodeai.com/outcomes/safe-recovery")
	if !ok || outcome.Status != OutcomeAccepted || outcome.Version != 2 {
		t.Fatalf("fixture outcome = %#v", outcome)
	}
	acceptance, ok := rebuilt.Acceptance("https://demo.trail.graycodeai.com/acceptances/safe-recovery-v1")
	if !ok || acceptance.OutcomeVersion != 1 {
		t.Fatalf("fixture acceptance = %#v", acceptance)
	}

	wrongWorkspace := cloneJSON(events)
	wrongWorkspace[0].Workspace = "https://demo.trail.graycodeai.com/workspaces/other"
	if _, err := Rebuild("https://demo.trail.graycodeai.com", clock, NewSequenceIDs("urn:fixture"), nil, wrongWorkspace); err == nil {
		t.Fatal("replay accepted payload whose workspace differs from its envelope")
	}
}
