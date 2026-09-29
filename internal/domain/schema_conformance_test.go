package domain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"time"
	"unicode/utf8"
)

// schemaHarness implements only the JSON Schema 2020-12 keywords used by the
// bounded Trail profile. It keeps conformance tests dependency-free; it is not
// a general-purpose JSON Schema implementation.
type schemaHarness struct {
	documents map[string]any
}

func loadSchemaHarness(t *testing.T) *schemaHarness {
	t.Helper()
	directory := filepath.Join("..", "..", "spec", "schema", "v0")
	paths, err := filepath.Glob(filepath.Join(directory, "*.schema.json"))
	if err != nil {
		t.Fatalf("list schemas: %v", err)
	}
	if len(paths) == 0 {
		t.Fatal("no protocol schemas found")
	}
	harness := &schemaHarness{documents: make(map[string]any, len(paths))}
	for _, path := range paths {
		encoded, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read schema %s: %v", path, err)
		}
		var document any
		if err := json.Unmarshal(encoded, &document); err != nil {
			t.Fatalf("decode schema %s: %v", path, err)
		}
		harness.documents[filepath.Base(path)] = document
	}
	return harness
}

func TestProtocolSchemasParseAndResolveLocalReferences(t *testing.T) {
	harness := loadSchemaHarness(t)
	for name, document := range harness.documents {
		walkJSON(document, func(value any) {
			object, ok := value.(map[string]any)
			if !ok {
				return
			}
			ref, ok := object["$ref"].(string)
			if !ok {
				return
			}
			if _, _, err := harness.resolve(name, ref); err != nil {
				t.Errorf("%s ref %q: %v", name, ref, err)
			}
		})
	}
}

func TestCanonicalEventBatchMatchesJSONSchemaProfile(t *testing.T) {
	harness := loadSchemaHarness(t)
	encoded, err := os.ReadFile(filepath.Join("..", "..", "spec", "fixtures", "v0", "valid", "safe-recovery-events.json"))
	if err != nil {
		t.Fatalf("read canonical events: %v", err)
	}
	var document any
	if err := json.Unmarshal(encoded, &document); err != nil {
		t.Fatalf("decode canonical events: %v", err)
	}
	if failures := harness.validate("event-batch.schema.json", document); len(failures) > 0 {
		t.Fatalf("canonical event batch violates schema:\n%s", strings.Join(failures, "\n"))
	}
}

func TestCommandExamplesMatchJSONSchemaAndDecode(t *testing.T) {
	harness := loadSchemaHarness(t)
	encoded, err := os.ReadFile(filepath.Join("..", "..", "spec", "fixtures", "v0", "valid", "command-examples.json"))
	if err != nil {
		t.Fatalf("read command examples: %v", err)
	}
	var document any
	if err := json.Unmarshal(encoded, &document); err != nil {
		t.Fatalf("decode command examples: %v", err)
	}
	if failures := harness.validate("command-batch.schema.json", document); len(failures) > 0 {
		t.Fatalf("command examples violate schema:\n%s", strings.Join(failures, "\n"))
	}
	var commands []json.RawMessage
	if err := json.Unmarshal(encoded, &commands); err != nil {
		t.Fatalf("decode raw command examples: %v", err)
	}
	if len(commands) != 13 {
		t.Fatalf("command example count = %d, want 13", len(commands))
	}
	types := make(map[string]struct{}, len(commands))
	for index, encodedCommand := range commands {
		command, result := DecodeWireCommand(encodedCommand)
		if result.Class != "" {
			t.Fatalf("command %d did not decode: %#v", index, result)
		}
		types[command.Type] = struct{}{}
	}
	if len(types) != len(commands) {
		t.Fatalf("command examples contain duplicate types: %d unique", len(types))
	}
}

type mutationFixture struct {
	Base struct {
		Fixture    string `json:"fixture"`
		EventIndex int    `json:"eventIndex"`
	} `json:"base"`
	Cases []struct {
		Name      string         `json:"name"`
		Mutations []jsonMutation `json:"mutations"`
		Expected  struct {
			SchemaValid   bool   `json:"schemaValid"`
			ReplayValid   bool   `json:"replayValid"`
			ErrorContains string `json:"errorContains"`
		} `json:"expected"`
	} `json:"cases"`
}

type jsonMutation struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value any    `json:"value"`
}

