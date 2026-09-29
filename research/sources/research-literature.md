# Research literature map

**Reviewed:** 2026-09-28  
**Method:** primary papers, author archives, standards bodies, and official research pages

## How to read this map

The literature supplies design constraints and hypotheses. It does not prove that a product interaction will succeed. Every transfer from theory to Trail is marked as a proposed implication and must be tested in simulation or a pilot.

## Coordination and cooperative work

| Source | Finding used | Proposed Trail implication |
|---|---|---|
| Malone and Crowston, coordination theory | coordination manages dependencies among activities | model dependency type and coordination mechanism, not only predecessor links |
| Schmidt and Bannon, CSCW and articulation work | cooperative work includes work needed to make other work possible | record clarification, handoff, negotiation, exception handling, and review load |
| Star and Griesemer, boundary objects | collaborating groups can share an object while retaining local interpretations | shared outcomes and commitments need stable identity plus local projections |
| Hutchins, distributed cognition | cognition can be distributed across people, artifacts, and environment | the workspace must preserve situation, rationale, cues, and resumption context |
| Grosz and Kraus, [Collaborative plans for complex group action](https://doi.org/10.1016/0004-3702(95)00103-4) | collaborators can work from partial plans, contract subactions, and coordinate without one complete centralized plan | allow partial plan disclosure, delegation, and progressive elaboration |
| Singh, [A conceptual analysis of commitments in multiagent systems](https://repository.lib.ncsu.edu/items/a51477a9-ddb3-4c93-9c26-268ee5f2fb01) | external social commitments differ from an agent's private mental state | represent debtor, creditor, antecedent, consequent, lifecycle, discharge, cancellation, and violation |
| Chopra and Singh, [Clouseau](https://doi.org/10.1609/aaai.v34i05.6215) | commitment specifications can generate decentralized protocols with safety and liveness properties | test whether Trail commitment types can compile to valid interaction protocols |

## Requirements, decisions, and process

| Field/source | Finding used | Proposed Trail implication |
|---|---|---|
| goal-oriented requirements engineering | goals, obstacles, responsibilities, and refinements retain intent beyond task lists | keep outcomes, assumptions, acceptance measures, and blockers first-class |
| requirements traceability | rationale and links from need through implementation and verification reduce loss of intent | create typed provenance from outcome to decision, work, artifact, evidence, and acceptance |
| adaptive case management | knowledge work evolves as information becomes available | support constrained state transitions and exceptions instead of requiring one workflow template |
| process mining | event traces can reveal actual processes rather than intended diagrams | derive process views and conformance warnings from events; do not silently rewrite policy |
| BPMN and Petri-net analysis | explicit transitions enable reachability, liveness, and deadlock analysis | model-check core lifecycle and high-risk policy templates |
| decision analysis and Pareto optimality | objectives often conflict and cannot be reduced safely to one scalar | show feasible tradeoffs and provenance for weights; keep a human choice unless policy delegates it |
| value of information | gathering evidence is valuable when it can change a consequential decision | prioritize tests, clarification, and reviews by expected decision value |

## Distributed systems and local ownership

| Source | Finding used | Proposed Trail implication |
|---|---|---|
| Kleppmann et al., [Local-first software](https://www.inkandswitch.com/essay/local-first/) | offline operation, longevity, privacy, and user agency require local ownership rather than a subordinate cache | support local projections, drafts, portable export, and graceful server absence |
| Shapiro et al., [CRDT study](https://inria.hal.science/inria-00555588) | specific algebraic conditions allow replicas to converge without coordination | use CRDTs for mergeable documents and sets only when invariant preservation is proven |
| Hellerstein and Alvaro, [Keeping CALM](https://doi.org/10.1145/3369736) | monotonic computations can avoid coordination; non-monotonic decisions generally cannot | let observations accumulate locally; coordinate exclusion, revocation, budgets, approvals, and final acceptance |
| Burrows, [Chubby lock service](https://research.google/pubs/the-chubby-lock-service-for-loosely-coupled-distributed-systems/) | leases and lock services require explicit failure semantics and operational discipline | exclusive work claims use leases plus monotonic fencing tokens at a home authority |
| Garcia-Molina and Salem, [Sagas](https://www.cs.princeton.edu/research/techreps/598) | long-lived distributed work can be decomposed into transactions with compensations | external side effects require recorded compensating actions and human escalation when compensation is impossible |
| Matrix federation and Radicle | signed partial-order histories enable independent nodes but make conflict, moderation, and deletion harder | federate narrow shared contracts with home authority rather than whole private graphs |

## Graphs, topology, and sheaves

| Source | Finding used | Proposed Trail implication |
|---|---|---|
| directed acyclic graphs and partial orders | valid execution order is often partial, allowing safe concurrency | compute readiness from typed conditions rather than a total schedule |
| hypergraphs | many real dependencies relate several sources and targets | model relation nodes or hyperedges for quorum, joint production, and alternatives |
| strongly connected components | dependency cycles identify mutually blocked or jointly negotiated work | collapse cycles in views and require an explicit loop-breaking decision |
| max-flow/min-cut and cut sets | small dependency sets can constrain large portions of a network | suggest high-leverage unblock actions with uncertainty and sensitivity shown |
| Hansen and Ghrist, [Opinion Dynamics on Discourse Sheaves](https://arxiv.org/abs/2005.12798) | local representations can be translated at boundaries and disagreement measured after translation | preserve local vocabularies; measure inconsistency only where shared commitments require compatibility |
| Ghrist and Riess, [Cellular Sheaves of Lattices and the Tarski Laplacian](https://arxiv.org/abs/2007.04099) | lattice-valued local data and fixed points support distributed consistency reasoning | explore policy/claim compatibility and local-to-global coherence in simulation, not in the first protocol core |

Sheaf mathematics is a design discipline and later analysis tool. Trail's user-facing ontology should remain outcomes, commitments, evidence, policies, and boundaries.

## Flow, scheduling, and control

| Source/field | Finding used | Proposed Trail implication |
|---|---|---|
| Little's Law | average work in progress relates throughput and cycle time when stability assumptions hold | show measured WIP, throughput, and time with the observation window and caveats |
| queueing theory | high utilization causes nonlinear waiting time and variability matters | admission control must include review and human-attention capacity |
| critical path and PERT/Monte Carlo | uncertain durations produce a distribution of completion time and criticality | simulate schedule ranges; never present one deterministic date as truth |
| theory of constraints | system throughput is bounded by constraints that may move | identify likely constraints and re-evaluate after interventions |
| feedback control | delayed noisy feedback and overcorrection can destabilize a system | apply hysteresis, rate limits, confidence thresholds, and minimum dwell time to automated allocation |
| model-predictive control | optimize over a changing horizon while respecting constraints | later scheduling may replan within explicit budgets and safety limits |

## Collective choice and incentives

| Field | Finding used | Proposed Trail implication |
|---|---|---|
| Arrow and social-choice results | no rank aggregation rule satisfies every desirable property under broad conditions | never label a single vote-derived rank as the objectively correct priority |
| mechanism design | actors respond strategically to allocation and reputation rules | expose incentives, test gaming, and avoid universal reputation scores |
| quadratic voting/funding | intensity can be expressed with increasing marginal cost but depends on identity and budget integrity | allow only as an optional scoped decision method with Sybil resistance and clear stakes |
| Ostrom, polycentric governance | durable commons often use multiple nested centers of decision making | policies belong to scopes; federation does not imply one global administrator |

## Human factors and human-agent teaming

| Source/field | Finding used | Proposed Trail implication |
|---|---|---|
| Hollnagel and Woods, joint cognitive systems | automation should improve the joint system's ability to remain in control | evaluate team outcome and recovery, not isolated model accuracy |
| Bainbridge, ironies of automation | automation can leave humans with rare, difficult interventions and degraded skill | preserve situation awareness, rehearsal, meaningful approvals, and takeover paths |
| Parasuraman et al., levels of automation | information acquisition, analysis, decision, and action can have different automation levels | configure autonomy per stage and risk, never as one workspace-wide percentage |
| common-ground research | collaboration fails when participants incorrectly assume shared understanding | record assumptions, audience, terminology mappings, and acknowledgments |
| transactive memory | teams perform better when they know who knows what | model capabilities and evidence by scope without creating a global person score |
| interruption and resumption research | interruptions impose reconstruction cost | generate resumption capsules containing goal, state, changes, blockers, decisions, and safe next actions |
| psychological safety research | hidden punishment and surveillance suppress problem reporting | separate learning metrics from punitive scoring; make inference and monitoring visible and contestable |

## Security, identity, and provenance

| Source | Finding used | Proposed Trail implication |
|---|---|---|
| Pang et al., [Zanzibar](https://research.google/pubs/zanzibar-googles-consistent-global-authorization-system/) | relationship-based authorization needs consistency tied to object changes | policy evaluation must observe causally sufficient membership and grant state |
| Birgisson et al., [Macaroons](https://research.google/pubs/macaroons-cookies-with-contextual-caveats-for-decentralized-authorization-in-the-cloud/) | delegated credentials can be attenuated with contextual caveats | explore short-lived, narrowed capability grants for agent execution |
| [SPIFFE](https://spiffe.io/docs/latest/spiffe-specs/) | workloads can receive short-lived cryptographic identities inside trust domains | use workload identity for hosted runners; distinguish workload, agent, operator, and user identities |
| [W3C PROV-O](https://www.w3.org/TR/prov-o/) | entities, activities, agents, derivation, attribution, and delegation form an interoperable provenance core | align evidence export with PROV concepts while keeping Trail's domain vocabulary concise |
| [in-toto Attestation Framework](https://in-toto.io/docs/specs/) and SLSA | signed statements describe how software artifacts were produced | reuse attestation envelopes for software evidence rather than inventing incompatible formats |
| [NIST AI RMF](https://www.nist.gov/itl/ai-risk-management-framework) | AI risks need continuous governance, mapping, measurement, and management | risk records and evaluation are lifecycle concerns, not a launch checklist |
| [OWASP Agentic threats](https://genai.owasp.org/resource/agentic-ai-threats-and-mitigations/) | agent goals, tools, identities, memory, communication, and multi-agent cascades expand the attack surface | make least agency, untrusted-content separation, sandboxing, independent verification, and kill controls baseline requirements |

## Counter-evidence and limits

1. Local-first systems can create difficult key recovery, retention, moderation, device-loss, and selective-deletion problems.
2. More provenance can become surveillance or unusable noise. Capture must be scoped, redacted, and summarized without losing verifiability.
3. Event sourcing increases migration, privacy, rebuild, and operational complexity. The journal should contain domain facts, not every UI gesture or secret.
4. Formal models can create false confidence when the model omits social or organizational reality.
5. Optimization can destabilize work, reward metric gaming, and obscure value conflict.
6. Agent evaluation from historical outcomes can encode biased task assignment and become a punitive reputation system.
7. Federation distributes administration and failure. It does not automatically improve privacy, safety, availability, or governance.
8. Evidence verifies declared checks; it does not prove the requirement was correct or the artifact is harmless.

## Research-to-requirement rule

No mathematical mechanism or paper becomes a core feature from elegance alone. It must have:

1. a user problem;
2. a falsifiable hypothesis;
3. a simpler baseline;
4. measurable benefit;
5. failure and abuse tests;
6. an explanation understandable to the affected person;
7. a rollback or disable path.

