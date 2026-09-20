# Technical Design: Read-Only Scan Foundation

## Decision summary

Implement one Go 1.24+ module at the repository root. `osdy scan` resolves exactly five home-relative roots and scans them serially through a descriptor-relative platform walker. Production traversal never uses `filepath.WalkDir`: after a trusted root is opened, directories are enumerated from their bound descriptors and descendants are opened relative to the parent descriptor with no-follow checks. Bounded regular-file workers may reopen only validated relative component chains from the trusted root; they never inspect an absolute descendant pathname. The scanner finalizes one immutable-by-convention snapshot and then presents it as text, stable JSON, or a minimal Bubble Tea viewer.

This strong traversal invariant replaces the rejected WU5A `serialScanner` candidate. The TUI coherence gate remains **passed**: the viewer receives a finalized snapshot, has only category/detail navigation and quit behavior, and introduces no selection or follow-up-operation domain.

## Scope and invariants

This design is limited to read-only local metadata observation. It adds no file mutation, arbitrary path, shell or external command, privilege, network, telemetry, background service, cleanup plan, or product-specific cache interpretation.

The implementation must preserve these invariants:

1. Exactly five built-in roots are resolved in canonical order; every finalized snapshot has five root observations.
2. Root and entry symbolic links are never followed or measured, and entries on another device or mount are never traversed or counted.
3. Once a directory descriptor is accepted, enumeration and child acquisition remain descriptor-relative; no absolute descendant pathname is reopened after validation.
4. Root-scoped visibility failures become snapshot facts; only initialization, invariant, finalization, serialization, or output failures are global.
5. Filesystem identity accounting happens after canonical ordering, so scheduling cannot choose attribution.
6. Text, JSON, and TUI consume one finalized snapshot and cannot scan, retry, or recompute statuses, findings, warnings, or totals.
7. Findings always use `manual_review` risk and estimate language, never safety or reclaim guarantees.

## Module and package topology

The implementation creates the following focused topology. Files are grouped by behavior rather than by individual type.

```text
go.mod
go.sum
cmd/osdy/
└── main.go                         # signal context, process exit only
internal/core/
├── types.go                        # IDs, enums, estimates, identity, warnings
├── snapshot.go                     # finding/root/snapshot constructors and accessors
├── types_test.go
└── snapshot_test.go
internal/scan/
├── builtins.go                     # fixed root table and lexical home resolution
├── scanner.go                      # platform-walker orchestration and bounded workers
├── descriptor.go                   # scan-owned walker/fact/job contracts; no syscalls
├── finalize.go                     # canonical sort, identity accounting, grouping
├── builtins_test.go
├── scanner_test.go                 # fake descriptor walker
└── finalize_test.go
internal/platform/macos/
├── walker_darwin.go                # descriptor-relative root acquisition/enumeration
├── walker_other.go                 # typed unsupported-environment result
├── metadata_darwin.go              # final-descriptor fact conversion, shared by walker
├── walker_darwin_test.go           # disposable macOS fixtures and syscall seams
└── metadata_darwin_test.go
internal/report/
├── json.go                         # schema-v1 ordered DTO projection
├── text.go                         # deterministic human report
├── report_test.go
└── testdata/                       # text and JSON goldens
internal/tui/
├── model.go                        # messages and Update state
├── view.go                         # pure read-only rendering
├── run.go                          # Bubble Tea terminal lifecycle
└── model_test.go
internal/cli/
├── command.go                      # Cobra command and validation
├── run.go                          # one-scan orchestration, streams, exit mapping
└── command_test.go
```

There is no separate rules package: five static definitions fit coherently in `internal/scan/builtins.go`. There is no package per enum, renderer, platform field, or TUI screen. `internal/core` imports only the standard library; dependency direction is `cli -> scan/report/tui/platform`, presentations and scan -> `core`, and `platform/macos` implements the scan-owned descriptor-walker port without a reverse dependency. Only platform files contain Darwin build tags, raw file descriptors, `openat`/directory-entry operations, or syscall constants; `internal/scan` imports none of them.

## Dependencies

| Dependency | Decision and purpose |
| --- | --- |
| Go 1.24+ standard library | Use `context`, `encoding/json`, `errors`, `io`, `io/fs`, `os`, `path/filepath` only for lexical root/display-path handling, `sort`, and `sync`. Production descendant traversal does not use `filepath.WalkDir`. No shelling out to `du`, cache tools, or `diskutil`. |
| `golang.org/x/sys/unix` (already present) | Supply Darwin descriptor-relative open, directory-entry, stat, statfs, and close operations behind `internal/platform/macos`; do not expose descriptors or syscall flags to `internal/scan`. |
| Bubble Tea v2 (`charm.land/bubbletea/v2`) | Own the minimal event loop and reliable terminal setup/restoration. Scanning is complete before the model is constructed. |
| Lip Gloss v2 (`charm.land/lipgloss/v2`) | Apply restrained status/category styles while keeping all wording in the local view package. |
| Bubbles v2 (`charm.land/bubbles/v2`) | Use only its viewport component because finding/warning detail can exceed terminal height. No list, table, selector, or form component is needed. |
| Cobra (`github.com/spf13/cobra`) | Validate the small command grammar and keep `RunE` thin. Viper and configuration frameworks are excluded. |

