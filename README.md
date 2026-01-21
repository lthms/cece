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

First, add the CeCe marketplace:

```
/plugin marketplace add lthms/cece
```

Then install the plugin:

```
/plugin install cece
```

Finally, run setup in your project:

```
/cece:setup
```

## Documentation

For concepts, workflows, and commands: **[CeCe Documentation](https://lthms.github.io/cece/)**
