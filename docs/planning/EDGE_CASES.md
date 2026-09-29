# Edge case and failure catalog

**Version:** 0.1  
**Date:** 2026-09-28  
**Status:** living catalog

No finite list can cover every failure in an open, adaptive, multi-agent distributed system. This catalog defines the minimum adversarial cases that must be resolved before each affected phase ships. Every case needs a test, model property, simulation, security exercise, or documented operational drill.

## Severity

- **C:** critical safety, authority, privacy, money, or irrecoverable integrity risk.
- **H:** high risk of incorrect work, lost data, broad outage, or material rework.
- **M:** degraded operation or confusing state with a recoverable path.
- **L:** minor friction or presentation issue.

## Identity and accountability

| ID | Case | Required behavior | Sev |
|---|---|---|---|
| EC-ID-01 | two principals have the same email or display name | keep separate identities until a verified binding is approved | H |
| EC-ID-02 | one person operates several agents | record distinct agent and run identities plus the accountable operator | H |
| EC-ID-03 | an agent changes model, prompt, tools, or runtime mid-run | version the execution profile; policy decides whether the run may continue | H |
| EC-ID-04 | service account is deleted after historical actions | preserve non-secret historical attribution and revocation state | M |
| EC-ID-05 | identity provider is unavailable | cached sessions follow bounded policy; no privilege expansion; break-glass is explicit | H |
| EC-ID-06 | a device is stolen while offline | revoke future sync/capabilities; encrypted local data and expiry limit exposure | C |
| EC-ID-07 | principal changes organization or role during a commitment | new commands use current authority; historical obligations and attribution remain | H |
| EC-ID-08 | remote node asserts an identity binding that conflicts locally | quarantine or keep separate; require an accepted trust path | C |
| EC-ID-09 | key rotates while messages are in flight | verify against key-validity interval and rotation proof; prevent rollback to old key | C |
| EC-ID-10 | operator and agent dispute who initiated an action | use signed/accepted delegation and run records; surface unresolved dispute | H |

## Outcomes, situations, and claims

| ID | Case | Required behavior | Sev |
|---|---|---|---|
| EC-OUT-01 | outcome criteria change after work starts | create a new version; show affected commitments and require renegotiation where material | H |
| EC-OUT-02 | outcome becomes impossible | record impossibility evidence, stop unsafe execution, and route release/renegotiation | H |
| EC-OUT-03 | two outcomes conflict through a shared resource or policy | expose conflict before scheduling and route an authorized decision | H |
| EC-OUT-04 | an observation later proves false | dispute/supersede it; recompute dependent views without erasing past decisions | H |
| EC-OUT-05 | information is valid in the past but recorded later | preserve valid time and transaction time | M |
| EC-OUT-06 | criteria are subjective or ambiguous | require named evaluator, rubric, examples, and allowed appeal/escalation | H |
| EC-OUT-07 | one criterion passes and another fails | keep criterion-level state; do not mark the whole outcome accepted | H |
| EC-OUT-08 | an outcome is duplicated during import | propose a mapping using source IDs/digests; never auto-merge on title alone | M |

## Negotiation and commitments

| ID | Case | Required behavior | Sev |
|---|---|---|---|
| EC-COM-01 | request is assigned but provider never accepts | remain a request/proposal, never an active commitment | H |
| EC-COM-02 | offer and withdrawal cross in flight | home-authority ordering/version rules decide; preserve both events and notify parties | H |
| EC-COM-03 | counteroffer references a stale version | return structured conflict with current version and preserved draft | M |
| EC-COM-04 | commitment deadline passes during a partition | mark locally uncertain; home authority records violation/extension after reconciliation | H |
| EC-COM-05 | debtor delegates to an agent then revokes delegation | agent loses future authority; underlying commitment remains unless separately changed | C |
| EC-COM-06 | creditor releases a commitment after provider completed work | preserve evidence and release timing; apply compensation/payment policy | H |
| EC-COM-07 | both parties believe the other must act next | detect unsatisfied antecedents/reciprocal wait and route negotiation | H |
| EC-COM-08 | a joint commitment loses one participant | policy defines replacement, reduced quorum, renegotiation, or failure | H |
| EC-COM-09 | party refuses an automated interpretation of terms | freeze automatic action and route the original accepted representation plus dispute | C |
| EC-COM-10 | commitment is legally or ethically invalid despite technical acceptance | policy can suspend/void future action while preserving audit and appeal | C |

