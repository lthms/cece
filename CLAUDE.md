# CeCe Development

This is the CeCe plugin repository. CeCe is a modal coding assistant framework
for Claude Code.

## What's Here

```
plugins/cece/
  commands/       Command definitions (.md files)
  agents/         Agent definitions (.md files)
.claude-plugin/   Marketplace metadata
```

The core files:
- `commands/setup.md` — Onboarding wizard, includes the cece.md template
- `commands/autonomous.md` — Issue-driven autonomous mode
- `commands/wizard.md` — Creates new commands interactively
- `agents/self-quality-assurance.md` — Reviews CeCe-managed files for quality

## Key Rules

**Never edit user files directly.** The `~/.claude/` directory and
`.claude/cece.local.md` belong to the user. If rules need to change, update the
template in `setup.md` instead.

**Use `self-quality-assurance` for CeCe files only.** It's scoped to commands,
agents, and embedded templates — not CLAUDE.md or user configs.

**Test manually.** No test suite exists. User will validate the changes.

## Writing Commands and Agents

Commands and agents are markdown files with YAML frontmatter. When writing or
modifying them:

- Run `self-quality-assurance` on your changes, loop until it is satisfied
