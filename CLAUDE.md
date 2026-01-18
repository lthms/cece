# CeCe Development

Instructions for developing CeCe itself.

## Project Structure

```
.claude/                    Project-local configuration (example)
.claude-plugin/             Marketplace metadata
plugins/cece/               Main plugin
  .claude-plugin/           Plugin manifest
  agents/                   Agent definitions
  commands/                 Command definitions
```

## Architecture

CeCe is a Claude Code plugin that provides a modal coding assistant framework.

**Core components:**

- `plugins/cece/commands/setup.md` — Setup wizard for new projects
- `plugins/cece/commands/autonomous.md` — Autonomous mode for issue-driven work
- `plugins/cece/commands/wizard.md` — Command creation wizard
- `plugins/cece/agents/prompt-reviewer.md` — Quality reviewer for instruction files

**User-facing files created by setup:**

- `~/.claude/CLAUDE.md` — References `@rules/cece.md`
- `~/.claude/rules/cece.md` — Core principles (transparency, modal behavior, git)
- `.claude/cece.local.md` — Per-project configuration

## Key Concepts

**Modal behavior:**
- Chat mode (🐱) is the default state
- Commands enter via `/cece:<name>` and define their own indicator
- Commands exit on completion or when user sends `stop`

**Transparency:**
- Agent uses dedicated accounts, never user's personal accounts
- Commits use configured `cece.name` and `cece.email`
- Work happens in forks owned by the agent's account

**Git identity:**
```bash
git config cece.name "CeCe"
git config cece.email "cece@example.com"
```

## Development Workflow

1. Create a branch following the `cece/<slug>` convention
2. Make changes to plugin files
3. Test by running `/cece:setup` in a test project
4. Test commands and agents manually

## Testing Changes

No automated tests. Verify changes by:

1. Running `/cece:setup` to check setup flow
2. Running `/cece:autonomous` to check autonomous mode
3. Running `/cece:wizard` to check command creation
4. Using the `prompt-reviewer` agent on modified instruction files

## Contribution Guidelines

**Instruction files:**
- Use imperative mood ("Run tests" not "Tests should be run")
- Use "you" when addressing the agent
- Avoid vague terms ("appropriate", "as needed")
- Put hard constraints first with NEVER/ALWAYS

**Commits:**
- Imperative mood in commit messages
- One logical change per commit
- Explain what and why

## Files Reference

| File | Purpose |
|------|---------|
| `plugins/cece/commands/setup.md` | Interactive project setup |
| `plugins/cece/commands/autonomous.md` | Issue-driven autonomous work |
| `plugins/cece/commands/wizard.md` | Command creation wizard |
| `plugins/cece/agents/prompt-reviewer.md` | Instruction quality review |
| `.claude-plugin/marketplace.json` | Plugin distribution metadata |
| `plugins/cece/.claude-plugin/plugin.json` | Plugin manifest |
