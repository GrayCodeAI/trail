# Trail research plan

**Owner:** Trail product and architecture  
**Started:** 2026-09-28  
**Current state:** desk research complete enough to define Phase 0; field validation pending

## Objective

Determine whether a federated commitment and situation fabric solves coordination problems for mixed human-agent teams better than a task manager or agent dashboard, and define the smallest safe product that can test that claim.

## Core hypotheses

| ID | Hypothesis | Falsifier | Evidence method |
|---|---|---|---|
| H-01 | Outcomes, commitments, and evidence retain intent better than task/status records | users repeatedly translate them back into ordinary tickets or cannot distinguish the terms | scenario interviews, prototype tasks, comprehension tests |
| H-02 | Evidence-backed acceptance reduces false completion and rework | review time rises without reducing escaped errors or rework | matched pilot workflows |
| H-03 | Dynamic autonomy outperforms a fixed human-agent ratio | users cannot predict or control authority, or outcomes worsen as autonomy changes | staged simulations and longitudinal pilot |
| H-04 | A typed dependency model improves readiness and unblock decisions | modeling cost exceeds planning benefit or users ignore relationship types | compare binary graph, expression model, and ordinary board |
| H-05 | Local-first drafts and projections improve resilience and ownership | sync/conflict complexity outweighs offline value for target users | offline diary study and failure drills |
| H-06 | Selective federation enables useful cross-organization commitments | required disclosure or administrative burden erases the benefit | two-organization pilot |
| H-07 | GraycodeAI's existing products let Trail remain a small coordination kernel | adapters require copying runtime, forge, or context responsibilities | adapter spike and contract review |
| H-08 | Recommendations can remain legible under graph and scheduling complexity | users cannot reproduce, contest, or appropriately ignore a recommendation | explanation and calibration tests |

## Research workstreams

### R1. Users and jobs

Target participants:

- engineering leads operating two or more coding agents;
- individual developers working across several repositories;
- product/engineering teams with external contributors;
- open source maintainers coordinating volunteers and bots;
- operations or research teams with evidence-heavy acceptance;
- security and platform administrators responsible for agent authority.

Questions:

- Where is outcome intent currently lost?
- Which work is performed only to coordinate other work?
- What counts as completion and who may decide?
- Which agent actions require approval, notification, or later review?
- What information may cross organizational boundaries?
- What makes an interrupted person or agent safe to resume?
- Which failures currently cause duplicated work, silent drift, or hidden cost?

Deliverables:

- 12 to 20 interview notes with consent and redaction;
- job stories and failure narratives;
- task-class/risk taxonomy;
- ranked pilot scenarios;
- explicit evidence against the thesis.

### R2. Domain language

Test these terms with examples, not definitions alone:

- situation;
- outcome;
- request and offer;
- commitment;
- dependency/condition;
- claim/lease;
- artifact and evidence;
- review and acceptance;
- dispute, violation, supersession, and revocation.

Exit signal: at least 80% of pilot participants correctly predict lifecycle behavior in scenario tests, and no critical pair remains consistently confused.

### R3. Existing-system benchmark

Build the same representative workflow in:

- one modern OSS project manager;
- GraphDone;
- one agent control plane;
- a Trail paper/interactive prototype.

Measure:

- capture steps and time;
- lost or duplicated facts;
- readiness errors;
- authority ambiguity;
- time to resume after interruption;
- time to decide whether an output is acceptable;
- export and recovery behavior.

### R4. Formal domain model

Produce:

- lifecycle state machines;
- typed dependency truth tables;
- commitment state transitions;
- authority and capability algebra;
- bitemporal rules;
- federation conflict rules;
- TLA+/Alloy or equivalent model for leases, acceptance, and revocation;
- property-based model tests independent of storage.

High-risk properties:

- one valid fencing token for exclusive work;
- expired/revoked authority cannot authorize a new side effect;
- acceptance cannot occur without the applicable evidence policy;
- duplicate delivery cannot duplicate a commitment or spend;
- a remote node cannot mutate a home-owned object by assertion;
- private content cannot be inferred from unauthorized projections.

### R5. Simulation and operations research

Build a simulator before production scheduling code. It should support:

