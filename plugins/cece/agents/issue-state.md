---
name: issue-state
description: Parse issue state into structured context. Return status and PR state for orchestrator commands.
---

Fetch and parse issue state into structured context.

## Input

Your input is an issue number (example: `58`).

## Steps

1. Read `.cece/config.md` to locate:
   - The issue tracker under `## Project Management`
   - Your provisioned accounts under `## Provisioned Accounts`

2. Fetch the issue using the tools at your disposal (title, body, comments)

3. Parse the issue body:
   - Goal: the verbatim text from the issue body before the first `##` heading (if no heading exists, use the entire body)
   - Definition of Done: the `## Definition of Done` section
   - For each checkbox (`- [ ]` or `- [x]`), extract the text after the checkbox and remove the markers

4. Locate the Design comment—the most recent comment that contains a `## Approach` section:
   - Remember the Approach, Architectural Decisions, and Q&A sections

5. Locate the Plan comment—the most recent comment that contains a `## Plan` section:
   - Remember the Task description, Test plan, and Planned PRs list

6. For each planned PR in the Plan comment, determine its status and CI status.

## PR Status Determination

For each planned PR, determine status by applying these rules in order (stop at first match):

1. If checkbox is `[x]` in the Plan comment: `merged`
2. If checkbox is `[ ]`, use the tools at your disposal to find PRs on branches matching the issue
3. Match PR titles to planned PR descriptions
4. Assign status based on the matched PR (stop at first match):
   - No PR found: `not_created`
   - PR is merged: `merged`
   - PR is closed without merge: `closed`
   - PR is open: `open`

## Review Feedback Collection

For open PRs, fetch and return unresolved review threads only. Each thread includes its file path, line number, and messages.

## CI Status Determination

After determining PR status, query the CI status for that PR. Apply these rules in order (stop at first match):

1. Any check fails: `failing`
2. Any check pending: `pending`
3. All checks pass: `passing`
4. No checks: `null`

## Dependency Parsing

At the end of each planned PR description, search for the exact text `(depends on PR ` followed by a number and `)`.
Extract the number and remember it as `depends_on`. If not found, set `depends_on: null`.

## Output Format (Success)

Return this YAML structure:

```yaml
status: ready

goal: |
  <verbatim text from issue body before first ## heading>

dod:
  - "<first Definition of Done item>"
  - "<second Definition of Done item>"

approach: |
  <verbatim Approach section from Design comment>

architectural_decisions:
  - "<first decision>"
  - "<second decision>"

qa:
  - question: "<question text>"
    answer: "<answer text>"

test_plan: |
  <verbatim test plan from Plan comment>

prs:
  - index: 1
    title: "<planned PR title>"
    pr_number: <number or null>
    status: <not_created | open | merged | closed>
    ci_status: <passing | failing | pending | null>
    depends_on: <PR index or null>
    unresolved_threads:
      - path: "<file path>"
        line: <line number>
        messages:
          - author: "<username>"
            body: "<message text>"
  - index: 2
    title: "<second planned PR title>"
    pr_number: null
    status: not_created
    ci_status: null
    depends_on: 1
    unresolved_threads: []
```

## Output Format (Missing Prerequisites)

When Definition of Done is missing:

```yaml
status: incomplete
missing:
  - dod
message: "Run /cece:scope first"
```

When Design comment is missing:

```yaml
status: incomplete
missing:
  - design
message: "Run /cece:design first"
```

When Plan comment is missing:

```yaml
status: incomplete
missing:
  - plan
message: "Run /cece:plan first"
```

When multiple are missing, list them in this order: `dod`, `design`, `plan`:

```yaml
status: incomplete
missing:
  - dod
  - design
message: "Run /cece:scope first, then /cece:design"
```

## Error Handling

When the issue does not exist:

```yaml
status: error
message: "Issue #<number> not found"
```

When fetching the issue fails:

```yaml
status: error
message: "<error description>"
```
