# Trail Protocol v0 JSON Schema

These Draft 2020-12 schemas cover the bounded executable profile, not the complete planned Trail domain.

| Schema | Purpose |
|---|---|
| `common.schema.json` | identifiers, time, delegation, policy decisions, and base envelopes |
| `resources.schema.json` | resources currently emitted by the reference kernel |
| `event.schema.json` | implemented event types and their typed payloads |
| `event-batch.schema.json` | canonical journal fixture shape |
| `command-envelope.schema.json` | language-neutral base command envelope |
| `command.schema.json` | implemented command types and their typed payloads |
| `command-batch.schema.json` | portable command-example batch |
| `command-sequence.schema.json` | stateful conformance fixture, policy setup, steps, and expected results |
| `result.schema.json` | accepted, conflict, denied, invalid, deferred, and unavailable results |

Unknown additive fields are allowed intentionally. Implementations must preserve them when round-trip behavior is required. Schema validity alone does not prove authority, ordering, idempotency, digest correctness, delegation continuity, workspace isolation, or domain invariants.

`requiredExtensions` is structurally valid when it contains extension URIs. An implementation must reject or quarantine a message when it does not implement every listed required extension.
