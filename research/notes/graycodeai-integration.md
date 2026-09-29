# GraycodeAI integration audit

**Observed:** 2026-09-28  
**Workspace:** `/Users/lakshmanpatel/Desktop/OSS2026/graycode-eco`

The detailed first pass is preserved in [the initial workspace map](../initial/graycodeai-workspace-map.md). This note turns it into integration requirements for Trail.

## Repository ownership

| Repository | Authority | Trail relationship |
|---|---|---|
| `rho/` | coding-agent conversation, tools, permissions, sessions, missions | participant and source of privacy-reduced graph/session events |
| `flux/` | provider transports, model catalog, streaming, retry, fallback, usage | remains behind agent runtimes; Trail does not call model vendors directly in v0 |
| `graycode-skills/` | skill/capability corpus | referenced for capability discovery through Rho or a registry adapter |
| `rover/` | isolated worktrees, execution DAGs, checks, evidence, bounded repair | execution provider after Trail authorizes a scoped run |
| `across/` | context packs, checkpoints, handoffs, memories, epistemic provenance | evidence and continuity provider; Trail stores references and disclosed claims |
| `trace/` | Git forge, issues, pull requests, reviews, CI, releases, signed mirrors | software-work connector and source of artifacts/evidence |
| `graycode-platform/` | company web, browser identity/BFF, organizations, billing, hosted ledger | optional hosted identity/control-plane integration through authenticated HTTP |
| `radius/` | uncommitted human-agent workspace plan | source of safety requirements; proposed to be superseded by Trail after an explicit repository decision |

## Observed constraints

1. The canonical `rho/ecosystem.yaml` currently lists Rho, Flux, Graycode Skills, Graycode Platform, Rover, and Across. Trace and Radius exist but are absent.
2. Only Rho has a compile-time dependency on Flux.
3. Graycode Platform's repository rules prohibit source dependency from other GraycodeAI projects. Trail integration must be HTTP unless that rule is changed in the Platform repository.
4. Rho's useful graph contract is under a Go `internal` package and cannot serve as Trail's public cross-repository API.
5. The existing graph envelope is narrower than Trail's intended commitments, claims, policies, disputes, delegation, revocation, and multi-party relations.
6. Trace and Across have existing local working-tree state that must not be modified by Trail setup work.
7. Radius has no committed implementation or remote. Its plan overlaps Trail and contains useful safety ideas: inbox delivery, idempotency, leases, credentials, approvals, and sandboxing.

## Required adapter contracts

### Rho

Inbound to Trail:

- actor/session identity and operator reference;
- proposed outcome, work decomposition, assumption, decision, or evidence reference;
- privacy-reduced run status and usage;
- `rho.graph/v1` import with an explicit semantic mapping.

Outbound from Trail:

- scoped work context;
- commitment and acceptance criteria;
- approved capability grant, time/budget limit, and disclosure policy;
- event callback or status endpoint.

Trail must not receive raw prompts, secrets, or full transcripts by default.

### Rover

Outbound from Trail:

- authorized run request;
- immutable plan version;
- repository/workspace reference;
- claim lease and fencing token;
- resource, cost, and time limits;
- checks and evidence policy;
- callback and cancellation identity.

Inbound to Trail:

- accepted/rejected dispatch receipt;
- run lifecycle events;
- checkpoint and heartbeat;
- artifacts and verification evidence;
- terminal result with failure classification.

Trail schedules; Rover executes. Trail accepts domain completion only after the applicable evidence/review policy passes.

### Across

- ingest immutable references to context packs, checkpoints, handoffs, and claims;
- preserve Across epistemic labels such as observed, stated, approved, inferred, disputed, superseded, and unknown;
- never rewrite an Across source claim silently;
- request selective bundles through an explicit disclosure boundary.

### Trace

- map repositories, issues, pull requests, reviews, commits, releases, checks, and identities to external references;
- use webhooks or an event cursor with deduplication and gap detection;
- preserve Trace as authority for Git state;
- create or update forge objects only through an authorized connector principal;
- link Trail acceptance to Trace evidence without pretending a merged pull request proves the business outcome.

### Graycode Platform

- authenticate Trail users and hosted nodes through documented HTTP APIs;
- keep billing/usage authority in Platform;
- avoid assuming Graycode Cloud project IDs are Trail workspace or outcome IDs;
- allow self-hosted Trail to operate without the hosted Platform;
- publish any shared protocol separately from Platform implementation code.

## Identity mapping

Each connector needs a mapping record:

```text
TrailPrincipal
  -> source system
  -> source principal ID
  -> verified binding method
  -> operator or organization
  -> allowed scopes
  -> valid interval
  -> revocation state
```

Do not merge identities based only on matching email, display name, model name, repository username, or agent label.

## Data ownership rule

Trail stores:

- stable external URI/ID;
- source system and authority;
- selected denormalized fields needed for a view;
- source version or cursor;
- disclosure classification;
- content digest when evidence integrity matters;
- mapping provenance.

Trail does not become authoritative for a sibling product's internal record. Connector projections must display freshness and degraded state.

## Failure behavior

| Failure | Required behavior |
|---|---|
| connector unreachable | keep Trail usable; mark projection stale; queue retry within policy |
| duplicate webhook | deduplicate without duplicating domain events or side effects |
| missed webhook | detect cursor/version gap and reconcile |
| source object deleted | retain a policy-compliant tombstone/reference without exposing removed content |
| identity mapping changes | create a new binding version; do not rewrite past actor attribution |
| source permission revoked | deny new access immediately after authoritative knowledge; remove cached content per retention policy |
| schema version unknown | quarantine, preserve envelope, and surface an operator action |
| source and Trail disagree | show source authority, mapping status, and dispute; do not use last-write-wins |

## Repository decision still required

`trail/` now contains the planned product boundary. `radius/` remains untouched. Before product implementation:

1. preserve any Radius planning material that is not already represented here;
2. explicitly decide whether Radius is archived, renamed, or retained for another scope;
3. initialize Trail's own repository only after the decision;
4. add Trail and Trace to the canonical ecosystem inventory through the appropriate repository change.

