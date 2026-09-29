# Trail concierge prototype specification

**Type:** disposable research prototype  
**Purpose:** validate language, decisions, recovery, and value before application code

## Constraints

- No production backend or autonomous side effects.
- One workspace, one outcome, and one pilot workflow at a time.
- Existing repositories, agents, CI, and trackers remain authoritative.
- Researchers may update the prototype manually behind the scenes.
- Every screen needs a keyboard-accessible list/table representation.
- The prototype must display uncertainty and missing information rather than inventing it.

## Core scenario

A human requests a change spanning two repositories. One or more agents propose and execute bounded work through existing GraycodeAI tools. Trail shows commitments, readiness, authority, runs, evidence, review, and acceptance. Midway through the scenario, one actor is interrupted and another must resume.

## Required surfaces

### 1. Outcome brief

Show:

- desired outcome in the requester's words;
- requester and acceptance owner;
- criteria with required evidence;
- scope, exclusions, deadline, risk, and budget;
- current lifecycle and version;
- unresolved questions.

Research question: can participants state what success means and notice when criteria change?

### 2. Commitment panel

Show request, offers, accepted commitment, beneficiary, responsible party, conditions, due point, and state. Allow propose, accept, decline, renegotiate, release, violate, supersede, and dispute in the scenario.

Research question: does this express accountability better than assignment alone?

### 3. Ready/blocked view

Show each work item in a list with:

- ready, blocked, uncertain, claimed, running, review, or terminal state;
- typed blocking condition;
- explanation and source freshness;
- available capable actors;
- authority and budget summary.

Research question: can users predict what may happen next and why?

### 4. Authority inspector

For a selected action, show:

- accountable principal;
- acting human/agent/service and operator;
- delegation chain;
- allowed action/resource/scope;
- lease/fencing token status;
- budget and expiry;
- policy decision and denial reason;
- revocation or exception path.

Research question: can a user determine whether the actor is currently allowed to act?

### 5. Run and evidence timeline

Separate:

- proposed action;
- authorized run;
- external side effect;
- executor completion claim;
- artifact and digest;
- deterministic verification;
- reviewer observation;
- acceptance decision.

Research question: do users stop treating run completion as accepted completion?

### 6. Review desk

Show the exact subject/version, applicable criteria, required and supplied evidence, limitations, conflicts, stale evidence, independent verifier, and actions: accept, reject, request change, dispute, or defer.

Research question: does structured evidence reduce review or increase confidence enough to repay its cost?

### 7. Resumption capsule

Generate a bounded handoff containing:

- current situation and uncertainty;
- desired outcome and active version;
- accepted commitments;
- completed, active, and ready work;
- decisions and rationale;
- current authority, budget, and lease state;
- artifacts/evidence with provenance;
- blockers, risks, and next decision;
- information intentionally omitted.

Research question: can a new participant resume safely without reading the full history?

### 8. Audit and export

Provide chronological and subject-based views of proposals, accepted mutations, policies, runs, evidence, reviews, disputes, and acceptance. Demonstrate an export that remains understandable without the prototype.

Research question: can participants reconstruct who knew, decided, authorized, performed, verified, and accepted what?

## Scenario variants

1. Criteria change after an agent has started.
2. Two agents claim the same exclusive work.
3. A lease expires while the old worker is offline.
4. CI passes but the acceptance criterion is not satisfied.
5. Evidence refers to an earlier commit.
6. The executor and verifier share the same faulty assumption.
7. A reviewer rejects only one criterion.
8. An acceptance decision is disputed later.
9. The budget is exhausted with partial useful work.
10. A private dependency blocks work but may not be disclosed.
11. A participant loses access during a run.
12. An external side effect succeeds before the run record is written.

## Prototype data model

Use static fixtures or a researcher-controlled table with these minimum objects:

```text
principal
outcome + version
criterion + evidence policy
request / offer / commitment
work item + typed relation
grant / claim / lease
run + side-effect status
artifact + digest
evidence + verifier result
review / acceptance / dispute
decision
event
```

Do not build reusable persistence, authentication, synchronization, federation, or plugin infrastructure for this prototype.

## Session script

1. Give the participant only the outcome brief.
2. Ask what can happen next and who may do it.
3. Ask them to negotiate/inspect a commitment.
4. Introduce a ready and a blocked work item.
5. Start a simulated agent run.
6. Inject one scenario variant without advance warning.
7. Ask another participant to resume using the capsule.
8. Present conflicting or incomplete evidence.
9. Ask the acceptance owner for a scoped decision.
10. Ask both participants to reconstruct the history from export.

Record errors, hesitations, facilitator help, terminology substitutions, time, and confidence. Do not coach during the measured portion.

## Prototype exit decision

Advance to a concierge pilot only if participants can understand the lifecycle, authority, evidence, and recovery model without repeated explanation. Revise vocabulary and interaction before technical architecture when they cannot.

