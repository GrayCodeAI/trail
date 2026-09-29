# Trail product requirements

**Version:** 0.1 research baseline  
**Date:** 2026-09-28  
**Status:** proposed; validation and prioritization pending

## Requirement language

- **MUST:** necessary for the stated phase to be considered safe and complete.
- **SHOULD:** expected unless an ADR records a justified exception.
- **MAY:** optional or experimental.
- **P0:** coordination-kernel requirement.
- **P1:** first useful single-node product.
- **P2:** local-first and agent-operation expansion.
- **P3:** federation.
- **P4:** ecosystem and optimization.

Requirement IDs remain stable. Removed requirements are marked superseded rather than renumbered.

## Product goals

1. Preserve a traceable path from situation and intent through commitment, execution, evidence, and acceptance.
2. Let humans and agents participate at ratios that change by workflow stage, risk, capability, and policy.
3. Protect authority, attention, privacy, budget, and the ability to recover control.
4. Let participants own usable local data and export complete authorized records.
5. Coordinate across independent organizations without centralizing their private internal work.
6. Integrate GraycodeAI's agent, execution, provenance, forge, skills, and hosted-platform products through narrow adapters.

## Non-goals for the first product

- full replacement for Jira, Slack, Notion, GitHub, or an ERP;
- model-provider routing or a new coding-agent runtime;
- Git hosting, CI, worktree management, or raw transcript storage;
- universal autonomous project planning;
- global organizational ranking or employee surveillance;
- cryptocurrency, token incentives, or public blockchain consensus;
- whole-workspace federation before selective commitments are proven;
- one visualization that renders every object and relation.

## Actors and roles

| Actor | Responsibilities and needs |
|---|---|
| contributor | proposes, commits, performs work, submits evidence, renegotiates honestly |
| outcome steward | defines desired state and measures, resolves ambiguity, accepts or delegates acceptance |
| requester/creditor | requests and relies on a promised consequent |
| provider/debtor | makes and manages a commitment |
| reviewer/verifier | evaluates evidence independently where policy requires |
| operator | monitors runs, failures, spend, queues, federation, and recovery |
| administrator | configures membership, trust, retention, integrations, and policy bounds |
| auditor | reconstructs decisions and actions without receiving unnecessary content |
| agent | acts as a named principal under an operator, runtime identity, capabilities, lease, and budget |
| connector | maps an external source under a service identity and explicit authority |
| remote node | independently administered trust domain exchanging selected contracts |

A principal may hold several roles. Role assignment is scoped, versioned, time-bounded where appropriate, and auditable.

## Core scenarios

### S-01: outcome to accepted result

A person describes a measurable outcome. Humans or agents propose a plan. Authorized actors form commitments, claim ready work, execute, attach evidence, request review, and receive scoped acceptance. Every transition remains traceable.

### S-02: changing autonomy

An agent performs low-risk steps without interruption. Risk, uncertainty, policy, cost, or disagreement increases. Trail pauses or narrows authority and routes a concise decision packet to a human. After a decision, automation can resume under a new grant.

### S-03: cross-organization delivery

Two nodes negotiate a commitment and acceptance policy. Each keeps private internal plans. They exchange only the shared contract, allowed status, artifacts/evidence, and accepted events.

### S-04: offline work

A participant reads authorized data and drafts updates offline. Mergeable edits synchronize later. A command affecting authority or scarce resources remains a pending proposal until the home node accepts it.

### S-05: failed run recovery

A run crashes, times out, loses its lease, or leaves an uncertain external side effect. Trail prevents a stale worker from continuing, distinguishes retryable execution from human recovery, and preserves evidence.

## Functional requirements

### Identity and principals

| ID | Requirement | Priority |
|---|---|---|
| FR-ID-001 | The system MUST assign stable IDs to people, agents, services, devices, workloads, organizations, and nodes without conflating these kinds. | P0 |
| FR-ID-002 | Every accepted mutation MUST identify the authenticated actor and the accountable principal/operator where they differ. | P0 |
| FR-ID-003 | A run MUST record the delegation chain from authorizing principal to agent to concrete runtime/workload. | P0 |
| FR-ID-004 | Identity bindings to external systems MUST be verified, versioned, revocable, and valid for a stated interval. | P1 |
| FR-ID-005 | Matching display name, email, agent label, or model name MUST NOT automatically merge principals. | P0 |
| FR-ID-006 | Historical attribution MUST remain verifiable after rename, membership change, key rotation, or deletion. | P1 |
| FR-ID-007 | Self-hosted use MUST NOT require a GraycodeAI cloud identity. | P1 |

