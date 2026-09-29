# Decentralized work-platform research

> **Archive note (2026-09-28):** These were the initial research documents. They were moved intact from the workspace-level `research/` directory into `trail/research/initial/`. The refreshed research index, requirements, architecture, and protocol draft begin at [Trail's main README](../../README.md).

Read these documents in this order:

1. [GraycodeAI workspace map](./graycodeai-workspace-map.md) — current repository ownership, dependency boundaries, overlap analysis, and the recommended place for Trail.
2. [Product name shortlist](./name-shortlist.md) — the Trail naming decision, namespace strategy, alternatives, taglines, and visual direction.
3. [Dynamic human-agent architecture](./dynamic-human-agent-architecture.md) — adaptive operating model supporting 90/10, 80/20, 50/50, and human-led work, with an architecture diagram, autonomy control, attention capacity, and version-one boundary.
4. [Beyond GraphDone: coordination fabric](./beyond-graphdone-coordination-fabric.md) — current product thesis, cross-disciplinary evidence, domain model, federation model, agent protocol, build sequence, and falsification plan.
5. [First-principles architecture](./first-principles-work-coordination-architecture.md) — deeper mathematics, scheduling, distributed-systems reasoning, threat model, and engineering requirements.
6. [OSS and research landscape](./decentralized-agentic-work-management-landscape-2026.md) — open-source comparisons, standards, agent platforms, social signals, and the initial GraphDone gap analysis.

The first two documents supersede the earlier graph-centered recommendation. The older material remains useful as evidence and for tracing how the design changed.

## Open the Markdown

- In VS Code, open a file and use **Markdown: Open Preview** (`Cmd+Shift+V` on macOS).
- On GitHub, open the file normally; Mermaid diagrams render in the page.
- Any Markdown viewer that supports Mermaid can render the architecture and lifecycle diagrams.

## Current architectural thesis

The platform uses a **fully dynamic human-agent ratio**. It can run at 90/10, 80/20, 50/50, or human-led operation depending on the work, risk, evidence, capacity, and governing policy. It should coordinate people, AI agents, services, and organizations through:

- a shared, evidence-backed situation model;
- explicit requests, offers, commitments, renegotiation, and acceptance;
- outcome-to-evidence traceability;
- capability-limited agent execution;
- local policy and nested governance;
- selective federation of boundary contracts;
- boards, tasks, graphs, schedules, and agent queues generated as views.

Its first product is a dynamic shared situation and commitment workspace across Git repositories and documents. The same protocol supports one person directing many agents, balanced teams, and human-led work.
