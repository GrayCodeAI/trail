# Trail Protocol v0

**Identifier:** `graycode.trail/0.1-draft`  
**Status:** exploratory draft; partially executable in the bounded reference profile and not stable  
**Date:** 2026-09-28

## Purpose

Trail Protocol defines a language-neutral contract for exchanging proposals and accepted facts about outcomes, commitments, work, authority, execution, evidence, and acceptance.

It is designed for:

- Trail clients and servers;
- GraycodeAI adapters such as Rho, Rover, Across, and Trace;
- self-hosted export/import;
- selective federation between independent Trail nodes;
- conformance and archival records.

It does not define model prompting, an agent runtime, Git operations, artifact storage, database layout, UI, or transport binding.

## Executable reference profile

The dependency-free Go package under `internal/domain` and `cmd/trail-sim` implements a deliberately narrow subset of this draft. It covers the flat transport-neutral command envelope, typed command dispatch, command outcomes, accepted event envelopes, optimistic versions, idempotency, explicit bilateral commitments, work traceability, capability expiry/revocation, leases and fencing, runs, artifact manifests, evidence, acceptance, and event replay. Draft 2020-12 schemas and portable fixtures live under `spec/schema/v0` and `spec/fixtures/v0`.

It does not yet implement principal or workspace resources, request/offer negotiation, signatures and canonicalization, bitemporal corrections, budgets as an authority-owned aggregate, disclosure enforcement, redaction, federation, transport bindings, or durable storage. Internal Go types are not a stable protocol API. The canonical executable journal is under `spec/fixtures/v0/valid/`.

## Normative language

`MUST`, `MUST NOT`, `SHOULD`, `SHOULD NOT`, and `MAY` are normative only when this draft is promoted to an implementation candidate. Until then they state intended constraints to validate.

## Design principles

1. Commands request change; accepted events record change.
2. Proposals, inferences, and remote assertions do not silently become authoritative state.
3. Every object has an authority and version.
4. Every mutation has an actor, policy decision, idempotency identity, and causal context.
5. Execution completion is separate from acceptance.
6. Concurrency conflicts remain explicit when meaning cannot merge.
7. Disclosure is part of the envelope and applies to derived data.
8. Serialization remains simpler than the full analytical hypergraph.

## Identifiers

Globally referenced objects use URIs. Local implementations may use UUIDv7/ULID internally but export a stable URI.

Examples:

```text
https://trail.example/workspaces/product
https://trail.example/outcomes/01K...
https://trail.example/commitments/01K...
urn:graycode:trail:event:01K...
urn:graycode:trail:run:01K...
```

An identifier is opaque. Consumers MUST NOT derive authorization, tenancy, time, or type solely from its string shape.

## Common event envelope

```json
{
  "specVersion": "graycode.trail/0.1-draft",
  "id": "urn:graycode:trail:event:01K8...",
  "type": "trail.commitment.activated",
  "subject": "https://a.example/commitments/01K8...",
  "subjectVersion": 4,
  "home": "https://a.example",
  "workspace": "https://a.example/workspaces/product",
  "actor": "https://a.example/principals/alice",
  "accountable": "https://a.example/principals/alice",
  "delegation": null,
  "recordedAt": "2026-09-28T10:31:22.418Z",
  "validTime": {
    "from": "2026-09-28T10:31:22.418Z",
    "to": null
  },
  "causedBy": "urn:graycode:trail:event:01K7...",
  "correlation": "urn:graycode:trail:negotiation:01K6...",
  "idempotencyKey": "sha256:...",
  "policy": {
    "id": "https://a.example/policies/commitment-v3",
    "version": 3,
    "decision": "urn:graycode:trail:policy-decision:01K8..."
  },
  "disclosure": {
    "classification": "shared",
    "audience": ["https://b.example"],
    "fields": ["commitment", "criteria", "deadline"]
  },
  "payload": {},
  "payloadDigest": "sha256:...",
  "signature": {
    "algorithm": "EdDSA",
    "keyId": "https://a.example/.well-known/trail/keys/2026-09#1",
    "value": "base64url..."
  }
}
```

