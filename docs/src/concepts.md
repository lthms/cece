# Concepts

This chapter explains CeCe's modal behavior in detail.

## Modes

CeCe operates in exactly one mode at a time. There are two types: chat mode and
command modes.

### Chat Mode

Chat mode is the default. CeCe starts here and returns here when command modes
exit.

**Indicator**: 🐱

Every response in chat mode begins with 🐱, so you always know when you're in
open collaboration.

**Behavior**:
- Discuss, analyze, suggest, and implement alongside you
- Ask questions freely
- You drive; CeCe assists

In chat mode, CeCe behaves like a typical AI assistant. You ask questions, CeCe
answers. You request changes, CeCe makes them. The interaction is open-ended and
exploratory.

### Command Modes

Command modes are for focused, autonomous work. Each command has a specific
purpose: planning an issue, executing a test plan, researching a topic.

**Entering a command mode**:

```
/cece:<mode-name> [arguments]
```

For example, `/cece:plan #42` enters planning mode for issue #42.

**Indicators**: Each command mode has its own indicator:
- `/cece:plan` → 📋
- `/cece:progress` → 🔥
- `/cece:research` → 🔬

The indicator appears at the start of every response, so you always know which
mode you're in.

**Behavior**:
- CeCe drives; you observe and can interrupt
- Work is bounded by the command's scope
- Progress is saved to a defined location (issue comments, files, etc.), making
  commands idempotent and resumable
- Permissions are explicit—the command prompt defines what CeCe can and cannot do

## Transitions

### Entering Command Modes

You enter command modes explicitly with slash commands. CeCe never enters a
command mode on its own.

If you're already in a command mode and try to enter another, CeCe will ask you
to type `stop` first. Only one command mode can be active at a time.

### Exiting Command Modes

Command modes exit in two ways:

1. **Task completion**: CeCe finishes the work and returns to chat mode
2. **Interruption**: You type `stop` to halt the work

When a command mode exits, CeCe returns to chat mode (🐱).

## Interruption

Type `stop` (case-insensitive) to interrupt any command mode.

When CeCe receives `stop`:
1. Halts current work
2. Saves progress to the mode's designated storage
3. Returns to chat mode
4. Confirms what was saved and how to resume

Interruption is always available. You never lose control, even during autonomous
execution.

## Resuming Work

Most command modes support resumption. After interrupting or if context is lost,
re-invoke the command with the same arguments:

```
/cece:progress #42
```

CeCe reads the saved state (from issue comments, files, or other storage) and
continues where it left off.
