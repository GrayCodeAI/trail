# Open questions and experiments

**Updated:** 2026-09-28

Questions are grouped by the decision they block. An answer needs evidence or an ADR; preference alone is insufficient for irreversible choices.

## Product

| ID | Question | Why it matters | Planned evidence | Blocks |
|---|---|---|---|---|
| OQ-P01 | Which user and workflow is the first paying or sustained adopter? | prevents a broad suite from replacing a sharp product | interviews and pilot commitment | Phase 1 scope |
| OQ-P02 | Do users understand “commitment,” or does another term work better? | this is the central interaction object | comprehension test | UI language and protocol naming |
| OQ-P03 | What is the smallest evidence packet that improves review? | excessive evidence creates attention debt | matched review study | default acceptance templates |
| OQ-P04 | Which work classes can operate at 90% agent effort safely? | dynamic autonomy needs concrete policies | task taxonomy and pilot data | autonomy presets |
| OQ-P05 | Should general chat exist in Trail or remain an integration? | chat can overwhelm durable state and duplicate other products | workflow prototype | product boundary |
| OQ-P06 | Which graph interactions are useful beyond ordinary list/board views? | graph UI has high complexity and accessibility cost | task-based usability test | visualization investment |

## Domain model

| ID | Question | Experiment/decision |
|---|---|---|
| OQ-D01 | Is `Commitment(debtor, creditor, antecedent, consequent)` expressive enough for joint and quorum commitments? | model real scenarios with relation nodes and compare complexity |
| OQ-D02 | Are request and offer durable objects or event sequences around commitments? | model negotiation, withdrawal, counteroffer, and expiry |
| OQ-D03 | How should conditional, alternative, quorum, and probabilistic dependencies serialize? | truth-table fixtures and human readability tests |
| OQ-D04 | Which facts require bitemporal representation in v0? | run correction, late evidence, backdated policy, and import scenarios |
| OQ-D05 | How are policy exceptions represented without normalizing unsafe bypasses? | formal state machine and governance review |
| OQ-D06 | Can acceptance be partial or scoped by criterion? | pilot with multi-artifact outcomes |

## Architecture

| ID | Question | Experiment/decision |
|---|---|---|
| OQ-A01 | PostgreSQL or embedded SQLite first? | benchmark local single-user, team server, event append, projection rebuild, and migration |
| OQ-A02 | PostgreSQL recursive queries or a separate graph projection? | measure representative reachability, cut sets, and rebuild cost before adopting a graph database |
| OQ-A03 | Yjs or Automerge for selected documents? | compare bundle size, offline sync, rich-text bindings, compaction, schema evolution, and non-JS support |
| OQ-A04 | Build a minimal durable runner or adapt Temporal? | compare operational weight, determinism constraints, timers, retries, cancellation, and self-hosting |
| OQ-A05 | Which policy engine fits relationship and contextual rules? | compare embedded code, Cedar, Rego/OPA, and Zanzibar-style systems against the policy corpus |
| OQ-A06 | What is the public protocol package location? | coordinate with Rho and ecosystem ownership; never depend on a sibling `internal` package |

## Federation

| ID | Question | Experiment/decision |
|---|---|---|
| OQ-F01 | Use ActivityStreams JSON-LD or a compact Trail JSON envelope? | implement canonical exchange both ways and compare interoperability, signature stability, and complexity |
| OQ-F02 | Which objects have home authority, and can home move? | migration/key continuity protocol exercise |
| OQ-F03 | How is right-to-erasure handled in append-only shared history? | legal/privacy review; encrypted payload deletion, tombstone, and redaction experiments |
| OQ-F04 | What metadata can be revealed for dependency readiness without exposing a private graph? | privacy threat model and inference tests |
| OQ-F05 | How are peer reputation and moderation scoped without a global score? | abuse scenarios and local policy model |
| OQ-F06 | What happens when a home node disappears permanently? | export, escrow, fork, and authority-transfer scenarios |

## Security and governance

| ID | Question | Experiment/decision |
|---|---|---|
| OQ-S01 | What principal binds an agent model, runtime, operator, and concrete run? | identity-chain prototype with Rho/Rover |
| OQ-S02 | What credential form supports attenuation and fast revocation? | compare short-lived OAuth tokens, macaroons/Biscuits, and service-minted capabilities |
| OQ-S03 | Which agent actions require an independent verifier? | risk-tier matrix and incident simulation |
| OQ-S04 | How can monitoring support safety without worker surveillance? | privacy impact assessment and user research |
| OQ-S05 | How are emergency stops authorized during identity-provider or policy-engine outage? | break-glass design and recovery drill |
| OQ-S06 | How can a compromised node rejoin after key rotation and audit? | node-compromise tabletop and protocol fixture |

## Business and ecosystem

| ID | Question | Decision needed |
|---|---|---|
| OQ-B01 | Is Trail open core, fully OSS, or protocol-open with hosted services? | licensing and business-model review before code reuse |
| OQ-B02 | Which license permits desired integrations while protecting the project? | legal review of AGPL, Apache-2.0, MPL-2.0, and dual licensing |
| OQ-B03 | Is the Trail name commercially clear under the GraycodeAI endorsement? | formal trademark and package-namespace review |
| OQ-B04 | Does Radius become an archive, rename, or distinct product? | explicit repository decision; preserve history and uncommitted files |
| OQ-B05 | Which optional account, organization, entitlement, and billing contracts should Trail consume from Graycode Platform? | ADR 0007 fixes independence; contract spike decides the bounded hosted integration |

## Decisions already strong enough to proceed

- use **Trail by GraycodeAI** as the working product identity;
- keep human-agent ratio dynamic and observed;
- make commitments and evidence central;
- use a signed event journal plus rebuildable projections;
- coordinate scarce/revocable rights at a home authority;
- make federation selective and defer it until local semantics work;
- integrate existing GraycodeAI products instead of merging their responsibilities;
- begin with a modular monolith and public protocol schemas;
- operate Trail as an independent product at `trail.graycodeai.com`, with optional Platform integration through versioned HTTP;
- exclude blockchain and token governance from the baseline.
