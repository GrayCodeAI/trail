# Trail risk register

**Updated:** 2026-09-28  
**Scale:** likelihood and impact are `low`, `medium`, `high`, or `critical`; values are pre-mitigation research estimates

## Product and adoption

| ID | Risk | Likelihood | Impact | Early signal | Treatment / decision gate |
|---|---|---:|---:|---|---|
| R-P01 | users perceive commitments/evidence as process overhead | high | critical | participants revert to comments and status fields | vocabulary/prototype tests; measure rework benefit before building execution |
| R-P02 | product expands into a weak clone of Jira, Slack, Notion, and GitHub | high | high | roadmap dominated by commodity suite features | enforce ADR 0001 and first-pilot job; integrate existing products |
| R-P03 | graph visualization becomes the product rather than a view | medium | high | core flows require canvas navigation | accessible list/table first; usability gate for every graph feature |
| R-P04 | no clear initial customer or pilot | high | critical | requirements grow without workflow evidence | Phase 0 Gate 0A blocks implementation |
| R-P05 | “decentralized” has no user value beyond self-hosting | medium | high | teams do not need cross-organization commitments | validate two-party scenario before federation investment |
| R-P06 | Trail name creates search/trademark conflict | medium | high | package/domain/trademark conflict appears | GraycodeAI endorsement and formal review before public launch |
| R-P07 | crowded agent-work products erase the differentiation | high | critical | roadmap converges on agents-as-assignees, chat, boards, and generic dashboards | enforce the evidence/acceptance wedge and require baseline comparison in the pre-build gate |
| R-P08 | YC/category attention is mistaken for customer demand | high | high | decisions cite trend reports, stars, or competitors without retained pilot use | require observed workflows, design partners, repeat use, and a buyer |

## Domain and correctness

| ID | Risk | Likelihood | Impact | Early signal | Treatment / decision gate |
|---|---|---:|---:|---|---|
| R-D01 | commitment model is too academic or incomplete | medium | critical | real scenarios need hidden exceptions/free text to function | scenario corpus, comprehension tests, formal lifecycle review |
| R-D02 | typed dependencies cost more to maintain than they save | high | high | stale edges and low use of readiness explanations | progressive typing; compare with binary/manual baseline |
| R-D03 | one event model cannot evolve safely | medium | critical | frequent breaking schema changes during prototype | delay production data; fixtures, upcasters/migrations, ADR 0002 gate |
| R-D04 | derived state has ambiguous authority | medium | high | readiness/violation differs across clients | specify one authoritative or explicitly projected source per fact |
| R-D05 | bitemporal history confuses users and developers | medium | medium | incorrect backdating/correction UX | keep advanced time fields behind clear workflows and fixtures |
| R-D06 | acceptance gives false assurance | high | critical | accepted outcomes fail in untested dimensions | bind scope/limitations; support dispute and post-acceptance evidence |

## Human-agent operation

| ID | Risk | Likelihood | Impact | Early signal | Treatment / decision gate |
|---|---|---:|---:|---|---|
| R-HA01 | dynamic autonomy is hard to understand | high | critical | users cannot state what an agent may currently do | stage-specific authority UI and comprehension gate |
| R-HA02 | approval fatigue produces rubber-stamping | high | critical | short repetitive approvals and escaped errors | batching, sampling, independent checks, redesign approval policy |
| R-HA03 | humans lose situation awareness and cannot take over | high | critical | slow/incorrect recovery after agent failure | resumption capsules, drills, checkpoint transparency, takeover metric |
| R-HA04 | agents create more review work than useful output | high | high | review queue grows faster than accepted throughput | admission control includes review capacity; accepted outcome per attention metric |
| R-HA05 | agent evaluation becomes worker surveillance or biased allocation | medium | critical | global scores or unexplained assignment disparities | task-scoped metrics, privacy/fairness review, no universal reputation |
| R-HA06 | multi-agent failures correlate through shared model/context | high | high | “independent” agents repeat same assumption | provenance diversity checks and truly independent verification |

## Distributed systems and data

| ID | Risk | Likelihood | Impact | Early signal | Treatment / decision gate |
|---|---|---:|---:|---|---|
| R-SYS01 | claims of exactly-once execution hide duplicate side effects | high | critical | retries create duplicate external objects/spend | explicit per-tool idempotency/reconciliation/compensation contracts |
| R-SYS02 | CRDT use violates authority or revocation invariants | medium | critical | offline acceptance/grant appears authoritative | strict mergeable/authority boundary and model/property tests |
| R-SYS03 | event journal conflicts with erasure/privacy obligations | medium | critical | sensitive payloads enter permanent history | payload separation/encryption/redaction; legal/privacy gate |
| R-SYS04 | projection drift produces wrong decisions | medium | high | rebuild differs from live state | deterministic reducers, shadow rebuild, cursor/version visibility |
| R-SYS05 | leases fail to stop stale external writes | medium | critical | enforcement target ignores fencing token | capability gateway validation at side-effect boundary |
| R-SYS06 | local-first sync/migration complexity delays product value | high | high | most engineering effort moves to replicas before pilot value | local-first follows useful single-node product; limit CRDT field set |
| R-SYS07 | graph/query scale drives premature separate database | medium | medium | architecture chosen from imagined scale | relational benchmark with representative data before graph store |

