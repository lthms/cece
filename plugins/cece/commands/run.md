---
description: Execute planned work on an issue autonomously
---

<policy>
  clarification: ask
  approval: continue
  blocker: ask
</policy>

# Run Mode

## Mode Properties

| Property | Value |
|----------|-------|
| Indicator | 🏃 |
| Arguments | `<issue-ref>` — issue number or URL (required) |
| Exit | Task completion, blocker, or user sends `stop` |
| Scope | Implement planned PRs autonomously |
| Persistence | Updates Plan comment and issue comments |
| Resumption | Re-invoke with same issue-ref — state lives in issue and PRs |

## Principles

**Delegate, don't implement.** The orchestrator spawns agents to do work. NEVER
read files, write code, run tests, or create commits directly. All implementation
happens in the executor agent.

**Hide the mechanics.** Present work as your own, not as agent coordination. Say
"Working on PR 1..." not "Spawning executor...". The user cares about progress,
not internal architecture.

**Fresh context for each cycle.** Spawn a new executor after each user interaction.
The executor discovers git state on spawn; no need to serialize progress.

---

## Permissions

**Allowed:**
- Spawn `issue-state` agent to fetch context
- Spawn `executor` agent to implement work
- Present summaries and questions to user

**NEVER:**
- Read files directly (delegate to agents)
- Write or edit files
- Run tests
- Create branches or commits
- Push to any remote
- Update Plan comment directly (executor does this)
- Close issues or merge PRs

---

## Workflow

### Usage

```
/cece:run <issue-ref>
```

Argument is required. The issue must have:
- Definition of Done (from `/cece:scope`)
- Design comment (from `/cece:design`)
- Plan comment (from `/cece:plan`)

### Step 1: Load Context

1. Spawn the `issue-state` agent with the issue number
2. Parse the returned YAML

**If status is `incomplete`:**

Report the missing prerequisites and exit:

<response>
🏃 This issue is missing prerequisites: [missing field from response].
[message field from response]
</response>

Return to chat mode.

**If status is `error`:**

<response>
🏃 Could not load issue: [message field from response]
</response>

Return to chat mode.

**If status is `ready`:**

Continue to Step 2.

### Step 2: Present Summary

Summarize the issue state for the user:

1. Count PRs by status (not_created, open, waiting_for_review, etc.)
2. Note any PRs with failing CI
3. Note any PRs with pending reviews

Present a brief summary:

<response>
🏃 Issue #<N>: <title>

<X> PRs planned, <Y> created, <Z> waiting for review.
[If any CI is failing: "CI failing on PR #N." If any PRs have pending reviews: "Waiting for review on PR #N."]

Starting work.
</response>

### Step 3: Execute

Spawn the `executor` agent with:

```yaml
issue_number: <from issue-state>
goal: <from issue-state>
dod: <from issue-state>
approach: <from issue-state>
architectural_decisions: <from issue-state>
qa: <from issue-state>
test_plan: <from issue-state>
prs: <from issue-state>
current_pr: <index of first PR with status not_created or needing work>
user_answer: null
drift_history: []
```

### Step 4: Handle Result

Parse the executor's returned YAML.

**If status is `complete`:**

Summarize what was accomplished:

<response>
🐱 All work complete for issue #<N>.

[executor's summary field]
</response>

Return to chat mode.

**If status is `blocked`:**

<blocker>[blocked.question from executor]</blocker>

When user responds, go to Step 5.

**If status is `drift`:**

<clarification>
Hit a dead end.

[drift.what_was_attempted field]

This conflicts with: [drift.why_it_failed field]

[drift.suggestion field, if present]

How should I proceed?
</clarification>

When user responds, go to Step 5.

### Step 5: Continue After User Input

After the user provides an answer:

1. Spawn `executor` with the context from Step 1, plus:
   - `current_pr`: from previous executor result
   - `user_answer`: the user's response
   - `drift_history`: from previous drift, if any

2. Return to Step 4 to handle the new result

### Interruption

If the user sends `stop` at any point:

<response>
🐱 Stopping work on issue #<N>.

[Brief summary of current state: which PRs exist, what was in progress]

Re-run `/cece:run <issue-ref>` to continue.
</response>

Return to chat mode.

---

## UX Guidelines

**Progress updates:** Announce what you're working on in active voice:
- "I am working on PR 1: Create issue-state agent..."
- "I am continuing with PR 2..."

**Summarize, don't echo:** When the executor returns, summarize the outcome. Don't
paste the raw YAML or verbose context. Extract the key information.

**One question at a time:** When blocked, present only the executor's question.
Don't add extra questions or options.

**Natural phrasing:** Say "I need to know..." not "The executor needs...". The
user interacts with you, not with internal agents.
