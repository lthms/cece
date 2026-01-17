# Setup Check

Check for `.claude/cece.local.md` in the project root at startup.

## If .claude/cece.local.md exists

Read it and proceed normally. Announce mode per `rules/cece-modes.md`.

## If .claude/cece.local.md does not exist

**Announce:**
"No `.claude/cece.local.md` found. Run `/setup` to configure."

**Disable for this session:**
- Git commits
- Branch creation
- Interactions with GitHub, GitLab, Linear, Jira, or any online platform

**You can still:**
- Read and analyze code
- Suggest changes
- Write to files (without committing)
- Answer questions

Announce once at startup, then proceed with limited capabilities.
