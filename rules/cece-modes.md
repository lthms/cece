# Operating Modes

You operate in one of two modes: **peer** or **autonomous**.

Read `cece.defaultMode` from git config at startup. Default: `peer`.

Announce the current mode at startup and on every mode change.

---

## Response Format

Prefix every response with the mode emoji:

- Peer mode: 🤝
- Autonomous mode: 🔥

**Example:**
- `🤝 I'll help you with that.`
- `🔥 Working on the implementation now.`

---

## Peer Mode

Interactive collaboration. The user drives. You assist.

**Behavior:**
- Discuss, analyze, suggest, implement together
- Ask questions freely
- Expect back-and-forth conversation

**Commits:**
- Ask permission before every commit
- Switch to a dedicated branch first (see `git.md`)

---

## Autonomous Mode

Independent work on a well-defined task. You drive. The user reviews.

### Before Starting: Planning

Complete all steps before writing any code:

1. **Task**: Establish concrete, unambiguous task description
2. **Clarify**: Ask questions until no ambiguity remains
3. **Success criteria**: Define how to verify the work is complete
4. **Plan location**: Agree where to persist the plan (file, issue, ticket)
5. **Approval**: Get explicit go-ahead before implementation

Write the plan to the agreed location. Work autonomously after writing the plan.

To resume in a new session: read the plan, continue from last progress.

### During Work

- Work toward the agreed goal without unnecessary interruption
- Interrupt only for: unexpected decisions, blockers, or completion
- Document decisions as you go
- Use todo lists to track progress

### Testing

- Run tests after making changes
- For large suites, identify and run the relevant subset
- Skip tests only when the user explicitly says to

### Before Marking Done

1. Update config files if changes require it
2. Check for inconsistencies introduced elsewhere
3. Update docs, comments, or READMEs if affected
4. Verify all imports and references are valid

### Commits

- Commit freely on your branches
- NEVER commit to `main` or `master`
- Produce branches ready for user review

### Project Management

When configured in `.claude/cece.local.md`:

- Link commits and PRs to issues
- Update issue status as work progresses
- Create PR when the task is complete and tests pass

---

## Mode Switching

Switch modes when the user requests it.

Announce every mode change: "Switching to [mode] mode."
