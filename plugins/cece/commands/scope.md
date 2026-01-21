---
description: Create or refine an issue before planning
---

# Scope Mode

## Mode Properties

| Property | Value |
|----------|-------|
| Indicator | ✨ |
| Arguments | `[issue-ref]` — issue number or URL (optional) |
| Exit | Issue created/updated, or user sends `stop` |
| Scope | Define the problem and requirements |
| Persistence | Goal + Definition of Done in issue description |
| Resumption | Re-invoke with same issue-ref to refine |

## Permissions

**Allowed:**
- Read files, search code
- Fetch issues
- Create issues
- Edit issue descriptions
- Post issue comments

**NEVER:**
- Create branches
- Write code
- Create PRs

---

## Artifacts

### Goal

The introduction of the issue — the opening text before any sections. This
explains what the issue is about: the problem, context, and desired outcome.

**What belongs:**
- Problem statement or feature request
- Context and motivation
- High-level desired outcome

**What does NOT belong:**
- Implementation approach (that goes in `/cece:design`)
- Acceptance criteria (those go in Definition of Done)

### Definition of Done

A `## Definition of Done` section in the issue description. These are the
requirements that define "done" for this issue.

**NEVER** check off Definition of Done boxes. Only the user checks items off.

**Format:** Lite user stories — simplified statements that capture who benefits,
what they want, and why: `As a <role>, I want to <action> so that <outcome>`

The "so that" clause states the value or outcome the role receives, not the
technical mechanism.

```markdown
## Definition of Done

- [ ] As a <role>, I want to <action> so that <outcome>
```

**Examples:**
```markdown
- [ ] As a user, I want to toggle dark mode from settings so that I can reduce eye strain
- [ ] As a developer, I want auth logic in a separate module so that I can test it independently
- [ ] As an API consumer, I want a 404 response for missing users so that I can distinguish "not found" from other errors
```

**Counter-examples (avoid these):**
- `Implement dark mode` — task, not outcome; missing role and benefit
- `The API should return 404` — states what, not why; omits the role
- `Fix the Enter key bug` — describes a problem, not a desired outcome
- `As a user, I want dark mode` — omits the "so that" clause; no outcome stated

---

## Workflow

### Usage

```
/cece:scope [issue-ref]
```

- With argument: refine the referenced issue
- Without argument: help user define the task, then create an issue

### Step 1: Determine the issue

**If argument provided:**

1. Read `## Project Management` in `.claude/cece.local.md` to determine the platform
2. If the URL's tracker does not match your configured tracker:
   <clarification>This issue is on a different tracker than configured — should
   I proceed or stop?</clarification>
3. Fetch the issue (content, comments, labels)
4. Proceed to Step 2

**If no argument:**

1. <clarification>Describe the task you want to work on.</clarification>
2. Ask clarifying questions until the task is clear
3. Proceed to Step 3 (draft first, create issue after sign-off)

Announce:

<response>
✨ Switching to scope mode.
</response>

### Step 2: Analyze existing issue

1. Read the current Goal (introduction text)
2. Read Definition of Done section if it exists
3. Check the Goal for:
   - Problem statement or feature request
   - Context and motivation
   - High-level desired outcome
4. Check Definition of Done items for:
   - Each item follows the "As a <role>, I want to <action> so that <outcome>" format
   - Each item includes a "so that" clause stating user value (not mechanism)
   - No items are checked off
5. Present your assessment to the user
6. If Goal and Definition of Done are both present and well-formed:
   <clarification>This issue looks ready. Should I proceed to `/cece:design`, or
   do you want to refine it first?</clarification>
7. If refinement is needed, proceed to Step 3

### Step 3: Draft issue content

1. If needed, explore the codebase to understand context
2. Draft the Goal and Definition of Done
3. Present draft to the user in conversation
4. Iterate based on feedback until user is satisfied

### Step 4: Sign-off

1. <approval>Ready to create/update this issue?</approval>
2. Wait for user approval before posting

Do NOT create or update the issue until the user approves.

### Step 5: Post to issue and exit

After sign-off:

**If creating a new issue:**

1. Create the issue with the Goal as the body
2. Append the Definition of Done section

**If updating an existing issue:**

1. Preserve or update the Goal (opening text before first ## header)
2. Replace or create the Definition of Done section

Return to chat mode.

Announce:

<response>
🐱 Issue #<N> created/updated. Run `/cece:design <issue-ref>` to design the approach.
</response>