## Dependencies and planning

| ID | Case | Required behavior | Sev |
|---|---|---|---|
| EC-DEP-01 | direct or indirect dependency cycle | detect SCC, block false readiness, and request a loop-breaking decision | H |
| EC-DEP-02 | dependency is deleted while downstream work runs | version impact, re-evaluate safety, and require policy decision for active run | H |
| EC-DEP-03 | any-of dependencies all fail | mark unsatisfied and identify new alternatives or escalation | M |
| EC-DEP-04 | quorum membership changes during evaluation | evaluate against the accepted policy/version; do not mix sets | H |
| EC-DEP-05 | dependency result becomes stale | expire readiness and notify active claims according to risk | H |
| EC-DEP-06 | two work items appear independent but share an external resource | require declared resource/write scope; detect conflicts where connector can report them | H |
| EC-DEP-07 | inferred dependency has low confidence | label inferred/confidence and prevent it from silently becoming a hard gate | M |
| EC-DEP-08 | huge graph creates an unreadable or expensive view | summarize by outcomes/SCCs/cut sets, paginate traversal, and preserve list equivalent | M |
| EC-DEP-09 | priority changes faster than work can switch | use minimum dwell time/hysteresis and count switching cost | M |
| EC-DEP-10 | optimizer recommends starvation of low-value maintenance | enforce fairness/risk/age constraints and show the tradeoff | H |

## Claims, leases, and concurrency

| ID | Case | Required behavior | Sev |
|---|---|---|---|
| EC-LEASE-01 | two actors claim exclusive work concurrently | home authority accepts one fencing token and rejects/conflicts the other | C |
| EC-LEASE-02 | worker continues after lease expiry | downstream gateways reject stale fencing token before side effect | C |
| EC-LEASE-03 | network partition prevents renewal but work is safe to preserve | checkpoint locally; stop new protected side effects; later propose/rebase results | H |
| EC-LEASE-04 | renewal response is lost | retry with same idempotency key; return current lease rather than issue a new token | C |
| EC-LEASE-05 | authority clock and worker clock differ | authority time defines lease; worker treats local time only as an early safety margin | C |
| EC-LEASE-06 | lease owner crashes after external side effect | reconcile target system using idempotency key/receipt before retry | C |
| EC-LEASE-07 | shared work does not require exclusivity | use nonexclusive participation/commitment rather than unnecessary locking | M |
| EC-LEASE-08 | administrator forcibly reassigns active work | fence old actor, preserve checkpoint, and record explicit override/risk | C |

## Agent execution and tools

| ID | Case | Required behavior | Sev |
|---|---|---|---|
| EC-RUN-01 | agent claims success but process exits nonzero or checks fail | record execution result; completion proposal fails evidence policy | H |
| EC-RUN-02 | process succeeds but produces no required artifact | remain unaccepted and request evidence/change | H |
| EC-RUN-03 | agent loops or recursively delegates | enforce step, time, token, cost, depth, and fan-out budgets | C |
| EC-RUN-04 | prompt/tool output injects instructions | keep data and authority channels separate; policy/tool gateway validates every action | C |
| EC-RUN-05 | tool schema changes during a run | pin tool version/capability contract or stop with structured incompatibility | H |
| EC-RUN-06 | model/provider becomes unavailable | follow approved fallback policy; record changed model or pause if profile matters | M |
| EC-RUN-07 | fallback model lacks required capability or data policy | deny fallback and request another plan | H |
| EC-RUN-08 | agent requests broader permission mid-run | create a separate approval request; current run remains bounded | C |
| EC-RUN-09 | cancellation races with a side effect | gateway uses fencing/idempotency; result becomes cancelled-with-effects and reconciles | C |
| EC-RUN-10 | run creates an unbounded volume of logs/artifacts | enforce quotas, truncate safely, preserve digests/diagnostics, and stop if evidence would be invalid | H |
| EC-RUN-11 | agent edits outside declared write set | sandbox/gateway blocks it; record violation and evaluate compromise | C |
| EC-RUN-12 | two agents produce valid but incompatible results | retain both branches, run integration evidence, and route selection/merge | H |
| EC-RUN-13 | human takes over an agent run | checkpoint, transfer claim/capabilities explicitly, and preserve attribution by action | H |
| EC-RUN-14 | agent runtime compromised | quarantine run/artifacts, revoke credentials, invalidate trust, and require independent re-verification | C |