func TestPortableInvalidEventFixtures(t *testing.T) {
	harness := loadSchemaHarness(t)
	fixtureDirectory := filepath.Join("..", "..", "spec", "fixtures", "v0", "invalid")
	encoded, err := os.ReadFile(filepath.Join(fixtureDirectory, "event-profile-cases.json"))
	if err != nil {
		t.Fatalf("read invalid fixture: %v", err)
	}
	var fixture mutationFixture
	if err := json.Unmarshal(encoded, &fixture); err != nil {
		t.Fatalf("decode invalid fixture: %v", err)
	}
	baseEncoded, err := os.ReadFile(filepath.Clean(filepath.Join(fixtureDirectory, fixture.Base.Fixture)))
	if err != nil {
		t.Fatalf("read base fixture: %v", err)
	}
	var baseEvents []any
	if err := json.Unmarshal(baseEncoded, &baseEvents); err != nil {
		t.Fatalf("decode base fixture: %v", err)
	}
	if fixture.Base.EventIndex < 0 || fixture.Base.EventIndex >= len(baseEvents) {
		t.Fatalf("base event index %d out of range", fixture.Base.EventIndex)
	}

	for _, testCase := range fixture.Cases {
		t.Run(testCase.Name, func(t *testing.T) {
			document := cloneAny(baseEvents[fixture.Base.EventIndex])
			for _, mutation := range testCase.Mutations {
				if err := applyJSONMutation(document, mutation); err != nil {
					t.Fatalf("apply mutation: %v", err)
				}
			}
			failures := harness.validate("event.schema.json", document)
			schemaValid := len(failures) == 0
			if schemaValid != testCase.Expected.SchemaValid {
				t.Fatalf("schemaValid=%v want %v; failures=%v", schemaValid, testCase.Expected.SchemaValid, failures)
			}
			if !schemaValid {
				return
			}

			eventBytes, err := json.Marshal(document)
			if err != nil {
				t.Fatalf("encode mutated event: %v", err)
			}
			var event Event
			if err := json.Unmarshal(eventBytes, &event); err != nil {
				t.Fatalf("decode mutated event: %v", err)
			}
			clock := ClockFunc(func() time.Time { return time.Date(2026, 9, 28, 10, 0, 0, 0, time.UTC) })
			rebuilt, replayErr := Rebuild("https://demo.trail.graycodeai.com", clock, NewSequenceIDs("urn:fixture"), nil, []Event{event})
			replayValid := replayErr == nil
			if replayValid != testCase.Expected.ReplayValid {
				t.Fatalf("replayValid=%v want %v; error=%v", replayValid, testCase.Expected.ReplayValid, replayErr)
			}
			if testCase.Expected.ErrorContains != "" && (replayErr == nil || !strings.Contains(replayErr.Error(), testCase.Expected.ErrorContains)) {
				t.Fatalf("replay error %q does not contain %q", replayErr, testCase.Expected.ErrorContains)
			}
			if testCase.Name == "unknown-additive-extension-is-preserved" {
				replayed := rebuilt.Events()
				if len(replayed) != 1 || len(replayed[0].Extensions) != 1 {
					t.Fatalf("optional extension was not preserved: %#v", replayed)
				}
			}
		})
	}
}

type commandSequenceFixture struct {
	SpecVersion   string    `json:"specVersion"`
	Home          string    `json:"home"`
	EventIDPrefix string    `json:"eventIdPrefix"`
	Clock         time.Time `json:"clock"`
	Policy        struct {
		ID      ID     `json:"id"`
		Version uint64 `json:"version"`
		Allow   []struct {
			Workspace ID          `json:"workspace"`
			Principal PrincipalID `json:"principal"`
			Actions   []Action    `json:"actions"`
		} `json:"allow"`
	} `json:"policy"`
	Cases []struct {
		Name  string `json:"name"`
		Steps []struct {
			Command  json.RawMessage `json:"command"`
			Expected struct {
				Class           ResultClass `json:"class"`
				Code            string      `json:"code"`
				EventType       string      `json:"eventType"`
				SameEventAsStep *int        `json:"sameEventAsStep"`
			} `json:"expected"`
		} `json:"steps"`
	} `json:"cases"`
}

func TestPortableCommandSequenceFixtures(t *testing.T) {
	path := filepath.Join("..", "..", "spec", "fixtures", "v0", "invalid", "command-sequence-cases.json")
	fixture := loadCommandSequenceFixture(t, path)
	runCommandSequenceFixture(t, fixture)
}

