# ADR 0005: Begin as a modular monolith

- **Status:** Proposed
- **Date:** 2026-09-28

## Context

Trail's domain boundaries and load are not yet validated. Identity, policy, commitments, leases, budgets, events, and outbox records benefit from a small transaction boundary. Premature services would create distributed failure before federation itself is ready.

## Decision

Begin with one API/kernel deployable, one worker deployable, one web application, one primary relational database, and one artifact-store abstraction. Enforce module boundaries in code and schemas. Keep adapters and federation at explicit ports.

Extract services only when measured scaling, isolation, deployment cadence, security, or team ownership requires it.

## Consequences

- local development and self-hosting remain simpler;
- transactional invariants are easier to enforce;
- background work still requires idempotent outbox processing;
- module discipline and architecture tests are necessary to avoid a tangled monolith;
- hosted scale may later extract federation, artifacts, search, analytics, or run control.

## Evidence gate

Accept after Phase 0 benchmarks show the relational/event design meets the first pilot and can rebuild projections within target.

## Reconsider if

- a required execution or policy component has a different security/availability boundary from day one;
- an existing managed component replaces a module without coupling domain semantics;
- verified scale requirements exceed a single authority/database partition before the first pilot.

