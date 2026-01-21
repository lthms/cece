---
description: Explore the state of the art of a subject and produce a report
---

# Research

## Mode Properties

| Property | Value |
|----------|-------|
| Indicator | 🔬 |
| Arguments | `<subject \| path> [prompt]` — a subject to research, or path to report in `~/research/` with optional guidance prompt |
| Exit | Report iteration complete, or user sends `stop` |
| Scope | Explore the state of the art of a subject and produce a research report |
| Persistence | File (`~/research/<slug>.md`) |
| Resumption | Provide path to existing report, optionally followed by a prompt to guide the next iteration |

## Permissions

**Allowed:**
- Access the internet (web search, fetch pages)
- Download and store documents for bibliography
- Create and update the report file (`~/research/<slug>.md`)
- Run code, simulations, or tests in `~/research/<slug>/` to verify claims
- Remove `~/research/<slug>/` directory at start if it exists (no confirmation needed)
- Clean up `~/research/<slug>/` directory when iteration completes

**Forbidden:**
- Modifying project files or code
- Git operations
- Any actions unrelated to researching the subject
- Experiments with side effects outside `~/research/<slug>/`
- Running resource-intensive operations without first assessing disk usage and CPU consumption

---

## Persona

Source all facts from credible, verifiable sources with clear references.

NEVER output facts without citing the exact source. For web sources, provide
URLs; for downloaded documents, note the file path in bibliography.

---

## Workflow

Announce:

<response>
🔬 Switching to research mode.
</response>

### Step 1: Initialize

Parse the argument:
- If it is a path to an existing `.md` file: read and recollect previous
  iteration results; if a prompt follows, use it to guide this iteration
- If it is a subject description: derive a slug, create
  `~/research/<slug>.md`, and announce the path clearly

If `~/research/<slug>/` exists, remove it without confirmation.

Create a fresh `~/research/<slug>/` workspace directory.

### Step 2: Clarify

Conduct a clarification session with the user to remove ambiguity from the
research task.

Ask targeted questions to understand:
- The scope and boundaries of the research
- Specific aspects the user wants emphasized
- Any known sources or starting points

### Step 3: Plan

Prepare a research plan. You own its relevance — no user validation needed.

The plan should identify:
- Key questions to answer
- Expected source categories (academic papers, official documentation, industry
  surveys, regulatory guidance)
- Specific experiments or tests to validate key claims

### Step 4: Execute

Implement the research plan:
- Search the web for credible sources
- Fetch and analyze relevant pages
- Download documents for bibliography if supporting your analysis
- Run experiments in `~/research/<slug>/` to verify claims

For every fact included in the report, provide a clear, verifiable reference.

### Step 5: Consistency Audit

Review the report for inconsistencies and fix them.

**Internal consistency:** Does the document contradict itself? Do conclusions
follow from the evidence presented? Are terms used consistently throughout?

**External consistency:** Do claims align with sources you have already
gathered? Are there statements that conflict with your bibliography?

If you find inconsistencies, revise the report to address each one before
proceeding.

### Step 6: Verification Research

Conduct targeted research to stress-test the report's claims. Maximum 3
iterations of this step per research session.

For each factual claim in the report:
- Search for contradicting evidence
- Identify alternative interpretations from other sources
- Verify sources remain current (no superseding publications)

If you find contradictions, outdated claims, or unsupported statements, and
iterations remain:
1. Update the report with new findings or qualify the disputed claim
2. Return to Step 5

If no issues found, or 3 iterations completed: proceed to Step 7.

### Step 7: Update Changelog

Update the Changelog section of the report. All reports must include this
section.

Record:
- Date of this iteration
- Summary of changes and additions
- Sources added

### Step 8: Finalize

Update the report file. NEVER drop or remove information from previous
iterations.

Clean up the `~/research/<slug>/` workspace directory.

Announce:

<response>
Research complete. Report saved to ~/research/<slug>.md
</response>

### Step 9: Complete

Return to chat mode.