Direct versions are pinned by `go.mod`/`go.sum` during implementation. Indirect Charm dependencies are accepted only as resolved transitive dependencies. No metadata, logging, UUID, time, or JSON framework is added.

## Domain contract

All domain structs have unexported fields. Constructors validate invariants and deep-copy incoming slices; accessors return values or copied slices. This makes snapshots immutable by convention and prevents presentation state from mutating scan facts.

### Named values and constructors

| Type | Allowed values and construction |
| --- | --- |
| `AreaID` | `npm-cache`, `homebrew-cache`, `gradle-caches`, `xcode-derived-data`, `core-simulator`; `ParseAreaID(string)` rejects every other value. `AreaRank` is a fixed switch, never a map iteration. |
| `RootStatus` | `scanned`, `missing`, `inaccessible`, `skipped`, `boundary_limited`, `partial`, `cancelled`; `ParseRootStatus` validates external values. |
| `RootReasonCode` | `completed`, `not_found`, `root_inspection_failed`, `root_symlink`, `root_device_boundary`, `root_not_directory`, `cancelled_before_start`, `device_boundary`, `entry_visibility_gap`, `cancelled_during_scan`. A root always has one reason. |
| `Outcome` | `complete`, `partial`, `cancelled`; it is derived by `NewSnapshot`, not selected by callers. |
| `Risk` | Only `manual_review` in schema v1; the named type prevents free-form labels. |
| `WarningCode` | `root_inaccessible`, `root_symlink`, `root_device_boundary`, `root_not_directory`, `entry_inaccessible`, `entry_changed`, `symlink_skipped`, `device_boundary`, `unsupported_entry_type`, `identity_unavailable`, `allocation_unavailable`, `hard_link_observed`, `identity_already_accounted`, `entry_limit_reached`, `path_budget_reached`, `cancelled_during_root`, `cancelled_before_root`, `apfs_estimate_uncertain`, `core_simulator_semantics`. `ParseWarningCode` rejects unknown codes. |
| `FilesystemIdentity` | Opaque `(device uint64, inode uint64)` from `NewFilesystemIdentity`; unknown identity is a separate optional value, not a zero-value sentinel. Identity is scan-internal accounting evidence and is not an actionable candidate ID. |
| `Allocation` | `KnownAllocation(bytes)` or `UnknownAllocation()`; unknown never falls back to logical bytes. |
| `ValueEstimate` | `known_bytes`, basis, and completeness (`complete`, `incomplete`, `unknown`). Logical basis is `regular_file_stat_size`; allocated basis is `stat_blocks_512`. |
| `Estimate` | Logical and allocated `ValueEstimate` plus `deduplication_complete`; `NewEstimate` rejects inconsistent completeness/basis combinations. |
| `Accounting` | Counts measured objects, canonically attributed objects, already-accounted paths, hard-link observations, and objects with unknown identity. |
| `Finding` | `NewFinding` requires a valid area, normalized home-relative display path, fixed built-in reason, `manual_review`, estimate, accounting, and sorted/deduplicated warning codes. |
| `RootObservation` | `NewRootObservation` validates status/reason combinations, area/display path, estimate, finding count, and warning codes. |
| `Snapshot` | `NewSnapshot(ruleSetVersion, policy, roots, findings, warnings)` requires exactly five unique roots, validates canonical facts, derives outcome/completeness and aggregate estimate, canonicalizes collections, and fixes schema version `1`. |

`Warning` contains area, display path, code, optional related display path, and a stable code-owned message. Raw operating-system error strings are not serialized; stable error classes and paths provide useful context without making fixtures or output platform-message dependent.

### Estimate semantics

Only regular-file data contributes bytes. Logical bytes are the sum of `lstat` file sizes. Allocated bytes are the sum of non-negative Darwin `st_blocks * 512` values when available. Directory metadata, symbolic links, sockets, devices, and other special objects contribute zero.

A complete zero-file root has complete zero logical and allocated estimates. Visibility gaps make the affected value incomplete. If no canonical measured file has allocation data, allocated completeness is `unknown`; mixed known/unknown allocation produces an `incomplete` known subtotal. Unknown identity does not suppress observed bytes, but sets `deduplication_complete=false` and emits a warning.

Every area with measured data carries the APFS caveat. `st_blocks` is observed allocation, not exclusive or guaranteed reclaimable storage: clones with different inodes may share extents, snapshots may retain blocks, compression and sparse files alter accounting, and hard links share one identity.

