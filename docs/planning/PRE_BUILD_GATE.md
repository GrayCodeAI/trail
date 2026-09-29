# Trail pre-build gate

**Decision:** do not start the production application yet  
**Next objective:** prove one narrow job, one user group, and the minimum safe technical path

## Why another gate is necessary

Current research supports the architecture, but it does not yet prove demand. The market already contains mature work managers, new agent-native trackers, control planes, shared-memory products, and open protocols. A broad build would compete on many commodity features before proving Trail's distinct value.

The next work should eliminate the largest product, market, and correctness uncertainties with interviews, concierge workflows, prototypes, and technical spikes.

## Recommended first wedge

### User

A software team or OSS maintainer coordinating several coding agents across two or more repositories.

### Pain

The team repeatedly loses outcome intent, duplicates work, cannot see current authority, struggles to resume interrupted runs, or discovers that “done” output lacks the evidence needed for acceptance.

### Job

Coordinate one cross-repository outcome from intent through bounded execution, evidence, review, and explicit acceptance.

### Existing GraycodeAI roles

- Rho acts as the coding-agent interface.
- Rover performs isolated execution and verification.
- Across supplies bounded continuity/provenance references.
- Trace supplies repository, change, review, and CI references.
- Trail owns the outcome, commitment, authority, readiness, evidence requirements, attention queue, and acceptance record.

## Gate A: problem evidence

Use the [problem interview guide](./INTERVIEW_GUIDE.md) and store redacted observations in the [field evidence ledger](../../research/evidence/README.md).

### Work

1. Interview 15 to 20 people across at least five teams that already use coding agents for real work.
2. Ask for the last concrete failure, artifact, timeline, and workaround rather than opinions about a proposed product.
3. Observe at least five real workflows from request through merge/deployment/acceptance.
4. Reconstruct the cost of coordination: interruptions, duplicate runs, rework, review time, waiting, escaped errors, and spend.
5. Identify the buyer, daily user, administrator, security approver, and person harmed by a mistake.
6. Record counterexamples where the existing tracker and agent tooling work well.

### Pass evidence

- at least five teams report the same high-cost failure pattern;
- the pattern occurs repeatedly, not as a one-off incident;
- at least three teams commit time and real workflows to a design-partner pilot;
- at least two provide redacted artifacts/events sufficient to reconstruct a workflow;
- at least one shows willingness to pay, sign a pilot agreement, or commit meaningful engineering time if the measured result improves;
- a named owner has budget or operational responsibility for the problem.

### Fail or revise

- interest is mostly about “AI project management” as a concept;
- teams ask for generic boards, chat, documents, or dashboards first;
- the pain disappears after configuring Jira/Linear plus an existing agent runner;
- no participant will expose a real workflow even under redaction;
- the affected user cannot influence adoption or identify a buyer.

## Gate B: vocabulary and interaction

Use the [concierge prototype specification](./CONCIERGE_PROTOTYPE.md).

### Prototype

Create a clickable, disposable prototype with only:

- outcome and acceptance criteria;
- request/offer/commitment;
- participant and authority explanation;
- ready/blocked/uncertain view;
- claim/lease and budget summary;
- run/evidence timeline;
- review, accept, reject, dispute, and supersede;
- resumption capsule and attention inbox.

Provide list/table forms for every relationship. The graph is an optional explanation view.

### Tests

- scenario comprehension without teaching the vocabulary first;
- keyboard and screen-reader walkthrough;
- time to capture an outcome;
- ability to predict who may act next and why;
- ability to tell proposed, executed, verified, and accepted states apart;
- takeover after an interrupted or failed agent;
- response to stale, conflicting, missing, or malicious evidence.

### Pass evidence

- at least 80% of participants correctly predict the core lifecycle;
- every participant can identify current authority and the acceptance owner;
- core capture is competitive with the baseline tracker for the pilot workflow;
- users can resume a simulated interrupted outcome without facilitator explanation;
- evidence adds less attention cost than the rework/failure it prevents in the chosen task class.

