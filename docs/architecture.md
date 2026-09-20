# OsdyCleaner: safe, reviewable macOS cleanup architecture

**Decision:** build a terminal-first, personal-use macOS cleaner in **Go 1.24+**: Bubble Tea v2, Bubbles v2, Lip Gloss v2, and Cobra. Phase 1 only discovers and reports. Every later cleanup action comes from an immutable scan snapshot, is bound to filesystem identity, and is revalidated immediately before mutation.

## Quick path

1. Run `osdy scan --category developer-cache` to create a read-only, versioned snapshot.
2. Review candidates and risks in the Bubble Tea TUI or JSON output.
3. Create a plan from that snapshot and, in Phase 2, execute it in Trash mode after confirmation.
4. Treat reported savings as an **estimate** until Trash is emptied.

> Go cannot encode as many invalid-state constraints at compile time as Rust. Runtime guard discipline—validated construction, identity binding, and mutation-time revalidation—is mandatory; language memory safety alone cannot prevent wrong-file deletion.

## Product boundary

| Topic | Decision |
| --- | --- |
| User and scope | One person's Mac; local filesystem only |
| Interface | Full-screen TUI plus scriptable CLI subcommands |
| First release | Read-only analysis; Trash-backed cleanup follows in Phase 2 |
| Destructive escape hatch | Permanent deletion, gated by stronger typed confirmation |
| Safety model | Rules, protected roots, identity-bound plans, and pre-mutation revalidation |

Goals are explainable findings, developer artifacts/caches plus personal-file review, responsive terminal use, recovery through Trash, and bounded non-privileged scans. Duplicate analysis is later.

Non-goals: antivirus, `sudo` or helpers, background/automatic cleaning, cloud/accounts/telemetry, shared-machine administration, arbitrary-path deletion, arbitrary shell execution, and external-volume traversal in MVP.

## Technology decision

| Option | Strengths | Decision |
| --- | --- | --- |
| **Go + Bubble Tea v2 + Bubbles v2 + Lip Gloss v2 + Cobra** | Productive terminal stack, straightforward distribution, explicit runtime safety boundaries | **Recommend** |
| Rust + Ratatui | Strongest alternative for stricter compile-time invariants | Alternative, not selected |
| Swift | Native Foundation access | Reject: terminal-first TUI choices are less established for this product |

Use one Go module and a deliberately small dependency set:

- **Bubble Tea v2, Bubbles v2, Lip Gloss v2:** TUI state, components, rendering, and terminal lifecycle.
- **Cobra:** shared command grammar for interactive and automation paths. Do not add Viper unless configuration later proves necessary.
- **`encoding/json`:** snapshots, plans, audits, and JSON output with explicit schema versions.
- **Standard library + existing `x/sys`:** a narrow Darwin descriptor-relative walker for targeted discovery; no generic framework.
- **A narrow `TrashBackend` port:** Phase 2 macOS support may call Foundation `FileManager.trashItem` through a small vetted adapter/bridge. Do not commit to an unvetted dependency or AppleScript string interpolation.

## Architecture and boundaries

```text
osdycleaner/                         # one Go module
├── cmd/osdy/                         # main package
├── internal/core/                    # domain, guards, audit, invariants
├── internal/rules/                   # built-in versioned rule data
├── internal/scan/                    # traversal and candidate collection
├── internal/plan/                    # snapshot-bound planning
├── internal/platform/macos/          # identity, Trash, vetted tool ports
├── internal/cli/                     # Cobra commands and JSON/text I/O
└── internal/tui/                     # Bubble Tea models and views
```

`internal/core` never mutates the filesystem or accepts command text. `internal/platform/macos` implements narrow interfaces such as `CandidateInspector`, `TrashBackend`, and `ToolAdapter`. Cobra and the TUI own presentation only; traversal and deletion run off the TUI update/render loop through bounded command/effect boundaries.

## Domain model and rule shape

Use named ID/value types, typed enums, unexported constructors, validation at creation, immutable-by-convention snapshots/plans, and exhaustive tests.

| Type | Required meaning |
| --- | --- |
| `Rule` | Versioned built-in discovery/action policy. |
| `Candidate` | Display path, filesystem identity, measurements, match explanation, risk, and scan metadata. |
| `ScanSnapshot` | Immutable-by-convention, versioned completed/cancelled scan record. |
| `CleanupPlan` | Immutable-by-convention selection from exactly one snapshot; candidate IDs and action intent, never loose paths. |
| `PlannedAction` | Candidate-bound backend request with expected identity/state and guards. |
| `OperationResult` | Idempotent observed result: moved, deleted, skipped, failed, or already completed. |
| `AuditRecord` | Append-only intent, confirmation, environment, result, warning, and timestamp report. |

A candidate identity is not its display string: store presentation path plus device, inode/file ID, object type, and required scan-time metadata. Plans bind `(snapshotID, candidateID, identity, ruleID, ruleVersion)` and cannot be created from a command-line path. Rescanning creates a new snapshot; it never refreshes a plan silently.

