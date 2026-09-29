# Trail pilot scorecard

Copy this section for each pilot. Register thresholds before the pilot begins.

## Pilot identity

| Field | Value |
|---|---|
| pilot ID | |
| team alias | |
| workflow | |
| start/end | |
| pilot owner | |
| acceptance owner | |
| participating roles | |
| agent/runtime tools | |
| repositories/systems | |
| risk class | |
| data classification | |

## Hypothesis

```text
For [specific user] performing [specific workflow],
Trail's [specific behavior] will improve [primary measure]
from [baseline] to [threshold], without worsening [guardrail].
```

## Pre-registered measures

Choose one primary metric. Secondary metrics help explain the result and cannot replace a failed primary metric after the fact.

| Measure | Definition | Baseline | Pass threshold | Result | Pass? |
|---|---|---:|---:|---:|---|
| primary | | | | | |
| time to commitment | request until responsible party and criteria are unambiguous | | | | |
| time to next action | time to correctly identify the next authorized ready action | | | | |
| resumption time | interruption until successor makes a correct next decision | | | | |
| duplicate/conflicting work | occurrences per outcome | | | | |
| false completion | completion claims rejected for missing/incorrect result | | | | |
| review attention | active reviewer minutes per accepted outcome | | | | |
| escaped acceptance error | material defect found after acceptance | | | | |
| human attention | total active human coordination minutes | | | | |
| agent cost | model/compute/tool cost per accepted outcome | | | | |
| reconstruction completeness | required facts correctly reconstructed | | | | |
| repeat use | participants voluntarily start a second workflow | no | yes | | |

## Safety and quality guardrails

Any critical failure blocks a pass regardless of the primary metric.

| Guardrail | Result | Evidence |
|---|---|---|
| no unauthorized side effect | | |
| no secret or prohibited-data exposure | | |
| exact accepted version remains identifiable | | |
| rejection preserves work and rationale | | |
| acceptance owner remains in control | | |
| participant can stop/take over | | |
| audit reconstruction is consistent | | |
| evidence limitations remain visible | | |
| accessibility blockers | | |
| approval/notification burden | | |

## Comprehension observations

| Question | Correct without help? | Time | Confidence | Notes |
|---|---|---:|---:|---|
| What is the current outcome version? | | | | |
| Who committed to what? | | | | |
| What may act next? | | | | |
| Why is this item blocked or uncertain? | | | | |
| What can this agent access and change? | | | | |
| Did execution finish? | | | | |
| Does evidence satisfy each criterion? | | | | |
| Has the result been accepted, by whom, and for which version? | | | | |
| How would you resume or contest this work? | | | | |

## Evidence inventory

| Evidence ID | Claim supported | Source | Redaction | Integrity/provenance | Limitation |
|---|---|---|---|---|---|
| | | | | | |

## Qualitative findings

### Behavior that created value


### Behavior that added overhead


### Existing tool that performed better


### Vocabulary participants actually used


### Strongest evidence against Trail


### Unexpected use or failure


## Decision

Select one and provide evidence:

- **Repeat:** run the same workflow with another outcome.
- **Narrow:** retain only the behavior responsible for measured value.
- **Revise:** change vocabulary or interaction and rerun the prototype.
- **Technical spike:** value exists but one safety/correctness uncertainty blocks progress.
- **Stop:** current alternatives solve the problem or costs exceed benefit.

### Rationale


### Required follow-up


### Decision owner and date