## Gate C: concierge pilot

Recruit with the [design-partner brief](./DESIGN_PARTNER_BRIEF.md) and register results in the [pilot scorecard](./PILOT_SCORECARD.md).

Run Trail manually before automating it. A founder/researcher can maintain the event/evidence ledger behind the prototype while participants use the workflow.

### Pilot sequence

1. Import one real outcome and its criteria.
2. Record participants, policies, and bounded agent authority.
3. Negotiate a commitment and create the work decomposition.
4. Let existing tools execute; do not replace their runtime.
5. Collect repository, test, review, and run evidence manually or read-only.
6. Produce a resumption capsule after a deliberate interruption.
7. Ask the authorized reviewer to accept, reject, or request change against the exact version.
8. Compare against a similar baseline workflow without Trail.

### Metrics

Choose thresholds before the pilot for:

- time from request to unambiguous commitment;
- time to determine what can act next;
- time to resume after interruption;
- duplicated or conflicting work;
- review time and review-queue age;
- false-completion and reopened-work rate;
- escaped acceptance errors;
- human attention minutes per accepted outcome;
- compute/model cost per accepted outcome;
- completeness of authority and evidence reconstruction;
- participant return rate for a second workflow.

### Pass evidence

- at least one preregistered failure metric improves materially over the baseline;
- no critical safety or authority regression occurs;
- participants voluntarily use Trail for a second real outcome;
- the acceptance owner says the evidence changed or accelerated a decision;
- the team can identify which product behavior created the benefit.

Positive feedback without repeat use does not pass this gate.

## Gate D: technical uncertainty spikes

Each spike should be disposable, time-bounded, and produce a decision record rather than application scaffolding.

| Spike | Question | Required result |
|---|---|---|
| lifecycle model | can commitment, claim, evidence, acceptance, dispute, and revocation remain correct under concurrency? | model/property tests for all P0 invariants |
| journal/projection | can a small event journal rebuild deterministically and migrate safely? | representative benchmark, corruption and rebuild report |
| authorization | can policy decisions express human, agent, service, delegation, lease, and evidence constraints? | OpenFGA/SpiceDB/AuthZEN/custom comparison with adversarial cases |
| durable execution | should Trail call Rover plus Temporal/Hatchet/Restate or implement a narrow worker? | crash/retry/cancel/uncertain-side-effect comparison and cost |
| evidence | can CI, review, run, and artifact attestations bind to an exact subject/version? | cross-language canonicalization and verifier fixtures |
| adapter boundary | can Rho, Rover, Across, and Trace integrate only through public contracts? | read-only end-to-end scenario with no internal imports |
| privacy | can a shared boundary report readiness without leaking private topology? | inference test corpus and explicit leakage results |
| export/recovery | can a user leave and reconstruct a useful record? | documented full export and restore demonstration |
| protocol stack | which versions of OpenAPI, MCP, A2A, and optional AG-UI interoperate with target clients? | compatibility matrix and pinned versions |

## Gate E: business and OSS strategy

Decide before accepting outside contributions or production data:

### Product model

- hosted service, self-hosted product, or both;
- which component is the paid value;
- expected buyer and purchasing path;
- whether pricing follows active humans, active agents, accepted outcomes, compute, or deployment;
- how a user gets value before a network of federated nodes exists.

### License matrix

Decide separately for:

- Trail protocol and schemas;
- reference SDKs and conformance fixtures;
- server/kernel;
- web client;
- official adapters;
- enterprise/hosted components.

Evaluate Apache-2.0/MIT, MPL/EPL, and AGPL tradeoffs. Do not copy code until dependency licenses, notices, trademarks, patents, and network-copyleft implications are recorded.

### Governance