func TestCanonicalCommandSequenceReplaysCanonicalEventJournal(t *testing.T) {
	fixtureDirectory := filepath.Join("..", "..", "spec", "fixtures", "v0", "valid")
	fixture := loadCommandSequenceFixture(t, filepath.Join(fixtureDirectory, "safe-recovery-commands.json"))
	eventsByCase := runCommandSequenceFixture(t, fixture)
	actual, exists := eventsByCase["safe-recovery"]
	if !exists {
		t.Fatal("canonical command fixture has no safe-recovery case")
	}

	encoded, err := os.ReadFile(filepath.Join(fixtureDirectory, "safe-recovery-events.json"))
	if err != nil {
		t.Fatalf("read canonical event journal: %v", err)
	}
	actualEncoded, err := json.MarshalIndent(actual, "", "  ")
	if err != nil {
		t.Fatalf("encode command-produced journal: %v", err)
	}
	actualEncoded = append(actualEncoded, '\n')
	if !bytes.Equal(actualEncoded, encoded) {
		t.Fatal("command fixture did not reproduce the canonical serialized event journal")
	}
}

func loadCommandSequenceFixture(t *testing.T, path string) commandSequenceFixture {
	t.Helper()
	harness := loadSchemaHarness(t)
	encoded, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read command sequences: %v", err)
	}
	var schemaDocument any
	if err := json.Unmarshal(encoded, &schemaDocument); err != nil {
		t.Fatalf("decode command sequence schema document: %v", err)
	}
	if failures := harness.validate("command-sequence.schema.json", schemaDocument); len(failures) > 0 {
		t.Fatalf("command sequence fixture violates schema:\n%s", strings.Join(failures, "\n"))
	}
	var fixture commandSequenceFixture
	if err := json.Unmarshal(encoded, &fixture); err != nil {
		t.Fatalf("decode command sequence fixture: %v", err)
	}
	return fixture
}

func runCommandSequenceFixture(t *testing.T, fixture commandSequenceFixture) map[string][]Event {
	t.Helper()
	eventsByCase := make(map[string][]Event, len(fixture.Cases))
	for _, testCase := range fixture.Cases {
		t.Run(testCase.Name, func(t *testing.T) {
			clock := ClockFunc(func() time.Time { return fixture.Clock })
			policy := NewRulePolicy(fixture.Policy.ID, fixture.Policy.Version)
			for _, permission := range fixture.Policy.Allow {
				policy.Allow(permission.Workspace, permission.Principal, permission.Actions...)
			}
			kernel := NewKernel(fixture.Home, clock, NewSequenceIDs(fixture.EventIDPrefix), policy)
			results := make([]Result, 0, len(testCase.Steps))
			for index, step := range testCase.Steps {
				result := kernel.ExecuteCommandJSON(step.Command)
				if result.Class != step.Expected.Class || result.Code != step.Expected.Code {
					t.Fatalf("step %d result class/code = %s/%s, want %s/%s", index, result.Class, result.Code, step.Expected.Class, step.Expected.Code)
				}
				if step.Expected.EventType != "" {
					if result.Event == nil || result.Event.Type != step.Expected.EventType {
						t.Fatalf("step %d event type = %#v, want %s", index, result.Event, step.Expected.EventType)
					}
				}
				if step.Expected.SameEventAsStep != nil {
					priorIndex := *step.Expected.SameEventAsStep
					if priorIndex < 0 || priorIndex >= len(results) || result.Event == nil || results[priorIndex].Event == nil || result.Event.ID != results[priorIndex].Event.ID {
						t.Fatalf("step %d did not return the event from step %d", index, priorIndex)
					}
				}
				results = append(results, result)
			}
			eventsByCase[testCase.Name] = kernel.Events()
		})
	}
	return eventsByCase
}

func (h *schemaHarness) validate(schemaName string, instance any) []string {
	schema, ok := h.documents[schemaName]
	if !ok {
		return []string{"unknown schema " + schemaName}
	}
	return h.validateNode(schemaName, schema, instance, "$")
}

