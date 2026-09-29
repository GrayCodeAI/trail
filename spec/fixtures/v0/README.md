# Trail Protocol v0 fixtures

These fixtures are executable outputs of the bounded Phase 1 reference model. They are conformance inputs for the draft protocol, not a stable storage format.

## Files

- `valid/safe-recovery-events.json`: a complete accepted-event journal from outcome creation through explicit commitment, scoped execution, independent evidence, and acceptance.
- `valid/safe-recovery-commands.json`: the self-contained policy, clock, command envelopes, and expected results that reproduce the canonical safe-recovery journal through the public protocol boundary.
- `valid/command-examples.json`: one schema-valid command for every command type in the bounded executable profile.
- `invalid/event-profile-cases.json`: portable mutations over the first canonical event, covering schema failures, semantic failures, optional extensions, and required-extension rejection.
- `invalid/command-sequence-cases.json`: portable stateful command cases for idempotency, optimistic concurrency, delegation, and extension negotiation.

Regenerate the canonical scenario and journal from the repository root:

```sh
go run ./cmd/trail-sim -mode scenario -output spec/fixtures/v0/valid/safe-recovery-commands.json
go run ./cmd/trail-sim -mode events -output spec/fixtures/v0/valid/safe-recovery-events.json
```

`make check` regenerates both into temporary files and fails if either checked-in fixture has drifted.

Verify the human-readable scenario:

```sh
go run ./cmd/trail-sim
```

Stateful adversarial cases such as reused idempotency keys, expired leases, stale fencing tokens, revocation, and stale evidence still live as executable Go tests in `internal/domain/kernel_test.go`. They will move to portable command-sequence fixtures after the typed command payload schemas stabilize.
