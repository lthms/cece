---
description: Check and configure CeCe setup for the current project
---

# Setup

Configure your setup for this project.

## Step 1: Check git config

Run these commands and check that each returns a value:

```bash
git config cece.name
git config cece.email
git config cece.defaultMode
```

If any are missing or empty:
1. Alert the user
2. Provide exact commands to set them:
   - `git config --global cece.name "CeCe"`
   - `git config --global cece.email "<email>"`
   - `git config --global cece.defaultMode "peer"`

Proceed only after `cece.name` and `cece.email` are set.

## Step 2: Configure git identity permissions

CeCe's commit identity mechanism uses environment variables (`GIT_COMMITTER_NAME`,
etc.) that require pre-approval in Claude Code settings. These patterns are not
auto-remembered like regular commands.

**Required permissions:**

```json
[
  "Bash(GIT_AUTHOR_NAME=*)",
  "Bash(GIT_AUTHOR_EMAIL=*)",
  "Bash(GIT_COMMITTER_NAME=*)",
  "Bash(GIT_COMMITTER_EMAIL=*)"
]
```

**Procedure:**

1. Read `.claude/settings.local.json` if it exists, otherwise start with
   `{"permissions": {"allow": []}}`
2. For each required permission, check if it exists in `permissions.allow`
3. Add any missing permissions
4. Write the file back (preserve formatting)

Report which permissions were added (if any).

## Step 3: Check or create .claude/cece.local.md

Look for `.claude/cece.local.md` in the project root.

**If it does not exist:**

Create `.claude/` directory if needed, then ask each question in order:

1. "What branch naming convention?" (e.g., `cece/<slug>`, `feature/<desc>`)
2. "What commit message style?" (e.g., conventional commits, imperative mood)
3. "Where is the upstream repository?" (full URL or `owner/repo`, e.g.,
   `github.com/user/project` or `gitlab.com/org/project`)
4. "Where is the issue tracker?" (full URL or `owner/repo` on the platform, e.g.,
   `github.com/user/project`, `gitlab.com/org/project`, `linear.app/team/project`,
   or None)
5. "Any PR template or required sections?" (path or description, or None)
6. "Which CLI tools and which account for each?" (e.g., `gh: cece-bot`)

Generate `.claude/cece.local.md`:

```markdown
# Project Configuration

## Git

Branch naming: <answer>
Commit style: <answer>
Upstream: <answer>

## Project Management

Issue tracker: <answer>
PR template: <answer>

## CLI Tools & Accounts

<tool>: <account>
```

**If it exists:**

Read the file and check for:
- `## Git` section with branch naming, commit style, and upstream
- `## Project Management` section with issue tracker
- `## CLI Tools & Accounts` section

Report any missing sections or placeholder values.
Offer to fix interactively.

## Step 4: Verify CLI tool authentication

For each tool in `.claude/cece.local.md`:

| Tool | Auth check command |
|------|-------------------|
| gh | `gh auth status` |
| glab | `glab auth status` |
| linear | `linear whoami` |

Run the check and compare the authenticated account with the configured account.
Alert the user if accounts mismatch or authentication is missing.

## Step 5: Review with prompt-reviewer

After any changes to `.claude/cece.local.md`:

1. Run `prompt-reviewer` on the file
2. Apply all critical issue fixes
3. Apply non-critical fixes unless they conflict with user intent
4. Re-run reviewer until no critical issues remain

## Output

Print summary:

```
Setup complete.

Git config:
  cece.name: <value> ✓
  cece.email: <value> ✓
  cece.defaultMode: <value> ✓

Git identity permissions: ✓ | <N added>

.claude/cece.local.md: <created|updated|valid>

CLI tools:
  <tool>: <account> ✓ | ✗ (reason)

Prompt quality: ✓ | <N issues fixed>
```