func (h *schemaHarness) validateNode(current string, rawSchema, instance any, path string) []string {
	if boolean, ok := rawSchema.(bool); ok {
		if boolean {
			return nil
		}
		return []string{path + ": schema is false"}
	}
	schema, ok := rawSchema.(map[string]any)
	if !ok {
		return []string{path + ": malformed schema node"}
	}
	var failures []string
	if ref, ok := schema["$ref"].(string); ok {
		resolved, resolvedFile, err := h.resolve(current, ref)
		if err != nil {
			failures = append(failures, path+": "+err.Error())
		} else {
			failures = append(failures, h.validateNode(resolvedFile, resolved, instance, path)...)
		}
	}
	if allOf, ok := schema["allOf"].([]any); ok {
		for _, candidate := range allOf {
			failures = append(failures, h.validateNode(current, candidate, instance, path)...)
		}
	}
	if anyOf, ok := schema["anyOf"].([]any); ok {
		matched := false
		for _, candidate := range anyOf {
			if len(h.validateNode(current, candidate, instance, path)) == 0 {
				matched = true
				break
			}
		}
		if !matched {
			failures = append(failures, path+": did not match anyOf")
		}
	}
	if oneOf, ok := schema["oneOf"].([]any); ok {
		matches := 0
		for _, candidate := range oneOf {
			if len(h.validateNode(current, candidate, instance, path)) == 0 {
				matches++
			}
		}
		if matches != 1 {
			failures = append(failures, fmt.Sprintf("%s: matched %d oneOf branches", path, matches))
		}
	}
	if expected, exists := schema["type"]; exists && !matchesSchemaType(instance, expected) {
		failures = append(failures, fmt.Sprintf("%s: type does not match %v", path, expected))
		return failures
	}
	if expected, exists := schema["const"]; exists && !reflect.DeepEqual(instance, expected) {
		failures = append(failures, fmt.Sprintf("%s: value does not match const", path))
	}
	if values, ok := schema["enum"].([]any); ok {
		matched := false
		for _, value := range values {
			if reflect.DeepEqual(instance, value) {
				matched = true
				break
			}
		}
		if !matched {
			failures = append(failures, path+": value is not in enum")
		}
	}

	switch value := instance.(type) {
	case map[string]any:
		if minimum, ok := schema["minProperties"].(float64); ok && len(value) < int(minimum) {
			failures = append(failures, path+": too few properties")
		}
		if required, ok := schema["required"].([]any); ok {
			for _, rawName := range required {
				name, _ := rawName.(string)
				if _, exists := value[name]; !exists {
					failures = append(failures, path+": missing required property "+name)
				}
			}
		}
		properties, _ := schema["properties"].(map[string]any)
		for name, propertySchema := range properties {
			if property, exists := value[name]; exists {
				failures = append(failures, h.validateNode(current, propertySchema, property, path+"/"+escapePointer(name))...)
			}
		}
		if nameSchema, exists := schema["propertyNames"]; exists {
			for name := range value {
				failures = append(failures, h.validateNode(current, nameSchema, name, path+"/{propertyName}")...)
			}
		}
		if additional, exists := schema["additionalProperties"]; exists {
			for name, property := range value {
				if _, known := properties[name]; known {
					continue
				}
				switch rule := additional.(type) {
				case bool:
					if !rule {
						failures = append(failures, path+": additional property "+name)
					}
				case map[string]any:
					failures = append(failures, h.validateNode(current, rule, property, path+"/"+escapePointer(name))...)
				}
			}
		}
	case []any:
		if minimum, ok := schema["minItems"].(float64); ok && len(value) < int(minimum) {
			failures = append(failures, path+": too few items")
		}
		if maximum, ok := schema["maxItems"].(float64); ok && len(value) > int(maximum) {
			failures = append(failures, path+": too many items")
		}
		if unique, _ := schema["uniqueItems"].(bool); unique {
			seen := make(map[string]struct{}, len(value))
			for _, item := range value {
				encoded, _ := json.Marshal(item)
				key := string(encoded)
				if _, duplicate := seen[key]; duplicate {
					failures = append(failures, path+": duplicate array item")
				}
				seen[key] = struct{}{}
			}
		}
		if itemSchema, exists := schema["items"]; exists {
			for index, item := range value {
				failures = append(failures, h.validateNode(current, itemSchema, item, fmt.Sprintf("%s/%d", path, index))...)
			}
		}
	case string:
		if minimum, ok := schema["minLength"].(float64); ok && utf8.RuneCountInString(value) < int(minimum) {
			failures = append(failures, path+": string is too short")
		}
		if pattern, ok := schema["pattern"].(string); ok {
			compiled, err := regexp.Compile(pattern)
			if err != nil {
				failures = append(failures, path+": schema pattern is invalid")
			} else if !compiled.MatchString(value) {
				failures = append(failures, path+": string does not match pattern")
			}
		}
		if format, _ := schema["format"].(string); format == "date-time" {
			if _, err := time.Parse(time.RFC3339Nano, value); err != nil {
				failures = append(failures, path+": invalid date-time")
			}
		}
	case float64:
		if minimum, ok := schema["minimum"].(float64); ok && value < minimum {
			failures = append(failures, path+": number is below minimum")
		}
	}
	return failures
}