```go
type Rule struct {
    id               RuleID
    version          RuleVersion
    category         Category
    roots            []RootSpec
    matcher          Matcher // typed predicates, never shell text
    exclusions       []Exclusion
    risk             RiskLevel
    recoverability   Recoverability
    defaultSelection SelectionPolicy
    actionBackend    ActionBackend // Trash, Permanent, or ToolAdapter
    guards           []Guard
    sizeEstimator    SizeEstimator
    supportNotes     SupportNotes
}
```

Only constructors in the owning package create valid rules, snapshots, candidates, and plans; callers receive copies or read-only accessors. Rule matchers are typed data, not command strings. Docker/npm adapters are allowlisted IDs with fixed arguments from typed inputs, no shell, advisory parsed output, and audit records.

## Categories and flow

| Category | MVP timing | Action strategy | Default |
| --- | ---: | --- | --- |
| Developer artifacts | Scan first | Trash selected known regenerable paths | Eligible low-risk caches may be preselected; review required |
| General caches/logs | Scan first | Trash with rule exclusions | Manual unless narrow/regenerable |
| Personal large/old files | Scan first | Report and manual Trash selection | **Manual** |
| Docker/CoreSimulator | Later adapters | Tool-aware semantics | **Manual** |
| App leftovers/duplicates | Later or experimental | Report/read-only until evidence proves safety | **Manual** |

Personal files (for example, files ≥1 GiB or untouched ≥180 days) are always manual: age and size are review cues, never deletion proof. Known regenerable caches are materially different.

```text
Idle → ResolvingRules → Scanning → Finalizing → SnapshotReady
                         │ cancel                 │ select
                         ▼                        ▼
                  CancelledSnapshot         DraftPlan → ValidatingPlan
                                                    │ valid
                                                    ▼
                                             AwaitingConfirmation → Executing → ResultsRecorded
```

Resolve enabled built-in rules and permitted roots; emit candidates, warnings, and progress incrementally; retain a cancelled snapshot only when marked incomplete. Only completed snapshots plan by default. Validate before confirmation, then revalidate each action immediately before mutation. Persist each result as known; isolated safe failures produce partial results, while cancellation or invariant failure stops execution.

## Safety invariants

- **Protected roots:** reject system roots, the home directory as a deletion target, exclusions, and anything outside the candidate's allowed root.
- **No mount crossing:** stay on the starting filesystem by default; external volumes are out of MVP.
- **No symlink following:** discovery and mutation never treat a symlink as its target.
- **Identity over display path:** plans bind expected identity and type; display paths are for people.
- **Traversal-resistant mutation:** where applicable, use Go 1.24+ `os.Root`/`os.OpenInRoot` APIs relative to an opened allowed root, then revalidate identity, type, containment, and guards immediately before mutation. `os.Root` is not a complete sandbox: mount boundaries still need explicit checks.
- **Permanent deletion:** never call `os.RemoveAll` on an arbitrary absolute path from a plan. Resolve and operate relative to the opened allowed root only after the above revalidation.
- **Hard links and APFS:** deduplicate aggregate accounting by device/inode. Logical and allocated sizes are estimates; clones, snapshots, compression, and hard links can make APFS reclaimed bytes lower or different.
- **No implicit privilege or arbitrary tools:** no `sudo`, passwords, helpers, escalation, shell, or user/rule-provided executable text.
- **Snapshot-bound, idempotent execution:** reject stale, altered, incomplete (unless explicitly supported later), or foreign-snapshot plans; report `already_moved`, `missing`, `identity_changed`, and similar observed states honestly.

## Trash and permanent deletion

### Default: Trash

The default backend moves selected items to macOS Trash. UI and CLI say **“moved to Trash”**, never “freed”: space is not reclaimed until the user empties Trash. Results distinguish `moved_to_trash`, `not_found`, `revalidation_failed`, and backend failure.

### Advanced: permanent deletion

Permanent deletion remains disabled until the Trash flow has test evidence. When enabled, it requires an exact count/estimate phrase such as `DELETE 3 ITEMS / 2.4 GiB`; rejects generic non-interactive `--yes` for personal/high-risk items; revalidates per item; records actual outcomes rather than claimed reclaimed bytes; and uses the same audit schema for failures and skips.

## TUI and CLI

The TUI maintains category navigation, candidate review, risk/details, selection, and operation status; confirmation is modal. It must show progress, discoveries, warnings, cancellation state, and size estimates. Keys remain `↑/↓` or `j/k`, `Tab`, `Space`, `a`, `/`, `Enter`, `r`, `c`, `Esc`, and `q`, with quit confirmation for an active operation.

```text
osdy
osdy scan [--category ID] [--format text|json] [--output SNAPSHOT]
osdy rules [--format text|json]
osdy doctor [--format text|json]
osdy plan create --snapshot SNAPSHOT --select CANDIDATE_ID...
osdy plan show --plan PLAN [--format text|json]
osdy clean --plan PLAN --backend trash
osdy clean --plan PLAN --backend permanent --confirm-phrase "..."
```