### Workspaces, situations, and outcomes

| ID | Requirement | Priority |
|---|---|---|
| FR-OUT-001 | A workspace MUST define a home authority, membership scope, policy set, disclosure defaults, and retention policy. | P0 |
| FR-OUT-002 | Users MUST be able to record a situation as versioned observations, assumptions, unknowns, risks, constraints, and claims with provenance. | P1 |
| FR-OUT-003 | An outcome MUST state a desired condition, measures/acceptance criteria, steward, scope, horizon, and status. | P0 |
| FR-OUT-004 | Criteria MUST support qualitative, quantitative, deterministic, and reviewer-judgment forms. | P1 |
| FR-OUT-005 | The system MUST distinguish a changed outcome from a changed plan and preserve both histories. | P0 |
| FR-OUT-006 | Outcomes MUST support decomposition, contribution, conflict, and supersession without requiring a tree. | P1 |
| FR-OUT-007 | A participant MUST be able to dispute an observation, criterion, or outcome without overwriting the original claim. | P1 |

### Requests, offers, and commitments

| ID | Requirement | Priority |
|---|---|---|
| FR-COM-001 | The system MUST represent requester/creditor, provider/debtor, antecedent conditions, promised consequent, timing, policy, and lifecycle for a commitment. | P0 |
| FR-COM-002 | A commitment MUST arise from an explicit authorized act, never from assignment, mention, model inference, or silence alone. | P0 |
| FR-COM-003 | Requests, offers, counteroffers, acceptance, rejection, withdrawal, expiry, and renegotiation MUST be distinguishable. | P1 |
| FR-COM-004 | Commitment changes MUST preserve what each party accepted and which version is active. | P0 |
| FR-COM-005 | Delegation MUST preserve accountability, scope, and whether the original debtor remains liable. | P1 |
| FR-COM-006 | Discharge, release, cancellation, violation, and impossibility MUST have distinct semantics. | P1 |
| FR-COM-007 | Joint, alternative, and quorum commitments SHOULD be representable through typed multi-party relations. | P2 |
| FR-COM-008 | A remote party MUST be able to reject a proposal without the proposal becoming authoritative state at its node. | P3 |

### Work decomposition and dependencies

| ID | Requirement | Priority |
|---|---|---|
| FR-DEP-001 | The system MUST support work items as mutable decompositions that trace to an outcome or commitment. | P0 |
| FR-DEP-002 | Dependencies MUST be typed and carry rationale, author, confidence, strength, and validity interval where applicable. | P1 |
| FR-DEP-003 | Readiness MUST support all-of, any-of, quorum, threshold, temporal, resource, policy, and evidence conditions. | P2 |
| FR-DEP-004 | Readiness calculations MUST expose the satisfied and unsatisfied terms. | P1 |
| FR-DEP-005 | Cycles and strongly connected components MUST be detected and shown as coordination loops rather than silently linearized. | P1 |
| FR-DEP-006 | The system MUST preserve manual override as an explicit policy exception with rationale and expiry. | P1 |
| FR-DEP-007 | A work item SHOULD declare affected resources or isolation boundaries before parallel execution. | P2 |

### Decisions, recommendations, and planning

| ID | Requirement | Priority |
|---|---|---|
| FR-DEC-001 | A decision MUST record question, options, criteria, evidence, authority, result, rationale, and validity/supersession. | P1 |
| FR-DEC-002 | Recommendations MUST record inputs, algorithm/model version, constraints, uncertainty, and explanation. | P2 |
| FR-DEC-003 | The system MUST distinguish feasibility filtering from ranking. | P2 |
| FR-DEC-004 | Multi-objective choices SHOULD expose Pareto-efficient alternatives instead of forcing one opaque score. | P2 |
| FR-DEC-005 | An automated recommendation MUST remain advisory unless a scoped policy explicitly delegates the decision. | P1 |
| FR-DEC-006 | Users MUST be able to contest, ignore, or supersede a recommendation without falsifying its history. | P2 |
| FR-DEC-007 | Scheduling estimates MUST retain ranges, assumptions, and calibration history rather than only one date. | P2 |
| FR-DEC-008 | Automated reallocation MUST use rate limits/hysteresis and MUST NOT thrash actors on small signal changes. | P2 |

