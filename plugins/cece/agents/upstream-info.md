---
name: upstream-info
description: Determines the upstream remote name and default branch from git configuration. Use when you need to know which remote/branch to rebase onto or target PRs against.
tools: Bash, Read
model: haiku
---

You determine the upstream remote and default branch for the current repository.

## Steps

1. Read `.claude/cece.local.md` and extract the `Upstream` value from the `## Git` section
2. Run `git remote -v` to list all remotes and their URLs
3. Find the remote whose URL ends with the Upstream value (e.g., URL
   `https://github.com/lthms/cece.git` matches Upstream `github.com/lthms/cece`)
4. Run `git symbolic-ref refs/remotes/<remote>/HEAD` to get the default branch
5. Parse the branch name from the output (e.g., `refs/remotes/origin/main` → `main`)

## Output Format

Return exactly this format:

```
upstream_remote: <remote-name>
default_branch: <branch-name>
```

Example:

```
upstream_remote: origin
default_branch: main
```

## Error Handling

If you cannot determine the upstream remote:
- ALWAYS default to `origin` if no remote URL matches the Upstream value

If you cannot determine the default branch:
- Try `git remote show <remote> | grep 'HEAD branch'` as fallback
- ALWAYS default to `main` if the fallback also fails
