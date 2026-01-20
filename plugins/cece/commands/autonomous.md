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
| Persistence | Plan comment + Q&A section on the issue |
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

Once the plan is agreed upon and you are working on a branch you own, execute
freely. NEVER ask for permission.

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

Autonomous mode maintains three artifacts on the issue. These persist across
sessions and provide state for resumption.

### Plan

A comment on the issue with the `## Work Plan` heading. Tracks what needs to be
done and current progress.

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

**Update when:**
- PR is completed (check it off, add link)
- Success criteria change
- Scope changes

### Q&A

A `## Q&A` section in the issue description. This is the decision log — key
decisions, constraints, and learnings that affect the task.

**Purpose:** Preserve context across sessions. When resuming, read this first and
refer to prior decisions before proposing changes to settled approaches.

**What belongs:**
- Decisions made during planning ("Chose X over Y because Z")
- Constraints discovered during implementation ("API doesn't support batch ops")
- Clarifications from user that affect approach
- Learnings from PR reviews that apply to remaining work

**What does NOT belong:**
- Verbose discussion logs
- Obvious or trivial decisions
- Temporary blockers that were resolved

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

**Update when:**
- After planning sign-off (initial decisions)
- After blocker resolution (constraints discovered)
- After significant review feedback (learnings)

### Comments

Posted on the issue or PRs during execution to communicate blockers, progress,
and decisions.

**Where to post:**
- Blockers → on the PR if one exists, otherwise on the issue
- Progress updates → on the issue (optional, for long tasks)
- Review responses → on the PR threads

---

## Workflow

Treat the issue as the single source for task definition, context, progress
tracking, and decision documentation.

### Usage

```
/cece:autonomous [issue-ref]
```

- With argument: work on the referenced issue
- Without argument: ask the user to describe the task, create an issue, then
  proceed

### Step 1: Determine the issue

**If argument provided:**

1. Read `## Project Management` in `cece.local.md` to determine the platform
2. If argument is a full URL that doesn't match configured tracker, tell the
   user and request confirmation before proceeding
3. Fetch the issue (content, comments, labels, linked PRs)

**If no argument:**

1. Ask the user to describe the task
2. Ask questions until the task is unambiguous
3. Create a new issue capturing the agreed task
4. Proceed with that issue

Announce:

<response>
🔥 Switching to autonomous mode.
</response>

### Step 2: Check for existing plan

Look for a Plan comment (see Artifacts) posted by your account.

**If plan exists:**

1. Read the Q&A section from the issue description
2. Validate plan completeness against required sections; if incomplete, ask user
   whether to update the plan before proceeding
3. Parse current state: which PRs are done, pending, any blockers
4. Check open PRs for unaddressed reviews
5. Present summary to user: what's planned, done, remaining, pending reviews
6. If reviews are pending, proceed to Step 5. Otherwise, proceed to Step 4.

**If no plan:**
- Proceed to Step 3 (Planning)

### Step 3: Planning

1. Draft the Plan (see Artifacts for required sections)
2. Present plan to the user in conversation
3. Wait for explicit sign-off before proceeding

**Test plan is mandatory.** If you cannot identify how to test the changes,
raise this during planning. The user must explicitly approve proceeding without
tests — NEVER assume it's okay.

Do NOT post the Plan to the issue until the user approves.

After sign-off:
- Post the Plan as a comment on the issue
- Create the Q&A section in the issue description with initial decisions

### Step 4: Execution

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

### Step 5: Handling Reviews

When PR reviews come in, evaluate each comment:

1. Does it change what "done" means? → Ask user before implementing
2. Does it require modifying architectural decisions? → Ask user before implementing
3. Otherwise → Implement the change

NEVER decline review feedback without user approval. If you believe a comment
should not be addressed, explain your reasoning to the user and let them decide.

After addressing comments:

4. Push fixes
5. Reply to each thread: explain what changed or why declined
6. Update Plan if scope or criteria changed
7. Update Q&A if review revealed significant learnings

### Step 6: Blockers

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

### Step 7: Completion

When all planned PRs are created:

1. Verify all PRs are checked off in Plan
2. Execute the test plan to verify all changes work
3. **Review each success criterion:**
   - Verify the implementation actually satisfies it
   - If any criterion is not fully met, raise a blocker
   - NEVER declare completion with unmet requirements
4. Return to chat mode
5. Present final summary: what was delivered, how each success criterion was met
6. Ask user what to do next

Never mark success criteria complete — only the user does that.

NEVER close issues; closure happens automatically when PRs merge.
