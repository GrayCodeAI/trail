# Trail repository instructions

## Current phase

Trail is in research and Phase 1 executable-domain validation. Phase 0 field-validation gates remain open. ADR 0006 authorizes only a dependency-free, in-memory Go reference kernel, simulator, protocol fixtures, and invariant tests. Do not add product APIs, durable storage, web or native UI, external-system adapters, generated clients, database migrations, infrastructure, or deployment configuration until the applicable gates in `docs/planning/PHASED_ROADMAP.md` are satisfied.

Reference code belongs under `internal/domain`, `cmd/trail-sim`, and `spec/fixtures`. It may prove or falsify protocol semantics; it must not establish an accidental production API or persistence format.

## Product boundary

Trail owns outcomes, situations, commitments, evidence references, acceptance, scoped policy, attention queues, and selective federation.

Keep these responsibilities in their existing GraycodeAI products:

- Rho: coding-agent interaction, tools, permissions, and sessions.
- Flux: model providers, catalogs, streaming, usage, retry, and fallback.
- Rover: isolated execution, DAG supervision, verification, and execution evidence.
- Across: engineering context, checkpoints, handoffs, and provenance bundles.
- Trace: Git hosting, issues, pull requests, reviews, releases, and CI.
- Graycode Skills: capability and skill catalog.
- Graycode Platform: company web, cross-product account and billing services. Trail is an independent product and integrates with Platform only through optional, versioned HTTP contracts.

Integrate through versioned protocols and adapters. Do not import private or `internal` packages from sibling repositories. Trail has its own repository, deployment, persistence, identity configuration, API, release cycle, and operational boundary. Graycode Platform integration is HTTP only unless its own repository rules are deliberately changed.

## Research rules

1. Prefer primary sources: specifications, papers, official repositories, official product documentation, and source code.
2. Record source URL, access date, source type, and the design claim it supports.
3. Mark statements as observed, inferred, proposed, decided, or open.
4. Use social posts only to discover sources or user problems. Do not treat popularity or an unverified post as architectural evidence.
5. Record counter-evidence and failure modes, not only supportive evidence.
6. Do not claim an exhaustive review of “all” software or literature. State scope and inclusion criteria.

## Architecture rules

1. Treat graph, board, list, timeline, and chat as projections.
2. Keep commands distinct from facts, proposals distinct from accepted mutations, and execution distinct from acceptance.
3. Every accepted mutation needs an authenticated principal, policy decision, schema version, idempotency key, and causal reference where applicable.
4. Use CRDTs only for state whose concurrent changes can merge without violating authority, budget, exclusivity, acceptance, or revocation rules.
5. Model agents as accountable principals with operators, capabilities, budgets, leases, and evidence requirements.
6. Make recommendations reproducible from recorded inputs and keep them advisory unless a policy explicitly delegates the decision.
7. Preserve bitemporal history: when a fact was valid and when Trail learned or recorded it.
8. Never expose private graph structure through derived scores, counts, error messages, or federation metadata.

## Change discipline

- Update requirements and edge cases before changing the protocol semantics.
- Record durable architecture choices as ADRs.
- Add a migration and compatibility note for every breaking protocol proposal.
- Link every phase exit criterion to tests, simulations, research evidence, or a named pilot result.
- Preserve the initial research archive under `research/initial/` as historical context.
