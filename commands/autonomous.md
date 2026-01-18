---
description: Switch to autonomous mode and work on an issue
---

# Autonomous

Use the issue as your source of truth for context, plan, progress, and
decisions.

## Usage

```
/autonomous [issue-ref]
```

- With argument: work on the referenced issue
- Without argument: clarify the task with the user, create issue, then proceed

## Step 1: Determine the issue

**If argument provided:**

1. Infer platform from `cece.local.md` issue tracker setting
2. If argument is a full URL that doesn't match configured tracker, tell the
   user and request confirmation before proceeding
3. Fetch the issue (content, comments, labels, linked PRs)

**If no argument:**

1. Ask the user to describe the task
2. Clarify until the task is unambiguous
3. Create a new issue capturing the agreed task
4. Proceed with that issue

## Step 2: Check for existing plan

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

## Step 3: Planning

Announce: "Switching to autonomous mode."

1. **Clarify** the task with the user until no ambiguity remains
2. **Draft plan** including:
   - Task summary (one sentence)
   - Success criteria (checkboxes)
   - Approach (high-level strategy)
   - Planned PRs (checkboxes with scope descriptions)
3. **Present plan** to user locally
4. **Wait for explicit sign-off** before proceeding

Do NOT post the plan to the issue until the user approves.

After sign-off, update the issue description to add a Q&A section summarizing
key clarifications.

## Step 4: Post plan to issue

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

## Step 5: Execution

Work through each planned PR:

1. **Branch**: Create or checkout branch per naming convention in `cece.local.md`
2. **Fork**: Fork the repository to your configured account and add it as a
   remote (see `rules/cece-git.md` - always work in your fork)
3. **Implement**: Write code, commit freely
4. **Test**: Run tests after changes
5. **PR**: When a PR scope is complete:
   - Create PR linking to the issue (use "Fixes #N" or "Part of #N")
   - Assign user as reviewer (if platform supports it)
   - Update plan comment: check off completed PR, add PR link
   - Post a progress comment capturing decisions and tradeoffs
6. **Repeat** for remaining PRs

## Step 6: Handling Reviews

When PR reviews come in:

1. Address review comments
2. If reviews change the design, update the issue description to reflect changes
3. Update plan comment if scope or approach changes
4. Push fixes and respond to reviewers

## Step 7: Blockers

If you encounter a blocker (tests fail unexpectedly, design question emerges,
missing information):

1. Post a blocker comment on the issue describing the situation
2. Stop and ask the user locally for clarification
3. Once resolved, update the issue and continue

## Step 8: Completion

When all planned PRs are created:

1. Post summary comment on the issue
2. Ask the user what to do next

NEVER check success criteria yourself; the user verifies completion.

NEVER close issues; closure happens automatically when PRs merge.
