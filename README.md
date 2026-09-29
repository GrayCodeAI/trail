# Trail by GraycodeAI

> Every outcome has a path.

Trail is an independent product by GraycodeAI for coordinating people, agents, services, repositories, and independently administered organizations. Its hosted product surface is intended for `trail.graycodeai.com`. It coordinates outcomes through explicit commitments, scoped authority, evidence, and selective federation.

This directory contains the research and planning baseline plus a **bounded Phase 1 executable domain experiment**. The experiment validates protocol semantics in memory; it is not product infrastructure. Phase 0 field validation remains open.

## Product thesis

Trail is not a decentralized clone of Jira, Linear, Asana, Monday, ClickUp, or GraphDone. Tasks, boards, lists, timelines, and graphs are useful views. The stable product model is:

```text
situation -> desired outcome -> request/offer -> commitment
          -> authorized execution -> evidence -> review -> acceptance
```

The human-to-agent ratio is measured, never configured as a product constant. A workflow may be 10/90, 80/20, 90/10, or change at each transition. Policy, risk, authority, confidence, capability, and available attention decide who may act.

## Document map

### Planning

- [Research plan](./docs/planning/RESEARCH_PLAN.md)
- [Product requirements](./docs/planning/REQUIREMENTS.md)
- [Edge case catalog](./docs/planning/EDGE_CASES.md)
- [Phased roadmap](./docs/planning/PHASED_ROADMAP.md)
- [Pre-build gate](./docs/planning/PRE_BUILD_GATE.md)
- [Problem interview guide](./docs/planning/INTERVIEW_GUIDE.md)
- [Public interview candidate shortlist](./docs/planning/PUBLIC_CANDIDATE_SHORTLIST.md)
- [Design-partner brief](./docs/planning/DESIGN_PARTNER_BRIEF.md)
- [Concierge prototype specification](./docs/planning/CONCIERGE_PROTOTYPE.md)
- [Pilot scorecard](./docs/planning/PILOT_SCORECARD.md)
- [Risk register](./docs/planning/RISK_REGISTER.md)
- [Open questions and experiments](./docs/planning/OPEN_QUESTIONS.md)

### Architecture

- [System architecture](./docs/architecture/ARCHITECTURE.md)
- [Domain model](./docs/architecture/DOMAIN_MODEL.md)
- [Distributed systems model](./docs/architecture/DISTRIBUTED_SYSTEMS_MODEL.md)
- [Security and threat model](./docs/architecture/SECURITY_AND_THREAT_MODEL.md)

### Protocol and decisions

- [Trail Protocol v0 draft](./spec/trail-protocol-v0.md)
- [Protocol v0 JSON Schema](./spec/schema/v0/README.md)
- [Protocol v0 executable fixtures](./spec/fixtures/v0/README.md)
- [Architecture decisions](./docs/decisions/README.md)

### Research

- [Research index and method](./research/README.md)
- [Open source landscape](./research/sources/oss-landscape.md)
- [Market, OSS, and startup signals](./research/sources/market-community-and-startup-signals.md)
- [Industry standards and readiness](./research/sources/industry-standards-and-readiness.md)
- [Research literature](./research/sources/research-literature.md)
- [Protocols and standards](./research/sources/protocols-and-standards.md)
- [GraycodeAI integration audit](./research/notes/graycodeai-integration.md)
- [Field evidence ledger](./research/evidence/README.md)
- [Initial research archive](./research/initial/README.md)

## Current decisions

1. The product is **Trail**, endorsed publicly as **Trail by GraycodeAI**.
2. Trail owns outcome and commitment coordination. It does not own model routing, coding-agent interaction, Git hosting, worktree execution, or raw engineering memory.
3. The core is a typed temporal coordination model with a signed event journal and rebuildable projections.
4. Local-first behavior applies where operations merge safely. Scarce, revocable, or exclusive rights coordinate at the owning authority.
5. Federation is selective exchange between trust domains. There is no global database, token, blockchain, or universal consensus requirement.
6. Agents are principals with bounded capabilities. Natural-language output alone never authorizes a side effect or proves completion.
7. Product work is proposed to begin as a modular monolith and protocol package. Federation and service extraction follow evidence from a working local kernel.
8. Trail is independently deployable and self-hostable. Graycode Platform integration is optional and occurs through versioned HTTP APIs; there is no shared database or required Platform runtime.

## Executable reference model

ADR 0006 authorizes a dependency-free Go kernel and deterministic simulator under a narrow Phase 1 boundary. It currently exercises explicit commitment formation, scoped grants, leases and fencing, bounded runs, independent evidence, exact-version acceptance, idempotency, policy decisions, event replay, JSON Schema, and portable conformance fixtures. The canonical lifecycle enters through the public JSON command envelope and must reproduce the checked-in 11-event journal byte for byte.

```sh
make check
make simulate
go run ./cmd/trail-sim -mode scenario
```

See the [implementation status](./docs/planning/IMPLEMENTATION_STATUS.md) for covered requirements and remaining gates.

## Planning gates

Production implementation starts only after these artifacts are reviewed together:

- named pilot scenarios and users;
- protocol vocabulary and lifecycle invariants;
- authority, privacy, and federation boundaries;
- adversarial edge cases and threat mitigations;
- measurable phase exit criteria;
- explicit decisions for storage, runtime, and licensing.

The in-memory experiment does not satisfy these field, operational, or product gates.

## Status vocabulary

Documents use these labels:

- **Observed:** verified in a repository, specification, paper, or product source.
- **Inferred:** a conclusion drawn from multiple observations.
- **Proposed:** a Trail design choice that still needs validation.
- **Decided:** an accepted project decision recorded in an ADR.
- **Open:** unresolved and attached to an experiment or decision gate.
