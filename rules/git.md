# Git Rules

## Configuration

- `cece.name`: Your display name (required)
- `cece.email`: Your email for authorship (required)

---

## Hard Constraints

**NEVER:**
- Commit to `main` or `master`
- Reset or stash changes you did not create
- Add user's `Signed-Off-By` without explicit approval
- Discard uncommitted changes without asking

**ALWAYS:**
- Commit with your identity as author, user as committer
- Follow branch naming from `.claude/cece.local.md`
- Alert the user when uncommitted changes exist that you did not make

---

## Commit Identity

Every commit:

| Field     | Value                        |
|-----------|------------------------------|
| Author    | `cece.name <cece.email>`     |
| Committer | `user.name <user.email>`     |

```bash
GIT_COMMITTER_NAME="$(git config user.name)" \
GIT_COMMITTER_EMAIL="$(git config user.email)" \
git commit --author="$(git config cece.name) <$(git config cece.email)>" \
  -m "commit message"
```

---

## Branches

**Protected:** `main`, `master`
- Never commit to these branches

**Your branches:**
- Named per `.claude/cece.local.md` convention
- Commit freely in autonomous mode
- In peer mode, ask permission before committing

---

## Commit History

Write commits that each represent one logical change.

**Commit messages:**
- Explain what changed and why
- NEVER use generic messages like "Address PR review feedback"
- Write each message so it explains the change without requiring context from
  other commits

**Handling PR reviews:**
- Use fixup commits to address review feedback
- Before requesting another review, squash fixups into the commits they fix
- Rewriting your own branch history between review rounds is allowed

---

## Destructive Operations

If uncommitted changes exist that you did not make:

1. Stop
2. Alert the user
3. Ask how to proceed
4. Never discard automatically

---

## Signed-Off-By

This trailer is a human attestation.

**User's Signed-Off-By:**
- Requires explicit approval every time
- When amending a commit, remove the user's Signed-Off-By first
- User must re-approve after modifications

**Your Signed-Off-By:**
- Add freely

```
Signed-Off-By: user <user@example.com>   # Requires approval
Signed-Off-By: CeCe <cece@example.com>   # Add freely
```
