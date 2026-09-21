# Read-Only Scan Specification

## Purpose

Define the first bounded OsdyCleaner capability: a local, non-privileged, read-only scan of exactly five built-in macOS developer/cache areas, represented by one versioned snapshot and presented consistently as text, JSON, or a live read-only terminal dashboard.

## Requirements

### Requirement: Fixed Built-In Scan Coverage

The scan SHALL resolve exactly the following built-in areas, in the listed canonical order, relative to the current user's local home directory:

| Area ID | Display name | Built-in root |
| --- | --- | --- |
| `npm-cache` | npm cache | `~/.npm` |
| `homebrew-cache` | Homebrew cache | `~/Library/Caches/Homebrew` |
| `gradle-caches` | Gradle caches | `~/.gradle/caches` |
| `xcode-derived-data` | Xcode DerivedData | `~/Library/Developer/Xcode/DerivedData` |
| `core-simulator` | CoreSimulator | `~/Library/Developer/CoreSimulator` |

The scan MUST create one root observation for every listed area and MUST NOT discover, configure, or accept any additional root. User-supplied filesystem paths and environment- or tool-configured cache overrides MUST NOT alter this root set.

#### Scenario: Resolve the complete built-in set

- GIVEN a supported local macOS environment with a resolvable current-user home
- WHEN a scan is started without a path
- THEN the snapshot contains one root observation for each of the five listed area IDs in canonical order
- AND no other root is scanned

#### Scenario: Reject an arbitrary path

- GIVEN a user supplies a positional path or path-selection option to `scan`
- WHEN input is validated
- THEN the request is rejected as invalid input before traversal begins
- AND the supplied path is not inspected

### Requirement: Strict Read-Only and Local Scope

The scan MUST only observe local filesystem metadata needed to produce the snapshot and MUST NOT intentionally create, modify, move, rename, delete, or change permissions on any filesystem object. It MUST NOT provide or initiate cleanup, deletion, Trash, selection, planning, confirmation, or execution behavior.

The capability MUST NOT use a network service, invoke a shell or external executable, request privilege escalation, emit telemetry, run as a daemon, or scan an external volume. It MUST NOT invoke npm, Homebrew, Gradle, Xcode, or CoreSimulator tools.

#### Scenario: Complete a scan without side effects

- GIVEN fixture roots with recorded contents and metadata
- WHEN the scan and each of its three presentations complete
- THEN the fixture contents and metadata remain unmodified by OsdyCleaner
- AND no network, shell, external-tool, privilege, telemetry, daemon, cleanup, deletion, or Trash operation occurs

### Requirement: Explicit Root Observation States

Each built-in root observation MUST have exactly one status from `scanned`, `missing`, `inaccessible`, `skipped`, `boundary_limited`, `partial`, or `cancelled`, plus a stable reason code where the status alone is insufficient.

The statuses SHALL mean:

- `scanned`: traversal completed without a visibility gap for every entry permitted by this specification.
- `missing`: the built-in root did not exist.
- `inaccessible`: the root itself existed but could not be opened or inspected.
- `skipped`: the root was intentionally not traversed because a safety boundary applied or cancellation occurred before it started.
- `boundary_limited`: traversal completed except for one or more entries on another mount or device.
- `partial`: traversal produced trustworthy facts but one or more in-scope entries could not be inspected for a reason other than a mount/device boundary.
- `cancelled`: traversal of that root stopped because the scan was cancelled.

When multiple conditions occur within a traversed root, `cancelled` MUST take precedence over `partial`, `partial` MUST take precedence over `boundary_limited`, and all observed conditions MUST remain available as warnings. A missing root is a complete observation of absence, not an error.

#### Scenario: Every built-in root is missing

- GIVEN none of the five built-in roots exists
- WHEN the scanner initializes and finalizes successfully
- THEN the snapshot contains exactly five `missing` root observations
- AND the snapshot is complete with zero measured totals and no findings
- AND absence is not reported as a global failure

#### Scenario: One root is inaccessible

- GIVEN one built-in root exists but is inaccessible and the other four roots are observable
- WHEN the scan runs
- THEN that root is recorded as `inaccessible` with useful warning context
- AND the remaining four roots are still scanned
- AND a trustworthy partial snapshot is finalized

### Requirement: Bounded Traversal

Traversal MUST remain beneath each built-in root, MUST NOT follow a root or entry that is a symbolic link, and MUST NOT cross from the permitted starting filesystem to another mount or device. A built-in root that is a symbolic link or is on an external volume MUST be `skipped` with a stable reason rather than traversed.

A skipped symbolic link MUST contribute neither its target's facts nor its target's size. A mount/device boundary MUST be recorded as a boundary warning and reflected in the root status. Isolated entry inspection failures MUST be recorded and MUST NOT prevent inspection of safe sibling entries or other roots.