### Claims, leases, and execution

| ID | Requirement | Priority |
|---|---|---|
| FR-RUN-001 | Exclusive work MUST be protected by a time-bounded lease and monotonically increasing fencing token issued by its home authority. | P0 |
| FR-RUN-002 | An expired, released, or superseded lease MUST NOT authorize a new external side effect. | P0 |
| FR-RUN-003 | Run requests MUST include immutable work/plan version, capabilities, budget, deadline, checkpoints, environment, and evidence policy. | P0 |
| FR-RUN-004 | Runs MUST support queued, starting, active, waiting, blocked, paused, cancelling, cancelled, failed, timed-out, and completed execution states. | P1 |
| FR-RUN-005 | Execution completion MUST be distinct from domain acceptance. | P0 |
| FR-RUN-006 | Every externally visible side effect MUST have an idempotency, reconciliation, or compensation strategy. | P0 |
| FR-RUN-007 | Retry policy MUST distinguish deterministic failure, transient failure, policy denial, exhausted budget, uncertain side effect, and human decision required. | P1 |
| FR-RUN-008 | Operators MUST be able to pause, cancel, quarantine, or kill a run within documented propagation bounds. | P1 |
| FR-RUN-009 | A run MUST emit heartbeats/checkpoints appropriate to its duration and risk; missing heartbeats MUST NOT alone prove failure. | P1 |
| FR-RUN-010 | Parallel execution MUST require ready dependencies, isolated write sets/environments, review capacity, and an integration policy. | P2 |
| FR-RUN-011 | Secrets MUST be delivered only at the authorized tool boundary and MUST NOT be stored in prompts, events, evidence, or logs. | P0 |

### Capabilities, policy, and budget

| ID | Requirement | Priority |
|---|---|---|
| FR-POL-001 | Authorization MUST evaluate principal, action, resource, context, delegation chain, current policy version, and revocation state. | P0 |
| FR-POL-002 | Agent capabilities MUST be least-privilege, purpose-scoped, time-bounded, resource-bounded, and revocable. | P0 |
| FR-POL-003 | Delegated capabilities MUST NOT exceed the delegator's authority and SHOULD support attenuation. | P0 |
| FR-POL-004 | Policy decisions MUST return allow/deny/defer plus a reason safe for the caller's disclosure scope. | P0 |
| FR-POL-005 | High-risk actions MUST support separation of proposer, executor, verifier, and approver. | P1 |
| FR-POL-006 | Budget reservation and spend MUST be atomic at the authority that owns the budget. | P1 |
| FR-POL-007 | New work MUST fail closed after capability expiry or known revocation. | P0 |
| FR-POL-008 | Break-glass access MUST be narrow, strongly authenticated, time-limited, visibly audited, and reviewed afterward. | P2 |
| FR-POL-009 | Policy simulation MUST be available before activating a policy change that can deny or expand material authority. | P2 |

### Artifacts, evidence, review, and acceptance

| ID | Requirement | Priority |
|---|---|---|
| FR-EVI-001 | Immutable artifacts MUST have a verified content digest, media type, size, storage locator, producer, and provenance. | P0 |
| FR-EVI-002 | Evidence MUST state the claim or criterion tested, method, result, evaluator, artifacts, validity interval, and uncertainty/limitations. | P0 |
| FR-EVI-003 | Evidence from an agent's natural-language self-report alone MUST NOT satisfy deterministic or independent verification requirements. | P0 |
| FR-EVI-004 | Acceptance MUST reference the exact criteria, evidence, policy version, subject version, and authorized acceptor. | P0 |
| FR-EVI-005 | Acceptance SHOULD support criterion-level results and partial acceptance without marking the entire outcome complete. | P2 |
| FR-EVI-006 | Conflicting evidence MUST remain visible and route to a defined resolution process. | P1 |
| FR-EVI-007 | Evidence freshness and environmental validity MUST be evaluated where results can become stale. | P1 |
| FR-EVI-008 | Reviewer independence requirements MUST be enforceable by identity/operator/runtime relationships. | P2 |
| FR-EVI-009 | A later finding MUST be able to dispute or supersede acceptance without rewriting the original event. | P1 |

