# Trail phased roadmap

**Version:** 0.1 research baseline  
**Date:** 2026-09-28  
**Planning principle:** evidence gates progress; calendar dates do not

## Sequence

```mermaid
flowchart LR
    R[0 · Research and falsification]
    P[1 · Protocol and executable domain]
    K[2 · Local coordination kernel]
    U[3 · Useful product and read adapters]
    E[4 · Bounded agent execution]
    L[5 · Local-first collaboration]
    F[6 · Selective federation]
    I[7 · Intelligence and ecosystem]

    R --> P --> K --> U --> E --> L --> F --> I
```

Research, security, accessibility, and operations continue through every phase. Each phase produces a usable or falsifiable result. A later phase does not justify leaving an earlier safety invariant incomplete.

## Phase 0: research and falsification

### Question

Is a commitment/evidence fabric a real product need, and what is the smallest pilot?

### Work

- run user and failure-story interviews described in the research plan;
- complete the [pre-build gate](./PRE_BUILD_GATE.md), including design partners, concierge workflows, and the build/no-build record;
- test vocabulary and paper/interactive flows;
- benchmark equivalent workflows in an OSS project manager, GraphDone, and an agent control plane;
- choose two pilot scenarios and task/risk classes;
- complete security/privacy impact framing;
- resolve Trail/Radius repository scope and ecosystem inventory;
- perform license and trademark review;
- turn open questions into bounded spikes.

### Artifacts

- research evidence ledger;
- validated job stories and anti-personas;
- task-class and risk taxonomy;
- revised requirements and edge-case priorities;
- product narrative and success/failure measures;
- decision on first deployment shape and license.

### Exit criteria

- Gates 0A, 0B, and 0C in the research plan pass;
- a pilot group commits to trying the workflow;
- users can distinguish request, commitment, execution completion, evidence, and acceptance;
- the team can state one reason to stop building Trail based on collected evidence;
- no implementation dependency is selected only from popularity or familiarity.

### Stop/revise signals

- ordinary issue trackers plus agent integrations solve the pilot with equal clarity and lower cost;
- users reject explicit commitments or evidence as overhead in every target scenario;
- the GraycodeAI boundary requires duplicating Rho/Rover/Across/Trace;
- meaningful federation requires unacceptable disclosure.

## Phase 1: protocol and executable domain model

### Question

Can Trail's semantics remain correct under concurrency, failure, revocation, and change?

### Work

- stabilize the minimum vocabulary in Trail Protocol v0;
- write JSON Schema, canonical examples, and invalid/adversarial fixtures;
- specify commands, events, errors, versioning, and capability negotiation;
- implement an in-memory reference model or simulator, not product infrastructure;
- model-check commitments, lifecycle, leases/fencing, budgets, acceptance, and revocation;
- specify bitemporal, supersession, dispute, and redaction behavior;
- define adapter mappings for Rho, Rover, Across, and Trace on paper;
- publish compatibility and change policy.

### Minimum object set

- principal and identity binding;
- workspace and policy reference;
- situation claim;
- outcome and criterion;
- request, offer, and commitment;
- work item and typed relation;
- capability/grant;
- claim/lease;
- run;
- artifact, evidence, review, and acceptance;
- decision;
- event and disclosure envelope.

### Exit criteria

- every P0 domain invariant has an executable property;
- adversarial fixture suite covers all critical edge cases applicable to this phase;
- two language implementations parse and validate canonical events;
- lifecycle model has no unintended terminal state, double acceptance, stale lease action, or authority escalation;
- at least two GraycodeAI adapter owners agree the public contract does not require internal imports;
- protocol changes can still be made cheaply because no production data exists.

## Phase 2: local coordination kernel

### Question

Can one node preserve authoritative, auditable coordination without agent execution or federation?

### Work

- implement authentication boundary and principal model;
- implement command handlers and domain invariants;
- append accepted facts to a durable event journal;
- build transactional outbox/inbox and idempotency records;
- build current-state, audit, list, and simple dependency projections;
- store artifact manifests and content digests;
- implement policy decision interface with a simple bounded policy corpus;
- add backup, restore, projection rebuild, migration, and corruption drills;
- add OpenTelemetry without content/secret leakage.

### Deployment

One modular monolith, one primary database, one artifact store abstraction, one background worker process. No service decomposition and no graph database unless benchmark evidence requires one.

### Demonstration

