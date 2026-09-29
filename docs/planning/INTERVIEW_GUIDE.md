# Trail problem interview guide

**Phase:** 0, problem validation  
**Interview length:** 45–60 minutes  
**Target:** people currently coordinating real work performed by coding agents

## Research objective

Determine whether agent-heavy software teams repeatedly experience expensive coordination failures that existing project managers, coding-agent tools, and run dashboards do not solve.

This is a problem interview. Do not pitch Trail until the workflow reconstruction is complete.

## Participant mix

Recruit 15–20 participants across at least five teams:

- engineering leads running several agents or parallel worktrees;
- developers using two or more coding-agent products;
- OSS maintainers coordinating people and bots;
- platform/security owners controlling agent credentials and environments;
- reviewers responsible for accepting generated changes;
- product or delivery owners accountable for the requested outcome.

Seek variation in team size, repositories, agent tools, deployment model, regulation, and autonomy. Record who owns the problem and who can authorize a purchase or deployment.

Use the [public candidate shortlist](./PUBLIC_CANDIDATE_SHORTLIST.md) as the initial recruiting pool while preserving the operator, domain-expert, and infrastructure-expert cohort balance.

## Screener

A participant qualifies when they can answer yes to most of these:

1. Have you used coding agents on production or maintained OSS work in the last 30 days?
2. Have two or more human/agent workers operated concurrently?
3. Has work crossed repository, service, team, or organizational boundaries?
4. Can you walk through a specific recent workflow using redacted artifacts?
5. Were you responsible for planning, execution, review, security, or acceptance?

Do not recruit only people who already agree with the Trail thesis.

## Consent and handling

Before recording or collecting artifacts, explain:

- the research purpose;
- what will be recorded;
- who can access the notes;
- how names, organizations, repositories, secrets, and customer data will be redacted;
- retention and deletion timing;
- that participation does not create a sales or pilot obligation;
- that the participant can skip a question or withdraw an artifact.

Record the consent scope in the evidence ledger. Never paste credentials, proprietary source, customer content, or production logs into the research corpus.

## Interview structure

### 1. Context — 5 minutes

- What is your role, and what outcome are you accountable for?
- Which coding agents, trackers, repositories, CI systems, and communication tools do you use?
- How many agent runs or agent-authored changes happen in a typical week?
- Who decides what agents may do?

### 2. Last concrete workflow — 20 minutes

Ask for the most recent meaningful outcome, then reconstruct it chronologically.

- What triggered the work?
- What was the desired outcome? How was success stated?
- Where did that intent live?
- Who agreed to do what? Was the agreement explicit?
- How was work decomposed and assigned?
- What could each human or agent access and change?
- How did a worker know it was ready to start?
- Which parts happened concurrently?
- What changed while work was underway?
- Where did decisions, assumptions, and exceptions get recorded?
- What did an interrupted participant need to resume?
- What did “done” mean to the executor?
- Which evidence did the reviewer inspect?
- Who had authority to accept the result?
- Was the exact accepted version recorded?
- What happened after acceptance or rejection?

Ask to see the relevant ticket, chat, transcript, run log, pull request, CI result, or checklist when consent permits. Observe before asking how it should be improved.

### 3. Failure reconstruction — 15 minutes

Ask for the last incident involving one or more of:

- duplicate or conflicting work;
- an agent acting without current authority;
- stale assumptions or context loss;
- work marked complete without acceptable evidence;
- unclear reviewer or approval owner;
- budget or rate-limit surprise;
- unsafe retry or duplicated external side effect;
- failure to resume after interruption;
- cross-team or cross-organization information loss;
- excessive approvals, notifications, or review backlog.

For that incident:

1. What was the first observable signal?
2. Which fact was missing, wrong, stale, private, or disputed?
3. Who noticed and how long did it take?
4. What was the consequence in time, money, quality, security, or trust?
5. Which workaround was used?
6. Why did the current tools not prevent or reveal it?
7. Has it happened before?
8. What would the team pay or change to avoid it?

### 4. Existing alternatives — 5 minutes

- What have you configured or purchased to solve this?
- Which part works well today?
- Why have you not solved the remaining part with Jira, Linear, GitHub, Slack, an agent control plane, or a script?
- What switching, security, or procurement constraint matters?
- If nothing changed, what would happen over the next six months?

### 5. Commitment — 5 minutes

Ask for behavioral evidence:

- May we observe or reconstruct another real workflow?
- Would you test a manual prototype with redacted data?
- Who else must participate for the workflow to be real?
- What result would justify a second use?
- Who could approve installation or purchase if it worked?

Do not count compliments or hypothetical willingness as a pilot commitment.

## Artifact reconstruction

For each workflow, capture:

| Field | Evidence sought |
|---|---|
| trigger | message, incident, request, issue, schedule, or observed condition |
| outcome | original words and later changes |
| criteria | explicit, inferred, changed, or missing |
| participants | people, agents, services, teams, and organizations |
| commitments | promise, owner, beneficiary, deadline, conditions |
| authority | identity, permission, delegation, approval, revocation |
| execution | runs, tools, environments, costs, retries, side effects |
| decisions | options, rationale, author, policy, effective version |
| evidence | artifacts, tests, reviews, attestations, provenance |
| acceptance | reviewer, scope, exact version, result, dispute |
| failures | delay, duplication, ambiguity, leak, rework, escaped defect |
| recovery | detection, takeover, compensation, resumption time |

## Post-interview scoring

Score independently before group discussion.

| Dimension | 0 | 1 | 2 | 3 |
|---|---|---|---|---|
| frequency | never | yearly | monthly | weekly/daily |
| consequence | negligible | inconvenience | material cost/rework | security, revenue, customer, or release impact |
| workaround | easy/configured | manual but adequate | painful and repeated | failing or impossible |
| ownership | nobody | user cares | named operational owner | owner controls budget/adoption |
| evidence | opinion only | one anecdote | artifacts for one workflow | repeated artifacts across workflows |
| commitment | compliment | follow-up call | prototype session | real pilot workflow/resources |

Record uncertainty and counter-evidence. A high total score is a prioritization signal, not proof by itself.

## Interviewer safeguards

- Ask about past behavior before future preference.
- Avoid introducing commitment/evidence terminology prematurely.
- Do not explain away a participant's confusion.
- Separate a requested feature from the underlying job.
- Ask what works well in current tools.
- Seek teams that do not have the problem.
- Record the participant's words before translating them into Trail vocabulary.
- Never infer willingness to pay from enthusiasm.

## Required output

Within 24 hours, add a redacted evidence entry, workflow reconstruction, score, strongest counter-evidence, and next action under `research/evidence/`. Link each supported or challenged hypothesis from the research plan.
