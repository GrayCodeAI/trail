# Architecture decision records

ADRs capture durable choices and the evidence that could reverse them. `Accepted` means the direction is settled for current planning. It does not bypass the research and implementation gates.

| ADR | Decision | Status |
|---|---|---|
| [0001](./0001-coordination-product-boundary.md) | Trail owns outcome and commitment coordination | Accepted |
| [0002](./0002-event-journal-and-projections.md) | accepted event journal with rebuildable projections | Proposed |
| [0003](./0003-selective-federation.md) | selective home-authority federation without global consensus | Accepted |
| [0004](./0004-dynamic-human-agent-participation.md) | human-agent participation is dynamic and stage-specific | Accepted |
| [0005](./0005-modular-monolith-first.md) | begin as a modular monolith | Proposed |
| [0006](./0006-bounded-executable-domain-experiment.md) | validate core semantics in a bounded in-memory reference kernel | Accepted |
| [0007](./0007-independent-graycodeai-product.md) | Trail is independently deployed under the GraycodeAI product family | Accepted |

## ADR lifecycle

- `Proposed`: under research or waiting for a gate.
- `Accepted`: current decision; implementation and docs align to it.
- `Superseded`: replaced by another ADR, with history preserved.
- `Rejected`: considered and deliberately not chosen.
- `Deprecated`: still present for compatibility but not for new use.

An ADR includes reconsideration triggers so evidence can change the choice without rewriting history.