#### Scenario: Encounter a symbolic-link entry

- GIVEN an in-scope directory contains a symbolic link to another location
- WHEN the directory is scanned
- THEN the symbolic link target is never traversed or measured
- AND the skipped link is represented by warning context associated with its root
- AND safe sibling entries continue to be inspected

#### Scenario: Encounter a mount boundary

- GIVEN an entry beneath a built-in root is reported on a different device from the permitted starting filesystem
- WHEN traversal reaches that entry
- THEN the entry and all of its descendants are skipped
- AND the root observation is `boundary_limited` unless a higher-precedence state applies
- AND no facts from the other device contribute to findings or totals

### Requirement: Explainable Findings and Honest Estimates

Every finding MUST identify its built-in area, display path, match reason, `manual_review` risk, estimate values and basis, and applicable warnings. No finding or presentation SHALL describe an item as safe to delete, cleanup-eligible, recommended for removal, or guaranteed to yield reclaimable space.

Logical size and allocated size, when available, MUST be distinguished. Unknown or unavailable allocated size MUST remain explicitly unknown rather than being replaced by logical size. Every presentation MUST explain that APFS clones, snapshots, compression, hard links, and filesystem accounting can make allocated estimates differ from space that could be reclaimed. CoreSimulator findings MUST additionally avoid inferring device lifecycle or product-specific cleanup semantics.

#### Scenario: Present an APFS-affected estimate

- GIVEN a finding has logical size and allocated-size evidence
- WHEN the finding is presented in text, JSON, or the viewer
- THEN both values are labeled by their measurement basis
- AND the presentation describes the value as an estimate requiring manual review
- AND it makes no promise of deletion safety or reclaimable space

### Requirement: Repeated Filesystem Identity Accounting

Within one snapshot, the same observed filesystem identity MUST contribute to aggregate logical and allocated totals at most once, including when multiple hard-link paths expose that identity. Canonical attribution MUST be selected after deterministic ordering so traversal or concurrency timing cannot choose which occurrence is counted.

Additional observed paths for the same identity MAY remain visible, but each MUST be marked as already accounted for and MUST add zero incremental bytes to every aggregate. If identity or allocation information cannot be observed, the snapshot MUST preserve that uncertainty in a warning and MUST NOT claim complete deduplication or exact reclaimable bytes. This accounting rule MUST NOT be presented as duplicate-file analysis.

#### Scenario: Observe two hard links to one identity

- GIVEN two fixture paths report the same filesystem device and object identity
- WHEN the snapshot is aggregated
- THEN the identity's logical and allocated values contribute exactly once
- AND the deterministically first occurrence receives canonical attribution
- AND any later occurrence is marked as already accounted for without being called a duplicate cleanup opportunity

### Requirement: Versioned Shared Snapshot Contract

Every successfully finalized scan MUST produce one immutable-by-convention snapshot with an explicit schema version. The snapshot MUST contain, at minimum, scan outcome and completeness, built-in rule-set version, all five root observations, findings, aggregate estimates, warnings, and any run metadata retained by the product.

The top-level outcome MUST be `complete`, `partial`, or `cancelled`. It MUST be `complete` when all roots are `scanned` or `missing`, `partial` when any non-cancellation visibility gap exists, and `cancelled` when cancellation determines the result. Estimates in a partial or cancelled snapshot MUST be labeled incomplete.

Text, JSON, and the terminal dashboard MUST consume this same snapshot contract as their sole source of scan facts. They MUST NOT rescan, maintain a presentation-specific result domain, or independently recompute findings, statuses, warnings, or totals. A breaking change to snapshot fields or semantics MUST use a new schema version.

#### Scenario: Render one snapshot in every presentation

- GIVEN one finalized snapshot with a fixed schema version
- WHEN it is rendered as text, JSON, and in the terminal dashboard
- THEN every presentation identifies the same schema version and scan outcome
- AND corresponding roots, findings, risks, estimates, completeness, and warnings agree
- AND changing presentation does not trigger another scan

### Requirement: Deterministic Snapshot and Serialization Order

For identical fixture contents, rule-set version, options, and fixed run metadata, snapshot collections and serialized output MUST be repeatable regardless of filesystem enumeration order or concurrent completion order.

Roots MUST use the canonical area order in this specification. Findings MUST be ordered by area order, then normalized display path, then a stable identity-independent tie breaker. Warnings MUST be ordered by associated area, path, and stable warning code. Runtime-varying metadata, if retained, MUST be explicit and controllable or fixed for deterministic verification.

