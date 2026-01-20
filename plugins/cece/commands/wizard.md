---
description: Create a new CeCe command mode through guided questions
---

# Wizard

Create a new command mode interactively.

## Mode Properties

| Property | Value |
|----------|-------|
| Indicator | 🧙 |
| Arguments | `[name] [scope...]` — name prefills question 1, remaining words prefill question 4 |
| Exit | Command file written, or user sends `stop` |
| Scope | Gather information and generate a command definition file |
| Persistence | None (ephemeral) |
| Resumption | Start over |

## Permissions

**Allowed:**
- Ask questions
- Write to `.claude/commands/` or `~/.claude/commands/`
- Run self-quality-assurance agent

**Forbidden:**
- Modify existing command files without confirmation

---

## Workflow

Announce:

<response>
🧙 Switching to wizard mode.
</response>

### Step 1: Parse arguments

If arguments provided:
- First word → `{name}` (prefills Step 2)
- Remaining words joined → `{scope}` (prefills Step 5)

Examples:
- `/cece:wizard` — no prefill
- `/cece:wizard focus` — name is "focus"
- `/cece:wizard focus Focus on a single question` — name is "focus", scope is
  "Focus on a single question"

If `{name}` is prefilled:
1. Validate format (lowercase letters and hyphens only)
2. Check if a command with that name already exists in both locations
3. If conflict found: warn the user immediately and ask whether to overwrite or
   choose a different name
4. If name is invalid: reject and ask for a different name

### Step 2: Name

Skip if prefilled from arguments.

Ask: "What should this command be called? (lowercase, hyphens allowed)"

Validate:
- Lowercase letters and hyphens only
- Check for conflicts in both `.claude/commands/` and `~/.claude/commands/`

If conflict: ask whether to overwrite or choose a different name.

Store as `{name}`.

### Step 3: Location

Use AskUserQuestion:
- Question: "Where should this command live?"
- Options:
  - Project — available only in this project (`.claude/commands/`)
  - Global — available in all projects (`~/.claude/commands/`)

Store as `{location}`.

### Step 4: Indicator

Use AskUserQuestion:
- Question: "Pick an indicator emoji for this command."
- Options:
  - 🔧 (tools/utility)
  - 🎯 (focused task)
  - 🔍 (investigation/search)
  - Other (let me type one)

If "Other": ask for the emoji in conversation.

Store as `{indicator}`.

### Step 5: Scope

Skip if prefilled from arguments.

Ask: "In one sentence, what does this command do? (from your perspective)"

Store as `{scope}`.

### Step 6: Arguments

Use AskUserQuestion:
- Question: "Does this command take arguments?"
- Options:
  - None — no arguments
  - Yes — I'll describe them

If "Yes": ask for argument description in conversation.

Store as `{arguments}`.

### Step 7: Exit conditions

Ask: "When is this command done? What are the exit conditions?"

Store as `{exit}`.

### Step 8: Persistence

Use AskUserQuestion:
- Question: "Where should this command store state for resumption?"
- Options:
  - None — ephemeral, no resumption needed
  - Issue — state lives in issue/ticket comments
  - File — state lives in a file in the repo
  - External — state lives elsewhere

If "None": set `{resumption}` to "Start over".

If "Issue", "File", or "External": ask how to resume from saved state.

Store `{persistence}` and `{resumption}`.

### Step 9: Permissions

Ask: "What actions should this command be allowed to do without confirmation?"

Store as `{allowed}`.

Ask: "What actions should be forbidden?"

Store as `{forbidden}`.

### Step 10: Workflow

Ask: "Describe what this command should do, step by step. Write freely — I'll
structure it afterward."

Take the user's free-form description and structure it into numbered steps.

Present the structured workflow:

<response>
Here's how I structured your workflow:

1. Step one description
2. Step two description
3. ...

Does this capture it?
</response>

If user accepts: proceed.

If user gives feedback: adjust the structure, present again.

Repeat until user accepts.

### Step 11: Generate

Generate the command file.

**Perspective transformation:** The user provided `{scope}` from their perspective.
- Use `{scope}` as-is for the frontmatter description (user-facing)
- Rewrite `{scope}` from CeCe's perspective for the Mode Properties scope

Example: User writes "Help me debug code" → description stays "Help me debug
code", scope becomes "Assist the user in debugging code by analyzing errors and
suggesting fixes".

Use this template:

~~~markdown
---
description: {scope}
---

# {name (title case)}

## Mode Properties

| Property | Value |
|----------|-------|
| Indicator | {indicator} |
| Arguments | {arguments} |
| Exit | {exit} |
| Scope | {scope rewritten from CeCe's perspective} |
| Persistence | {persistence} |
| Resumption | {resumption} |

## Permissions

**Allowed:**
{allowed as bullet list}

**Forbidden:**
{forbidden as bullet list}

---

## Workflow

Announce:

<response>
{indicator} Switching to {name} mode.
</response>

{workflow as numbered steps with ### headers}
~~~

### Step 12: Review

Run `self-quality-assurance` on the generated content.

Apply all fixes that do not alter the user's intended meaning.

For fixes that would change meaning, ask the user before applying.

### Step 13: Write

Create the directory if needed:
- Project: `.claude/commands/`
- Global: `~/.claude/commands/`

Write the file as `{name}.md`.

### Step 14: Confirm

Announce:

<response>
🧙 Command created: /cece:{name}

Location: {location path}
Indicator: {indicator}

You can now use this command by typing /cece:{name}
</response>

Return to chat mode.
