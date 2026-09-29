# Industry standards and readiness map

**Reviewed:** 2026-09-28  
**Purpose:** distinguish implementation standards, design references, buyer expectations, and later certifications

Standards are inputs to product design. Listing one does not imply certification or legal compliance. Exact legal duties depend on Trail's users, deployments, data, jurisdictions, and role in each AI system.

## Readiness tiers

### Tier 1: use in the first implementation

| Area | Standard or specification | Required Trail posture |
|---|---|---|
| HTTP API | [OpenAPI](https://spec.openapis.org/oas/) | pin a supported release after generator/tooling tests; publish command/query APIs and errors |
| schemas | [JSON Schema 2020-12](https://json-schema.org/draft/2020-12) | canonical event and object validation, fixtures, and forward-compatibility rules |
| event envelope | [CloudEvents](https://cloudevents.io/) | reuse envelope conventions where compatible; keep Trail authority and temporal fields |
| agent tools | [MCP 2026-07-28](https://modelcontextprotocol.io/specification/2026-07-28) | narrow query/proposal/action surface; negotiate extensions; never expose raw storage |
| agent delegation | [A2A 1.0.0](https://a2a-protocol.org/v1.0.0/specification/) | optional adapter for discovery and task/artifact exchange; Trail remains the durable business contract |
| provenance | [W3C PROV-O](https://www.w3.org/TR/prov-o/) | map principals, activities, artifacts, derivation, and delegation for interchange |
| authentication | OpenID Connect, WebAuthn, OAuth 2.0 with [RFC 9700](https://www.rfc-editor.org/info/rfc9700/) | standards-based login and delegated API access; PKCE, exact redirects, short-lived/sender-constrained credentials where practical |
| authorization | [AuthZEN Authorization API 1.0](https://openid.net/wg/authzen/specifications/) concepts | keep policy decision and policy enforcement separate; evaluate direct compatibility after a spike |
| observability | [OpenTelemetry](https://opentelemetry.io/docs/specs/status/) | stable traces, metrics, and logs with content/secret redaction; domain audit remains separate |
| accessibility | [WCAG 2.2](https://www.w3.org/TR/WCAG22/) AA | keyboard and non-drag alternatives, focus, target size, accessible authentication, understandable errors |
| secure development | [NIST SSDF 1.1](https://csrc.nist.gov/pubs/sp/800/218/final) | security requirements, protected build environments, provenance, vulnerability response, and tracked design decisions |
| agent security | [OWASP Agentic Security Initiative](https://genai.owasp.org/initiatives/agentic-security-initiative/) | threat-model memory, tools, identity, human oversight, multi-agent behavior, and MCP boundaries |
| evidence integrity | [in-toto](https://in-toto.io/docs/specs/) and SLSA concepts | signed attestations for software evidence; do not confuse attestation with outcome acceptance |

AG-UI is promising for streaming state and human steering, but it should remain an adapter experiment until its stability and Trail's UI needs are demonstrated.

### Tier 2: design for from the beginning

| Area | Reference | Design consequence now |
|---|---|---|
| information security | [ISO/IEC 27001:2022](https://www.iso.org/standard/27001) | asset inventory, risk ownership, access reviews, incident response, supplier controls, backup and continuity evidence |
| AI management | [ISO/IEC 42001:2023](https://www.iso.org/standard/81230.html) | AI inventory, accountable owners, impact/risk assessment, monitoring, change control, and improvement loop |
| AI risk | [NIST AI RMF](https://www.nist.gov/itl/ai-risk-management-framework) and GenAI Profile | govern, map, measure, and manage risks across the lifecycle; document limitations and evaluation evidence |
| service controls | [AICPA SOC suite](https://www.aicpa-cima.com/resources/landing/system-and-organization-controls-soc-suite-of-services) | make security, availability, processing integrity, confidentiality, and privacy controls auditable |
| identity lifecycle | [SCIM 2.0 RFC 7644](https://www.rfc-editor.org/info/rfc7644/) | keep tenant/user/group lifecycle compatible with enterprise provisioning and rapid deprovisioning |
| workload identity | [SPIFFE](https://spiffe.io/docs/latest/spiffe-specs/) | distinguish human, agent, service, node, device, and concrete worker identities |
| software supply chain | SLSA, SBOMs using SPDX/CycloneDX, signed releases | retain build provenance and component inventory from the first release pipeline |
| privacy | GDPR-style purpose limitation, minimization, retention, access, export, correction, and erasure | separate mutable/sensitive payloads from durable audit references and document every telemetry field |
| reliability | SLOs, error budgets, recovery objectives, restore drills | define correctness and recovery targets before claiming production readiness |

Formal ISO certification and a SOC 2 examination are later business decisions. Their control evidence is expensive to reconstruct, so foundational records should exist from the first production pilot.

### Tier 3: interoperability and enterprise expansion

| Area | Reference | Trail use |
|---|---|---|
| project management | [ISO 21502:2020](https://www.iso.org/standard/74947.html), currently marked for revision | vocabulary and import/export mapping for governance, benefits, outcomes, risk, issues, change, and lifecycle; do not reproduce the standard internally |
| process and decisions | BPMN 2.0.2 and DMN | import/export defined business processes and decisions; do not make BPMN the adaptive agent runtime |
| lifecycle integration | OSLC Core and Change Management | integrate with enterprise requirements/change systems using linked resources |
| federation | ActivityPub and ForgeFed patterns | addressed delivery and forge interop experiments after a useful single-node product |
| user-controlled data | Solid and local-first protocols | reference for portable/user-owned data; no premature conformance claim |
| collaborative documents | Automerge or Yjs | merge descriptions, comments, and selected plans only after field-level conflict analysis |

## Regulatory watch

The EU AI Act is already relevant to product design. The European Commission states that enforcement powers and certain transparency duties began applying on 2026-08-02. Interactive AI systems may need to disclose that a person is interacting with AI. High-risk duties depend on the use case and have later dates under the 2026 changes.

Trail should therefore have, before an external pilot:

- visible human-versus-agent identity and action attribution;
- an AI-system inventory and named operator/provider roles;
- model/provider/version and policy provenance for consequential runs;
- documented intended use, prohibited use, limitations, and risk class;
- human oversight and effective stop/takeover controls;
- evaluation records for accuracy, robustness, security, and harmful bias where relevant;
- retention and disclosure rules for prompts, outputs, traces, evidence, and derived metrics;
- an exportable audit trail and a process for contesting automated decisions.

Employment, education, essential services, biometrics, law enforcement, migration, credit, health, and safety-related workflows can create much higher regulatory exposure. They should be excluded from the initial pilot unless specialist legal and domain review is part of the pilot.

## Enterprise buyer baseline

Before asking an organization to install a pilot, Trail should answer:

1. Where is each class of data stored, replicated, logged, backed up, and deleted?
2. Which subprocessors and model providers receive content?
3. Can models train on customer data, and how is that technically controlled?
4. How are SSO, MFA, service accounts, SCIM, role changes, and termination handled?
5. Can an administrator read content, and is break-glass access audited?
6. How are secrets brokered, scoped, rotated, and excluded from prompts and traces?
7. What is the incident, vulnerability-disclosure, and patch policy?
8. Can the customer export all data and verify a restore?
9. Which events prove authorization, execution, evidence, review, and acceptance?
10. What happens when a model, provider, connector, node, or policy service fails?
11. Which accessibility standard is tested, and by whom?
12. What is supported, experimental, deprecated, and compatible across releases?

## Protocol corrections from this review

- A2A's released specification is now **1.0.0**; older v0.3 references are stale.
- MCP's 2026-07-28 core and extensions need separate version negotiation; Tasks moved into an official extension proposal rather than remaining an assumed core primitive.
- IBM's Agent Communication Protocol has merged into A2A. Trail should not implement ACP as a parallel generic agent protocol.
- AuthZEN Authorization API 1.0 became a final OpenID specification in January 2026 and is now a serious policy interoperability candidate.
- OpenAPI 3.2 exists, but Trail should choose its pinned version from ecosystem compatibility tests rather than adopting the newest release automatically.
- ISO 21502:2020 entered revision status in September 2026; use it as guidance and monitor its replacement.

## Readiness evidence, not badges

For each applicable standard or framework, maintain a small control record:

```text
control or requirement
applicability and rationale
owner
implementation reference
verification evidence
known gap
residual risk
review date
```

This produces useful engineering evidence before the company decides whether formal certification is commercially justified.

