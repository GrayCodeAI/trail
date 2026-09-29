package domain

import (
	"bytes"
	"encoding/json"
	"io"
	"time"
)

// WireCommand is the transport-neutral command envelope from Trail Protocol v0.
// Transport adapters decode into this type and then invoke ExecuteCommandJSON;
// they never append events or mutate projections directly.
type WireCommand struct {
	SpecVersion        string                     `json:"specVersion"`
	ID                 ID                         `json:"id"`
	Type               string                     `json:"type"`
	Subject            ID                         `json:"subject"`
	ExpectedVersion    uint64                     `json:"expectedVersion"`
	Actor              PrincipalID                `json:"actor"`
	Accountable        PrincipalID                `json:"accountable"`
	Delegation         []DelegationStep           `json:"delegation,omitempty"`
	Workspace          ID                         `json:"workspace"`
	IdempotencyKey     string                     `json:"idempotencyKey"`
	SubmittedAt        time.Time                  `json:"submittedAt"`
	CausedBy           ID                         `json:"causedBy,omitempty"`
	Correlation        ID                         `json:"correlation,omitempty"`
	Extensions         map[string]json.RawMessage `json:"extensions,omitempty"`
	RequiredExtensions []ID                       `json:"requiredExtensions,omitempty"`
	Payload            json.RawMessage            `json:"payload"`
}

func (command WireCommand) meta() CommandMeta {
	return CommandMeta{
		SpecVersion:     command.SpecVersion,
		ID:              command.ID,
		Type:            command.Type,
		Subject:         command.Subject,
		ExpectedVersion: command.ExpectedVersion,
		Actor: ActorContext{
			Actor:       command.Actor,
			Accountable: command.Accountable,
			Delegation:  command.Delegation,
		},
		Workspace:      command.Workspace,
		IdempotencyKey: command.IdempotencyKey,
		SubmittedAt:    command.SubmittedAt,
		CausedBy:       command.CausedBy,
		Correlation:    command.Correlation,
	}
}

func DecodeWireCommand(encoded []byte) (WireCommand, Result) {
	decoder := json.NewDecoder(bytes.NewReader(encoded))
	var command WireCommand
	if err := decoder.Decode(&command); err != nil {
		return WireCommand{}, invalid("MALFORMED_COMMAND_JSON", "command is not valid JSON")
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return WireCommand{}, invalid("MALFORMED_COMMAND_JSON", "command contains more than one JSON value")
	}
	if len(command.RequiredExtensions) > 0 {
		return WireCommand{}, invalid("UNSUPPORTED_REQUIRED_EXTENSION", "command requires an extension this implementation does not support")
	}
	if len(command.Payload) == 0 || !json.Valid(command.Payload) {
		return WireCommand{}, invalid("INVALID_COMMAND_PAYLOAD", "command payload must be a JSON object")
	}
	var object map[string]any
	if err := json.Unmarshal(command.Payload, &object); err != nil || object == nil {
		return WireCommand{}, invalid("INVALID_COMMAND_PAYLOAD", "command payload must be a JSON object")
	}
	return command, Result{}
}

func (k *Kernel) ExecuteCommandJSON(encoded []byte) Result {
	wire, decodeResult := DecodeWireCommand(encoded)
	if decodeResult.Class != "" {
		return decodeResult
	}
	meta := wire.meta()
	switch wire.Type {
	case CommandCreateOutcome:
		var command CreateOutcomeCommand
		if !decodeCommandPayload(wire.Payload, &command) {
			return invalid("INVALID_COMMAND_PAYLOAD", "outcome creation payload is invalid")
		}
		command.Meta = meta
		return k.CreateOutcome(command)
	case CommandReviseOutcome:
		var command ReviseOutcomeCommand
		if !decodeCommandPayload(wire.Payload, &command) {
			return invalid("INVALID_COMMAND_PAYLOAD", "outcome amendment payload is invalid")
		}
		command.Meta = meta
		return k.ReviseOutcome(command)
	case CommandAcceptOutcome:
		var command AcceptOutcomeCommand
		if !decodeCommandPayload(wire.Payload, &command) {
			return invalid("INVALID_COMMAND_PAYLOAD", "acceptance payload is invalid")
		}
		command.Meta = meta
		return k.AcceptOutcome(command)
	case CommandProposeCommitment:
		var command ProposeCommitmentCommand
		if !decodeCommandPayload(wire.Payload, &command) {
			return invalid("INVALID_COMMAND_PAYLOAD", "commitment proposal payload is invalid")
		}
		command.Meta = meta
		return k.ProposeCommitment(command)
	case CommandAcceptCommitment:
		var command AcceptCommitmentCommand
		if !decodeCommandPayload(wire.Payload, &command) {
			return invalid("INVALID_COMMAND_PAYLOAD", "commitment acceptance payload is invalid")
		}
		command.Meta = meta
		return k.AcceptCommitment(command)
	case CommandCreateWork:
		var command CreateWorkCommand
		if !decodeCommandPayload(wire.Payload, &command) {
			return invalid("INVALID_COMMAND_PAYLOAD", "work creation payload is invalid")
		}
		command.Meta = meta
		return k.CreateWork(command)
	case CommandIssueGrant:
		var command IssueGrantCommand
		if !decodeCommandPayload(wire.Payload, &command) {
			return invalid("INVALID_COMMAND_PAYLOAD", "grant issue payload is invalid")
		}
		command.Meta = meta
		return k.IssueGrant(command)
	case CommandRevokeGrant:
		var command RevokeGrantCommand
		if !decodeCommandPayload(wire.Payload, &command) {
			return invalid("INVALID_COMMAND_PAYLOAD", "grant revocation payload is invalid")
		}
		command.Meta = meta
		return k.RevokeGrant(command)
	case CommandAcquireLease:
		var command AcquireLeaseCommand
		if !decodeCommandPayload(wire.Payload, &command) {
			return invalid("INVALID_COMMAND_PAYLOAD", "claim acquisition payload is invalid")
		}
		command.Meta = meta
		return k.AcquireLease(command)
	case CommandReleaseLease:
		var command ReleaseLeaseCommand
		if !decodeCommandPayload(wire.Payload, &command) {
			return invalid("INVALID_COMMAND_PAYLOAD", "claim release payload is invalid")
		}
		command.Meta = meta
		return k.ReleaseLease(command)
	case CommandStartRun:
		var command StartRunCommand
		if !decodeCommandPayload(wire.Payload, &command) {
			return invalid("INVALID_COMMAND_PAYLOAD", "run request payload is invalid")
		}
		command.Meta = meta
		return k.StartRun(command)
	case CommandCompleteRun:
		var command CompleteRunCommand
		if !decodeCommandPayload(wire.Payload, &command) {
			return invalid("INVALID_COMMAND_PAYLOAD", "run result payload is invalid")
		}
		command.Meta = meta
		return k.CompleteRun(command)
	case CommandSubmitEvidence:
		var command SubmitEvidenceCommand
		if !decodeCommandPayload(wire.Payload, &command) {
			return invalid("INVALID_COMMAND_PAYLOAD", "evidence submission payload is invalid")
		}
		command.Meta = meta
		return k.SubmitEvidence(command)
	default:
		return invalid("UNKNOWN_COMMAND_TYPE", "command type is not implemented by this protocol profile")
	}
}

func decodeCommandPayload(payload json.RawMessage, target any) bool {
	return json.Unmarshal(payload, target) == nil
}
