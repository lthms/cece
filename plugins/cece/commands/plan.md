---
description: Plan work on an issue collaboratively before execution
---

# Plan Mode

## Mode Properties

| Property | Value |
|----------|-------|
| Indicator | 📋 |
| Arguments | `[issue-ref]` — issue number or URL |
| Exit | Plan posted to issue, or user sends `stop` |
| Scope | Collaborative planning for a task |
| Persistence | Plan comment + Q&A section on the issue |
| Resumption | Re-invoke with same issue-ref to revise plan |

## Permissions

**Allowed:**
- Read files, search code
- Fetch issues
- Post issue comments
- Edit issue descriptions

**NEVER:**
- Create branches
- Write code
- Create PRs

---

## Artifacts

### Plan

A comment on the issue with the `## Work Plan` heading.

**Required sections:**
- Task summary (one sentence)
- Success criteria (checkboxes)
- Approach (high-level strategy)
- Test plan (how changes will be validated)
- Planned PRs (checkboxes with scope descriptions)

**Format:**
```markdown
## Work Plan

**Task**: <summary>

**Success criteria**:
- [ ] Criterion 1
- [ ] Criterion 2

**Approach**: <strategy>

**Test plan**: <validation method, or "User approved: no tests" if waived>

**Planned PRs**:
- [ ] PR 1: <scope>
- [ ] PR 2: <scope>
```

### Q&A

A `## Q&A` section in the issue description. This is the decision log — key
decisions, constraints, and learnings from planning.

**Format:**
```markdown
## Q&A

- **<question>?** <answer/decision>
- **<question>?** <answer/decision>
```

Example:
```markdown
- **Why not use the built-in cache?** It doesn't support TTL, so we use Redis.
- **Should we backfill existing records?** No, only new records get the flag.
```

---

## Workflow

### Usage

```
/cece:plan [issue-ref]
```

- With argument: plan work for the referenced issue
- Without argument: help user describe the task, create an issue, then plan

### Step 1: Determine the issue

**If argument provided:**

1. Read `## Project Management` in `cece.local.md` to determine the platform
2. If the URL's tracker does not match your configured tracker, tell the user
   and request confirmation before proceeding
3. Fetch the issue (content, comments, labels, linked PRs)

**If no argument:**

1. Ask the user to describe the task
2. Ask questions until the task is unambiguous
3. Create a new issue capturing the agreed task
4. Proceed with that issue

Announce:

<response>
📋 Switching to plan mode.
</response>

### Step 2: Check for existing plan

Look for a Plan comment posted by your account.

**If plan exists:**

1. Read the Q&A section from the issue description
2. Read all comments on the issue (including review feedback, blockers, updates)
3. Analyze whether the plan needs revision:
   - Are there unresolved blockers or constraints discovered?
   - Has feedback suggested scope changes?
   - Are success criteria still accurate?
   - Is the test plan still valid?
4. Present the existing plan to the user with your assessment:
   - If plan looks current: suggest proceeding to `/cece:progress`
   - If plan may need updates: explain what seems outdated and suggest revising
5. Wait for user to confirm their intent before proceeding
6. If user wants to revise, proceed to Step 3

**If no plan:**
- Proceed to Step 3

### Step 3: Draft plan

1. Explore the codebase to understand the task
2. Draft the Plan (see Artifacts for required sections)
3. Present plan to the user in conversation
4. Iterate based on feedback until user is satisfied

**Test plan is mandatory.** If you cannot identify how to test the changes,
raise this during planning. The user must explicitly approve proceeding without
tests — NEVER skip this step.

### Step 4: Sign-off

1. Ask for explicit sign-off on the plan
2. Wait for user approval before posting

Do NOT post the Plan to the issue until the user approves.

### Step 5: Post to issue

After sign-off:

1. Post the Plan as a comment on the issue
2. Create or update the Q&A section in the issue description with decisions made
   during planning

Announce:

<response>
📋 Plan posted to issue. Run `/cece:progress <issue-ref>` to start execution.
</response>

### Step 6: Exit

Return to chat mode.
