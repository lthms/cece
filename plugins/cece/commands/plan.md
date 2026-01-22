---
description: Plan work on an issue with concrete PRs and test strategy
---

<policy>
  clarification: ask
  approval: ask
  blocker: ask
</policy>

# Plan Mode

## Mode Properties

| Property | Value |
|----------|-------|
| Indicator | 📋 |
| Arguments | `<issue-ref>` — issue number or URL (required) |
| Exit | Plan comment posted, or user sends `stop` |
| Scope | Create concrete execution plan with test strategy |
| Persistence | Plan file + Plan comment (Task + Test plan + Planned PRs) |
| Resumption | Re-invoke with same issue-ref to revise plan |

## How Plan Mode Works

This command uses Claude's native planning features for safe, reviewable planning.

**What happens:**

1. You explore the codebase using read-only tools
2. You draft the plan to a file the user can review and edit
3. You request approval before posting
4. Once approved, the plan is posted to the issue

This gives the user complete control: they see the full plan before it affects
the issue tracker.

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
- [ ] PR 2: <scope> (depends on PR 1)
```

**Dependencies:** When a PR depends on another, add `(depends on PR N)` at the
end of its line, where N is the PR number in the planned list (1, 2, 3, etc.).
This syntax is case-sensitive. When `/cece:progress` pushes changes to a base
branch, it parses this syntax to identify and auto-rebase dependent branches.

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

### Step 3: Enter plan mode

Use `EnterPlanMode` to enter read-only planning mode. During planning, explore
the codebase using Read, Glob, Grep, and Task tools only. Do not create branches,
write code, or create commits.

### Step 4: Check for existing plan

Look for a Plan comment posted by your configured account that contains a
`## Plan` heading.

**If plan exists:**

1. Read all comments on the issue (including review feedback and blockers)
2. Check for linked PRs — if any exist, work has already started
3. Analyze whether the plan needs revision. Check:
   - Whether the Design comment was edited since the plan was posted (compare
     timestamps or content)
   - Whether there are unresolved comments on linked PRs indicating blockers
   - Whether scope changed based on feedback (new Definition of Done items or
     modified existing items)
   - Whether the tests specified in the plan still work with the current codebase
4. Present the existing plan to the user with your assessment:
   - If PRs exist: tell the user that revising the plan may require rework on
     existing PRs
   - If all checks pass: suggest proceeding to `/cece:progress`
   - If any check fails: identify which specific items changed and recommend
     revising
5. Wait for user confirmation before proceeding
6. If user wants to revise, proceed to Step 5

**If no plan:**
- Proceed to Step 5

### Step 5: Draft plan

1. Review the Approach and Architectural Decisions from the Design comment
2. Map Definition of Done items to concrete PRs
3. For each planned PR, check if implementing it would require violating an
   Architectural Decision from the Design comment. If so, raise a blocker:
   "PR N would violate Architectural Decision X — revise the plan or design?"
4. Identify dependencies between PRs — if a PR builds on changes from another,
   mark it with `(depends on PR N)`
5. Write the plan to the plan file:
   - Task: one sentence stating what the code will do when completed
   - Test plan: specific test commands to run, or manual verification steps
   - Planned PRs: breakdown of work into reviewable units, with dependencies noted

**Test plan is mandatory.** Specify either: (a) test commands with expected
outcomes, (b) manual verification steps, or (c) the literal phrase
`User approved: no tests` if testing is waived. If you cannot identify a test
approach, raise this as a blocker.

**If the Approach is infeasible** — constraints make the design impossible to
implement — <blocker>The approach cannot be implemented due to these
constraints. Run `/cece:design` to revise the approach before planning.</blocker>

### Step 6: Exit plan mode and get approval

1. Use `ExitPlanMode` to signal the plan is complete
2. This triggers Claude's native approval flow — the user reviews the plan file
3. Wait for user approval before proceeding

Do NOT post the Plan comment until the user approves.

### Step 7: Post to issue and exit

After approval:

**If creating a new plan:**

1. Post the Plan comment on the issue (using the content from the plan file)

**If updating an existing plan:**

1. Edit the existing Plan comment with the revised content
2. If changes are significant, post a comment explaining what changed

Return to chat mode.

Announce:

<response>
🐱 Plan posted to issue #<N>. Run `/cece:progress <issue-ref>` to start execution.
</response>
