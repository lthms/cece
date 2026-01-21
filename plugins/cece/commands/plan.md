---
description: Plan work on an issue with concrete PRs and test strategy
---

# Plan Mode

## Mode Properties

| Property | Value |
|----------|-------|
| Indicator | 📋 |
| Arguments | `<issue-ref>` — issue number or URL (required) |
| Exit | Plan comment posted, or user sends `stop` |
| Scope | Create concrete execution plan with test strategy |
| Persistence | Plan comment (Task + Test plan + Planned PRs) |
| Resumption | Re-invoke with same issue-ref to revise plan |

## Permissions

**Allowed:**
- Read files, search code
- Fetch issues
- Post issue comments
- Edit own Plan comment

**NEVER:**
- Create branches
- Write code
- Create PRs
- Create issues (use `/cece:scope` first)
- Edit issue description (use `/cece:scope` for that)
- Edit Design comment (use `/cece:design` for that)

---

## Artifacts

### Goal

The introduction of the issue — the opening text before any sections. Created by
`/cece:scope`. This explains the problem, context, and desired outcome.

**Read-only:** Only `/cece:scope` creates or modifies this section.

### Definition of Done

A `## Definition of Done` section in the issue description. Created by
`/cece:scope`. These define what "done" means.

**Read-only:** Only `/cece:scope` creates or modifies this section.

### Design Comment

A comment on the issue containing Approach, Architectural Decisions, and Q&A.
Created by `/cece:design`.

**Read-only:** Only `/cece:design` creates or modifies this comment.

### Plan Comment

A comment on the issue with the `## Plan` heading. This is the single artifact
owned by `/cece:plan`.

**Required sections:**
- Task (one sentence summary)
- Test plan (specific validation method: test commands, manual steps, or "User approved: no tests")
- Planned PRs (checkboxes with scope descriptions)

**Format:**
```markdown
## Plan

**Task**: <summary>

**Test plan**: <validation method, or "User approved: no tests" if waived>

**Planned PRs**:
- [ ] PR 1: <scope>
- [ ] PR 2: <scope>
```

---

## Workflow

### Usage

```
/cece:plan <issue-ref>
```

Argument is required. The issue must have:
- Definition of Done (from `/cece:scope`)
- Design comment (from `/cece:design`)

### Step 1: Load the issue

1. Read `## Project Management` in `.claude/cece.local.md` to determine the platform
2. If the URL's tracker does not match your configured tracker:
   <clarification>This issue is on a different tracker than configured — should
   I proceed or stop?</clarification>
3. Fetch the issue (content, comments, labels, linked PRs)

Announce:

<response>
📋 Switching to plan mode.
</response>

### Step 2: Validate issue readiness

1. Read the Definition of Done section from the issue description
2. Find the Design comment posted by your configured account (from `## Identity`
   in `.claude/cece.local.md`)
3. Read the Approach, Architectural Decisions, and Q&A from the Design comment

**If Definition of Done is missing or empty:**

<response>
📋 This issue has no Definition of Done. Run `/cece:scope <issue-ref>` first to define requirements.
</response>

Return to chat mode.

**If Design comment is missing:**

<response>
📋 This issue has no design. Run `/cece:design <issue-ref>` first to settle on an approach.
</response>

Return to chat mode.

### Step 3: Check for existing plan

Look for a Plan comment posted by your configured account that contains a
`## Plan` heading.

**If plan exists:**

1. Read all comments on the issue (including review feedback, blockers, updates)
2. Check for linked PRs — if any exist, work has already started
3. Analyze whether the plan needs revision:
   - Has the Design been updated since the plan was created?
   - Are there unresolved blockers?
   - Has feedback suggested scope changes?
   - Is the test plan still valid?
4. Present the existing plan to the user with your assessment:
   - If PRs exist: warn that revising the plan may require rework on existing PRs
   - If all checks pass (no design changes, no blockers, no scope changes, test plan valid): suggest proceeding to `/cece:progress`
   - If any check fails: identify which specific items changed and recommend revising
5. Wait for user to confirm their intent before proceeding
6. If user wants to revise, proceed to Step 4

**If no plan:**
- Proceed to Step 4

### Step 4: Draft plan

1. Review the Approach and Architectural Decisions from the Design comment
2. Map Definition of Done items to concrete PRs
3. Verify planned PRs respect Architectural Decisions — if a PR would violate one,
   raise this with the user before finalizing the plan
4. Draft the Plan comment:
   - Task: one sentence stating what the code will do when completed
   - Test plan: specific test commands to run, or manual verification steps
   - Planned PRs: breakdown of work into reviewable units
5. Present plan to the user in conversation
6. Iterate based on feedback until user is satisfied

**Test plan is mandatory.** NEVER proceed without specifying either: (a) test
commands with expected outcomes, (b) manual verification steps, or (c) explicit
user approval for "User approved: no tests". If you cannot identify a test
approach, raise this as a blocker before finalizing the plan.

**If the Approach is infeasible** — constraints make the design impossible to
implement — <blocker>The approach cannot be implemented due to these
constraints. Run `/cece:design` to revise the approach before planning.</blocker>

### Step 5: Sign-off

1. <approval>Ready to post this plan to the issue?</approval>
2. Wait for user approval before posting

Do NOT post the Plan comment until the user approves.

### Step 6: Post to issue and exit

After sign-off:

**If creating a new plan:**

1. Post the Plan comment on the issue

**If updating an existing plan:**

1. Edit the existing Plan comment with the revised content
2. Optionally post a brief comment noting what changed (if changes are significant)

Return to chat mode.

Announce:

<response>
🐱 Plan posted to issue #<N>. Run `/cece:progress <issue-ref>` to start execution.
</response>
