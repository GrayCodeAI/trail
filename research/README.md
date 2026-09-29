# Trail research index

**Research cut:** 2026-09-28  
**Status:** active, living corpus

## Purpose

This corpus supports a decision about what Trail should be, which guarantees it should provide, and what must be learned before implementation. It does not attempt to enumerate every repository, paper, product, or social post ever published. That claim would be neither testable nor maintainable.

The review is designed for coverage across distinct evidence classes:

1. mature open source work-management systems;
2. agent operation and coding-agent control planes;
3. local-first and replicated data systems;
4. federated collaboration and linked-data protocols;
5. multi-agent commitments and human-machine teamwork;
6. distributed systems, scheduling, control, and decision theory;
7. identity, authorization, provenance, and agent security;
8. the current GraycodeAI repositories and their ownership boundaries.

## Method

### Inclusion criteria

A source is included when it contributes at least one of:

- a production-tested mechanism Trail may need;
- a distinct domain model or product interaction;
- a protocol with an implementation or standards process;
- peer-reviewed or historically important theory;
- a documented failure, security constraint, or counterexample;
- direct evidence about an existing GraycodeAI component.

### Evidence order

1. normative specifications and repository source;
2. peer-reviewed papers and official research archives;
3. official documentation and product definitions;
4. maintainer articles and incident reports;
5. community discussion and social posts as discovery signals only.

### Review questions

For every candidate, ask:

- What is the authoritative unit: task, document, event, issue, goal, commitment, or conversation?
- Who owns a mutation, and what happens during partitions?
- How are identity, authority, revocation, and delegation represented?
- Can work be resumed and audited after failures?
- How is completion distinguished from a status claim?
- What data can leave one organization or device?
- Which design depends on centralized control despite self-hosting?
- Which behavior has evidence, and which is marketing or aspiration?

## Research outputs

- [Open source landscape](./sources/oss-landscape.md)
- [Market, OSS, and startup signals](./sources/market-community-and-startup-signals.md)
- [Industry standards and readiness](./sources/industry-standards-and-readiness.md)
- [Research literature](./sources/research-literature.md)
- [Protocols and standards](./sources/protocols-and-standards.md)
- [GraycodeAI integration audit](./notes/graycodeai-integration.md)
- [Field evidence ledger](./evidence/README.md)
- [Initial research archive](./initial/README.md)

## Strongest synthesis

The evidence supports a **federated commitment and situation fabric**:

- established work managers show the value of fast capture, flexible views, search, imports, and predictable workflows;
- agent control planes show the need for atomic claims, durable runs, budgets, approvals, scoped secrets, and completion evidence;
- local-first research supports user-owned replicas for mergeable knowledge;
- CALM and CRDT research warn against merging scarce or revocable authority without coordination;
- multi-agent research makes social commitments more interoperable than private plans or chat logs;
- federation protocols demonstrate inbox/outbox delivery and home authorities, while also exposing moderation, privacy, deletion, and schema-evolution problems;
- human factors research favors adjustable autonomy, legible handoffs, recovery support, and protection of human attention.

## Research gaps

The next evidence must come from pilots, simulations, and adversarial prototypes:

- whether teams naturally use request, offer, commitment, and acceptance language;
- whether evidence-backed acceptance reduces rework enough to justify added structure;
- whether cross-organization coordination creates value without disclosing private topology;
- which task classes tolerate high agent autonomy;
- whether people understand dynamic authority and can recover control under time pressure;
- the smallest interoperability profile that Rho, Rover, Across, and Trace can support without coupling their internals.
