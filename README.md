# CeCe

A modal coding assistant for Claude Code.

CeCe provides a structured framework for working with Claude Code, emphasizing
transparency (agent actions are always identifiable) and modal operation (chat
mode for collaboration, command modes for focused tasks).

## Installation

Install the plugin via Claude Code:

```
/install-plugin https://github.com/lthms/cece
```

Then configure your git identity:

```bash
git config --global cece.name "CeCe"
git config --global cece.email "your-cece-email@example.com"
```

## Project Setup

In each project where you want to use CeCe:

```
/cece:setup
```

This creates `.claude/cece.local.md` with project-specific configuration
(branch naming, commit style, issue tracker, CLI accounts).

## Usage

### Chat Mode (Default)

CeCe starts in chat mode (🐱). You drive, CeCe assists. Ask questions, discuss
approaches, implement together.

### Autonomous Mode

For focused work on a specific task:

```
/cece:autonomous [issue-ref]
```

CeCe works independently (🔥) toward the goal, creating branches and PRs in its
own fork. Provide an issue reference or describe the task to create one.

Send `stop` to interrupt and return to chat mode.

### Creating Commands

To create a new command mode:

```
/cece:wizard [name] [scope...]
```

The wizard (🧙) guides you through defining a command interactively. Arguments
are optional — provide a name to prefill it, or name and scope together.

Examples:
- `/cece:wizard` — full interactive flow
- `/cece:wizard focus` — prefills name as "focus"
- `/cece:wizard focus Focus on a single question` — prefills name and scope

## Key Principles

**Transparency:** CeCe uses dedicated accounts and git identity. Actions are
always identifiable as coming from an agent.

**Fork workflow:** CeCe works in forks it owns, never pushing to repositories
owned by others.

**Protected branches:** CeCe never commits to `main` or `master`.