A human creates an outcome, negotiates a commitment, records work and evidence, requests review, and accepts or rejects the exact version. Restart and rebuild produce the same projections.

### Exit criteria

- command/event atomicity survives injected process crashes;
- duplicate commands and outbox delivery do not duplicate state;
- projection rebuild is deterministic and meets the initial throughput target;
- backup restore reconstructs identities, events, policies, artifacts, and audit;
- all P0 security requirements and critical storage/policy edge cases pass;
- no client, agent, or MCP tool can write the event store or projection directly.

## Phase 3: useful product and read-only adapters

### Question

Does Trail improve a real human workflow before it can autonomously act?

### Work

- build outcome, now/ready, commitment, evidence, review, and decision surfaces;
- add accessible list/table equivalents before advanced graph visualization;
- implement resumption capsules and actionable attention queue;
- add comments as non-authoritative collaboration around domain objects;
- implement import/export and usable local backups;
- add read-only adapters for Rho graph exports and one Git/Trace provider;
- show connector freshness, provenance, unresolved identity, and mapping status;
- run the first pilot with agents observed as participants but without Trail dispatch authority.

### Experience constraints

- common capture within a small, tested interaction budget;
- keyboard-first operation;
- no graph-only workflow;
- explanation for readiness, denial, recommendation, and acceptance;
- observed/inferred/stale/disputed/accepted labels visible;
- notification volume and review queue measured.

### Exit criteria

- pilot users complete the core workflow without facilitator intervention;
- Trail reduces at least one selected failure measure versus the baseline;
- user comprehension meets the vocabulary threshold;
- WCAG 2.2 AA audit passes for core flows;
- read adapters cannot mutate source systems;
- export can reconstruct a useful record without the running server.

### Product decision

If users value the views but avoid commitments and evidence, revise the ontology before adding execution. If they value only agent monitoring, narrow or stop rather than building another control-plane clone.

## Phase 4: bounded agent execution

### Question

Can Trail safely authorize work through existing runtimes and keep acceptance independent?

### Work

- implement least-privilege capability grants, lease/fencing, budget reservation, and run lifecycle;
- integrate one execution path, preferably Rover with Rho as the acting agent;
- add checkpoint, heartbeat, cancellation, timeout, retry, reconciliation, and compensation paths;
- add secret brokering at the tool boundary;
- import Rover evidence and Across provenance references;
- add risk-tier policy templates and independent verification;
- add operator controls and incident runbooks;
- map controls to current OWASP agentic threats and NIST AI RMF lifecycle functions.

### Autonomy ladder

1. **Observe:** read context and produce no durable proposal.
2. **Propose:** create plans, work, decisions, or evidence proposals.
3. **Execute bounded:** act within a claim, sandbox, capability, and budget.
4. **Verify:** run deterministic checks or scoped independent evaluation.
5. **Decide bounded:** make delegated low-risk decisions under explicit policy.
6. **Accept:** permitted only for specifically delegated criteria and risk classes.

Each stage is independently configured. A workspace-wide “90% autonomous” switch is prohibited.

### Exit criteria

- stale/expired workers cannot perform protected new side effects;
- cancellation and uncertain-side-effect races reconcile safely;
- agents cannot expand capabilities, exceed reserved budget, or access undeclared resources;
- execution completion cannot bypass evidence or acceptance policy;
- high-risk executor/verifier separation is enforced;
- kill, quarantine, secret rotation, and compromised-run drills pass;
- cost per accepted outcome and human attention per outcome improve over baseline for at least one task class.

## Phase 5: local-first collaboration

### Question

Can participants retain useful local ownership without weakening authority semantics?

### Work

- add encrypted device-local projection and pending proposal queue;
- choose and integrate a CRDT for descriptions, comments, and selected planning documents;
- implement offline drafts and conflict explanations;
- separate mergeable local facts from authority-coordinated commands in APIs and UX;
- add device enrollment, key rotation, loss/revocation, cache policy, and compaction;
- test long-offline schema upgrade and data export;
- run partition, reordering, duplicate, quota, device-loss, and erasure drills.

### Exit criteria

- two devices edit offline and later converge on selected documents;
- exclusive claims, budgets, grants, revocations, and acceptance never merge through a CRDT;
- rejected offline commands preserve the user's draft and explanation;
- revoked devices cannot resume synchronization;
- local encrypted data and keys meet the threat model;
- long-offline clients have an upgrade or export path without data loss.

