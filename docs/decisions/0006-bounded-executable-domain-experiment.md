# ADR 0006: Start with a bounded executable domain experiment

- **Status:** Accepted
- **Date:** 2026-09-28

## Context

Trail's research and architecture baseline is complete enough to make its central safety claims executable. Phase 0 still lacks direct design-partner interviews, concierge trials, and a recorded build/no-build decision. Building product infrastructure before that evidence would make weak assumptions expensive. Keeping every rule only in prose would also leave concurrency, authority, fencing, evidence, and acceptance semantics untested.

The Phase 1 roadmap already permits an in-memory reference model or simulator before product infrastructure. The user explicitly authorized beginning implementation on 2026-09-28.

## Decision

Implement a dependency-free Go reference kernel that exercises the smallest complete Trail lifecycle:

1. define an outcome and criteria;
2. form a commitment through explicit party acceptance;
3. create work that traces to the outcome or active commitment;
4. issue a scoped, expiring capability;
5. acquire an exclusive lease with a monotonically increasing fencing token;
6. start and complete a bounded run;
7. submit version-bound evidence; and
8. accept an exact outcome version through an authorized policy path.

The experiment includes deterministic clocks and identifiers, an accepted-event journal, idempotent command handling, replayable in-memory projections, adversarial tests, protocol fixtures, and a small simulator.

This decision does not authorize network APIs, database schemas, production authentication, durable storage, user interfaces, execution adapters, federation, deployment manifests, or hosted services. Types under `internal/domain` are deliberately not a stable public Go API. JSON fixtures describe the draft protocol and may change while the protocol identifier remains a draft.

## Consequences

- central invariants become reviewable and testable before storage or UI choices constrain them;
- protocol gaps can be found with cheap, deterministic scenarios;
- the repository gains a Go module and build commands solely for the reference experiment;
- simulator success is technical evidence only and does not satisfy Phase 0 product validation;
- production code may later replace the experiment rather than inherit its internal structure.

## Evidence gate

Keep the experiment only if it exposes useful semantic failures, supports the Phase 1 adversarial fixtures, or materially improves design-partner comprehension. Before expanding beyond this boundary, record the Phase 0 build/no-build result and the next applicable ADR.

## Reconsider if

- users reject explicit commitment, evidence, or acceptance semantics in the selected pilots;
- the model cannot express the pilot without duplicating Rho, Rover, Across, or Trace;
- another implementation language makes independent conformance substantially easier;
- reference code begins driving production storage or API compatibility by accident.