### Finding granularity

A finding groups one immediate child of a built-in root and all safe regular files beneath that child. A regular file directly under the root is its own group. An empty root has no findings but still has a complete root observation. This yields useful cache/artifact groups without turning the public snapshot into an exhaustive file inventory.

Nested paths remain visible only when needed for warning context. A repeated identity's later path gets `identity_already_accounted`, contributes zero bytes, and increments the containing finding's already-accounted count. This is identity accounting, never duplicate-file analysis.

## Snapshot schema and deterministic output

Schema v1 uses only ordered structs, scalar values, and non-nil slices. Maps may be used internally for lookup but are never marshaled or iterated to emit facts. There is no wall clock, hostname, username, random run ID, or absolute home path in v1. Display paths use `~` plus slash-separated lexical relative paths; names are not case-folded or Unicode-normalized.

The top-level JSON field order is fixed by a dedicated DTO declaration:

```text
schema_version
rule_set_version
outcome
complete
scan_policy
estimate
roots
findings
warnings
```

Nested field order is likewise declared once:

- policy: `no_symlink_following`, `same_device_only`, `max_entries_per_root`, `max_path_bytes_per_root`, `max_depth_per_root`;
- estimate: `logical`, `allocated`, `deduplication_complete`;
- value estimate: `known_bytes`, `basis`, `completeness`;
- root: `area_id`, `display_name`, `display_path`, `status`, `reason`, `estimate`, `finding_count`, `warning_codes`;
- finding: `area_id`, `display_path`, `reason`, `risk`, `estimate`, `accounting`, `warning_codes`;
- warning: `area_id`, `path`, `code`, `related_path`, `message`.

`related_path` is explicitly `null` when absent, and empty collections encode as `[]`, not `null`. JSON is produced in memory with `encoding/json`, followed by one newline. Identical finalized snapshots therefore produce identical bytes.

Canonical collection order is:

1. roots by the five-area rank;
2. findings by area rank, normalized display path, then fixed finding reason;
3. warnings by area rank, normalized path, warning code, related path, then stable message;
4. warning-code lists lexically sorted and deduplicated.

Text follows those same collections and prints schema/outcome first, then all roots, findings, estimate bases/completeness, and all warnings. It includes “Manual review required” and “Estimate—not guaranteed reclaimable space,” plus the APFS explanation. Text formatting never becomes an input to JSON or TUI.

## Built-in root resolution

`scan.HomeResolver` has one method returning the current home. Production construction wraps `os.UserHomeDir`; tests inject a fixture resolver whose home is always a `t.TempDir()` descendant. No flag, positional argument, environment override owned by OsdyCleaner, config file, or tool lookup can supply roots.

`ResolveBuiltins` validates an absolute, inspectable home and joins this static ordered table without `EvalSymlinks`:

| Area | Relative path | Finding reason |
| --- | --- | --- |
| npm cache | `.npm` | `built_in_npm_cache_entry` |
| Homebrew cache | `Library/Caches/Homebrew` | `built_in_homebrew_cache_entry` |
| Gradle caches | `.gradle/caches` | `built_in_gradle_cache_entry` |
| Xcode DerivedData | `Library/Developer/Xcode/DerivedData` | `built_in_xcode_derived_data_entry` |
| CoreSimulator | `Library/Developer/CoreSimulator` | `built_in_core_simulator_entry` |

`ResolveBuiltins` performs only lexical validation and returns relative component chains. Platform acquisition establishes trust in the required order: open `/` first, open every absolute home component relative to the current descriptor, validate the final home descriptor, and then open each built-in suffix component relative to that home descriptor. Every component uses read-only `O_NOFOLLOW|O_DIRECTORY|O_CLOEXEC|O_NONBLOCK` semantics. A symbolic-link component makes that root `skipped/root_symlink`; a missing component makes it `missing`; no target is resolved. Home locality and mount identity are accepted only from its final descriptor. Each root must match the trusted home device/mount; every descendant must match the accepted root boundary. A non-local home is globally unsupported, while a separately mounted built-in root is `skipped/root_device_boundary`.

The platform session owns `/`, home, root, traversal, and worker-anchor descriptors. Neither `internal/scan` nor a callback receives an integer FD or `*os.File`. Extending the current macOS component-acquisition code is preferred, but its current “reopen an absolute path and retain every component FD” ownership model must not be copied into traversal: acquisition advances one descriptor at a time, closes the superseded descriptor after the next component is validated, and transfers only an opaque root capability to the platform walker.

## Descriptor-walker contract

The scan package owns a narrow, platform-neutral port; Darwin implements it and non-Darwin returns a typed unsupported-environment error before home resolution or traversal. The API names are illustrative, but ownership is normative:

