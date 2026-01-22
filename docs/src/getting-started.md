# Getting Started

This guide walks you through installing CeCe and running your first commands.

## Installation

CeCe can be installed in two ways: with the CLI wrapper (recommended) or as a
plugin only.

### CLI Wrapper (Recommended)

The `cece` CLI wrapper provides the full CeCe experience:

- Automatically installs the CeCe plugin on first run
- Injects CeCe's system prompt (modal behavior, identity, git rules)
- Loads your project configuration from `.cece/config.md`

Install with Go:

```bash
go install github.com/lthms/cece/cmd/cece@latest
```

Then use `cece` instead of `claude`:

```bash
cece
```

### Plugin Only

If you only want CeCe's commands without the system prompt (no identity
management, git rules, or modal behavior), install the plugin directly in
Claude Code:

```
/plugin marketplace add lthms/cece
/plugin install cece
```

Note: With plugin-only installation, you get the slash commands but not the
full CeCe experience. The CLI wrapper is recommended for the intended workflow.

## Configuration

When using the CLI wrapper, CeCe reads project configuration from
`.cece/config.md` in your project root. This file defines:

- **Identity**: Name and email for commits
- **Git**: Branch naming, commit style, upstream repository
- **Git Strategy**: How CeCe pushes changes (fork, remote, or custom)
- **Project Management**: Issue tracker location
- **Provisioned Accounts**: Platform accounts CeCe can use

On first run without a config file, CeCe will guide you through setup in chat
mode. You can also create the file manually:

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

## First Commands

After setup, try these commands to see CeCe in action:

### Chat Mode

By default, CeCe operates in chat mode (indicated by 🐱). Ask questions, request
code reviews, or discuss implementation approaches. CeCe assists while you
drive.

### Working on Issues

CeCe uses a multi-phase workflow for issue-driven development:

**1. Scope** — Create or refine an issue:

```
/cece:scope #42
```

CeCe enters scope mode (indicated by ✨) and helps you define what needs to be
done, producing a clear Definition of Done.

**2. Design** — Make architectural decisions:

```
/cece:design #42
```

CeCe enters design mode (indicated by 🧠) and works through technical approach,
architectural decisions, and open questions.

**3. Plan** — Break work into PRs:

```
/cece:plan #42
```

CeCe enters planning mode (indicated by 📋) and creates a concrete plan with
test strategy and PR breakdown.

**4. Progress** — Execute the plan:

```
/cece:progress #42
```

CeCe enters progress mode (indicated by 🔥) and executes autonomously, creating
branches, writing code, running tests, and opening PRs. Type `stop` at any time
to interrupt.

## Interrupting Commands

In any command mode, type `stop` to halt CeCe's work. CeCe saves progress and
returns to chat mode, telling you what was saved and how to resume.
