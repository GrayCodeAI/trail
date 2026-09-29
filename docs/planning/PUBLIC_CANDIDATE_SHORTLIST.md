# Trail public interview candidate shortlist

**Prepared:** 2026-09-28  
**Status:** research shortlist only; nobody has been contacted  
**Rule:** use public community channels with moderator permission where appropriate; never use bug reports or unrelated pull requests for recruitment

## Cohort design

The first 20 interviews should not all be competitors or protocol builders. Use three cohorts:

1. **Operators/design partners:** people doing real work with several coding agents. Their behavior tests demand.
2. **Domain experts:** maintainers building agent-native work products. Their experience reveals failure patterns and crowded features.
3. **Infrastructure experts:** durable execution and authorization maintainers. Their input tests technical assumptions, not product demand.

Target at least 10 operator interviews, six domain-expert interviews, and four infrastructure interviews. A maintainer may fit several cohorts, but count each interview once according to its primary evidence.

## Scoring

Score out of 10:

- workflow relevance: 0–3;
- likelihood of concrete artifacts/experience: 0–3;
- appropriate public access route: 0–2;
- independence from Trail's proposed solution: 0–2.

A high score determines outreach order. It does not imply endorsement or willingness to participate.

## A. Operator and design-partner pools

These entries identify public communities rather than claiming access to their users. Ask moderators before posting a research invitation in a community.