### Envelope rules

- `id`, `type`, `subject`, `home`, `workspace`, `actor`, `recordedAt`, `idempotencyKey`, `policy`, and `payload` are required for accepted domain events.
- `accountable` identifies the principal responsible for the act; it equals `actor` for direct action and differs only with an explicit delegation chain.
- `subjectVersion` increases at the home authority for each accepted mutation to that aggregate.
- `validTime` is required when the payload asserts a fact about a time distinct from `recordedAt`.
- `causedBy` references the direct accepted/proposal event that caused the transition where known.
- `correlation` groups a negotiation, run, import, or distributed interaction.
- `delegation` references the accepted authority chain when the actor acts for another principal.
- `payloadDigest` covers canonical payload bytes. Federation messages require it.
- signatures are required for federation/export profiles and optional inside one trusted process boundary.
- disclosure metadata does not itself grant access; it constrains an already authorized exchange.

## Common command envelope

```json
{
  "specVersion": "graycode.trail/0.1-draft",
  "id": "urn:graycode:trail:command:01K8...",
  "type": "trail.commitment.accept",
  "subject": "https://a.example/commitments/01K8...",
  "expectedVersion": 3,
  "actor": "https://b.example/principals/provider",
  "accountable": "https://b.example/principals/provider",
  "delegation": null,
  "workspace": "https://a.example/workspaces/product",
  "idempotencyKey": "client-generated-stable-value",
  "submittedAt": "2026-09-28T10:31:20Z",
  "payload": {}
}
```

Commands MAY be transported through HTTP, MCP, A2A, a local queue, or federation. `accountable` is required and the authority validates any delegation connecting it to `actor`. The authority returns a structured result.

## Command result

Exactly one result class:

```text
Accepted(event, projectionHint?)
Conflict(expectedVersion, currentVersion, currentRef, resolutionOptions)
Denied(policyRef, safeReasonCode, approvalOrAppealPath?)
Invalid(invariantCode, fieldErrors)
Deferred(authority, receipt, nextCheck?)
Unavailable(retryClass, safeRetryAfter?)
```

Clients never infer acceptance from HTTP 2xx alone, a connection close, a chat response, or an agent's natural-language statement.

## Common types

### PrincipalRef

```json
{
  "id": "https://a.example/principals/agent-7",
  "kind": "agent",
  "home": "https://a.example",
  "operator": "https://a.example/principals/alice"
}
```

Kinds: `person`, `agent`, `service`, `connector`, `organization`, `group`, `device`, `workload`, `node`.

### SubjectRef

```json
{
  "id": "https://a.example/outcomes/01K...",
  "type": "outcome",
  "version": 7,
  "home": "https://a.example"
}
```

### TemporalRef

```json
{
  "validFrom": "2026-09-28T00:00:00Z",
  "validTo": null,
  "recordedAt": "2026-09-28T10:31:22Z"
}
```

### ProvenanceRef

```json
{
  "source": "https://trace.example/repos/x/pulls/42",
  "sourceType": "trace.pull-request",
  "sourceVersion": "sha256:...",
  "activity": "urn:graycode:trail:run:01K...",
  "attributedTo": "https://a.example/principals/agent-7",
  "epistemicStatus": "observed"
}
```

Epistemic status values: `observed`, `stated`, `inferred`, `simulated`, `approved`, `disputed`, `superseded`, `unknown`.

### Disclosure

```json
{
  "classification": "private|workspace|restricted|shared|public",
  "audience": ["principal-or-node-uri"],
  "purpose": ["coordination", "verification"],
  "fields": ["json-pointer-or-profile-name"],
  "expiresAt": null
}
```

## Domain resources

### Outcome