- maintainer and decision rights;
- RFC/ADR and compatibility process;
- contribution agreement or DCO choice;
- security disclosure and supported-version policy;
- telemetry and privacy policy;
- release signing, SBOM, provenance, and dependency policy;
- community moderation and code of conduct;
- trademark use for “Trail by GraycodeAI.”

### Distribution experiment

Before building a suite, test whether the target users respond to:

- a concise problem page and workflow demo;
- direct outreach to agent-heavy OSS maintainers and small engineering teams;
- a manual “coordination audit” of one failed or delayed agent workflow;
- an import-only report showing missing authority, stale work, evidence gaps, and resumption cost;
- a public protocol/RFC discussion focused on evidence-backed acceptance.

Measure qualified conversations, shared real workflows, pilot commitments, repeat use, and referrals. Page views and stars are secondary.

## Gate F: security, privacy, and industry readiness

Before an external pilot:

- classify the pilot's AI Act role/use case and exclude high-risk domains by default;
- document AI disclosure, human oversight, stop/takeover, and contest paths;
- complete a data-flow and privacy-impact review;
- define tenant isolation, administrator access, retention, deletion, export, and backup behavior;
- define model-provider and subprocessor data handling;
- define secret brokering and prompt/tool/evidence redaction;
- map initial controls to NIST SSDF, NIST AI RMF, OWASP agentic guidance, ISO/IEC 27001, and ISO/IEC 42001;
- plan WCAG 2.2 AA verification for every core flow;
- create incident, vulnerability disclosure, and compromised-agent/node runbooks;
- choose supported deployment environments and recovery objectives.

Certification is not required to validate the product. Engineering evidence that later supports certification is required.

## Build/no-build decision record

Complete this table with links to evidence:

| Decision area | Required answer | Status |
|---|---|---|
| target user | one narrow role/team profile with observed workflow | open |
| urgent problem | repeated costly failure with a named owner | open |
| alternative | why configuration of an existing tracker/control plane is insufficient | open |
| wedge | one end-to-end job Trail performs | proposed |
| repeat demand | second voluntary workflow from pilot users | open |
| measurable value | preregistered metric improved against baseline | open |
| comprehension | lifecycle and authority threshold passed | open |
| technical safety | P0 invariants and recovery spikes passed | in progress: ADR 0006 reference kernel covers the first lifecycle, fencing, revocation, evidence, acceptance, and replay invariants |
| ecosystem boundary | public contracts agreed with sibling products | open |
| business | buyer, deployment, pricing hypothesis, and pilot terms | open |
| OSS | license, governance, contribution, security, and release model | open |
| compliance | applicable roles/data/jurisdictions and initial controls | open |
| stop condition | evidence that would cancel or narrow the build | open |

Production implementation begins only when every row is either passed with evidence or deliberately narrowed in an ADR.

## If the gate passes: first build

The first product slice should be:

- one node and one organization;
- one cross-repository software outcome at a time;
- human-created outcome and explicit acceptance owner;
- read-only Rho/Rover/Across/Trace adapters;
- manual commitments and policy decisions;
- event journal and rebuildable list/audit projections;
- evidence manifest bound to exact versions;
- manual review, accept, reject, dispute, and export;
- no autonomous dispatch until this is useful repeatedly.

Defer advanced scheduling, rich graph visualization, local-first editing, federation, marketplaces, generic plugins, mobile apps, chat, documents, calendar, billing, and a broad PM hierarchy.

## Immediate sequence

1. Review and select the first wave from the [public candidate shortlist](./PUBLIC_CANDIDATE_SHORTLIST.md).
2. Write the interview guide and evidence-consent/redaction process.
3. Select five past agent-work failures to reconstruct.
4. Build the disposable interaction prototype.
5. Register baseline metrics and pilot thresholds.
6. Secure three design-partner commitments.
7. Run the concierge pilot.
8. Execute only the technical spikes needed to answer pilot-blocking questions.
9. Review the build/no-build table and record the decision.
