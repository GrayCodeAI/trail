# ADR 0004: Human-agent participation is dynamic and stage-specific

- **Status:** Accepted
- **Date:** 2026-09-28

## Context

A fixed 80/20 or 90/10 human-agent ratio ignores task risk, capability, uncertainty, policy, and stage. One workflow can be mostly automated in execution while retaining human authority over intent and acceptance. Another can use agents only for analysis.

Human factors research also shows that automation can remove situation awareness and leave people with rare, difficult interventions.

## Decision

Trail measures participation; it does not enforce a target ratio. Authority is configured separately for observe, analyze, propose, decide, execute, verify, and accept.

Risk, policy, confidence, budget, capability, and attention capacity can narrow or expand delegated authority within fixed safety bounds. People retain visible takeover, revocation, explanation, and appeal paths.

## Consequences

- a workspace-wide autonomy percentage is prohibited;
- the UI must show current authority and why human attention is requested;
- resumption and takeover quality are product metrics;
- automated allocation needs hysteresis and rate limits;
- participation metrics must not become hidden performance scores.

## Reconsider if

- users cannot understand stage-specific authority despite tested interaction improvements;
- dynamic policies cause more error or cognitive load than a small set of fixed, scoped modes;
- regulatory contexts require fixed human control for a particular task class, in which case scoped policy can enforce it without changing the general model.