```json
{
  "id": "https://a.example/outcomes/01K...",
  "type": "outcome",
  "version": 3,
  "desiredCondition": "Users can recover an interrupted agent run safely.",
  "steward": "https://a.example/principals/alice",
  "criteria": [
    {
      "id": "criterion:resume-without-duplicate-effect",
      "statement": "Recovery produces no duplicate protected side effect.",
      "evidencePolicy": "https://a.example/evidence-policies/fault-injection-v1"
    }
  ],
  "horizon": {"due": null, "timezone": "Asia/Kolkata"},
  "status": "active"
}
```

Outcome status: `draft`, `active`, `at-risk`, `submitted`, `accepted`, `rejected`, `cancelled`, `superseded`.

### SituationClaim

```json
{
  "id": "https://a.example/claims/01K...",
  "subject": "https://a.example/systems/worker-1",
  "predicate": "trail:hasStatus",
  "value": "unreachable",
  "epistemicStatus": "observed",
  "confidence": 1.0,
  "provenance": {},
  "validTime": {"from": "...", "to": null}
}
```

Disputes and corrections reference and supersede claims; they do not mutate the original statement.

### Request and Offer

Requests and offers share:

```json
{
  "id": "...",
  "kind": "request|offer",
  "from": "principal-uri",
  "to": ["principal-or-audience-uri"],
  "consequent": {},
  "antecedent": {},
  "criteria": ["criterion-ref"],
  "expiresAt": "...",
  "status": "open"
}
```

Status: `draft`, `open`, `countered`, `accepted`, `rejected`, `withdrawn`, `expired`, `superseded`.

### Commitment

```json
{
  "id": "https://a.example/commitments/01K...",
  "type": "commitment",
  "version": 4,
  "debtor": "https://b.example/principals/team-b",
  "creditor": "https://a.example/principals/team-a",
  "antecedent": {
    "type": "all",
    "terms": [{"type": "accepted-event", "ref": "..."}]
  },
  "consequent": {
    "type": "outcome-contribution",
    "outcome": "https://a.example/outcomes/01K...",
    "criteria": ["criterion:api-compatible"]
  },
  "deadline": "2026-10-30T12:00:00Z",
  "acceptancePolicy": "https://a.example/policies/acceptance-v2",
  "delegation": {"allowed": true, "originalDebtorRemainsLiable": true},
  "status": "active"
}
```

### Relation

Relation objects represent binary and multi-party relations:

```json
{
  "id": "https://a.example/relations/01K...",
  "type": "trail:enables",
  "participants": [
    {"role": "source", "ref": "work:a"},
    {"role": "source", "ref": "work:b"},
    {"role": "target", "ref": "work:c"}
  ],
  "expression": {"operator": "all"},
  "rationale": "Both schemas are required for integration.",
  "confidence": 1.0,
  "epistemicStatus": "approved"
}
```

Initial relation types:

- `contributesTo`;
- `decomposes`;
- `enables`;
- `blocks`;
- `requiresEvidence`;
- `conflictsWith`;
- `usesResource`;
- `supersedes`;
- `disputes`;
- `derivedFrom`;
- `actsOnBehalfOf`.

Extensions use namespaced URIs and declare semantics.

### CapabilityGrant

```json
{
  "id": "https://a.example/grants/01K...",
  "issuer": "https://a.example/principals/alice",
  "subject": "https://a.example/principals/agent-7",
  "actions": ["trail.run.checkpoint", "trace.pull-request.create"],
  "resources": ["https://trace.example/repos/x"],
  "purpose": ["commitment:01K..."],
  "constraints": {
    "maxCost": {"amount": "20.00", "currency": "USD"},
    "network": ["api.github.com"],
    "writePaths": ["src/**", "tests/**"]
  },
  "delegationDepth": 0,
  "validFrom": "...",
  "expiresAt": "...",
  "status": "active"
}
```

### ClaimLease

```json
{
  "id": "https://a.example/leases/01K...",
  "work": "https://a.example/work/01K...",
  "holder": "https://a.example/principals/agent-7",
  "issuedBy": "https://a.example",
  "fencingToken": 42,
  "issuedAt": "...",
  "expiresAt": "...",
  "status": "active"
}
```

