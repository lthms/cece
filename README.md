# CeCe

Think vim, but for AI. CeCe is a modal coding assistant for Claude Code.

Chat mode for collaboration. Command modes for focused, autonomous work. Each
mode has clear boundaries, explicit transitions, and predictable behavior.

CeCe explores a few ideas worth considering:

- **Modal behavior** — Separating open collaboration from focused execution
- **Agent identity** — Dedicated accounts and git authorship for transparent,
  auditable contributions
- **Fork-first workflow** — CeCe pushes to its own fork, keeping your branches
  clean until you merge

## Installation

### CLI Wrapper (Recommended)

The `cece` CLI wrapper provides the full CeCe experience. It automatically
installs the plugin, injects CeCe's system prompt, and loads your project
configuration.

```bash
go install github.com/lthms/cece/cmd/cece@latest
```

Then use `cece` instead of `claude`:

```bash
cece
```

The wrapper will auto-install the CeCe plugin on first run.

### Plugin Only

If you only want the commands without the system prompt (identity, modes, git
rules), install the plugin directly in Claude Code:

```
/plugin marketplace add lthms/cece
/plugin install cece
```

## Configuration

Create `.cece/config.md` in your project root to configure CeCe's identity and
git workflow. On first run without a config, CeCe will guide you through setup.

Example configuration:

```markdown
# Project Configuration

## Identity

Name: CeCe
Email: cece@example.com

## Git

Branch naming: cece/<issue-id>-<description>
Commit style: Imperative mood
Upstream: owner/repo

## Git Strategy

Strategy: fork
Fork account: cece-username
Fork remote name: cece

## Project Management

Issue tracker: github:owner/repo

## Provisioned Accounts

GitHub: cece-username
```

## Documentation

For concepts, workflows, and commands: **[CeCe Documentation](https://lthms.github.io/cece/)**
