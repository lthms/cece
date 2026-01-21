---
description: Execute work on an issue with an existing plan
---

<policy>
  clarification: ask
  approval: continue
  blocker: ask
</policy>

# Progress Mode

## Mode Properties

| Property | Value |
|----------|-------|
| Indicator | 🔥 |
| Arguments | `<issue-ref>` — issue number or URL (required) |
| Exit | Task completion, or user sends `stop` |
| Scope | Execute planned work independently |
| Persistence | Updates Plan comment; can post comments |
| Resumption | Re-invoke with same issue-ref |

## Principles

**Requirements are commitments.** Once the user approves the plan, every Definition
of Done item is a promise. NEVER drop, weaken, or partially implement a requirement
without explicit user approval. If you cannot deliver exactly what was agreed,
raise it as a blocker — do not silently reduce scope.

**Persistence over convenience.** When implementation is difficult, explore
alternative approaches before concluding something is impossible. Only raise a
blocker after genuine effort to solve the problem.

---

## Permissions

Execute freely without asking for approval once work begins.

**Allowed:**
- Read files, search code
- Write/edit files
- Run tests
- Create branches (per naming convention)
- Commit to your branches
- Push per `## Git Strategy` in `.claude/cece.local.md`
- Create/update PRs
- Post issue comments
- Edit own Plan comment (to check off completed PRs)

**NEVER:**
- Commit to `main` or `master`
- Push to unauthorized remotes
- Close issues
- Merge PRs
- Edit issue description (use `/cece:scope` for that)
- Edit Design comment (use `/cece:design` for that)
- Violate an Architectural Decision without raising a blocker

---

## Artifacts

### Goal

The introduction of the issue — the opening text before any sections. Created by
`/cece:scope`. This explains the problem, context, and desired outcome.

**Read-only:** Only `/cece:scope` creates or modifies this section.

### Definition of Done

A `## Definition of Done` section in the issue description. Created by
`/cece:scope`. These define what "done" means.

**Read-only:** Only `/cece:scope` modifies this section. If you discover that a
requirement needs changing, tell the user to run `/cece:scope`.

**NEVER** check off Definition of Done boxes. Only the user checks items off.

### Design Comment

A comment on the issue containing Approach, Architectural Decisions, and Q&A.
Created by `/cece:design`.

**Read-only:** Only `/cece:design` modifies this comment. If you discover
constraints that should be recorded in Q&A or require changes to Architectural
Decisions, tell the user to run `/cece:design`.

### Plan Comment

A comment on the issue with the `## Plan` heading. Created by `/cece:plan`,
updated by `/cece:progress`.

**Update when:**
- PR is completed (check it off, add link)
- PR scope changes based on review feedback

### Comments

Posted on the issue or PRs during execution.

**Where to post:**
- Blockers → on the PR if one exists, otherwise on the issue
- Progress updates → on the issue (optional, for long tasks)
- Review responses → on the PR threads

---

## Workflow

### Usage

```
/cece:progress <issue-ref>
```

Argument is required. The issue must have:
- Definition of Done (from `/cece:scope`)
- Design comment (from `/cece:design`)
- Plan comment (from `/cece:plan`)

### Step 1: Load context

1. Read `## Project Management` in `.claude/cece.local.md` to determine the platform
2. If the URL's tracker does not match your configured tracker, tell the user
   and ask whether to proceed with the mismatched tracker or stop
3. Fetch the issue (content, comments, labels, linked PRs)
4. Read the Definition of Done section from the issue description
5. Find the Design comment posted by your configured account (from `## Identity`
   in `.claude/cece.local.md`)
6. Find the Plan comment posted by your configured account

**If Definition of Done is missing:**

<response>
🔥 This issue has no Definition of Done. Run `/cece:scope <issue-ref>` first.
</response>

Return to chat mode.

**If Design comment is missing:**

<response>
🔥 This issue has no design. Run `/cece:design <issue-ref>` first.
</response>

Return to chat mode.

**If Plan comment is missing:**

<response>
🔥 No plan found for this issue. Run `/cece:plan <issue-ref>` first.
</response>

Return to chat mode.

### Step 2: Validate and resume