## Phase 6: selective federation

### Question

Can two independently administered nodes complete a shared commitment while keeping private work private?

### Work

- implement node discovery, protocol negotiation, keys, rotation, and trust configuration;
- implement addressed signed inbox/outbox with receipts, retry, deduplication, replay defense, quarantine, and backpressure;
- federate the minimum contract objects and selected evidence;
- implement proposal/accept/reject rather than remote direct mutation;
- implement peer block, node suspension, tombstone, dispute, revocation, and authority migration experiments;
- privacy-test error messages, counts, timing, and derived status for topology leakage;
- publish conformance fixtures and run a second independent implementation.

### Minimum federation scenario

1. Node A publishes a bounded request.
2. Node B offers and both accept a commitment version.
3. The home authority issues a lease for the shared work.
4. B completes private internal work without exposing its graph.
5. B submits allowed artifacts/evidence.
6. A accepts, rejects, or requests change.
7. Both retain the shared signed history and can export it.

### Exit criteria

- all critical federation edge cases pass;
- neither node needs administrator access to the other;
- unrelated private objects, identities, counts, and topology are absent from exchange and inference tests;
- partitions converge to an explicit shared result or visible dispute, never silent last-write-wins;
- peer abuse can be rate-limited/blocked without disabling local work;
- two independent implementations pass conformance.

## Phase 7: decision intelligence and ecosystem

### Question

Which advanced analyses measurably improve outcomes without reducing agency or stability?

### Work

- typed dependency expressions and scalable reachability;
- SCC, cut-set, bottleneck, and criticality analysis;
- stochastic schedule simulation and calibration;
- Pareto choice and value-of-information suggestions;
- capacity-aware dynamic allocation with hysteresis;
- task-class-specific, privacy-preserving actor reliability estimates;
- plugin/adapter SDK and conformance program;
- MCP, A2A, ForgeFed, OSLC, and BPMN profiles where pilots require them;
- hosted multi-region and enterprise operational profiles only from measured need.

### Admission rule for every algorithm

An algorithm ships only when it beats a simpler baseline on a pilot measure, exposes inputs and uncertainty, has gaming/fairness/stability tests, remains contestable, and can be disabled without corrupting domain state.

### Exit criteria

- recommendations are reproducible from stored inputs and versioned code/model identity;
- calibration and outcome benefit exceed chosen baseline;
- no universal hidden priority or reputation score is introduced;
- allocation remains stable under noisy inputs and manipulation tests;
- plugin capabilities are isolated and revocable;
- protocol governance and compatibility policy support external implementers.

## Continuous workstreams

### Security and privacy

Threat modeling, dependency/supply-chain review, secret scanning, adversarial fixtures, privacy impact, incident drills, and independent audit.

### Reliability and operations

Backups, restores, migrations, projection rebuild, chaos/fault injection, capacity, SLOs, and on-call runbooks.

### Accessibility and human factors

Keyboard/screen-reader testing, cognitive load, interruption/resumption, approval quality, attention load, and takeover recovery.

### Research and governance

Evidence ledger updates, counter-evidence, protocol change process, ADRs, pilot ethics, and public compatibility documentation.

## Suggested initial team capabilities

The work needs capabilities, not fixed titles:

- product discovery and CSCW research;
- domain modeling and distributed systems;
- security, identity, authorization, and privacy;
- accessible interaction design and frontend engineering;
- event-driven backend and data engineering;
- agent runtime/sandbox integration;
- formal methods, simulation, or operations research;
- developer relations/protocol documentation once external integrations begin.

One person may cover several capabilities early. Security and user research need independent review even in a small team.

## Phase metrics

| Dimension | Measures |
|---|---|
| outcome | time to validated outcome, accepted outcome rate, 30/90-day survival, rework |
| coordination | commitment renegotiation, readiness errors, blocked time, dependency age |
| attention | decisions/reviews per outcome, interruption rate, queue wait, approval error |
| agents | accepted result per cost/time, retries, prevented unsafe actions, verifier disagreement |
| resilience | recovery time, duplicate side effects, lost drafts, rebuild/restore success |
| federation | delivery/convergence lag, duplicate/quarantine rate, disclosure incidents |
| trust | contested decisions, explanation comprehension, takeover success, policy-denial quality |

Metrics are scoped by task class and risk. They are diagnostic signals, not worker rankings.
