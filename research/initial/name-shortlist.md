# Product name shortlist

**Prepared:** 2026-09-28  
**Naming target:** decentralized, adaptive coordination among humans, agents, services, and organizations

## Decision: Trail, a GraycodeAI product

The product name is **Trail**. **GraycodeAI** is the parent brand.

- A trail is the time-ordered path from intent through commitments, actions, evidence, review, and acceptance.
- It fits a platform where humans and agents leave verifiable provenance while moving through a changing graph.
- It remains understandable to nontechnical users while the architecture can retain its deeper sheaf, graph, causal, and control-theory foundations.
- The GraycodeAI endorsement resolves most repository, package, search, and product-family ambiguity created by the crowded standalone word.

Recommended presentation:

> **Trail**  
> Every outcome has a path.
>
> A GraycodeAI product

Use **Trail** throughout the interface and documentation. Use **Trail by GraycodeAI** in search metadata, launch material, legal pages, package descriptions, and other contexts where the parent brand supplies identity.

## Mathematical foundation: cellular sheaves

**SheafWork** remains the strongest standalone alternative when the brand itself must express the mathematical foundation.

- In applied topology, a **cellular sheaf on a graph** attaches local state to vertices and edges plus maps that describe when neighboring views are compatible.
- A **global section** is a coherent assignment across the network. That is a useful model for independent teams and agents coordinating without surrendering all local state to one central plan.
- A **sheaf Laplacian** measures disagreement after translating different local representations into a shared boundary context.
- “Work” makes the product category clear. The name does not restrict the product to tickets, boards, software teams, or a fixed human-to-agent ratio.

Recommended presentation:

> **SheafWork**  
> Local autonomy. Global coherence.

Plain-language descriptive line:

> Coordinate people and agents through outcomes, commitments, and evidence.

### The mathematical idea behind the name

Let each participant or workspace vertex `v` hold local state `x_v`. For an edge `e = (u,v)`, restriction maps translate the two local states into the context shared on that edge. A simple disagreement energy is

```text
E(x) = Σₑ ‖Rᵤ→ₑ xᵤ − Rᵥ→ₑ xᵥ‖² = xᵀ L_F x
```

Low energy means the local states agree where agreement is required. It does **not** require every participant to store the same state, use the same ontology, or have the same authority. This is the exact design principle needed for decentralized work.

The platform does not need to expose sheaf terminology in the product interface. It can present outcomes, commitments, policies, evidence, and exceptions while using the mathematics as a design discipline.

## How graph theory, mathematics, and physics should shape the product

| Foundation | Product interpretation |
|---|---|
| typed temporal hypergraph | work includes multi-party, time-varying relations; one flat task graph is insufficient |
| partial orders | dependencies constrain valid execution order without demanding one total plan |
| cellular sheaves | local vocabularies and permissions can remain local while boundary conditions are checked |
| Laplacian energy | quantify disagreement at coordination boundaries; do not collapse every difference into an error |
| constraints and invariants | safety, budget, authority, privacy, and acceptance rules remain true as participation changes |
| control theory | observe, estimate, plan, act, measure, and correct through feedback |
| queueing theory | admission control and concurrency respond to load, risk, and reviewer capacity |
| causal graphs | decisions and evidence record what changed what, with time-respecting provenance |
| statistical mechanics analogy | local interactions can produce global coordination; treat this as an analogy, not a physical law |
| Noether-inspired design | allow behavior and topology to change while preserving declared governance invariants |

The human/agent ratio should be an observed function, not a product constant:

```text
r(t, scope) = agent effort / (agent effort + human effort)
```

It may be 80/20, 90/10, 10/90, or change during one workflow. Authority, risk, confidence, and policy determine who may act at each moment.

## Namespace and collision strategy

The standalone word **Trail** is already used by work-management, agent, governance, and developer products. Use the Graycode namespace consistently on every public and technical surface:

| Surface | Name |
|---|---|
| product name | **Trail** |
| navigation label | **Trail** |
| endorsed form | **Trail by GraycodeAI** |
| command line | `graycode trail` |
| repository | `graycodeai/trail` |
| npm | `@graycode/trail` |
| Python distribution | `graycode-trail` |
| protocol identifier | `graycode.trail` |

Avoid claiming the unqualified `trail` package, executable, domain, or repository name. Namespacing reduces product confusion, but formal trademark review remains necessary before a commercial launch.

## Scientific shortlist

| Rank | Name | Scientific basis | Product strength | Preliminary collision note |
|---:|---|---|---|---|
| 1 | **Trail** | causal paths, temporal graphs, and verifiable provenance | clearest product name inside the established GraycodeAI family | use the GraycodeAI endorsement where identity or discovery matters |
| 2 | **SheafWork** | cellular sheaves, graph Laplacians, local-to-global consistency | best standalone mathematical name | `.com`, GitHub, npm, and PyPI appeared open |
| 3 | **NoetherFlow** | symmetry, invariants, conservation, dynamic flow | strongest physics identity; good for a technical developer audience | `.com`, GitHub, npm, and PyPI appeared open |
| 4 | **Causalattice** | causal order plus lattice-valued local decisions | highly distinctive and closely related to policies and distributed agreement | `.com`, GitHub, npm, and PyPI appeared open; spelling needs teaching |
| 5 | **SheafFlow** | local-to-global consistency plus dynamic work flow | mathematically accurate and more abstract than SheafWork | `.com`, GitHub, npm, and PyPI appeared open |

