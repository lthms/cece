---
description: Execute work on an issue with an existing plan
---

# Progress Mode

## Mode Properties

| Property | Value |
|----------|-------|
| Indicator | 🔥 |
| Arguments | `<issue-ref>` — issue number or URL (required) |
| Exit | Task completion, or user sends `stop` |
| Scope | Execute planned work independently |
| Persistence | Updates Plan comment + Q&A section on the issue |
| Resumption | Re-invoke with same issue-ref |

## Principles

**Requirements are commitments.** Once the user approves the plan, every success
criterion is a promise. NEVER drop, weaken, or partially implement a requirement
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
- Push per `## Git Strategy` in `cece.local.md`
- Create/update PRs
- Post issue comments

**NEVER:**
- Commit to `main` or `master`
- Push to unauthorized remotes
- Close issues
- Merge PRs

---

## Artifacts

### Plan

A comment on the issue with the `## Work Plan` heading. Created by `/cece:plan`.

**Update when:**
- PR is completed (check it off, add link)
- Success criteria change
- Scope changes

### Q&A

A `## Q&A` section in the issue description. This is the decision log — key
decisions, constraints, and learnings.

**Purpose:** Preserve context across sessions. When resuming, read this first and
refer to prior decisions before proposing changes to settled approaches.

**What belongs:**
- Constraints discovered during implementation ("API doesn't support batch ops")
- Clarifications from user that affect approach
- Learnings from PR reviews that apply to remaining work

**Update when:**
- After blocker resolution (constraints discovered)
- After significant review feedback (learnings)

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

Argument is required. The issue must have an existing Plan (created by
`/cece:plan`).

### Step 1: Load plan

1. Read `## Project Management` in `cece.local.md` to determine the platform
2. Fetch the issue (content, comments, labels, linked PRs)
3. Find the Plan comment posted by your account

**If no plan exists:**

<response>
🔥 No plan found for this issue. Run `/cece:plan <issue-ref>` first.
</response>

Return to chat mode.

### Step 2: Validate and resume

1. Read the Q&A section from the issue description
2. Validate plan completeness; if incomplete, ask user whether to run
   `/cece:plan` to update it before proceeding
3. Parse current state: which PRs are done, pending, any blockers
4. Check open PRs for unaddressed reviews
5. Present summary to user: what's planned, done, remaining, pending reviews

Announce:

<response>
🔥 Resuming progress on issue.
</response>

6. If reviews are pending, proceed to Step 4. Otherwise, proceed to Step 3.

### Step 3: Execution

Work through each planned PR:

1. **Branch**: Create or checkout branch per naming convention in `cece.local.md`
2. **Git setup**: Read `## Git Strategy` from `cece.local.md` and prepare
3. **Implement**: Write code, commit freely
4. **Test**: Execute the test plan. If tests fail, fix before proceeding. If
   test plan cannot be executed, raise as blocker — do not skip.
5. **PR**: When scope is complete:
   - Create PR linking to the issue ("Fixes #N" or "Part of #N")
   - Assign user as reviewer
   - Update Plan: check off completed PR, add link
6. **Repeat** for remaining PRs

### Step 4: Handling Reviews

When PR reviews come in, evaluate each comment:

1. Does it change what "done" means? → Ask user before implementing
2. Does it require modifying architectural decisions? → Ask user before implementing
3. Otherwise → Implement the change

NEVER decline review feedback without user approval. If you believe a comment
should not be addressed, present your reasoning to the user and request approval
before declining.

After addressing comments:

4. Push fixes
5. Reply to each thread: explain what changed or why declined
6. Update Plan if scope or criteria changed
7. Update Q&A if review revealed significant learnings

### Step 5: Blockers

A blocker is anything that prevents full implementation of a requirement:
- Tests fail unexpectedly
- Design question emerges
- Missing information
- Technical constraints that force a compromise

**NEVER silently compromise.** If you cannot implement exactly what was asked,
raise it as a blocker. Partial solutions require explicit user approval.

When blocked:

1. Post blocker as comment (on PR if exists, otherwise on issue)
2. Ask user for clarification in conversation
3. Present options when possible
4. Once user approves an option, continue
5. Update Q&A with the constraint and decision

### Step 6: Completion

When all planned PRs are created:

1. Verify all PRs are checked off in Plan
2. Execute the test plan to verify all changes work
3. **Review each success criterion:**
   - Confirm the implementation meets the requirement exactly
   - If any criterion is not fully satisfied, raise a blocker
   - NEVER declare completion with unmet requirements
4. Return to chat mode
5. Present final summary: what was delivered, how each success criterion was met
6. Ask user what to do next

Never mark success criteria complete — only the user does that.

NEVER close issues; closure happens automatically when PRs merge.