1. Read the Approach, Architectural Decisions, and Q&A from the Design comment
2. Read the Plan comment (task, test plan, planned PRs)
3. Parse current state: which PRs are done, pending, any blockers
4. Check open PRs for unaddressed reviews
5. Present summary to user: what's planned, done, remaining, pending reviews
6. Announce:

<response>
🔥 Resuming progress on issue.
</response>

7. Proceed to Step 3. (Step 4 applies when reviews arrive during execution.)

### Step 3: Execution

**Before any implementation:**

1. Extract every Definition of Done item from the issue description
2. Create a todo list in conversation with one item per Definition of Done requirement
3. These todos track requirement coverage — mark each complete only after the
   code is committed (and tests pass, unless waived)

Work through each planned PR:

1. **Branch**: Create or checkout branch per naming convention in `.claude/cece.local.md`
2. **Git setup**: Read `## Git Strategy` from `.claude/cece.local.md` and prepare
3. **Implement**: Write code to implement the planned PR, committing as you progress
4. **Test**: Execute the test plan. If tests fail, fix before proceeding.
   - If test plan says "User approved: no tests", skip testing for this PR
   - If test plan cannot be executed for other reasons, raise as blocker
5. **PR**: When PR scope is complete:
   - **Gate**: Before creating the PR, confirm which Definition of Done items this
     PR implements. Verify the PR fully implements those items. If incomplete,
     either complete the missing work, split across multiple PRs, or raise a
     blocker if a constraint prevents completion.
   - Create PR linking to the issue ("Fixes #N" or "Part of #N")
   - Assign user as reviewer
   - Update Plan comment: check off completed PR, add link
6. **Repeat** for remaining PRs

### Step 4: Handling Reviews

When PR reviews come in, evaluate each comment:

1. Does it change what "done" means? → <clarification>This review feedback changes the Definition of Done — should I implement it?</clarification>
2. Would it violate an Architectural Decision? → <blocker>This review feedback conflicts with an Architectural Decision — how should I proceed?</blocker>
3. Does it add work beyond the planned scope? → <clarification>This review feedback adds work beyond the planned scope — should I implement it?</clarification>
4. Otherwise → Implement the change

NEVER decline review feedback without user approval. If you believe a comment
should not be addressed, present your reasoning to the user and request approval
before declining.

After addressing comments:

4. Push fixes to your branch per `## Git Strategy` in `.claude/cece.local.md`
5. In each thread, explain what you changed or why you declined the feedback (with user approval)
6. Update the Plan comment if PR scope changed based on review
7. If review requires changes to Definition of Done, tell the user to run `/cece:scope`
8. If review requires changes to Approach, Architectural Decisions, or Q&A, tell the
   user to run `/cece:design`

### Step 5: Blockers

A blocker is anything that prevents full implementation of a requirement:
- Tests fail unexpectedly
- Design question emerges
- Missing information
- Technical constraints that force a compromise
- Implementation would violate an Architectural Decision

**NEVER silently compromise.** If you cannot implement exactly what was asked,
raise it as a blocker. Partial solutions require explicit user approval.

When blocked:

1. Post blocker as comment (on PR if exists, otherwise on issue)
2. <blocker>Cannot implement the requirement as specified — what constraint prevents completion and how should I proceed?</blocker>
3. Present options when possible
4. Once user approves an option, continue
5. If the decision should be recorded, tell the user to run `/cece:design` to
   update Q&A

### Step 6: Completion

When all planned PRs are created:

1. **Pre-check**: Re-fetch the Definition of Done from the issue description. For
   each item, identify the specific code and tests that implement it. Concrete
   implementation means: the feature works in code, tests pass (unless waived),
   and the requirement is fully addressed — not partially. If you cannot point to
   concrete implementation, the item is not met — raise a blocker before proceeding.
2. Verify all PRs are checked off in the Plan comment
3. Run the full test plan to verify all PRs work together (skip if "User approved: no tests")
4. **Review each Definition of Done item:**
   - Confirm the implementation meets the requirement exactly
   - If any item is not fully satisfied, raise a blocker
   - NEVER declare completion with unmet requirements
5. Present final summary: what was delivered, how each Definition of Done item was met

Return to chat mode.

Announce:

<response>
🐱 All work complete for issue #<N>. [Summary of what was delivered]
</response>

**NEVER** mark Definition of Done checkboxes complete — only the user does that.

NEVER close issues; closure happens automatically when PRs merge.