## Federation

| ID | Risk | Likelihood | Impact | Early signal | Treatment / decision gate |
|---|---|---:|---:|---|---|
| R-F01 | selective data still leaks private topology | high | critical | counts/timing/readiness reveal hidden work | inference threat tests and coarse signed boundary claims |
| R-F02 | node/key compromise poisons shared history | medium | critical | unexplained valid signatures or key rollback | rotation history, suspension, quarantine, compromise runbook |
| R-F03 | schema evolution fragments the network | high | high | peers quarantine normal updates after releases | negotiation, compatibility window, conformance fixtures, two implementations |
| R-F04 | home node disappearance strands shared objects | medium | high | parties cannot amend/close commitments | export and explicit authority-transfer/fork protocol |
| R-F05 | moderation conflicts with contractual history | medium | high | peer block is mistaken for commitment release | keep communication trust and domain obligation as separate state |
| R-F06 | federation operations overwhelm small self-hosted nodes | high | high | inbox backlog/storage/CPU grows from one peer | size limits, quotas, backpressure, peer blocking, narrow contract profile |

## Security and privacy

| ID | Risk | Likelihood | Impact | Early signal | Treatment / decision gate |
|---|---|---:|---:|---|---|
| R-SEC01 | prompt injection causes an authorized side effect | high | critical | action rationale traces to retrieved content | authority/content separation, least-agency gateway, adversarial tests |
| R-SEC02 | secret leaks through prompt, tool output, artifact, or trace | high | critical | scanner detects credential pattern | brokered credentials, structured redaction, egress and evidence scanning |
| R-SEC03 | compromised plugin/skill/adapter gains broad access | high | critical | manifest/update expands capability unexpectedly | signed/pinned supply chain, sandbox, reapproval on scope change |
| R-SEC04 | administrator can silently access all content | medium | critical | routine operations require content-level superuser | separate admin/content roles, JIT access, audited break-glass |
| R-SEC05 | policy outage or bad update grants or denies dangerously | medium | critical | broad denial/allow anomaly after policy change | fail closed, simulation, staged activation, recovery policy |
| R-SEC06 | signature/canonicalization differs across implementations | medium | critical | same event verifies in one language only | one canonical profile and cross-language adversarial fixtures |

## Ecosystem and delivery

| ID | Risk | Likelihood | Impact | Early signal | Treatment / decision gate |
|---|---|---:|---:|---|---|
| R-E01 | Trail duplicates or couples to sibling repository internals | high | high | imports from `internal` packages or shared database | public protocol/adapters; repository boundary review |
| R-E02 | Radius and Trail remain overlapping products | high | high | two roadmaps use same workspace/agent concepts | explicit disposition ADR before implementation |
| R-E03 | Graycode Platform becomes mandatory for self-hosted use | medium | high | core login/billing API dependency appears | independent identity/local deployment; Platform via optional HTTP |
| R-E04 | protocol design outruns implementation feedback | high | high | vocabulary grows without two concrete workflows | keep v0 minimal, executable fixtures, pilot-driven extensions |
| R-E05 | premature microservices consume the team | medium | high | cross-service failures before product fit | modular monolith gate and extraction criteria |
| R-E06 | OSS license choice blocks ecosystem or business plan | medium | critical | desired reuse/distribution conflicts appear | counsel/license decision in Gate 0C before copying code |
| R-E07 | MCP, A2A, AG-UI, or authorization protocol churn breaks adapters | high | high | core domain types begin mirroring a protocol draft | isolate adapters, pin negotiated versions, maintain fixtures, and keep Trail semantics independent |
| R-E08 | unclear OSS governance discourages adoption or creates control disputes | medium | high | contributions arrive before license and decision rights are understood | publish license matrix, governance, contribution, security, and trademark rules before public launch |

## Risk review cadence

- Phase 0: review after each research batch and vocabulary test.
- Build phases: review at design start, before pilot, and after incidents/major scope change.
- Security model refresh: whenever tools, authority, memory, external inputs, federation, identity, or autonomy changes.
- A critical risk without a named owner, treatment, verification method, and accepted residual exposure blocks the affected phase.