```go
type DescriptorWalker interface {
    WalkRoot(ctx context.Context, home AbsoluteComponents, root RelativeComponents, limits WalkLimits, visit Visitor) (RootFacts, error)
}

type Visitor interface {
    Directory(EntryFact) error // already opened, identity- and boundary-checked
    Regular(RegularJob) error  // relative chain plus expected enumeration identity
    Skipped(SkipFact) error
}
```

`AbsoluteComponents` and `RelativeComponents` are validated value types: components are non-empty and cannot be `.`, `..`, contain `/`, or contain NUL. `EntryFact` contains normalized relative components, display-safe kind, and expected `(device,inode)`. `RegularJob` additionally carries the accepted identity chain for every ancestor plus the final expected file identity; it contains no descriptor. The opaque platform call owns all descriptors until return. Visitor facts may be retained; capabilities may not.

| Boundary decision | Required behavior |
| --- | --- |
| Trusted acquisition | Open `/`, then home and root components with no-follow/directory/cloexec/nonblock flags; accept facts only from final descriptors. |
| Directory enumeration | Read directory records from the already-open directory descriptor, validate each kernel-provided name, reject malformed records, and sort the current directory's accepted names before callbacks/descent. Directory-entry kind is a hint only. |
| Child directory | Open with `openat(parentFD, name, O_RDONLY | O_NOFOLLOW | O_DIRECTORY | O_CLOEXEC | O_NONBLOCK)`, then`fstat`/`fstatfs`; require the enumerated inode when supplied to equal the final descriptor identity and require the root device/mount/local boundary before enumeration. |
| Regular file | Open no-follow/nonblock relative to its parent to classify and capture expected identity. A bounded worker receives only the relative component chain plus expected ancestor/file identities, safely reopens it from the private trusted-root anchor, validates each directory identity before opening the next component, compares final-descriptor identity/boundary, records metadata, and reads no content. |
| Symlink or special object | `ELOOP` or final symlink classification is `symlink_skipped`; unsupported or unopenable special objects contribute no bytes and produce the existing scoped warning/partial state. |
| Concurrent rename | An opened descriptor remains bound to the opened object. Renaming its pathname cannot redirect its enumeration. A child that disappears before open, or whose enumerated and final identities/kinds differ, is `entry_changed`, contributes no bytes, and does not cause target traversal. Safe siblings continue. |

Darwin directory records supply name, inode, and type where available. Names and record lengths are validated before use; `.` and `..` are ignored. Type never authorizes descent. Directories and accepted regular files require final-descriptor `fstat`; an unavailable/zero enumeration inode is conservatively treated as changed for directories and as incomplete evidence for regular files rather than weakening the no-redirection claim. `fstatfs` supplies locality and a root mount key (device plus filesystem identity where available). A child outside that key emits `device_boundary` and is never enumerated or counted.

## Scanner pipeline

Roots remain serial in canonical area order. Within one root, the descriptor walker uses deterministic, depth-first traversal and bounded regular-file workers:

1. **Acquire trust:** check cancellation, acquire home and root descriptor-relatively, and classify root missing, inaccessible, symlink, non-directory, non-local, or mount mismatch before enumeration.
2. **Enumerate bound directory:** read and validate entries from the accepted directory FD, stop at entry/path limits, sort names lexically, and process them without constructing an absolute descendant path.
3. **Open and gate child:** open the child relative to its parent FD with no-follow flags; derive kind and identity from the final descriptor, compare the directory record identity when available, and check local device/mount scope before descent or accounting.
4. **Descend or queue:** recurse into an accepted child directory while retaining only the DFS ancestor descriptors. For a regular file, close the classification descriptor after creating a bounded job containing relative components and the accepted ancestor/file identity chain; a fixed worker safely reopens from the private root anchor, rolls descriptors one component at a time, validates each ancestor before continuing, and compares the final identity before accepting metadata.
5. **Collect deterministically:** a bounded result channel is drained concurrently, but completion order is ignored. Raw facts are sorted by normalized display path and kind before cross-root identity attribution.
6. **Stop, drain, and close:** cancellation or a limit stops new opens and jobs immediately. Close job input, allow only active syscalls to return, drain results, join fixed workers, close child-before-parent in LIFO order, and return cancelled/partial facts.
7. **Finalize:** group canonically attributed observations into immediate-child findings, discard current-root raw facts, and construct the shared snapshot after all five root observations exist.

No goroutine is created per file. Worker count and channels are fixed. At most one descriptor per DFS ancestor, the platform's small trusted anchors, and a constant number per active worker are live. Worker path acquisition closes each superseded intermediate descriptor, so worker FD use is constant rather than proportional to path depth. A depth limit must be added to `WalkLimits`; exceeding it is the same bounded partial condition as entry/path exhaustion.