JSON MUST be one valid document, MUST use a stable schema-declared field order and canonical collection order, and MUST produce identical bytes for identical fixed snapshots. It MUST NOT rely on unordered object or map traversal for serialized order.

#### Scenario: Stable JSON under reordered discovery

- GIVEN two runs over identical fixtures and fixed run metadata
- AND filesystem enumeration and worker completion arrive in different orders
- WHEN each snapshot is serialized as JSON
- THEN the two JSON byte sequences are identical
- AND their root, finding, warning, and accounting order is the specified canonical order

### Requirement: Equivalent Text and JSON Reports

Text and JSON reports MUST expose the snapshot version, outcome/completeness, all five root statuses, findings, estimate bases, manual-review risks, and warnings. Formatting MAY differ, but neither report may contradict, omit the existence of, or invent a snapshot fact.

JSON output MUST preserve typed values and explicit unknown values rather than parsing human-formatted text. Text output MUST use non-actionable language and MUST make partial or cancelled status prominent.

#### Scenario: Compare human and machine-readable reports

- GIVEN a partial snapshot containing one inaccessible root and one finding
- WHEN text and JSON reports are produced
- THEN both report the same root status, finding facts, estimate basis, manual-review risk, warning, and partial outcome
- AND only their presentation formatting differs

### Requirement: Live Read-Only Terminal Dashboard

When terminal mode is selected, Bubble Tea MUST open before scan work begins. Scan I/O MUST run only in Bubble Tea commands/effects; `Update` MUST be state-only and `View` MUST be pure. Text and JSON modes MUST remain noninteractive and unchanged.

Live progress MUST be limited to the five canonical categories in their specified serial order and activity for the active category. It MUST NOT fabricate per-file percentages, a pre-scan pass, or other scan facts. On cancellation, the terminal program MUST wait for scanner shutdown before finalizing or restoring the terminal.

The terminal program MUST render one finalized schema-v1 snapshot as the sole authoritative source for outcome, estimates, roots, findings, and warnings. Progress is transient UI state and MUST NOT alter snapshot facts. It MUST display complete, partial, and cancelled snapshots without changing their facts.

The finalized read-only dashboard MUST expose estimates, outcome, findings, warnings, all canonical categories in order, and selected finding or warning detail. Its responsive cyber-neon wide and compact presentations MUST preserve the same textual meaning without color. It MUST use clear phrases equivalent to “Read-only scan,” “Manual review required,” and “Estimate—not guaranteed reclaimable space.” Controls MUST provide `h`/`l` or arrows, `j`/`k` or arrows, `Page Up`/`Page Down`, `Tab`, and `q`. Controls and content MUST NOT use cleanup, selection, action, confirmation, deletion, removal, Trash, or reclaim commands or affordances. Quitting MUST NOT start, authorize, or imply another operation.

#### Scenario: Open terminal mode before scanning

- GIVEN `scan` is invoked with terminal input and output
- WHEN terminal mode starts
- THEN Bubble Tea opens before scan I/O begins
- AND scan I/O is performed by commands/effects rather than `Update` or `View`
- AND text and JSON mode behavior remains noninteractive and unchanged

#### Scenario: Show honest live progress

- GIVEN terminal mode is scanning the fixed built-in roots
- WHEN progress is rendered
- THEN it identifies only completed canonical categories and active-category activity in canonical serial order
- AND it shows no per-file percentage, pre-scan pass, or fabricated scan fact

#### Scenario: Cancel an active terminal scan

- GIVEN terminal mode has an active scanner
- WHEN cancellation is requested
- THEN no new scanner work is begun
- AND the program waits for scanner shutdown before finalizing or restoring the terminal
- AND any finalized result remains a cancelled, incomplete snapshot

#### Scenario: Review a finalized dashboard

- GIVEN one finalized schema-v1 snapshot with findings and warnings
- WHEN the dashboard renders in wide or compact layout
- THEN both layouts preserve the same textual outcome, estimates, canonical category order, findings, warnings, and selected detail without relying on color
- AND the dashboard consumes that snapshot as its sole result source
- AND controls provide category selection, selected-detail scrolling, findings/warnings switching, and quit only
- AND no selection or action wording or affordance is present

#### Scenario: View a cancelled snapshot

- GIVEN a valid cancelled snapshot with incomplete observations
- WHEN the dashboard opens
- THEN cancellation and incompleteness are prominent
- AND available partial facts and warnings match the shared snapshot
- AND quitting performs no further operation

### Requirement: Isolated Errors and Global Failures

Missing roots, inaccessible roots, inaccessible entries, symbolic links, mount boundaries, and other root- or entry-scoped inspection errors MUST be represented in a finalized snapshot and MUST NOT become global failures while trustworthy aggregation remains possible.

