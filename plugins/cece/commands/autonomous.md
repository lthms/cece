---
description: Switch to autonomous mode and work on an issue
---

# Autonomous Mode

## Mode Properties

| Property | Value |
|----------|-------|
| Indicator | 🔥 |
| Arguments | `[issue-ref]` — issue number or URL |
| Exit | Task completion, or user sends `stop` |
| Scope | Independent work on a well-defined task |
| Persistence | Post work plan as issue comment; address feedback in PRs |
| Resumption | Re-invoke with same issue-ref; plan comment provides state |

## Permissions

Once the plan is agreed upon and you are working on a branch you own, NEVER ask
for permission. Execute freely.

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

## Workflow

Treat the issue as the single source for task definition, context, progress
tracking, and decision documentation.

### Usage

```
/cece:autonomous [issue-ref]
```

- With argument: work on the referenced issue
- Without argument: ask the user to describe the task, create an issue, then proceed

### Step 1: Determine the issue

**If argument provided:**

1. Infer platform from `cece.local.md` issue tracker setting
2. If argument is a full URL that doesn't match configured tracker, tell the
   user and request confirmation before proceeding
3. Fetch the issue (content, comments, labels, linked PRs)

**If no argument:**

1. Ask the user to describe the task
2. Ask questions until the task is unambiguous
3. Create a new issue capturing the agreed task
4. Proceed with that issue

### Step 2: Check for existing plan

Look for a comment on the issue with the `## Work Plan` heading posted by your
account (as configured in `cece.local.md`).

**If plan exists:**
- Parse success criteria and PR checklist
- Identify which PRs are done, which are pending, and any blockers
- Check open PRs linked to this issue for unaddressed reviews
- If reviews exist, go to Step 6 (Handling Reviews)
- Otherwise, skip to Step 5 to continue execution

**If no plan:**
- Proceed to planning (Step 3)

### Step 3: Planning

Announce:

<response>
🔥 Switching to autonomous mode.
</response>

1. **Draft plan** including:
   - Task summary (one sentence)
   - Success criteria (checkboxes)
   - Approach (high-level strategy)
   - Planned PRs (checkboxes with scope descriptions)
3. **Present plan** to the user in the conversation
4. **Wait for explicit sign-off** before proceeding

Do NOT post the plan to the issue until the user approves.

After sign-off, update the issue description with a "Q&A" section listing all
clarifications made during planning in "Question? Answer" format.

### Step 4: Post plan to issue

After user sign-off:

1. Post the approved plan as a comment on the issue
2. Use this format:

```markdown
## Work Plan

**Task**: <summary>

**Success criteria**:
- [ ] Criterion 1
- [ ] Criterion 2

**Approach**: <strategy>

**Planned PRs**:
- [ ] PR 1: <scope>
- [ ] PR 2: <scope>
```

### Step 5: Execution

Work through each planned PR:

1. **Branch**: Create or checkout branch per naming convention in `cece.local.md`
2. **Git setup**: Read `## Git Strategy` from `cece.local.md` and prepare:
   - **fork**: Create fork if needed, add as `cece` remote, verify access
   - **remote**: Verify the specified remote is accessible
   - **custom**: Execute any setup instructions provided
3. **Implement**: Write code, commit freely
4. **Test**: Run tests after changes
5. **PR**: When a PR scope is complete:
   - Create PR linking to the issue (use "Fixes #N" or "Part of #N")
   - Assign user as reviewer (if platform supports it)
   - Update plan comment: check off completed PR, add PR link
6. **Repeat** for remaining PRs

### Step 6: Handling Reviews

When PR reviews come in, evaluate each comment before acting:

1. Does it change what "done" means for a success criterion or planned PR? Ask
   the user before implementing.
2. Does it require modifying architectural decisions documented in the plan?
   Ask the user before implementing.
3. Otherwise: implement the change or explain why not.

After addressing comments:

4. Push fixes
5. Reply to each review thread: explain what you changed or why you declined
6. If a comment added, removed, or changed a success criterion or PR scope,
   update the plan comment on the issue

### Step 7: Blockers

A blocker is anything that prevents full implementation of a requirement. This
includes:

- Tests fail unexpectedly
- Design question emerges
- Missing information
- Technical constraints that force a compromise (circular dependencies,
  incompatible APIs, platform limitations)

**NEVER silently compromise.** If you cannot implement exactly what was asked,
raise it as a blocker. Partial solutions require explicit user approval.

**Anti-pattern:** User asks for X. You hit a constraint. You implement partial-X
and continue as if it were the solution. This is wrong — raise the constraint as
a blocker first.

When blocked:

1. If working on a PR, post the blocker as a comment on the PR
2. If no PR exists yet, post the blocker on the issue
3. Ask the user for clarification in the conversation
4. Present options when possible (e.g., "Option A: ..., Option B: ..., Option C: other approach")
5. Once the user approves an option, continue with the chosen approach

### Step 8: Completion

When all planned PRs are created:

1. Verify all PRs are checked off in the plan comment
2. Return to chat mode
3. Confirm completion and ask the user what to do next

Run tests to verify code works. Never mark success criteria complete — only the
user marks them by checking them off.

NEVER close issues; closure happens automatically when PRs merge.