- typed dependency expressions;
- stochastic service time and failure;
- human review queues and interruption cost;
- agent cost, rate limits, retries, and confidence;
- changing capability and availability;
- work preemption and cancellation;
- communication and integration overhead;
- partitions and delayed federation;
- adversarial or strategically reported priorities.

Compare simple baselines with advanced methods:

- FIFO and manual choice;
- static priority;
- critical path probability;
- bottleneck/cut-set suggestions;
- Pareto choice sets;
- value-of-information actions;
- constrained dynamic allocation with hysteresis.

### R6. Security and abuse

Threat-model:

- humans, agents, connectors, runtimes, nodes, model providers, and external content;
- trust boundaries and credential flows;
- prompt injection and tool-description poisoning;
- confused deputy and delegation-chain abuse;
- malicious or compromised federation peers;
- evidence forgery and verifier collusion;
- insider surveillance and retaliation;
- budget/resource exhaustion;
- recovery, key rotation, and node compromise.

Deliverables:

- threat model tied to requirements;
- abuse-case tests;
- least-agency capability profiles;
- incident and kill-control runbooks;
- privacy impact assessment for telemetry and performance data.

### R7. Interoperability spikes

Read-only spikes only during Phase 0:

1. Import a Rho graph export into a Trail semantic mapping.
2. Map one Trace issue, pull request, review, and CI result to Trail references.
3. Represent one Rover run and its evidence as a proposed completion.
4. Reference one Across bundle while preserving its epistemic labels.
5. Round-trip canonical Trail events through JSON Schema and two languages.

Do not dispatch real side effects during the research phase.

### R8. Federation experiment

Use two independently configured nodes and a deliberately small shared contract:

- Node A offers a bounded piece of work.
- Node B accepts and creates a commitment.
- A home authority issues a claim/lease.
- B submits a content-addressed artifact and evidence.
- A accepts, rejects, or requests change.
- both nodes retain the shared history while internal child work remains private.

Inject duplicate delivery, reordering, clock skew, key rotation, revocation, node outage, schema mismatch, content removal, and a malicious peer.

## Pilot scenarios

### P1. Cross-repository software change

A human defines an outcome spanning two repositories. A planner proposes work. Two agents operate in isolated worktrees. Rover verifies. Trace records pull requests. Across preserves handoff context. A human accepts the outcome based on combined evidence.

### P2. External contributor contract

One organization requests a deliverable from another without disclosing its internal roadmap. Both negotiate acceptance criteria, exchange only the shared commitment, and retain their own local work structures.

### P3. High-autonomy maintenance

Agents handle low-risk dependency updates and tests. Human attention is requested only for policy exceptions, conflicting evidence, budget changes, security-sensitive diffs, and acceptance samples.

### P4. Incident recovery

An agent or node fails after performing an external side effect but before recording completion. The system prevents unsafe re-execution, reconstructs the situation, and routes a compensating or human action.

### P5. Offline field work

A participant records observations and drafts work offline, reconnects after concurrent changes, and sees semantic conflicts without losing local work.

## Evidence ledger fields

Every future research entry should record:

```text
id
question or claim
source type
source URL or repository/path/commit
publication and access dates
observed result
counter-evidence
applicability to Trail
confidence
affected requirement/decision
reviewer
next review date
```

## Decision gates

### Gate 0A: problem evidence

- at least three recurring failure modes across five or more target teams;
- one pilot group willing to test an evidence-backed commitment flow;
- terminology comprehension threshold met or vocabulary revised;
- explicit evidence that existing task managers and agent dashboards leave the chosen problem unsolved.

### Gate 0B: domain safety

- lifecycle and authority invariants model-checked;
- edge-case catalog reviewed by product, distributed-systems, and security perspectives;
- every high-impact side effect has an authority, idempotency, and recovery contract;
- privacy and retention requirements cover events, artifacts, telemetry, and federation.

### Gate 0C: implementation choice

- protocol v0 reviewed by at least two future adapter owners;
- storage decision benchmarked with a rebuild test;
- CRDT choice proven on representative documents and schema migration;
- build versus reuse decision for durable execution;
- license and trademark review complete;
- Radius disposition and ecosystem inventory decisions recorded.

Implementation may begin only after all three gates are satisfied or an ADR records a deliberately narrower experiment.

