# ADR 0001: Trail owns outcome and commitment coordination

- **Status:** Accepted
- **Date:** 2026-09-28

## Context

The workspace already contains model routing, coding-agent interaction, execution supervision, engineering continuity, skills, a Git forge, and hosted platform infrastructure. Building another all-in-one project manager or agent runtime would duplicate those products and obscure ownership.

Research across GraphDone, traditional work managers, CSCW, requirements engineering, and multi-agent commitments suggests that tasks and graph nodes alone lose social obligation, shared situation, evidence, and acceptance authority.

## Decision

Trail owns:

- situations and desired outcomes;
- requests, offers, commitments, delegation, renegotiation, release, and violation;
- typed relations and readiness;
- scoped policy, attention, evidence, review, and acceptance;
- selected cross-organization boundary contracts;
- list, board, timeline, inbox, and graph projections of those facts.

Trail integrates Rho, Rover, Across, Trace, Flux/Skills, and Graycode Platform. It does not absorb their runtimes or storage.

## Consequences

- the first product can be narrower than a full project-management suite;
- adapters and a public protocol are core work;
- ordinary tasks remain useful but are subordinate to outcome/commitment semantics;
- chat and Git content should be referenced or projected rather than copied wholesale;
- Trail's usefulness must be validated independently from the value of existing GraycodeAI tools.

## Reconsider if

- user research shows commitments/evidence add no benefit beyond a well-integrated issue tracker;
- adapters cannot preserve necessary semantics without owning the external systems;
- one existing GraycodeAI product deliberately expands to own this domain.

