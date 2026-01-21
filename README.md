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

### Do Mode

For quick, autonomous task execution:

```
/cece:do <prompt>
```

CeCe executes the task immediately (⚡), working autonomously without
intermediate checkpoints. Use this for self-contained tasks that don't need
issue tracking.

Send `stop` to interrupt and return to chat mode.

### Plan and Progress Modes

For issue-driven work with structured planning:

```
/cece:plan <issue-ref>
/cece:progress <issue-ref>
```

**Plan mode** (📋) helps you collaboratively design an implementation approach
before writing code. CeCe explores the codebase, drafts a work plan, adds a
Definition of Done to the issue description, and posts the plan after you approve.

**Progress mode** (🔥) executes the plan independently, creating branches and
PRs per your configured git strategy. It tracks progress on the issue and
handles reviews.

This two-phase workflow separates planning from execution, giving you control
over the approach before CeCe starts implementing.

Send `stop` to interrupt either mode and return to chat mode.

> **Note:** The previous `/cece:autonomous` command is deprecated. Use
> `/cece:plan` followed by `/cece:progress` instead.

### Quick Fix

For addressing trivial PR review comments:

```
/cece:quick-fix <pr-ref>
```

CeCe classifies review comments (🪄) as trivial or non-trivial. Typos, explicit
formatting fixes, and exact rename requests are addressed immediately. Logic
changes, ambiguous suggestions, and multi-file requests are deferred for your
attention.

### Research Mode

For exploring a subject and producing a research report:

```
/cece:research <subject | path> [prompt]
```

CeCe works as a researcher (🔬), gathering information from credible sources,
verifying claims through experiments, and producing a report in `~/research/`.

Provide a subject to start new research, or a path to an existing report to
continue iterating on it. An optional prompt can guide the iteration.

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

**Configurable git strategy:** CeCe supports three push strategies configured
per-project in `.claude/cece.local.md`:
- **Fork**: Work in a fork owned by CeCe's account (recommended)
- **Remote**: Push to an existing remote the user specifies
- **Custom**: Follow user-provided instructions for complex workflows

**Protected branches:** CeCe never pushes to `main` or `master` in command
modes. Individual commands define their own commit restrictions.
