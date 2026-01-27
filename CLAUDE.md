# CeCe Development

This is the CeCe plugin repository. CeCe is a modal coding assistant framework
for Claude Code.

## What's Here

```
cmd/              CLI wrapper to inject CeCe system prompt
plugins/cece/
  commands/       Command definitions (.md files)
  agents/         Agent definitions (.md files)
.claude-plugin/   Marketplace metadata
```

The core files:
- `cmd/cece/system_prompt.md` — CeCe system prompt, injected by the `cece` CLI wrapper
- `commands/plan.md` — Collaborative planning for issues
- `commands/progress.md` — Execute work on an issue with an existing plan
- `commands/research.md` — Research mode for exploring subjects
- `commands/wizard.md` — Creates new commands interactively
- `agents/self-quality-assurance.md` — Reviews CeCe-managed files for quality

**Deprecated:**
- `commands/autonomous.md` — Deprecated in favor of `plan.md` + `progress.md`.
  Do not reference or modify these files when making changes to the codebase.

## Key Rules

**Never edit user files directly.** The `~/.claude/` directory and
`.cece/config.md` belong to the user.

**Use `self-quality-assurance` for CeCe files only.** It's scoped to commands
and agents — not CLAUDE.md, system prompts, or user configs.

**Test manually.** No test suite exists. User will validate the changes.

## Writing Commands and Agents

Commands and agents are markdown files with YAML frontmatter. When writing or
modifying them:

- Prefer XML tags over Markdown headings for structure
- Run `self-quality-assurance` on your changes, loop until it is satisfied

## Writing Skills

Skills live in `plugins/cece/skills/<name>/SKILL.md`. When writing or
modifying them:

- Use authoritative language in the `description` frontmatter field (e.g.,
  "Required skill to...") to increase the odds Claude loads the skill
  automatically
- Prefer XML tags over Markdown headings for structure
- Run `self-quality-assurance` on your changes, loop until it is satisfied