Cancellation is checked before directory reads, after each returned entry batch, before child opens, before enqueue, and before a worker accepts a job. A blocking kernel call may finish, but its result is ignored after cancellation except for cleanup. Operation errors remain primary; cleanup proceeds child-before-parent. If no earlier error exists, a close failure becomes the scoped traversal error and invalidates facts whose ownership was not safely finalized. If an earlier error exists, close failures are retained as secondary diagnostic evidence without replacing its stable class.

Concurrency cannot affect attribution because worker completion order is never consulted by identity accounting. The `seen identity -> canonical display path` map is populated only while consuming the sorted stream.

## macOS descriptor adapter

The Darwin implementation extends or replaces the existing component metadata adapter behind the scan-owned port. It uses the already-present `golang.org/x/sys/unix` operations for descriptor-relative open and directory reads, plus typed `fstat`, `fstatfs`, and close operations. It validates numeric ranges, exposes bytes only for regular files, never invents unavailable values, and never reads regular-file content. Test-only syscall seams may substitute operations, but production construction has no fake walk callback, fake stat type, pathname revalidation hook, or unreachable changed-error class.

The adapter retains reachable stable classes: missing, symlink, non-directory, device/mount mismatch, non-local filesystem, inaccessible, invalid directory record/metadata, cancellation, limit, and close failure. The scanner maps these to existing snapshot warnings and statuses; `entry_changed` is produced from the adapter's real enumeration-identity/final-descriptor comparison, not from a synthetic metadata hook.

`os.Root` is explicitly not an approved substitute. It may be reconsidered only if an implementation proves that directory enumeration is bound to the accepted root/directory descriptors, child opens cannot be redirected, final descriptor identities are comparable with enumeration records, mount/local checks occur before descent, and ownership/close/cancellation semantics satisfy this section. Path confinement alone is insufficient.

## Exact status and outcome derivation

Preflight statuses are terminal: missing root -> `missing/not_found`; uninspectable root -> `inaccessible/root_inspection_failed`; symbolic-link component -> `skipped/root_symlink`; other-device root -> `skipped/root_device_boundary`; non-directory root -> `skipped/root_not_directory`; cancellation before starting -> `skipped/cancelled_before_start`.

For a traversed root, precedence is exact:

1. cancellation observed -> `cancelled/cancelled_during_scan`;
2. any non-boundary visibility gap, unknown required metadata, unsupported entry, or resource limit -> `partial/entry_visibility_gap`;
3. otherwise any descendant device boundary -> `boundary_limited/device_boundary`;
4. otherwise -> `scanned/completed`.

All lower-precedence conditions remain in warnings. A permitted skipped symlink entry is warned but does not by itself make a root partial because its target is outside observable scope. APFS and hard-link caveats also do not change status.

Global derivation is equally fixed:

- `cancelled` if cancellation stopped a root or prevented a later root from starting; completed roots remain, the active root is cancelled, and every later root is skipped with cancellation context;
- otherwise `complete` only when every root is `scanned` or `missing`;
- otherwise `partial`.

`complete` is true only for outcome `complete`. Partial and cancelled aggregate estimates are incomplete. All-missing is complete with zero totals and no findings. Cancellation outranks any partial or boundary condition.

## Shared presentations and TUI

The application invokes the scanner exactly once. Text and JSON render to an in-memory byte slice before stdout is touched. The TUI model is constructed only from the returned `core.Snapshot`; `Init` performs no I/O and returns no scan command.

The model stores the snapshot, active area rank, current pane (`summary`, `area`, or `warnings`), viewport, width, height, and quitting flag. It stores no candidate selection, operation intent, or alternate totals. Messages are only Bubble Tea key/window messages plus viewport messages.

`Update` behavior is limited to:

- left/right or `h/l`: move among summary and the five canonical areas;
- up/down or `j/k`, Page Up/Page Down: scroll details;
- Tab: switch between area facts and warnings/details;
- `q`, Escape, or Ctrl-C: quit without another command;
- window-size messages: resize the viewport.

`View` is pure over model state and snapshot accessors. It prominently renders “Read-only scan,” outcome/completeness, “Manual review required,” and “Estimate—not guaranteed reclaimable space.” The footer says only `category`, `details`, `scroll`, and `quit`; there are no checkbox, cursor-selection, confirmation, or follow-up-operation affordances. Cancelled snapshots render cancellation and incompleteness before available facts.

`run.go` validates an interactive terminal, starts Bubble Tea with alternate-screen ownership, and relies on Bubble Tea's normal quit/context/error path to restore terminal state. It emits diagnostics only after `Run` returns and restoration has occurred. No custom raw-mode or ANSI lifecycle is added; SIGKILL remains inherently unrestorable. A viewer error is a global reporting failure and never triggers a rescan or fallback operation.

## CLI, streams, signals, and exits

The only command is:

```text
osdy scan [--format text|json|tui]
```

The default is `text`. Cobra uses `NoArgs`; any positional value, unknown path-like flag, or format outside the three values is rejected before scanner construction. TUI additionally requires character-device stdin/stdout and otherwise returns invalid invocation code 2 before traversal.

`cmd/osdy/main.go` creates a `signal.NotifyContext` for interrupt and termination signals, delegates to `internal/cli`, stops signal delivery, and calls `os.Exit`. Cobra usage and automatic error printing are silenced; `internal/cli` owns one concise stderr diagnostic. There is no progress output in v1.

For text/JSON, the finalized representation alone goes to stdout. Snapshot warnings remain inside it. JSON is exactly one document plus newline. Global diagnostics go only to stderr. Serialization finishes before writing; a write failure is reported as code 1, although an operating-system partial write cannot be retracted.

Exit mapping is centralized and tested:

| Condition | Code |
| --- | ---: |
| Complete snapshot, including all missing | 0 |
| Global initialization/finalization/serialization/output/viewer failure | 1 |
| Cobra arguments, positional path, format, or non-interactive TUI request invalid | 2 |
| Valid non-cancelled partial snapshot | 3 |
| Non-Darwin or unsupported home filesystem | 4 |
| Valid cancelled snapshot, or signal interruption of the viewer | 130 |

A root warning never maps to 1. A cancelled snapshot maps to 130 even when it also contains partial conditions. In TUI mode a finalized cancelled snapshot is a valid model input and displays correctly; if the process signal context already requests immediate termination, the runner may exit without entering a new alternate-screen session, still returning 130.

## Performance and resource limits

Defaults are injected scanner options, not package globals:

| Control | Initial default |
| --- | ---: |
| Metadata workers per active root | 8 |
| Job/result channel capacity | 256 each |
| Concurrent roots | 1 |
| Active descriptor budget per root | 96 |
| Maximum traversal depth per root | 64 components |
| Visited entries per root | 250,000 |
| Cumulative relative-path bytes per root | 64 MiB |

Regular files are opened only for metadata and their content is never read. No goroutine is created per entry, and no wall-clock timeout is imposed. Queue capacity provides backpressure. Goroutine count is bounded by workers plus walker/collector coordination. Directory FDs are bounded by DFS depth rather than total tree size; workers use rolling descriptor acquisition and a shared private root anchor. The platform enforces the explicit descriptor budget before opening, so an unexpectedly low process limit becomes a scoped partial result rather than uncontrolled exhaustion. Memory is `O(unique identities + current-root observations + one bounded directory batch + finalized grouped facts)`.

Workers and queue sizes are implementation tuning and do not enter the snapshot. Fact-changing entry/path/depth limits are fields in `scan_policy`; a semantic built-in-policy change bumps the rule-set version. Hitting an entry, path, depth, descriptor, job, or channel bound stops new work for that root, emits stable warning evidence using the existing resource-limit vocabulary, and makes it partial. Cancellation tests synchronize on fake walker/worker channels rather than flaky sleep deadlines; cancellation stops new jobs immediately, while an already-running filesystem syscall is allowed to return.

## Strict TDD and test seams

Implementation proceeds RED -> GREEN -> TRIANGULATE -> REFACTOR within each review unit:

1. Table-driven `core` tests establish enum parsing, constructor rejection, estimate states, status precedence, exact outcome derivation, deep-copy accessors, and all-missing behavior.
2. Finalizer tests feed reordered raw observations and controlled identities to prove canonical attribution, cross-area hard links, zero incremental bytes, warning sorting/deduplication, and unknown identity/allocation behavior.
3. Built-in resolver tests use only fixture home resolvers and assert exact order/relative component chains. A poison resolver/scanner proves invalid CLI paths fail before platform construction.
4. Scanner tests inject a fake `DescriptorWalker`, never `filepath.WalkDir`. Scripted facts cover normal/empty roots, inaccessible root and entry, root/entry symlinks, mount boundaries, special entries, limits, repeated identities, deterministic callback permutations, and stable warning/fact mapping. The fake exposes no FD or syscall behavior and therefore tests only the consumer contract.
5. Cancellation tests block a fake walker or worker job, cancel deterministically, release it, and assert completed/active/unstarted root states, stopped enqueueing, drained channels, and no leaked or newly started work.
6. Report goldens use a fixed constructed snapshot. Tests serialize discovery permutations to byte-identical text/JSON, parse JSON for typed/null fields, and verify all snapshot facts are represented. Golden updates require an explicit test `-update` flag, followed by a run without it.
7. TUI tests call `Model.Update` directly for navigation, resize, scroll, quit, complete/partial/cancelled views, and required/forbidden language. Use a full interactive harness only if direct tests cannot prove a terminal lifecycle regression; do not introduce it by default.
8. CLI tests inject scanner, renderer, viewer runner, terminal detector, stdout, and stderr. They cover every exit class, one scan call, JSON stream purity, output failures, signal cancellation, and terminal errors.
9. Darwin-tagged adapter tests use real disposable directory trees plus narrow descriptor-operation seams. They prove trusted `/`/home/root acquisition order, exact component flags, descriptor-bound enumeration, malformed-name rejection, lexical sibling order, child `openat` parentage, final-descriptor facts, inode mismatch handling, symlink non-following, mount/local gating before descent, rolling worker reopen from relative chains, regular-file metadata without content reads, depth/FD limits, cancellation checkpoints, and child-before-parent close/error precedence. Rename fixtures open a directory, rename/replace its pathname, and prove enumeration remains on the opened object; disappearance/replacement before child open produces `entry_changed` with no target descent. Tests never scan the real home, require elevated privileges, or require a real external mount.