| ID | Public pool | Target participant | Why this workflow matters | Public route | Score | Priority |
|---|---|---|---|---|---:|---|
| OP-01 | [OpenHands](https://github.com/OpenHands/OpenHands) community | maintainer or team using issue-to-PR agents across maintained repositories | real issue, agent, sandbox, PR, review, interruption, and acceptance trail | [OpenHands Slack](https://go.openhands.dev/slack) | 10 | P0 |
| OP-02 | [Cline](https://github.com/cline/cline) community | developer using sessions, subagents/teams, approval policies, or scheduled automation | explicit sessions, cost, checkpoints, approvals, resumption, and team-agent features | [GitHub Discussions](https://github.com/cline/cline/discussions) or [Discord](https://discord.gg/cline) | 10 | P0 |
| OP-03 | [OpenCode](https://github.com/anomalyco/opencode) community | maintainer operating the agent on several repositories or concurrent changes | large active user pool for agent work, review, handoff, and local control | [OpenCode Discord](https://opencode.ai/discord) | 9 | P0 |
| OP-04 | [Paperclip](https://github.com/paperclipai/paperclip) community | operator who has run several agents with budgets, approvals, and issues | directly exercises agent organization, control, and governance pain | [GitHub Discussions](https://github.com/paperclipai/paperclip/discussions) or [Discord](https://discord.gg/m4HZY7xNG3) | 9 | P0 |
| OP-05 | [Mission Control](https://github.com/builderz-labs/mission-control) community | team supervising long-running agents across runtimes | run visibility, spend, failures, alerts, schedules, and operator recovery | [GitHub Discussions](https://github.com/builderz-labs/mission-control/discussions) | 9 | P0 |
| OP-06 | [It's a Plan](https://github.com/croffasia/itsaplan) community | self-hosting team assigning issues to internal or external coding agents | conventional PM plus permissions, runs, MCP, webhooks, and five forge integrations | [GitHub Discussions](https://github.com/croffasia/itsaplan/discussions) | 9 | P0 |
| OP-07 | [Aider](https://github.com/Aider-AI/aider) community | OSS maintainer using Aider repeatedly across branches or repositories | independent high-use coding-agent workflow; useful counterweight to multi-agent-product enthusiasts | [Aider Discord](https://discord.gg/Y7X7bhMQFV) | 8 | P1 |
| OP-08 | [Hatchet](https://github.com/hatchet-dev/hatchet) community | application team running durable AI-agent workflows in production | concrete retries, queues, concurrency, rate limits, failures, and observability | [GitHub Discussions](https://github.com/hatchet-dev/hatchet/discussions) or [Discord](https://hatchet.run/discord) | 8 | P1 |
| OP-09 | [Restate](https://github.com/restatedev/restate) community | team operating durable agents or stateful workflows | concrete experience with recovery, replay, idempotency, and long-running state | [GitHub Discussions](https://github.com/restatedev/restate/discussions) or [community Slack](https://join.slack.com/t/restatecommunity/shared_invite/zt-2v9gl005c-WBpr167o5XJZI1l7HWKImA) | 8 | P1 |
| OP-10 | [Agent Base](https://github.com/AgentOrchestrator/AgentBase) users | developer operating parallel agents with isolated edits and centralized approvals | close match for interruption, conflicting work, shared context, and attention-queue problems | repository's documented contribution/feedback route; ask maintainers before recruiting | 7 | P1 |
| OP-11 | [Agent Orchestrator](https://github.com/useagent/ao) users | team supervising worktree-isolated coding agents and PR repair | direct parallel-work and human-escalation workflow | [project Discord](https://discord.gg/UZv7JjxbwG) | 7 | P2 |
| OP-12 | [Plane](https://github.com/makeplane/plane) community | team augmenting an OSS issue tracker with external coding agents | tests whether an existing tracker plus integrations already solves the Trail wedge | [Plane forum](https://forum.plane.so) or [organization discussions](https://github.com/orgs/makeplane/discussions) | 8 | P1 |

### Operator selection filters

Prefer participants who can show a workflow from the last 30 days and meet at least three conditions:

- two or more concurrent agent runs;
- two or more repositories or services;
- human review or acceptance responsibility;
- interruption or handoff;
- approvals, secrets, budget, or permissions;
- rejected or reopened agent work;
- a real downstream side effect;
- use by more than one person.

Avoid counting hobby experimentation with no consequential output as design-partner evidence. It can still inform usability.

## B. Domain-expert shortlist

Public handles below came from the top non-bot GitHub contributor list on 2026-09-28. Contribution rank does not prove current maintainer authority; verify the person's documented role before addressing them as a maintainer.

| ID | Project | Public contributor starting point | Topic to investigate | Public route | Score | Priority |
|---|---|---|---|---|---:|---|
| EX-01 | [Paperclip](https://github.com/paperclipai/paperclip) | `@cryppadotta`, `@devinfoley`, `@nickyleach` | atomic checkout, agent hierarchy, budget/approval burden, completion truth | project Discussions | 9 | P0 |
| EX-02 | [Paca](https://github.com/Paca-AI/paca) | `@pikann` | agents as Scrum teammates, configuration/plugins, human-agent equality limits | project Discussions | 8 | P0 |
| EX-03 | [Mission Control](https://github.com/builderz-labs/mission-control) | `@0xNyk`, `@bhavikprit` | cross-runtime runs, spend, failure recovery, schedules, and alerts | project Discussions | 9 | P0 |
| EX-04 | [It's a Plan](https://github.com/croffasia/itsaplan) | `@croffasia`, `@OthmanAdi` | agent permissions, run triggers, forge integrations, signed webhook delivery | project Discussions | 9 | P0 |
| EX-05 | [OpenHands](https://github.com/OpenHands/OpenHands) | `@hieptl`, `@tofarr` | issue-to-PR automation, sandbox boundaries, feedback/review loop, enterprise triggers | OpenHands Slack | 8 | P1 |
| EX-06 | [Cline](https://github.com/cline/cline) | `@saoudrizwan`, `@abeatrix`, `@celestial-vault` | session persistence, agent teams, tool approvals, checkpoints, cost visibility | project Discussions | 8 | P1 |
| EX-07 | [OpenCode](https://github.com/anomalyco/opencode) | `@thdxr`, `@adamdotdevin` | local agent operation, cross-session state, review and collaboration boundaries | official Discord | 7 | P1 |
| EX-08 | [Plane](https://github.com/makeplane/plane) | `@anmolsinghbhatia`, `@aaryan610`, `@sriramveeraghanta` | why teams adopt/self-host PM; import, migration, notification, and agent integration expectations | public forum/discussions | 7 | P1 |
| EX-09 | [OpenProject](https://github.com/opf/openproject) | `@oliverguenther`, `@ulferts`, `@myabc` | enterprise permissions, audit, lifecycle rigor, procurement, and data sovereignty | [OpenProject community](https://community.openproject.org/projects/openproject/forums) | 7 | P2 |
| EX-10 | [AG-UI](https://github.com/ag-ui-protocol/ag-ui) | `@ranst91`, `@contextablemark`, `@mme` | human steering, run-state events, frontend synchronization, interoperability boundaries | project Discussions or [Discord](https://discord.gg/Jd3FzfdJa8) | 7 | P1 |
| EX-11 | [A2A](https://github.com/a2aproject/A2A) | `@holtskinner`, `@kthota-g`, `@amye` | task ownership, identity, long-running work, artifacts, extension governance, cross-org gaps | [A2A Discussions](https://github.com/a2aproject/A2A/discussions) | 7 | P1 |

Domain experts can validate mechanisms and warn about failed approaches. They cannot substitute for design partners using Trail's proposed workflow.

## C. Infrastructure-expert shortlist

| ID | Project | Public contributor starting point | Question for Trail | Public route | Score | Priority |
|---|---|---|---|---|---:|---|
| IN-01 | [Hatchet](https://github.com/hatchet-dev/hatchet) | `@abelanger5`, `@mrkaye97` | which execution guarantees, concurrency controls, and recovery tools should Trail reuse? | project Discussions | 8 | P1 |
| IN-02 | [Restate](https://github.com/restatedev/restate) | `@tillrohrmann`, `@AhmedSoliman`, `@slinkydeveloper` | how should uncertain side effects, replay, idempotency, and stateful actors map to Trail? | project Discussions/community Slack | 8 | P1 |
| IN-03 | [OpenFGA](https://github.com/openfga/openfga) | `@jon-whit`, `@adriantam` | can relationship tuples express Trail membership/delegation while capabilities handle concrete runs? | [OpenFGA community](https://openfga.dev/community) | 7 | P2 |
| IN-04 | [SpiceDB](https://github.com/authzed/spicedb) | `@josephschorr`, `@vroldanbet`, `@jakedt` | consistency, caveats, revocation, schema migration, and authorization explanation tradeoffs | project Discussions or [Discord](https://authzed.com/discord) | 7 | P2 |

## First outreach wave

Recruit no more than six conversations before reviewing the interview script:

1. OpenHands operator;
2. Cline team/multi-agent operator;
3. OpenCode or Aider OSS maintainer using agents on real work;
4. Paperclip or Mission Control operator;
5. one Paperclip/Mission Control/It's a Plan domain expert;
6. one Hatchet or Restate infrastructure expert.

This mix exposes vocabulary and recruiting mistakes before approaching the whole list.

## Public recruitment message — operator

Customize this only after checking the community's rules:

> We are researching coordination failures in software teams that run several coding agents. We are looking for a 45-minute conversation about one recent real workflow: how the outcome was assigned, interrupted, reviewed, and accepted. This is research rather than a product demo. We will redact project and organization details, will not request credentials or proprietary source, and will share the synthesized findings with participants. If you have operated concurrent agent work across repositories and are open to a conversation, please reply through the channel permitted by the community moderators.

## Public invitation — expert

> We are researching the boundary between agent execution and accountable work coordination. Your project has implemented mechanisms relevant to this question. We would value a 45-minute technical interview about failure modes, design tradeoffs, and approaches you rejected. We are not asking for roadmap commitments or support, and we will attribute comments only with explicit permission.

## Outreach safeguards

- Obtain explicit user authorization before posting or sending any message.
- Read each community's rules and ask moderators when research recruitment is not clearly allowed.
- Do not scrape or use personal email addresses.
- Do not mass-DM contributors.
- Do not describe a top contributor as a maintainer without verification.
- Do not imply partnership, endorsement, or affiliation.
- Keep competitor interviews separate from confidential Trail design-partner data.
- Offer an anonymized findings summary rather than payment promises unless compensation is deliberately approved.
- Track invitation, response, consent, scheduled, completed, declined, and withdrawn states.

## Candidate tracker

| ID | Contacted? | Route | Response | Qualified? | Scheduled | Completed | Evidence IDs | Follow-up |
|---|---|---|---|---|---|---|---|---|
| OP-01 | no | | | | | | | |
| OP-02 | no | | | | | | | |
| OP-03 | no | | | | | | | |
| OP-04 | no | | | | | | | |
| OP-05 | no | | | | | | | |
| OP-06 | no | | | | | | | |
| OP-07 | no | | | | | | | |
| OP-08 | no | | | | | | | |
| OP-09 | no | | | | | | | |
| OP-10 | no | | | | | | | |
| OP-11 | no | | | | | | | |
| OP-12 | no | | | | | | | |

## Snapshot limitations

- Repository activity and stars are discovery signals and change continuously.
- A public community may prohibit research recruitment.
- Contributors may have changed roles or may not operate the product themselves.
- Competitor maintainers have unusually deep knowledge and unusually high solution bias.
- Public OSS users underrepresent regulated, private, and large-enterprise deployments.
- The shortlist must expand through referrals from qualified operators, while preserving the cohort balance.