### Physics-first alternative: NoetherFlow

Emmy Noether's theorem connects continuous symmetries with conserved quantities. The product analogy is useful: teams, agents, workflows, and graph topology may change, while declared invariants such as authority, budget, privacy, and auditability remain conserved.

Use this name if the initial audience is mainly developers and infrastructure teams. Pronunciation and spelling add friction for a broad business audience.

Suggested line:

> **NoetherFlow**  
> Dynamic work. Preserved guarantees.

### Graph-first alternative: Causalattice

This coined name combines **causal** and **lattice**. Causal structure captures time-respecting dependencies. Lattices model partial information, policy joins and meets, and monotone convergence toward decisions.

It is distinctive and compact, but people may initially parse or spell it incorrectly.

Suggested line:

> **Causalattice**  
> Structure every action. Trace every outcome.

## Names rejected or downgraded after screening

| Name | Reason |
|---|---|
| Sheaf | active Sheaf Labs agent-governance platform, plus current CRM, SaaS-management, model-serving, and programming-language products; `.com`, npm, and PyPI are registered |
| GraphDone-like compounds | make the graph the identity of the product rather than its underlying model |
| CausalWeave | used publicly since 2025 as a named AI architecture or concept |
| PhaseWeave | title of a 2026 ISCA computer-architecture system and research artifact |
| HyperWeave | active agent-visualization product, package, and trademark usage |
| EigenWork | active technical hiring platform |
| FieldGraph | active GIS and CAD products |
| InvariantFlow | `.com` is registered and a same-name GitHub repository exists |
| NodalFlux | `.com` is registered and a same-name GitHub repository exists |
| Causara | multiple active AI and decision-intelligence products |
| Meldwork | active multi-agent workspace and other collaboration usage |
| Knotwork | several current AI automation and agent products |
| Pactora | several active contract-intelligence products |
| Cohora | active ecommerce platform and an agent-collaboration project |
| Praxora | multiple active AI, productivity, consultancy, and trademark uses |
| Sympraxis | registered workflow-management software trademark |
| IntentMesh | several closely related agent-governance projects and companies |
| AccordMesh | current AI meeting-assistant repository and product usage |
| CommitWeave | existing Git tool, VS Code extension, and mobile app |

## Naming system

- parent brand: **GraycodeAI**;
- product: **Trail**;
- endorsed form: **Trail by GraycodeAI**;
- self-hosted coordination server: **Trail Node**;
- interoperability specification: **Trail Protocol**;
- command line: `graycode trail`;
- repository: `graycodeai/trail`.

Avoid separately branding the scheduler, policy service, agent runtime, and database until users encounter them as independent products.

## Voice and visual direction

- Show local regions connected through small shared boundaries, rather than a generic node-link graph.
- Let regions change size and connection density to express variable human and agent participation.
- Use flow lines only where they communicate causality, load, or state transition.
- Avoid atoms, quantum imagery, robot heads, checkboxes, kanban columns, blockchain cubes, and decorative equations.
- Lead public language with work, outcomes, commitments, evidence, and agreement. Explain sheaf theory in technical material.

## Candidate taglines

1. **Every outcome has a path.**
2. **Follow work from intent to evidence.**
3. **Where people and agents move work forward.**
4. **Coordinate every step. Verify every outcome.**
5. **The verifiable path through work.**

The recommended launch combination is:

> **Trail**  
> **Every outcome has a path.**  
> A GraycodeAI product

## Research basis

- Hansen and Ghrist, [Toward a Spectral Theory of Cellular Sheaves](https://arxiv.org/abs/1808.01513): cellular sheaves extend spectral graph theory and support consensus, sparsification, and effective-resistance analysis.
- Ghrist and Riess, [Cellular Sheaves of Lattices and the Tarski Laplacian](https://arxiv.org/abs/2007.04099): local-to-global constraints, consensus, and distributed optimization over networks.
- Hansen and Ghrist, [Opinion Dynamics on Discourse Sheaves](https://arxiv.org/abs/2005.12798): networked agents can communicate through different local representations and converge through sheaf diffusion.
- Oxford researchers, [Consensus dynamics on temporal hypergraphs](https://ora.ox.ac.uk/objects/uuid%3A51092ab7-2904-481d-8ff4-6825ce1555f4): time-dependent multiway interactions behave differently from their static pairwise projections.
- U.S. Department of Energy, [Symmetry in Physics](https://www.energy.gov/science/doe-explainssymmetry-physics): a concise account of Noether's connection between continuous symmetries and conservation laws.