func (h *schemaHarness) resolve(current, ref string) (any, string, error) {
	parts := strings.SplitN(ref, "#", 2)
	file := current
	if parts[0] != "" {
		file = filepath.Base(parts[0])
	}
	document, exists := h.documents[file]
	if !exists {
		return nil, "", fmt.Errorf("referenced schema %q not found", file)
	}
	if len(parts) == 1 || parts[1] == "" {
		return document, file, nil
	}
	if !strings.HasPrefix(parts[1], "/") {
		return nil, "", fmt.Errorf("unsupported ref fragment %q", parts[1])
	}
	currentValue := document
	for _, rawToken := range strings.Split(strings.TrimPrefix(parts[1], "/"), "/") {
		token := unescapePointer(rawToken)
		object, ok := currentValue.(map[string]any)
		if !ok {
			return nil, "", fmt.Errorf("ref %q traverses a non-object", ref)
		}
		currentValue, ok = object[token]
		if !ok {
			return nil, "", fmt.Errorf("ref %q token %q not found", ref, token)
		}
	}
	return currentValue, file, nil
}

func matchesSchemaType(value, expected any) bool {
	if values, ok := expected.([]any); ok {
		for _, candidate := range values {
			if matchesSchemaType(value, candidate) {
				return true
			}
		}
		return false
	}
	name, _ := expected.(string)
	switch name {
	case "object":
		_, ok := value.(map[string]any)
		return ok
	case "array":
		_, ok := value.([]any)
		return ok
	case "string":
		_, ok := value.(string)
		return ok
	case "integer":
		number, ok := value.(float64)
		return ok && number == math.Trunc(number)
	case "number":
		_, ok := value.(float64)
		return ok
	case "boolean":
		_, ok := value.(bool)
		return ok
	case "null":
		return value == nil
	default:
		return false
	}
}

func walkJSON(value any, visit func(any)) {
	visit(value)
	switch typed := value.(type) {
	case map[string]any:
		for _, child := range typed {
			walkJSON(child, visit)
		}
	case []any:
		for _, child := range typed {
			walkJSON(child, visit)
		}
	}
}

func cloneAny(value any) any {
	encoded, _ := json.Marshal(value)
	var clone any
	_ = json.Unmarshal(encoded, &clone)
	return clone
}

func applyJSONMutation(document any, mutation jsonMutation) error {
	if !strings.HasPrefix(mutation.Path, "/") {
		return fmt.Errorf("mutation path %q is not a JSON pointer", mutation.Path)
	}
	tokens := strings.Split(strings.TrimPrefix(mutation.Path, "/"), "/")
	current := document
	for _, rawToken := range tokens[:len(tokens)-1] {
		token := unescapePointer(rawToken)
		object, ok := current.(map[string]any)
		if !ok {
			return fmt.Errorf("path %q traverses a non-object", mutation.Path)
		}
		current, ok = object[token]
		if !ok {
			return fmt.Errorf("path %q does not exist", mutation.Path)
		}
	}
	object, ok := current.(map[string]any)
	if !ok {
		return fmt.Errorf("path %q parent is not an object", mutation.Path)
	}
	name := unescapePointer(tokens[len(tokens)-1])
	switch mutation.Op {
	case "add":
		object[name] = mutation.Value
	case "replace":
		if _, exists := object[name]; !exists {
			return fmt.Errorf("replace path %q does not exist", mutation.Path)
		}
		object[name] = mutation.Value
	case "remove":
		if _, exists := object[name]; !exists {
			return fmt.Errorf("remove path %q does not exist", mutation.Path)
		}
		delete(object, name)
	default:
		return fmt.Errorf("unsupported mutation op %q", mutation.Op)
	}
	return nil
}

func escapePointer(value string) string {
	return strings.ReplaceAll(strings.ReplaceAll(value, "~", "~0"), "/", "~1")
}

func unescapePointer(value string) string {
	return strings.ReplaceAll(strings.ReplaceAll(value, "~1", "/"), "~0", "~")
}