## Budgets and resources

| ID | Case | Required behavior | Sev |
|---|---|---|---|
| EC-BUD-01 | two runs reserve the last budget concurrently | atomic authority accepts only reservations within limit | C |
| EC-BUD-02 | provider reports delayed or corrected cost | record adjustment, alert overrun, and update future authority; never rewrite original usage | H |
| EC-BUD-03 | currency/unit changes | preserve unit and conversion source/time; avoid implicit aggregation | H |
| EC-BUD-04 | cancellation avoids future cost but incurred cost remains | release only unused reservation; retain actual spend | M |
| EC-BUD-05 | shared resource capacity drops suddenly | pause/admit according to policy and replan without overbooking | H |
| EC-BUD-06 | human review queue is saturated | stop admitting work that requires that review, batch safely, or use an approved alternative | H |
| EC-BUD-07 | budget authority is unavailable | deny new spend; allow only explicitly pre-reserved bounded work | C |

## Evidence, review, and acceptance

| ID | Case | Required behavior | Sev |
|---|---|---|---|
| EC-EVI-01 | artifact changes at same URL | digest mismatch invalidates reference; preserve original expected digest | C |
| EC-EVI-02 | artifact is unavailable but digest remains | show unverifiable availability; acceptance policy decides whether cached/replicated copy is enough | H |
| EC-EVI-03 | verifier is the same agent/operator when independence is required | reject policy satisfaction and route an independent verifier | C |
| EC-EVI-04 | deterministic checks pass but semantic reviewer rejects | preserve both results; acceptance remains unresolved/rejected per policy | H |
| EC-EVI-05 | reviewers disagree | apply named quorum/tie/escalation policy, never average away the disagreement | H |
| EC-EVI-06 | evidence was valid but environment has changed | expire or scope validity and request re-verification | H |
| EC-EVI-07 | acceptance is issued against a stale subject version | reject with conflict; reviewer must inspect current version | C |
| EC-EVI-08 | malicious artifact exploits reviewer tooling | scan/sandbox preview and treat artifact as untrusted | C |
| EC-EVI-09 | evidence contains secrets or personal data | quarantine/redact under policy, preserve safe integrity metadata, notify owner | C |
| EC-EVI-10 | accepted result later causes harm | open dispute/incident, supersede status where authorized, preserve original acceptance and evidence | C |
| EC-EVI-11 | reviewer rubber-stamps repeated requests | measure review quality/sampling and redesign policy; do not treat click count as assurance | H |

## Policy, capability, and governance

| ID | Case | Required behavior | Sev |
|---|---|---|---|
| EC-POL-01 | policy changes between command submission and commit | evaluate against a causally current policy or return conflict; record version used | C |
| EC-POL-02 | revocation races with queued command | ordering at authority decides; no execution after accepted revocation | C |
| EC-POL-03 | capability is valid but purpose/condition is not | deny with scoped reason | C |
| EC-POL-04 | delegated capability exceeds parent | reject derivation and audit attempt | C |
| EC-POL-05 | policy engine fails or times out | fail closed for mutation; allow safe cached reads under policy | C |
| EC-POL-06 | policy would lock out every administrator | preflight simulation and protected recovery path prevent activation | C |
| EC-POL-07 | break-glass used during outage | bound action, notify designated parties, and require post-event review | C |
| EC-POL-08 | policy reason leaks private membership/resource existence | return disclosure-safe denial; keep detailed reason in authorized audit | H |
| EC-POL-09 | democratic vote attempts to authorize a restricted action | voting can recommend only; capability/policy authority remains binding | C |
| EC-POL-10 | malicious majority changes local governance | honor scope and constitutional constraints; support fork/export/appeal where defined | C |

## Offline and synchronization

