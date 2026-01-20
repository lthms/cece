---
description: Address trivial PR review comments (typos, formatting, naming)
---

# Quick Fix

## Mode Properties

| Property | Value |
|----------|-------|
| Indicator | 🪄 |
| Arguments | `<pr-ref>` — PR number or URL |
| Exit | All comments classified and trivial ones addressed |
| Scope | Address only trivial review comments; defer everything else |
| Persistence | None |
| Resumption | Re-invoke with same PR reference |

## Permissions

**Allowed:**
- Read files, search code
- Write/edit files
- Create fixup commits on the current branch
- Push to the branch (per `## Git Strategy` in `cece.local.md`)
- Reply to PR review threads

**Forbidden:**
- Commit to `main` or `master`
- Push to unauthorized remotes
- Merge or close the PR
- Implement anything beyond trivial fixes

---

## What Counts as Trivial

ALWAYS defer unless the comment exactly matches the criteria below.

**Trivial (implement):**
- Typos in comments or strings
- Formatting/style issues the reviewer explicitly pointed out
- Variable/function names the reviewer explicitly named as replacements
- "Add a comment here" only if the reviewer provides exact wording

**Not trivial (defer):**
- Any logic change, no matter how small
- "This could be simplified" without exact replacement text
- Suggestions where you must judge whether the change improves behavior
- Anything touching test files
- Ambiguous or open-ended requests
- Requests spanning multiple files

Classify ambiguous comments as not trivial.

---

## Workflow

Announce:

<response>
🪄 Addressing trivial review comments.
</response>

### Step 1: Fetch PR and comments

1. Read `## Project Management` in `.claude/cece.local.md` to determine the platform
2. Use the platform CLI (`gh` for GitHub, `glab` for GitLab) to fetch the PR diff, changed files, and all review comments
3. If the PR cannot be found, announce the error and exit
4. Identify unresolved review threads

If all threads are resolved, announce and exit:

<response>
No unresolved comments.
</response>

### Step 2: Classify each thread

For each unresolved thread:

1. Read the comment and any replies
2. Determine if it meets the **trivial** criteria above
3. Record classification: `trivial` or `deferred`


### Step 3: Implement trivial fixes

For each thread classified as `trivial`:

1. Make the change in the relevant file
2. If your change differs from the reviewer's exact request, undo and defer

After all trivial fixes are applied:

1. Create fixup commits grouping related changes
2. ALWAYS use `fixup! <original commit subject>` when the parent commit is identifiable; otherwise use a single-line imperative message (e.g., "Fix typo in login comment")

If no threads were classified as trivial, skip to Step 5.

### Step 4: Push

Push the fixup commits to the PR branch.

### Step 5: Reply to threads

Reply to threads:

**For trivial (addressed):**

Use the commit SHA from the push output:

<response>
Fixed: <one-line description, e.g., "Renamed foo to bar"> (<commit-sha>).
</response>

**For deferred:**

<response>
This is not trivial. I will get back to you.
</response>

### Step 6: Report to user

Present a summary:

<response>
Addressed (N):
- <file>:<line> — <brief description of fix>
- ...

Deferred (M):
- <file>:<line> — <reason: logic change | ambiguous request | test file | multi-file span>
- ...
</response>

If any threads were deferred, remind the user they need attention.

### Step 7: Exit

Return to chat mode.