### Human-agent participation and attention

| ID | Requirement | Priority |
|---|---|---|
| FR-HA-001 | The system MUST measure human and agent effort/decisions by time, stage, scope, and task class without enforcing a target ratio. | P2 |
| FR-HA-002 | Autonomy MUST be configurable independently for observe, analyze, propose, decide, execute, verify, and accept stages. | P1 |
| FR-HA-003 | Risk or uncertainty thresholds MUST be able to narrow authority or request human attention. | P1 |
| FR-HA-004 | Human decision requests MUST include context, options, consequences, evidence, uncertainty, deadline, and safe default. | P1 |
| FR-HA-005 | Attention queues MUST support capacity, urgency, batching, quiet hours, delegation, and escalation. | P1 |
| FR-HA-006 | Approval requests MUST be meaningful; repeated low-value approvals SHOULD be sampled, delegated, or redesigned rather than normalized. | P2 |
| FR-HA-007 | A resumption capsule MUST summarize current outcome, active commitments, changes since last view, decisions, blockers, risks, and safe next actions. | P1 |
| FR-HA-008 | Users MUST be able to see and revoke an agent's current authority and stop its active runs. | P1 |
| FR-HA-009 | Participation and performance metrics MUST NOT become a hidden global reputation or punitive worker score. | P0 |

### Views and collaboration

| ID | Requirement | Priority |
|---|---|---|
| FR-UX-001 | The same authorized facts MUST support list, board, timeline, inbox, evidence, decision, and graph projections. | P1 |
| FR-UX-002 | The default view SHOULD emphasize outcomes, ready commitments, blockers, reviews, and exceptions. | P1 |
| FR-UX-003 | Graph views MUST progressively disclose detail, preserve stable layout where practical, and provide equivalent list/table navigation. | P2 |
| FR-UX-004 | Every status or recommendation MUST offer an explanation path appropriate to the user's authorization. | P1 |
| FR-UX-005 | Users MUST be able to distinguish observed, stated, inferred, simulated, disputed, stale, and accepted information. | P1 |
| FR-UX-006 | Comments and discussions MAY remain flexible, but only authorized commands can change durable domain state. | P1 |
| FR-UX-007 | Destructive or high-impact actions MUST show scope, affected objects, reversibility, and authority before execution. | P1 |
| FR-UX-008 | Notifications MUST be deduplicated, groupable, routable, and tied to an action or explicit informational purpose. | P1 |

### Offline and local-first behavior

| ID | Requirement | Priority |
|---|---|---|
| FR-LOC-001 | Authorized users SHOULD be able to read cached projections and create drafts without a network. | P2 |
| FR-LOC-002 | Mergeable document edits MUST converge after duplicate, reordered, or delayed delivery. | P2 |
| FR-LOC-003 | Offline actions affecting exclusive claims, budgets, grants, revocations, policy, or acceptance MUST remain proposals until accepted by the home authority. | P2 |
| FR-LOC-004 | Sync MUST preserve rejected local work and explain why it was not accepted. | P2 |
| FR-LOC-005 | Device loss and revocation MUST stop future synchronization and provide policy-driven remote key/session invalidation. | P2 |
| FR-LOC-006 | Local storage MUST respect workspace encryption, retention, and selective-cache policy. | P2 |
| FR-LOC-007 | Users MUST be able to export their authorized data in documented, usable formats without an active hosted subscription. | P1 |

### Federation

