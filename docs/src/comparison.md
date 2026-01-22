# Comparison

CeCe isn't the only AI coding assistant, and it's not trying to be the best at
everything. This chapter explains the ideas CeCe explores and how they differ
from other approaches.

## What CeCe Explores

CeCe focuses on three ideas that may be worth considering:

### Modal Behavior

Most AI coding assistants operate in a single mode. You chat, ask questions, and
request changes in an open-ended conversation. This works well for exploration
but can be chaotic for complex, multi-step work.

CeCe separates **open collaboration** (chat mode) from **focused execution**
(command modes). When you enter a command mode, CeCe takes the lead on a bounded
task with explicit permissions and clear exit conditions. You can interrupt at
any time with `stop`.

This isn't better or worse—it's a different tradeoff. Some developers prefer
staying in control at all times. Others appreciate handing off well-defined work
while retaining the ability to interrupt.

### Agent Identity

When an AI assistant creates commits, opens PRs, or posts comments, whose
identity appears? Usually yours, because the assistant uses your credentials.

CeCe maintains its own identity:
- Its own git author (name and email)
- Its own GitHub/GitLab account
- Clear indication when work comes from an agent

This makes agent-authored work transparent. When you see a commit from CeCe, you
know it wasn't written by a human. When reviewers see a PR from CeCe's account,
they know they're reviewing agent work.

This transparency has tradeoffs. Some teams prefer attribution to go to the
developer who directed the work. CeCe's approach prioritizes visibility over
convenience.

### Fork-First Workflow

Many AI assistants push directly to your branches. This is convenient but can
clutter your branch history with agent commits that later need rebasing or
squashing.

CeCe pushes to its own fork by default. Your branches stay clean until you
explicitly merge CeCe's work. You review the PR, and only then do the changes
enter your branch.

This adds a step—you must merge the PR—but keeps your local work separate from
agent-generated changes. If you prefer direct pushes, CeCe supports that too.

## What CeCe Doesn't Do

CeCe isn't optimized for:

- **Minimal setup**: CeCe requires a configuration file (`.cece/config.md`)
  before full functionality. The CLI wrapper guides you through setup on first
  run, but assistants that work out of the box have lower friction.
- **Invisible assistance**: CeCe's identity is always visible. If you want
  agent work to appear as your own, CeCe isn't the right fit.

## Should You Give It a Try?

CeCe is an experiment. The ideas it explores—modal behavior, agent identity,
fork-first workflow—aren't new, but their combination creates a particular
workflow that some developers may find useful.

If you value clear boundaries between collaboration and execution, transparent
attribution, and clean branch history, CeCe's approach might work for you. If
you prefer seamless, invisible assistance, other tools are better suited.

Try it and see.