Narrow package tests run first, then `go test ./...` and `go vet ./...`. Goldens are rerun without update. No acceptance test needs network, external tools, elevated permission, a personal path, or a real external volume.

## Failure modes

| Failure | Designed behavior |
| --- | --- |
| Unsupported OS/non-local home | No traversal or snapshot; stderr diagnostic and code 4. |
| Home/rule initialization invalid | No traversal or snapshot; wrapped global error and code 1. |
| Root missing/inaccessible/redirected | Descriptor-relative acquisition never follows the redirected component; finalize the scoped observation and continue other roots. |
| Entry permission or special type | The adapter returns a reachable scoped class; the scanner emits the stable warning, continues safe siblings, and makes the root partial when visibility is incomplete. |
| Child disappears before relative open | Emit `entry_changed`, contribute no bytes, do not traverse any replacement target, and continue safe siblings. |
| Child identity/kind changes | Compare directory-record evidence with the final descriptor and worker reopen; mismatch emits `entry_changed`, contributes no bytes, and makes the root partial. |
| Opened directory is renamed/replaced by pathname | Continue enumerating the already-open descriptor; the new pathname cannot redirect enumeration. A later child open remains relative to that descriptor. |
| Symlink or device/mount boundary | Never measure or enumerate the target; warning and skipped/boundary status. |
| Allocation/identity unavailable | Preserve known subtotal, mark uncertainty, never substitute values. |
| Entry/path/depth/descriptor budget exhausted | Stop new work for the current root, close owned descriptors, and finalize trustworthy partial facts. |
| Cancellation | Stop enqueueing, drain completed facts, finalize cancelled snapshot when invariants hold. |
| Finalizer invariant violation | No purported snapshot; code 1. |
| JSON/text serialization failure | No write attempted; code 1. |
| stdout or terminal failure | Report separately on stderr; code 1; never rescan. |
| Concurrent filesystem mutation | Accepted descriptors remain bound; evidence that changes before child acceptance becomes `entry_changed`. The scan is observational, not transactional. |

## WU5A disposition

The current `serialScanner` candidate is **not accepted** and must not be incrementally promoted to production:

| Verification failure | Why it blocks acceptance |
| --- | --- |
| Invalid root validation order | It calls `inspect(rootPath)` before a trusted `/` -> home -> built-in descriptor chain exists, so the first accepted root fact is not anchored to the required capability boundary. |
| `filepath.WalkDir` race | After a callback validates a directory, `WalkDir` reopens that directory by pathname to enumerate it. A concurrent rename, symlink replacement, or mount swap can therefore redirect which names are discovered before later callback pruning. Root/directory metadata callbacks cannot close this gap. |
| Fake production seams | Its injected `walk func` and path metadata function permit test states that do not demonstrate production descriptor ownership, open flags, enumeration source, or close behavior. Passing those tests is not evidence for the production invariant. |
| Incorrect completeness and identity | The candidate reports scanned/complete zero-byte roots without regular-file collection, treats limited metadata as sufficient identity evidence, and does not establish canonical file identity/completeness from accepted final descriptors. |
| Missing changed-directory evidence | Tests fake callback entries but do not open a real directory, rename/replace its pathname, and prove enumeration remains bound to the original descriptor or that a vanished child becomes `entry_changed`. |

The candidate may donate domain mapping code only after its traversal assumptions are removed. Production `filepath.WalkDir`, path-based descendant inspection, and tests that merely wrap `filepath.WalkDir` must be deleted from the scanner boundary.

## Descriptor traversal risks