A global failure MUST be reserved for inability to initialize the supported scan contract or inability to finalize or serialize a trustworthy snapshot. An unsupported operating environment MUST be distinguished from a runtime failure. When no trustworthy snapshot can be finalized, no partial document may be represented as valid snapshot output.

#### Scenario: Continue after an entry error

- GIVEN one entry cannot be inspected while safe sibling entries and other roots remain observable
- WHEN scanning continues
- THEN the affected root is `partial`
- AND a scoped warning identifies the visibility gap
- AND safe sibling facts and all other root observations remain in the finalized partial snapshot

#### Scenario: Fail global initialization

- GIVEN the current-user home or valid built-in rule set cannot be resolved
- WHEN scan initialization is attempted
- THEN traversal does not begin
- AND no valid snapshot is emitted
- AND the failure is reported as a global diagnostic rather than a root observation

### Requirement: Cancellation Produces an Incomplete Snapshot

The scanner MUST honor cancellation promptly, stop beginning new inspection work, and make no claim that unvisited data was scanned. If trustworthy collected facts can be finalized, it MUST produce a snapshot with outcome `cancelled`, incomplete estimates, completed root observations preserved, an active root marked `cancelled`, and not-yet-started roots marked `skipped` with a cancellation reason.

Cancellation MUST NOT discard already finalized warnings or findings and MUST NOT start a viewer action, cleanup, retry, or any other operation. A cancelled snapshot MUST never be represented as complete.

#### Scenario: Cancel during traversal

- GIVEN at least one root has completed, one root is active, and later roots have not started
- WHEN cancellation is received
- THEN new inspection work stops
- AND a trustworthy finalized result has outcome `cancelled` and incomplete estimates
- AND completed, active, and unstarted roots are respectively preserved, `cancelled`, and `skipped` with cancellation context
- AND no subsequent operation begins

### Requirement: Stable Streams and Scan Exit Classes

For non-interactive text and JSON modes, the finalized snapshot representation MUST be written to stdout. Snapshot warnings MUST remain in that representation; transient progress and diagnostics MUST be written only to stderr. JSON stdout MUST contain exactly one JSON document and no progress, decoration, or diagnostic text.

If a global initialization or finalization failure prevents a trustworthy snapshot, stdout MUST NOT contain a purported snapshot and stderr MUST contain the diagnostic. The terminal dashboard MAY render on the attached interactive terminal, but non-render diagnostics MUST remain separate.

The `scan` exit classes SHALL be stable:

| Code | Meaning |
| --- | --- |
| `0` | A complete snapshot was finalized, including a valid all-missing result. |
| `1` | A global runtime, finalization, serialization, or output failure prevented successful reporting. |
| `2` | Command input was invalid, including any arbitrary path input or unsupported format value. |
| `3` | A non-cancelled partial snapshot was finalized because visibility was incomplete. |
| `4` | The operating environment is unsupported. |
| `130` | Cancellation determined the scan outcome. |

A root-scoped warning MUST NOT produce code `1`. Cancellation MUST take precedence over partial-result code `3` when a valid cancelled snapshot is reported.

#### Scenario: Keep JSON stdout parseable during a partial scan

- GIVEN one root is inaccessible and a valid partial snapshot can be finalized
- WHEN JSON output is requested
- THEN stdout contains exactly one parseable JSON snapshot
- AND warning facts are contained in that document
- AND diagnostics are confined to stderr
- AND the process exits with code `3`

#### Scenario: Report cancellation through streams and exit status

- GIVEN cancellation occurs after trustworthy facts have been collected
- WHEN JSON or text reporting finalizes the cancelled snapshot
- THEN stdout contains that valid incomplete snapshot representation
- AND stderr contains only separate diagnostics, if any
- AND the process exits with code `130`

### Requirement: Disposable and Deterministic Verification

All automated filesystem tests for this capability MUST use controlled fixtures and temporary directories and MUST never inspect or mutate the developer's or test runner's real home directory. Tests MUST cover normal and empty roots, every-root-missing, an inaccessible root and entry, symbolic links, simulated or controlled mount/device boundaries, repeated identities/hard links, cancellation, deterministic text and JSON, shared snapshot equivalence, and read-only dashboard language.

Fixture tests MUST NOT require network access, external developer tools, privilege escalation, real external volumes, or mutation of personal files. Any run metadata used in output verification MUST be fixed.

#### Scenario: Run the acceptance suite in isolation

- GIVEN a test runner with a real home directory containing arbitrary personal data
- WHEN the read-only scan acceptance suite runs
- THEN all scanned roots resolve within disposable fixtures or temporary directories controlled by the tests
- AND the real home directory is neither inspected nor mutated
- AND deterministic output assertions use fixed metadata
