# Getting Started

This guide walks you through installing CeCe and running your first commands.

## Installation

First, add the CeCe marketplace:

```
/plugin marketplace add lthms/cece
```

Then install the plugin:

```
/plugin install cece
```

## Setup

Before using CeCe's full capabilities, run the setup wizard:

```
/cece:setup
```

The wizard does two things:

1. **Initializes CeCe's core rules** in your global `~/.claude/` directory
   (creates or updates `CLAUDE.md` and `rules/cece.md`)
2. **Creates per-project configuration** in `.claude/cece.local.md` with:
   - **Git identity**: CeCe's name and email for commits
   - **Git strategy**: How CeCe pushes changes (fork, remote, or custom)
   - **Project management**: Which issue tracker you use

Run setup in each project where you want to use CeCe. The local configuration
file is yours to edit if your settings change.

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
