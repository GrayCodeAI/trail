# Trail implementation status

**Date:** 2026-09-28  
**Current scope:** ADR 0006 bounded executable domain experiment

## Gate status

| Area | Status | Evidence or next gate |
|---|---|---|
| research and architecture baseline | complete enough for semantic experiments | research corpus, requirements, edge cases, architecture, Protocol v0, ADRs 0001–0007 |
| Phase 0 field validation | open | design-partner interviews, concierge trials, pilot commitment, build/no-build record |
| Phase 1 executable domain | in progress | Go reference kernel, protocol-driven simulator, JSON Schema profile, reproducible canonical journal, portable fixtures, adversarial invariant tests |
| Phase 2 local kernel | not authorized | requires applicable Phase 0/1 evidence; no database, API, authentication, worker, or durable event store yet |
| Phase 3 and later product phases | not started | follow the evidence gates in `PHASED_ROADMAP.md` |
| independent product boundary | accepted | ADR 0007: own repository, deployment, data, identity configuration, API, and release lifecycle under `trail.graycodeai.com` |

## Implemented slice

| Capability | Requirement/invariant | Executable evidence |
|---|---|---|
| command envelope and typed results | accepted/conflict/denied/invalid/deferred are distinct | `internal/domain/model.go`, policy-denial and version-conflict tests |
| accepted event journal | mutations carry actor, accountable principal, policy, version, idempotency, time, payload digest | lifecycle and replay tests; canonical event fixture |
| idempotency | same command returns the original event; changed bytes under the same key fail; concurrent retries append once | `TestIdempotencyAndOptimisticConcurrency`, `TestConcurrentDuplicateCommandAppendsOneEvent` |
| explicit commitments | proposal does not activate work; both parties accept the same terms version | `TestCompleteLifecycleSeparatesExecutionEvidenceAndAcceptance` |
| traced work | work references an active outcome and/or active commitment | lifecycle test |
| least-privilege grants | subject, action, exact resource, purpose, validity, and revocation are checked | lifecycle and revocation tests |
| exclusive claims | authority issues monotonically increasing fencing tokens | `TestFencingTokenAndAuthorityTimeStopStaleWorker` |
| protected effects | run, current lease/token, and effective grant are rechecked at the tool boundary | fencing and revocation tests |
| execution vs acceptance | a completed run leaves the outcome active | complete-lifecycle test |
| evidence | exact subject/run versions, artifact manifest, deterministic method, evaluator authority, independence, validity | evidence test |
| acceptance | steward authority, exact outcome version, every criterion, current passing evidence, policy version | lifecycle and stale-evidence tests |
| replay/tamper detection | projections rebuild from ordered accepted events; altered payload digest fails | lifecycle and tamper tests |
| language-neutral schema | Draft 2020-12 envelope, result, implemented resource, event, and batch schemas | `spec/schema/v0`; schema reference and canonical-batch tests |
| portable event conformance | schema failure, optional/required extensions, delegation, workspace, policy, and digest cases | `spec/fixtures/v0/invalid/event-profile-cases.json` |
| transport-neutral commands | flat protocol envelope decodes and dispatches all 13 implemented command types | `internal/domain/protocol.go`, typed command schemas and examples |
| portable command sequences | duplicate delivery, changed idempotent payload, stale version, delegation, and required-extension cases | `spec/fixtures/v0/invalid/command-sequence-cases.json` |
| protocol-to-journal vertical slice | a self-contained policy, clock, and 11-command lifecycle enters through the wire decoder and reproduces the canonical journal exactly | `safe-recovery-commands.json`, `safe-recovery-events.json`, simulator drift check, replay equality test |
| bounded interleaving exploration | 256 orderings of revocation, claim release, deadline, completion, and protected-effect checks never restore lost authority | `TestProtectedEffectAuthorityNeverReturnsAcrossInterleavings` |

## Known omissions before Phase 1 exit

- principal, identity-binding, workspace, request, offer, situation claim, relation, decision, review, and disclosure resources;
- commitment amendment, release, violation, impossibility, dispute, and supersession transitions;
- the full run state machine, checkpoint records, cancellation races, reconciliation receipts, and budget authority;
- final cross-language canonical JSON/signature rules and capability negotiation;
- portable command-sequence fixtures for remaining stateful edge cases such as expiry, revocation, stale fencing, and stale evidence;
- bitemporal claims, evidence supersession, redaction, erasure, and retention behavior;
- independent parser/validator in a second language;
- formal or independent model checking beyond the bounded exhaustive Go interleaving test;
- paper adapter mappings approved by Rho, Rover, Across, and Trace owners.

## Next implementation order

1. Extend portable command sequences to expiry, revocation, stale fencing, and stale evidence.
2. Add the missing minimum object types without adding transport or persistence.
3. Implement a second fixture parser/validator after the schema stops moving.

Phase 0 field work continues in parallel. Technical conformance cannot replace proof that the workflow solves a user problem.
