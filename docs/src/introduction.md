# Introduction

Think vim, but for AI. CeCe is a modal coding assistant built as a plugin for
Claude Code.

## The Core Idea

CeCe has a **chat mode** for open-ended collaboration and **command modes** for
focused, autonomous work. This isn't new—Cline has Plan/Act modes, Roo Code has
five role-based modes, Cursor has layout modes. But CeCe's modes are different:
they're workflow phases, not task types.

In chat mode, you're in control. CeCe discusses, suggests, and implements
alongside you. In command modes, CeCe takes the lead on a specific task—scoping
an issue, designing an approach, planning PRs, or executing work—while you
retain the ability to interrupt at any time.

## Why Modes?

Structured workflows help with complex tasks. CeCe's modal design
is about enacting these workflows with simple slash commands. This brings:

- **Clear boundaries**: You always know which mode you're in. Chat mode shows
  a 🐱 indicator; command modes show their own indicators.
- **Explicit transitions**: You enter command modes deliberately with slash
  commands like `/cece:plan`. You exit by completing the task or typing `stop`.
- **Predictable behavior**: Each command mode has defined permissions and
  constraints. You know what CeCe can and cannot do.

## What CeCe Explores

CeCe isn't trying to be the most capable coding assistant. Instead, it explores
a few ideas that may be worth considering:

- **Modal behavior**: Separating open collaboration from focused execution
- **Agent identity**: CeCe has its own git identity and accounts, making
  agent-authored work transparent
- **Fork-first workflow**: CeCe pushes to its own fork, keeping your branches
  clean until you merge

These ideas aren't unique to CeCe, but the combination creates a particular
workflow worth trying.
