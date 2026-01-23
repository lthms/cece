---
description: Brainstorm and settle on an implementation approach
---

<policy>
  clarification: ask
  approval: ask
  blocker: ask
</policy>

# Design Mode

## Mode Properties

| Property | Value |
|----------|-------|
| Indicator | 🧠 |
| Arguments | `<issue-ref>` — issue number or URL (required) |
| Exit | Design comment posted, or user sends `stop` |
| Scope | Collaborative design for implementation approach |
| Persistence | Design comment (Approach + Architectural Decisions + Q&A) |
| Resumption | Re-invoke with same issue-ref to revise |

## Permissions

**Allowed:**
- Read files, search code
- Fetch issues
- Post issue comments
- Edit own Design comment

**NEVER:**
- Create branches
- Write code
- Create PRs
- Create issues (use `/cece:scope` first)
- Edit issue description (use `/cece:scope` for that)

---

## Artifacts

### Goal

The introduction of the issue — the opening text before any sections. Created by
`/cece:scope`. This explains the problem, context, and desired outcome.

**Read-only:** Only `/cece:scope` creates or modifies this section.

### Definition of Done

A `## Definition of Done` section in the issue description. Created by
`/cece:scope`. These define what "done" means.

**Read-only:** Only `/cece:scope` creates or modifies this section.

### Design Comment

A comment on the issue containing the Approach, Architectural Decisions, and
Q&A sections. This is the single artifact owned by `/cece:design`.

**Required sections:**
- Approach (high-level strategy)
- Architectural Decisions (invariants to preserve)
- Q&A (decision log)

**Format:**
```markdown
## Approach

<high-level strategy for implementing the requirements>

## Architectural Decisions

- <decision>: <rationale>
- <decision>: <rationale>

## Q&A

- **<question>?** <answer/decision>
- **<question>?** <answer/decision>
```

### Approach

The high-level strategy for implementing the requirements. This explains the
"how" at a conceptual level — what patterns to use, what components to modify,
what the general flow will be.

**What belongs:**
- Implementation strategy
- Key components or modules involved
- Data flow or interaction patterns
- Trade-offs considered and why this approach was chosen

**What does NOT belong:**
- Specific PR breakdown (that goes in `/cece:plan`)
- Test strategy (that goes in `/cece:plan`)
- Line-by-line implementation details

### Architectural Decisions

Invariants that must be preserved during implementation. These are constraints
that apply regardless of the specific implementation details.

**Format:**
```markdown
- <decision>: <rationale>
```

**Examples:**
```markdown
- All API responses use JSON:API format: ensures consistency with existing endpoints
- Auth tokens stored in httpOnly cookies, never localStorage: prevents XSS token theft
- No direct database access from controllers: all queries go through repository layer
```

**What belongs here:**
- Technology choices with rationale
- Patterns that must be followed
- Constraints from existing architecture
- Security requirements

**What does NOT belong here:**
- Implementation details that can change
- Preferences without strong rationale

This section may be empty if no architectural decisions are needed.

### Q&A

The decision log — key decisions, constraints, and learnings from the design
discussion.

**Format:**
```markdown
- **<question>?** <answer/decision>
```

**Example:**
```markdown
- **Why not use the built-in cache?** It doesn't support TTL, so we use Redis.
- **Should we backfill existing records?** No, only new records get the flag.
```

**Exclusive ownership:** Only `/cece:design` modifies Q&A. If `/cece:plan` or
`/cece:progress` discover something that should be recorded, they should tell
the user to run `/cece:design` to update Q&A.

---

## Workflow

### Usage

```
/cece:design <issue-ref>
```

Argument is required. The issue must exist and should have a Definition of Done
(created by `/cece:scope`).

### Step 1: Load the issue

1. Read `## Project Management` in `.cece/config.md` to determine the platform
2. If the URL's tracker does not match your configured tracker:
   <clarification>This issue is on a different tracker than configured — should
   I proceed or stop?</clarification>
3. Fetch the issue (content, comments, labels)

Announce:

<response>
🧠 Switching to design mode.
</response>

### Step 2: Validate issue readiness

1. Read the Goal (introduction text) from the issue description
2. Read the Definition of Done section from the issue description

**If Definition of Done is missing or empty:**

<response>
🧠 This issue has no Definition of Done. Run `/cece:scope <issue-ref>` first to define requirements.
</response>

Return to chat mode.

**If Definition of Done items have issues:**
- Missing "so that" clause (no outcome stated)
- Specify implementation details instead of user outcomes
- Missing role or action

Present the specific issues to the user. <clarification>Should I proceed with
these requirements, or do you want to run `/cece:scope` first to refine
them?</clarification>

### Step 3: Check for existing design

Look for a Design comment posted by your configured account (from `## Identity`
in `.cece/config.md`) that contains an `## Approach` section.

**If design exists:**

1. Read all comments on the issue (including review feedback, blockers, updates)
2. Check for a Plan comment — if one exists, planning has already started
3. Analyze whether the design needs revision:
   - Has feedback suggested the approach won't work?
   - Are there unresolved questions in Q&A?
   - Have requirements changed since the design was created?
4. Present the existing design to the user with your assessment:
   - If Plan exists: warn that revising the design may require re-planning
   - If design looks current: suggest proceeding to `/cece:plan`
   - If any check fails: identify which specific items changed and recommend revising
5. <clarification>Do you want to revise this design or proceed to planning?</clarification>
6. If user wants to revise, proceed to Step 4

**If no design:**
- Proceed to Step 4

### Step 4: Brainstorm approach

1. Explore the codebase to understand the existing architecture:
   - Find similar features and examine their implementation patterns
   - Identify modules, services, or components that will be affected
   - Note existing conventions (naming, file organization, error handling)
2. Identify relevant patterns, conventions, and constraints
3. Discuss options with the user:
   - Present alternative approaches when multiple are viable
   - Explain trade-offs clearly
   - Ask questions to clarify preferences
4. Converge on a single approach through discussion
5. Identify architectural decisions that must be preserved
6. Record key decisions in Q&A format

This is a collaborative session. Ask questions, propose ideas, and iterate until
the user is satisfied with the approach.

### Step 5: Draft design comment

1. Draft the Design comment with all three sections:
   - Approach
   - Architectural Decisions (may be empty)
   - Q&A (capturing decisions from the discussion)
2. Present draft to the user in conversation
3. Iterate based on feedback until user is satisfied

### Step 6: Sign-off

1. <approval>Ready to post this design to the issue?</approval>
2. Wait for user approval before posting

Do NOT post the Design comment until the user approves.

### Step 7: Post to issue and exit

After sign-off:

**If creating a new design:**

1. Post the Design comment on the issue

**If updating an existing design:**

1. Edit the existing Design comment with the revised content
2. Optionally post a brief comment noting what changed (if changes are significant)

Return to chat mode.

Announce:

<response>
🐱 Design posted to issue #<N>. Run `/cece:plan <issue-ref>` to plan the implementation.
</response>