| ID | Case | Required behavior | Sev |
|---|---|---|---|
| EC-OFF-01 | two users edit the same text offline | CRDT merge, per-user undo, and visible attribution/version history | M |
| EC-OFF-02 | offline user accepts work with stale authority | store as proposal; home authority revalidates and may reject | C |
| EC-OFF-03 | local draft references deleted object | preserve draft and offer retarget/recovery, never silently discard | M |
| EC-OFF-04 | client is offline beyond schema compatibility window | require upgrade/export path; preserve readable local data | H |
| EC-OFF-05 | sync sends updates twice or out of order | operations converge/idempotently apply | H |
| EC-OFF-06 | local storage quota is exhausted | stop caching safely, preserve pending writes, explain recovery | H |
| EC-OFF-07 | device clock is far wrong | use logical/authority ordering; display clock anomaly | M |
| EC-OFF-08 | a revoked user reconnects with old local content | reject unauthorized writes/reads; apply retention and cache-removal protocol | C |
| EC-OFF-09 | CRDT state contains data later subject to erasure | compact/rekey/redact according to tested policy; do not promise impossible deletion | C |

## Federation and trust domains

| ID | Case | Required behavior | Sev |
|---|---|---|---|
| EC-FED-01 | duplicate/replayed signed event | deduplicate by event ID and reject outside replay policy | C |
| EC-FED-02 | events arrive causally out of order | hold pending dependencies or request missing facts; never invent order | H |
| EC-FED-03 | remote schema version is unknown | quarantine, preserve envelope, negotiate capability, and expose operator action | H |
| EC-FED-04 | signature valid but sender lacks domain authority | reject based on object home/relationship/local policy | C |
| EC-FED-05 | remote node is compromised | suspend peer/key, quarantine new events, assess previously accepted facts | C |
| EC-FED-06 | sender key rotates without a trusted continuity proof | stop acceptance and require re-establishment of trust | C |
| EC-FED-07 | home node is permanently gone | use documented export/fork/authority-transfer process; never elect authority implicitly | H |
| EC-FED-08 | two nodes claim to be home after migration | verify signed transfer/version chain; quarantine split-brain mutations | C |
| EC-FED-09 | remote object is deleted or redacted | apply disclosure/retention policy, tombstone references, and stop serving removed content | H |
| EC-FED-10 | peer floods inbox or expensive signatures | rate limit, quota, backpressure, and block without harming local work | C |
| EC-FED-11 | error response reveals private graph topology | use constant/safe errors and privacy-tested metadata | C |
| EC-FED-12 | shared commitment depends on private internal work | expose only agreed readiness/status/evidence claims, never child topology | H |
| EC-FED-13 | nodes disagree on time or deadline | use contract-defined time authority and signed timestamps; surface ambiguity | H |
| EC-FED-14 | peer blocks another while shared work remains | freeze new exchange and route closure, export, or dispute policy | H |
| EC-FED-15 | malicious peer sends valid but harmful content | content remains untrusted; scan/sandbox and apply local moderation | C |

## Connectors and external systems

| ID | Case | Required behavior | Sev |
|---|---|---|---|
| EC-CON-01 | webhook is duplicated | idempotent ingest; no duplicate notification, spend, or commitment | H |
| EC-CON-02 | webhook delivery gap | detect cursor/version discontinuity and reconcile from source | H |
| EC-CON-03 | source object changes while Trail command is pending | compare source version; reject or rebase explicitly | H |
| EC-CON-04 | two-way sync creates an update loop | tag origin/correlation and suppress reflected no-op changes | H |
| EC-CON-05 | source API rate limits or changes | backoff, surface stale state, and preserve local core operation | M |
| EC-CON-06 | source permissions are revoked | stop new fetch/write and remove cached data according to policy | C |
| EC-CON-07 | connector secret rotates | resume without changing connector identity or losing cursor | H |
| EC-CON-08 | source deletes an artifact after acceptance | preserve authorized digest/provenance and show availability loss | H |
| EC-CON-09 | external side effect succeeds but connector times out | reconcile using idempotency/source query before retry | C |
| EC-CON-10 | imported identity cannot be resolved | create an explicit unresolved principal; never assign to a guessed user | H |
| EC-CON-11 | mapped fields have different semantics | preserve source field and mapping version; mark lossy transformation | M |

## Time, ordering, and history