### Run

```json
{
  "id": "urn:graycode:trail:run:01K...",
  "work": {"id": "...", "version": 8},
  "plan": {"id": "...", "version": 2, "digest": "sha256:..."},
  "executor": "https://a.example/principals/agent-7",
  "operator": "https://a.example/principals/alice",
  "workload": "spiffe://a.example/trail/runner/01K...",
  "lease": {"id": "...", "fencingToken": 42},
  "grants": ["grant-ref"],
  "budget": {"reservation": "..."},
  "evidencePolicy": "...",
  "status": "active"
}
```

### Artifact and Evidence

```json
{
  "artifact": {
    "id": "urn:cid:bafy...",
    "digest": "sha256:...",
    "mediaType": "application/vnd.git.patch",
    "size": 48220,
    "locations": ["https://trace.example/..."],
    "producedBy": "urn:graycode:trail:run:01K...",
    "classification": "restricted"
  },
  "evidence": {
    "id": "https://a.example/evidence/01K...",
    "subject": {"id": "...", "version": 8},
    "criterion": "criterion:api-compatible",
    "method": "trace.ci/check-suite",
    "result": "pass",
    "evaluator": "https://trace.example/services/ci",
    "artifacts": ["urn:cid:bafy..."],
    "limitations": ["does not test third-party rate limits"],
    "validUntil": "2026-10-05T00:00:00Z"
  }
}
```

### Review and Acceptance

Acceptance is criterion-scoped:

```json
{
  "id": "https://a.example/acceptances/01K...",
  "subject": {"id": "https://a.example/work/01K...", "version": 8},
  "criterionResults": [
    {
      "criterion": "criterion:api-compatible",
      "result": "accepted",
      "evidence": ["https://a.example/evidence/01K..."]
    }
  ],
  "policy": {"id": "...", "version": 2},
  "acceptedBy": "https://a.example/principals/reviewer",
  "acceptedAt": "...",
  "scope": "criterion"
}
```

## Command families

Initial command names:

### Situation/outcome

- `trail.situation.claim.propose`
- `trail.situation.claim.dispute`
- `trail.situation.claim.supersede`
- `trail.outcome.create`
- `trail.outcome.amend`
- `trail.outcome.submit`
- `trail.outcome.cancel`

### Negotiation/commitment

- `trail.request.create|withdraw|reject`
- `trail.offer.create|withdraw|reject`
- `trail.negotiation.counter`
- `trail.commitment.propose`
- `trail.commitment.accept|amend|delegate|release|cancel`
- `trail.commitment.mark-impossible|report-violation`

### Work/execution

- `trail.work.create|amend|cancel`
- `trail.relation.propose|accept|dispute|remove`
- `trail.claim.acquire|renew|release|force-reassign`
- `trail.run.request|checkpoint|pause|resume|cancel|report-result`
- `trail.budget.reserve|record-spend|release`

### Evidence/acceptance

- `trail.artifact.register`
- `trail.evidence.submit|dispute|supersede`
- `trail.review.request|submit`
- `trail.acceptance.accept|reject|request-change|dispute`

### Governance

- `trail.grant.issue|attenuate|revoke`
- `trail.policy.publish|activate|retire`
- `trail.disclosure.grant|revoke`
- `trail.identity.bind|revoke`

The final schema may use a smaller command set with typed payload actions. The names above make missing semantics visible during research.

## Event families

Accepted events use past tense, for example:

- `trail.outcome.created`, `trail.outcome.amended`;
- `trail.request.opened`, `trail.offer.countered`;
- `trail.commitment.proposed`, `trail.commitment.party-accepted`, `trail.commitment.activated`, `trail.commitment.satisfied`, `trail.commitment.violated`;
- `trail.work.ready`, `trail.claim.issued`, `trail.claim.released`, `trail.claim.expired`;
- `trail.run.started`, `trail.run.checkpointed`, `trail.run.completed`, `trail.run.failed`;
- `trail.artifact.registered`, `trail.evidence.submitted`;
- `trail.review.completed`, `trail.acceptance.recorded`, `trail.acceptance.disputed`;
- `trail.grant.issued`, `trail.grant.revoked`;
- `trail.policy.activated`;
- `trail.federation.receipt.recorded`.