`scan` is read-only and serializes a schema version. `plan create` accepts IDs from one snapshot, and `clean` accepts a plan—neither accepts paths. JSON is one structured stdout document; progress/diagnostics use stderr. Exit classes: `0` complete, `1` runtime failure, `2` invalid input/confirmation, `3` partial operation, `4` unsupported environment, `130` cancelled.

## Performance and testing

Start from targeted roots, not the whole disk. On Darwin, depth-bounded descriptor DFS enumerates opened directories and opens each child relative to its parent; it never reopens an accepted directory by absolute pathname, follows no symlink target, and crosses no mount even during concurrent rename. Bound descriptors, workers, and queues/backpressure; keep scan/deletion work off the TUI loop. If a child disappears or its identity changes before safe open, record observational `entry_changed`. Show logical and allocated estimates when available and label their basis.

`filepath.WalkDir` does not follow static symlinks, but its callback validation can race a later pathname `ReadDir`; it is not the production walker. It may be a test/reference utility only when it cannot imply production safety.

| Layer | Coverage |
| --- | --- |
| Domain | Rule matching, constructors/validation, guards, selection, immutable snapshots/plans, result transitions |
| Safety | Root containment, identity-bound plans, stale-plan rejection, no arbitrary-path execution |
| Integration | Disposable fixtures for descriptor traversal, atomic rename/symlink swap, close counts, FD budget, mount abstraction, hard links, cancellation, and partial failures; never personal paths |
| Backends | Fakes that assert revalidation order and idempotent results |
| TUI | Direct Bubble Tea model `Init`/`Update`/`View` state tests, golden output, and fake command/effect boundaries |
| macOS | Disposable-fixture Trash tests only; never personal files |

Test permanent deletion only in disposable fixtures and after equivalent Trash coverage. Assert “moved,” not “freed.”

## Roadmap and Phase 1 acceptance

1. **Phase 1 — read-only:** one module, domain model, built-in rules, targeted scans, JSON/text snapshots, audit reports, and acceptance tests. No cleanup backend.
2. **Phase 2 — recovery first:** plans, review UI, narrow macOS Trash adapter, root-relative revalidation, audits, and disposable-fixture tests.
3. **Phase 3 — interfaces/adapters:** TUI polish, stable CLI, `doctor`, constrained Docker/npm adapters, and CoreSimulator research.
4. **Phase 4 — advanced analysis:** cautious app leftovers, experimental read-only duplicates, and evidence-based permanent deletion.
5. **Phase 5 — distribution:** arm64 packaging, docs, signing/notarization decisions, and release support policy.

Phase 1 is accepted when:

- [ ] `osdy scan` mutates nothing and requires no privilege, network, or shell.
- [ ] Built-in rules expose ID, version, category, roots, risk, support notes, and estimate basis.
- [ ] Text and JSON output are immutable-by-convention, explicitly schema-versioned snapshots.
- [ ] Candidates show display path, reason, risk, estimates, and warnings.
- [ ] Depth-bounded descriptor traversal never reopens accepted directories by absolute pathname, follows no symlink target or mount under concurrent rename, bounds FDs/workers, and records `entry_changed` when a child disappears or changes identity before safe open.
- [ ] Aggregate accounting deduplicates hard links and labels APFS reclaim estimates accurately.
- [ ] Cancellation creates an explicitly incomplete/cancelled result and never starts cleanup.
- [ ] Domain, safety, temporary-directory integration, and Bubble Tea state/golden tests cover these invariants.

Recommended defaults remain arm64 first, macOS 14+, no external volumes in MVP, report-file audit retention, a 1 GiB large-file cue, a 180-day old-file cue, deferred duplicates, and permanent deletion disabled until Trash evidence exists.

## Sources

- Bubble Tea v2 repository and release: <https://github.com/charmbracelet/bubbletea>, <https://github.com/charmbracelet/bubbletea/releases/tag/v2.0.0>
- Lip Gloss v2 release: <https://github.com/charmbracelet/lipgloss/releases/tag/v2.0.0>
- Cobra: <https://github.com/spf13/cobra>
- Go traversal-resistant APIs: <https://go.dev/blog/osroot>, <https://pkg.go.dev/os>
- `filepath.WalkDir` behavior: <https://pkg.go.dev/path/filepath>
- Apple Trash API: <https://developer.apple.com/documentation/foundation/filemanager/trashitem(at:resultingitemurl:)>
- APFS storage behavior: <https://developer.apple.com/documentation/foundation/about-apple-file-system>
- `stat(2)` links and allocation metadata: <https://developer.apple.com/library/archive/documentation/System/Conceptual/ManPages_iPhoneOS/man2/stat.2.html>
- Docker cleanup semantics: <https://docs.docker.com/reference/cli/docker/system/prune/>
- npm cache semantics: <https://docs.npmjs.com/cli/v11/commands/npm-cache/>