| Risk | Mitigation and acceptance evidence |
| --- | --- |
| Darwin directory-record parsing is subtle | Keep parsing in the platform adapter, validate record lengths/names, fuzz parser inputs, and verify against disposable real directories. |
| Descriptor exhaustion on deep trees | DFS retains descriptors only to bounded depth, workers roll component descriptors, and an explicit active-FD budget fails the root partial before uncontrolled exhaustion. |
| Worker reopen observes a replacement | Jobs carry expected ancestor and file identities; every component is reopened no-follow from the trusted root, each ancestor is checked before continuing, and any identity/boundary mismatch becomes `entry_changed` with zero bytes. |
| Concurrent mutation prevents a transactionally complete view | State explicitly remains observational; descriptor binding prevents redirection, while disappeared or changed entries become partial evidence rather than guessed facts. |
| Darwin-only implementation complicates portability | Keep build tags/syscalls in the platform package and return a typed unsupported result elsewhere; later ports implement the same scan-owned contract. |
| Close errors obscure primary failures | Use one ownership ledger, LIFO close, primary-error precedence, and test exact close counts/order. |

## Security and privacy boundary

The binary observes metadata only beneath fixed home-relative roots. It performs no network request, telemetry, shell/executable call, privilege request, config-root lookup, or write to scanned storage. Home prefixes are rendered as `~`, but nested cache names and warning paths can still be sensitive; reports stay local unless the user redirects or transmits them.

Schema v1 intentionally omits actionable candidate IDs and does not expose identities as authorization tokens. A report is evidence for manual review, not permission for a later operation. Later mutation work must define a separate identity-bound plan and revalidation contract.

## Rollback and compatibility

Rollback removes/disables `scan`, the module/packages, and publication of schema v1. There is no filesystem recovery, migration, daemon, or application state to unwind. Existing redirected reports remain inert documents and must never be accepted as later operation authorization. A breaking snapshot meaning or field change requires a new schema version; built-in policy changes that preserve schema shape require a new rule-set version.

## Rejected alternatives

- Arbitrary paths, whole-disk discovery, environment/tool cache overrides, and concurrent roots expand privacy/safety scope and are rejected.
- Shelling out to `du`, `find`, `diskutil`, or product tools violates the local typed adapter boundary.
- Following root/entry symlinks, crossing devices, or equating logical with unknown allocation would contradict the snapshot's evidence claims.
- One goroutine per entry, unbounded channels, or streaming results directly to stdout make resource use or output order depend on discovery timing.
- Map-backed JSON and presentation-specific result structs risk non-determinism and disagreement.
- A package per type/rule/renderer adds navigation cost without an independent boundary.
- A live dashboard, finding selector, or follow-up workflow would create a second state model. The chosen finalized-snapshot viewer passes the proposal's coherence gate without those semantics.
- Deferring the viewer is unnecessary because the bounded model is a pure projection of the same snapshot; richer TUI behavior remains deferred.
- `filepath.WalkDir` and any traversal that reopens an accepted directory by absolute pathname are rejected because callback validation occurs before pathname-based enumeration and cannot prevent rename/symlink/mount redirection.
- `os.Root`/`os.OpenInRoot` are rejected as a substitute unless their concrete use proves descriptor-bound directory enumeration, final identity comparison, mount/local checks before descent, bounded ownership, and cancellation semantics; confinement alone is insufficient.

## Review workload and implementation slicing

The likely implementation is approximately 2,200-3,000 authored Go/test/golden lines, excluding `go.sum`; it is unavoidably above the repository's 400-line review-risk threshold as one change. Do not hide this in a single implementation review. The tasks artifact should split it into coherent, independently tested units targeting roughly 250-400 authored lines each:

1. module bootstrap plus core named values/constructor tests;
2. snapshot/estimate construction and status/outcome tests;
3. deterministic finalizer and identity-accounting tests;
4. built-in resolver plus scan-owned descriptor-walker contract and fake tests;
5. Darwin trusted-root acquisition and descriptor enumeration with disposable rename/symlink fixtures;
6. bounded regular-file reopen workers, mount/identity checks, close/error/limit/cancellation evidence;
7. deterministic JSON/text DTOs and goldens;
8. read-only Bubble Tea viewport and direct Update tests;
9. Cobra orchestration, streams, signals, exits, and final acceptance checks.

If any unit exceeds 400 authored lines, split its tests and green implementation into adjacent review commits/PRs while retaining RED/GREEN evidence and a compiling boundary. The viewer is the first scope to simplify, not the five roots, snapshot integrity, or scanner safety.

### Next task-planning recommendation

Replace the failed WU5A task rather than amending its acceptance claim. Plan the next review unit as **descriptor walker contract plus Darwin root/directory proof**, before regular-file accounting: (1) remove production `filepath.WalkDir` and path-inspection seams, (2) add the fake descriptor walker for scan-domain status/callback tests, (3) implement trusted `/` -> home -> root acquisition and descriptor-bound lexical directory enumeration, and (4) land disposable macOS rename/symlink/close-order fixtures. Do not start worker accounting until this unit independently proves that replacing an opened directory pathname cannot redirect enumeration and that a child disappearance yields `entry_changed` without target traversal.
