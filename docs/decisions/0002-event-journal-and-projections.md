# ADR 0002: Accepted event journal and rebuildable projections

- **Status:** Proposed
- **Date:** 2026-09-28

## Context

Trail needs audit, causality, offline/federated exchange, correction without history loss, and multiple views. A mutable graph or board database cannot by itself show which authorized facts produced current state.

Event sourcing also adds migration, privacy, retention, replay, and operational complexity. The journal must be deliberately small and domain-focused.

## Decision

Persist each accepted domain transition as an immutable event. In the same transaction, persist aggregate state/snapshot, idempotency result, and outbox records. Build graph, board, list, timeline, search, metrics, and audit views as versioned projections.

Do not record every UI gesture, prompt token, log line, or secret as a domain event. Store removable/encrypted content separately from structural event metadata when erasure may be required.

## Consequences

- current state can be rebuilt and explained;
- event schemas and reducer versions require strong compatibility discipline;
- projection lag and rebuild become operational concerns;
- privacy/redaction must be designed rather than assuming append-only means permanent content;
- external systems remain authoritative for their own events and are referenced with provenance.

## Evidence gate

Accept after a prototype proves:

- crash-safe event/state/outbox atomicity;
- deterministic rebuild across schema versions;
- acceptable storage and rebuild cost;
- redaction/retention behavior;
- simpler mutable-state alternatives fail audit or federation requirements.

## Reconsider if

- a bitemporal relational model provides the required audit/export with materially lower complexity;
- domain event evolution cannot be operated safely by the likely team;
- privacy obligations make durable structural history unsuitable for the target market.

