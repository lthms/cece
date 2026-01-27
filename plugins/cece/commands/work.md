---
description: Execute work on an issue with an existing plan
---

# Work Mode

<mode>
indicator: 🔨
argument: <issue-ref> — issue number or URL (required)
exit: All planned PRs created, or user sends `stop`
</mode>

<variables>
- `<issue-ref>` — issue number or URL provided as argument
- `<N>` — issue number extracted from issue-ref
- `<Upstream>` — value from `.cece/config.md` `## Git` section (e.g., `lthms/cece`)
- `<upstream_remote>` — git remote whose URL contains `<Upstream>` (e.g., `origin`)
- `<default_branch>` — primary branch of upstream (e.g., `main` or `master`)
- `<branch>` — branch name for a specific PR
- `<base-ref>` — ref to rebase onto (`<upstream_remote>/<default_branch>` or another PR's branch)
</variables>

<forbidden>
- Commit to `main` or `master`
- Push to remotes not listed in `.cece/config.md`
- Close issues
- Merge PRs
- Check off Definition of Done boxes
- Edit issue description
- Edit Design comment
- Skip a phase without completing its exit condition
</forbidden>

<required>
- Use todo list to track Definition of Done items
- Sync all open PR branches before implementation
- Verify each PR fully implements its scope before creating
- Use `--author="CeCe <cece@soap.coffee>"` for all commits
</required>

---

## Artifacts

<artifact name="Definition of Done">
Location: Issue description, `## Definition of Done` section
Owner: `/cece:scope`
Access: Read-only. Track progress using todo list.
</artifact>

<artifact name="Design comment">
Location: Comment with `## Approach` heading
Owner: `/cece:design`
Access: Read-only. If implementation reveals design issues, tell user to run `/cece:design`.
</artifact>

<artifact name="Plan comment">
Location: Comment with `## Plan` heading
Owner: `/cece:plan` creates, `/cece:work` updates
Access: Update PR links when created. Check off PRs when merged.
</artifact>

---

## Phase 1: Load

<entry>Command invoked with issue reference</entry>

<steps>
1. Read `.cece/config.md`:
   - `Upstream` from `## Git`
   - `Strategy` and `Remote name` from `## Git Strategy`
   - Platform from `## Project Management`

2. Fetch issue (content, comments, linked PRs)

3. Locate required artifacts:

   | Artifact | How to find | Required |
   |----------|-------------|----------|
   | Definition of Done | `## Definition of Done` in issue body | Yes |
   | Design comment | Comment containing `## Approach` | Yes |
   | Plan comment | Comment containing `## Plan` | Yes |

4. If artifact missing, announce and exit:

   <missing artifact="Definition of Done">
   <response>
   🔨 No Definition of Done. Run `/cece:scope <issue-ref>` first.
   </response>
   Return to chat mode.
   </missing>

   <missing artifact="Design comment">
   <response>
   🔨 No design. Run `/cece:design <issue-ref>` first.
   </response>
   Return to chat mode.
   </missing>

   <missing artifact="Plan comment">
   <response>
   🔨 No plan. Run `/cece:plan <issue-ref>` first.
   </response>
   Return to chat mode.
   </missing>
</steps>

<exit>
All artifacts loaded.

<response>
🔨 Working on #<N>.
</response>
</exit>

---

## Phase 2: Sync

<entry>Phase 1 complete</entry>

<steps>
1. Determine `<upstream_remote>` and `<default_branch>`:

   ```bash
   # Find remote whose URL contains <Upstream>
   git remote -v | grep <Upstream>
   # Result: the remote name is <upstream_remote>

   # Get default branch
   git remote show <upstream_remote> | grep 'HEAD branch'
   # Result: the branch name is <default_branch>
   ```

2. Fetch upstream before any branch or merge status checks:

   ```bash
   git fetch <upstream_remote>
   ```

3. Check merge status: for each PR in Plan comment that is not already checked
   off, check its status on the hosting platform. For any that show "merged",
   edit the Plan comment to tick its checkbox.

4. For each PR in Plan comment that exists but is not merged:

   a. Checkout the branch:
      ```bash
      git checkout <branch>
      ```

   b. Determine base ref:

      | Condition | Base ref |
      |-----------|----------|
      | PR has no dependency | `<upstream_remote>/<default_branch>` |
      | PR depends on merged PR | `<upstream_remote>/<default_branch>` |
      | PR depends on open PR | That PR's branch |

   c. Check if behind (upstream already fetched in step 2):
      ```bash
      git merge-base --is-ancestor <base-ref> HEAD
      # Exit 0 = up to date, Exit 1 = behind
      ```

   d. If behind, rebase:
      ```bash
      git rebase <base-ref>
      # If conflicts: resolve, git add, git rebase --continue
      # If unresolvable: git rebase --abort, raise blocker
      git push --force-with-lease
      ```

5. If rebase conflicts cannot be resolved:

   <blocker>
   Rebase conflict in `<branch>`. Files: [list conflicts]. How should I resolve?
   </blocker>

   Wait for user input before continuing.
</steps>

<exit>
All existing PR branches up to date.

<response>
🔨 Branches synced.
</response>
</exit>

---

## Phase 3: Implement

<entry>Phase 2 complete</entry>

<steps>
1. Create todo list from Definition of Done:
   - One item per requirement
   - Mark complete only after code committed AND tests pass

2. For each planned PR (respecting dependency order):

   <step name="pre-check">
   Before starting work, check PR status on the hosting platform:

   - If this PR shows "merged", skip to the next PR.
   - If a predecessor PR shows "merged", rebase onto
     `<upstream_remote>/<default_branch>` instead of the predecessor's branch.
   </step>

   <step name="branch">
   | Condition | Command |
   |-----------|---------|
   | Branch exists | `git checkout <branch>` |
   | Branch new | `git checkout -b <branch> <upstream_remote>/<default_branch>` |
   </step>

   <step name="implement">
   Write code for this PR's scope (from Plan comment). Commit incrementally:
   ```bash
   git commit --author="CeCe <cece@soap.coffee>" -m "<message>"
   ```
   </step>

   <step name="test">
   Execute test plan from Plan comment:

   | Test plan says | Action |
   |----------------|--------|
   | Specific commands | Run them. Fix failures before proceeding. |
   | Manual steps | Describe what to verify. |
   | "User approved: no tests" | Skip testing. |
   </step>

   <step name="verify">
   Before creating PR, confirm:
   - All commits address this PR's scope only
   - Tests pass (or waived)
   - Every Definition of Done item claimed by this PR is fully implemented

   If any fails: fix before proceeding.
   </step>

   <step name="create-pr">
   Create PR:
   - Target: `<default_branch>`
   - Body: "Fixes #<N>" if this is the final PR, otherwise "Part of #<N>"
   - Reviewer: assign user

   Update Plan comment: add PR link. Do not check off (checked when merged).
   </step>

   <step name="rebase-dependents">
   If Plan marks other PRs as `(depends on PR X)` where X is this PR's number:

   For each dependent:
   ```bash
   git checkout <dependent-branch>
   git rebase <this-branch>
   # Resolve conflicts or abort and raise blocker
   git push --force-with-lease
   git checkout <this-branch>
   ```
   </step>

3. Repeat for next planned PR.
</steps>

<exit>
All planned PRs created and pushed. Proceed to Phase 4.
</exit>

---

## Phase 4: Complete

<entry>Phase 3 complete</entry>

<steps>
1. Verify completion:
   - All PRs from Plan are created and linked
   - All Definition of Done items implemented (check todo list)
   - All tests pass

2. If verification fails:
   - Identify gap
   - Return to Phase 3 to address
   - Re-verify

3. Summarize:
   - What was delivered
   - Which Definition of Done items each PR addresses
</steps>

<exit>
<response>
🐱 All work complete for #<N>.
</response>

Return to chat mode.
</exit>

---

## Blockers

A blocker prevents completing a phase.

| Type | Example | Action |
|------|---------|--------|
| Missing info | Unclear requirement | Ask user |
| Technical | Cannot implement as designed | Present options, wait |
| Conflict | Rebase conflict | Show files, ask resolution |
| Test failure | Tests fail after attempts | Show failure, ask |

When blocked:
1. Announce what blocks progress
2. Present options if applicable
3. Wait for user
4. Resume current phase after resolution

---

## Reviews

When user reports PR feedback:

1. Evaluate each comment:

   | Type | Action |
   |------|--------|
   | Clarification | Respond in thread |
   | Bug/style fix | Implement, push, respond |
   | Scope expansion | Ask: should this be in scope? |
   | Conflicts with Design | Raise blocker |

2. After addressing:
   - Push fixes
   - Rebase dependents (if any)
   - Respond in each thread

3. If review changes what a PR delivers, tell user to run `/cece:plan` to update the plan