| ID | Requirement | Priority |
|---|---|---|
| FR-FED-001 | Every federated object MUST have a stable URI, home node, owning scope, authoritative version, and disclosure policy. | P3 |
| FR-FED-002 | A remote node MUST send proposals or accepted facts within its authority; it MUST NOT directly overwrite another node's home-owned object. | P3 |
| FR-FED-003 | Delivery MUST tolerate retry, duplication, reordering, delay, and temporary partition. | P3 |
| FR-FED-004 | Nodes MUST verify sender identity, signature, audience, schema version, replay window, and local policy before applying a message. | P3 |
| FR-FED-005 | Unknown, invalid, or suspicious messages MUST be quarantined with safe operator diagnostics. | P3 |
| FR-FED-006 | Federation MUST support key rotation, node suspension, peer blocking, object revocation, tombstones, and dispute. | P3 |
| FR-FED-007 | Shared contracts MUST expose only explicitly disclosed fields and MUST resist topology inference through counts, errors, and derived metrics. | P3 |
| FR-FED-008 | A participant MUST be able to leave a federation while preserving authorized history and respecting retention/deletion policy. | P3 |
| FR-FED-009 | Protocol capabilities and versions MUST be discovered and negotiated. | P3 |
| FR-FED-010 | Two independent implementations MUST pass conformance fixtures before federation is called stable. | P3 |

### Integrations and import/export

| ID | Requirement | Priority |
|---|---|---|
| FR-INT-001 | Connectors MUST act under named service principals with explicit scopes. | P1 |
| FR-INT-002 | Inbound events MUST be deduplicated and cursor/version gaps reconciled. | P1 |
| FR-INT-003 | Mappings MUST preserve source IDs, source authority, version, timestamp, provenance, and deletion state. | P1 |
| FR-INT-004 | Trail MUST show connector freshness and degraded/stale state. | P1 |
| FR-INT-005 | Import MUST support dry-run, mapping preview, identity resolution, validation errors, and restartable batches. | P2 |
| FR-INT-006 | Export MUST include schemas, events or authoritative snapshots, artifact manifests, identity bindings allowed by policy, and integrity digests. | P2 |
| FR-INT-007 | MCP/A2A adapters MUST expose bounded domain operations rather than direct database or arbitrary tool access. | P1 |
| FR-INT-008 | GraycodeAI adapters MUST preserve the ownership boundaries in the integration audit. | P1 |

### Audit, administration, and lifecycle

| ID | Requirement | Priority |
|---|---|---|
| FR-ADM-001 | Administrators MUST be able to inspect policy versions, grants, revocations, active leases, budget reservations, runs, federation peers, and connector health. | P1 |
| FR-ADM-002 | Audit views MUST reconstruct who knew, proposed, authorized, executed, verified, and accepted what, and when. | P1 |
| FR-ADM-003 | Audit export MUST redact content the auditor is not authorized to see while preserving verifiable event structure where possible. | P2 |
| FR-ADM-004 | Retention policy MUST distinguish domain metadata, content, artifacts, secrets, telemetry, local caches, and federation copies. | P1 |
| FR-ADM-005 | Schema and projection migrations MUST be restartable, observable, and reversible or restore-tested. | P1 |
| FR-ADM-006 | Projection rebuild MUST be supported from accepted events and versioned transformation code. | P0 |
| FR-ADM-007 | Deletion/redaction MUST preserve required integrity and legal/audit facts while removing content according to policy. | P2 |

## Nonfunctional requirements

### Security and privacy

| ID | Requirement/initial target |
|---|---|
| NFR-SEC-001 | Deny by default at every command, artifact, connector, tool, and federation boundary. |
| NFR-SEC-002 | Encrypt network traffic and sensitive persistent data; separate key management from application data. |
| NFR-SEC-003 | Do not place credentials, raw secrets, private keys, or sensitive tool output in durable events or model context. |
| NFR-SEC-004 | Support key and credential rotation without rewriting historical actor identity. |
| NFR-SEC-005 | Treat prompts, retrieved content, tool descriptions, artifacts, remote events, and agent messages as untrusted input. |
| NFR-SEC-006 | Record security-relevant actions in an append-only, tamper-evident audit stream with access controls. |
| NFR-SEC-007 | Minimize collection of worker behavior and content; provide visibility, purpose, retention, access, and contestability for monitoring. |
| NFR-SEC-008 | Threat models must be refreshed when tools, authority, memory, external inputs, federation, or autonomy change. |

### Reliability and consistency

