# Roadmap

This document captures practical directions for evolving `adr` from a file generator into a more complete ADR workflow tool.

## Priorities

The highest-value next steps are:

1. Improve day-to-day navigation and editing.
2. Add validation and lifecycle support.
3. Improve discoverability and reporting.
4. Add team and CI-oriented workflows.

## Implemented

These roadmap items are already done:

- `adr show <id>`
  Show one ADR by number, full stem, or slug.

- `adr edit <id>`
  Open one ADR in `$VISUAL`, `$EDITOR`, or a small OS-specific fallback editor list.

- `adr last`
  Show the newest ADR directly.

- `adr drop-last`
  Delete the newest ADR only when it is still in a non-final state.

## Core CLI Improvements

These commands would make the tool more useful during normal ADR work:

- `adr next`
  Print the next available ADR number for scripting.

- `adr search <query>`
  Search ADR titles and content without requiring external tools.

- `adr rename <id> <new-title>`
  Rename the ADR file while preserving the numeric prefix.

## Validation and Health Checks

These features help keep the ADR set consistent over time:

- `adr validate`
  Check for malformed filenames, duplicate numbers, numbering gaps, missing `Status:` lines, and missing required sections.

- `adr doctor`
  Report repository ADR health, such as stale proposed decisions, unknown statuses, and inconsistent formatting.

- `adr stats`
  Show counts by status, age, and recent activity.

## Lifecycle Management

These features move the tool beyond file creation and into actual decision management:

- `adr status <id> <new-status>`
  Update the ADR status line safely and consistently.

- Status transition rules
  Support explicit allowed transitions such as `Proposed -> Accepted -> Deprecated` or `Superseded`.

- Superseding support
  Add commands like `adr supersede <old> <new>` to create and maintain cross-references between decisions.

- Metadata support
  Add optional fields such as owner, tags, date accepted, related systems, and links to other ADRs.

- Aging and reminders
  Flag ADRs that have stayed in `Proposed` for too long.

## Templates and Output

These improvements would make the tool easier to adapt and easier to consume:

- Multiple templates
  Support options such as `record`, `proposal`, or `spike`.

- Custom project templates
  Allow repository-specific templates without requiring source changes.

- Generated index
  Build a `.adr/index.md` or README section summarizing all ADRs.

- Documentation site generation
  Generate a browsable ADR site from `.adr/`.

## Team and CI Use Cases

These features make the tool more useful in collaborative environments:

- CI validation
  Fail pull requests when ADR files are malformed or missing required structure.

- Pull request assistance
  Detect changes that likely require a new ADR.

- Release note integration
  Include accepted ADRs since the last release tag.

- Repository reporting
  Summarize ADR activity over time for maintainers.

## Suggested Implementation Order

If implemented incrementally, a good sequence would be:

1. `validate`
2. `status` updates and transition rules
3. `supersede` and relationship support
4. `stats`, `doctor`, and generated index output
5. CI and documentation integration
