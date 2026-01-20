---
description: Execute a task autonomously without intermediate checkpoints
---

# Do Mode

## Mode Properties

| Property | Value |
|----------|-------|
| Indicator | ⚡ |
| Arguments | `<prompt>` — complete task description with relevant context |
| Exit | Task completion, or user sends `stop` |
| Scope | Single task execution |
| Persistence | None (ephemeral) |
| Resumption | Start over |

## Permissions

**Allowed:**
- Read files, search code
- Write/edit files
- Run commands and tests
- Git operations (commit, push, branch)
- Complete the task without stopping for intermediate approvals

**Ask questions only for:**
- Choices where you cannot determine the user's preference from context
- Missing information required to proceed
- Irreversible actions with significant consequences

## Workflow

### Step 1: Announce and analyze

Announce:

<response>
⚡ Executing task.
</response>

Break the task into subtasks if needed. Use the todo list to track progress.

### Step 2: Execute

Work through each subtask:

1. **Act**: Perform the work (read, write, run commands)
2. **Verify**: Check the result before moving on
3. **Continue**: Move to the next subtask

### Step 3: Handle blockers

A blocker is anything that prevents completing the task as described.

When blocked:

1. Determine if the blocker is resolvable without user input
2. If resolvable: fix it and continue
3. If not resolvable: stop, explain the blocker, ask the user how to proceed

### Step 4: Complete

When finished:

1. Summarize what was done
2. Note any issues encountered and how they were resolved
3. Return to chat mode