| ID | Requirement/initial target |
|---|---|
| NFR-REL-001 | Acknowledged domain events survive one process crash and restart in the supported deployment profile. |
| NFR-REL-002 | Commands are idempotent within their documented retention window. |
| NFR-REL-003 | Projection lag is observable; safety-critical decisions read causally sufficient authoritative state. |
| NFR-REL-004 | Backups and restore are tested, including keys, event data, artifacts, and schema versions. |
| NFR-REL-005 | External delivery is at least once with deduplication; no claim of exactly-once side effects is permitted without a target-system proof. |
| NFR-REL-006 | Clock skew cannot extend authority; leases rely on authority-side time and fencing rather than client wall clocks. |

### Performance and scale

Initial targets are hypotheses to validate, not promises:

| ID | Target |
|---|---|
| NFR-PERF-001 | local navigation and cached reads respond within 100 ms at p95 on a reference laptop |
| NFR-PERF-002 | ordinary authoritative commands complete within 300 ms p95 excluding external policy/connector latency |
| NFR-PERF-003 | graph/list projection supports 100,000 visible domain objects per workspace without rendering the full graph |
| NFR-PERF-004 | projection rebuild processes at least 10,000 small events/second on a reference development machine |
| NFR-PERF-005 | federation backpressure prevents one peer from exhausting worker, storage, or review capacity |
| NFR-PERF-006 | scheduling and recommendation jobs have bounded input, time, memory, and cost budgets |

### Accessibility and usability

| ID | Requirement |
|---|---|
| NFR-A11Y-001 | Core web flows MUST target WCAG 2.2 AA. |
| NFR-A11Y-002 | Every graph function MUST have a keyboard and screen-reader usable list/table equivalent. |
| NFR-A11Y-003 | Status, risk, actor kind, and evidence state MUST NOT rely on color alone. |
| NFR-A11Y-004 | Reduced motion, high zoom, narrow viewport, and non-pointer input MUST be supported. |
| NFR-A11Y-005 | Explanations and denial messages MUST use plain language and include a safe next action when one exists. |

### Portability and operation

| ID | Requirement |
|---|---|
| NFR-OPS-001 | A single-node self-hosted deployment MUST be supported before distributed production profiles. |
| NFR-OPS-002 | Core functionality MUST operate without proprietary hosted dependencies. |
| NFR-OPS-003 | Data formats and APIs MUST be versioned and documented with compatibility policy. |
| NFR-OPS-004 | Metrics, traces, and logs MUST use stable correlation IDs and avoid content/secret leakage. |
| NFR-OPS-005 | Operators MUST have health, backlog, projection lag, lease, budget, connector, and federation delivery signals. |
| NFR-OPS-006 | Upgrades MUST have preflight checks, backup/restore guidance, and tested rollback or forward-recovery. |

## Domain invariants

1. Every accepted mutation has an authenticated actor, an idempotency key, an applicable policy decision, and a schema version.
2. A command is a request. Only an accepted event changes authoritative domain state.
3. An exclusive claim has at most one currently valid fencing token at its home authority.
4. A stale fencing token cannot authorize a later side effect even if its holder did not receive cancellation.
5. Reserved plus spent budget cannot exceed the authorized limit.
6. Acceptance references the exact subject version and the evidence required by the applicable policy version.
7. An artifact called immutable has a verified digest; a locator alone is insufficient.
8. Agent completion is a proposal until accepted by an authorized policy path.
9. Revoked or expired capability cannot authorize a new command.
10. A remote message is never trusted solely because it arrived over an authenticated connection.
11. No remote or local projection can become a second hidden write authority.
12. Private content is never included in an event, score, count, error, or federation message without a matching disclosure grant.
13. Historical events are not rewritten to resolve disagreement; correction uses dispute, supersession, redaction, or a compensating event.
14. A user can distinguish stale, inferred, simulated, and authoritative information.
15. Dynamic autonomy never transfers legal or policy authority beyond an explicit grant.

## Traceability and verification

Each requirement implemented later must link to:

- one or more user scenarios;
- its domain invariant or policy basis;
- protocol command/event/schema;
- automated test, model-check property, or usability/security evaluation;
- telemetry showing whether it works in production without exposing private content;
- the phase exit criterion that depends on it.

The first traceability matrix will be created after pilot scenarios are selected, because prioritization without a target user would produce false completeness.