| ID | Case | Required behavior | Sev |
|---|---|---|---|
| EC-TIME-01 | daylight-saving or timezone change crosses deadline | store instants plus intended timezone/calendar semantics | M |
| EC-TIME-02 | event recorded late or backdated | preserve recorded time and valid time; prevent unauthorized backdating | H |
| EC-TIME-03 | concurrent events have no total order | keep partial order; order only where authority/invariant requires it | M |
| EC-TIME-04 | policy or schema effective date is in future | activate at defined authority time; show pending change | M |
| EC-TIME-05 | legal hold conflicts with normal deletion time | hold only required scope, audit access, and release when authority ends | C |
| EC-TIME-06 | history becomes too large for replay | snapshot/projection checkpoints remain verifiable against event ranges | H |

## Storage, migration, and recovery

| ID | Case | Required behavior | Sev |
|---|---|---|---|
| EC-DATA-01 | process crashes between domain write and message publish | transactional outbox ensures eventual delivery after committed event | C |
| EC-DATA-02 | projection update fails | authoritative event remains; retry/rebuild and show lag | H |
| EC-DATA-03 | projection code has a bug | version projections, rebuild corrected view, retain audit of repair | H |
| EC-DATA-04 | schema migration stops midway | restart or restore safely; writes follow explicit compatibility mode | C |
| EC-DATA-05 | backup restores data but not keys | restore drill treats keys/trust metadata as required or documents irrecoverable encryption | C |
| EC-DATA-06 | artifact store and database disagree | reconcile manifests/digests; never claim evidence available until verified | H |
| EC-DATA-07 | corrupted event detected | stop unsafe replay, isolate range, verify signatures/checksums, recover from replica/backup | C |
| EC-DATA-08 | unknown event type appears during rebuild | preserve and stop/skip only under explicit forward-compatibility rule | H |
| EC-DATA-09 | tenant/workspace data crosses boundary through cache/index | treat as security incident; isolate indexes and test authorization at query result | C |

## User experience and accessibility

| ID | Case | Required behavior | Sev |
|---|---|---|---|
| EC-UX-01 | color-blind or screen-reader user needs dependency state | redundant text/icon semantics and accessible table/path explanation | H |
| EC-UX-02 | large graph causes motion or cognitive overload | progressive disclosure, reduced motion, stable layout, focus mode | M |
| EC-UX-03 | optimistic UI command is rejected later | revert authoritative display while preserving the user's draft and reason | H |
| EC-UX-04 | user approves the wrong scope | show concrete affected resources and require stronger confirmation for material expansion | C |
| EC-UX-05 | too many notifications hide a critical exception | severity/action routing, deduplication, batching, and escalation | H |
| EC-UX-06 | agent and human edit same description | collaborative merge with clear attribution and local undo | M |
| EC-UX-07 | explanation requires private data the viewer lacks | provide a safe partial explanation and authorized escalation path | H |
| EC-UX-08 | localized text changes meaning of policy or status | stable machine semantics; reviewed translations and source-language reference | H |
| EC-UX-09 | human resumes after long absence | resumption capsule shows changes, invalidated assumptions, and safe next step | H |

## Safety, abuse, and social failure

| ID | Case | Required behavior | Sev |
|---|---|---|---|
| EC-SAF-01 | manager uses telemetry for hidden worker ranking | prohibit undisclosed derived scores; purpose/access/appeal controls | C |
| EC-SAF-02 | agent fabricates evidence or cites nonexistent artifact | digest/availability verification and independent checks | C |
| EC-SAF-03 | several agents reinforce the same false assumption | preserve source diversity, require independent evidence, and detect shared provenance | C |
| EC-SAF-04 | malicious user creates approval fatigue | rate limits, batching, sampling, and policy redesign | H |
| EC-SAF-05 | model output includes harassment or sensitive inference | moderation and redaction policy; do not turn inference into accepted fact | H |
| EC-SAF-06 | optimization shifts all undesirable work to one person | fairness/rotation constraints and transparent workload impact | H |
| EC-SAF-07 | participants game reported confidence or priority | compare calibration/outcomes, preserve incentives, and avoid automatic universal scoring | H |
| EC-SAF-08 | emergency stop itself is abused | separation of duties, scoped stop, strong authentication, audit, and recovery path | C |
| EC-SAF-09 | compliance request conflicts with participant privacy | apply jurisdiction/workspace policy, data minimization, and auditable disclosure | C |

## Exit rule

An affected phase cannot ship while a critical case lacks all of:

1. an explicit domain result;
2. an enforcing component and trust boundary;
3. an automated property/test where feasible;
4. operator and user-visible recovery behavior;
5. an observability signal;
6. a documented residual risk owner.