Derived states such as readiness or violation-by-time may be emitted by an authorized deterministic service with input versions and policy, or computed as projections. The specification must choose one source of authority per derived fact.

## Federation interaction profile

Remote nodes exchange messages containing one of:

- a proposal command addressed to the subject home;
- an accepted event for an object the sender owns;
- a receipt for a prior message;
- a request for missing causal data;
- a revocation/tombstone/dispute within sender authority;
- a capability/version discovery document.

A valid signature proves control of the signing key, not authority over the subject. The receiver checks:

1. message size and basic parse limits;
2. spec/extension compatibility;
3. recipient and replay window;
4. key validity and signature;
5. sender/node/principal binding;
6. subject home and sender authority;
7. causal dependencies and version;
8. disclosure and local policy;
9. payload/domain invariants;
10. idempotency/inbox result.

## Versioning

- protocol versions use `major.minor` while draft;
- incompatible semantic changes increment major;
- additive optional fields/extensions increment minor;
- unknown optional fields are preserved where round-trip is required and ignored semantically;
- unknown required extensions cause `Invalid` or federation quarantine;
- resources include their own schema version when evolution differs from envelope version;
- event meaning is immutable after publication; corrections create a new version/event;
- protocol release and Trail product release are independent.

## Canonicalization and signatures

The executable Go profile currently normalizes parsed JSON objects before computing `payloadDigest`, making the digest insensitive to insignificant whitespace and object-key order. This normalization is only an interim fixture rule and is not yet a cross-language signing profile.

The final implementation candidate must choose one canonical representation and test it across languages. Candidate approaches:

- JSON Canonicalization Scheme plus detached signature over envelope/payload digest;
- signed CBOR/COSE profile;
- HTTP message signatures plus signed domain digest.

Do not sign parser-dependent raw JSON text without canonicalization. Do not use JSON-LD canonicalization unless interoperability benefit justifies its complexity and denial-of-service controls.

## Privacy and redaction

- secrets are forbidden in protocol payloads;
- artifact bytes are referenced, not embedded by default;
- sensitive fields may live in encrypted payload objects whose key can be destroyed;
- a redaction event identifies removed fields and legal/policy authority without repeating the content;
- public error responses do not reveal hidden objects, membership, or topology;
- export filters must not break referential integrity silently: missing objects become authorized redacted placeholders.

## Conformance fixtures

The protocol package must include:

- minimal valid command/event for every family;
- unknown additive field;
- unknown required extension;
- duplicate command/message;
- changed payload under reused idempotency key;
- stale subject version;
- invalid delegation chain;
- revoked/expired grant;
- expired lease and stale fencing token;
- acceptance against stale subject/evidence;
- backdated valid-time claim;
- dispute and supersession chain;
- federation replay, missing predecessor, wrong audience, bad signature, rotated key;
- disclosure-redacted export;
- old reader/new writer compatibility cases.

## Open issues before implementation candidate

1. Final vocabulary for request/offer/commitment and joint commitments.
2. URI versus compact ID rules for local/offline creation.
3. JSON Schema versus another schema system and code-generation policy.
4. Canonicalization/signature profile.
5. Exact bitemporal fields per resource.
6. Policy-decision representation and causal consistency token.
7. ActivityPub/ActivityStreams reuse in federation.
8. CRDT document reference/version representation.
9. Artifact digest/CID profile and encryption metadata.
10. Redaction semantics under audit, legal hold, and federation.
11. Stable mapping to W3C PROV, in-toto/SLSA, MCP, A2A, and ForgeFed.
12. Whether readiness transitions are events or purely derived projections.
