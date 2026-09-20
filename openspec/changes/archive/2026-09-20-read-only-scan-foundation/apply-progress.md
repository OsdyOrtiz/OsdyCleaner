# Apply Progress: Read-Only Scan Foundation

> **Recovery provenance — exact prior prose is unavailable.** This concise cumulative receipt replaces an accidentally overwritten progress file; it is not a recreation of lost prose. It was reconstructed from the 51-line recovery backup, `tasks.md`, current accepted source hashes, runtime records, and the supplied session/Engram facts. Backup: `/tmp/osdy-apply-progress-overwrite-recovery.md`, SHA-256 `0db0f3edf98f1089d22f302746df87a5a8f235084dd5ddb701c1836cdc1c552a`. Unavailable historical timings and per-boundary file hashes are explicitly not invented. Runtime attempt tokens are deliberately omitted.

## Current task and boundary summary

- `tasks.md` currently has 20 checked and 31 unchecked checkbox rows: WU1 through WU3 (including the recorded WU3A/WU3B split) and WU4A are accepted; WU4B through WU10 and all three parent-owned rows are unchecked.
- The accepted implementation boundary is WU1 → WU2A → WU2B → WU3A → WU3B → WU4A. Product files are coherent through corrected WU4A only.
- Module identity is `github.com/osdy/OsdyCleaner`. Delivery remains sequential `stacked-to-main`; no unresolved module choice exists.
- No WU4B, traversal, report, TUI, CLI, mutation, commit, PR, or release began. No runtime attempt token is recorded in this receipt.

## Historical accepted evidence

The runtime records provide objective outcomes and evidence revisions. “Unavailable” means the exact historical value was not present in the permitted recovery sources; it is not inferred from current files.

| Accepted unit | Accepted boundary and line count | Evidence revision | Recorded result and verification | Runtime / rollback summary |
| --- | --- | --- | --- | --- |
| WU1 | Final 380 authored lines as recorded by Engram observation `10299`; not independently reconstructable from Git/runtime records | `sha256:2ebfd03b8355f802a4287e4205c4a37c37f1b85abd8309baea3df3e0630ea186` | Zero-value filesystem identity was remediated; focused, core, full, and vet checks passed. Exact timings unavailable. | N/A: pure core value boundary. Roll back WU1 module/core named-value files. |
| WU2A | Final 398 authored lines; recorded/reconstructed-current boundary | `sha256:7b78beec26eecb3c9039b3652bde91ec3638bc4f4f128c3fbff6f2014577656e` | Estimates, findings, roots, warnings, and copy invariants passed; WU2B symbols were absent. Focused, core, full, and vet passed. Exact timings unavailable. | N/A: domain boundary. Roll back WU2A facts while retaining WU1. |
| WU2B | Final exactly 400 authored lines after size correction; recorded/reconstructed-current boundary | Interim `sha256:f10f98ac3dcaa2d3ba88c41223e0e1ef7323218b3d0d7f4e8b953784e736802b`; final `sha256:2c63312e1dbfc8653e6b36126cbca0f4256c95cd49d888eaf41576cbe5b945ab` | Immutable snapshot/policy work and aggregate-completeness correction passed; focused/core/full/vet/gofmt and independent semantic verification passed. Exact timings unavailable. | N/A: immutable snapshot boundary. Roll back WU2B policy/snapshot additions while retaining WU2A. |
| WU3A | Final 399 authored lines after fail-closed and total-order remediation; recorded/reconstructed-current boundary | `sha256:e6c55ab731e07fac1fbba64e613062650e9543d42493e5581e50f526d38859ba` | Invalid raw paths now fail closed and ordering is total; focused, core, full, vet, and gofmt passed. Recorded timings: regressions 0.376s, focused 0.241s, core 0.267s, full core 0.221s/scan 0.405s. | N/A: pure finalizer tests. Roll back WU3A finalizer implementation/tests while retaining WU2B. |
| WU3B | +83/-0 authored lines; recorded/reconstructed-current boundary | `sha256:85da59872d70d8a463b7df105cd48a26cbe83cb9c32b09b88cc0a5c9be26920f` | Triangulation/refactor focused, package, core, full, and vet checks passed. Exact timings unavailable. | N/A: pure finalizer boundary. Roll back the adjacent WU3B triangulation/refactor changes only. |
| WU4A | Original +164 lines; corrected cumulative 182 authored lines; recorded/reconstructed-current boundary | Failed `sha256:088de4076c84954699f661379dd6bd2820465dd861e2b442a3fa16614673d40e`; corrected `sha256:ccc51e5791985cf3819c19c68128ee0efcbfcf6b8cf0f1fe08389fbbe60cf980` | Original evidence failed independent clean-home verification. The corrected resolver/table boundary passed focused, package, core, full, vet, format, and source-safety checks. | N/A: injected resolver/table boundary. Roll back only the lexical-clean guard and correction tests from the two WU4A files. |

### Historical availability limits

- Exact historical file hashes for WU1 through WU3B, except facts explicitly supplied above, are unavailable in the recovery sources.
- Exact historical command timings are unavailable except the WU3A timings recorded above and the WU4A timings preserved below.
- WU1's final 380-line total is sourced specifically from Engram observation `10299`; runtime records alone establish the 369-line initial candidate and bounded remediation, not the final total.
- These boundaries are **recorded/reconstructed-current** from receipts, runtime records, explicitly cited Engram facts, and current files; no Git diff provenance is claimed.

## Current accepted identity

These hashes were computed during this recovery; they identify the current product boundary through WU4A and the authoritative task list.

| File | SHA-256 |
| --- | --- |
| `go.mod` | `a8e4209f8cee87445d578dec35fb1368cd9b1a4c1aaf65303575eae650ee73af` |
| `internal/core/types.go` | `e525fd428f3e5549681d1995ea2a61820e10e5ff382cb4bd2c01b9a1ee86b7ea` |
| `internal/core/types_test.go` | `e79a225ca2178da8b050217093823070b207d0675b93c3ada998fbd54a456431` |
| `internal/core/snapshot.go` | `674278a981ae16ed11c7aa3ec6cce02ce7789e415b6a866433439d0bd67654c1` |
| `internal/core/snapshot_test.go` | `fb935962f06f9533e00f0db5b04f7219ef3f156465ca3394cb540329fae3153b` |
| `internal/scan/finalize.go` | `aa993faf1842cfeca3a6984bf29e7001ae530a691f3caf431240cc747a87d625` |
| `internal/scan/finalize_test.go` | `2ec604ae858ecad200e58a6ffc083898a35a0050c070fcc66fa500e614db3719` |
| `internal/scan/builtins.go` | `d72919bc648354201453dec7c126f9286b580f7d73fcf115ccc9695de8d08e82` |
| `internal/scan/builtins_test.go` | `137804f3de2c145d5501932ecf72ac57726d16f5ba86aa0ed8e809214fcff135` |
| `openspec/changes/read-only-scan-foundation/tasks.md` | `b26f02582b03780405c44cf02d14417fe0842974d6c8c326ec04ae80dac9b672` |

## WU4A clean-home correction

- Failed evidence remediated: `sha256:088de4076c84954699f661379dd6bd2820465dd861e2b442a3fa16614673d40e`.
- Defect: an absolute but lexically unclean home reached the inspector because `filepath.Join` silently cleaned it. This could accept `/fixture/./home`, `/fixture/../home`, redundant separators, and a non-root trailing separator.
- Remediation: after non-empty, trimmed, and absolute validation, `ResolveBuiltins` now rejects a home when `home != filepath.Clean(home)` before calling `InspectHome`. It does not normalize the input. The root `/` remains valid because `filepath.Clean("/") == "/"`; non-root trailing separators are rejected.
- Preserved behavior: the exact five static definitions and their order, injected resolver/inspector seams, copied definitions, resolver-error wrapping, and syscall-free/no-real-home behavior remain unchanged.

### Independent correction TDD evidence

| Stage | Command or evidence | Result |
| --- | --- | --- |
| Safety net | Pre-existing WU4A focused suite was previously green; the current focused WU4A suite also passed before settlement. | No pre-existing failure observed. |
| RED | Added table rows for `/fixture/./home`, `/fixture/../home`, `/fixture//home`, and `/fixture/home/`; each asserts an error and zero inspector calls. | `go test ./internal/scan -run '^TestResolveBuiltinsHomeValidation$' -count=1` exited 1. Each new row reported `invalid home invoked inspector 1 times`. |
| GREEN | Added the lexical identity guard in `ResolveBuiltins`. | `go test ./internal/scan -run '^(TestResolveBuiltins | TestResolveBuiltinsHomeValidation | TestResolveBuiltinsPoisonCases)$' -count=1`passed (`ok github.com/osdy/OsdyCleaner/internal/scan 0.265s`). |
| TRIANGULATE | Added `TestResolveBuiltinsCleanRootHome`, proving `/` is accepted, invokes the inspector once, returns five definitions, and keeps `/.npm`. | Included in the final focused suite; PASS. |
| REFACTOR | Ran `gofmt`; the one-condition guard was retained without further change. | All listed checks pass and `gofmt -d` has no output. |

### Verification, diagnostics, accounting, and rollback

- Focused: `go test ./internal/scan -run '^(TestResolveBuiltins|TestResolveBuiltinsHomeValidation|TestResolveBuiltinsPoisonCases)$' -count=1` -> PASS (`0.254s`).
- Scan: `go test ./internal/scan -count=1` -> PASS (`0.223s`).
- Core: `go test ./internal/core -count=1` -> PASS (`0.252s`).
- Full: `go test ./... -count=1` -> PASS (`internal/core 0.225s`, `internal/scan 0.405s`).
- Static/format: `go vet ./...` -> PASS with no diagnostics; `gofmt -w internal/scan/builtins.go internal/scan/builtins_test.go` and `gofmt -d` -> clean.
- Source diagnostics: search for `os.UserHomeDir`, `Lstat`, `Statfs`, `EvalSymlinks`, `WalkDir`, and `os.Stat` in the two WU4A files returned no matches. No `gopls` run was required or available.
- Runtime: N/A — this pure injected resolver/table boundary has no executable or scanner runtime harness, and no production home or syscall is exercised.
- Correction-only accounting against supplied pre-correction hashes: `builtins.go` +1/-1 and `builtins_test.go` +19/-1 = **20 additions + 2 deletions = 22 authored changed lines**.
- Revised cumulative WU4A accounting from absent files: `builtins.go` 57 additions and `builtins_test.go` 125 additions = **182 additions + 0 deletions = 182 authored changed lines**, within the 400-line cap.
- Corrected hashes: `internal/scan/builtins.go` `d72919bc648354201453dec7c126f9286b580f7d73fcf115ccc9695de8d08e82`; `internal/scan/builtins_test.go` `137804f3de2c145d5501932ecf72ac57726d16f5ba86aa0ed8e809214fcff135`.
- Corrected evidence revision: `sha256:ccc51e5791985cf3819c19c68128ee0efcbfcf6b8cf0f1fe08389fbbe60cf980`, computed from the ordered SHA-256 manifest of those two corrected WU4A files; it is distinct from and remediates the failed evidence revision.
- Rollback boundary: remove the lexical `filepath.Clean` identity clause and the correction-only four rejection rows, root-semantic case, and adjusted copy fixture from the two WU4A files. This restores the prior WU4A candidate only; WU4B metadata adapters and WU5 scanner work remain absent.
- Task reconciliation: `tasks.md` was re-read. All four WU4A rows remain visibly `- [x]`; all four WU4B rows remain visibly `- [ ]`; no checkbox was changed by this recovery.

### Exact next row

- [ ] **WU4B RED:** In Darwin-tagged `internal/platform/macos/metadata_darwin_test.go`, add disposable `t.TempDir()` file/directory/hard-link tests and controlled adapter-block cases for missing paths, symlinks, non-directories, device mismatch, non-local filesystems, inaccessible metadata, unknown identity/allocation, negative or overflowing block counts, and link counts; run the WU4B focused command and capture a non-zero failure caused only by missing metadata APIs. <!-- sdd-owner: implementation -->

WU4B is intentionally deferred. No verification, review, receipt approval, commit, push, PR, or release activity began beyond the recorded WU4A evidence.

---

## WU4B receipt — Darwin and non-Darwin metadata adapters

**Receipt predecessor:** SHA-256 `d64579eb04033b171ebe6c23869d20325d587994430c23db40d581d0a4ff17f0` (the complete recovered receipt was read at 86 lines before this append). This receipt is appended only; no prior text was replaced. No runtime attempt token is persisted.

### Scope, contract, and completed tasks

- Completed and visibly checked in `tasks.md`: WU4B RED, GREEN, TRIANGULATE, and REFACTOR. WU5 and later implementation rows and every parent-owned row remain unchecked.
- Added the narrow `macos.MetadataAdapter` one-path boundary. Darwin `Inspect` uses non-following `os.Lstat`, checked `syscall.Stat_t` fields, and `syscall.Statfs`; it does not walk, read content, resolve a home, mutate fixtures, or follow a symbolic link target.
- The stable typed error sentinels classify missing paths, symlinks, non-directories, device mismatch, non-local filesystems, and inaccessible metadata. On non-Darwin, `NewMetadataAdapter` returns `UnsupportedEnvironmentError`, which unwraps to `ErrUnsupportedEnvironment`, before any filesystem or home operation.
- Logical bytes, identity, and allocation are only reported when observed. Negative or overflowing block counts produce `core.UnknownAllocation()` rather than wrapping or substituting logical bytes; unavailable identity remains unknown.

### TDD Cycle Evidence

| Task | Test file / layer | Safety net | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- | --- | --- |
| WU4B | `internal/platform/macos/metadata_darwin_test.go` / Darwin unit filesystem fixture plus controlled adapter seams | N/A (three new files; existing core and scan safety net passed) | `go test ./internal/platform/macos -run '^(TestDarwinMetadataDisposable | TestDarwinMetadataClassification | TestDarwinMetadataUncertainty)$' -count=1` exited 1 only because `NewMetadataAdapter`,`MetadataAdapter`,`newMetadataAdapter`, and typed metadata errors were absent. | Same focused command passed after the minimal adapter implementation. | The focused command passed after hard-link identity/link-count, content/mode/mtime preservation, typed error, unknown identity/allocation, negative-block, and overflowing-block cases were added. | Exact required format/package/cross-compile command passed; no behavior-changing refactor was needed after formatting. |

### Verification and diagnostics

- `gofmt -w internal/platform/macos/metadata_darwin.go internal/platform/macos/metadata_other.go internal/platform/macos/metadata_darwin_test.go && go test ./internal/platform/macos -count=1 && GOOS=linux go test -c -o /dev/null ./internal/platform/macos` -> PASS (`go test` Darwin package `0.241s`; Linux compile succeeded and left no repository artifact).
- Focused command above -> PASS (`0.251s` after triangulation).
- `go test ./internal/core -count=1` -> PASS (`0.236s`); `go test ./internal/scan -count=1` -> PASS (`0.242s`); `go test ./... -count=1` -> PASS (core `0.234s`, macos `0.601s`, scan `0.420s`).
- `go vet ./...` -> PASS with no diagnostics. `gopls` was unavailable, so no LSP check could be run.
- Source-boundary scan found no `WalkDir`, content read/write, removal, real-home resolution, symlink resolution, shell, or network calls in the two production adapter files.
- Fixture immutability: `TestDarwinMetadataDisposable` writes only under `t.TempDir()`, creates one hard link, then proves the observed regular fixture's content, mode, and mtime are unchanged after adapter calls.
- Runtime: N/A — this unit has no executable, scanner traversal, or CLI boundary; its disposable metadata tests are the designed bounded harness.

### Accounting, identity, and rollback

- Authored additions from the absent three-file WU4B baseline: `metadata_darwin.go` +108, `metadata_other.go` +56, `metadata_darwin_test.go` +197; deletions 0; **361 authored changed lines**, within the 400-line cap.
- File hashes: `metadata_darwin.go` `d046c8be6db96872c1bee7128f78b67d903c9179aee8cf3025ea8051a7ba0089`; `metadata_other.go` `a1b5d36c1dd06ce5d8c8aaf871ea8c988f52ee4c22c5fbb96ea57dc0034c01ee`; `metadata_darwin_test.go` `0a94d63ee9414eb8e28fc46c6a65a87702957d6d38ca624b67e9324a91426351`.
- Ordered product-manifest evidence revision: `sha256:601a9b4b128624406308099b14f47b52b857552547732a0c5fc2a4f672b22a29`. Updated task artifact SHA-256: `74e408511f32f90d845444a7a648df3e3b5dcca2893115e227c04dd1077f3569`.
- Rollback boundary: remove exactly `internal/platform/macos/metadata_darwin.go`, `internal/platform/macos/metadata_other.go`, and `internal/platform/macos/metadata_darwin_test.go`; accepted WU4A built-ins and earlier core/finalizer work remain intact.

### Status and remaining work

- Consumed status: authoritative OpenSpec change `read-only-scan-foundation`; WU1–WU4A accepted, recovery closed, WU4B was the exact first unchecked implementation unit, sequential `stacked-to-main`, one writer, and WU4B was assigned as the independent <=400-line slice. Strict TDD was active from `openspec/config.yaml`. No unsafe `actionContext` or edit-root warning was supplied; edits stayed within the three WU4B allowed files plus required OpenSpec artifacts.
- Workload / PR boundary: WU4B only, independent 361 authored lines, sequential stacked-to-main work-unit boundary; no commit or PR created.
- Exact next unchecked implementation row: `- [ ] **WU5 RED:** Add table-driven temporary-home and fake-metadata tests for normal and empty roots, every-root-missing, inaccessible root/entry, non-directory root, root/entry symlinks, lexical escape callbacks, different-device root/entry, changed entries, unsupported special types, sibling continuation, serial root order, bounded goroutine/queue use, and before/after fixture metadata; run the focused command and record missing-scanner failures. <!-- sdd-owner: implementation -->`

---

## Planning rescope — authorized adjacent WU4C safety remediation

**Append provenance:** The user supplied SHA-256 `7f1d20a7416f0f9e9a3225b976fa6d9fcaf3db468ce122e6eed5913a824d37b3` as the current complete receipt predecessor for this append. The WU4B verifier found the earlier WU4B receipt's embedded predecessor claim non-reproducible; that embedded claim is therefore **superseded** and is not repeated here as verified evidence. This planning-only append preserves all prior recovery and WU4B receipt text byte-for-byte and records no new implementation acceptance.

### Authorized plan change

- The active `tasks.md` now inserts WU4C RED/GREEN/TRIANGULATE/REFACTOR immediately after the four checked WU4B rows and before WU5.
- WU4C is an independent 275–390 authored-changed-line slice with its own mandatory <=400-line accounting, rollback, focused verification, Darwin package verification, Linux cross-build, full suite, vet, and format diagnostics.
- The remediation requires fail-closed descriptor-backed metadata acquisition with `O_NOFOLLOW` and safe flags, descriptor `Fstat`/`Fstatfs`, `Lstat` identity/type revalidation, regular-file-only byte exposure, stable sentinel classification plus concrete typed path errors, descriptor closure on every path, and adversarial no-target-following race evidence.
- WU4B's four checked task rows remain unchanged as historical attempted implementation. Its receipt remains failed historical evidence; no checked state was reverted and no acceptance is inferred from it.
- WU5 and every later implementation row remain unchecked. WU5 now explicitly depends on accepted WU4C in addition to accepted WU4A and the preserved WU4B adapter surfaces.
- All three parent-owned rows remain unchecked. No product code, proposal, specification, or design artifact was edited by this planning rescope.

### Reconciled task ledger

- Total: **55 checkbox rows**.
- Implementation-owned: **52 rows** — **24 checked** historical/accepted rows through WU4B and **28 unchecked** rows from WU4C through WU10.
- Parent-owned: **3 unchecked rows**.
- Overall: **24 checked, 31 unchecked**.
- Preservation proof: the pre-existing 51 checkbox rows retain their exact text, checked state, order, and terminal ownership markers; only four new unchecked implementation-owned WU4C rows were inserted.
- Pre-rescope task artifact SHA-256 recorded by the WU4B receipt: `74e408511f32f90d845444a7a648df3e3b5dcca2893115e227c04dd1077f3569`.

### Exact next implementation row

```markdown
- [ ] **WU4C RED:** In Darwin-tagged `internal/platform/macos/metadata_darwin_test.go`, add tests first proving directories and FIFO/socket/device-like special entries report zero logical bytes with unknown allocation; a concrete adapter-owned typed path error supports both `errors.Is` against stable sentinel classes and `errors.As` for path/operation context; zero device/inode identity, zero link count, negative logical size, negative block count, non-`syscall.Stat_t`, and `statfs` failure fail closed; and a synchronized adversarial `Lstat`→symlink swap never follows or accepts target metadata; run the WU4C focused command and capture a non-zero failure caused only by the missing safety behavior. <!-- sdd-owner: implementation -->
```

Apply must resume at WU4C RED. WU5 is not authorized until all four WU4C rows have accepted evidence within the independent <=400-line guard.

---

## WU4C receipt — descriptor-backed metadata safety remediation

**Receipt predecessor:** supplied failed WU4B verification evidence `sha256:7f1d20a7416f0f9e9a3225b976fa6d9fcaf3db468ce122e6eed5913a824d37b3`. The full existing apply-progress artifact was read before this append; this receipt is appended only and does not alter the recovered history, WU4B receipt, or planning-rescope prose. No runtime token is persisted.

### Completed scope and exact API

- Completed and visibly checked: WU4C RED, GREEN, TRIANGULATE, and REFACTOR. WU5 remains unchecked; no WU5 source was created or edited.
- Darwin `MetadataAdapter.Inspect(path, requireDirectory)` performs `Lstat`, then `syscall.Open(path, O_RDONLY|O_NOFOLLOW|O_CLOEXEC|O_NONBLOCK, 0)`, descriptor `Fstat`, and descriptor `Fstatfs`; it never reads descriptor contents. `defer close(fd)` closes every successfully opened descriptor, including fstat/fstatfs/revalidation failures.
- A descriptor is accepted only when its nonzero device, inode, link count, nonnegative size/blocks, and file-type bits match the preceding Lstat; its device must match the adapter device. Type/identity changes fail closed.
- `MetadataPathError` is adapter-owned and exposes `Operation` and `Path`; its `Unwrap` retains stable sentinels for `errors.Is`, while callers can use `errors.As` for context. Raw OS errors are classified and are not exposed as domain facts. Non-Darwin retains typed `UnsupportedEnvironmentError` and now shares the concrete path-error API.
- Only `info.Mode().IsRegular()` exposes logical bytes or checked `blocks*512` allocation. Directories and the disposable FIFO expose zero logical bytes and unknown allocation. Overflow remains unknown; no logical fallback is invented.

### TDD Cycle Evidence

| Task | Layer / safety net | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- | --- |
| WU4C RED/GREEN/TRIANGULATE/REFACTOR | Darwin package unit tests; pre-change `go test ./internal/platform/macos -count=1` PASS | Added the five named WU4C tests first; focused command exited 1 with only missing `newDescriptorMetadataAdapter`, `descriptorOps`, and `MetadataPathError` symbols. | Same focused command PASS after descriptor-backed implementation. | PASS after FIFO nonblocking fixture, typed errors.Is/errors.As, zero/negative/non-Stat_t/statfs fail-closed cases, no-follow swap seam, overflow-unknown, and close-count seams. | Exact required format/focused/package/Linux/full/vet diagnostic passed. |

### Verification and fixture proof

- `gofmt -w internal/platform/macos/metadata_darwin.go internal/platform/macos/metadata_other.go internal/platform/macos/metadata_darwin_test.go && go test ./internal/platform/macos -run '^(TestDarwinMetadataRegularOnlyBytes|TestDarwinMetadataTypedErrors|TestDarwinMetadataFailClosedFields|TestDarwinMetadataNoFollowSwap|TestDarwinMetadataDescriptorLifecycle)$' -count=1 && go test ./internal/platform/macos -count=1 && GOOS=linux go test -c -o /tmp/osdycleaner-macos-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-macos-linux.test && go test ./... -count=1 && go vet ./... && test -z "$(gofmt -d internal/platform/macos/metadata_darwin.go internal/platform/macos/metadata_other.go internal/platform/macos/metadata_darwin_test.go)"` -> PASS. Focused/package/full tests passed; Linux compiled with no test files; vet and format had no diagnostics.
- RED command: `go test ./internal/platform/macos -run '^(TestDarwinMetadataRegularOnlyBytes|TestDarwinMetadataTypedErrors|TestDarwinMetadataFailClosedFields|TestDarwinMetadataNoFollowSwap|TestDarwinMetadataDescriptorLifecycle)$' -count=1` -> exit 1 before production changes, missing only the new descriptor adapter/error API.
- The disposable test fixtures use only `t.TempDir()`. The existing regular-file fixture asserts unchanged content/mode/mtime; the WU4C FIFO is created only under that temporary directory. The swap seam asserts `O_NOFOLLOW` is present and returns `ELOOP`, so no target descriptor/facts can be accepted. Close-count seams prove one close after open for success, fstat failure, and fstatfs failure. `/tmp/osdycleaner-macos-linux.test` was removed and its absence was confirmed.
- Runtime: N/A — no executable or scanner exists; the synchronized disposable descriptor fixtures are the bounded harness. No real home, external mount fixture, shell, network, telemetry, privilege, traversal, or mutation was used.

### Correction accounting, identity, and rollback

- Correction-only accounting against the supplied WU4B hashes: `metadata_darwin.go` +114/-82, `metadata_other.go` +12/-0, `metadata_darwin_test.go` +115/-57 = **241 additions + 139 deletions = 380 authored changed lines**, within the independent 400-line cap.
- Current hashes: `metadata_darwin.go` `154a481233e5f3573a04ef0686df3f8d0ba4ff4aba33bfc6dc0e9da854728c42`; `metadata_other.go` `bf0b0e3d7d271106ca9507306af952ffd96bfc932b93805d512d42946b18efba`; `metadata_darwin_test.go` `4850c143f93e919d7dc02afb6b7bdbd550a417678534ab19e572c0ac17f889d1`.
- Rollback boundary: restore exactly the WU4B hashes supplied for those three files. This removes descriptor acquisition, typed path errors, and WU4C fixtures while preserving WU4A and earlier work. WU5 must remain untouched.

### Status and remaining work

- Consumed authoritative OpenSpec status: `applyState: ready`, change `read-only-scan-foundation`, repo-local workspace and allowed root `/Users/osdy/Documents/GitHub/OsdyCleaner`, strict TDD active, no blocked reasons. Delivery decision was the user-authorized sequential stacked-to-main WU4C slice.
- Parent-owned lifecycle rows remain deferred and byte-for-byte unchanged. The persisted task artifact was re-read after updates: all four WU4C rows are `[x]`, but the next WU5 RED row is `[ ]`.
- Workload / PR boundary: WU4C only, independently measured at 380 authored changes, sequential stacked-to-main. No commit or PR was created. The implementation is ready for independent verification; route lifecycle ownership to the parent. Do not start WU5 in this delegated scope.

---

## Planning rescope — authorized adjacent WU4D component-resolution remediation

**Append provenance:** The full existing apply-progress artifact was read before this append. The user supplied the abbreviated current baselines `tasks.md` `sha256:0ddfe277...` and apply-progress/WU4C failed evidence `sha256:5d38e230...`; these prefixes are recorded exactly as supplied and no unprovided suffix is invented. This planning-only note is appended after all prior content and records no implementation acceptance.

### Superseded historical evidence

- WU4C failed evidence `sha256:5d38e230...` is historical superseded evidence, not accepted evidence.
- WU4C's non-reproducible predecessor and baseline claims are historical superseded claims and must not authorize WU5.
- WU4B and WU4C checkbox rows remain checked solely as preserved attempted-implementation history. No current checkbox text, state, ownership marker, or relative order was changed.

### Authorized plan change

- Four unchecked implementation-owned WU4D rows were inserted immediately after WU4C and before WU5 in strict RED → GREEN → TRIANGULATE → REFACTOR order.
- WU4D is an independent forecast of 290–395 authored changed lines with a mandatory maximum of 400 additions plus deletions; it does not merge accounting with WU4C.
- WU4D replaces unsafe pathname/final-component-only acquisition with component-wise Darwin `openat`-style descriptor-relative resolution anchored at a trusted absolute root descriptor. It validates components and safe flags, closes every intermediate/final descriptor, accepts facts only from final-descriptor `Fstat`/`Fstatfs`, and treats `Lstat`/pathname facts as initial-failure classification only.
- WU4D verification covers stable concrete `MetadataPathError` `errors.Is`/`errors.As` classes; ancestor/final symlink swaps; identity/type and applicable same-inode hard-link redirects; flags; zero/negative/overflow and special-type fields; close errors/counts; intermediate descriptors; and Linux unsupported typing.
- WU4D explicitly excludes leaked descriptors, blocking FIFO/device access, path-based `Statfs`, content reads, traversal/walk, and mutation.
- WU5 now depends on accepted WU4D; checked WU4B/WU4C rows and failed WU4C evidence cannot satisfy that dependency.
- Forecast totals are now 3,260–4,200 additions, 10–140 deletions, and 3,270–4,340 authored changed lines, excluding generated `go.sum`.
- No product code, proposal, specification, or design artifact was edited by this planning rescope.

### Reconciled task ledger and preservation proof

- Total: **59 checkbox rows**.
- Implementation-owned: **56 rows** — **28 checked** pre-existing rows through WU4C and **28 unchecked** rows from WU4D through WU10.
- Parent-owned: **3 unchecked rows**.
- Overall: **28 checked, 31 unchecked**.
- Preservation proof: all 55 pre-existing checkbox rows retain their exact text, checked state, relative order, and terminal ownership markers; the only checkbox additions are the four new unchecked WU4D implementation rows between WU4C and WU5.
- Baseline identity recorded exactly as supplied: task artifact `sha256:0ddfe277...`; receipt/WU4C failed evidence `sha256:5d38e230...`. Full suffixes and post-edit hashes were not available through the injected file tools and are not fabricated.

### Exact next implementation row

```markdown
- [ ] **WU4D RED:** In `internal/platform/macos/metadata_darwin_test.go` and `metadata_other_test.go`, add failing tests first for trusted-root component-wise resolution; rejection of `.`, `..`, empty/repeated/trailing suffix components; required intermediate/final `O_NOFOLLOW|O_CLOEXEC|O_NONBLOCK` and appropriate `O_DIRECTORY` flags; stable concrete `MetadataPathError` `errors.Is`/`errors.As` classes for missing, symlink, non-directory, device mismatch, non-local, inaccessible, changed-between-observations, and invalid metadata; final-descriptor-only accepted facts; and Linux unsupported `errors.Is`/`errors.As`; run the focused command and capture a non-zero RED caused only by the missing WU4D behavior. <!-- sdd-owner: implementation -->
```

Apply must resume at WU4D RED. WU5 is not authorized until all four WU4D rows have accepted evidence within the independent 400-line guard.

---

## WU4D blocked-for-split receipt — component-resolution remediation

**Receipt predecessor:** exact pre-edit backup `/tmp/osdy-wu4d-baseline-1788066883-49900/apply-progress.md`, SHA-256 `9ed808e0259b062bab414706ef77d56555141426c0517321a41efd65e4fae6df`. The predecessor compared byte-for-byte equal immediately before this append. This is an append-only blocked receipt; no runtime token is persisted.

### Status, scope, and integrity

- Consumed authoritative native OpenSpec status: `read-only-scan-foundation`, `applyState: ready`, repo-local workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, allowed root the workspace, strict TDD enabled, and no blocked reasons. The supplied sequential `stacked-to-main` WU4D delivery path resolved the high-risk workload gate.
- Before any edit, exact copies of this receipt and the three WU4C metadata baseline files were created under `/tmp/osdy-wu4d-baseline-1788066883-49900/`; `SHA256SUMS` records their hashes. The baseline product hashes are `metadata_darwin.go` `154a481233e5f3573a04ef0686df3f8d0ba4ff4aba33bfc6dc0e9da854728c42`, `metadata_other.go` `bf0b0e3d7d271106ca9507306af952ffd96bfc932b93805d512d42946b18efba`, and `metadata_darwin_test.go` `4850c143f93e919d7dc02afb6b7bdbd550a417678534ab19e572c0ac17f889d1`.
- A test-first candidate was discarded, not accepted: its focused Darwin command passed and its Linux cross-compile passed, but exact reconstructed-baseline accounting was `226 additions + 246 deletions + 25 additions in new metadata_other_test.go = 497 authored changed lines`. This exceeds WU4D's hard 400-line limit. The candidate product/test files were restored byte-for-byte from the saved baseline and `metadata_other_test.go` was removed.
- Post-restore hashes equal the three baseline hashes above. No WU4D task checkbox was changed; WU5 remains untouched and unchecked. Parent-owned lifecycle rows were deferred unchanged.

### TDD evidence for discarded candidate

| Stage | Evidence |
| --- | --- |
| Safety net | Existing Darwin package tests were readable and the baseline had no pending source diff against the saved copy. |
| RED | `go test ./internal/platform/macos -run '^(TestDarwinMetadataComponentResolution | TestDarwinMetadataAdversarialResolution | TestDarwinMetadataAcceptedFacts | TestDarwinMetadataErrorClasses | TestDarwinMetadataDescriptorLifecycle)$' -count=1` exited 1 only for new component resolver symbols and descriptor seams. Linux unsupported test RED exited 1 only for `ErrMetadataInvalid`. |
| GREEN | The discarded candidate's same focused Darwin command passed; `GOOS=linux go test -c -o /tmp/osdycleaner-macos-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-macos-linux.test` passed. |
| TRIANGULATE / REFACTOR | Not accepted or continued after the independent accounting measurement exceeded the hard guard. |

### Required split and remaining task rows

`blocked_for_split`: authorize adjacent RED/GREEN-preserving WU4D slices before another product edit. The exact unchecked WU4D rows remain:

- [ ] **WU4D RED:** In `internal/platform/macos/metadata_darwin_test.go` and `metadata_other_test.go`, add failing tests first for trusted-root component-wise resolution; rejection of `.`, `..`, empty/repeated/trailing suffix components; required intermediate/final `O_NOFOLLOW|O_CLOEXEC|O_NONBLOCK` and appropriate `O_DIRECTORY` flags; stable concrete `MetadataPathError` `errors.Is`/`errors.As` classes for missing, symlink, non-directory, device mismatch, non-local, inaccessible, changed-between-observations, and invalid metadata; final-descriptor-only accepted facts; and Linux unsupported `errors.Is`/`errors.As`; run the focused command and capture a non-zero RED caused only by the missing WU4D behavior. <!-- sdd-owner: implementation -->
- [ ] **WU4D GREEN:** In the allowed metadata files, replace pathname/final-component-only acquisition with the minimum trusted-root descriptor-relative resolver: open each validated component with Darwin `openat`-style APIs and safe flags, close every intermediate descriptor, obtain mode/device/inode/link/size/blocks only from final-descriptor `Fstat` and locality only from final-descriptor `Fstatfs`, use `Lstat`/pathname facts solely to classify initial failures, reject device/non-local/type/identity/size/block invariants fail closed, and return stable concrete classified errors without content reads, path-based `Statfs`, traversal, mutation, leaked descriptors, or blocking FIFO/device opens; rerun the focused command and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU4D TRIANGULATE:** Extend the disposable synchronized seams with ancestor and final-component symlink swaps, identity/type changes, an applicable same-inode hard-link redirect case, zero/negative/overflow device/inode/link/size/block fields, directories/FIFO/socket/device-like special types, every intermediate/final close count, and open/fstat/fstatfs/close failures; prove a close failure after otherwise successful acquisition returns inaccessible with no accepted `Metadata`, an earlier primary failure remains primary while close is still attempted, no target facts are accepted, and Linux unsupported construction preserves `errors.Is`/`errors.As`; rerun the focused command plus `GOOS=linux go test -c -o /tmp/osdycleaner-macos-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-macos-linux.test` and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU4D REFACTOR:** Run `gofmt -w internal/platform/macos/*.go && go test ./internal/platform/macos -run '^(TestDarwinMetadataComponentResolution|TestDarwinMetadataAdversarialResolution|TestDarwinMetadataAcceptedFacts|TestDarwinMetadataErrorClasses|TestDarwinMetadataDescriptorLifecycle)$' -count=1 && go test ./internal/platform/macos -count=1 && GOOS=linux go test -c -o /tmp/osdycleaner-macos-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-macos-linux.test && go test ./... -count=1 && go vet ./... && test -z "$(gofmt -d internal/platform/macos)"`; expect focused/platform/Linux/full/vet/format PASS, source checks finding no path-based `Statfs`, content reads, traversal/walk, mutation, unsafe final-only open, or descriptor leak/blocking path, and record exact commands/results, runtime N/A, file hashes, rollback boundary, and WU4D additions plus deletions at no more than 400 before authorizing WU5. <!-- sdd-owner: implementation -->

- Workload / PR boundary: no accepted WU4D slice exists. Runtime is N/A because no executable or scanner boundary exists; disposable candidate fixtures only were used and are removed. WU5 is not authorized.

---

## Planning rescope — user-authorized WU4D1 and WU4D2 split

This is a planning-only append. No product, proposal, specification, or design file was edited, no implementation was accepted, and no checkbox was checked.

- The original blocked WU4D receipt and its four unchecked rows are superseded by the authorized split; they remain historical blocked evidence and are not accepted.
- The 497-authored-line candidate remains fully restored to the WU4C baseline. Recorded baseline product hashes remain `metadata_darwin.go` `154a481233e5f3573a04ef0686df3f8d0ba4ff4aba33bfc6dc0e9da854728c42`, `metadata_other.go` `bf0b0e3d7d271106ca9507306af952ffd96bfc932b93805d512d42946b18efba`, and `metadata_darwin_test.go` `4850c143f93e919d7dc02afb6b7bdbd550a417678534ab19e572c0ac17f889d1`.
- `/tmp/osdy-wu4d-baseline-1788066883-49900/` must remain until WU4D1 and WU4D2 are each independently verified. Its recorded apply-progress predecessor hash is `9ed808e0259b062bab414706ef77d56555141426c0517321a41efd65e4fae6df`.
- `tasks.md` replaces only the four unchecked WU4D checkbox rows with eight unchecked implementation-owned rows: WU4D1 RED/GREEN/TRIANGULATE/REFACTOR followed by WU4D2 RED/GREEN/TRIANGULATE/REFACTOR. Every other checkbox row retains its prior text, state, order, and terminal ownership marker.
- The reconciled ledger is 63 rows: 60 implementation-owned and 3 parent-owned; 28 checked and 35 unchecked overall. WU4D1 and WU4D2 are independently capped at 400 authored additions plus deletions.
- WU4D1 must finish compiling and passing as a safety-improved component-wise resolver and accepted-facts boundary. WU4D2 depends on accepted WU4D1 and completes concrete typed-error and descriptor-lifecycle hardening. WU5 depends on accepted WU4D2 and remains unstarted.
- Forecast totals are 3,370–4,380 additions, 15–160 deletions, and 3,385–4,540 authored changed lines, excluding generated `go.sum`.
- Current post-edit artifact hashes are unavailable through the injected file-only tools and are not fabricated.

### Exact next implementation row

```markdown
- [ ] **WU4D1 RED:** In `internal/platform/macos/metadata_darwin_test.go`, add failing tests first for clean absolute component parsing and rejection of `.`, `..`, empty/repeated/trailing suffix components; exact root/intermediate/final `O_NOFOLLOW|O_CLOEXEC|O_NONBLOCK` flags with `O_DIRECTORY` only where required; root/intermediate/final descriptor open/close counts; ancestor and final symlink no-follow; final-descriptor-only `Fstat`/`Fstatfs` accepted facts; regular-file-only bytes; type/identity/device/local revalidation; and the minimum stable missing/symlink/non-directory/inaccessible resolver error classes; run the WU4D1 focused command and capture a non-zero RED caused only by missing WU4D1 behavior. <!-- sdd-owner: implementation -->
```

Apply must resume at WU4D1 RED. WU4D2 and WU5 are not authorized until their stated predecessor has independent accepted evidence within the 400-line guard.

---

## WU4D1 receipt — component-wise resolver and accepted descriptor facts

**Receipt predecessor / prefix proof:** this append began from the WU4D1 baseline copy `/tmp/osdy-wu4d1-baseline-Nqk9pbi1/openspec/changes/read-only-scan-foundation/apply-progress.md`, SHA-256 `0b3f05f16631adb30afd4d7c15dca3eb619252895e3656fd25896ba4d43f5190`; that copy is retained for independent verification. The original required baseline remains at `/tmp/osdy-wu4d-baseline-1788066883-49900/`, with the recorded WU4C metadata hashes unchanged at WU4D1 start. This receipt is append-only and no runtime token is persisted.

### Completed scope and exact API

- Completed and visibly checked: WU4D1 RED, GREEN, TRIANGULATE, and REFACTOR only. WU4D2, WU5, and later rows remain unchecked.
- `MetadataAdapter.Inspect(path, requireDirectory)` now accepts only a clean non-root absolute path split into nonempty, non-`.`/`..` components. It does not use `filepath.Clean` to accept input.
- It opens trusted `/`, every intermediate, and the final component descriptor-relatively using `O_RDONLY|O_NOFOLLOW|O_CLOEXEC|O_NONBLOCK`; `O_DIRECTORY` is applied to root/intermediates and to a directory-required final component. `golang.org/x/sys/unix.Openat` provides the Darwin openat binding.
- Every acquired descriptor is closed in reverse chain order on the exercised success, open, fstat, and fstatfs paths. Close-error precedence and adversarial lifecycle hardening remain explicitly deferred to WU4D2.
- Only final-descriptor `Fstat` and `Fstatfs` form accepted mode/device/identity/link/size/allocation/locality facts. Regular files alone expose bytes; directories and FIFO-like special entries expose zero logical bytes and unknown allocation. No path-based `Statfs`, content read, walk, mutation, or accepted Lstat fact exists.

### TDD Cycle Evidence

| Task | Test file/layer | Safety net | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- | --- | --- |
| WU4D1 RED/GREEN/TRIANGULATE/REFACTOR | `internal/platform/macos/metadata_darwin_test.go` / Darwin descriptor unit seams and disposable fixtures | `go test ./internal/platform/macos -count=1` PASS before edits | Exact WU4D1 focused command exited 1 with only missing `parseCleanAbsolutePath`, `componentOps`, and `newComponentMetadataAdapter` symbols. | Exact focused command PASS after the minimum clean component resolver and final-descriptor acquisition. | PASS after ancestor/final no-follow seams, exact flags, root/intermediate/final close count, regular/directory/FIFO cases, and open/fstat/fstatfs error rows. | `gofmt`, focused/package/full/vet/format all PASS; no behavior expansion. |

### Verification, descriptor proof, and diagnostics

- Exact refactor command: `gofmt -w internal/platform/macos/metadata_darwin.go internal/platform/macos/metadata_darwin_test.go internal/platform/macos/metadata_other.go && go test ./internal/platform/macos -run '^(TestDarwinMetadataComponentParsing|TestDarwinMetadataOpenatFlags|TestDarwinMetadataDescriptorChain|TestDarwinMetadataNoFollowComponents|TestDarwinMetadataAcceptedFacts|TestDarwinMetadataRegularOnlyBytes|TestDarwinMetadataRevalidation|TestDarwinMetadataResolverErrors)$' -count=1 && go test ./internal/platform/macos -count=1 && go test ./... -count=1 && go vet ./... && test -z "$(gofmt -d internal/platform/macos)"` -> PASS.
- Focused tests prove lexical rejection, exact root/intermediate/final flags, descriptor-chain open/close counts, no-follow classification for ancestor/final seam failures, final descriptor accepted facts, regular-only bytes, device revalidation, and minimum resolver errors. `TestDarwinMetadataRegularOnlyBytes` uses only physical `t.TempDir()` fixture paths to avoid macOS `/var` aliasing; it proves special types cannot block or report bytes.
- `gopls` was unavailable; Go compiler/package tests plus `go vet ./...` were clean. Runtime evidence: N/A — this is the designed disposable descriptor harness; no scanner or executable exists.
- Source scan found no pathname `Statfs`, `Lstat`, `os.Stat`, content read, walk, symlink resolution, or mutation in Darwin production metadata code.

### Accounting, hashes, and rollback

- Reproducible product accounting relative to `/tmp/osdy-wu4d1-baseline-Nqk9pbi1`: `metadata_darwin.go` +81/-34, `metadata_darwin_test.go` +98/-94, `metadata_other.go` +0/-0, `go.mod` +5/-0; `go.sum` is generated and excluded. **184 additions + 128 deletions = 312 authored changed lines**, within the independent 400-line cap.
- Final hashes: `metadata_darwin.go` `4cf7dcda259cf88217d3ac49871cb6bddae7939fdd6ed4fbacbf6fa9f1bdd221`; `metadata_other.go` `bf0b0e3d7d271106ca9507306af952ffd96bfc932b93805d512d42946b18efba`; `metadata_darwin_test.go` `46684b27a940083a64a36978cbac21c7b8048c5cf72a2c3f7d7f97670116afeb`; `go.mod` `438c539e796cc69735ce3dab3340e2b39944cd8f69c5b2069f7f60111f1d0330`; generated `go.sum` `50525b09f4448026339b45ad88801ff0fae28a671b620e7e3cca358e7703941f`.
- Rollback: restore the three metadata files and apply-progress copy from `/tmp/osdy-wu4d1-baseline-Nqk9pbi1/`, restore `go.mod` and remove generated `go.sum` if reverting this new dependency, then retain both `/tmp` baselines until WU4D2 independent verification. No commit or PR was created.

### Status and remaining work

- Consumed authoritative native status: `applyState: ready`, `artifactStore: openspec`, change `read-only-scan-foundation`, repo-local workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, allowed root the workspace, strict TDD active, and no blocker. The active-attempt continuation was acquired using the parent-bound token; it is intentionally not recorded here.
- Workload/PR boundary: WU4D1 only, sequential `stacked-to-main`, 312 authored changed lines. WU4D2 is required before WU5 and remains deferred.
- Persisted task reconciliation follows; each listed completed WU4D1 row is `[x]`, and all remaining rows are `[ ]`.

#### Exact remaining unchecked rows

- [ ] **WU4D2 RED:** In `internal/platform/macos/metadata_darwin_test.go` and `metadata_other_test.go`, add failing concrete `errors.Is`/`errors.As` cases for all eight metadata classes; zero device/inode/link count, negative size/blocks, allocation overflow, and non-`syscall.Stat_t`; close failure as primary only after otherwise successful acquisition while an earlier open/fstat/fstatfs/revalidation failure remains primary; exact close counts for every adversarial root/intermediate/final path; and Linux unsupported `errors.Is`/`errors.As`; run the WU4D2 focused command and capture a non-zero RED caused only by missing WU4D2 behavior. <!-- sdd-owner: implementation -->
- [ ] **WU4D2 GREEN:** Implement the minimum remaining concrete `MetadataPathError` classification, fail-closed numeric/non-Stat validation, deferred-close error capture with primary-error precedence, and typed non-Darwin unsupported behavior while preserving WU4D1 resolver flags, accepted-fact rules, and all close paths; rerun the WU4D2 focused command and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU4D2 TRIANGULATE:** Add combined open/fstat/fstatfs/revalidation/close failures, all root/intermediate/final close-count permutations, repeated close-error and invalid-field cases, and Linux construction assertions; prove no close error masks an earlier primary failure, successful acquisition plus close failure returns inaccessible with no accepted `Metadata`, unsupported errors retain both `errors.Is` and `errors.As`, and no descriptor or target fact escapes; rerun the focused command plus `GOOS=linux go test -c -o /tmp/osdycleaner-macos-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-macos-linux.test` and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU4D2 REFACTOR:** Run `gofmt -w internal/platform/macos/*.go && go test ./internal/platform/macos -run '^(TestDarwinMetadataErrorClasses|TestDarwinMetadataInvalidFacts|TestDarwinMetadataCloseFailurePrecedence|TestDarwinMetadataAdversarialLifecycle|TestUnsupportedEnvironmentErrors)$' -count=1 && go test ./internal/platform/macos -count=1 && GOOS=linux go test -c -o /tmp/osdycleaner-macos-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-macos-linux.test && go test ./... -count=1 && go vet ./... && test -z "$(gofmt -d internal/platform/macos)"`; expect focused/platform/Linux/full/vet/format PASS, record exact commands/results, runtime N/A, accepted-WU4D1 predecessor hashes, final hashes, rollback boundary, and WU4D2 additions plus deletions at no more than 400; retain the baseline reproduction directory through this independent verification, then authorize WU5. <!-- sdd-owner: implementation -->
- [ ] **WU5 RED:** Add table-driven temporary-home and fake-metadata tests for normal and empty roots, every-root-missing, inaccessible root/entry, non-directory root, root/entry symlinks, lexical escape callbacks, different-device root/entry, changed entries, unsupported special types, sibling continuation, serial root order, bounded goroutine/queue use, and before/after fixture metadata; run the focused command and record missing-scanner failures. <!-- sdd-owner: implementation -->
- [ ] **WU5 GREEN:** Implement `filepath.WalkDir` orchestration, preflight, defensive containment, synchronous directory boundary checks, fixed worker/job/result bounds, backpressure, `lstat`-only regular-file jobs, safe stop/drain/join mechanics, root-scoped warning conversion, and finalizer integration; rerun and expect PASS with no target or other-device bytes. <!-- sdd-owner: implementation -->
- [ ] **WU5 TRIANGULATE:** Add completion-order permutations and combined entry-error/symlink/boundary fixtures proving safe siblings and later roots continue, status remains boundary-limited unless a higher gap applies, special/changed entries add no bytes, and all five observations finalize; rerun and expect PASS without inspecting the production home. <!-- sdd-owner: implementation -->
- [ ] **WU5 REFACTOR:** Run `gofmt -w internal/scan/scanner.go internal/scan/scanner_test.go internal/scan/finalize.go internal/scan/finalize_test.go && go test ./internal/scan -count=1`; expect PASS, bounded goroutine ownership, and a coherent non-runnable scanner package before advancing to WU6. <!-- sdd-owner: implementation -->
- [ ] **WU6 RED:** Add deterministic channel-synchronized tests for cancellation before a root, during an active metadata call, after an earlier root, and alongside partial/boundary warnings; add entry/path limit tests and run the focused command to record failures for missing stop/state behavior without sleep-based timing. <!-- sdd-owner: implementation -->
- [ ] **WU6 GREEN:** Implement context checks before new work, immediate enqueue stop, bounded active-call completion, channel drain/join, cancellation warnings, skipped later roots, limit warnings, scan-policy fields, and exact `cancelled > partial > boundary_limited > scanned` derivation; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU6 TRIANGULATE:** Add reordered worker completion, blocked-result-channel, unknown-required-metadata, combined limit/boundary/error, and repeated-cancel cases; run `go test -race ./internal/scan -run '^(TestScannerStatusPrecedence|TestScannerEntryLimit|TestScannerPathBudget|TestScannerCancellation|TestScannerNoLeakedWork|TestScannerCompletionOrderStable)$' -count=1` and expect PASS with no race, leaked work, discarded finalized fact, or newly started root. <!-- sdd-owner: implementation -->
- [ ] **WU6 REFACTOR:** Run `gofmt -w internal/scan/scanner.go internal/scan/scanner_test.go internal/scan/finalize.go internal/scan/finalize_test.go && go test -race ./internal/scan -count=1`; expect PASS and verify partial/cancelled aggregate estimates remain incomplete. <!-- sdd-owner: implementation -->
- [ ] **WU7 RED:** Add fixed-snapshot report tests for complete, partial, and cancelled outcomes, exact top-level/nested field order, typed bytes/booleans, explicit unknowns, `related_path: null`, `[]` collections, one trailing newline, all five roots, and text/JSON fact equivalence; run the focused command and record missing-renderer/golden failures. <!-- sdd-owner: implementation -->
- [ ] **WU7 GREEN:** Implement ordered DTO projection in `internal/report/json.go` and canonical text rendering in `internal/report/text.go`, then run `go test ./internal/report -run '^(TestJSONReport|TestTextReport|TestReportDeterminism|TestReportEquivalence)$' -update -count=1`; inspect all six allowed goldens and rerun the focused command without `-update`, expecting PASS. <!-- sdd-owner: implementation -->
- [ ] **WU7 TRIANGULATE:** Add discovery/collection permutations and assertions for byte-identical JSON, JSON parseability, schema/outcome/completeness, risks, bases, warnings, prominent partial/cancelled text, APFS clone/snapshot/compression/hard-link caveats, CoreSimulator manual-review wording, and no deletion-safety/reclaim promise; rerun without `-update` and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU7 REFACTOR:** Run `gofmt -w internal/report/json.go internal/report/text.go internal/report/report_test.go && go test ./internal/report -count=1`; expect PASS without golden updates, map-backed serialization, absolute home paths, wall-clock/host/user/random fields, or presentation-derived scan facts. <!-- sdd-owner: implementation -->
- [ ] **WU8 RED:** Add direct model tests for pure `Init`, canonical category movement, h/j/k/l and arrows, page scrolling, Tab details, resize, q/Escape/Ctrl-C quit, complete/partial/cancelled views, and terminal rejection; run the focused command and record missing-model/dependency failures. <!-- sdd-owner: implementation -->
- [ ] **WU8 GREEN:** Pin approved compatible v2 dependency versions in `go.mod`, implement snapshot-only model/update/view/runner files, use only the viewport component, and rerun the focused command; expect PASS with no scan command from `Init` and no command after quit. <!-- sdd-owner: implementation -->
- [ ] **WU8 TRIANGULATE:** Assert prominent “Read-only scan,” “Manual review required,” and “Estimate—not guaranteed reclaimable space” wording, cancellation/incompleteness before facts, canonical area/warning agreement, and absence of cleanup/select/action/confirm/delete/remove/Trash affordances while allowing the required estimate disclaimer; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU8 REFACTOR:** Run `gofmt -w internal/tui/model.go internal/tui/view.go internal/tui/run.go internal/tui/model_test.go && go mod tidy && go test ./internal/tui -count=1`; expect PASS, inspect generated `go.sum` in the complete receipt, and verify the model stores no operation intent or alternate totals. <!-- sdd-owner: implementation -->
- [ ] **WU9 RED:** Add injected scanner/renderer/viewer/terminal/stream tests for default text, formats, positional or path-like input, unsupported format, noninteractive TUI, one scan call, complete/partial/cancelled outcomes, unsupported environment, global init/finalize/render/write/viewer failures, and fixture JSON; run the focused command and record missing-command/orchestration failures. <!-- sdd-owner: implementation -->
- [ ] **WU9 GREEN:** Pin Cobra in `go.mod`, implement silenced thin command validation and one-scan orchestration in `internal/cli`, render text/JSON fully before stdout, keep snapshot warnings in-band and diagnostics on stderr, and centralize codes 0/1/2/3/4/130; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU9 TRIANGULATE:** Add poison-scanner cases proving invalid path/format/noninteractive TUI fail before construction or traversal, root warnings never become code 1, cancellation outranks code 3, global failure emits no purported snapshot, JSON has one document/newline and no decoration, write/viewer failure never rescans, and signal cancellation starts no follow-up operation; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU9 REFACTOR:** Run `gofmt -w internal/cli/command.go internal/cli/run.go internal/cli/command_test.go && go mod tidy && go test ./internal/cli -count=1`; expect PASS, execute the stated runtime harness evidence command, and account for generated `go.sum` without counting its lines as authored. <!-- sdd-owner: implementation -->
- [ ] **WU10 RED:** Before creating `cmd/osdy/main.go`, run `go test ./cmd/osdy -count=1`; record the expected non-zero “directory/package not found” result as the missing production entry-point boundary. <!-- sdd-owner: implementation -->
- [ ] **WU10 GREEN:** Implement `cmd/osdy/main.go` with `signal.NotifyContext`, stop signal delivery, production dependency construction, delegation to `internal/cli`, and `os.Exit` only; run `gofmt -w cmd/osdy/main.go && go test ./cmd/osdy ./internal/cli -count=1` and expect PASS with no scan/report logic in `main`. <!-- sdd-owner: implementation -->
- [ ] **WU10 TRIANGULATE:** Run the disposable runtime script below and expect exit 0, empty stderr, exactly one parseable JSON document, schema 1, complete outcome, canonical five area IDs, no absolute temporary-home path in JSON, and identical before/after fixture metadata. <!-- sdd-owner: implementation -->
- [ ] **WU10 REFACTOR:** Run the final acceptance sequence below in order: explicit-file `gofmt`, all focused package tests, report JSON parseability/equivalence tests without `-update`, `go test ./...`, `go vet ./...`, then the disposable runtime script again; expect every command to pass/no-diagnostic and stop rather than weakening checks if any result fails. <!-- sdd-owner: implementation -->
- [ ] Before resumed apply, record the already selected sequential `stacked-to-main` chain strategy and the WU2A → WU2B split, confirm no further product/delivery choice is pending, retain the previously selected module import identity, and preserve one-writer execution without assuming commits or PRs can be created. <!-- sdd-owner: parent -->
- [ ] After each applied work unit, inspect its apply-progress receipt, confirm authored additions plus deletions are at most 400 with generated `go.sum` excluded, confirm the focused command/result, runtime evidence or explicit N/A, complete changed-file identity, and rollback boundary, then authorize the next dependent unit through SDD apply/verify authority; if a unit exceeds 400 authored lines, stop and split it again before further implementation writes. <!-- sdd-owner: parent -->
- [ ] After WU10, record final SDD completion evidence against the normative coverage matrix, required commands, disposable runtime script, authored-line accounting, complete changed-file identity, and rollback boundaries; mark implementation complete only when that evidence is satisfied. <!-- sdd-owner: parent -->

---

## WU4D2 receipt — typed-error and descriptor lifecycle hardening

**Receipt predecessor / prefix proof:** Before this append, the complete receipt was re-read and byte-for-byte matched the retained WU4D2 baseline copy at `/tmp/osdy-wu4d2-baseline-1788068463-3237/openspec/changes/read-only-scan-foundation/apply-progress.md`; both SHA-256 values were `9eb6af5ca3222f6d94617d9322d214b1df77ad9fafc1a676ee7cc385e6c000a0`. This append does not replace prior receipt text. Runtime attempt tokens are deliberately omitted.

### Completed scope and API/error semantics

- Completed and visibly checked in `tasks.md`: WU4D2 RED, GREEN, TRIANGULATE, and REFACTOR only. WU5 and all later implementation rows remain unchecked; parent-owned lifecycle rows remain unchanged.
- Added stable `ErrMetadataChanged` and `ErrMetadataInvalid` sentinels on Darwin and non-Darwin. All eight classes (missing, symlink, non-directory, device mismatch, non-local, inaccessible, changed, invalid) return adapter-owned `*MetadataPathError` values with nonempty operation and input path; `errors.Is` reaches only the stable class and `errors.As` reaches the context type.
- Zero device/inode/link fields, negative size/blocks, allocation overflow, and an injected non-`*syscall.Stat_t` fact type fail closed as invalid metadata. Overflow no longer becomes an unknown allocation after otherwise accepted acquisition.
- Descriptor closure is deferred across the complete acquired root/intermediate/final chain. A close failure after otherwise valid acquisition returns zero `Metadata` plus inaccessible; open, fstat, fstatfs, and revalidation primary errors remain their original class while every acquired descriptor is still closed exactly once.
- Non-Darwin construction and `Inspect` return `UnsupportedEnvironmentError` without filesystem access; it supports both `errors.Is(err, ErrUnsupportedEnvironment)` and `errors.As`.

### TDD Cycle Evidence

| Task | Test file/layer | Safety net | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- | --- | --- |
| WU4D2 RED/GREEN/TRIANGULATE/REFACTOR | `metadata_darwin_test.go` Darwin descriptor unit seams; `metadata_other_test.go` non-Darwin unit | `go test ./internal/platform/macos -count=1` PASS before edits. | Focused RED exited 1 solely because `ErrMetadataChanged`, `ErrMetadataInvalid`, and the revalidation seam were absent. | Exact focused command PASS after stable classes, invalid-fact rejection, deferred close handling, and typed non-Darwin test. | PASS after fstatfs/revalidation plus close-error combinations, repeated close failure, root/intermediate/final count permutations, overflow, and non-Stat assertions; Linux compile PASS. | `gofmt`, focused/package/full/vet/source checks PASS with no behavior expansion. |

### Verification and lifecycle proof

- `go test ./internal/platform/macos -run '^(TestDarwinMetadataErrorClasses|TestDarwinMetadataInvalidFacts|TestDarwinMetadataCloseFailurePrecedence|TestDarwinMetadataAdversarialLifecycle|TestUnsupportedEnvironmentErrors)$' -count=1` -> PASS.
- `go test ./internal/platform/macos -run '^TestDarwinMetadataLifecycleCombinations$' -count=1` -> PASS.
- `go test ./internal/platform/macos -count=1`, `GOOS=linux go test -c -o /tmp/osdycleaner-macos-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-macos-linux.test`, `go test ./... -count=1`, and `go vet ./...` -> PASS. The Linux temporary artifact was removed and confirmed absent.
- `gofmt -d internal/platform/macos` produced no output. Source scan found no pathname `Statfs`, `Lstat`, `os.Stat`, content read, walk, mutation, or symlink-resolution calls in Darwin production metadata source. `gopls` was not available.
- Lifecycle fixtures prove: root/intermediate/final chains close in reverse acquisition order; an intermediate open failure closes exactly acquired descriptors; fstat/fstatfs/revalidation primaries survive close failures; and otherwise-successful close failure returns no facts. Runtime N/A: disposable descriptor fixtures are the bounded harness; no scanner or executable exists.

### Accounting, baselines, hashes, rollback, and status

- New retained independent baseline: `/tmp/osdy-wu4d2-baseline-1788068463-3237`, containing metadata sources/tests, `go.mod`, `go.sum`, `tasks.md`, this receipt's predecessor, and `SHA256SUMS`. The required older baselines `/tmp/osdy-wu4d-baseline-1788066883-49900/` and `/tmp/osdy-wu4d1-baseline-Nqk9pbi1/` were retained.
- Relative to that WU4D2 baseline: `metadata_darwin.go` +33/-12, `metadata_other.go` +2/-0, `metadata_darwin_test.go` +151/-0, `metadata_other_test.go` +23/-0: **209 additions + 12 deletions = 221 authored changed lines**, below the 400-line hard cap. `go.mod`/`go.sum` are unchanged.
- Final hashes: `metadata_darwin.go` `47f6b087cda8eb5f89e8a8d386573d46aea941f978b166f46778f2b6b86b3dfc`; `metadata_other.go` `81e4a9cb7469076a4cfb17eddea6dead533938bc08428954c600ab3a5f591ad0`; `metadata_darwin_test.go` `77237653df8f6e02b515baab07a94b277e85316d3e01349ff69c6e720d81aea8`; `metadata_other_test.go` `85c4721901e7c95cb56aa35a14b31e16197cedfa0b1478f9263ac8de66892fab`. Ordered product-manifest evidence revision: `sha256:d39b74730d6c4174e383edecf59e56b39209f103d58976842ae07f6b65963f9b`.
- Rollback boundary: restore the three WU4D1 files and apply-progress from `/tmp/osdy-wu4d1-baseline-Nqk9pbi1/`, remove `metadata_other_test.go`, and retain all three baseline directories through independent verification. This removes only WU4D2 classes/lifecycle hardening and retains WU4D1's component resolver.
- Consumed authoritative status: `artifactStore: openspec`, `applyState: ready`, change `read-only-scan-foundation`, repo-local workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, allowed root the workspace, strict TDD active, and no action-context warnings. Delivery was the user-authorized `stacked-to-main` WU4D2 work-unit slice; no commit, PR, review, or WU5 work was created.

---

## Planning correction — reachable metadata errors and scanner-owned entry changes

**Append provenance:** The complete receipt was read before this append. The user-supplied current receipt baseline is SHA-256 `311412225ececf023de7f3f87de0072636b30669fad4a7915186cb647554e0d9`. This planning-only note appends to history; it accepts no implementation and edits no product, proposal, specification, or design file. The available file tools do not compute SHA-256, so no current `tasks.md` hash is fabricated.

### Independent D2 verification result

- **FAIL:** WU4D2 introduced artificial `statType` and `revalidate` seams, impossible non-`syscall.Stat_t` runtime simulation, and adapter-level `ErrMetadataChanged` behavior with no production trigger.
- Reachable metadata behavior remains required: missing, symlink, non-directory, device mismatch, non-local filesystem, inaccessible, invalid numeric facts, descriptor closure and close-failure precedence, plus typed unsupported environment.
- Accepted final-descriptor facts remain stable after pathname replacement because the opened descriptor stays bound to its object. The adapter has no enumeration baseline and therefore cannot honestly own `entry_changed`.
- WU5 owns concurrent-change classification by comparing saved enumeration kind/identity with later descriptor-backed metadata; changed kind/identity, disappearance, or inconsistent symlink/non-directory state contributes zero, emits `entry_changed`, yields `partial/entry_visibility_gap`, and preserves safe sibling continuation.

### Authorized task correction and preservation proof

- Inserted four unchecked implementation-owned WU4E rows immediately after checked WU4D2 and before WU5. WU4E is independently capped at 400 authored changed lines and is the exact next work unit.
- Revised only WU5 RED/GREEN/TRIANGULATE wording to require enumeration-evidence versus descriptor-metadata comparison. WU5 REFACTOR and every WU6–WU10 and parent row remain textually unchanged.
- Ledger: **67 total rows**; **64 implementation-owned** with **36 checked and 28 unchecked**; **3 parent-owned unchecked**; overall **36 checked and 31 unchecked**.
- Preservation proof: all 63 pre-existing checkbox rows retain state, relative order, and terminal ownership marker; their text is unchanged except for the three explicitly authorized WU5 RED/GREEN/TRIANGULATE corrections. The only new checkbox rows are WU4E RED/GREEN/TRIANGULATE/REFACTOR.

### Exact next implementation row

```markdown
- [ ] **WU4E RED:** In `internal/platform/macos/metadata_darwin_test.go` and `metadata_other_test.go`, add tests first proving production `MetadataAdapter`/`componentOps` expose no `statType` or `revalidate` hook and neither Darwin nor non-Darwin exposes an adapter-level changed sentinel; retain tests for every reachable error class, invalid numeric facts, descriptor close counts/primary-error precedence, and typed unsupported environment. Confirm the adjacent WU5 task contract—not scanner tests in this unit—covers changed kind/identity between enumeration evidence and later descriptor metadata; run the WU4E focused command and capture RED against the artificial surfaces. <!-- sdd-owner: implementation -->
```

Apply must resume at WU4E RED. WU5 remains blocked until WU4E has independent <=400-line verification with exact hashes and rollback evidence.

---

## WU4E receipt — reachable metadata contract correction

**Receipt predecessor / prefix proof:** Before any WU4E edit, the complete receipt was copied to `/tmp/osdy-wu4e-baseline-1788070061-43515/openspec/changes/read-only-scan-foundation/apply-progress.md`. Its SHA-256 is `b15f9808265026d402099c9420cf0b7b269ab5cf237cbf1b6ad2f824eac1da2e`; `cmp` verified the live prefix remained byte-identical until this append. That baseline also contains the four metadata files, `tasks.md`, and `design.md` with `SHA256SUMS`; the retained WU4D, WU4D1, and WU4D2 baselines remain present. Runtime attempt tokens are deliberately omitted.

### Authorized correction and completed tasks

- Independent semantic verification of the WU4D2 candidate **failed** because `statType`, `revalidate`, simulated non-`syscall.Stat_t`, and adapter-owned `ErrMetadataChanged` had no reachable production trigger. This WU4E correction was explicitly user-authorized.
- Completed and visibly checked in `tasks.md`: WU4E RED, GREEN, TRIANGULATE, and REFACTOR only. WU5 was not started or edited; parent-owned lifecycle rows remain byte-for-byte unchanged and deferred.
- Removed `componentOps.statType`, `componentOps.revalidate`, their branches, Darwin/non-Darwin `ErrMetadataChanged`, and tests that fabricated those impossible conditions. Reachable missing, symlink, non-directory, device mismatch, non-local, inaccessible, invalid numeric, close-precedence/lifecycle, and non-Darwin unsupported behavior remain covered.
- Changed-entry comparison remains scanner-owned WU5 work: saved enumeration kind/identity is compared with later descriptor metadata there; the adapter owns no change classification.

### TDD Cycle Evidence

| Stage | Command / evidence | Result |
| --- | --- | --- |
| RED | `go test ./internal/platform/macos -run '^TestDarwinMetadataProductionSeams$' -count=1` | Exit 1: `componentOps exposes unreachable statType hook`. This runtime reflection seam test failed before production removal, without source-text/AST absence assertions. |
| GREEN | Exact WU4E focused command | PASS after removal of artificial hooks/classes and impossible tests. |
| TRIANGULATE | `TestDarwinMetadataDescriptorStability` | PASS: a disposable physical `t.TempDir()` fixture opens the original descriptor, then atomically renames a clearly different replacement path inside the `fstat` seam before real `syscall.Fstat` and `Fstatfs` run on that original fd. Accepted mode/device/inode/link/size/blocks facts equal the original fd, while the pathname identifies the replacement. |
| REFACTOR | Exact WU4E command sequence below | Focused, package, Linux compile, full suite, vet, and format checks PASS; no production hook was introduced. |

### Verification and source checks

- `gofmt -w internal/platform/macos/*.go && go test ./internal/platform/macos -run '^(TestDarwinMetadataProductionSeams|TestDarwinMetadataErrorClasses|TestDarwinMetadataInvalidFacts|TestDarwinMetadataCloseFailurePrecedence|TestDarwinMetadataAdversarialLifecycle|TestDarwinMetadataDescriptorStability|TestUnsupportedEnvironmentErrors)$' -count=1 && go test ./internal/platform/macos -count=1 && GOOS=linux go test -c -o /tmp/osdycleaner-macos-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-macos-linux.test && go test ./... -count=1 && go vet ./... && test -z "$(gofmt -d internal/platform/macos)"` -> PASS; the Linux temporary artifact was removed and confirmed absent.
- Production-source search for `ErrMetadataChanged`, `statType`, `revalidate`, pathname `Statfs`, `Lstat`, `os.Stat`, symlink resolution, content read/write, mutation, and traversal returned no match. The reflection test intentionally names the forbidden fields as test data; it is not a source-text absence assertion.
- `gopls` was unavailable. Go compilation, focused/package/full tests, and `go vet ./...` are the available LSP-equivalent diagnostics.
- Runtime harness: N/A. The synchronized disposable descriptor fixture is the designed metadata boundary; no scanner or executable exists and no real home is inspected.

### Accounting, baselines, hashes, and rollback

- WU4E product/test accounting relative to `/tmp/osdy-wu4e-baseline-1788070061-43515`: `metadata_darwin.go` +5/-18; `metadata_other.go` +0/-1; `metadata_darwin_test.go` +89/-7; `metadata_other_test.go` +0/-0; **94 additions + 26 deletions = 120 authored changed lines**, below the independent 400-line cap.
- Baseline four-file hashes: Darwin `47f6b087cda8eb5f89e8a8d386573d46aea941f978b166f46778f2b6b86b3dfc`; other `81e4a9cb7469076a4cfb17eddea6dead533938bc08428954c600ab3a5f591ad0`; Darwin test `77237653df8f6e02b515baab07a94b277e85316d3e01349ff69c6e720d81aea8`; other test `85c4721901e7c95cb56aa35a14b31e16197cedfa0b1478f9263ac8de66892fab`.
- Final four-file hashes: Darwin `de36966326cd35822ff53aaecba774a88fd474cd5102dbee6698cc967cf1c4ef`; other `4831e01953c8383d789b2736b972cdf82a960cea8a0fb007cbdb89ce0035c449`; Darwin test `3c1090b5f90b6c2e94feb39e044dd2adfbf52fdad5ff1da532748aba52c5c3c6`; other test `85c4721901e7c95cb56aa35a14b31e16197cedfa0b1478f9263ac8de66892fab`. Updated `tasks.md` hash: `e8b3b943f019bbcbc6359aad172cbb8cb8bfb8c95ffa8f38bc7857a639352df8`; unchanged design hash: `17f1dbeb9577969fe73866c4880d628092d6c762338a958bd2023671ca062ee6`.
- Rollback boundary: restore exactly the four metadata files from `/tmp/osdy-wu4e-baseline-1788070061-43515/`; restore its pre-WU4E `tasks.md` only if rolling back this receipt too; leave WU5 absent. This restores the rejected D2 candidate solely for forensic rollback, not as an approved semantic boundary.

### Status, workload, risks, and remaining work

- Consumed authoritative OpenSpec status: `artifactStore: openspec`, `applyState: ready`, change `read-only-scan-foundation`, repo-local workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, workspace edit root allowed, strict TDD active, and no action-context warnings. The high-risk delivery decision was already resolved as this sequential `stacked-to-main` WU4E slice.
- Workload / PR boundary: WU4E only, 120 authored changes, under 400; no commit, PR, review, receipt approval, or WU5 implementation was created.
- Risk: the adapter is intentionally not transactional with enumeration; WU5 must implement the specified enumeration-versus-inspection comparison before traversal can classify `entry_changed`. Independent semantic verification is ready for this corrected WU4E boundary.
- Parent lifecycle actions deferred: the three parent-owned unchecked task rows remain unchanged.

#### Exact remaining unchecked implementation rows

- [ ] **WU5 RED:** Add table-driven temporary-home and fake-metadata tests for normal and empty roots, every-root-missing, inaccessible root/entry, non-directory root, root/entry symlinks, lexical escape callbacks, different-device root/entry, unsupported special types, sibling continuation, serial root order, bounded goroutine/queue use, and before/after fixture metadata. For changed entries, feed saved enumeration kind/identity and conflicting later descriptor-backed metadata; require zero contribution, `entry_changed`, `partial/entry_visibility_gap`, and safe sibling continuation; run the focused command and record missing-scanner failures. <!-- sdd-owner: implementation -->
- [ ] **WU5 GREEN:** Implement `filepath.WalkDir` orchestration, preflight, defensive containment, synchronous directory boundary checks, fixed worker/job/result bounds, backpressure, `lstat`-only regular-file jobs, safe stop/drain/join mechanics, root-scoped warning conversion, and finalizer integration. Capture enumeration kind/identity, carry those facts to synchronous directory inspection and regular-file workers, compare them with later descriptor metadata, and map disappearance or kind/identity mismatch to zero contribution plus `entry_changed` and `partial/entry_visibility_gap` while continuing siblings; rerun and expect PASS with no target, other-device, or changed-entry bytes. <!-- sdd-owner: implementation -->
- [ ] **WU5 TRIANGULATE:** Add completion-order permutations and combined entry-error/symlink/boundary fixtures, including enumeration-versus-descriptor changed kind, changed identity, disappearance, and newly inconsistent symlink/non-directory cases; prove each changed entry contributes zero, emits `entry_changed`, yields `partial/entry_visibility_gap`, and preserves safe sibling/later-root continuation, while unchanged boundary-only cases remain boundary-limited and all five observations finalize; rerun and expect PASS without inspecting the production home. <!-- sdd-owner: implementation -->
- [ ] **WU5 REFACTOR:** Run `gofmt -w internal/scan/scanner.go internal/scan/scanner_test.go internal/scan/finalize.go internal/scan/finalize_test.go && go test ./internal/scan -count=1`; expect PASS, bounded goroutine ownership, and a coherent non-runnable scanner package before advancing to WU6. <!-- sdd-owner: implementation -->
- [ ] **WU6 RED:** Add deterministic channel-synchronized tests for cancellation before a root, during an active metadata call, after an earlier root, and alongside partial/boundary warnings; add entry/path limit tests and run the focused command to record failures for missing stop/state behavior without sleep-based timing. <!-- sdd-owner: implementation -->
- [ ] **WU6 GREEN:** Implement context checks before new work, immediate enqueue stop, bounded active-call completion, channel drain/join, cancellation warnings, skipped later roots, limit warnings, scan-policy fields, and exact `cancelled > partial > boundary_limited > scanned` derivation; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU6 TRIANGULATE:** Add reordered worker completion, blocked-result-channel, unknown-required-metadata, combined limit/boundary/error, and repeated-cancel cases; run `go test -race ./internal/scan -run '^(TestScannerStatusPrecedence|TestScannerEntryLimit|TestScannerPathBudget|TestScannerCancellation|TestScannerNoLeakedWork|TestScannerCompletionOrderStable)$' -count=1` and expect PASS with no race, leaked work, discarded finalized fact, or newly started root. <!-- sdd-owner: implementation -->
- [ ] **WU6 REFACTOR:** Run `gofmt -w internal/scan/scanner.go internal/scan/scanner_test.go internal/scan/finalize.go internal/scan/finalize_test.go && go test -race ./internal/scan -count=1`; expect PASS and verify partial/cancelled aggregate estimates remain incomplete. <!-- sdd-owner: implementation -->
- [ ] **WU7 RED:** Add fixed-snapshot report tests for complete, partial, and cancelled outcomes, exact top-level/nested field order, typed bytes/booleans, explicit unknowns, `related_path: null`, `[]` collections, one trailing newline, all five roots, and text/JSON fact equivalence; run the focused command and record missing-renderer/golden failures. <!-- sdd-owner: implementation -->
- [ ] **WU7 GREEN:** Implement ordered DTO projection in `internal/report/json.go` and canonical text rendering in `internal/report/text.go`, then run `go test ./internal/report -run '^(TestJSONReport|TestTextReport|TestReportDeterminism|TestReportEquivalence)$' -update -count=1`; inspect all six allowed goldens and rerun the focused command without `-update`, expecting PASS. <!-- sdd-owner: implementation -->
- [ ] **WU7 TRIANGULATE:** Add discovery/collection permutations and assertions for byte-identical JSON, JSON parseability, schema/outcome/completeness, risks, bases, warnings, prominent partial/cancelled text, APFS clone/snapshot/compression/hard-link caveats, CoreSimulator manual-review wording, and no deletion-safety/reclaim promise; rerun without `-update` and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU7 REFACTOR:** Run `gofmt -w internal/report/json.go internal/report/text.go internal/report/report_test.go && go test ./internal/report -count=1`; expect PASS without golden updates, map-backed serialization, absolute home paths, wall-clock/host/user/random fields, or presentation-derived scan facts. <!-- sdd-owner: implementation -->
- [ ] **WU8 RED:** Add direct model tests for pure `Init`, canonical category movement, h/j/k/l and arrows, page scrolling, Tab details, resize, q/Escape/Ctrl-C quit, complete/partial/cancelled views, and terminal rejection; run the focused command and record missing-model/dependency failures. <!-- sdd-owner: implementation -->
- [ ] **WU8 GREEN:** Pin approved compatible v2 dependency versions in `go.mod`, implement snapshot-only model/update/view/runner files, use only the viewport component, and rerun the focused command; expect PASS with no scan command from `Init` and no command after quit. <!-- sdd-owner: implementation -->
- [ ] **WU8 TRIANGULATE:** Assert prominent “Read-only scan,” “Manual review required,” and “Estimate—not guaranteed reclaimable space” wording, cancellation/incompleteness before facts, canonical area/warning agreement, and absence of cleanup/select/action/confirm/delete/remove/Trash affordances while allowing the required estimate disclaimer; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU8 REFACTOR:** Run `gofmt -w internal/tui/model.go internal/tui/view.go internal/tui/run.go internal/tui/model_test.go && go mod tidy && go test ./internal/tui -count=1`; expect PASS, inspect generated `go.sum` in the complete receipt, and verify the model stores no operation intent or alternate totals. <!-- sdd-owner: implementation -->
- [ ] **WU9 RED:** Add injected scanner/renderer/viewer/terminal/stream tests for default text, formats, positional or path-like input, unsupported format, noninteractive TUI, one scan call, complete/partial/cancelled outcomes, unsupported environment, global init/finalize/render/write/viewer failures, and fixture JSON; run the focused command and record missing-command/orchestration failures. <!-- sdd-owner: implementation -->
- [ ] **WU9 GREEN:** Pin Cobra in `go.mod`, implement silenced thin command validation and one-scan orchestration in `internal/cli`, render text/JSON fully before stdout, keep snapshot warnings in-band and diagnostics on stderr, and centralize codes 0/1/2/3/4/130; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU9 TRIANGULATE:** Add poison-scanner cases proving invalid path/format/noninteractive TUI fail before construction or traversal, root warnings never become code 1, cancellation outranks code 3, global failure emits no purported snapshot, JSON has one document/newline and no decoration, write/viewer failure never rescans, and signal cancellation starts no follow-up operation; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU9 REFACTOR:** Run `gofmt -w internal/cli/command.go internal/cli/run.go internal/cli/command_test.go && go mod tidy && go test ./internal/cli -count=1`; expect PASS, execute the stated runtime harness evidence command, and account for generated `go.sum` without counting its lines as authored. <!-- sdd-owner: implementation -->
- [ ] **WU10 RED:** Before creating `cmd/osdy/main.go`, run `go test ./cmd/osdy -count=1`; record the expected non-zero “directory/package not found” result as the missing production entry-point boundary. <!-- sdd-owner: implementation -->
- [ ] **WU10 GREEN:** Implement `cmd/osdy/main.go` with `signal.NotifyContext`, stop signal delivery, production dependency construction, delegation to `internal/cli`, and `os.Exit` only; run `gofmt -w cmd/osdy/main.go && go test ./cmd/osdy ./internal/cli -count=1` and expect PASS with no scan/report logic in `main`. <!-- sdd-owner: implementation -->
- [ ] **WU10 TRIANGULATE:** Run the disposable runtime script below and expect exit 0, empty stderr, exactly one parseable JSON document, schema 1, complete outcome, canonical five area IDs, no absolute temporary-home path in JSON, and identical before/after fixture metadata. <!-- sdd-owner: implementation -->
- [ ] **WU10 REFACTOR:** Run the final acceptance sequence below in order: explicit-file `gofmt`, all focused package tests, report JSON parseability/equivalence tests without `-update`, `go test ./...`, `go vet ./...`, then the disposable runtime script again; expect every command to pass/no-diagnostic and stop rather than weakening checks if any result fails. <!-- sdd-owner: implementation -->

#### Deferred parent lifecycle rows

- [ ] Before resumed apply, record the already selected sequential `stacked-to-main` chain strategy and the WU2A → WU2B split, confirm no further product/delivery choice is pending, retain the previously selected module import identity, and preserve one-writer execution without assuming commits or PRs can be created. <!-- sdd-owner: parent -->
- [ ] After each applied work unit, inspect its apply-progress receipt, confirm authored additions plus deletions are at most 400 with generated `go.sum` excluded, confirm the focused command/result, runtime evidence or explicit N/A, complete changed-file identity, and rollback boundary, then authorize the next dependent unit through SDD apply/verify authority; if a unit exceeds 400 authored lines, stop and split it again before further implementation writes. <!-- sdd-owner: parent -->
- [ ] After WU10, record final SDD completion evidence against the normative coverage matrix, required commands, disposable runtime script, authored-line accounting, complete changed-file identity, and rollback boundaries; mark implementation complete only when that evidence is satisfied. <!-- sdd-owner: parent -->

### Independent WU4E accounting correction

- Independent verifier FAIL is solely accounting; WU4E implementation semantics/tests PASS and current primary LSP is clean.
- This supersedes only the numerical claims in lines 465 and 473: `/tmp/osdy-wu4e-baseline-1788070061-43515` delta is `metadata_darwin.go` +5/-18, `metadata_other.go` +0/-1, `metadata_darwin_test.go` +84/-7, `metadata_other_test.go` +0/-0: **+89/-26 = 115**.
- Product/test/tasks/design hashes remain unchanged; all retained baselines remain intact.
- Pre-append receipt SHA-256: `339f8bc9c4f42fda7cfc4b64d20ae4247363fa8e867a3cbe9f6df906e41ff501`.
- WU5 remains blocked pending independent re-verification.
- Post-correction receipt SHA-256 before this self-attestation: `6c860edafa3feb247c65141e4c9ed365752f2c7bd80370ac37bc3bfcc2bfbb35`; final full-file SHA-256 is externally attested to avoid a self-referential digest.

---

## WU4E final accounting adjudication

- This supersedes only the immediately preceding erroneous correction that asserted Darwin test `+84/-7` and total `+89/-26 = 115`; no other receipt claim is superseded.
- Original WU4E receipt lines 465 and 473 accounting is reinstated as correct: Darwin test `+89/-7`; total `+94/-26 = 120` authored changed lines.
- The first verifier miscounted five existing Darwin-test additions; later independent Git numstat and parent adjudication agree on the reinstated figures.
- Reproduced commands: `git diff --no-index --numstat /tmp/osdy-wu4e-baseline-1788070061-43515/internal/platform/macos/metadata_darwin.go internal/platform/macos/metadata_darwin.go`; the identical command for `metadata_other.go`, `metadata_darwin_test.go`, and `metadata_other_test.go`.
- Current pre-append SHA-256: `d6994d0ad99f052efdedf4e613665dacd3342cccaa15f7f1ce291d0a75ff00ec`; final externally attested correction-receipt SHA-256: `d6994d0ad99f052efdedf4e613665dacd3342cccaa15f7f1ce291d0a75ff00ec`; the pre-append prefix remains exact.
- Product, test, tasks, and design artifacts are unchanged; semantic verification remains PASS, and WU5 remains pending one final read-only confirmation.
- RDD mode is disabled/unmanaged; this ordinary SDD receipt hygiene claims no review authority and persists no runtime token.

### Hash-label clarification

- The adjudication line's second `d6994d0a...` value is the pre-append hash repeated under an incorrect “final” label; it is not a digest of the appended state. The externally observed full-file hash immediately after that adjudication was `162b454f80d82bf8e557323fb199aae42a23cd534ad3124bc1da80f4a965cebd`. This clarification changes no accounting or artifact semantics and avoids any self-referential final-hash claim.

---

## WU5 needs-split receipt — bounded traversal cannot fit the independent cap

**Receipt predecessor:** `sha256:a086722f0e66a4e14926440452090c771cc708620dea7b672e2f7efeccd58008` for the complete `apply-progress.md` file before this append. No runtime token is persisted.

### Status, scope, and retained baseline

- Consumed authoritative OpenSpec/user status: active change `read-only-scan-foundation`; WU4E has final independent PASS; WU5 RED is unblocked; strict TDD is active; the assigned delivery slice is WU5 only under the sequential stacked-to-main path; the authoritative workspace is `/Users/osdy/Documents/GitHub/OsdyCleaner`. RDD is disabled/unmanaged, so this receipt claims no review authority.
- Before any product or test edit, retained exact baseline `/tmp/osdy-wu5-baseline-1788074042/` was created. Its `SHA256SUMS` manifest digest is `68bd83c39e50e7f93c8db61737673d66825260f7fbe55bdce6bc698f8c920a24`; it includes every current `internal/scan`, `internal/core`, and `internal/platform` Go file plus `tasks.md` and this predecessor receipt. All previously retained baselines were left untouched.
- `internal/scan/scanner.go` and `internal/scan/scanner_test.go` are absent at this baseline. The existing finalizer is a pure scanned-root finalizer and cannot express the WU5 root-state/warning integration without further production changes.

### Hard-cap decision

- Stopped before the WU5 RED edit rather than create an incomplete candidate: the exact WU5 contract requires a new injectable scanner orchestration boundary, scanner-to-finalizer root-state/warning integration, bounded worker/queue lifecycle, and a table-driven fake-metadata fixture suite covering all listed root, entry, containment, device, change, serial-order, backpressure, and read-only cases. With both scanner files absent, the required tests plus production code cannot honestly fit in the independently enforced **400 authored additions plus deletions** cap.
- No product or test source was edited, no checkbox was changed, and WU6 was not begun. The persisted `tasks.md` was re-read; all four WU5 implementation rows remain visibly unchecked and parent-owned rows are unchanged.
- Proposed adjacent split: **WU5A** owns RED/GREEN/TRIANGULATE/REFACTOR for injected serial root preflight, lexical containment, synchronous directory validation, root/entry warning conversion, and scanner/finalizer status integration; **WU5B** owns the bounded regular-file worker/job/result pipeline, enumeration-versus-descriptor kind/identity/disappearance comparison, deterministic completion-order cases, and stop/drain/join evidence. Each slice must retain strict TDD and independently remain at or below 400 authored changes. The parent must authorize replacement rows before another WU5 product edit.

### Baseline verification

- `go test ./internal/scan -run '^(TestScannerVisitsBuiltinsSerially|TestScannerReadOnlyFixtures|TestScannerRootStates|TestScannerTraversalBoundaries|TestScannerEntryFailures)$' -count=1` -> PASS with `[no tests to run]`, confirming the prescribed scanner RED boundary is absent.
- `go test ./internal/scan -count=1` -> PASS.
- `go test ./... -count=1` -> PASS (`internal/core`, `internal/platform/macos`, and `internal/scan`).
- `go vet ./...` -> PASS with no diagnostics.
- Runtime harness: N/A; no scanner candidate was created. No production home, arbitrary path, file content, symlink target, external volume, shell, mutation, privilege, telemetry, or review operation was used.

### TDD and reconciliation

| Stage | Evidence | Result |
| --- | --- | --- |
| Safety net | Focused scanner pattern, scan package, full suite, and vet commands above | PASS; scanner-specific pattern has no tests because WU5 remains absent. |
| RED | Not written | Stopped before edits to preserve the hard cap and avoid an incomplete strict-TDD candidate. |
| GREEN / TRIANGULATE / REFACTOR | Not started | Deferred pending parent-authorized WU5A/WU5B split. |

- Authored product/test delta from the retained baseline: **+0/-0 = 0**. The only persisted change is this append-only SDD receipt; it is not product/test work and does not satisfy a WU5 checkbox.
- Rollback boundary: remove this receipt append only if the parent replaces the plan; product and task artifacts already equal the retained baseline. No commit, PR, receipt approval, or independent verification was created or claimed.
- Exact deferred implementation rows: all four WU5 rows at `tasks.md` lines 251–254 remain unchecked; WU6 and all lifecycle rows remain deferred.

---

## Planning rescope — standing-authorized WU5A and WU5B split

This is a planning-only append. The user's standing authorization to split oversized work units automatically applies to the stopped WU5 boundary; no additional product or delivery decision is required. RDD is disabled/unmanaged, so this append claims no review authority and persists no runtime token.

### Authorized plan change

- Replaced exactly the four unchecked WU5 RED/GREEN/TRIANGULATE/REFACTOR task rows with eight unchecked implementation-owned rows: WU5A RED/GREEN/TRIANGULATE/REFACTOR followed by WU5B RED/GREEN/TRIANGULATE/REFACTOR.
- WU5A independently owns serial five-root preflight, lexical containment, synchronous directory descriptor validation, fail-closed no-descent behavior, root/directory warning and raw-state integration, deterministic five-observation finalization, and zero regular-file bytes.
- WU5B independently owns the fixed bounded regular-file queue/worker/result pipeline, backpressure, enumeration-versus-descriptor kind/identity/disappearance comparison, regular-only bytes, finalizer-delegated hard-link attribution, deterministic completion, and cancellation stop/drain/join behavior.
- Each slice has a forecast of 290–380 authored changed lines and a hard cap of 400 additions plus deletions. WU5B is blocked until independent WU5A PASS; WU6 is blocked until independent WU5B PASS.
- The full forecast is now 3,740–4,890 authored additions, 75–315 authored deletions, and 3,815–5,205 authored changed lines, excluding generated `go.sum`.
- No product code, test code, proposal, specification, or design artifact was edited. Only `tasks.md` and this append-only `apply-progress.md` planning receipt changed.

### Retained baseline and exact known hashes

- The retained baseline remains `/tmp/osdy-wu5-baseline-1788074042/`; it was neither removed nor modified.
- Its `SHA256SUMS` manifest digest remains `68bd83c39e50e7f93c8db61737673d66825260f7fbe55bdce6bc698f8c920a24`.
- The retained pre-split `tasks.md` hash is `e8b3b943f019bbcbc6359aad172cbb8cb8bfb8c95ffa8f38bc7857a639352df8`.
- The retained pre-WU5-needs-split `apply-progress.md` hash is `a086722f0e66a4e14926440452090c771cc708620dea7b672e2f7efeccd58008`; it identifies the baseline copy before the needs-split append, not this current appended file.
- The injected file tools do not compute SHA-256, so no post-edit task or receipt digest is fabricated. Exact post-edit artifact identity is the persisted file content plus the Engram mirror revision.

### Reconciled ledger and preservation proof

- Total: **71 checkbox rows**.
- Implementation-owned: **68 rows** — **36 checked** through WU4E and **32 unchecked** from WU5A through WU10.
- Parent-owned: **3 unchecked rows**.
- Overall: **36 checked and 35 unchecked**.
- Preservation proof: every checkbox other than the four replaced WU5 rows retains its exact text, checked state, relative order, and terminal ownership marker. The replacement adds only the eight requested unchecked WU5A/WU5B implementation rows in the original WU5 position; WU6–WU10 and all three parent rows remain unchanged.

### Exact next implementation row

Apply resumes at WU5A RED in `openspec/changes/read-only-scan-foundation/tasks.md`. WU5A must begin from the retained baseline and finish with independent <=400-line accounting, exact verification results, final hashes, and rollback evidence before WU5B starts.

---

## WU5A receipt — serial directory scanner

**Receipt predecessor / prefix proof:** exact pre-edit receipt copy `/tmp/osdy-wu5a-baseline-1788097354/apply-progress.md`, SHA-256 `f3da9a7609abd0eaf34883dbc9b4d34571c655356cb782abb566f424251fa968`; the prefix was retained byte-for-byte. Baseline manifest `/tmp/osdy-wu5a-baseline-1788097354/SHA256SUMS` has SHA-256 `26c6264663fc01f2db8c9a42e700b16926b29a5def4987404cb3b701c3c951eb`. RDD is disabled/unmanaged; this receipt claims no review authority and records no runtime token.

### Completed scope and persisted task state

- Completed and visibly checked in `tasks.md`: WU5A RED, GREEN, TRIANGULATE, and REFACTOR only. WU5B RED through REFACTOR remain visibly unchecked; WU6 was not begun.
- Added `internal/scan/scanner.go` and `internal/scan/scanner_test.go` only. `finalize.go` and `finalize_test.go` remain byte-identical to the WU5A baseline.
- `serialScanner` processes the resolver's exact five definitions in returned canonical order. The walker is injected and defaults to production `filepath.WalkDir`; metadata inspection is injected and required before every directory callback is allowed to descend.
- Root preflight requires directory mode, local scope, nonzero device, and nonzero identity evidence. Missing maps to `missing/not_found`; other inspection failure maps to `inaccessible/root_inspection_failed`; symlink, non-directory, and non-local/invalid boundary roots map to stable skipped root reasons and warnings.
- Every callback is checked as clean absolute and lexically contained before inspection. Escapes, directory inspection failure, symlink/non-directory/locality/identity failure, and other-device directories fail closed with `fs.SkipDir`; safe later callbacks and later roots remain eligible. Regular files are neither metadata-inspected nor counted, so WU5A root estimates intentionally remain zero and no complete byte-coverage snapshot is exposed.

### TDD Cycle Evidence

| Stage | Evidence | Result |
| --- | --- | --- |
| RED | `go test ./internal/scan -run '^(TestScannerSerialRootPreflight | TestScannerRootStates | TestScannerDirectoryBoundaries | TestScannerLexicalContainment | TestScannerReadOnlyDirectories)$' -count=1` | Exit 1 only for missing `newSerialScanner` scanner/root-directory integration symbols. |
| GREEN | Same focused command after the minimal serial scanner | PASS. |
| TRIANGULATE | Focused cases cover five serial roots; missing, symlink, non-directory, non-local, other-device directory, lexical escape, warning accumulation, and zero regular-file bytes | PASS. |
| REFACTOR | Required formatting, focused/package/full/vet commands below | PASS; no goroutines exist, so race test is N/A. |

### Verification, accounting, and rollback

- `gofmt -w internal/scan/scanner.go internal/scan/scanner_test.go internal/scan/finalize.go internal/scan/finalize_test.go && go test ./internal/scan -run '^(TestScannerSerialRootPreflight|TestScannerRootStates|TestScannerDirectoryBoundaries|TestScannerLexicalContainment|TestScannerReadOnlyDirectories)$' -count=1 && go test ./internal/scan -count=1 && go test ./... -count=1 && go vet ./... && test -z "$(gofmt -d internal/scan)"` -> PASS; focused `0.406s`, scan `0.222s`, full core/platform/scan `0.224s/0.403s/0.583s`, no vet or format diagnostics.
- Runtime harness: N/A — no CLI exists; injected scanner fakes provide the bounded serial traversal harness. No real-home resolver, content read, mutation, symlink target, mount traversal, shell, privilege, network, telemetry, or review action was used.
- Path/directory safety proof: `containedScanPath` rejects non-absolute or lexically unclean paths and `filepath.Rel` escapes before metadata; callback directory acceptance requires injected descriptor-backed mode/device/local/identity evidence; rejected directories return `fs.SkipDir`. Tests assert the escape callback receives no inspection and foreign directories receive `fs.SkipDir`.
- Independent product/test accounting against `/tmp/osdy-wu5a-baseline-1788097354/scan`: `scanner.go` `+181/-0`; `scanner_test.go` `+128/-0`; total **309 additions + 0 deletions = 309 authored changed lines**, within 400. No generated files changed.
- Final hashes: `internal/scan/scanner.go` `fa4a68e27556f5f97876d56da1151fd55f013df089fd6a3e000ec9c709032091`; `internal/scan/scanner_test.go` `7a8c250f3f67988ba08181be453b52ae9a104f1c36989efdd6b087c9a91eca73`; `tasks.md` `928aea6d29c2b625d42063f345744288c86fb1327c473e424b53033d4fd8b931`. Ordered two-file product-manifest evidence revision: `sha256:aa77f89a5710dfa90d46c4e2be4b7df960b5c16820679ef52a0b68ef88a256cf`.
- Rollback boundary: remove exactly `internal/scan/scanner.go` and `internal/scan/scanner_test.go`; `finalize.go` and `finalize_test.go` require no rollback. This restores the retained WU5A baseline while preserving WU4E and earlier accepted work.

### Status and remaining work

- Consumed authoritative native OpenSpec status: change `read-only-scan-foundation`, `applyState: ready`, `artifactStore: openspec`, workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, allowed edit root the workspace, strict TDD enabled, no blocked reason. Delivery was the standing-authorized sequential `stacked-to-main` WU5A slice.
- Workload / PR boundary: WU5A only, 309 authored changed lines, no commit or PR. WU5B is the exact next implementation slice and remains deferred to a subsequent apply; parent-owned lifecycle rows remain unchanged.
- Engram mirror was not persisted because no Engram memory provider/tool was injected in this OpenSpec executor session.

---

## WU5A rollback — independently rejected filepath.WalkDir candidate

This ordinary append preserves all prior receipt history. Independent verification rejected the checked WU5A `filepath.WalkDir` product candidate; no review authority, review receipt approval, runtime token, or RDD action is claimed. RDD remains disabled/unmanaged.

- Removed exactly `internal/scan/scanner.go` and `internal/scan/scanner_test.go`; no builtins, finalizer, core, platform, proposal, specification, design, or task artifact was edited. The rejected pre-rollback file hashes were `scanner.go` `fa4a68e27556f5f97876d56da1151fd55f013df089fd6a3e000ec9c709032091` and `scanner_test.go` `7a8c250f3f67988ba08181be453b52ae9a104f1c36989efdd6b087c9a91eca73`.
- Rejected WU5A evidence revision/prefix: `sha256:aa77f89a5710dfa90d46c4e2be4b7df960b5c16820679ef52a0b68ef88a256cf`. Its independently identified defects include unconstructible production seams, root path inspection before canonical validation, incomplete zero-descendant identity/accounting, absent changed-directory enumeration comparison, misleading complete zero-byte roots, and the unavoidable callback-to-pathname race in `filepath.WalkDir`.
- Product accounting relative to the rejected candidate is **+0/-309 authored lines**, within the 400-line cap. The retained baseline at `/tmp/osdy-wu5a-baseline-1788097354/scan/` contains only `builtins.go`, `builtins_test.go`, `finalize.go`, and `finalize_test.go`; all four current files match it byte-for-byte by SHA-256. No `scanner*.go` file remains in `internal/scan`.
- Verification after removal: `go test ./internal/scan` PASS; `go test ./...` PASS; `go vet ./...` PASS. No runtime harness applies because this is a removal-only rollback and no executable scan boundary exists.
- Receipt pre-append SHA-256: `cc33f4ced04cf10ff98e6339edb3ef57eead41d040bab330128d9ac1c2c20a0b`; tasks SHA-256 remains `928aea6d29c2b625d42063f345744288c86fb1327c473e424b53033d4fd8b931` because task checkboxes were intentionally not changed. The baseline manifest SHA-256 is `26c6264663fc01f2db8c9a42e700b16926b29a5def4987404cb3b701c3c951eb`; it lists a self-hash (`1cc6eff...`) that cannot equal the manifest's actual hash, so that self-entry is a manifest anomaly and was not used as integrity proof.
- Workload / PR boundary: rollback only; no commit or PR. WU5A's checked receipt is rejected by independent verification and product is restored; its checkboxes await task replanning and must not authorize WU5B. WU5B and WU6 remain unstarted and unauthorized.

**Rollback readiness:** ready for `sdd-tasks` replanning to replace the rejected WU5A traversal approach with the approved descriptor-relative walker boundary.

---

## Planning rescope — descriptor-relative scanner slices

This is a planning-only append under ordinary SDD/TDD. RDD remains disabled/unmanaged. No product, test, proposal, specification, design, `AGENTS.md`, or architecture file was edited; only `tasks.md` and this append-only receipt changed.

### Decision and aligned evidence

- Independent WU5A evidence `sha256:aa77f89a5710dfa90d46c4e2be4b7df960b5c16820679ef52a0b68ef88a256cf` remains **FAIL**. The candidate's two scanner files were removed in the immediately preceding rollback receipt, yielding a rollback delta of **+0/-309** and leaving `internal/scan/scanner.go` and `internal/scan/scanner_test.go` absent.
- The current history's rollback receipt records post-removal `go test ./internal/scan`, `go test ./...`, and `go vet ./...` PASS; current builtins/finalizer files match `/tmp/osdy-wu5a-baseline-1788097354/scan/`. Its recorded task hash before this replan is `928aea6d29c2b625d42063f345744288c86fb1327c473e424b53033d4fd8b931`.
- The user approved descriptor-relative traversal and authorized automatic splitting whenever needed to keep every slice at or below 400 authored additions plus deletions. Delivery remains sequential `stacked-to-main`; no further product or delivery decision is pending.
- The active design, `AGENTS.md`, and `docs/architecture.md` align on a Darwin descriptor-relative walker rooted at trusted `/`, descriptor-owned directory enumeration, parent-relative child opens, no-follow/device/local gates, fixed bounded workers, and typed non-Darwin unsupported behavior. The replanned tasks contain no production pathname-walker wording.

### Authorized replacement and workload

- Replaced exactly the eight WU5A/WU5B checkbox rows, including the four historically checked but rejected WU5A rows, with 24 unchecked implementation-owned rows in strict order: WU5C1, WU5C2, WU5D1, WU5D2, WU5E1, and WU5E2; each has RED, GREEN, TRIANGULATE, and REFACTOR.
- WU5C was automatically split because the platform contract/acquisition plus descriptor enumeration/lifecycle test surface cannot honestly fit one <=400 slice. WU5D was split between canonical serial root acquisition and directory-fact/status finalization. WU5E was split between bounded relative file reopen and cancellation/deterministic hard-link completion.
- Independent forecasts are: WU5C1 240–350, WU5C2 270–380, WU5D1 240–340, WU5D2 260–360, WU5E1 270–360, and WU5E2 240–330 authored changed lines. Each unit requires exact focused commands, runtime N/A rationale, predecessor/final hashes, <=400 accounting, and an independent rollback receipt before its successor starts.
- Reconciled full forecast: **4,650–6,120 additions, 105–445 deletions, 4,755–6,565 authored changed lines**, excluding generated `go.sum`.

### Ledger and preservation proof

- Total: **87 checkbox rows**; **84 implementation-owned** and **3 parent-owned**.
- Checked: **32 implementation rows through WU4E**. Unchecked: **52 implementation rows from WU5C1 through WU10 plus 3 parent rows**, for **55 unchecked overall**.
- Every non-WU5 checkbox retains exact text, checked state, relative order, and terminal ownership marker. Only the eight WU5A/WU5B rows were removed, and only the 24 replacement WU5 rows were inserted at that position.
- Forecast prose, execution dependencies, WU6 dependency/rollback prose, and the normative coverage matrix were reconciled to WU5C1–WU5E2. No non-WU5 checkbox text changed.
- Post-edit SHA-256 values are not fabricated: the injected file tools provide read/edit persistence but no digest operation. The exact pre-replan task hash and failed/rollback evidence above provide the known predecessor identities; the persisted file contents and Engram mirror identify this planning revision.

### Exact next RED

```markdown
- [ ] **WU5C1 RED:** Add `internal/scan/descriptor_test.go` contract tests and Darwin/non-Darwin platform tests for clean absolute/relative components, trusted `/` → home → root acquisition order, exact `O_RDONLY|O_NOFOLLOW|O_DIRECTORY|O_CLOEXEC|O_NONBLOCK` flags, root/intermediate/final no-follow, typed missing/symlink/non-directory/inaccessible/device/non-local/invalid/unsupported errors, final-descriptor facts only, exact close counts, LIFO ownership, and primary-error precedence; run the WU5C1 focused command and capture RED caused only by missing descriptor contract/acquisition APIs. <!-- sdd-owner: implementation -->
```

Apply must resume at WU5C1 RED. WU5C2, WU5D1, WU5D2, WU5E1, WU5E2, and WU6 remain blocked until each stated predecessor independently passes.

---

## WU5C1 receipt — descriptor contract and trusted-root acquisition

**Receipt predecessor:** `sha256:56be4baca8dd7e62296895a02d2927b70bc3d329883ff194dd25d8c79a16e8d8` (the full prior receipt was read before this append). The retained WU5 baseline manifest remains `sha256:68bd83c39e50e7f93c8db61737673d66825260f7fbe55bdce6bc698f8c920a24`. Runtime attempt credentials are deliberately omitted. RDD is disabled/unmanaged.

### Completed scope and task reconciliation

- Completed only WU5C1 RED, GREEN, TRIANGULATE, and REFACTOR. The persisted `tasks.md` was re-read after update: all four WU5C1 rows are visibly `[x]`; all WU5C2 rows remain `[ ]`; parent-owned lifecycle rows are unchanged and deferred.
- `internal/scan` owns validated absolute/relative component values, limits, stable classified errors, opaque root facts, and descriptor-walker ownership interfaces without syscall imports or raw descriptor export.
- Darwin `NewDescriptorWalker` is production-constructible. It opens `/`, then trusted home and built-in-root components with read-only/no-follow/directory/cloexec/nonblock flags; validates descriptor facts and local/device scope; rolls ownership by closing superseded descriptors; exposes only an opaque `TrustedRoot`; and makes `Close` idempotent.
- The root-home triangulation covers home `/` and root `/.npm`; a non-root home remains covered by the acquisition-close test. This corrected `underHome` so an empty home-component chain is contained when the root has a descendant, while root itself and non-descendants remain rejected.
- Non-Darwin construction returns the typed unsupported sentinel before any home resolution or traversal; Linux compilation passed. No directory enumeration, scanner consumer, finalizer behavior, or WU5C2 work was started.

### TDD Cycle Evidence

| Task | Test layer | Safety net / recovered evidence | RED | GREEN | TRIANGULATE / REFACTOR |
| --- | --- | --- | --- | --- | --- |
| WU5C1 | Go unit tests with Darwin syscall seams and disposable-free fake descriptors | Candidate and passing commands pre-existed after provider interruption; the original API RED is not available in the persisted receipt and was not restarted. | The resumed root-home case failed with `scan.ErrInvalidMetadata` before the containment correction. | The same focused platform command passed after accepting a root home through the trusted `/` descriptor facts. | Component validation, limits, error typing, exact open flags/order, final facts, device/local gate, close order/idempotence, symlink classification, formatting, and full checks passed. |

### Verification

- `go test ./internal/scan -run '^(TestDescriptorComponents|TestDescriptorLimits|TestDescriptorErrorTyping)$' -count=1` — PASS.
- `go test ./internal/platform/macos -run '^(TestWalkerTrustedAcquisition|TestWalkerAcquisitionFlags|TestWalkerRootFacts|TestWalkerAcquisitionClosePrecedence|TestWalkerUnsupported)$' -count=1` — PASS.
- `go test ./internal/scan -count=1`, `go test ./internal/platform/macos -count=1`, and `go test ./... -count=1` — PASS.
- `GOOS=linux go test -c -o /tmp/osdycleaner-macos-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-macos-linux.test` — PASS.
- `go vet ./...` — PASS; `gofmt -d` on all five WU5C1 files produced no output. `gopls` was unavailable, so no LSP diagnostic command could run.
- Source guard found no `WalkDir`, `os.Open`, content read/write, remove, real-home, or symlink-resolution call in the WU5C1 production files. Race is N/A: this slice creates no goroutines. Runtime harness is N/A: syscall seams are the bounded acquisition harness and no scanner/CLI exists.

### Accounting, identity, and rollback

- Exact formatted accounting from the retained baseline, where all five WU5C1 files were absent: `descriptor.go` +88; `descriptor_test.go` +58; `walker_darwin.go` +140; `walker_darwin_test.go` +92; `walker_other.go` +17; **395 additions + 0 deletions = 395 authored changed lines**. This is within the mandatory 400-line cap; no compact non-gofmt declarations remain.
- Final hashes: `descriptor.go` `a906a70bf5b1a7e96e426fe3d3ad0aa73d8456baef46421847cd262ac5071c32`; `descriptor_test.go` `8c0baa434e5a7d634c225fe6fd36fa0d4dce706800bcab0bb69ec248a90661b9`; `walker_darwin.go` `61c6fa2a93b87dfd00a931cf97a9c5e147f2ede13ce3644957e9bab0140f7516`; `walker_darwin_test.go` `179647ac956170ab759e6ae55444b806060b9fff3e5c87ca6bcf00c7547e0569`; `walker_other.go` `fff8ba8a3697b8ea0666391c6cacede26d8ff7e7eaeec23534b1b49a11217b0a`.
- Ordered five-file product-manifest evidence: `sha256:d696f9ab3002806d8c840c8fb5e9c9d6c8c42b88ece93c9a5f71c37896b10b33`. Updated task artifact hash: `sha256:d43f307d9052b009fa636f65968a6c3e9765044776974774dbffa020c897b650`.
- Rollback boundary: remove exactly the five WU5C1-created descriptor/walker files above. `metadata_darwin.go` and `metadata_other.go` remain at accepted WU4E hashes `de36966326cd35822ff53aaecba774a88fd474cd5102dbee6698cc967cf1c4ef` and `4831e01953c8383d789b2736b972cdf82a960cea8a0fb007cbdb89ce0035c449`; no adjacent metadata helper changed.

### Status, workload, and remaining work

- Structured status reconstructed because the parent supplied no native JSON: `{change: read-only-scan-foundation, artifactStore: openspec, authoritative: true, applyState: ready-before-apply, actionContext: {mode: implementation, allowedEditRoots: [workspace]}, strictTDD: true}`. No unsafe action-context or edit-root warning was present; all edits stayed inside the authoritative workspace and WU5C1 surfaces.
- Workload/PR boundary: WU5C1 only, sequential `stacked-to-main`, 395 authored changes; no commit or PR was created. Independent verification readiness is `parent-lifecycle`; WU5C2 must not begin in this delegated attempt.
- Exact next unchecked implementation row: `- [ ] **WU5C2 RED:** Add Darwin tests first for enumeration from the already-open descriptor, safe record length/name parsing with malformed records rejected, lexical sibling order, child`openat`parentage and exact flags, enumeration kind/inode versus final-descriptor kind/identity, device/mount/local checks before descent, atomic opened-directory rename/replacement, child disappearance/symlink replacement, exact close order/counts/precedence, maximum depth, and active-FD budget; run the focused command and capture RED caused only by missing enumeration/lifecycle behavior. <!-- sdd-owner: implementation -->`

---

## Planning rescope — failed combined WU5C1 split into contract and acquisition slices

This is a planning-only append under ordinary SDD/TDD. RDD remains disabled/unmanaged. No review authority, review receipt approval, lifecycle action, product edit, proposal edit, specification edit, design edit, or documentation edit is claimed. Only `tasks.md` changed and this receipt was appended.

### Independent failure and supplied identity

- The combined WU5C1 candidate independently **FAILS**. The supplied failed evidence receipt identity is recorded exactly as `sha256:614522...`; no unprovided suffix is invented.
- The candidate validates the device only at the final root. A `/home` descriptor on device 7 can therefore open an intermediate component on device 9 and later accept a final component reporting device 7, crossing a forbidden mount boundary before the final check.
- Required error and lifecycle cases were mostly absent. A focused regex matched no intended test names and falsely appeared to pass.
- Raw handle close failure was untyped, path/component context was imprecise, and cancellation, limit, close-count, close-precedence, and repeated-close combinations were incomplete.
- The formatted combined candidate was already **395 authored lines**: `internal/scan/descriptor.go` plus `descriptor_test.go` were 146 lines, while `walker_darwin.go`, `walker_darwin_test.go`, and `walker_other.go` were 249 lines. This leaves no honest room to remediate the combined boundary under the 400-line guard.
- Known failed-candidate file hashes remain: `descriptor.go` `a906a70bf5b1a7e96e426fe3d3ad0aa73d8456baef46421847cd262ac5071c32`; `descriptor_test.go` `8c0baa434e5a7d634c225fe6fd36fa0d4dce706800bcab0bb69ec248a90661b9`; `walker_darwin.go` `61c6fa2a93b87dfd00a931cf97a9c5e147f2ede13ce3644957e9bab0140f7516`; `walker_darwin_test.go` `179647ac956170ab759e6ae55444b806060b9fff3e5c87ca6bcf00c7547e0569`; `walker_other.go` `fff8ba8a3697b8ea0666391c6cacede26d8ff7e7eaeec23534b1b49a11217b0a`. The rejected five-file manifest remains `sha256:d696f9ab3002806d8c840c8fb5e9c9d6c8c42b88ece93c9a5f71c37896b10b33`, and the known pre-rescope task hash is `sha256:d43f307d9052b009fa636f65968a6c3e9765044776974774dbffa020c897b650`.

### Authorized replacement, rollback, and sequence

- Standing user split authorization selects the conservative 12-row plan: WU5C1A contract only, WU5C1B Darwin trusted-root acquisition, and WU5C1C adversarial acquisition triangulation/refactor. WU5C2 enumeration and every later checkbox remain after them unchanged.
- Exactly the four checked combined WU5C1 checkbox rows were removed and replaced by 12 unchecked implementation-owned rows in strict WU5C1A RED/GREEN/TRIANGULATE/REFACTOR → WU5C1B RED/GREEN/TRIANGULATE/REFACTOR → WU5C1C RED/GREEN/TRIANGULATE/REFACTOR order.
- Exact next product action, not performed by this planning append: freshly retain a reproducible baseline because the old `/tmp` baseline is absent; then remove `internal/platform/macos/walker_darwin.go`, `internal/platform/macos/walker_darwin_test.go`, and `internal/platform/macos/walker_other.go`, while retaining `internal/scan/descriptor.go` and `internal/scan/descriptor_test.go` only as the unaccepted WU5C1A candidate.
- After rollback, run the WU5C1A focused contract command against exact named tests before accepting any contract behavior. Do not begin WU5C1B until WU5C1A has independent <=400-line evidence, exact hashes, verification, runtime N/A, and rollback receipt; do not begin WU5C1C or WU5C2 until each predecessor independently passes.

### Forecast, ledger, and preservation proof

- Forecasts: WU5C1A 130–190 additions, 10–30 deletions, 140–220 changes; WU5C1B 250–330 additions, 20–50 deletions, 270–380 changes; WU5C1C 170–250 additions, 10–40 deletions, 180–290 changes. Each is independently capped at 400 authored additions plus deletions.
- Reconciled full forecast: **4,970–6,580 additions, 135–525 deletions, 5,105–7,105 authored changed lines**, excluding generated `go.sum`. Budget risk remains High; chained PRs remain recommended; delivery remains `ask-on-risk` resolved by standing authorization; chain strategy remains `stacked-to-main`; decision before apply remains No.
- Exact ledger: **95 checkbox rows** total; **92 implementation-owned** and **3 parent-owned**. There are **32 checked implementation rows**, **60 unchecked implementation rows**, and **3 unchecked parent rows**, for **32 checked / 63 unchecked overall**.
- Preservation proof: every non-C1 checkbox retains exact text, checked state, relative order, and terminal ownership marker. WU5C2 and all later checkbox rows are byte-for-byte unchanged and remain after the 12 new C1 rows. Only forecast/dependency/coverage prose was reconciled outside the replaced C1 section.
- Post-edit artifact SHA-256 values are not fabricated: the injected file tools expose read/edit persistence but no digest operation. The known pre-rescope task hash, supplied failed receipt prefix, rejected file hashes, and rejected five-file manifest above are the available exact identities.

### Exact next verification action

After the acquisition rollback and fresh baseline retention, run:

```sh
go test ./internal/scan -run '^(TestDescriptorComponents|TestDescriptorRootPolicy|TestDescriptorRelativeChains|TestDescriptorLimits|TestDescriptorFactsAndIdentity|TestDescriptorOwnership|TestDescriptorPathError)$' -count=1
```

Capture genuine RED/GREEN evidence against exact test names; unmatched-regex success is not acceptable. WU5C1B, WU5C1C, WU5C2, and review/lifecycle work remain unauthorized until their stated predecessor and authority gates are satisfied.

---

## WU5C1A receipt — platform-neutral descriptor contract correction

The combined WU5C1 failed evidence `sha256:614522...` is remediated **only partially** by this contract-only WU5C1A slice; WU5C1B, WU5C1C, WU5C2, and later work remain unstarted. RDD is disabled/unmanaged. No review receipt, commit, PR, or token is recorded.

### Completed work and contract

- Removed exactly the rejected incomplete acquisition files: `internal/platform/macos/walker_darwin.go`, `internal/platform/macos/walker_darwin_test.go`, and `internal/platform/macos/walker_other.go`; accepted metadata files were untouched.
- `AbsoluteComponents` accepts only exact clean absolute paths and immutable component copies. `/` remains a valid trusted-home anchor, while `ValidateScanRoot` honestly rejects it as a scan root. `RelativeComponents` is a non-empty immutable name chain and rejects empty, dot, parent, slash, and NUL components.
- `WalkLimits` requires positive depth and positive active-descriptor caps independently, matching future DFS semantics without inventing a `descriptors >= depth` rule. `Identity` and `RootFacts` expose validation/constructors so zero or non-local facts are rejected before acceptance.
- `TrustedRoot` transfers facts and Close capability only, and `DescriptorWalker` has no platform resource or platform-constant surface. `PathError` is constructed with non-empty normalized context and retains sentinel classification through `errors.Is` and `errors.As`.
- Persisted task reconciliation: WU5C1A RED, GREEN, TRIANGULATE, and REFACTOR are visibly `[x]` in `tasks.md`; every WU5C1B-and-later row remains `[ ]`.

### TDD Cycle Evidence

| Task | Layer | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- | --- |
| WU5C1A | Go unit contract | Exact seven-name focused command genuinely failed to compile because root-policy, facts/identity, and PathError contract APIs were absent. | Minimum syscall-free contract passed the focused command. | Added retained-copy, invalid-chain, independent-limit, fake opaque-close, and nested error-context cases; focused PASS. | Reduced the contract to the <=400 authorized slice; gofmt, focused, package, full, and vet PASS. |

### Verification

- Discovered before relying on PASS: `TestDescriptorComponents`, `TestDescriptorRootPolicy`, `TestDescriptorRelativeChains`, `TestDescriptorLimits`, `TestDescriptorFactsAndIdentity`, `TestDescriptorOwnership`, and `TestDescriptorPathError`; the exact focused regex matched all seven.
- `go test ./internal/scan -run '^(TestDescriptorComponents|TestDescriptorRootPolicy|TestDescriptorRelativeChains|TestDescriptorLimits|TestDescriptorFactsAndIdentity|TestDescriptorOwnership|TestDescriptorPathError)$' -count=1` — PASS.
- `go test ./internal/scan -count=1`, `go test ./... -count=1`, and `go vet ./...` — PASS. `gofmt -d` was empty. LSP is N/A because `gopls` is unavailable on PATH. Race/runtime are N/A: this contract slice creates no goroutines, CLI, or platform runtime boundary.

### Baseline, accounting, and rollback

- Fresh retained baseline: `/tmp/osdy-wu5c1a-baseline-1788101729` contains all five pre-correction candidate files plus `tasks.md`, prior `apply-progress.md`, `design.md`, and `docs/architecture.md`. Its `SHA256SUMS` lists no self-hash; manifest SHA-256 is `96a52f5ca9f55502c7a206f34ad7d5c3bcef38439bba6ba3acf57210834613e8`.
- Fresh final two-file manifest SHA-256 is `6f486b4bd88d53a2ba2c21cb117d3420034feb3b03ad9aa5d426bf630a336592`; final hashes are `descriptor.go` `38823c837ab22c3373ab3c148bccf938b13c5361bf290233d77f43dbda362d47` and `descriptor_test.go` `ef645382fa9f6d6a42c5456e95fa16894ee92ab32888ea513c190c77b7e8522b`.
- Relative to the fresh baseline: `descriptor.go` `+51/-5`; `descriptor_test.go` `+83/-12`; deleted walkers `+0/-140`, `+0/-92`, and `+0/-17`; total **+134/-266 = 400 authored changed lines**, exactly within the cap and standing `stacked-to-main` boundary.
- Rollback boundary: restore the two descriptor files from `/tmp/osdy-wu5c1a-baseline-1788101729` and restore the three listed walker files from the same baseline; metadata remains untouched.

### Status and remaining work

- Consumed authoritative OpenSpec status for `read-only-scan-foundation`: artifact store `openspec`, apply route `apply`, no blocked reason, authoritative workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`; edits stayed there. Standing split authorization resolves the high-risk/chained forecast as this WU5C1A `stacked-to-main` slice.
- Exact next unchecked implementation rows are the four WU5C1B rows in `tasks.md`; they are deferred, as are WU5C1C onward and parent-owned lifecycle actions. Independent readiness is `parent-lifecycle`, not authorization to begin WU5C1B in this delegated run.
- Receipt predecessor hash: `46d877fb33f7758a7942f4c31b10b135565c74034aa279cfadb937ef4b30ff79`.

---

## Planning rescope — bounded WU5C1A2 contract correction after independent C1A failure

This is a planning-only append under ordinary SDD/TDD. RDD remains disabled/unmanaged. No product, platform, proposal, specification, design, or documentation file was edited; only `tasks.md` changed and this receipt was appended. No review approval, lifecycle action, implementation acceptance, commit, PR, or release is claimed.

### Independent verifier result

- Independent C1A verification is **FAIL** despite the focused/package/full/vet/format checks and exact +134/-266=400 accounting passing.
- `NewPathError` silently fabricated operation `scan`, path `<unknown>`, and class `ErrInvalidMetadata` for empty/nil inputs. That behavior hides producer defects instead of preserving precise fail-closed context.
- `WalkLimits` lacked exported API documentation and exact edge evidence for depth=0, descriptors=0, `{1,1}`, and asymmetric values. Its required semantics are now explicit: `MaxDepth` is DFS component depth, while `MaxDescriptors` is the simultaneous active capability budget; neither dimension implies the other.
- `TestDescriptorOwnership` directly wrapped a fake root, did not perform meaningful opaque transfer through a fake `DescriptorWalker`, and allowed each repeated `Close` to call the underlying close again. The corrected contract must require idempotent close with one underlying invocation and the same cached typed result.
- Checked WU5C1A remains historical failed-candidate evidence. Because its independent accounting is already exactly 400, standing user split authorization requires adjacent WU5C1A2 rather than modifying cumulative C1A.

### Authorized correction and dependency boundary

- Four unchecked implementation-owned WU5C1A2 rows were inserted immediately after the four checked WU5C1A rows and before the unchanged WU5C1B rows in strict RED → GREEN → TRIANGULATE → REFACTOR order.
- WU5C1A2 may edit only `internal/scan/descriptor.go` and `internal/scan/descriptor_test.go`. It selects the panic-free fail-closed constructor shape `NewPathError(operation, path, class) (*PathError, error)` and adds no platform code.
- WU5C1A2 starts from checked C1A hashes: `descriptor.go` `38823c837ab22c3373ab3c148bccf938b13c5361bf290233d77f43dbda362d47` and `descriptor_test.go` `ef645382fa9f6d6a42c5456e95fa16894ee92ab32888ea513c190c77b7e8522b`.
- The retained C1A baseline remains `/tmp/osdy-wu5c1a-baseline-1788101729`; its no-self-hash manifest SHA-256 is `96a52f5ca9f55502c7a206f34ad7d5c3bcef38439bba6ba3acf57210834613e8`, and the checked C1A final two-file manifest is `6f486b4bd88d53a2ba2c21cb117d3420034feb3b03ad9aa5d426bf630a336592`.
- Apply must retain a fresh two-file WU5C1A2 baseline matching those exact hashes. Rollback restores only those two files; WU5C1B and all platform files remain untouched.
- WU5C1B now depends on independently accepted WU5C1A2. WU5C1C, WU5C2, and every later unit preserve their existing transitive order and remain blocked.

### Forecast, counts, coverage, and preservation

- WU5C1A2 forecast: 90–150 additions, 10–40 deletions, 100–190 authored changed lines, independently capped at 400 and not cumulative with C1A.
- Reconciled total forecast: **5,060–6,730 additions, 145–565 deletions, 5,205–7,295 authored changed lines**, excluding generated `go.sum`. Budget risk remains High; chained PRs remain recommended; delivery remains resolved `ask-on-risk`; chain strategy remains `stacked-to-main`; decision before apply remains No.
- Exact ledger: **99 checkbox rows** total; **96 implementation-owned** and **3 parent-owned**. There are **44 checked implementation rows**, **52 unchecked implementation rows**, and **3 unchecked parent rows**, for **44 checked / 55 unchecked overall**.
- Coverage now assigns fail-closed path-error construction, exact limit semantics/edge matrices, meaningful fake-walker ownership transfer, and idempotent cached typed close to WU5C1A2 before Darwin acquisition.
- Preservation proof: all 95 pre-existing checkbox rows retain exact text, checked state, relative order, and terminal ownership marker. In particular, every WU5C1B, WU5C1C, WU5C2, and later checkbox row is byte-for-byte unchanged and remains in its prior state/order; only the four WU5C1A2 rows were added.
- Exact post-edit `tasks.md` and apply-progress SHA-256 values are unavailable from the injected read/edit-only file tools and are not fabricated. The exact source/baseline hashes above are the reproducible identities available to this planning transaction.

### Exact next implementation row

```markdown
- [ ] **WU5C1A2 RED:** In `internal/scan/descriptor_test.go`, add exact named tests requiring `NewPathError(operation, path, class) (*PathError, error)` to reject empty operation, empty path, and nil class without panic or fabricated operation/path/error; require `WalkLimits` API comments and exact validation cases for depth=0, descriptors=0, `{1,1}`, depth-heavy/descriptor-light, and depth-light/descriptor-heavy values with semantics observable as DFS component depth and simultaneous active capability budget; and require a meaningful fake `DescriptorWalker` to transfer one opaque `TrustedRoot` whose repeated `Close` calls invoke the underlying close once and return the same cached typed result; run the focused command and capture genuine failures against the checked C1A candidate. <!-- sdd-owner: implementation -->
```

Apply resumes only at this WU5C1A2 RED row. Its exact focused command is:

```sh
go test ./internal/scan -run '^(TestDescriptorPathErrorRejectsInvalidContext|TestDescriptorLimitSemantics|TestDescriptorOwnershipTransfer|TestDescriptorCloseIdempotence)$' -count=1
```

WU5C1B, WU5C1C, WU5C2, later implementation, and parent lifecycle actions remain deferred until WU5C1A2 has independent <=400-line evidence, exact hashes, verification, runtime N/A, and two-file rollback evidence.

---

## WU5C1A2 receipt — contract hardening

**Receipt prefix / fresh baseline:** This append began after reading the complete existing receipt. Fresh two-file baseline `/tmp/osdy-wu5c1a2-baseline-1788102961` retains exactly `internal/scan/descriptor.go` and `internal/scan/descriptor_test.go`, whose hashes matched the checked C1A predecessor: `38823c837ab22c3373ab3c148bccf938b13c5361bf290233d77f43dbda362d47` and `ef645382fa9f6d6a42c5456e95fa16894ee92ab32888ea513c190c77b7e8522b`. Its self-excluding `SHA256SUMS` manifest hash is `6f486b4bd88d53a2ba2c21cb117d3420034feb3b03ad9aa5d426bf630a336592`. RDD is disabled/unmanaged; no RDD artifact or action was created.

### Completed implementation tasks and API

- Completed and visibly checked in `tasks.md`: WU5C1A2 RED, GREEN, TRIANGULATE, and REFACTOR. WU5C1B and every later implementation row remain unchecked; all three parent-owned lifecycle rows remain byte-for-byte deferred.
- `NewPathError(operation, path, class)` now returns `(*PathError, error)`. It returns `(nil, ErrInvalidMetadata)` for an empty operation, empty path, or nil class and therefore fabricates neither context nor classification. Valid nested path errors retain outer operation/path while `errors.Is` reaches stable wrapped classes and `errors.As` returns `*PathError`.
- `WalkLimits.MaxDepth` documents the maximum DFS component depth below an acquired root; `MaxDescriptors` documents the maximum simultaneously active opaque capabilities. Each dimension independently requires a positive value.
- `NewTrustedRoot(RootFacts, func() error)` is the reusable `internal/scan` constructor available to future platform adapters without exposing a descriptor, raw handle, or platform constant. It transfers only `TrustedRoot`; its close callback is private capability-boundary state. `sync.Once` invokes that callback once, and repeated `Close` calls return the identical cached error interface, preserving typed `errors.Is`/`errors.As` behavior.

### TDD Cycle Evidence

| Task | Test file / layer | Safety net | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- | --- | --- |
| WU5C1A2 RED/GREEN/TRIANGULATE/REFACTOR | `internal/scan/descriptor_test.go` / platform-neutral unit contract | Pre-change C1A focused suite passed: `go test ./internal/scan -run '^(TestDescriptorComponents | TestDescriptorRootPolicy | TestDescriptorRelativeChains | TestDescriptorLimits | TestDescriptorFactsAndIdentity | TestDescriptorOwnership | TestDescriptorPathError)$' -count=1`. | Added the four exact named tests first; focused command exited 1 at compile time because the old constructor returned one value and `newTrustedRoot` did not exist. | After minimum constructor, comments, and opaque wrapper work, the focused command passed. | Added invalid-context permutations, nested wrapped class/context `Is`/`As`, zero/one/asymmetric limits, one-transfer fake-walker handoff, and success/failing cached-close cases; focused command passed. | `gofmt`, focused/package/full/vet/format checks passed with no behavior expansion. |

### Ownership proof and verification

- `fakeDescriptorWalker` returns its sole opaque `TrustedRoot` exactly once; a second acquisition returns `ErrInvalidMetadata` without incrementing its acquisition counter. The consuming capability observes the supplied `RootFacts` and owns the transferred close call.
- Failing close proof: the callback returns one `*PathError` wrapping `ErrInaccessible`; two `Close` calls run the callback once, return the same error interface, satisfy `errors.Is`, and retain the identical typed error under `errors.As`. Successful close is likewise cached after one callback invocation.
- Focused: `go test ./internal/scan -run '^(TestDescriptorPathErrorRejectsInvalidContext|TestDescriptorLimitSemantics|TestDescriptorOwnershipTransfer|TestDescriptorCloseIdempotence)$' -count=1` -> PASS.
- Scan: `go test ./internal/scan -count=1` -> PASS.
- Full: `go test ./... -count=1` -> PASS (`internal/core`, `internal/platform/macos`, and `internal/scan`).
- Static/format: `go vet ./...` -> PASS; `gofmt -w internal/scan/descriptor.go internal/scan/descriptor_test.go` and final `gofmt -d` -> clean.
- LSP: N/A — `gopls` is not installed. Runtime/race: N/A — this two-file platform-neutral contract owns no executable, scanner runtime, goroutine, or race harness.
- Source boundary scan found no syscall, unix import, build tag, platform constant, or raw-handle surface in the two allowed files.

### Accounting, hashes, rollback, and status

- Independent accounting from the fresh two-file baseline: `descriptor.go` +34/-11 and `descriptor_test.go` +133/-20 = **167 additions + 31 deletions = 198 authored changed lines**, within the 400-line cap.
- Final hashes: `internal/scan/descriptor.go` `f0bbd0127bfd5a3c50faa96daae6202f91898bc292d8b6c7105995d44f9a3e93`; `internal/scan/descriptor_test.go` `5328215c44b189b6ca1194aafbc4d68a03d2ba60b2bb886ce9ebb9b3d2e00df6`; ordered final two-file manifest `sha256:3893ecde8fcea9a84e06f93929757aa79d9a73094173ebf9ee0243bb199c7c0b`.
- Rollback boundary: restore only the two files from `/tmp/osdy-wu5c1a2-baseline-1788102961`; no walker/platform/later-work-unit file is changed.
- Consumed authoritative native status before edits: `artifactStore: openspec`, `applyState: ready`, `nextRecommended: apply`, repo-local workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, allowed root the workspace, strict TDD active, and no blockers. The parent-provided WU5C1A2 sequential `stacked-to-main` slice resolves the high-risk workload guard; no action-context warning occurred.
- Workload / PR boundary: WU5C1A2 only, independent 198 authored lines, no commit, PR, review, RDD action, or lifecycle validation. The next implementation work is WU5C1B and requires parent lifecycle/independent readiness handling.

#### Exact remaining next implementation row

- [ ] **WU5C1B RED:** In Darwin-tagged `internal/platform/macos/walker_darwin_test.go`, add disposable seam tests first for trusted `/` → home → root component acquisition order, exact no-follow/directory/cloexec/nonblock flags, root/home identity-device/locality gates, cancellation before/within acquisition, positive depth/descriptor limits, typed component-specific `errors.Is`/`errors.As`, one-time typed close caching, all-path LIFO close/error precedence, and no raw capability escape; run the focused command and capture a non-zero RED caused only by missing walker APIs. <!-- sdd-owner: implementation -->

---

## WU5C1B blocked-for-split receipt — Darwin trusted-root acquisition

- Status consumed: authoritative OpenSpec `applyState: ready`, repo-local workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, allowed root the workspace, strict TDD enabled; the parent-bound attempt token was continued and settled failed because this independent slice exceeded its line guard. RDD is disabled/unmanaged.
- RED: `walker_darwin_test.go` was added first; the exact focused command failed only for absent `newDescriptorWalker`, `darwinWalker`, `walkerOps`, and `safeDirectoryFlags` symbols.
- GREEN/TRIANGULATE candidate: descriptor-relative `/` → home → root acquisition used exact safe directory flags, descriptor-only `Fstat`/`Fstatfs` facts, post-home device/local/type/nonzero gates, cancellation checks, opaque `NewTrustedRoot` handoff, and typed close handling. The named focused suite passed, including the core dev9 intermediate gate that stopped before the poison child open.
- REFACTOR verification passed: focused, package, `GOOS=linux go test -c` (compile proof only, not executable semantic coverage), full suite, vet, and gofmt. Runtime harness is N/A: no executable/scanner exists and tests used fake descriptor operations only.
- Candidate accounting from the absent three-file baseline was `walker_darwin.go` +203, `walker_other.go` +19, and `walker_darwin_test.go` +212: **434 additions, 0 deletions, 434 authored lines**. This exceeds the mandatory independent 400-line cap.
- Candidate hashes before rollback: `walker_darwin.go` `5bdb97aa9aa8ca1963700d700165a8137eda3a8768575e6844182ca5f2ed8870`; `walker_other.go` `6f2e043c0c3186042436a711e18e73b085a332e7d1d856bc72ea105002f012e4`; `walker_darwin_test.go` `2aa1f1a29d38ec7a514b71c0da9baa184e436a9bb87446721f17c160b65ac279`. Candidate manifest revision: `sha256:cf0fddeaf92c9a53c1a1a1a15a0649a10719f115bedf77344a4f39787ad4bb3e`.
- Rollback completed: all three newly created walker files were removed; no scan-contract, metadata, tasks, or parent-owned lifecycle artifact was changed. All WU5C1B rows remain visibly unchecked.
- Required delivery decision: split WU5C1B into a smaller approved implementation/test work unit or grant an explicit `size:exception`; do not accept this 434-line candidate. Parent-owned lifecycle work remains deferred unchanged.

### TDD Cycle Evidence

| Task | Test layer | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- | --- |
| WU5C1B (discarded) | Darwin fake descriptor-operation unit suite | Exact focused command failed for absent walker symbols. | Exact focused command passed after minimum acquisition code. | Named error, cancellation, limit, post-home dev9 gate, and close cases passed. | Focused/package/Linux compile/full/vet/gofmt passed, but accounting invalidated the candidate. |

---

## Planning rescope — standing split authorization applied to oversized WU5C1B

This is a planning-only append. RDD remains disabled/unmanaged. Only `tasks.md` was edited and this receipt was appended; no product, proposal, specification, design, documentation, implementation acceptance, review, commit, PR, release, or lifecycle action is claimed.

### Failed candidate evidence and rollback boundary

- The discarded WU5C1B candidate passed its focused suite, the full `internal/platform/macos` package, Linux bounded compile proof, `go test ./...`, `go vet ./...`, and gofmt, including the core device-9 intermediate poison gate that proved the poison child did not open.
- Independent evidence is failed at `sha256:ef65933786a8a6cc9b27542f7aa79cd771d3ecde646fa0cf77c6f33cc502c7e3` because the candidate measured **+434/-0**, over the mandatory 400-authored-line guard.
- Candidate file identities before rollback remain: `walker_darwin.go` `5bdb97aa9aa8ca1963700d700165a8137eda3a8768575e6844182ca5f2ed8870`, `walker_other.go` `6f2e043c0c3186042436a711e18e73b085a332e7d1d856bc72ea105002f012e4`, and `walker_darwin_test.go` `2aa1f1a29d38ec7a514b71c0da9baa184e436a9bb87446721f17c160b65ac279`; the candidate manifest revision is `sha256:cf0fddeaf92c9a53c1a1a1a15a0649a10719f115bedf77344a4f39787ad4bb3e`.
- Rollback remains complete: no `walker_darwin.go`, `walker_other.go`, or `walker_darwin_test.go` exists, and every replacement B row is unchecked.

### Authorized split and preserved boundary

- Standing split authorization replaces exactly the four unchecked WU5C1B checkbox rows with eight unchecked implementation-owned rows: WU5C1B1 RED/GREEN/TRIANGULATE/REFACTOR followed by WU5C1B2 RED/GREEN/TRIANGULATE/REFACTOR.
- WU5C1B1 owns core Darwin acquisition and essential acceptance gates at an independent 250–350 authored-change forecast: constructible `NewDescriptorWalker`, root/home ancestry and pre-filesystem validation, exact `/` plus relative component order/flags, component-specific final-descriptor facts, every post-home device/local/type/nonzero gate before deeper open, core typed errors, typed unsupported compile/semantic boundary, opaque `NewTrustedRoot` handoff, and the device-9 poison-child stop. Enumeration and exhaustive cancellation/limits/close matrices remain absent.
- WU5C1B2 depends on accepted B1 and owns exhaustive cancellation checkpoints, depth/descriptor limits, all open/fstat/fstatfs/validation/handoff close counts, primary-plus-cleanup precedence, and idempotent cached typed handle close at an independent 140–240 authored-change forecast.
- WU5C1C's four adversarial atomic-replacement/nested-mount fixture rows remain byte-for-byte unchanged and now depend on accepted B2. Every WU5C2 and later checkbox row remains byte-for-byte unchanged in text, state, order, and terminal ownership marker.

### Forecast, ledger, and exact next action

- Reconciled total forecast: **5,200–6,930 additions, 125–575 deletions, 5,325–7,505 authored changed lines**, excluding generated `go.sum`. Budget risk remains High; chained PRs remain recommended; delivery is resolved `ask-on-risk`; chain strategy remains `stacked-to-main`; decision before apply remains No.
- Exact ledger: **103 checkbox rows** total: **100 implementation-owned** and **3 parent-owned**. There are **48 checked implementation rows**, **52 unchecked implementation rows**, and **3 unchecked parent rows**, for **48 checked / 55 unchecked overall**.
- Preservation proof: only the four current B rows were replaced by eight B1/B2 rows; C1C adversarial fixture rows and every later checkbox remain unchanged. Dependency/coverage/forecast prose was reconciled without product, design, spec, or docs edits.
- Exact next action is WU5C1B1 RED; B2, C1C, C2, and all later work remain blocked on sequential independent acceptance.

```markdown
- [ ] **WU5C1B1 RED:** In `internal/platform/macos/walker_darwin_test.go`, add the exact named tests `TestNewDescriptorWalker`, `TestWalkerRootHomeAncestryPreFSValidation`, `TestWalkerAcquisitionOrderAndFlags`, `TestWalkerComponentFacts`, `TestWalkerEveryPostHomeComponentGate`, `TestWalkerCoreErrorClassification`, `TestWalkerUnsupportedBoundary`, and `TestWalkerIntermediateDevicePoison`; require a constructible `NewDescriptorWalker`, root/home ancestry and pre-filesystem validation, exact `/` plus component relative-open order and `O_RDONLY|O_NOFOLLOW|O_DIRECTORY|O_CLOEXEC|O_NONBLOCK` flags, component-specific final-descriptor facts, device/local/type/nonzero gates after home before every deeper open, core typed missing/symlink/non-directory/inaccessible/device/non-local/invalid classification, a typed unsupported compile/semantic boundary, and device-9 intermediate poison whose child must not open; run the B1 focused command and capture genuine missing-acquisition RED. <!-- sdd-owner: implementation -->
```

---

## WU5C1B1 receipt — core Darwin trusted-root acquisition

**Receipt predecessor:** the cumulative apply-progress artifact was read before this append. This independent B1 receipt remediates failed rolled-back combined B evidence `sha256:ef65933786a8a6cc9b27542f7aa79cd771d3ecde646fa0cf77c6f33cc502c7e3` only for the completed core-acquisition slice; B2 and later work remain unaccepted. RDD is disabled/unmanaged. No runtime token, commit, review receipt, or lifecycle validation is persisted.

### Completed implementation tasks

- Persisted and re-read as checked: WU5C1B1 RED, GREEN, TRIANGULATE, and REFACTOR.
- Added only `internal/platform/macos/walker_darwin.go`, `walker_other.go`, and `walker_darwin_test.go`. Darwin acquisition validates context, home/root ancestry and limits before filesystem operations; it opens `/`, home components, then root suffixes with `O_RDONLY|O_NOFOLLOW|O_DIRECTORY|O_CLOEXEC|O_NONBLOCK`.
- Every acquired component is validated from descriptor `fstat`/`fstatfs`; home establishes the device, and each post-home component is directory, local, identity-valid, and on that device before the next open. Intermediate device poison stops before opening its child.
- Only the final root descriptor facts are transferred through `scan.NewTrustedRoot`; no directory enumeration, ReadDir, content read, path stat, worker, cancellation matrix, limit matrix, or exhaustive close-lifecycle work was added. `MaxDepth` remains below-root traversal capacity and does not reject a long absolute acquisition chain; B1 requires two active capabilities.
- Errors use `scan.NewPathError` with accumulated component paths and stable missing/symlink/non-directory/inaccessible/device/non-local/invalid classes. Cleanup errors are not returned raw and do not replace a primary acquisition error.
- Non-Darwin builds a constructible walker whose acquisition returns a typed `scan.ErrUnsupported` before filesystem work; Linux cross-compilation is the B1 proof.

### TDD Cycle Evidence

| Task | Test file / layer | Safety net | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- | --- | --- |
| WU5C1B1 RED/GREEN/TRIANGULATE/REFACTOR | `internal/platform/macos/walker_darwin_test.go` / Darwin descriptor-operation unit seams | New walker files; existing macOS and scan package tests passed before creation. | Exact focused command exited 1 because `NewDescriptorWalker`, `trustedRootWalker`, `walkerOps`, and `newTrustedRootWalker` were absent. | Exact focused command passed after minimal descriptor-relative acquisition. | The eight exact named tests cover invalid pre-FS inputs, absolute-depth independence, order/flags, final facts, post-home device/local/type/identity gates, typed errors, unsupported boundary, and device-9 poison. | Formatted and reran focused/package/Linux/full/vet checks with no behavior expansion. |

### Verification

- RED: `go test ./internal/platform/macos -run '^(TestNewDescriptorWalker|TestWalkerRootHomeAncestryPreFSValidation|TestWalkerAcquisitionOrderAndFlags|TestWalkerComponentFacts|TestWalkerEveryPostHomeComponentGate|TestWalkerCoreErrorClassification|TestWalkerUnsupportedBoundary|TestWalkerIntermediateDevicePoison)$' -count=1` -> exit 1, missing acquisition symbols only.
- Focused GREEN/TRIANGULATE/REFACTOR: same command -> PASS.
- `go test ./internal/platform/macos -count=1` -> PASS.
- `GOOS=linux go test -c -o /tmp/osdycleaner-macos-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-macos-linux.test` -> PASS; temporary output removed.
- `go test ./... -count=1` -> PASS; `go vet ./...` -> PASS; `gofmt -d` over all three B1 files -> clean. LSP diagnostics were supplied as clean after writes; no separate `gopls` executable was required.
- Runtime harness: N/A — B1 has no executable, scanner, or enumeration boundary; synchronized descriptor-operation seams are its bounded harness.

### Accounting, hashes, and rollback

- Fresh baseline: all three B1 walker files were absent after the oversized +434/-0 combined candidate rollback. This B1 slice is **381 additions + 0 deletions = 381 authored changed lines**, within the 400-line limit.
- SHA-256: `walker_darwin.go` `ba2910bfee5639b492e1ef718c6c05694583ea304e1ae78f779802cb0d4316a4`; `walker_other.go` `32a7f4dc2f588ac226b178b3cec700cb4309ca770eeec6c1fac418bbd05c8b72`; `walker_darwin_test.go` `7bfb70431fe1dbe1158d50a3ec15c74a89c13a63e48f57c7af76d4fad501b562`.
- Ordered product-manifest evidence revision: `sha256:efec28beedda1ae132979df8db0794f7502b130d072c404870a9b7f4ceecf562`. Tasks artifact after checkbox update: `40b39436ad81e4c03cd77e7792ce8e422b651c7ee246a5bda043c5a16fa56b70`.
- Rollback boundary: remove exactly the three B1 walker files. This returns to the accepted WU5C1A2 contract without changing scan contract files, B2, enumeration, or lifecycle artifacts.

### Remaining work and status

- Consumed authoritative native status: change `read-only-scan-foundation`, OpenSpec store, `applyState: ready`, workspace root `/Users/osdy/Documents/GitHub/OsdyCleaner`, allowed root the repository, strict TDD active, and no action-context warnings. The sequential `stacked-to-main` delivery path resolves the high workload forecast for this assigned B1 boundary.
- Ledger after reconciliation: 103 total / 52 checked; 100 implementation / 52 checked; 3 parent / 0 checked. Parent-owned lifecycle actions remain deferred and byte-for-byte unchanged.
- Exact next unchecked B2 rows remain deferred:
  - [ ] **WU5C1B2 RED:** In `internal/platform/macos/walker_darwin_test.go`, add the exact named tests `TestWalkerAcquisitionCancellationCheckpoints`, `TestWalkerAcquisitionLimits`, `TestWalkerAcquisitionHandleClose`, `TestWalkerAcquisitionLifecycle`, `TestWalkerAcquisitionCloseCounts`, and `TestWalkerAcquisitionErrorPrecedence`; require cancellation before/open/fstat/fstatfs/validation/handoff, every depth/descriptor limit edge, exact handle-close counts for success and every open/fstat/fstatfs/validation/handoff failure, primary-plus-cleanup precedence, and idempotent typed handle close with one underlying close and the cached result. <!-- sdd-owner: implementation -->
  - [ ] **WU5C1B2 GREEN:** Make only the minimum `walker_darwin.go`/`walker_other.go` corrections needed for cancellation checkpoints, limit enforcement, complete acquisition ownership cleanup, typed primary-error preservation with secondary close evidence, and idempotent cached typed handle close; preserve B1 order, flags, gates, facts, unsupported behavior, and no-enumeration boundary. <!-- sdd-owner: implementation -->
  - [ ] **WU5C1B2 TRIANGULATE:** Extend the synchronized matrix across every component position and cancellation/limit boundary, all success/open/fstat/fstatfs/validation/handoff/close combinations, repeated successful and failing close calls, and combined primary-plus-cleanup errors; prove exact opens/closes, no deeper work after stop, no capability leak, and stable `errors.Is`/`errors.As` context and precedence. <!-- sdd-owner: implementation -->
  - [ ] **WU5C1B2 REFACTOR:** Run `gofmt -w internal/platform/macos/walker_darwin.go internal/platform/macos/walker_other.go internal/platform/macos/walker_darwin_test.go && go test ./internal/platform/macos -run '^(TestWalkerAcquisitionCancellationCheckpoints|TestWalkerAcquisitionLimits|TestWalkerAcquisitionHandleClose|TestWalkerAcquisitionLifecycle|TestWalkerAcquisitionCloseCounts|TestWalkerAcquisitionErrorPrecedence)$' -count=1 && go test ./internal/platform/macos -count=1 && GOOS=linux go test -c -o /tmp/osdycleaner-macos-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-macos-linux.test && go test ./... -count=1 && go vet ./... && test -z "$(gofmt -d internal/platform/macos/walker_darwin.go internal/platform/macos/walker_other.go internal/platform/macos/walker_darwin_test.go)"`; expect focused/platform/Linux compile/full/vet/format PASS and persist accepted-B1 baseline/final hashes, exact <=400 accounting, runtime N/A, and rollback before WU5C1C. <!-- sdd-owner: implementation -->
- Workload / PR boundary: WU5C1B1 only; independently 381 authored changed lines, sequential stacked-to-main. No commit or PR was created. Return to parent lifecycle; do not start B2 in this delegated scope.

---

## Planning correction — bounded WU5C1B1A after independent B1 verifier FAIL

This is a planning-only append. Only `tasks.md` was edited and this apply-progress artifact was appended; no product code, proposal, specification, design, implementation acceptance, review, commit, PR, release, or lifecycle action is claimed. RDD remains disabled/unmanaged.

### Independent verifier result and retained B1 evidence

- Independent B1 verifier result: **FAIL**. `AcquireRoot` ignores an intermediate `close(current)` error after the next descriptor has opened and passed fact validation, then continues acquisition. That reachable handoff can leak the previous descriptor while deeper acquisition proceeds, so B2 cannot be used to defer the ownership failure.
- B1 remains checked historical implementation evidence at **381 additions + 0 deletions = 381 authored changed lines** against its absent-file baseline. Its focused/platform/Linux compile/full/vet/gofmt checks, flags, mount/device gates, and poison-child stop otherwise passed.
- Exact B1 hashes retained as the B1A baseline: `internal/platform/macos/walker_darwin.go` `ba2910bfee5639b492e1ef718c6c05694583ea304e1ae78f779802cb0d4316a4`; `internal/platform/macos/walker_other.go` `32a7f4dc2f588ac226b178b3cec700cb4309ca770eeec6c1fac418bbd05c8b72`; `internal/platform/macos/walker_darwin_test.go` `7bfb70431fe1dbe1158d50a3ec15c74a89c13a63e48f57c7af76d4fad501b562`. Ordered B1 product manifest: `sha256:efec28beedda1ae132979df8db0794f7502b130d072c404870a9b7f4ceecf562`.

### Inserted correction, dependencies, forecast, and preservation

- Inserted exactly four unchecked implementation-owned WU5C1B1A rows in RED → GREEN → TRIANGULATE → REFACTOR order immediately after checked B1 and before unchanged unchecked B2 rows.
- B1A is bounded to the reachable handoff leak: after a newly opened descriptor is fact-validated, failure closing the superseded descriptor must fail acquisition closed, close the new descriptor exactly once, perform no deeper open, and return no `TrustedRoot`. Earlier primary failures remain primary while every possible cleanup close is attempted. B2 retains the exhaustive cancellation/limit/lifecycle matrix.
- B2 now directly depends on independently accepted B1A; C remains unchanged and depends on B2, therefore transitively on B1A. All B2 and C checkbox rows retain exact text, unchecked state, relative order, and ownership markers. Every other pre-existing checkbox row is likewise preserved.
- Reconciled forecast: **5,270–7,050 additions, 130–600 deletions, 5,400–7,650 authored changed lines**, excluding generated `go.sum`; risk High, chained PRs Yes, delivery `ask-on-risk`, strategy `stacked-to-main`, decision before apply No.
- Reconciled ledger: **107 total rows** = **104 implementation-owned** plus **3 parent-owned**; **52 checked** and **55 unchecked** overall. B1A is independently forecast at 70–120 additions, 5–25 deletions, 75–145 authored changed lines and must be measured at no more than 400 against the exact three-file B1 baseline.

### Exact next implementation row

```markdown
- [ ] **WU5C1B1A RED:** In `internal/platform/macos/walker_darwin_test.go`, add exact `TestWalkerAcquisitionHandoffCloseFailure` and `TestWalkerAcquisitionHandoffPrimaryPrecedence` cases where the next component is opened and fact-validated before closing the previous descriptor fails; require acquisition to return a typed inaccessible close `PathError`, close the new descriptor exactly once, perform no deeper open, and return no `TrustedRoot`; combine an earlier primary failure with cleanup close failures to prove the primary is preserved while every possible close is attempted, focusing only on this reachable handoff leak rather than B2's exhaustive matrix. <!-- sdd-owner: implementation -->
```

Apply must resume at WU5C1B1A RED. B2 and C remain blocked until their sequential predecessor independently passes. RDD remains disabled/unmanaged.

---

## WU5C1B1A receipt — fail-closed acquisition handoff close

**Receipt prefix proof:** the complete prior `apply-progress.md` SHA-256 was `98751225a0011163385175fa06d64bd973af1f2241fa82dff1509c6ff781bb5d` before this append. A fresh exact three-file B1 baseline was retained at `/tmp/osdy-wu5c1b1a-baseline-XKhN67`; its `SHA256SUMS` manifest lists only the three product files and therefore excludes the manifest itself. No runtime attempt token is persisted.

### Completed scope and persisted task evidence

- Completed and visibly checked: WU5C1B1A RED, GREEN, TRIANGULATE, and REFACTOR in `tasks.md`. The artifact was re-read after the checkbox update; all four B1A rows are `[x]`.
- B2 remains untouched and visibly unchecked, beginning with `- [ ] **WU5C1B2 RED:** In \`internal/platform/macos/walker_darwin_test.go\`, add the exact named tests ...`.
- `AcquireRoot` now checks the successful-handoff `close(current)`. If it fails, it closes the already validated `next` descriptor once and returns a typed `scan.PathError` with operation `close`, the relinquished component path, and `scan.ErrInaccessible`; it does not reassign `current`, open deeper, or transfer a `TrustedRoot`.
- Earlier primary failures remain unchanged: the existing fact-error path closes `next` then `current` where possible and returns the earlier fact error without exposing cleanup failures. No B2 cancellation, limit, idempotent-handle, enumeration, scanner, or finalizer scope was added.

### TDD Cycle Evidence

| Stage | Evidence | Result |
| --- | --- | --- |
| RED | Added `TestWalkerAcquisitionHandoffCloseFailure` and `TestWalkerAcquisitionHandoffPrimaryPrecedence` before production changes, then ran `go test ./internal/platform/macos -run '^(TestWalkerAcquisitionHandoffCloseFailure | TestWalkerAcquisitionHandoffPrimaryPrecedence)$' -count=1`. | Exit 1: both handoff-close subcases returned a non-nil `TrustedRoot` and nil error, proving the ignored `close(current)` failure. |
| GREEN | Replaced only the ignored successful-handoff close with checked fail-closed cleanup and typed close context in `walker_darwin.go`; reran the focused command. | PASS. |
| TRIANGULATE | The focused seams cover root-to-first-home and post-home handoff paths, repeated current/next close errors, an earlier `fstat` primary plus cleanup close errors, no deeper opens, no capability escape, and successful transferable handoff close counts. | PASS. |
| REFACTOR | Ran gofmt and the required focused, platform, Linux compile, full-suite, vet, and format checks. | PASS; `gopls` is unavailable. |

### Verification and lifecycle proof

- Focused: `go test ./internal/platform/macos -run '^(TestWalkerAcquisitionHandoffCloseFailure|TestWalkerAcquisitionHandoffPrimaryPrecedence)$' -count=1` -> PASS.
- Platform: `go test ./internal/platform/macos -count=1` -> PASS.
- Linux: `GOOS=linux go test -c -o /tmp/osdycleaner-macos-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-macos-linux.test` -> PASS; the temporary binary is absent afterward.
- Full: `go test ./... -count=1` -> PASS.
- Static/format: `go vet ./...` -> PASS; `test -z "$(gofmt -d internal/platform/macos/walker_darwin.go internal/platform/macos/walker_other.go internal/platform/macos/walker_darwin_test.go)"` -> PASS.
- LSP: `gopls` unavailable; compiler/package tests and vet are clean.
- Runtime harness: N/A — this descriptor-operation seam is the bounded handoff ownership harness; no scanner or CLI/runtime boundary exists. RDD is off.

### Accounting, identity, and rollback

- Independent exact-baseline accounting: `walker_darwin.go` +3/-1, `walker_other.go` +0/-0, `walker_darwin_test.go` +71/-2 = **74 additions + 3 deletions = 77 authored changed lines**, within the 400-line cap.
- Final three-file hashes: `internal/platform/macos/walker_darwin.go` `12daa7e9020a4c7dd8acaaa5c8f61f2838d07474945f97e72af680fbdde51ae1`; `walker_other.go` `32a7f4dc2f588ac226b178b3cec700cb4309ca770eeec6c1fac418bbd05c8b72`; `walker_darwin_test.go` `56ea7bd5fb6a6ef391173e3935b24c8d3e52ae156da74fbafefbc447fcf4c057`.
- Ordered final three-file manifest: `/tmp/osdy-wu5c1b1a-final-manifest.txt`, SHA-256 `e1e28596b6888006684cf4e0091d77a4c1203af9b7da2abe4004ab5f91c999bf`; the manifest excludes itself.
- Rollback boundary: restore the exact three files from `/tmp/osdy-wu5c1b1a-baseline-XKhN67/`, whose hashes are the B1 values recorded in `tasks.md`; this removes only the B1A tests and fail-closed handoff correction.

### Status, workload, and deferred lifecycle

- Consumed authoritative native status: `changeName: read-only-scan-foundation`, `artifactStore: openspec`, `applyState: ready`, `nextRecommended: apply`, workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, `actionContext.mode: repo-local`, and allowed root `/Users/osdy/Documents/GitHub/OsdyCleaner`; no action-context warning or blocked reason applied.
- The active parent-bound attempt was authenticated with its provided continuation token and returned `proceed`. Delivery was the explicitly authorized sequential `stacked-to-main` B1A work-unit slice, at 77 authored lines; no commit, PR, review, receipt approval, validation gate, or lifecycle action was created.
- Parent-owned lifecycle rows remain deferred byte-for-byte. B2 is the next sequential implementation dependency but is outside this delegated scope; return lifecycle ownership to the parent before any further apply.

### WU5C1B1A accounting correction

- The receipt's per-file additions/deletions were transposed by a non-`git --numstat` counting script; the total remains correct.
- Authoritative `git diff --no-index --numstat` against `/tmp/osdy-wu5c1b1a-baseline-XKhN67` reports `walker_darwin.go` **+4/-1**, `walker_other.go` **+0/-0**, and `walker_darwin_test.go` **+70/-2**, for the unchanged **+74/-3 = 77** authored-change total.
- All final product hashes, tests, task checkboxes, rollback boundary, and scope claims in the preceding receipt remain unchanged.

---

## WU5C1B1A baseline integrity repair receipt

- Independent semantic verification is **PASS**; integrity is **FAIL only** because the original baseline directory contained only its manifest.
- For reproducibility, the incomplete original is superseded by the exact hash-verified reconstructed baseline at `/tmp/osdy-wu5c1b1a-reconstructed-baseline`.
- Reconstruction used the inverse of the sole B1A production handoff branch, removed exactly the two B1A tests and their fstat/close recording seam, and cut this receipt before the B1A separator; it makes no Git-history claim.
- Reconstructed hashes: `walker_darwin.go` `ba2910bfee5639b492e1ef718c6c05694583ea304e1ae78f779802cb0d4316a4`; `walker_other.go` `32a7f4dc2f588ac226b178b3cec700cb4309ca770eeec6c1fac418bbd05c8b72`; `walker_darwin_test.go` `7bfb70431fe1dbe1158d50a3ec15c74a89c13a63e48f57c7af76d4fad501b562`; predecessor receipt `98751225a0011163385175fa06d64bd973af1f2241fa82dff1509c6ff781bb5d`.
- Self-excluding reconstructed `SHA256SUMS` hash: `ca64fa5ab222590a46e77ae99f8d80e78a505cb46243ebfa8adc75a968c65cbb`; all identities match the pre-existing B1A receipt.
- Exact `git diff --no-index --numstat` is `walker_darwin.go +4/-1`, `walker_other.go +0/-0`, `walker_darwin_test.go +70/-2`, total `+74/-3 = 77`; no product or task semantics changed.
- Prefix proof: the reconstructed predecessor receipt is byte-exact against the current receipt prefix for 159155 bytes.
- Pre-append receipt SHA-256: `72d8272539d7073f679d5f7839641e6bfbd003181381834bd404bb6aaa5850cc`; the final receipt digest is externally attested after this append and deliberately not self-embedded.
- B2 remains blocked pending one read-only integrity re-verification; no B2 work was performed.
- RDD is disabled/unmanaged.

---

## WU5C1B2 receipt — acquisition cancellation, limits, and lifecycle hardening

**Baseline integrity:** `/tmp/osdy-wu5c1b2-baseline-zkOPZP/SHA256SUMS` was verified before writes: `walker_darwin.go`, `walker_other.go`, `walker_darwin_test.go`, `tasks.md`, and `apply-progress.md` all matched. The immutable manifest hash supplied for that baseline was `0b609bc3cfca53c43812da6c4ebc7dbe5b6cc8910a47cbc8401a7bde5a52ec2c`.

### Completed scope and persisted task evidence

- Completed and visibly checked in `tasks.md`: WU5C1B2 RED, GREEN, TRIANGULATE, and REFACTOR. No WU5C1C, enumeration, scanner, finalizer, CLI, or parent-owned row was changed.
- Acquisition now observes context before component opens, immediately after opens, before and after descriptor `fstat`, before and after `fstatfs`/fact validation, and before capability handoff. Cancellation closes acquired descriptors and returns the stable typed inaccessible acquisition error without deeper work or a capability.
- Acquisition still uses the documented `WalkLimits` semantics: nonpositive depth/descriptors are invalid, and acquisition requires two simultaneous rolling capabilities; no traversal-depth interpretation or contract API changed.
- Every owned descriptor is closed on open, `fstat`, `fstatfs`, validation, device-gate, handoff, cancellation, and constructor failure paths. Earlier typed primary errors are first in `errors.Join`; typed close failures remain secondary evidence without replacing `errors.Is`/`errors.As` primary context.
- The transferred close callback now maps its one underlying close failure to a typed `scan.PathError`; existing `scan.NewTrustedRoot` caches that exact typed result through `sync.Once`, so repeated `Close` calls invoke one underlying close.

### TDD Cycle Evidence

| Stage | Evidence | Result |
| --- | --- | --- |
| RED | Added the exact six B2 named tests before production changes and ran the focused command. | Exit 1: cancellation checkpoints returned a capability, handle close exposed raw `EIO`, and lifecycle paths exposed missing closure/checkpoint behavior. |
| GREEN | Added reachable context checks, complete cleanup with joined secondary close evidence, and typed transferred-close wrapping in `walker_darwin.go`. | Focused command PASS. |
| TRIANGULATE | Exercised cancellation labels, invalid/one descriptor limit edges, open/`fstat`/`fstatfs`/validation/handoff lifecycle cases, close counts, repeated typed close caching, and primary-plus-cleanup precedence. | Focused command PASS. |
| REFACTOR | Applied `gofmt` and ran focused, package, Linux compile, full-suite, vet, and format checks. | PASS; `gopls` is unavailable. |

### Verification and scope checks

- `go test ./internal/platform/macos -run '^(TestWalkerAcquisitionCancellationCheckpoints|TestWalkerAcquisitionLimits|TestWalkerAcquisitionHandleClose|TestWalkerAcquisitionLifecycle|TestWalkerAcquisitionCloseCounts|TestWalkerAcquisitionErrorPrecedence)$' -count=1` -> PASS.
- `go test ./internal/platform/macos -count=1` -> PASS.
- `GOOS=linux go test -c -o /tmp/osdycleaner-macos-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-macos-linux.test` -> PASS; temporary binary absent afterward.
- `go test ./... -count=1` -> PASS; `go vet ./...` -> PASS; required `gofmt -d` check -> clean.
- Runtime harness: N/A — deterministic descriptor-operation seams are the bounded acquisition harness; no scanner or CLI exists. RDD is disabled/unmanaged.
- Scope scan found no `WalkDir`, pathname file reads, mutation, or `filepath` traversal in the two walker product files.

### Accounting, identity, and rollback

- Exact `git diff --no-index --numstat` against the immutable baseline: `walker_darwin.go` **+71/-30**, `walker_other.go` **+0/-0**, and `walker_darwin_test.go` **+151/-56**; total **+222/-86 = 308 authored changed lines**, within the 400-line cap.
- Final SHA-256 hashes: `internal/platform/macos/walker_darwin.go` `d931f9c27e665e06bf19ae15f8b16f21d22d232d6776324f7c828d0d34854dbd`; `walker_other.go` `32a7f4dc2f588ac226b178b3cec700cb4309ca770eeec6c1fac418bbd05c8b72`; `walker_darwin_test.go` `23afee440df8c992556d97ea53c66e84215215b29fc3720c65e95668d0d38fb1`.
- Ordered final three-file manifest digest: `sha256:84d101c47fee75c389fd55ba35f836240eb4fa43ec664a965a1666235c52b08e`.
- Rollback boundary: restore exactly the three allowed files from `/tmp/osdy-wu5c1b2-baseline-zkOPZP/`; this removes only B2 acquisition lifecycle hardening and its tests, preserving B1A and all earlier accepted work.

### Status and deferred lifecycle

- Consumed authoritative native OpenSpec status: `applyState: ready`, `nextRecommended: apply`, workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, `actionContext.mode: repo-local`, and allowed root the repository; no blocked reason or action-context warning applied.
- Workload/PR boundary: WU5C1B2 only, independently measured at 308 authored changed lines under the authorized sequential `stacked-to-main` path. No commit, PR, review, receipt approval, or lifecycle gate was created.
- Parent-owned lifecycle rows remain byte-for-byte deferred. WU5C1C is the next unchecked implementation slice and is not started by this delegation.

---

## WU5C1B2 independent semantic FAIL and correction-plan receipt

- Provenance: independent verifier `subtask_gentle-ai-verify_1788110698244_822f2808` rejected checked B2 semantically; command, hash, and accounting evidence remains historical only and does not establish acceptance or satisfy C's dependency.
- Failures: validation/handoff cancellation cases remap to `fstatfs` and do not prove their named checkpoints or no deeper work; limits cover only `{0,2}`, `{1,0}`, `{1,1}`; lifecycle/close cases omit every component-position × open/fstat/fstatfs/validation/handoff × cleanup-close combination.
- Current failed B2 baseline remains +222/-86 = **308** authored lines with product hashes: `walker_darwin.go` `d931f9c27e665e06bf19ae15f8b16f21d22d232d6776324f7c828d0d34854dbd`, `walker_other.go` `32a7f4dc2f588ac226b178b3cec700cb4309ca770eeec6c1fac418bbd05c8b72`, and `walker_darwin_test.go` `23afee440df8c992556d97ea53c66e84215215b29fc3720c65e95668d0d38fb1`; ordered product manifest digest `84d101c47fee75c389fd55ba35f836240eb4fa43ec664a965a1666235c52b08e`.
- Standing split authorization inserts unchecked, single-writer WU5C1B2A (reachable cancellation + exact limit semantics, forecast 150–260) and WU5C1B2B (exhaustive lifecycle/close matrix, forecast 200–340) before unchanged unchecked C; each is capped at 400 and requires strict TDD plus independent verification before the next slice.
- Parent settlement prerequisite: before B2A writes, create a byte-exact immutable baseline of the current failed-B2 product files plus amended `tasks.md` and this receipt, verify a self-excluding manifest, and record exact file/manifest hashes and the byte-exact pre-append prefix length/hash; repeat an immutable accepted-B2A baseline before B2B.
- Settlement evidence for each correction: genuine named RED, focused/platform/Linux/full/vet/gofmt results, reached-checkpoint and exact no-deeper-operation proof, complete file hashes, `git diff --no-index --numstat` additions+deletions, ≤400 result, rollback restoration proof, independent semantic verdict, and preservation of B1A/no-enumeration.
- Ledger after planning: 115 checkbox rows = 112 implementation-owned + 3 parent-owned; 60 checked + 55 unchecked. Checked B2 remains visibly historical/failed; next action is WU5C1B2A RED after parent baseline settlement. RDD remains disabled/unmanaged; unrelated `.DS_Store` remains untouched and unverified.

---

## WU5C1B2A depth/acquisition task correction provenance receipt

- Authorization and source revision: this append-only receipt records the user-authorized reset/correction from revision `sha256:4c37cb...1458bd`; all receipt bytes before this separator are preserved unchanged.
- Pre-write discovery: accepted `WalkLimits.MaxDepth` evidence in this receipt defines maximum DFS component depth below an acquired root, and `design.md` lines 208–218 consume depth only during descriptor enumeration/descent. The pre-write B2A task text nevertheless required acquisition exact-depth success and one-below failure, which contradicted that accepted meaning.
- Corrected task scope only: `tasks.md` now requires `AcquireRoot` to validate positive `MaxDepth` without charging `/` → home → root acquisition components; acquisition exercises only simultaneous-capability `MaxDescriptors` boundaries (`1` fails before a second capability, `2` exact-boundary success, `64` success) with positive `MaxDepth` `1`/`64` invariance. Exact descendant DFS-depth success and one-below failure are assigned to WU5C2 enumeration/descent.
- Contract evidence: existing `internal/scan/descriptor.go` and `internal/scan/descriptor_test.go` are read-only B2A evidence for zero-value invalidity, positive asymmetric `{1,64}`/`{64,1}` validity, depth-below-root documentation, and simultaneous-capability documentation. B2A writable surfaces are reduced to `internal/platform/macos/walker_darwin.go` and `walker_darwin_test.go`.
- Exact task-ledger effect: 11 non-ledger Markdown rows were corrected; no checkbox, check state, ownership marker, work-unit order, B2B dependency, C dependency, chain strategy, or ≤400 forecast changed. Ledger remains 115 checkbox rows = 112 implementation-owned + 3 parent-owned; 60 checked + 55 unchecked.
- Scope settlement: no Go/product/test/design/spec/proposal/documentation file, commit, review, PR, or lifecycle action was changed or started. Only `openspec/changes/read-only-scan-foundation/tasks.md` was corrected and this receipt was appended.
- Hash settlement still required: before B2A apply, parent must compute and record full SHA-256 hashes for corrected `tasks.md` and appended `apply-progress.md`, then create and verify the already-required byte-exact immutable failed-B2 baseline/self-excluding manifest. The delegated file tools did not expose a digest operation, so no uncomputed full hash is invented here.

---

## WU5C1B2A receipt — corrected cancellation reachability and acquisition capability limits

**Receipt prefix / baseline integrity:** The complete preceding receipt was read before this append. Its byte-exact prefix was 176958 bytes with SHA-256 `29637395d10dee28a3ce72d41e33096462a0039fca223c38f533a1a6ce856ab6`. The immutable corrected baseline `/tmp/osdy-wu5c1b2a-corrected-baseline-SJKWQ2` was verified before any write: all seven manifest entries matched, and its self-excluding `SHA256SUMS` hash was `b49b27c5d409a511bdb11d9abcc8282bd84249425c387ac65059f29a46242c07`. RDD remains off; no commit, PR, review, scanner, finalizer, enumeration, CLI, or lifecycle validation was started.

### Completed implementation tasks and persisted evidence

- Completed and visibly checked in `tasks.md`: WU5C1B2A RED, GREEN, TRIANGULATE, and REFACTOR. The persisted tasks artifact SHA-256 after the updates is `7b2b4c70ff1fcf7561c3700de0e784ec93cfce39f4604f2ce361f8584b29178b`.
- Added only a narrow optional synchronized checkpoint hook to the Darwin descriptor-operation seam. Production always executes real `ctx.Err()` checks at `before`, `after-open`, `after-fstat`, `after-fstatfs`, `after-validation`, and `before-handoff`; the test hook only triggers cancellation at those real checks.
- Cancellation tests prove every component position for each post-operation checkpoint, plus before acquisition and before handoff. Each records the named reached checkpoint, absence of its immediately deeper operation, one close per owned descriptor, nil `TrustedRoot`, and a typed `scan.PathError` retaining `acquire` / `scan.ErrInaccessible` context.
- `MaxDepth` is validated only as positive and is not charged by `/` → home → root acquisition. Existing read-only `TestDescriptorLimitSemantics` remains the contract evidence for zero-value invalidity, asymmetric positive limits, DFS-below-root documentation, and simultaneous-capability documentation.
- Rolling acquisition now treats `MaxDescriptors` as active opaque capability capacity: 1 acquires and fact-validates the current root capability, closes it once, then fails before the second simultaneous open with no capability; 2 is the exact success boundary; 64 succeeds. Positive depths 1 and 64 produce unchanged acquisition.

### TDD Cycle Evidence

| Task | Test file / layer | Safety net | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- | --- | --- |
| WU5C1B2A RED/GREEN/TRIANGULATE/REFACTOR | `internal/platform/macos/walker_darwin_test.go` / Darwin descriptor-operation unit seam | `go test ./internal/platform/macos -count=1` and `go test ./internal/scan -run '^TestDescriptorLimitSemantics$' -count=1` passed before changes. | Exact new tests were written first; focused command exited 1 because `walkerRecord.checkpoint` and `walkerRecord.reached` did not exist. | Added the production checkpoint seam, real context checks, and rolling descriptor-limit stop; focused command passed. | Expanded exact post-operation cancellation coverage through all four acquired component positions and limit rows 1/2/64 with MaxDepth 1/64; focused command passed. | `gofmt`, focused, scan/platform, Linux compile/remove, full, vet, and final format checks passed; no further refactor was needed. |

### Verification

- RED: `go test ./internal/platform/macos -run '^(TestWalkerAcquisitionCancellationReachability|TestWalkerAcquisitionLimitSemantics)$' -count=1` -> exit 1, missing test seam fields only.
- Focused GREEN/TRIANGULATE/REFACTOR: that exact command -> PASS.
- Descriptor contract: `go test ./internal/scan -run '^TestDescriptorLimitSemantics$' -count=1` -> PASS.
- Scan/platform: `go test ./internal/scan ./internal/platform/macos -count=1` -> PASS.
- Linux bounded compile: `GOOS=linux go test -c -o /tmp/osdycleaner-macos-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-macos-linux.test` -> PASS; the output file was confirmed absent.
- Full: `go test ./... -count=1` -> PASS. Static: `go vet ./...` -> PASS. Format: final `gofmt -d` over both writable files was clean.
- LSP: N/A — `gopls` is not installed. Runtime harness: N/A — the deterministic descriptor-operation seam is the bounded unit harness and this scope has no scanner or CLI runtime boundary.
- Scope scan found no enumeration, pathname traversal, file reads, or mutation in the writable walker files. `internal/scan/descriptor.go` and `descriptor_test.go` remain byte-identical to the immutable baseline.

### Accounting, identity, rollback, and status

- Exact `git diff --no-index --numstat` against the immutable corrected baseline: `walker_darwin.go` **+31/-21**, `walker_other.go` **+0/-0**, and `walker_darwin_test.go` **+99/-22**, totaling **+130/-43 = 173 authored changed lines**, within the 400-line cap.
- Final SHA-256: `walker_darwin.go` `beca553e58802696ba9af5d21281ef8707623d25d5ee92af302886e099ab644b`; `walker_other.go` `32a7f4dc2f588ac226b178b3cec700cb4309ca770eeec6c1fac418bbd05c8b72`; `walker_darwin_test.go` `200af6c8e2a830c678e671a7af6f2bc9ee28bef868fff7a0d9031ad8c9e89ab6`. Ordered three-file product manifest `/tmp/osdy-wu5c1b2a-corrected-final-manifest.txt` SHA-256: `sha256:61122ec11b4520ed905867277b9b59aafe11bc7a06aa19b5125dfdd973999813`.
- Rollback boundary: restore exactly `walker_darwin.go`, `walker_other.go`, and `walker_darwin_test.go` from `/tmp/osdy-wu5c1b2a-corrected-baseline-SJKWQ2/`; this removes only WU5C1B2A corrections and tests, preserving B1A and leaving B2B/C/enumeration/scanner/finalizer/CLI untouched.
- Consumed authoritative native OpenSpec status before writes: `artifactStore: openspec`, `applyState: ready`, `nextRecommended: apply`, repo-local workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, allowed root that workspace, strict TDD enabled, and no blocked reason or action-context warning. The forecast is High/chained, but standing `stacked-to-main` authorization explicitly covers this B2A slice.
- Workload / PR boundary: WU5C1B2A only, 173 authored changes under the 400-line cap. B2B remains unchecked and requires independent verification before it may begin; parent-owned lifecycle rows remain deferred byte-for-byte. No approval or acceptance receipt is claimed.

### Immediate remaining implementation work (deferred)

The next sequential B2B rows remain unchecked and are not authorized by this completed B2A slice:

- [ ] **WU5C1B2B RED:** Add exact table-driven `TestWalkerAcquisitionLifecycleComponentMatrix` and `TestWalkerAcquisitionClosePrecedenceMatrix` covering every component position and every success/open/fstat/fstatfs/validation/handoff primary with each possible secondary cleanup-close failure; require exact operation sequence, no deeper operation/open, all possible closes attempted exactly once, no `TrustedRoot`, and stable `errors.Is`/`errors.As` primary context. <!-- sdd-owner: implementation -->
- [ ] **WU5C1B2B GREEN:** Make only lifecycle/cleanup corrections exposed by the exhaustive RED; preserve the first typed primary, join typed secondary close evidence, close every owned descriptor once, cache repeated successful/failing handle-close results, and preserve B1A flags/gates/order and no-enumeration. <!-- sdd-owner: implementation -->
- [ ] **WU5C1B2B TRIANGULATE:** Complete the Cartesian component/operation/cleanup matrix, including success with each close position, primary plus multiple secondary closes, repeated handle close, root→home and post-home handoffs, exact no-deeper-work assertions, and no capability/fact escape. <!-- sdd-owner: implementation -->
- [ ] **WU5C1B2B REFACTOR:** Run `gofmt -w internal/platform/macos/walker_darwin.go internal/platform/macos/walker_darwin_test.go internal/platform/macos/walker_other.go && go test ./internal/platform/macos -run '^(TestWalkerAcquisitionLifecycleComponentMatrix|TestWalkerAcquisitionClosePrecedenceMatrix)$' -count=1 && go test ./internal/platform/macos -count=1 && GOOS=linux go test -c -o /tmp/osdycleaner-macos-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-macos-linux.test && go test ./... -count=1 && go vet ./...`; verify B1A/no-enumeration, record exact hashes/+−/rollback against immutable B2A, and stop for independent verification before C. <!-- sdd-owner: implementation -->

All later unchecked implementation rows and the three parent-owned lifecycle rows remain deferred unchanged; this executor did not initiate B2B or any parent lifecycle action.

---

## WU5C1B2B receipt — exhaustive acquisition lifecycle and close matrix

### Status, scope, and task settlement

- Consumed authoritative native OpenSpec status for `read-only-scan-foundation`: artifact store `openspec`, `applyState: ready`, `nextRecommended: apply`, no blocked reasons, repository-local workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, and no action-context warning. Strict TDD was active from `openspec/config.yaml`.
- The parent-authorized `stacked-to-main` delivery path covers this single B2B slice. The exact five-file B2A baseline `/tmp/osdy-wu5c1b2b-baseline-GsIKcn` and manifest hash `4945f59a62ae37b433788b7ebd672d614d15222218345dc5d22a849509559c1a` were verified before edits.
- Completed and visibly checked: WU5C1B2B RED, GREEN, TRIANGULATE, and REFACTOR. The persisted task artifact was updated immediately after successful verification; WU5C1C and all later implementation rows remain unchecked. Parent-owned lifecycle rows remain unchanged and deferred.

### TDD Cycle Evidence

| Task | Test file / layer | Safety net | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- | --- | --- |
| WU5C1B2B | `internal/platform/macos/walker_darwin_test.go` / Darwin unit syscall seam | `go test ./internal/platform/macos -count=1` PASS (`0.414s`) | The new focused matrix initially exited 1. Exact failure examples were `fd 1 close count=0 trace=[open /]` for an attempted-but-not-acquired open and `scan fstat "/": scan path metadata is invalid` for validation; this corrected the test oracle to distinguish attempted opens and validation-before-`fstatfs`, not production behavior. | The corrected generated lifecycle and precedence matrices passed without a production edit; no artificial failure class or domain error was introduced. | 54 generated cases: lifecycle matrix 24 (five positions × five stages minus final handoff) and close matrix 30 (24 single-cleanup, four reachable cancellation primary plus two secondary closes, two transferred-handle close cases). Cases assert exact trace, path, operation, class, stop boundary, close-once ownership, and typed secondary traversal. | `gofmt`, focused/platform/Linux/full/vet all passed; no behavior-changing refactor was needed. |

### Implementation and verification evidence

- Added only the two required compact table-driven tests and test-seam operation trace. `TestWalkerAcquisitionLifecycleComponentMatrix` spans trusted root, both home components, both suffix/final components, root-to-home and post-home handoffs. `TestWalkerAcquisitionClosePrecedenceMatrix` spans each reachable primary cleanup position, all reachable two-secondary cancellation cleanup branches, and cached repeated transferred-root close for both nil and typed close error results.
- Each generated failure case asserts the first `scan.PathError` operation/path and the required stable class (`ErrInaccessible`, or `ErrInvalidMetadata` for descriptor validation), `TrustedRoot == nil`, no deeper open, exact trace, and exactly-once close for every successfully acquired descriptor. Secondary close evidence is traversed through joined errors so typed close contexts remain discoverable without displacing the first primary context.
- No production correction was warranted: the pre-existing rolling descriptor ownership already closed all acquired descriptors once, joined typed cleanup evidence after the primary, and cached transferred handle close results. No close result is ignored: acquisition close errors are returned or retained as joined secondary evidence, while the transferred capability returns its cached result on every later `Close` call.
- Focused: `go test ./internal/platform/macos -run '^(TestWalkerAcquisitionLifecycleComponentMatrix|TestWalkerAcquisitionClosePrecedenceMatrix)$' -count=1` -> PASS (`0.387s`).
- Platform: `go test ./internal/platform/macos -count=1` -> PASS (`0.266s`).
- Linux: `GOOS=linux go test -c -o /tmp/osdycleaner-macos-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-macos-linux.test` -> PASS; output was removed and confirmed absent.
- Full: `go test ./... -count=1` -> PASS (core `0.669s`, macos `0.252s`, scan `0.456s`). `go vet ./...` -> PASS. `gofmt -d` for all three allowed files -> clean. `gopls` was unavailable, so primary compiler/package diagnostics are the LSP fallback.
- Runtime harness: N/A — this is a deterministic injected Darwin acquisition seam with no scanner, enumeration, CLI, real home, or executable boundary.

### Accounting, hashes, rollback, and remaining work

- Accounted only the three allowed product/test files against the supplied immutable baseline: `walker_darwin.go` +0/-0, `walker_other.go` +0/-0, `walker_darwin_test.go` +266/-1, for **266 additions + 1 deletion = 267 authored changed lines**, within the 400-line cap.
- Final SHA-256: `internal/platform/macos/walker_darwin.go` `beca553e58802696ba9af5d21281ef8707623d25d5ee92af302886e099ab644b`; `walker_other.go` `32a7f4dc2f588ac226b178b3cec700cb4309ca770eeec6c1fac418bbd05c8b72`; `walker_darwin_test.go` `a330a5d6d91c04d8f676e1b4b0650e16665758dffbe6eff164bef82cedddfb0b`.
- Rollback boundary: restore these three files from `/tmp/osdy-wu5c1b2b-baseline-GsIKcn`; this removes only the B2B matrix evidence and leaves B2A behavior intact. No commit, PR, review, receipt approval, scanner, enumeration, finalizer, or CLI work occurred.
- Exact next unchecked implementation row: `- [ ] **WU5C1C RED:** In \`internal/platform/macos/walker_darwin_test.go\`, add synchronized adversarial tests for atomic component replacement with no-follow, \`/home\` device 7 → intermediate device 9 → final device 7 with a poison child that must never open, cancellation/limit checkpoints at every acquisition phase, and every success/open/fstat/fstatfs/validation/close primary-plus-secondary combination; run the focused command and capture any missing safety or lifecycle behavior. <!-- sdd-owner: implementation -->`
- Settlement state: B2B implementation evidence is complete and task checkboxes reconcile; independent semantic verification and all parent lifecycle activity remain parent-owned. Route next to `parent-lifecycle`, not C.

---

## WU5C1B2B independent FAIL and B2B1 verification-only insertion receipt

### Independent verdict and historical reclassification

- Provenance: independent verifier `subtask_gentle-ai-verify_1788112997031_636b270f` returned **FAIL** for checked WU5C1B2B. This receipt records no acceptance, approval, review, commit, PR, implementation, or lifecycle completion.
- The claimed B2B RED was not genuine behavioral RED: its failures were wrong test oracles for attempted opens and validation ordering, and correcting those oracles produced PASS without any production change. The four checked B2B rows remain append-only historical test evidence but are reclassified as verification/triangulation only, never behavioral GREEN and never dependency acceptance.
- The matrix omitted exactly 12 primary × cleanup combinations: fstat, fstatfs, and validation primary failures at component positions 1–4, where both `next` and `current` are owned, but close failure was injected only for `next`. Existing cancellation multi-close cases do not cover an owned-current close secondary for these primaries.

### Reproducible accounting and superseded external claim

- Independently reproducible B2B accounting remains `walker_darwin.go` +0/-0, `walker_other.go` +0/-0, and `walker_darwin_test.go` +266/-1: **266 additions + 1 deletion = 267 authored changed lines**, within 400.
- Independently reproducible current SHA-256 values remain: `internal/platform/macos/walker_darwin.go` `beca553e58802696ba9af5d21281ef8707623d25d5ee92af302886e099ab644b`; `walker_other.go` `32a7f4dc2f588ac226b178b3cec700cb4309ca770eeec6c1fac418bbd05c8b72`; `walker_darwin_test.go` `a330a5d6d91c04d8f676e1b4b0650e16665758dffbe6eff164bef82cedddfb0b`.
- An external subagent summary asserted a final manifest identified only as `sha256:720554…`, but the B2B receipt did not persist that manifest and the independent verifier could not reproduce it. Append-only provenance rejects and supersedes that external manifest claim; no replacement manifest digest is invented. The three file hashes and +266/-1 baseline diff above are the current reproducible identity/accounting evidence.

### Authorized B2B1 plan and gate

- Standing `stacked-to-main` split authorization inserts one bounded sequential WU5C1B2B1 before C, forecast **60–140 additions, 0 deletions, 60–140 authored changed lines**, with only `internal/platform/macos/walker_darwin_test.go` writable.
- B2B1 uses exact `TestWalkerAcquisitionPrimaryCleanupCartesian`, or an explicit generated extension of `TestWalkerAcquisitionClosePrecedenceMatrix`, for all 12 omitted combinations. Every case must assert primary-first `errors.Is`/`errors.As`, a discoverable typed current-close secondary, exactly one attempted close each for next/current, no deeper operation/open, and nil capability/no fact escape; it must also verify multiple secondaries and existing handle-close caching without duplicating B2B.
- Exact focused command: `go test ./internal/platform/macos -run '^(TestWalkerAcquisitionPrimaryCleanupCartesian|TestWalkerAcquisitionClosePrecedenceMatrix)$' -count=1`.
- This is an honest non-behavior VERIFY → TRIANGULATE → REFACTOR unit. If correct added tests PASS on unchanged production, record verification-only PASS and no RED obligation. If any correct expectation FAILS, stop and create a new genuine test-first behavior-correction unit before any production edit; any behavior change still requires strict RED→GREEN.
- Parent settlement gate: before B2B1 writes, retain and verify a byte-exact immutable baseline of the three current walker files plus amended `tasks.md` and this append-only receipt using a self-excluding manifest. After B2B1, record exact hashes and +/− accounting, preserve the +266/-1 B2B baseline evidence, and obtain independent semantic acceptance. Checked failed B2B does not satisfy this gate, and C remains blocked until B2B1 is independently accepted.
- Ledger after amendment: **118 checkbox rows = 115 implementation-owned + 3 parent-owned; 60 checked + 58 unchecked**. The three new unchecked rows are WU5C1B2B1 VERIFY, TRIANGULATE, and REFACTOR; all later checkbox states remain unchanged. RDD remains off, and unrelated repository metadata remains untouched.

### B2B1 planning-ledger count correction addendum

- Supersession: verifier `subtask_gentle-ai-verify_1788113439986_0ca340bc` found that the B2B1 planning receipt's `60 checked + 58 unchecked` statement did not reflect the actual current task rows. That statement remains historical receipt text but is superseded for current ledger counts by this addendum.
- Exact method: count every `openspec/changes/read-only-scan-foundation/tasks.md` line whose Markdown row begins with `- [x]` or `- [ ]`; classify ownership only by its exact terminal `<!-- sdd-owner: implementation -->` or `<!-- sdd-owner: parent -->` marker; then partition each ownership class by `[x]` versus `[ ]` without changing any row.
- Authoritative recomputation: **118 total = 68 checked + 50 unchecked; 115 implementation-owned = 68 checked + 47 unchecked; 3 parent-owned = 0 checked + 3 unchecked**.
- Scope: this correction changes no checkbox state, ownership marker, B2B/B2B1/C semantics, product, test, design, implementation, documentation, commit, review, or lifecycle state.

---

## WU5C1B2B1 receipt — primary × owned-current-close Cartesian verification

**Receipt predecessor:** baseline directory `/tmp/osdy-wu5c1b2b1-baseline-Sv5prz`; its ordered five-file `SHA256SUMS` manifest SHA-256 was verified as `2c0ffe173fe58b5c042b5d288af226213385e39ab6669f2ff5b68d83ec39ec42`. This is an append-only verification-only receipt. No runtime token is recorded.

### Completed task scope and result

- Completed and visibly checked: WU5C1B2B1 VERIFY, TRIANGULATE, and REFACTOR. WU5C1C and all subsequent rows remain unchecked; parent-owned lifecycle rows remain unchanged.
- Added `TestWalkerAcquisitionPrimaryCleanupCartesian` to `internal/platform/macos/walker_darwin_test.go`: generated 3 × 4 = **12 exact cases** for fstat, fstatfs, and validation primary failures at acquisition component positions 1–4.
- Each case injects the secondary close failure on the owned **current** descriptor only, verifies its case-specific current FD/path, requires one close attempt for current and next, and uses the recorded exact operation trace to prove all owned cleanup and no deeper open after the primary failure.
- Every case asserts nil `TrustedRoot` (therefore no trusted capability/facts escape), primary-first classified `scan.PathError` through `errors.Is`/`errors.As`, and discoverable typed current-close secondary evidence. The validation primary is correctly represented by the production contract as operation `fstat` with class `ErrInvalidMetadata`; the initial test expectation labelled it `validation`, was an incorrect oracle, and was corrected without any production change.
- The existing `TestWalkerAcquisitionClosePrecedenceMatrix` was rerun unchanged, retaining its multiple-secondary and transferred-handle repeated-close caching coverage. This is verification-only PASS, not RED/GREEN behavior work.

### Verification evidence

| Check | Result |
| --- | --- |
| Focused: `go test ./internal/platform/macos -run '^(TestWalkerAcquisitionPrimaryCleanupCartesian | TestWalkerAcquisitionClosePrecedenceMatrix)$' -count=1` | PASS |
| Platform: `go test ./internal/platform/macos -count=1` | PASS |
| Linux compile/remove: `GOOS=linux go test -c -o /tmp/osdycleaner-macos-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-macos-linux.test` | PASS; artifact absent after removal |
| Full: `go test ./... -count=1` | PASS (`core`, `platform/macos`, `scan`) |
| Static: `go vet ./...` | PASS |
| Format: `gofmt -w internal/platform/macos/walker_darwin_test.go` and `gofmt -d` | PASS; no diff |
| LSP primary: `gopls check internal/platform/macos/walker_darwin_test.go` | Not run: `gopls` is unavailable in this environment |

Runtime harness: N/A — the synchronized scripted descriptor-operation seam is the bounded acquisition verification harness; no scanner, CLI, enumeration, or executable boundary was started.

### Accounting, hashes, and rollback

- Baseline manifest verification covered all five ordered files: `walker_darwin.go`, `walker_other.go`, `walker_darwin_test.go`, `tasks.md`, and `apply-progress.md`; each recomputed hash matched `/tmp/osdy-wu5c1b2b1-baseline-Sv5prz/SHA256SUMS`.
- Product accounting against that baseline: `walker_darwin.go` +0/-0; `walker_other.go` +0/-0; `walker_darwin_test.go` +41/-0; **41 authored changed lines**, within the 400-line cap.
- Current ordered product hashes: `walker_darwin.go` `beca553e58802696ba9af5d21281ef8707623d25d5ee92af302886e099ab644b`; `walker_other.go` `32a7f4dc2f588ac226b178b3cec700cb4309ca770eeec6c1fac418bbd05c8b72`; `walker_darwin_test.go` `5a4559bb248b0a3cf463a8fa5272bab8f18affc4a052a45111542dfc7f6ad831`.
- The reproducible ordered product-manifest SHA-256, computed from the ordered `sha256sum` rows above, is `22809ffed2e773eda91d666da9519a79d469fbca29c678ea868aa24d8bd8da88`.
- Production integrity: `walker_darwin.go` and `walker_other.go` are byte-identical to the baseline, confirmed by `cmp` and matching SHA-256 values. No production behavior, scanner, finalizer, enumeration, CLI, spec, design, or proposal file changed.
- Rollback boundary: remove only the B2B1 Cartesian test and its test-only rows from `internal/platform/macos/walker_darwin_test.go`, restoring baseline hash `a330a5d6d91c04d8f676e1b4b0650e16665758dffbe6eff164bef82cedddfb0b`; do not change either production walker file.

### Status and settlement

- Consumed authoritative native OpenSpec status: change `read-only-scan-foundation`, `artifactStore: openspec`, `applyState: ready`, repo-local workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, allowed edit root the workspace, strict TDD configured, and no action-context warnings. The standing `stacked-to-main` authorization resolved the high-risk workload gate for this assigned ≤400-line slice.
- Strict-TDD exception is task-defined: this verification-only correct-PASS unit has no fabricated RED/GREEN cycle. The only test-oracle adjustment was validated against the specified typed-error contract; no correct expectation failed and no correction unit is required.
- Workload / PR boundary: WU5C1B2B1 only, 41 authored changed lines, test-only, no commit or PR. No review, receipt approval, verification actor, or delivery lifecycle gate was started.
- Settlement: implementation-owned B2B1 rows are persisted `[x]`; C remains blocked pending parent-owned independent verification and must not be started by this delegated executor.

---

## WU5C1B2B1 independent FAIL and B2B1A verification-correction plan

### Verifier provenance and non-acceptance

- Independent verifier `subtask_gentle-ai-verify_1788113912097_c4b25af5` returned **FAIL** for checked verification-only WU5C1B2B1. This append records no acceptance, implementation, design, documentation, commit, review, PR, or lifecycle completion.
- The sole defect is verification incompleteness: `assertSecondaryCloseEvidence` verifies typed `scan.PathError` close operation/path but not the close error class. It must require `errors.Is(typedClose, scan.ErrInaccessible)` for all 12 current-close Cartesian cases.
- All other B2B1 semantics/checks passed. Preserve its three checked rows, +41/-0 test-only accounting, and hashes as historical attempted evidence independently failed/not accepted; B2B1 satisfies no dependency.

### Authorized minimal B2B1A plan and settlement gate

- Insert sequential WU5C1B2B1A before C with a **10–40 additions, 0 deletions** forecast and only `internal/platform/macos/walker_darwin_test.go` writable. The unit is verification-only VERIFY → REFACTOR; no fake RED/GREEN is permitted.
- VERIFY minimally adds the class assertion to `assertSecondaryCloseEvidence` and reruns exact `TestWalkerAcquisitionPrimaryCleanupCartesian` plus `TestWalkerAcquisitionClosePrecedenceMatrix`. If the correct assertion fails, stop for a separately planned genuine behavior correction and do not edit production.
- REFACTOR runs gofmt, the focused tests, full platform tests, Linux compile/remove, full tests, vet, and a clean format diff. Record exact additions/deletions, final test hash, runtime N/A, rollback, and `cmp`/SHA-256 proof that both production walker files remain byte-identical.
- B2B1 baseline identity remains: `walker_darwin.go` `beca553e58802696ba9af5d21281ef8707623d25d5ee92af302886e099ab644b`; `walker_other.go` `32a7f4dc2f588ac226b178b3cec700cb4309ca770eeec6c1fac418bbd05c8b72`; `walker_darwin_test.go` `5a4559bb248b0a3cf463a8fa5272bab8f18affc4a052a45111542dfc7f6ad831`; ordered production manifest `22809ffed2e773eda91d666da9519a79d469fbca29c678ea868aa24d8bd8da88`.
- C now depends on independently accepted B2B1A. Checked B2, B2B, and B2B1 remain historical failed evidence and cannot satisfy that gate.
- Reconciled planning ledger: **120 total = 68 checked + 52 unchecked; 117 implementation-owned = 68 checked + 49 unchecked; 3 parent-owned = 0 checked + 3 unchecked**. The only new task rows are two unchecked implementation-owned B2B1A rows; every pre-existing checkbox state and ownership marker is preserved.

---

## B2B1A ledger correction addendum

Independent verifier `subtask_gentle-ai-verify_1788114253204_96e71a17` established the actual current ledger as **120 total = 71 checked + 49 unchecked; 117 implementation-owned = 71 checked + 46 unchecked; 3 parent-owned = 0 checked + 3 unchecked**. This addendum supersedes only the planning count recorded at line 1345; history remains append-only. No task row, checkbox state, ownership marker, dependency, semantic requirement, or product artifact changed.

---

## WU5C1B2B1A verification-only progress record

### Scope and status

- Consumed authoritative native OpenSpec status before edits: `changeName: read-only-scan-foundation`, `artifactStore: openspec`, `applyState: ready`, `nextRecommended: apply`, repo-local workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, and allowed edit root that workspace. No action-context warnings or blocked reasons were present.
- Strict TDD is configured, but this task is explicitly verification-only `VERIFY → REFACTOR`; no RED/GREEN claim is made and no production code was edited.
- Baseline `/tmp/osdy-wu5c1b2b1a-baseline-pmtDhA` was present. Every `SHA256SUMS` row verified, and its manifest SHA-256 was confirmed as `7987d355d4f3d4b7b2a7b5a3b2ad3bb16cd7da989af6031ee9252eb91955112d` before the test-only edit.

### Completed implementation tasks and persistence

- Completed and persisted immediately: WU5C1B2B1A VERIFY and WU5C1B2B1A REFACTOR. Both rows are visibly `[x]` in `tasks.md`.
- The exact class check now applies to the selected typed close error itself: `errors.Is(typedClose, scan.ErrInaccessible)`. Thus an aggregate cannot satisfy the class through another joined `ErrInaccessible` while the selected `*scan.PathError` close candidate fails it.
- `internal/platform/macos/walker_darwin_test.go` is the only product/test surface changed. `walker_darwin.go` and `walker_other.go` were confirmed byte-identical with `cmp` against the baseline.

### Verification evidence

| Check | Result |
| --- | --- |
| Pre-edit focused safety net | Both named B2B1/B2B1A focused tests passed before the assertion edit. |
| Post-edit exact focused command | PASS |
| Platform | `go test ./internal/platform/macos -count=1` PASS |
| Linux compile/remove | `GOOS=linux go test -c -o /tmp/osdycleaner-macos-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-macos-linux.test` PASS; output absence confirmed |
| Full | `go test ./... -count=1` PASS |
| Static | `go vet ./...` PASS |
| Format | `gofmt -w internal/platform/macos/walker_darwin_test.go` plus final `gofmt -d` clean |
| LSP | `gopls` unavailable; Go compiler/package diagnostics and `go vet` were clean |

Runtime harness: N/A — the deterministic scripted descriptor-operation seam is the bounded verification harness; no scanner, enumeration, CLI, or executable boundary was started.

### Accounting, hashes, rollback, and boundary

- Exact baseline diff: `walker_darwin.go` +0/-0; `walker_other.go` +0/-0; `walker_darwin_test.go` **+5/-2**; total **7 authored changed lines**, within the 40-line forecast and 400-line cap.
- SHA-256: `walker_darwin.go` `beca553e58802696ba9af5d21281ef8707623d25d5ee92af302886e099ab644b`; `walker_other.go` `32a7f4dc2f588ac226b178b3cec700cb4309ca770eeec6c1fac418bbd05c8b72`; `walker_darwin_test.go` `9d751137039eabce28e444393d40c3378687029174f8c7ad1604d622bf6afdb7`.
- Pre-append ordered five-file evidence manifest (production files, test, tasks, prior apply-progress) SHA-256: `2ca2c88a75f72dadf8d3528bf4bbac909d652860f553a9cf06b3e4505e050bea`.
- Rollback boundary: restore only `internal/platform/macos/walker_darwin_test.go` from the baseline, returning it to `5a4559bb248b0a3cf463a8fa5272bab8f18affc4a052a45111542dfc7f6ad831`; do not modify either production walker file.
- Workload/PR boundary: WU5C1B2B1A only; test-only; no commit, PR, review, validation actor, or parent lifecycle action was started. WU5C1C remains unchecked and must not be started in this scope.

### Deferred lifecycle actions

The three parent-owned lifecycle rows remain byte-for-byte unchanged and unchecked. Independent verification/acceptance and all subsequent lifecycle actions are parent-owned.

### Exact remaining unchecked implementation rows

- [ ] **WU5C1C RED:** In `internal/platform/macos/walker_darwin_test.go`, add synchronized adversarial tests for atomic component replacement with no-follow, `/home` device 7 → intermediate device 9 → final device 7 with a poison child that must never open, cancellation/limit checkpoints at every acquisition phase, and every success/open/fstat/fstatfs/validation/close primary-plus-secondary combination; run the focused command and capture any missing safety or lifecycle behavior. <!-- sdd-owner: implementation -->
- [ ] **WU5C1C GREEN:** In `walker_darwin.go`, make only the minimum corrections exposed by the adversarial RED: gate every accepted post-home component before deeper open, stop poison-child acquisition, preserve component-specific typed context, honor cancellation/limits, and close all ownership paths with typed primary precedence; rerun the focused command and expect PASS without enumeration. <!-- sdd-owner: implementation -->
- [ ] **WU5C1C TRIANGULATE:** Add disposable integration fixtures for ancestor/final symlink and atomic replacement, nested mount simulation returning to the home device, repeated cancellation, each tight-limit permutation, idempotent repeated close, and all close-error combinations; prove exact open/close counts, no poison-target facts, no deeper open after a failed gate, and stable `errors.Is`/`errors.As` classes/context. <!-- sdd-owner: implementation -->
- [ ] **WU5C1C REFACTOR:** Run `gofmt -w internal/platform/macos/walker_darwin.go internal/platform/macos/walker_darwin_test.go internal/platform/macos/walker_other.go && go test ./internal/platform/macos -run '^(TestWalkerAcquisitionAtomicReplacement|TestWalkerAcquisitionNestedMountPoison|TestWalkerAcquisitionCancellationLimitMatrix|TestWalkerAcquisitionCloseMatrix|TestWalkerAcquisitionDisposableFixtures)$' -count=1 && go test ./internal/platform/macos -count=1 && GOOS=linux go test -c -o /tmp/osdycleaner-macos-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-macos-linux.test && go test ./... -count=1 && go vet ./...`; expect PASS, verify no enumeration API/read exists, and record accepted-WU5C1B2B1A/final hashes, <=400 accounting, runtime N/A, and rollback before WU5C2. <!-- sdd-owner: implementation -->
- [ ] **WU5C2 RED:** Add Darwin tests first for enumeration from the already-open descriptor, safe record length/name parsing with malformed records rejected, lexical sibling order, child `openat` parentage and exact flags, enumeration kind/inode versus final-descriptor kind/identity, device/mount/local checks before descent, atomic opened-directory rename/replacement, child disappearance/symlink replacement, exact close order/counts/precedence, exact maximum DFS component depth success below the acquired root and one-below failure during enumeration/descent, and active-FD budget; run the focused command and capture RED caused only by missing enumeration/lifecycle behavior. <!-- sdd-owner: implementation -->
- [ ] **WU5C2 GREEN:** Implement descriptor-owned directory reads and validated record conversion, sort each accepted batch, open child directories relative to the parent with no-follow safe flags, accept only final-descriptor facts, return typed changed/boundary/invalid-record facts, enforce depth/FD bounds, and close child before parent without exposing descriptors to scan callbacks; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU5C2 TRIANGULATE:** Add real disposable fixtures proving an opened directory continues enumerating its original object after pathname rename/replacement and a vanished or symlink-replaced child yields changed evidence with no target descent; combine malformed records, different-device/non-local facts, operation plus close errors, deep trees, and tight FD budgets; rerun the focused command and expect deterministic PASS with safe siblings retained. <!-- sdd-owner: implementation -->
- [ ] **WU5C2 REFACTOR:** Run `gofmt -w internal/scan/descriptor.go internal/scan/descriptor_test.go internal/platform/macos/walker_darwin.go internal/platform/macos/walker_darwin_test.go && go test ./internal/platform/macos -run '^(TestWalkerDescriptorEnumeration|TestWalkerDirectoryRecords|TestWalkerChildOpenParentage|TestWalkerDirectoryReplacement|TestWalkerChildChanges|TestWalkerBoundaryGates|TestWalkerDirectoryLifecycle|TestWalkerDepthFDBudget)$' -count=1 && go test ./internal/platform/macos -count=1 && go test ./... -count=1 && go vet ./...`; expect PASS, record race as N/A because this slice creates no goroutines, and persist accepted-WU5C1/final hashes, <=400 accounting, runtime N/A, and rollback before WU5D1. <!-- sdd-owner: implementation -->
- [ ] **WU5D1 RED:** Add fake-walker tests for canonical built-in validation before platform construction/call, exact five-root serial order, normal/missing/inaccessible/symlink/non-directory/device-boundary/non-local acquisition states, cancellation before a root, poison resolver/platform seams, and five ordered honest incomplete zero-file transitional observations; run the focused command and capture RED caused only by missing scanner consumer APIs. <!-- sdd-owner: implementation -->
- [ ] **WU5D1 GREEN:** Implement the minimum serial scanner consumer with injected home/built-ins/walker, validated absolute home and relative built-in components, one active root, stable acquisition-state mapping, later-root continuation, and five ordered transitional observations; do not add file workers or claim scanned/complete zero-byte coverage; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU5D1 TRIANGULATE:** Add reordered fake definitions, duplicate/extra/unclean roots, root-level combined errors, cancellation after an earlier root, and poison callbacks; prove invalid canonical inputs make zero platform calls, safe later roots continue, no production home is read, and warnings/root states are deterministic; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU5D1 REFACTOR:** Run `gofmt -w internal/scan/scanner.go internal/scan/scanner_test.go internal/scan/descriptor.go && go test ./internal/scan -run '^(TestScannerCanonicalPreflight|TestScannerSerialRoots|TestScannerAcquisitionStates|TestScannerFiveTransitionalObservations|TestScannerNoRealHome)$' -count=1 && go test ./internal/scan -count=1 && go test ./... -count=1 && go vet ./...`; expect PASS, race N/A because roots remain serial without goroutines, and record predecessor/final hashes, <=400 accounting, runtime N/A, and rollback before WU5D2. <!-- sdd-owner: implementation -->
- [ ] **WU5D2 RED:** Add fake-walker/finalizer tests for validated relative components and containment, synchronous directory kind/metadata/identity facts, enumeration-versus-open kind/identity mismatch, disappearance, symlink/special entry, different-device/non-local boundary, malformed/typed platform errors, deterministic warnings, safe siblings/later roots, and exact `partial > boundary_limited > scanned` mapping across all five final observations; run the focused command and capture RED caused only by missing directory-consumer integration. <!-- sdd-owner: implementation -->
- [ ] **WU5D2 GREEN:** Map immutable walker directory/skip/error facts into raw scanner state and the existing finalizer, reject unclean/escaping components, classify changed/disappeared directories as zero-contribution `entry_changed`, retain lower-precedence warnings, continue safe siblings/later roots, and finalize five ordered observations without regular-file bytes; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU5D2 TRIANGULATE:** Add combined changed/symlink/boundary/inaccessible fixtures and callback-order permutations; prove deterministic warning order, correct partial/boundary/root statuses, no target facts, no unsafe descent request, safe continuation, and no complete claim while regular-file evidence is absent; rerun and expect PASS without real-home inspection. <!-- sdd-owner: implementation -->
- [ ] **WU5D2 REFACTOR:** Run `gofmt -w internal/scan/scanner.go internal/scan/scanner_test.go internal/scan/finalize.go internal/scan/finalize_test.go && go test ./internal/scan -run '^(TestScannerDirectoryFacts|TestScannerDirectoryChanges|TestScannerDirectoryBoundaries|TestScannerRootStatusMapping|TestScannerFiveFinalObservations)$' -count=1 && go test ./internal/scan -count=1 && go test ./... -count=1 && go vet ./...`; expect PASS, race N/A because no goroutines exist, and record accepted-WU5D1/final hashes, <=400 accounting, runtime N/A, and rollback before WU5E1. <!-- sdd-owner: implementation -->
- [ ] **WU5E1 RED:** Add synchronized scanner and Darwin tests for fixed worker/job/result counts, bounded capacities and backpressure, validated relative identity-chain jobs, trusted-root private-anchor reopen, rolling ancestor descriptors, exact no-follow flags, ancestor/final identity and device/local checks, enumeration-versus-open kind/identity/disappearance, regular-only logical/allocation facts, no content reads, and exact close counts; run both focused commands and capture RED caused only by the missing file pipeline/reopen APIs. <!-- sdd-owner: implementation -->
- [ ] **WU5E1 GREEN:** Implement the minimum fixed worker and bounded channel pipeline plus platform relative reopen: stop absolute descendant inspection, reopen each component from the private root anchor, close rolling descriptors, accept only unchanged same-boundary regular final-descriptor facts, and return changed/boundary/error results with zero bytes otherwise; rerun both focused commands and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU5E1 TRIANGULATE:** Add blocked enqueue/result cases, ancestor/final rename or symlink replacement, disappearance, non-regular entries, different-device/non-local facts, metadata/close failures, deep identity chains, and queue-pressure permutations; prove bounded goroutines/FDs, no target/content facts, no dropped accepted result, and deterministic zero contribution for changed entries; rerun the focused commands and `go test -race ./internal/scan -run '^(TestScannerBoundedFileJobs|TestScannerFileBackpressure|TestScannerRegularOnlyFacts|TestScannerFileEntryChanges)$' -count=1`, expecting PASS. <!-- sdd-owner: implementation -->
- [ ] **WU5E1 REFACTOR:** Run `gofmt -w internal/scan/descriptor.go internal/scan/scanner.go internal/scan/scanner_test.go internal/platform/macos/walker_darwin.go internal/platform/macos/walker_darwin_test.go && go test ./internal/scan -run '^(TestScannerBoundedFileJobs|TestScannerFileBackpressure|TestScannerRegularOnlyFacts|TestScannerFileEntryChanges)$' -count=1 && go test ./internal/platform/macos -run '^(TestWalkerRelativeFileReopen|TestWalkerFileIdentityChain|TestWalkerRegularFileFacts|TestWalkerFileReopenLifecycle)$' -count=1 && go test -race ./internal/scan -count=1 && go test ./... -count=1 && go vet ./...`; expect PASS and record exact predecessor/final hashes, <=400 accounting, runtime N/A, and rollback before WU5E2. <!-- sdd-owner: implementation -->
- [ ] **WU5E2 RED:** Add channel-synchronized tests for worker/result completion permutations, same-identity hard links, changed/disappeared file warnings, cancellation before enqueue/after acceptance/while blocked, stopped new jobs, result draining, worker joining, no leaked goroutines, and deterministic finalizer attribution/warnings; run the focused command and capture RED caused only by missing completion/cancellation integration. <!-- sdd-owner: implementation -->
- [ ] **WU5E2 GREEN:** Implement cancellation-aware ownership that stops enqueueing and worker acceptance, closes job input once, drains every accepted result, joins the fixed workers, canonically sorts facts before existing hard-link finalization, and preserves completed roots plus active/later cancellation states without scheduling-dependent attribution; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU5E2 TRIANGULATE:** Add blocked job/result channels, repeated cancellation, mixed boundary/partial/change errors, worker completion permutations, cross-area hard links, and close/error combinations; prove cancellation precedence, no discarded accepted fact, no post-cancel open/job, deterministic warnings/totals, regular-only bytes, and no goroutine/descriptor leak; rerun the focused command and `go test -race ./internal/scan -run '^(TestScannerFileCompletionPermutations|TestScannerHardLinkFinalization|TestScannerCancellationStopDrainJoin|TestScannerNoPipelineLeaks|TestScannerDeterministicWarnings)$' -count=1`, expecting PASS. <!-- sdd-owner: implementation -->
- [ ] **WU5E2 REFACTOR:** Run `gofmt -w internal/scan/scanner.go internal/scan/scanner_test.go internal/scan/finalize.go internal/scan/finalize_test.go && go test ./internal/scan -run '^(TestScannerFileCompletionPermutations|TestScannerHardLinkFinalization|TestScannerCancellationStopDrainJoin|TestScannerNoPipelineLeaks|TestScannerDeterministicWarnings)$' -count=1 && go test -race ./internal/scan -count=1 && go test ./... -count=1 && go vet ./...`; expect PASS, record accepted-WU5E1/final hashes, <=400 accounting, runtime N/A, and rollback. WU6 remains blocked until this independent receipt passes. <!-- sdd-owner: implementation -->
- [ ] **WU6 RED:** Add deterministic channel-synchronized tests for cancellation before a root, during an active metadata call, after an earlier root, and alongside partial/boundary warnings; add entry/path limit tests and run the focused command to record failures for missing stop/state behavior without sleep-based timing. <!-- sdd-owner: implementation -->
- [ ] **WU6 GREEN:** Implement context checks before new work, immediate enqueue stop, bounded active-call completion, channel drain/join, cancellation warnings, skipped later roots, limit warnings, scan-policy fields, and exact `cancelled > partial > boundary_limited > scanned` derivation; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU6 TRIANGULATE:** Add reordered worker completion, blocked-result-channel, unknown-required-metadata, combined limit/boundary/error, and repeated-cancel cases; run `go test -race ./internal/scan -run '^(TestScannerStatusPrecedence|TestScannerEntryLimit|TestScannerPathBudget|TestScannerCancellation|TestScannerNoLeakedWork|TestScannerCompletionOrderStable)$' -count=1` and expect PASS with no race, leaked work, discarded finalized fact, or newly started root. <!-- sdd-owner: implementation -->
- [ ] **WU6 REFACTOR:** Run `gofmt -w internal/scan/scanner.go internal/scan/scanner_test.go internal/scan/finalize.go internal/scan/finalize_test.go && go test -race ./internal/scan -count=1`; expect PASS and verify partial/cancelled aggregate estimates remain incomplete. <!-- sdd-owner: implementation -->
- [ ] **WU7 RED:** Add fixed-snapshot report tests for complete, partial, and cancelled outcomes, exact top-level/nested field order, typed bytes/booleans, explicit unknowns, `related_path: null`, `[]` collections, one trailing newline, all five roots, and text/JSON fact equivalence; run the focused command and record missing-renderer/golden failures. <!-- sdd-owner: implementation -->
- [ ] **WU7 GREEN:** Implement ordered DTO projection in `internal/report/json.go` and canonical text rendering in `internal/report/text.go`, then run `go test ./internal/report -run '^(TestJSONReport|TestTextReport|TestReportDeterminism|TestReportEquivalence)$' -update -count=1`; inspect all six allowed goldens and rerun the focused command without `-update`, expecting PASS. <!-- sdd-owner: implementation -->
- [ ] **WU7 TRIANGULATE:** Add discovery/collection permutations and assertions for byte-identical JSON, JSON parseability, schema/outcome/completeness, risks, bases, warnings, prominent partial/cancelled text, APFS clone/snapshot/compression/hard-link caveats, CoreSimulator manual-review wording, and no deletion-safety/reclaim promise; rerun without `-update` and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU7 REFACTOR:** Run `gofmt -w internal/report/json.go internal/report/text.go internal/report/report_test.go && go test ./internal/report -count=1`; expect PASS without golden updates, map-backed serialization, absolute home paths, wall-clock/host/user/random fields, or presentation-derived scan facts. <!-- sdd-owner: implementation -->
- [ ] **WU8 RED:** Add direct model tests for pure `Init`, canonical category movement, h/j/k/l and arrows, page scrolling, Tab details, resize, q/Escape/Ctrl-C quit, complete/partial/cancelled views, and terminal rejection; run the focused command and record missing-model/dependency failures. <!-- sdd-owner: implementation -->
- [ ] **WU8 GREEN:** Pin approved compatible v2 dependency versions in `go.mod`, implement snapshot-only model/update/view/runner files, use only the viewport component, and rerun the focused command; expect PASS with no scan command from `Init` and no command after quit. <!-- sdd-owner: implementation -->
- [ ] **WU8 TRIANGULATE:** Assert prominent “Read-only scan,” “Manual review required,” and “Estimate—not guaranteed reclaimable space” wording, cancellation/incompleteness before facts, canonical area/warning agreement, and absence of cleanup/select/action/confirm/delete/remove/Trash affordances while allowing the required estimate disclaimer; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU8 REFACTOR:** Run `gofmt -w internal/tui/model.go internal/tui/view.go internal/tui/run.go internal/tui/model_test.go && go mod tidy && go test ./internal/tui -count=1`; expect PASS, inspect generated `go.sum` in the complete receipt, and verify the model stores no operation intent or alternate totals. <!-- sdd-owner: implementation -->
- [ ] **WU9 RED:** Add injected scanner/renderer/viewer/terminal/stream tests for default text, formats, positional or path-like input, unsupported format, noninteractive TUI, one scan call, complete/partial/cancelled outcomes, unsupported environment, global init/finalize/render/write/viewer failures, and fixture JSON; run the focused command and record missing-command/orchestration failures. <!-- sdd-owner: implementation -->
- [ ] **WU9 GREEN:** Pin Cobra in `go.mod`, implement silenced thin command validation and one-scan orchestration in `internal/cli`, render text/JSON fully before stdout, keep snapshot warnings in-band and diagnostics on stderr, and centralize codes 0/1/2/3/4/130; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU9 TRIANGULATE:** Add poison-scanner cases proving invalid path/format/noninteractive TUI fail before construction or traversal, root warnings never become code 1, cancellation outranks code 3, global failure emits no purported snapshot, JSON has one document/newline and no decoration, write/viewer failure never rescans, and signal cancellation starts no follow-up operation; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU9 REFACTOR:** Run `gofmt -w internal/cli/command.go internal/cli/run.go internal/cli/command_test.go && go mod tidy && go test ./internal/cli -count=1`; expect PASS, execute the stated runtime harness evidence command, and account for generated `go.sum` without counting its lines as authored. <!-- sdd-owner: implementation -->
- [ ] **WU10 RED:** Before creating `cmd/osdy/main.go`, run `go test ./cmd/osdy -count=1`; record the expected non-zero “directory/package not found” result as the missing production entry-point boundary. <!-- sdd-owner: implementation -->
- [ ] **WU10 GREEN:** Implement `cmd/osdy/main.go` with `signal.NotifyContext`, stop signal delivery, production dependency construction, delegation to `internal/cli`, and `os.Exit` only; run `gofmt -w cmd/osdy/main.go && go test ./cmd/osdy ./internal/cli -count=1` and expect PASS with no scan/report logic in `main`. <!-- sdd-owner: implementation -->
- [ ] **WU10 TRIANGULATE:** Run the disposable runtime script below and expect exit 0, empty stderr, exactly one parseable JSON document, schema 1, complete outcome, canonical five area IDs, no absolute temporary-home path in JSON, and identical before/after fixture metadata. <!-- sdd-owner: implementation -->
- [ ] **WU10 REFACTOR:** Run the final acceptance sequence below in order: explicit-file `gofmt`, all focused package tests, report JSON parseability/equivalence tests without `-update`, `go test ./...`, `go vet ./...`, then the disposable runtime script again; expect every command to pass/no-diagnostic and stop rather than weakening checks if any result fails. <!-- sdd-owner: implementation -->

### Exact deferred parent-owned rows

- [ ] Before resumed apply, record the already selected sequential `stacked-to-main` chain strategy and the WU2A → WU2B split, confirm no further product/delivery choice is pending, retain the previously selected module import identity, and preserve one-writer execution without assuming commits or PRs can be created. <!-- sdd-owner: parent -->
- [ ] After each applied work unit, inspect its apply-progress receipt, confirm authored additions plus deletions are at most 400 with generated `go.sum` excluded, confirm the focused command/result, runtime evidence or explicit N/A, complete changed-file identity, and rollback boundary, then authorize the next dependent unit through SDD apply/verify authority; if a unit exceeds 400 authored lines, stop and split it again before further implementation writes. <!-- sdd-owner: parent -->
- [ ] After WU10, record final SDD completion evidence against the normative coverage matrix, required commands, disposable runtime script, authored-line accounting, complete changed-file identity, and rollback boundaries; mark implementation complete only when that evidence is satisfied. <!-- sdd-owner: parent -->

---

## B2B1A receipt table correction

- Replaced exactly the malformed pre-edit focused safety-net row with pipe-free accurate prose.
- Baseline `144d9d6b09e88c63cb09696db670972e2cc3f2bcecd0090690bb9939dcbfb230`; verifier `subtask_gentle-ai-verify_1788114734368_3c0fe9de` identified the third-cell pipe.
- No product, test, or task semantic change occurred.
- The final artifact hash is deliberately not embedded and remains externally computable.

---

## Planning reset — WU5C1C verification-first split after independent FAIL

### Reset authorization, provenance, and non-acceptance

- Explicit user reset authorization was consumed for planning only; its authorization token is deliberately not persisted. No Go/product/test/design/docs/commit/review work occurred.
- Independent verifier `subtask_gentle-ai-verify_1788116436728_35d0b50f` returned **FAIL** for the WU5C1C candidate. Failed evidence is `sha256:a126a6cb18d2d79f905714e7bc613ec8378efd11dc662245c18cbf114848cd99`; it remains historical evidence with no acceptance and satisfies no dependency.
- The rejected candidate's fake ELOOP did not perform synchronized real ancestor/final pathname replacement and its `t.TempDir()` was not used by fake open facts. Its mount poison stopped at device 9 and never modeled or observed final device 7 or a poison child. Cancellation occurred before descriptor-limit distinctions, and the close matrix omitted required success/open/close and primary-secondary dimensions.
- The failed candidate added **123 authored lines** and was restored completely. The exact C baseline is restored with no C row checked and no production acceptance claimed.
- Settlement of this reset must remediate the failed evidence revision `sha256:a126a6cb18d2d79f905714e7bc613ec8378efd11dc662245c18cbf114848cd99`; no later receipt may treat that revision or the historical C rows as accepted evidence.

### Exact restored baseline identities

- `internal/platform/macos/walker_darwin.go`: `beca553e58802696ba9af5d21281ef8707623d25d5ee92af302886e099ab644b`.
- `internal/platform/macos/walker_other.go`: `32a7f4dc2f588ac226b178b3cec700cb4309ca770eeec6c1fac418bbd05c8b72`.
- `internal/platform/macos/walker_darwin_test.go`: `9d751137039eabce28e444393d40c3378687029174f8c7ad1604d622bf6afdb7`.
- The verified pre-B2B1A baseline manifest remains `7987d355d4f3d4b7b2a7b5a3b2ad3bb16cd7da989af6031ee9252eb91955112d`; the pre-append five-file evidence manifest remains `2ca2c88a75f72dadf8d3528bf4bbac909d652860f553a9cf06b3e4505e050bea`. These historical manifest identities are retained without claiming they digest the newly amended planning artifacts.
- Exact post-plan `tasks.md` and `apply-progress.md` digests are unavailable through the injected file-only tools and are not fabricated.

### Authorized sequential replacement

- The four historical unchecked WU5C1C RED/GREEN/TRIANGULATE/REFACTOR rows are superseded by six unchecked implementation-owned rows in this exact order: WU5C1C1 VERIFY, TRIANGULATE, REFACTOR, then WU5C1C2 VERIFY, TRIANGULATE, REFACTOR.
- WU5C1C1 is test-only, forecast **160–280 authored lines**, and uses actual `t.TempDir()` component trees. Synchronization-only wrappers must delegate facts to real Darwin `openat`, `fstat`, `fstatfs`, and `close`; they pause only to permit real opened-ancestor/final rename/replacement and must prove no-follow plus retained-descriptor binding. A separate deterministic mount/device seam must explicitly declare and inspect the full **7→9→7** script while the acquisition trace consumes only 7→9; it must prove the poison child never opens and scripted final facts are never accepted.
- WU5C1C2 is test-only, forecast **120–240 authored lines**, and cross-products every reachable acquisition checkpoint with exact-boundary and one-too-tight descriptor limits so cancellation and limit outcomes are distinguishable. It then covers only remaining success/open/fstat/fstatfs/validation/handoff/close primary-secondary dimensions not already accepted, reusing helpers without superficial duplicate cardinality.
- Both units use VERIFY → TRIANGULATE → REFACTOR. Correct PASS on unchanged production is accepted verification and has no fabricated RED obligation. Any correct failure stops immediately and requires a separate independently bounded strict RED → GREEN behavior-correction unit before any production edit; production edits are unauthorized in C1 and C2.
- WU5C1C1 depends on independently accepted WU5C1B2B1A. WU5C1C2 depends on independently accepted WU5C1C1. Existing WU5C2 enumeration now depends on independently accepted WU5C1C2. The historical failed C candidate satisfies none of these gates.

### Forecast, ledger, and preservation

- Reconciled total forecast: **5,770–8,000 additions, 150–660 deletions, 5,920–8,660 authored changed lines**, excluding generated `go.sum`. Budget risk remains High; chained PRs remain recommended; delivery remains resolved `ask-on-risk`; chain strategy remains `stacked-to-main`; decision before apply remains No.
- Exact ledger: **122 total rows = 71 checked + 51 unchecked**. **119 implementation-owned = 71 checked + 48 unchecked**; **3 parent-owned = 0 checked + 3 unchecked**.
- Exact new unchecked rows are WU5C1C1 VERIFY/TRIANGULATE/REFACTOR and WU5C1C2 VERIFY/TRIANGULATE/REFACTOR. Every other checkbox state and ownership marker is preserved. The three parent-owned rows and WU5C2/later rows remain unchecked.
- No review actor, commit, PR, lifecycle gate, implementation acceptance, product edit, test edit, design edit, or documentation edit was started by this planning reset.

### Settlement needs and next action

1. Retain and verify a fresh immutable C1 baseline matching all three restored source/test hashes above plus amended `tasks.md` and this append-only receipt.
2. Apply WU5C1C1 only, record exact additions/deletions and final hashes, prove both production files byte-identical, and obtain independent semantic acceptance.
3. Apply WU5C1C2 only after accepted C1, record exact non-duplicative matrix coverage, additions/deletions, final hashes, production byte-identity, and obtain independent semantic acceptance.
4. Mark failed evidence `sha256:a126a6cb18d2d79f905714e7bc613ec8378efd11dc662245c18cbf114848cd99` remediated only after both C1 and C2 independently pass; only then may WU5C2 enumeration begin.

---

## Current C replan ledger correction

Verifier `subtask_gentle-ai-verify_1788117151301_2565ddb1` established the current ledger as **122 total = 73 checked + 49 unchecked; 119 implementation-owned = 73 checked + 46 unchecked; 3 parent-owned = 0 checked + 3 unchecked**. This supersedes only stale current C replan count claims; all historical counts remain unchanged. No checkbox row, task semantic, or product change occurred.

---

## WU5C1C1 verification-only receipt — real descriptor binding and mount poison rejection

- Baseline directory `/tmp/osdy-wu5c1c1-baseline-flCQDB` was verified: its `SHA256SUMS` digest is `d877e90184c39c55ebe18e335ecb29325a17af8eaf08dde4e521c9d9a6fefc62` and its declared walker hashes match the restored production files.
- Added only `internal/platform/macos/walker_darwin_test.go`: **+226/-0 = 226 authored test-only lines**, within the 400-line C1 cap. Production walker files are byte-identical to the baseline by `cmp` and SHA-256.
- Completed persisted implementation rows: WU5C1C1 VERIFY, TRIANGULATE, and REFACTOR. C2 and every later row remain unchecked; parent-owned lifecycle rows are deferred unchanged.

### TDD Cycle Evidence

| Stage | Evidence | Result |
| --- | --- | --- |
| VERIFY | Exact four-test focused command | PASS on unchanged production; no fabricated RED. |
| TRIANGULATE | Real ancestor/final replacements plus symlink targets and separate poison script | PASS; no content read or enumeration. |
| REFACTOR | Platform, Linux compile, full suite, vet, and gofmt checks | PASS; `gopls` unavailable. |

- Focused command passed: the four named C1 tests only.
- Platform package, Linux cross-compile with temporary artifact removal, full suite, vet, and gofmt verification all passed.
- Real mutation/identity evidence: each fixture is a physical `t.TempDir()` canonicalized before root-to-home-to-root acquisition; mutation runs synchronously after real ancestor or final `walkerOpen`. The original path is renamed and replaced by a symlink; retained descriptors retain the original root identity, record no replacement-target facts, use exact safe flags for every real open, and close every acquired FD instance exactly once.
- Mount poison evidence: declared device script is `[7,9,7]`; acquisition consumes only `7,9`, returns `ErrDeviceBoundary` at `poison`, performs four opens/fstat/fstatfs calls through the rejected intermediate, never opens `child`, and never consumes or accepts final device-7 facts.
- Runtime harness: N/A — disposable Darwin descriptor fixtures are the designated boundary. Rollback: remove only the four named C1 tests and synchronization-only helpers, restoring baseline test hash `9d751137039eabce28e444393d40c3378687029174f8c7ad1604d622bf6afdb7`.
- Final hashes: walker_darwin `beca553e58802696ba9af5d21281ef8707623d25d5ee92af302886e099ab644b`; walker_other `32a7f4dc2f588ac226b178b3cec700cb4309ca770eeec6c1fac418bbd05c8b72`; walker test `eb8ce26e6974646ffe067834390b095c0764e1897ebab43d118d151d1c360b95`. Evidence manifest hash `445e1a0214e6f8ff0a8b0f6046a5451bc3a177ab9b706fb82a2792906470ecc1`. Updated tasks hash `28265c67020fdda7bfe8f7d59829abfa943382e7b55c594a02e0c1b271e57368`.
- Structured status consumed: authoritative OpenSpec change `read-only-scan-foundation`, apply state ready, repo-local allowed workspace, strict TDD active, and resolved stacked-to-main delivery. No action-context warnings. Workload boundary: C1 only; no commit, review, enumeration, or production edit.

---

## Planning correction — minimal WU5C1C1A after independent C1 verifier FAIL

### Scope, provenance, and non-acceptance

- This is planning-only. No Go product/test implementation, proposal, specification, design, documentation, commit, review, PR, release, or lifecycle action was performed. The supplied authorization token is deliberately not persisted.
- Independent verifier `subtask_gentle-ai-verify_1788117808925_f5d42ba5` returned **FAIL** for checked WU5C1C1. Its three checked rows remain historical verification-attempt evidence only; C1 is independently failed/not accepted and satisfies no dependency.
- C1 exact accounting is **+233/-0 = 233 authored test-only lines**, superseding only the receipt's inaccurate +226/-0 claim. All production and check results otherwise passed.
- The real ancestor/final replacement evidence is incomplete because it does not, after rename, stat/fstat the renamed-original pathname/FD and explicitly compare that `TrustedRoot` retained identity equals the renamed original identity and differs from the replacement target identity.
- The mount poison evidence is incomplete because it infers stopping from calls rather than recording and asserting consumed device trace exactly `[7,9]` and explicitly proving the final scripted `7` slot is unconsumed.
- No acceptance is recorded for C1, and no historical receipt or hash is rewritten.

### Authorized minimal correction and sequence

- Inserted WU5C1C1A immediately after checked failed C1 and before C2. It is test-only, limited to `internal/platform/macos/walker_darwin_test.go`, forecast **20–80 additions, 0 deletions**, and contains exactly two unchecked implementation-owned rows: VERIFY then REFACTOR. No fake RED/GREEN/TRIANGULATE or production edit is authorized.
- C1A retains the exact focused C1 command and four names: `TestWalkerAcquisitionRealAncestorReplacement`, `TestWalkerAcquisitionRealFinalReplacement`, `TestWalkerAcquisitionRealNoFollowBinding`, and `TestWalkerAcquisitionMountPoisonScript`.
- VERIFY must stat the renamed-original pathname and replacement target and fstat retained descriptor evidence after each real rename-replacement, then explicitly assert retained `TrustedRoot` identity equals renamed-original identity and differs replacement-target identity. The mount test must record consumed devices exactly `[7,9]`, retain declared script `[7,9,7]`, and explicitly assert the final `7` slot/index remains unconsumed.
- Any genuine correction failure stops immediately with no product edit or acceptance and requires a separately planned behavior correction. C2 now depends only on independently accepted C1A; checked failed C1 cannot satisfy that gate.

### Hashes, accounting, rollback, and mappings

- Failed C1 source identities remain: `walker_darwin.go` `beca553e58802696ba9af5d21281ef8707623d25d5ee92af302886e099ab644b`; `walker_other.go` `32a7f4dc2f588ac226b178b3cec700cb4309ca770eeec6c1fac418bbd05c8b72`; `walker_darwin_test.go` `eb8ce26e6974646ffe067834390b095c0764e1897ebab43d118d151d1c360b95`.
- Historical C1 evidence manifest remains `445e1a0214e6f8ff0a8b0f6046a5451bc3a177ab9b706fb82a2792906470ecc1`, and historical C1-updated tasks hash remains `28265c67020fdda7bfe8f7d59829abfa943382e7b55c594a02e0c1b271e57368`; neither is claimed to digest this planning correction.
- C1A rollback removes only its explicit post-rename identity comparisons and consumed-slot assertions, restoring failed C1 test hash `eb8ce26e6974646ffe067834390b095c0764e1897ebab43d118d151d1c360b95`; both production hashes must remain byte-identical.
- Coverage mappings now distinguish failed C1 from corrective C1A and assign retained-versus-renamed/replacement identity proof plus exact consumed/unconsumed device trace to C1A before C2 matrices.
- Reconciled forecast: **5,790–8,080 additions, 150–660 deletions, 5,940–8,740 authored changed lines**, excluding generated `go.sum`; budget risk High, chained PRs Yes, resolved `stacked-to-main`, decision before apply No.
- Exact ledger after insertion: **124 total = 73 checked + 51 unchecked**; **121 implementation-owned = 73 checked + 48 unchecked**; **3 parent-owned = 0 checked + 3 unchecked**. Only the two C1A rows are new; C1's three rows remain checked historical failed evidence and every other row retains state, order, and ownership.
- Exact post-edit `tasks.md` and append-only `apply-progress.md` digests are unavailable through the injected file-only tools and are not fabricated.

### Settlement needs and next action

1. Retain a fresh immutable C1A baseline matching the three failed-C1 file hashes above and the amended task/receipt artifacts.
2. Apply only WU5C1C1A VERIFY and REFACTOR; record exact C1A additions/deletions, final test hash, unchanged production hashes, runtime N/A, focused/platform/Linux/full/vet/format results, and rollback evidence.
3. Obtain independent semantic acceptance of C1A. If identity or consumed-slot assertions genuinely fail, stop and plan a separate correction without production edits in C1A.
4. Begin C2 only after independently accepted C1A; C1 remains failed/not accepted regardless of its checked historical rows.

---

## Current C1A planning ledger arithmetic correction

The current C1A planning ledger is **124 total = 76 checked + 48 unchecked; 121 implementation-owned = 76 checked + 45 unchecked; 3 parent-owned = 0 checked + 3 unchecked**. This supersedes only the stale current C1A task-count claim; all earlier ledger counts remain preserved as historical snapshots.

Method: counted every Markdown checkbox row in the current `tasks.md`, partitioned rows by terminal `sdd-owner` marker, and then partitioned each ownership group by `[x]` versus `[ ]`. This planning correction changed no checkbox row, state, order, ownership marker, task semantic, product file, or test file.

---

## WU5C1C1A apply progress — verification-only correction

- Scope completed: only `WU5C1C1A VERIFY` and `WU5C1C1A REFACTOR`; both persisted implementation-owned rows are `[x]`. No C1A behavioral correction, C2/later task, production edit, commit, review actor, receipt, or approval was created.
- Changed file: `internal/platform/macos/walker_darwin_test.go` only. The four existing real C1 tests now obtain post-mutation identities with real `syscall.Stat` from renamed-original and replacement paths and assert `TrustedRoot == renamed-original != replacement`; the mount fixture records scripted fact requests and asserts declared `[7,9,7]`, consumed `[7,9]`, and an unconsumed final `7`.
- Focused safety net before test edits: exact four-test command PASS (`0.367s`). Focused verification after correction: PASS (`0.408s`). Package PASS (`0.245s`); Linux cross-compile PASS; full PASS (core `0.438s`, macos `0.253s`, scan `0.620s`); `go vet ./...` PASS; `gofmt -d` clean. `gopls` was unavailable, so LSP checking could not run.
- Runtime: N/A — disposable Darwin `t.TempDir()` descriptor fixtures are the bounded harness; no executable/scanner runtime exists.
- Baseline `/tmp/osdy-wu5c1c1a-baseline-8Nxfij` manifest verified with all five entries OK. C1A accounting against its baseline is `+44/-0` authored test lines, within the `20–80` forecast; superseded C1 accounting is `+233/-0`.
- Production byte identity: `cmp` against the baseline passed for both production walkers. SHA-256: `walker_darwin.go` `beca553e58802696ba9af5d21281ef8707623d25d5ee92af302886e099ab644b`; `walker_other.go` `32a7f4dc2f588ac226b178b3cec700cb4309ca770eeec6c1fac418bbd05c8b72`; final test `walker_darwin_test.go` `46398466d805bd97dcf3e7ef499df5fd9aa4d224beebb1aaf3a31adeb9b9046c`.
- Workload / PR boundary: WU5C1C1A only, test-only `+44/-0`, sequential `stacked-to-main`; no commit or PR. Roll back by restoring the baseline test hash `eb8ce26e6974646ffe067834390b095c0764e1897ebab43d118d151d1c360b95` or removing the C1A-only assertions/helpers; production remains byte-identical.
- Structured status consumed: authoritative OpenSpec `read-only-scan-foundation`, `applyState: ready`, `artifactStore: openspec`, repo-local action context rooted at `/Users/osdy/Documents/GitHub/OsdyCleaner` with that workspace allowed; no action-context warning. Strict-TDD config is active, with this task-specific verification-only exception `VERIFY → REFACTOR` and no fabricated RED/GREEN.
- Deferred lifecycle actions: all parent-owned rows are unchanged. Remaining C2 rows (not started):
  - [ ] **WU5C1C2 VERIFY:** Add `TestWalkerAcquisitionCheckpointDescriptorLimitMatrix` and `TestWalkerAcquisitionRemainingLifecycleMatrix` in `internal/platform/macos/walker_darwin_test.go`; cross every reachable before/open/fstat/fstatfs/validation/handoff checkpoint with exact-boundary and one-too-tight descriptor limits, asserting whether cancellation or limit wins and proving the outcomes are not masked. Add only missing success/open/fstat/fstatfs/validation/handoff/close primary-secondary rows not already accepted. Correct PASS is accepted verification; any correct failure stops for a separate strict RED → GREEN correction plan with no production edit in this unit. <!-- sdd-owner: implementation -->
  - [ ] **WU5C1C2 TRIANGULATE:** Reuse existing trace/typed-error/close helpers to prove exact operation/path/class, open/close counts, no deeper operation, nil capability on failure, transferred ownership after success, primary-first context with discoverable secondaries, and distinct cancellation-versus-limit results; document exclusions for the 12 already accepted B2B1A Cartesian rows so cardinality is substantive rather than duplicated. <!-- sdd-owner: implementation -->
  - [ ] **WU5C1C2 REFACTOR:** Run `gofmt -w internal/platform/macos/walker_darwin_test.go`, the exact focused command, `go test ./internal/platform/macos -count=1`, `GOOS=linux go test -c -o /tmp/osdycleaner-macos-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-macos-linux.test`, `go test ./... -count=1`, `go vet ./...`, and clean `gofmt -d`; record exact accounting/test hash, runtime N/A, rollback, and production `cmp`/hash byte-identity before independent acceptance authorizes WU5C2 enumeration. <!-- sdd-owner: implementation -->

---

## WU5C1C2 apply progress — verification-only checkpoint and lifecycle matrices

- Completed persisted implementation rows: `WU5C1C2 VERIFY`, `TRIANGULATE`, and `REFACTOR`; each is visibly `[x]` in `tasks.md`. WU5C2's four enumeration rows remain `[ ]`; no WU5C2 code, test, enumeration, review, commit, or production edit was started.
- Changed file: `internal/platform/macos/walker_darwin_test.go` only. Authored accounting against `/tmp/osdy-wu5c1c2-baseline-zUC4SZ` is **+112/-0 = 112** test-only lines, below the 400-line cap.
- Baseline verification: every `SHA256SUMS` entry passed; its canonical five-file manifest was `419c3b84cce5b63eceaf9191261ac07189dda05454cf5f86d6b75dc0d2d562bb`.
- `TestWalkerAcquisitionCheckpointDescriptorLimitMatrix` contains **44** generated rows: the 22 reachable checkpoint-position pairs (`before` root; four after-operation checkpoints at five positions each; `before-handoff` final) crossed with descriptor maxima 2 and 1. With maximum 2, each reached checkpoint cancels first and returns typed `acquire`/`ErrInaccessible`; with maximum 1, only root `before`/after-operation checkpoints are reached and cancel first, while every later/unreachable checkpoint returns typed `acquire` `/`/`ErrInvalidMetadata`. Thus cancellation wins when its checkpoint occurs before the limit check, and the one-capability limit wins when cancellation is unreachable. Every failure has nil root, exact one-close ownership for every opened descriptor, and no deeper work; `MaxDepth: 64` is held invariant and uncharged.
- `TestWalkerAcquisitionRemainingLifecycleMatrix` adds exactly two non-duplicative primary/secondary rows: open failure at `/home/alice` with a typed previous-current close secondary at `/home`, and handoff current-close primary at `/home/alice/cache` with typed transferred-next close secondary at `/home/alice/cache/final`. It asserts `errors.Is`, `errors.As`, primary-first context, exact close ownership, and nil transfer. The accepted 12 B2B1A `fstat`/`fstatfs`/`validation × current-close` rows, existing multiple-secondary rows, and idempotent transferred-handle close-cache coverage are explicitly excluded; accepted C1A replacement and mount-trace coverage is unchanged.
- Focused command: `go test ./internal/platform/macos -run '^(TestWalkerAcquisitionCheckpointDescriptorLimitMatrix|TestWalkerAcquisitionRemainingLifecycleMatrix|TestWalkerAcquisitionPrimaryCleanupCartesian|TestWalkerAcquisitionClosePrecedenceMatrix)$' -count=1` — PASS (`0.253s` on final run).
- Platform package: `go test ./internal/platform/macos -count=1` — PASS (`0.230s`). Linux compile: `GOOS=linux go test -c -o /tmp/osdycleaner-macos-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-macos-linux.test` — PASS. Full suite: `go test ./... -count=1` — PASS (core `0.630s`, macos `0.264s`, scan `0.438s`). `go vet ./...` — PASS. `gofmt -d internal/platform/macos/walker_darwin_test.go` — clean. `gopls` was unavailable, so no LSP check could run.
- Production byte identity: `cmp` against the baseline passed. SHA-256: `walker_darwin.go` `beca553e58802696ba9af5d21281ef8707623d25d5ee92af302886e099ab644b`; `walker_other.go` `32a7f4dc2f588ac226b178b3cec700cb4309ca770eeec6c1fac418bbd05c8b72`. Final test SHA-256: `43065a984ada9d107e954f970762d064bc0f80719c343fcf396a549c47fbc164`; updated tasks SHA-256: `90c4d265a8d4ee3f9816090672aa0ed9bf7c7d4b3d8083042fea1464e53a02f3`; bounded evidence manifest: `sha256:635b485e51986f9610a943d77eeb6d1f8423ce0ba44a3ae9cf83f328858013f3`.
- Runtime: N/A — deterministic scripted and disposable Darwin descriptor tests are the bounded harness; no executable, real home, enumeration, network, shell, or production mutation ran. Rollback: restore the baseline `walker_darwin_test.go` or remove only the two C2 tests and their helpers; production files remain untouched.
- Workload / PR boundary: WU5C1C2 only, test-only +112/-0, sequential `stacked-to-main`; no commit or PR. Structured status consumed: authoritative OpenSpec `read-only-scan-foundation`, `applyState: ready`, repo-local allowed workspace, strict TDD active, no action-context warning. No Engram memory tools were injected, so no Engram mirror was possible. Parent-owned lifecycle rows remain byte-for-byte deferred. Route to parent lifecycle for independent acceptance before WU5C2; do not begin enumeration.

---

## Planning correction — minimal WU5C1C2A after independent C2 verifier FAIL

### Scope, provenance, and non-acceptance

- This settlement is planning-only. No Go product or test implementation, proposal, specification, design, general documentation, commit, review, PR, release, or lifecycle action was performed. The supplied authorization token was not persisted.
- Independent verifier `subtask_gentle-ai-verify_1788119235934_81a16f3d` returned **FAIL** for checked WU5C1C2. Its three checked rows remain historical verification-attempt evidence only; C2 is independently failed/not accepted and satisfies no dependency.
- Exact C2 accounting is **+116/-0 = 116 authored test-only lines**, superseding only the receipt's inaccurate +112/-0 claim. All production and check results otherwise passed.
- Every reached-cancellation matrix row must explicitly assert its exact operation trace, exact open count, immediate next operation absent, and immediate next open absent; class and close assertions alone are insufficient.
- The lifecycle inventory must explicitly enumerate every primary-secondary dimension and map each cell to accepted WU5C1B1A, WU5C1B2A, WU5C1B2B1A, or WU5C1C1A evidence, or to exactly one of C2's two new lifecycle rows. Broad comments do not prove omission-free coverage.
- No acceptance is recorded for C2, and no historical receipt or hash is rewritten.

### Authorized minimal correction and sequence

- Inserted WU5C1C2A immediately after checked failed C2 and before WU5C2 enumeration. It is test-only, limited to `internal/platform/macos/walker_darwin_test.go`, forecast **20–80 additions, 0 deletions**, and contains exactly two unchecked implementation-owned rows: VERIFY then REFACTOR. No fake RED, GREEN, TRIANGULATE, production edit, enumeration, design, docs, commit, or review is authorized.
- C2A retains the exact focused C2 command: `go test ./internal/platform/macos -run '^(TestWalkerAcquisitionCheckpointDescriptorLimitMatrix|TestWalkerAcquisitionRemainingLifecycleMatrix|TestWalkerAcquisitionPrimaryCleanupCartesian|TestWalkerAcquisitionClosePrecedenceMatrix)$' -count=1`.
- Any genuine correction failure stops immediately with no production edit or acceptance and requires a separately planned behavior correction. WU5C2 now depends independently on accepted C2A; checked failed C2 cannot satisfy that gate.

### Hashes, accounting, rollback, mappings, and ledger

- Failed C2 source identities remain: `walker_darwin.go` `beca553e58802696ba9af5d21281ef8707623d25d5ee92af302886e099ab644b`; `walker_other.go` `32a7f4dc2f588ac226b178b3cec700cb4309ca770eeec6c1fac418bbd05c8b72`; `walker_darwin_test.go` `43065a984ada9d107e954f970762d064bc0f80719c343fcf396a549c47fbc164`.
- Historical C2 bounded evidence manifest remains `sha256:635b485e51986f9610a943d77eeb6d1f8423ce0ba44a3ae9cf83f328858013f3`, and historical C2-updated tasks hash remains `90c4d265a8d4ee3f9816090672aa0ed9bf7c7d4b3d8083042fea1464e53a02f3`; neither is claimed to digest this planning correction.
- C2A rollback removes only its explicit trace/open/no-deeper assertions and exhaustive lifecycle inventory mapping, restoring failed C2 test hash `43065a984ada9d107e954f970762d064bc0f80719c343fcf396a549c47fbc164`; both production hashes must remain byte-identical.
- Coverage mappings and the chain now distinguish failed C2 from corrective C2A and assign omission-free checkpoint/lifecycle proof to C2A before WU5C2.
- Reconciled forecast is **5,810–8,160 additions, 150–660 deletions, 5,960–8,820 authored changed lines**, excluding generated `go.sum`; budget risk High, chained PRs Yes, resolved `stacked-to-main`, decision before apply No.
- Exact ledger after inserting two rows is **126 total = 81 checked + 45 unchecked; 123 implementation-owned = 81 checked + 42 unchecked; 3 parent-owned = 0 checked + 3 unchecked**. Only the two C2A rows are new; C2's three rows remain checked historical failed evidence and every other row retains state, order, and ownership.
- Exact post-edit artifact digests are unavailable through the injected file-only tools and are not fabricated.

### Exact scope and settlement

1. Apply only WU5C1C2A VERIFY and REFACTOR in `internal/platform/macos/walker_darwin_test.go`.
2. Record exact C2A additions/deletions, final test hash, unchanged production hashes, runtime N/A, focused/platform/Linux/full/vet/format results, and rollback evidence.
3. Obtain independent semantic acceptance of C2A; if a corrected assertion fails, stop without production changes.
4. Begin WU5C2 only after independently accepted C2A. C2 remains failed/not accepted regardless of its checked historical rows.

---

## WU5C1C2A apply progress — verification-only trace and lifecycle correction

- Completed persisted implementation rows: `WU5C1C2A VERIFY` and `WU5C1C2A REFACTOR`; both are visibly `[x]` in `tasks.md`. WU5C2 remains `[ ]`; no production, enumeration, design, documentation, commit, or review edit was made.
- Baseline was verified before edits: every `/tmp/osdy-wu5c1c2a-baseline-5kfBXn/SHA256SUMS` entry matched and the supplied baseline-manifest SHA-256 `972d9ddac9f890a035f33aeaef15568bea245dd9afd80795a81a41b4716c9522` matched `SHA256SUMS`.
- Test-only accounting against that baseline is **+70/-10 = 80 authored changed lines** in `internal/platform/macos/walker_darwin_test.go`; it is within the C2A 20–80 forecast. Historical failed C2 accounting remains corrected at **+116/-0**, superseding the inaccurate +112/-0 claim.
- Reached cancellation evidence now asserts exact trace, exact open count, and explicit absent next `open`/`fstat` for every `before`, `after-open`, `after-fstat`, `after-fstatfs`, `after-validation`, and `before-handoff` matrix row. Tight-limit rows separately assert `ErrInvalidMetadata`, `acquire /`, one open, and stopped trace `[open /, fstat 1, fstatfs 1, close 1]`, proving limit precedence.
- Executable inventory keys: `success`, `open`, `fstat`, `fstatfs`, `validation`, `handoff/close-current`, `cleanup-close-next/current`, `multiple-cleanup`, `transferred-close-success`, `transferred-close-failure`, `transferred-close-cache`. It fails on unknown, empty-evidence, duplicate, or missing keys; `fstat`, `fstatfs`, and `validation` explicitly identify the accepted 12 B2B1A current-close rows as excluded (four each), while C2's `open` and `handoff/close-current` remain case-specific lifecycle rows.
- Verification PASS: focused C2 command; `go test ./internal/platform/macos -count=1`; `GOOS=linux go test -c -o /tmp/osdycleaner-macos-linux.test ./internal/platform/macos` followed by removal; `go test ./... -count=1`; `go vet ./...`; and clean `gofmt -d internal/platform/macos/walker_darwin_test.go`. `gopls` was unavailable, so LSP diagnostics could not run.
- Production byte identity: `cmp` passed against baseline. SHA-256 remains `walker_darwin.go` `beca553e58802696ba9af5d21281ef8707623d25d5ee92af302886e099ab644b` and `walker_other.go` `32a7f4dc2f588ac226b178b3cec700cb4309ca770eeec6c1fac418bbd05c8b72`; final test SHA-256 is `95e56aaef496274eb1f45e9d23ad3260a07b44f1432d9093d26b22644d4338f1`.
- Runtime: N/A — deterministic scripted descriptor tests are the bounded harness; no executable, real-home access, enumeration, network, shell, or production mutation ran. Rollback: restore the baseline `walker_darwin_test.go` (or remove only C2A trace/inventory assertions), restoring failed-C2 test hash `43065a984ada9d107e954f970762d064bc0f80719c343fcf396a549c47fbc164`; production files remain unchanged.
- Workload / PR boundary: WU5C1C2A only, test-only 80 authored changed lines, sequential `stacked-to-main`; no commit or PR. Structured status consumed: authoritative OpenSpec `read-only-scan-foundation`, `applyState: ready`, repo-local workspace `/Users/osdy/Documents/GitHub/OsdyCleaner` in allowed edit roots, strict-TDD active with this task's verification-only VERIFY → REFACTOR exception, and no action-context warnings. No Engram memory tools were injected, so no Engram mirror was possible. Parent-owned lifecycle rows remain deferred byte-for-byte.
- Remaining implementation tasks include the four unchecked WU5C2 enumeration rows; independent acceptance remains parent-owned before WU5C2 begins.

### TDD Cycle Evidence

| Task | Test file / layer | Safety net | VERIFY | REFACTOR |
| --- | --- | --- | --- | --- |
| WU5C1C2A | `internal/platform/macos/walker_darwin_test.go` / Darwin scripted unit harness | Focused C2 suite PASS before edits | PASS after exact trace/open/no-deeper and inventory assertions | gofmt and all required checks PASS |

---

## WU5C2 apply progress — blocked for sequential split

- Structured status produced because no native status JSON was supplied: `{change: read-only-scan-foundation, artifactStore: openspec, authoritative: true, applyState: ready-before-apply, actionContext: {mode: implementation, allowedEditRoots: [workspace]}, strictTDD: true}`. No unsafe action-context or edit-root warning was found.
- The six-file baseline `/tmp/osdy-wu5c2-baseline-o82OM0` was verified with `sha256sum -c SHA256SUMS`: all six entries passed. The four allowed product/test files match that baseline byte-for-byte: `descriptor.go`, `descriptor_test.go`, `walker_darwin.go`, and `walker_darwin_test.go`.
- Before a substantial write, the exact eight-test contract was forecast at roughly 240–300 production/contract lines plus 340–480 meaningful Darwin test/fixture/seam lines, or 580–780 authored lines. Record parsing, descriptor-bound DFS, identity/boundary gates, and the required real rename/replacement evidence cannot honestly fit the independent 400-line cap together.
- No RED test, production code, scanner, finalizer, D1/later, CLI, review, commit, or task checkbox was changed. The four WU5C2 implementation rows remain visibly unchecked; parent-owned lifecycle rows remain byte-for-byte deferred.
- Sequential split proposal: **WU5C2A** (depends on accepted C2A; allowed surfaces `internal/scan/descriptor.go`, `internal/scan/descriptor_test.go`, `internal/platform/macos/walker_darwin.go`, `internal/platform/macos/walker_darwin_test.go`) owns validated bound-directory record enumeration, lexical ordering, child `openat` parentage/flags, final identity/kind comparison, and typed changed/invalid/boundary facts with scripted seams. **WU5C2B** depends on accepted WU5C2A and uses the same surfaces to own DFS depth and simultaneous-FD budget enforcement, LIFO child-before-parent closure and precedence, plus disposable rename/replacement and vanished/symlink-child evidence. Each slice must start with its own RED and remain at or below 400 authored lines.
- Runtime and verification: N/A; this was a pre-write workload gate, so no focused, platform, full, vet, format, Linux compile, or LSP command was run. Baseline remains restored and unchanged. No Engram tools were injected.
- Workload / PR boundary: no accepted WU5C2 candidate; proposed sequential `stacked-to-main` WU5C2A then WU5C2B. Next action requires parent approval to amend the task plan with that dependency and assigned slice.

---

## Planning rescope — authorized WU5C2A and WU5C2B split

### Authorization, baseline, and non-acceptance

- This is planning-only under standing user split authorization and the already recorded explicit continuation/reset. Reset revision: `sha256:21a23fbbca9599085cd4d6896f05d519a8da04b52176beae7cb4f43aa566ea45`. No runtime token is persisted.
- The blocked unsplit WU5C2 forecast remains **580–780 authored changed lines**, above the 400-line cap. Its four unchecked rows are superseded, not accepted, by two sequential independently bounded strict-TDD units.
- Immutable baseline `/tmp/osdy-wu5c2-baseline-o82OM0` has manifest SHA-256 `6e8a3785e9d5560c75680a57b314f70dafc8bf1ccc78890ff8214145ec41207b`. Exact baseline entries are: `descriptor.go` `f0bbd0127bfd5a3c50faa96daae6202f91898bc292d8b6c7105995d44f9a3e93`; `descriptor_test.go` `5328215c44b189b6ca1194aafbc4d68a03d2ba60b2bb886ce9ebb9b3d2e00df6`; `walker_darwin.go` `beca553e58802696ba9af5d21281ef8707623d25d5ee92af302886e099ab644b`; `walker_darwin_test.go` `95e56aaef496274eb1f45e9d23ad3260a07b44f1432d9093d26b22644d4338f1`; baseline `tasks.md` `7f6dbc5594c86cc7646d57457d4720c5dfc5a9e5d4af383751cd8c2c665ab91d`; baseline `apply-progress.md` `ccbbaae5d0729a4257175ae7a640614d95736bf0bee700fa4a36f7d76ac75041`.
- The four product/test files remain byte-identical to that baseline because this rescope edited only `tasks.md` and this append-only receipt. No RED, product/test change, scanner/finalizer work, command execution, commit, review, or acceptance occurred.

### Authorized sequential units

- **WU5C2A core enumeration**, forecast **280–380** authored changed lines, depends on independently accepted C2A. It owns the scan-level opaque enumeration fact extension without raw FD/capability escape; Darwin already-open-FD `getdirentries`/dirent seam; safe record length/name validation including malformed, dot, dotdot, slash, and NUL cases; lexical ordering; exact child `openat` parentage/flags; enumeration kind/inode versus final-descriptor identity/kind; device/mount/local/type/identity gates before descent; and typed changed/boundary/invalid facts. Its focused evidence is `DescriptorEnumerationFacts`, `DescriptorEnumerationOwnership`, `WalkerDescriptorEnumeration`, `WalkerDirectoryRecords`, `WalkerChildOpenParentage`, `WalkerChildChanges`, and `WalkerBoundaryGates`. Real rename/replacement and depth/FD lifecycle fixtures do not belong to A.
- **WU5C2B lifecycle/adversarial**, forecast **240–360** authored changed lines, depends on independently accepted A. It owns exact descendant depth below the acquired root; simultaneous active-FD budget; child-before-parent LIFO closure; primary-plus-secondary precedence; safe siblings after isolated errors; real opened-directory rename/replacement retained enumeration; and vanished/symlink child no-target descent. Its focused evidence is `WalkerDirectoryReplacement`, `WalkerDirectoryLifecycle`, and `WalkerDepthFDBudget`, with disposable triangulation only.
- Both units use strict RED → GREEN → TRIANGULATE → REFACTOR, one writer, no generic traversal framework, no `filepath.WalkDir`, no regular-file content read, and no goroutines. WU5D1 now depends on independently accepted WU5C2B; accepted A alone cannot authorize it.

### Reconciled forecasts, mappings, and ledger

- Forecast totals are **6,030–8,500 additions**, **180–680 deletions**, and **6,210–9,180 authored changed lines**, excluding generated `go.sum`; risk High, chained PRs Yes, chain `stacked-to-main`, decision before apply No.
- Normative mappings now assign opaque descriptor-bound lexical enumeration, exact child parentage, final identity/type comparison, and pre-descent boundary gates to A; exact descendant depth/FD budget, LIFO closure/error precedence, retained opened-directory identity, and replacement no-target descent to B.
- Exact current ledger, counted from the persisted checkbox rows: **130 total = 81 checked + 49 unchecked**; **127 implementation-owned = 81 checked + 46 unchecked**; **3 parent-owned = 0 checked + 3 unchecked**. The replacement is exactly four unchecked WU5C2 rows removed/superseded and eight unchecked WU5C2A/B rows inserted; all other rows and states, including all parent rows, are preserved.

### Settlement needs and exact next action

1. Apply only WU5C2A RED, GREEN, TRIANGULATE, and REFACTOR from the immutable product baseline; record exact additions/deletions, predecessor/final hashes, focused/full/vet/format results, runtime N/A, and rollback.
2. Obtain independent semantic acceptance of A before any B edit; if A exceeds 400 authored lines or leaks B fixture/lifecycle scope, stop and resplit without acceptance.
3. Apply B only from independently accepted A, with its own strict-TDD evidence, ≤400 accounting, exact hashes, disposable fixture proof, and rollback; then obtain independent B acceptance.
4. Authorize WU5D1 only after independently accepted B. No settlement, acceptance, commit, PR, or review is recorded by this planning receipt.

**Exact next implementation row:** `WU5C2A RED` in `openspec/changes/read-only-scan-foundation/tasks.md`.

---

## WU5C2 split planning ledger correction — superseding current count

Verifier `subtask_gentle-ai-verify_1788120832955_3dd6f571` established the current ledger as **130 total = 83 checked + 47 unchecked; 127 implementation-owned = 83 checked + 44 unchecked; 3 parent-owned = 0 checked + 3 unchecked**. The method counted every Markdown checkbox row in current `tasks.md`, partitioned by terminal `sdd-owner` marker, then by `[x]` versus `[ ]` state. This supersedes only the stale current WU5C2 split count; all earlier counts remain preserved as historical snapshots. No checkbox row, row state, ownership marker, task semantic, product file, or test file changed.

---

## WU5C2A receipt — Core Descriptor-Owned Enumeration

**Predecessor / baseline proof:** The six immutable entries in `/tmp/osdy-wu5c2a-baseline-8wv3At/SHA256SUMS` verified with `shasum -a 256 -c`; its manifest SHA-256 is `766dd7f8b270ada46bd5839f3ebb0a619624faf8c92d55995e57b2d4cd128dba`. The four allowed product/test files matched that baseline before RED. No runtime token is persisted.

### Completed tasks and API

- Persisted task updates: WU5C2A RED, GREEN, TRIANGULATE, and REFACTOR are visibly `[x]`; WU5C2B remains `[ ]`.
- `scan.EnumerationFact` has private fields, validating construction and copying an optional identity; callers receive immutable scalar facts only: name, kind, identity availability, descent eligibility, and stable typed skip class.
- Darwin production uses `unix.Getdirentries` on the already-open descriptor. The narrow raw-record parser rejects short, zero, overflowing, and truncated records and invalid names, ignores only exact `.` and `..`, then lexically sorts facts.
- Each directory child uses the exact parent FD with `O_RDONLY|O_NOFOLLOW|O_DIRECTORY|O_CLOEXEC|O_NONBLOCK`; only child `fstat` and `fstatfs` form facts. Kind/inode mismatch, zero/unavailable record identity, device boundary, non-local filesystem, and operation failures become ineligible typed facts before descent.

### TDD Cycle Evidence

| Task | Test layer | Safety net | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- | --- | --- |
| WU5C2A | Go unit tests in `descriptor_test.go` and `walker_darwin_test.go` | Both exact focused commands passed with no matching tests before edits. | Both focused commands exited 1 solely for absent enumeration APIs, record helpers, and walker behavior. | Both focused commands passed after the minimal immutable fact and descriptor-bound parser/gate implementation. | Added record-boundary, exact-dot, NUL/slash, lexical sibling, kind/inode mismatch, device, and non-local table cases; focused commands passed. | `gofmt`, focused, combined packages, full suite, vet, Linux compile, and format check passed. |

### Verification and accounting

- Focused scan and Darwin commands: PASS.
- `go test ./internal/scan ./internal/platform/macos -count=1`, `go test ./... -count=1`, and `go vet ./...`: PASS.
- `GOOS=linux go test -c -o /tmp/osdycleaner-macos-linux.test ./internal/platform/macos` passed; the temporary binary was removed. `gopls` was unavailable. Runtime harness: N/A — deterministic record/syscall seams only.
- Baseline-relative accounting: `descriptor.go` +35/-0, `descriptor_test.go` +29/-0, `walker_darwin.go` +117/-1, `walker_darwin_test.go` +93/-1; **274 additions + 2 deletions = 276 authored changed lines**, below 400.
- Final hashes: `descriptor.go` `eaef5480cbd0381649f1bfdfbc78708b868c1e3b92bff16a73e2a898b3b4c5ea`; `descriptor_test.go` `2cf87a50122416037ed60684b2d35d6bd9d82f2c7674bffd6ebef2e8cbe1a497`; `walker_darwin.go` `d80e9955c7e261ba138d9e9a550ce212a9b14fbc2b81d9b45811f2e50917dba7`; `walker_darwin_test.go` `25ef09a08e0c3c3d5b82642f7c23218b49379d977e8ff8a2438ad37153133373`.
- No B scope: no real replacement fixture, depth/FD lifecycle, scanner/finalizer, worker/goroutine, content read, `WalkDir`, commit, or review action was added. Rollback restores the four allowed files from `/tmp/osdy-wu5c2a-baseline-8wv3At/`; the accepted acquisition boundary remains.
- Status consumed: authoritative OpenSpec `applyState: ready`, `repo-local` workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, allowed root the workspace, strict TDD enabled, no blocked reasons. The high-risk workload delivery path was the user-authorized `stacked-to-main` WU5C2A slice. No action-context warning applied.
- Engram mirror was not performed because no Engram memory tool was injected for this OpenSpec-only run.

### Remaining implementation work

- [ ] **WU5C2B RED:** In `internal/platform/macos/walker_darwin_test.go` (and only minimal contract assertions in `internal/scan/descriptor_test.go` if required), add focused `TestWalkerDirectoryReplacement`, `TestWalkerDirectoryLifecycle`, and `TestWalkerDepthFDBudget` tests first for exact depth zero/maximum/one-beyond below the acquired root, simultaneous FD budget exact-boundary/one-too-tight behavior, child-before-parent LIFO close order/counts, operation-primary plus typed close-secondary precedence, close-primary after success, safe siblings after isolated failures, real opened-directory rename/replacement retention, and vanished/symlink child no-target descent; run the focused command and capture RED caused only by missing B behavior. <!-- sdd-owner: implementation -->

---

## Planning correction — WU5C2A1 directory-record terminator remediation

### Verifier provenance and non-acceptance

- This was planning-only. No Go product/test implementation, proposal, specification, design, general documentation, commit, review, PR, release, or lifecycle action occurred. The supplied authorization token remains private and is not persisted.
- Verifier `subtask_gentle-ai-verify_1788121542056_f459f57c` returned **FAIL** for checked WU5C2A. Darwin parsing and the valid-record helper accept an unterminated record when `reclen == nameOffset+namlen`; therefore A's four checked rows remain historical implementation evidence only, record no acceptance, and satisfy no dependency.
- Historical A accounting remains exact at **+274/-2 = 276 authored changed lines**. Its final hashes remain `descriptor.go` `eaef5480cbd0381649f1bfdfbc78708b868c1e3b92bff16a73e2a898b3b4c5ea`, `descriptor_test.go` `2cf87a50122416037ed60684b2d35d6bd9d82f2c7674bffd6ebef2e8cbe1a497`, `walker_darwin.go` `d80e9955c7e261ba138d9e9a550ce212a9b14fbc2b81d9b45811f2e50917dba7`, and `walker_darwin_test.go` `25ef09a08e0c3c3d5b82642f7c23218b49379d977e8ff8a2438ad37153133373`. Historical baseline manifest `766dd7f8b270ada46bd5839f3ebb0a619624faf8c92d55995e57b2d4cd128dba` is retained without acceptance.

### Authorized A1 correction

- Inserted exactly four unchecked implementation-owned WU5C2A1 rows before unchanged WU5C2B in strict RED → GREEN → TRIANGULATE → REFACTOR order. A1 is limited to `internal/platform/macos/walker_darwin.go` and `internal/platform/macos/walker_darwin_test.go`, forecast **20–70 additions, 0–20 deletions**, with no B lifecycle/fixture, scanner, finalizer, contract, design, docs, generated, commit, or review surface.
- Exact focused evidence is `TestWalkerDirectoryRecordTermination` plus existing `TestWalkerDirectoryRecords`. Acceptance requires overflow-safe `reclen >= nameOffset+namlen+1`, exact NUL at `nameOffset+namlen`, valid helper emission of that terminator, and rejection of missing/non-NUL/truncated/overflowing terminators before any child open while preserving embedded-NUL and slash rejection semantics.
- WU5C2B now depends on independently accepted A1; failed A alone cannot authorize B. Rollback restores A's two walker hashes above and removes only A1 terminator validation/helper/tests.

### Forecast, mappings, and current ledger

- Replacing A's prior forecast with actual +274/-2 and adding A1 yields **6,064–8,494 additions**, **162–672 deletions**, and **6,226–9,166 authored changed lines**, excluding generated `go.sum`. Risk remains High; chained PRs remain Yes; chain remains resolved `stacked-to-main`; decision before apply remains No.
- Coverage now assigns parser termination, no-child-open malformed rejection, and preservation of valid lexical sibling semantics to A1 before B's lifecycle/adversarial evidence.
- Recounted current ledger: **134 total = 87 checked + 47 unchecked; 131 implementation-owned = 87 checked + 44 unchecked; 3 parent-owned = 0 checked + 3 unchecked**. Only four unchecked A1 rows were added; all historical A and parent states are preserved.
- Exact next implementation row is `WU5C2A1 RED`. No A1 acceptance, B authorization, product/test edit, command result, commit, or review is recorded by this planning receipt.

---

## WU5C2A1 receipt — Darwin directory-record terminator correction

### Status, scope, and task reconciliation

- Consumed authoritative OpenSpec status reconstructed from the active artifacts: `change=read-only-scan-foundation`, `artifactStore=openspec`, `applyState=ready`, workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, strict TDD enabled by `openspec/config.yaml`, no blocked reasons, and the parent-authorized `stacked-to-main` WU5C2A1 slice resolving the high-workload guard. No `actionContext` warning or restricted edit root was supplied; edits remained in the two assigned Darwin walker files plus required OpenSpec artifacts.
- Completed and visibly checked in `tasks.md`: WU5C2A1 RED, GREEN, TRIANGULATE, and REFACTOR. WU5C2B and every later implementation task remain unchecked. Parent-owned lifecycle rows were not edited.
- Historical WU5C2A remains failed/not accepted evidence: verifier `subtask_gentle-ai-verify_1788121542056_f459f57c` found the missing terminator defect. Its immutable failed-A walker hashes were `walker_darwin.go` `d80e9955c7e261ba138d9e9a550ce212a9b14fbc2b81d9b45811f2e50917dba7` and `walker_darwin_test.go` `25ef09a08e0c3c3d5b82642f7c23218b49379d977e8ff8a2438ad37153133373`.

### TDD Cycle Evidence

| Task | Test layer | Safety net | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- | --- | --- |
| WU5C2A1 | Darwin package unit tests in `walker_darwin_test.go` | `go test ./internal/platform/macos -run '^TestWalkerDirectoryRecords$' -count=1` PASS before edits. | Added exact `TestWalkerDirectoryRecordTermination`; focused command exited 1. The exact failures were `missing_terminator` and `non-NUL_terminator`, each reporting accepted child facts and `err=<nil>`. | Parser now requires `nameLength <= recordLength-direntHeader-1` and `raw[direntHeader+nameLength] == 0`; `dirRecord` emits the required terminator. Focused command PASS. | Covered zero `namlen` through existing `TestWalkerDirectoryRecords`, maximum representable name, minimum valid record, nonzero garbage padding after a terminator, embedded NUL/slash rejection, and a valid-plus-malformed batch with zero child opens. Focused command PASS. | Formatted both files and ran focused, A-focused scan/platform, combined, Linux compile/remove, full suite, vet, and format checks; all PASS. |

### Parser boundary and verification

- Required parser boundary: `nameOffset+namlen+1 <= reclen <= remaining bytes`, with an exact NUL at `record[nameOffset+namlen]`. Padding after that terminator is intentionally allowed, including non-NUL padding; the terminator is not required to occupy `reclen-1`.
- Malformed records return typed `scan.ErrInvalidMetadata` before `enumerateDirectory` reaches any child `open`; the malformed mixed batch similarly makes zero child opens. Exact `.` and `..` ignore behavior and embedded-NUL/slash rejection remain unchanged.
- `go test ./internal/platform/macos -run '^(TestWalkerDirectoryRecordTermination|TestWalkerDirectoryRecords)$' -count=1` -> PASS.
- `go test ./internal/scan -run '^(TestDescriptorEnumerationFacts|TestDescriptorEnumerationOwnership)$' -count=1` -> PASS.
- `go test ./internal/platform/macos -run '^(TestWalkerDescriptorEnumeration|TestWalkerDirectoryRecords|TestWalkerChildOpenParentage|TestWalkerChildChanges|TestWalkerBoundaryGates)$' -count=1` -> PASS.
- `go test ./internal/scan ./internal/platform/macos -count=1` -> PASS.
- `GOOS=linux go test -c -o /tmp/osdycleaner-walker-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-walker-linux.test` -> PASS; removal was confirmed.
- `go test ./... -count=1` -> PASS. `go vet ./...` -> PASS. `gofmt -d` over both modified files produced no output.
- LSP: `gopls` was unavailable, so no LSP diagnostic command could run; compiler, focused/package/full tests, and vet were clean.
- Runtime harness: N/A — this correction uses deterministic raw-directory-record and poison child-open seams only; it does not invoke an executable scanner, real home, mount, privilege, or external volume.

### Accounting, identity, rollback, and deferred work

- Baseline `/tmp/osdy-wu5c2a1-baseline-VDEmju` verified with `sha256sum -c SHA256SUMS`; its manifest SHA-256 is `7f21fc84f1fc99ffd5bc57fc8708acadf1669046524f107b89319eb218434a79`.
- Correction accounting from that baseline: `walker_darwin.go` `+6/-2`; `walker_darwin_test.go` `+60/-1`; total **66 additions + 3 deletions = 69 authored changed lines**, within the 400-line cap.
- Final correction hashes: `internal/platform/macos/walker_darwin.go` `412a412f0e2554d05f9198bdf6381a4fd75a08f7c01eb21f161ceefbc99de020`; `internal/platform/macos/walker_darwin_test.go` `7809445656752727c40f8d4be6bbbdd11c8f78d3a91437d88930a7bcf9cf7d00`.
- Rollback boundary: restore exactly those two walker files from `/tmp/osdy-wu5c2a1-baseline-VDEmju`; this removes only the terminator validation, valid-record helper correction, and A1 tests while preserving failed-A historical candidate state.
- No WU5C2B/later, scanner, finalizer, commit, review, receipt approval, or lifecycle gate was started. Workload / PR boundary is one 69-line WU5C2A1 stacked-to-main correction; no commit or PR was created.
- Engram mirror could not be performed because the injected Engram provider was unreachable; OpenSpec persistence above was performed.

### Exact remaining implementation work

- [ ] **WU5C2B RED:** In `internal/platform/macos/walker_darwin_test.go` (and only minimal contract assertions in `internal/scan/descriptor_test.go` if required), add focused `TestWalkerDirectoryReplacement`, `TestWalkerDirectoryLifecycle`, and `TestWalkerDepthFDBudget` tests first for exact depth zero/maximum/one-beyond below the acquired root, simultaneous FD budget exact-boundary/one-too-tight behavior, child-before-parent LIFO close order/counts, operation-primary plus typed close-secondary precedence, safe sibling continuation, and real opened-directory rename/replacement/no-target descent; run the focused command and capture a genuine RED. <!-- sdd-owner: implementation -->

---

## WU5C2A1 accounting adjudication addendum

- This receipt-only addendum supersedes the incorrect `+66/-3=69` claim; no product, test, task, design, documentation, review, or lifecycle artifact was changed.
- Independent reproduction from `/tmp/osdy-wu5c2a1-baseline-VDEmju` measures production `+6/-2` and test `+64/-1`, totaling **+70/-3=73** authored changed lines.
- Exact current hashes: `walker_darwin.go` `412a412f0e2554d05f9198bdf6381a4fd75a08f7c01eb21f161ceefbc99de020`; `walker_darwin_test.go` `7809445656752727c40f8d4be6bbbdd11c8f78d3a91437d88930a7bcf9cf7d00`.
- Semantic verifier PASS except integrity: `subtask_gentle-ai-verify_1788122606468_ea94f7a9`; failed evidence `sha256:b037c7390b56d0d92df223774145af745f614aff31eab6abd8bec8da873f26c9`.
- User-authorized reset revision: `sha256:bb7aa1149c06ca29e7b8ac7f20fb0a97cbf0291ef3c12c0f62dbec3ef20cce90`.
- Task checkboxes are preserved; WU5C2B remains blocked until independent adjudication PASS, and this addendum does not settle any attempt or approve a receipt.

---

## WU5C2B apply progress — blocked by mandatory pre-write split forecast

### Status and baseline consumed

- Consumed authoritative native OpenSpec status for `read-only-scan-foundation`: `artifactStore=openspec`, `applyState=ready`, `nextRecommended=apply`, strict TDD enabled by `openspec/config.yaml`, and `actionContext.mode=repo-local` with allowed workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`. No action-context warning applied.
- Continued only the supplied active-attempt context; no acquire, settle, review, receipt, commit, or lifecycle action was performed.
- Verified `/tmp/osdy-wu5c2b-baseline-kLCdoo/SHA256SUMS` before any product/test write: all six entries passed `shasum -a 256 -c SHA256SUMS`; the `SHA256SUMS` file digest is the supplied `e178b1cc4f02cf5d7c5727cb29500115800c0c98d4e2dca546e3f585e44c4ddf`.
- Current allowed product/test files are byte-identical to that baseline: `descriptor.go` `eaef5480cbd0381649f1bfdfbc78708b868c1e3b92bff16a73e2a898b3b4c5ea`, `descriptor_test.go` `2cf87a50122416037ed60684b2d35d6bd9d82f2c7674bffd6ebef2e8cbe1a497`, `walker_darwin.go` `412a412f0e2554d05f9198bdf6381a4fd75a08f7c01eb21f161ceefbc99de020`, and `walker_darwin_test.go` `7809445656752727c40f8d4be6bbbdd11c8f78d3a91437d88930a7bcf9cf7d00`.

### Mandatory honest pre-write forecast

| Required B behavior | Plausible production lines | Plausible named-test/fixture lines | Total plausible authored lines |
| --- | ---: | ---: | ---: |
| Serial descriptor-bound DFS frame/child state, exact descendant depth, and pre-open active-FD reservation/release | 105–145 | 55–80 | 160–225 |
| LIFO child-before-parent close ledger, operation-primary/typed-close-secondary joining, close-primary-after-success, and isolated-sibling continuation | 75–110 | 85–125 | 160–235 |
| Disposable real opened-directory rename/replacement retention proof, plus vanished and symlink-replaced child no-target descent proof | 25–45 | 135–190 | 160–235 |
| Cross-cutting focused helpers, trace/count assertions, and adversarial triangulation | 15–30 | 55–85 | 70–115 |
| **Total** | **220–330** | **330–480** | **550–810** |

- The low-end total is 550 authored lines, already 150 lines above the independent 400-line cap. The required real `t.TempDir()` rename/replacement proof cannot honestly be omitted or compressed into scripted-only evidence, and the separate close/error matrix must remain readable rather than merging unrelated assertions.
- Therefore no RED test, production code, Go/test-file edit, task checkbox update, scanner/finalizer work, review, commit, or verification command was run. `WU5C2B RED`, `GREEN`, `TRIANGULATE`, and `REFACTOR` remain visibly unchecked in `tasks.md`.

### Required task replan under standing split authorization

Split WU5C2B before implementation into two serial strict-TDD units, each with independent baseline, accounting, verification, and acceptance:

1. **WU5C2B1 — DFS descriptor lifecycle and bounded sibling continuation**: `walker_darwin.go` and `walker_darwin_test.go`; exact descendant depth, pre-open simultaneous-FD cap, LIFO closure/counts, primary-plus-typed-secondary errors, close primary after successful work, and isolated safe sibling continuation. Forecast 250–360 authored lines. Excludes all real pathname-mutation fixtures.
2. **WU5C2B2 — Real descriptor-binding mutation fixtures**: `walker_darwin_test.go` plus only any minimal pre-approved fixture seam; `t.TempDir()` opened-directory rename/replacement retention identity, vanished child, symlink-replaced child, and no replacement-target descent. Forecast 180–280 authored lines. Depends on independently accepted B1 and adds no DFS semantics.

- This proposed split preserves the user-authorized stacked-to-main delivery mode and does not itself accept, settle, or approve either slice. The parent must amend/replan task rows and assign the next bounded slice before code changes.
- Engram mirroring was attempted through the injected provider but it was unreachable at `http://127.0.0.1:7437`; no Engram artifact was claimed persisted.
- Deferred lifecycle actions: all parent-owned task rows remain byte-for-byte unchanged.

---

## Planning rescope — untouched oversized WU5C2B split into B1 and B2

### Scope, reset, and immutable provenance

- Planning only: no Go/test, proposal, specification, design, documentation, commit, review, verification command, lifecycle action, or checkbox acceptance occurred. The private continuation credential was consumed only as delegated context and is deliberately not persisted.
- Re-read the authoritative specification, design, tasks, project config, A1 accounting adjudication, and blocked B receipt before amending the plan.
- Accepted WU5C2A1 is carried forward at corrected **+70/-3=73** authored lines. Final hashes remain `internal/platform/macos/walker_darwin.go` `412a412f0e2554d05f9198bdf6381a4fd75a08f7c01eb21f161ceefbc99de020` and `internal/platform/macos/walker_darwin_test.go` `7809445656752727c40f8d4be6bbbdd11c8f78d3a91437d88930a7bcf9cf7d00`. Semantic verifier `subtask_gentle-ai-verify_1788122606468_ea94f7a9` is PASS under user-authorized reset revision `sha256:bb7aa1149c06ca29e7b8ac7f20fb0a97cbf0291ef3c12c0f62dbec3ef20cce90`; failed pre-reset integrity evidence remains `sha256:b037c7390b56d0d92df223774145af745f614aff31eab6abd8bec8da873f26c9` as historical provenance only.
- Verified blocked-B provenance is retained from `/tmp/osdy-wu5c2b-baseline-kLCdoo`; its `SHA256SUMS` digest is `e178b1cc4f02cf5d7c5727cb29500115800c0c98d4e2dca546e3f585e44c4ddf`. The blocked receipt established an honest combined-B forecast of **550–810 authored lines** and that all Go/test files were unchanged.
- Byte-identical blocked baseline hashes remain: `descriptor.go` `eaef5480cbd0381649f1bfdfbc78708b868c1e3b92bff16a73e2a898b3b4c5ea`; `descriptor_test.go` `2cf87a50122416037ed60684b2d35d6bd9d82f2c7674bffd6ebef2e8cbe1a497`; `walker_darwin.go` `412a412f0e2554d05f9198bdf6381a4fd75a08f7c01eb21f161ceefbc99de020`; `walker_darwin_test.go` `7809445656752727c40f8d4be6bbbdd11c8f78d3a91437d88930a7bcf9cf7d00`.

### Task replacement and independent boundaries

- The original four unchecked WU5C2B rows remain explicitly unchecked historical **DO NOT EXECUTE** rows. They document the untouched oversized/non-accepting combined plan and satisfy no dependency.
- Added four unchecked WU5C2B1 strict-TDD rows. B1 owns deterministic core DFS depth exactly below root, descriptor capacity before child open, exact LIFO ownership, operation-primary plus typed-secondary close precedence, close-primary after success, and safe sibling continuation. Forecast is **260–380 authored lines**. Its focused rollback restores the exact accepted-A1/baseline Go/test hashes and removes no A1 behavior. Real rename/replacement fixtures are excluded.
- Added four unchecked WU5C2B2 verification-first rows. B2 owns disposable real `t.TempDir()` opened-directory rename/replacement identity binding, vanished and symlink-replaced child changed/no-target descent, safe siblings, and only-needed combined depth/FD or error cases. Forecast is **180–300 authored lines**, preferably test-only. It must not fabricate RED: unchanged B1 production PASS is valid test-only verification; only a correct failing fixture may become genuine RED before a minimal strict-TDD production correction. Its rollback restores exact independently accepted B1 hashes.
- Dependency chain is now accepted A1 → independently accepted B1 → independently accepted B2 → WU5D1. WU5D1 and downstream rollback/mapping references were updated to B2, not historical combined B.

### Forecast, ledger, and preservation

- Recounted active implementation forecast replaces A1's stale 20–70 additions/0–20 deletions with actual +70/-3 and replaces stale combined-B 220–330 additions/20–30 deletions with B1 240–360/+20 plus B2 180–300/+0. New total is **6,314–8,824 additions, 165–645 deletions, 6,479–9,469 authored changed lines**, excluding generated `go.sum`. The historical 550–810 combined-B estimate is recorded but not double-counted.
- Risk remains High; chained PRs remain recommended; delivery remains `ask-on-risk` resolved by standing authorization; chain strategy remains `stacked-to-main`; decision before apply remains No.
- Exact persisted ledger is **142 checkbox rows = 139 implementation-owned + 3 parent-owned**. Implementation rows are **87 checked + 52 unchecked**; parent rows are **0 checked + 3 unchecked**; overall is **87 checked + 55 unchecked**.
- Parent-owned rows and every unrelated implementation checkbox retain their state, order, and terminal ownership marker. No parent action was launched.
- Exact next executable implementation row is WU5C2B1 RED. Historical B rows are never executable; B2 and WU5D1 remain blocked on sequential independent acceptance.

---

## WU5C2B split-planning provenance and ledger correction addendum

- Independent verifier `subtask_gentle-ai-verify_1788123837977_09dea9f2` failed the planning artifacts on ledger and reset-revision provenance only; this addendum corrects those planning facts without settling the attempt.
- This addendum supersedes the immediately preceding incorrect ledger values of **87 checked / 55 unchecked overall** and **87 checked / 52 unchecked implementation-owned**. The recounted actual ledger is **142 total = 91 checked + 51 unchecked; 139 implementation-owned = 91 checked + 48 unchecked; 3 parent-owned = 0 checked + 3 unchecked**.
- No checkbox state changed. The active implementation forecast remains **6,314–8,824 additions, 165–645 deletions, 6,479–9,469 authored changed lines** because its formula does not depend on checkbox states.
- The actual oversized WU5C2B reset and B1/B2 split revision is `sha256:21a23fbbca9599085cd4d6896f05d519a8da04b52176beae7cb4f43aa566ea45`; it supersedes any attribution of the B split to another revision.
- Revision `sha256:bb7aa1149c06ca29e7b8ac7f20fb0a97cbf0291ef3c12c0f62dbec3ef20cce90` is retained solely as the WU5C2A1 accounting reset/adjudication revision and does not authorize or identify the B split.
- All otherwise verified B/B1/B2 scopes, forecasts, dependencies, baseline manifests, file hashes, rollback boundaries, and next-action gating remain unchanged. No Go/test/design/docs/commit/review/lifecycle action occurred, and no settlement was performed.

---

## WU5C2B1 receipt — core DFS, descriptor budget, and lifecycle

### Status, scope, and pre-write integrity

- Consumed native authoritative OpenSpec status: `change=read-only-scan-foundation`, `artifactStore=openspec`, `applyState=ready`, `nextRecommended=apply`, strict TDD enabled by `openspec/config.yaml`, `repo-local` workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, and allowed edit root equal to that workspace. No status blockers or action-context warnings were present.
- The parent supplied the standing `stacked-to-main` delivery decision and the active WU5C2B1 continuation token. The continuation acquire returned `proceed`; no new attempt was created. This receipt does not perform review, validation, commit, PR, or settlement.
- Before any Go/test write, `/tmp/osdy-wu5c2b1-baseline-tJIx3p/SHA256SUMS` verified with `shasum -a 256 -c SHA256SUMS`; all six entries were `OK`, and its SHA-256 was `26b18511ca40bc9949af4fb7ab5705ef9200b498b02181d2fb95e688ce3bb86d`. The workspace copies of all six manifest files were byte-identical to that baseline.

### Detailed forecast and boundary

| Surface | Base forecast | Contingency | Forecast with buffer |
| --- | ---: | ---: | ---: |
| Minimal typed traversal-limit contract | 1–6 | 0–4 | 1–10 |
| Serial DFS frames, child gate, ownership/close aggregation | 65–105 | 20–30 | 85–135 |
| Deterministic test seam and named depth/FD/lifecycle/sibling tests | 105–145 | 35–55 | 140–200 |
| **Total authored additions plus deletions** | **171–256** | **55–89** | **226–345** |

The buffered high estimate was 345, below the 400-line hard cap; therefore the authorized B1 slice proceeded. B2 real rename/replacement fixtures, scanner/finalizer, workers/goroutines, content reads, `WalkDir`, docs/design, commit, review, and RDD were excluded.

### TDD Cycle Evidence

| Stage | Evidence | Result |
| --- | --- | --- |
| RED | Added the four exact named B1 tests before production behavior. | The exact focused command exited 1 only because `trustedRootWalker.walkDirectory`, `scan.ErrLimit`, and the deterministic directory-map seam were absent. |
| GREEN | Added the minimum serial `walkFrame` DFS and `ErrLimit` contract. | The exact focused command passed. |
| TRIANGULATE | Added depth exact/one-beyond, root-counted FD exact/one-tight, multi-close, close-after-success, changed child, and later sibling traces. | The exact focused command passed. |
| REFACTOR | Formatted only B1-touched walker/descriptor files and ran required checks. | Focused, package, full, vet, Linux compile/remove, and format checks passed; `gopls` was unavailable. |

### Parser and lifecycle semantics

- The accepted A1 directory-record parser remains unchanged: it rejects malformed/non-terminated records before child opening, retains lexical accepted-name order, and is not relaxed by B1.
- DFS starts at acquired root depth zero; a child at `depth+1 == MaxDepth` is accepted, while a deeper child produces typed `ErrLimit` evidence and is not opened.
- The acquired root consumes one active capability. Before every child `openat`, B1 rejects `active >= MaxDescriptors`; no child open occurs when capacity is absent.
- Each successful child open has one caller-owned close after recursive work. Recursive returns close deepest child first, yielding exact child-before-parent LIFO traces and no double close.
- The first operation failure remains the joined primary error; typed `close` `PathError` values remain discoverable as secondary evidence. A close error is primary only after otherwise successful child work. Open, changed/gate, and limit failures continue lexical safe siblings.

### Verification, accounting, hashes, and rollback

- RED: `go test ./internal/platform/macos -run '^(TestWalkerDFSDepth|TestWalkerDescriptorCapacity|TestWalkerDirectoryLifecycle|TestWalkerSiblingContinuation)$' -count=1` -> exit 1 with only missing B1 symbols/seam listed above.
- GREEN/TRIANGULATE and final focused: same command -> PASS.
- Required refactor checks -> PASS: `gofmt -w internal/scan/descriptor.go internal/platform/macos/walker_darwin.go internal/platform/macos/walker_darwin_test.go`; `go test ./internal/scan ./internal/platform/macos -count=1`; `go test ./... -count=1`; `go vet ./...`; `GOOS=linux go test -c -o /tmp/osdycleaner-wu5c2b1-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-wu5c2b1-linux.test`; and clean `gofmt -d` output. `gopls` was not installed.
- Runtime: N/A — deterministic descriptor-operation seams are this core platform unit's bounded harness; no scanner or executable boundary exists.
- Accounting against the verified B1 baseline: `descriptor.go` +1/-0; `descriptor_test.go` +0/-0; `walker_darwin.go` +79/-0; `walker_darwin_test.go` +112/-1; **192 additions + 1 deletion = 193 authored changed lines**, below the 400-line cap.
- Final hashes: `internal/scan/descriptor.go` `6f7741457a761c95e0532e271394883b3a6b6eab64343d8479ca8e4b09f14a5b`; `internal/scan/descriptor_test.go` `2cf87a50122416037ed60684b2d35d6bd9d82f2c7674bffd6ebef2e8cbe1a497`; `internal/platform/macos/walker_darwin.go` `08f24772d963d8b44db42f677f750da100e8f04c12b8926fc5502b8d33f84610`; `internal/platform/macos/walker_darwin_test.go` `e87eaa3a56896decf2d205c0c2163a4e50075b460c9ab6f407715d4cada030ee`.
- Rollback boundary: restore these four Go/test files from `/tmp/osdy-wu5c2b1-baseline-tJIx3p/`, removing only B1 DFS/limit/lifecycle behavior and tests while preserving accepted A1 parser semantics. No commit or PR was created.
- Task reconciliation: all four implementation-owned WU5C2B1 rows are visibly `[x]`; B2 remains unchecked. Parent-owned lifecycle rows were deferred and retained byte-for-byte. Engram mirroring was attempted through the injected provider but it was unreachable, so only OpenSpec persistence was performed.

### Exact remaining unchecked rows

- [ ] **WU5C2B HISTORICAL RED — DO NOT EXECUTE:** The combined RED would have mixed core DFS/FD/lifecycle seams with real rename/replacement fixtures and was blocked before any test write because the honest 550–810-line forecast exceeds 400; it is superseded by B1 then B2 and cannot be accepted. <!-- sdd-owner: implementation -->
- [ ] **WU5C2B HISTORICAL GREEN — DO NOT EXECUTE:** No combined production implementation exists; all Go/test files remained byte-identical to manifest `e178b1cc4f02cf5d7c5727cb29500115800c0c98d4e2dca546e3f585e44c4ddf`, and only independently accepted B1 may establish core DFS/lifecycle behavior. <!-- sdd-owner: implementation -->
- [ ] **WU5C2B HISTORICAL TRIANGULATE — DO NOT EXECUTE:** No combined real fixture or adversarial test was authored; independent B2 owns those fixtures only after accepted B1 and must not fabricate RED when B1 production already satisfies them. <!-- sdd-owner: implementation -->
- [ ] **WU5C2B HISTORICAL REFACTOR — DO NOT EXECUTE:** No combined verification, acceptance, rollback, commit, or review is permitted; retain this row as blocked provenance and use the separately bounded B1/B2 rollback boundaries below. <!-- sdd-owner: implementation -->
- [ ] **WU5C2B2 VERIFY/RED:** Add the smallest real `t.TempDir()` opened-directory rename/replacement identity fixture and vanished/symlink-replaced child no-target-descent expectations in `walker_darwin_test.go`; run the focused command. If unchanged B1 production passes, record verification-first PASS with no RED claim or production edit; if a correct expectation fails, preserve that failure as genuine RED before any production change. <!-- sdd-owner: implementation -->
- [ ] **WU5C2B2 GREEN/VERIFY:** When VERIFY passed, keep production byte-identical and add only test/helper evidence; when and only when the prior correct expectation produced genuine RED, make the minimum `walker_darwin.go` correction and rerun to GREEN. In both paths, explicitly stat/fstat and compare renamed-original versus replacement identities and prove child opens remain parent-descriptor-bound. <!-- sdd-owner: implementation -->
- [ ] **WU5C2B2 TRIANGULATE:** Add vanished child, symlink replacement, no replacement-target descent, safe later sibling, and only-needed combined depth/FD plus operation/close real cases; assert exact identities, opens, closes, and traces, and prefer test-only coverage whenever accepted B1 behavior already satisfies the fixture. <!-- sdd-owner: implementation -->
- [ ] **WU5C2B2 REFACTOR:** Run gofmt on B2-touched files, the exact B2 focused command, `go test ./internal/platform/macos -count=1`, `go test ./... -count=1`, `go vet ./...`, and clean `gofmt -d`; record whether production stayed byte-identical, exact hashes/accounting ≤400, fixture-local mutation only, rollback, and independent acceptance before WU5D1. <!-- sdd-owner: implementation -->
- [ ] **WU5D1 RED:** Add fake-walker tests for canonical built-in validation before platform construction/call, exact five-root serial order, normal/missing/inaccessible/symlink/non-directory/device-boundary/non-local acquisition states, cancellation before a root, poison resolver/platform seams, and five ordered honest incomplete zero-file transitional observations; run the focused command and capture RED caused only by missing scanner consumer APIs. <!-- sdd-owner: implementation -->
- [ ] **WU5D1 GREEN:** Implement the minimum serial scanner consumer with injected home/built-ins/walker, validated absolute home and relative built-in components, one active root, stable acquisition-state mapping, later-root continuation, and five ordered transitional observations; do not add file workers or claim scanned/complete zero-byte coverage; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU5D1 TRIANGULATE:** Add reordered fake definitions, duplicate/extra/unclean roots, root-level combined errors, cancellation after an earlier root, and poison callbacks; prove invalid canonical inputs make zero platform calls, safe later roots continue, no production home is read, and warnings/root states are deterministic; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU5D1 REFACTOR:** Run `gofmt -w internal/scan/scanner.go internal/scan/scanner_test.go internal/scan/descriptor.go && go test ./internal/scan -run '^(TestScannerCanonicalPreflight|TestScannerSerialRoots|TestScannerAcquisitionStates|TestScannerFiveTransitionalObservations|TestScannerNoRealHome)$' -count=1 && go test ./internal/scan -count=1 && go test ./... -count=1 && go vet ./...`; expect PASS, race N/A because roots remain serial without goroutines, and record predecessor/final hashes, <=400 accounting, runtime N/A, and rollback before WU5D2. <!-- sdd-owner: implementation -->
- [ ] **WU5D2 RED:** Add fake-walker/finalizer tests for validated relative components and containment, synchronous directory kind/metadata/identity facts, enumeration-versus-open kind/identity mismatch, disappearance, symlink/special entry, different-device/non-local boundary, malformed/typed platform errors, deterministic warnings, safe siblings/later roots, and exact `partial > boundary_limited > scanned` mapping across all five final observations; run the focused command and capture RED caused only by missing directory-consumer integration. <!-- sdd-owner: implementation -->
- [ ] **WU5D2 GREEN:** Map immutable walker directory/skip/error facts into raw scanner state and the existing finalizer, reject unclean/escaping components, classify changed/disappeared directories as zero-contribution `entry_changed`, retain lower-precedence warnings, continue safe siblings/later roots, and finalize five ordered observations without regular-file bytes; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU5D2 TRIANGULATE:** Add combined changed/symlink/boundary/inaccessible fixtures and callback-order permutations; prove deterministic warning order, correct partial/boundary/root statuses, no target facts, no unsafe descent request, safe continuation, and no complete claim while regular-file evidence is absent; rerun and expect PASS without real-home inspection. <!-- sdd-owner: implementation -->
- [ ] **WU5D2 REFACTOR:** Run `gofmt -w internal/scan/scanner.go internal/scan/scanner_test.go internal/scan/finalize.go internal/scan/finalize_test.go && go test ./internal/scan -run '^(TestScannerDirectoryFacts|TestScannerDirectoryChanges|TestScannerDirectoryBoundaries|TestScannerRootStatusMapping|TestScannerFiveFinalObservations)$' -count=1 && go test ./internal/scan -count=1 && go test ./... -count=1 && go vet ./...`; expect PASS, race N/A because no goroutines exist, and record accepted-WU5D1/final hashes, <=400 accounting, runtime N/A, and rollback before WU5E1. <!-- sdd-owner: implementation -->
- [ ] **WU5E1 RED:** Add synchronized scanner and Darwin tests for fixed worker/job/result counts, bounded capacities and backpressure, validated relative identity-chain jobs, trusted-root private-anchor reopen, rolling ancestor descriptors, exact no-follow flags, ancestor/final identity and device/local checks, enumeration-versus-open kind/identity/disappearance, regular-only logical/allocation facts, no content reads, and exact close counts; run both focused commands and capture RED caused only by the missing file pipeline/reopen APIs. <!-- sdd-owner: implementation -->
- [ ] **WU5E1 GREEN:** Implement the minimum fixed worker and bounded channel pipeline plus platform relative reopen: stop absolute descendant inspection, reopen each component from the private root anchor, close rolling descriptors, accept only unchanged same-boundary regular final-descriptor facts, and return changed/boundary/error results with zero bytes otherwise; rerun both focused commands and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU5E1 TRIANGULATE:** Add blocked enqueue/result cases, ancestor/final rename or symlink replacement, disappearance, non-regular entries, different-device/non-local facts, metadata/close failures, deep identity chains, and queue-pressure permutations; prove bounded goroutines/FDs, no target/content facts, no dropped accepted result, and deterministic zero contribution for changed entries; rerun the focused commands and `go test -race ./internal/scan -run '^(TestScannerBoundedFileJobs|TestScannerFileBackpressure|TestScannerRegularOnlyFacts|TestScannerFileEntryChanges)$' -count=1`, expecting PASS. <!-- sdd-owner: implementation -->
- [ ] **WU5E1 REFACTOR:** Run `gofmt -w internal/scan/descriptor.go internal/scan/scanner.go internal/scan/scanner_test.go internal/platform/macos/walker_darwin.go internal/platform/macos/walker_darwin_test.go && go test ./internal/scan -run '^(TestScannerBoundedFileJobs|TestScannerFileBackpressure|TestScannerRegularOnlyFacts|TestScannerFileEntryChanges)$' -count=1 && go test ./internal/platform/macos -run '^(TestWalkerRelativeFileReopen|TestWalkerFileIdentityChain|TestWalkerRegularFileFacts|TestWalkerFileReopenLifecycle)$' -count=1 && go test -race ./internal/scan -count=1 && go test ./... -count=1 && go vet ./...`; expect PASS and record exact predecessor/final hashes, <=400 accounting, runtime N/A, and rollback before WU5E2. <!-- sdd-owner: implementation -->
- [ ] **WU5E2 RED:** Add channel-synchronized tests for worker/result completion permutations, same-identity hard links, changed/disappeared file warnings, cancellation before enqueue/after acceptance/while blocked, stopped new jobs, result draining, worker joining, no leaked goroutines, and deterministic finalizer attribution/warnings; run the focused command and capture RED caused only by missing completion/cancellation integration. <!-- sdd-owner: implementation -->
- [ ] **WU5E2 GREEN:** Implement cancellation-aware ownership that stops enqueueing and worker acceptance, closes job input once, drains every accepted result, joins the fixed workers, canonically sorts facts before existing hard-link finalization, and preserves completed roots plus active/later cancellation states without scheduling-dependent attribution; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU5E2 TRIANGULATE:** Add blocked job/result channels, repeated cancellation, mixed boundary/partial/change errors, worker completion permutations, cross-area hard links, and close/error combinations; prove cancellation precedence, no discarded accepted fact, no post-cancel open/job, deterministic warnings/totals, regular-only bytes, and no goroutine/descriptor leak; rerun the focused command and `go test -race ./internal/scan -run '^(TestScannerFileCompletionPermutations|TestScannerHardLinkFinalization|TestScannerCancellationStopDrainJoin|TestScannerNoPipelineLeaks|TestScannerDeterministicWarnings)$' -count=1`, expecting PASS. <!-- sdd-owner: implementation -->
- [ ] **WU5E2 REFACTOR:** Run `gofmt -w internal/scan/scanner.go internal/scan/scanner_test.go internal/scan/finalize.go internal/scan/finalize_test.go && go test ./internal/scan -run '^(TestScannerFileCompletionPermutations|TestScannerHardLinkFinalization|TestScannerCancellationStopDrainJoin|TestScannerNoPipelineLeaks|TestScannerDeterministicWarnings)$' -count=1 && go test -race ./internal/scan -count=1 && go test ./... -count=1 && go vet ./...`; expect PASS, record accepted-WU5E1/final hashes, <=400 accounting, runtime N/A, and rollback. WU6 remains blocked until this independent receipt passes. <!-- sdd-owner: implementation -->
- [ ] **WU6 RED:** Add deterministic channel-synchronized tests for cancellation before a root, during an active metadata call, after an earlier root, and alongside partial/boundary warnings; add entry/path limit tests and run the focused command to record failures for missing stop/state behavior without sleep-based timing. <!-- sdd-owner: implementation -->
- [ ] **WU6 GREEN:** Implement context checks before new work, immediate enqueue stop, bounded active-call completion, channel drain/join, cancellation warnings, skipped later roots, limit warnings, scan-policy fields, and exact `cancelled > partial > boundary_limited > scanned` derivation; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU6 TRIANGULATE:** Add reordered worker completion, blocked-result-channel, unknown-required-metadata, combined limit/boundary/error, and repeated-cancel cases; run `go test -race ./internal/scan -run '^(TestScannerStatusPrecedence|TestScannerEntryLimit|TestScannerPathBudget|TestScannerCancellation|TestScannerNoLeakedWork|TestScannerCompletionOrderStable)$' -count=1` and expect PASS with no race, leaked work, discarded finalized fact, or newly started root. <!-- sdd-owner: implementation -->
- [ ] **WU6 REFACTOR:** Run `gofmt -w internal/scan/scanner.go internal/scan/scanner_test.go internal/scan/finalize.go internal/scan/finalize_test.go && go test -race ./internal/scan -count=1`; expect PASS and verify partial/cancelled aggregate estimates remain incomplete. <!-- sdd-owner: implementation -->
- [ ] **WU7 RED:** Add fixed-snapshot report tests for complete, partial, and cancelled outcomes, exact top-level/nested field order, typed bytes/booleans, explicit unknowns, `related_path: null`, `[]` collections, one trailing newline, all five roots, and text/JSON fact equivalence; run the focused command and record missing-renderer/golden failures. <!-- sdd-owner: implementation -->
- [ ] **WU7 GREEN:** Implement ordered DTO projection in `internal/report/json.go` and canonical text rendering in `internal/report/text.go`, then run `go test ./internal/report -run '^(TestJSONReport|TestTextReport|TestReportDeterminism|TestReportEquivalence)$' -update -count=1`; inspect all six allowed goldens and rerun the focused command without `-update`, expecting PASS. <!-- sdd-owner: implementation -->
- [ ] **WU7 TRIANGULATE:** Add discovery/collection permutations and assertions for byte-identical JSON, JSON parseability, schema/outcome/completeness, risks, bases, warnings, prominent partial/cancelled text, APFS clone/snapshot/compression/hard-link caveats, CoreSimulator manual-review wording, and no deletion-safety/reclaim promise; rerun without `-update` and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU7 REFACTOR:** Run `gofmt -w internal/report/json.go internal/report/text.go internal/report/report_test.go && go test ./internal/report -count=1`; expect PASS without golden updates, map-backed serialization, absolute home paths, wall-clock/host/user/random fields, or presentation-derived scan facts. <!-- sdd-owner: implementation -->
- [ ] **WU8 RED:** Add direct model tests for pure `Init`, canonical category movement, h/j/k/l and arrows, page scrolling, Tab details, resize, q/Escape/Ctrl-C quit, complete/partial/cancelled views, and terminal rejection; run the focused command and record missing-model/dependency failures. <!-- sdd-owner: implementation -->
- [ ] **WU8 GREEN:** Pin approved compatible v2 dependency versions in `go.mod`, implement snapshot-only model/update/view/runner files, use only the viewport component, and rerun the focused command; expect PASS with no scan command from `Init` and no command after quit. <!-- sdd-owner: implementation -->
- [ ] **WU8 TRIANGULATE:** Assert prominent “Read-only scan,” “Manual review required,” and “Estimate—not guaranteed reclaimable space” wording, cancellation/incompleteness before facts, canonical area/warning agreement, and absence of cleanup/select/action/confirm/delete/remove/Trash affordances while allowing the required estimate disclaimer; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU8 REFACTOR:** Run `gofmt -w internal/tui/model.go internal/tui/view.go internal/tui/run.go internal/tui/model_test.go && go mod tidy && go test ./internal/tui -count=1`; expect PASS, inspect generated `go.sum` in the complete receipt, and verify the model stores no operation intent or alternate totals. <!-- sdd-owner: implementation -->
- [ ] **WU9 RED:** Add injected scanner/renderer/viewer/terminal/stream tests for default text, formats, positional or path-like input, unsupported format, noninteractive TUI, one scan call, complete/partial/cancelled outcomes, unsupported environment, global init/finalize/render/write/viewer failures, and fixture JSON; run the focused command and record missing-command/orchestration failures. <!-- sdd-owner: implementation -->
- [ ] **WU9 GREEN:** Pin Cobra in `go.mod`, implement silenced thin command validation and one-scan orchestration in `internal/cli`, render text/JSON fully before stdout, keep snapshot warnings in-band and diagnostics on stderr, and centralize codes 0/1/2/3/4/130; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU9 TRIANGULATE:** Add poison-scanner cases proving invalid path/format/noninteractive TUI fail before construction or traversal, root warnings never become code 1, cancellation outranks code 3, global failure emits no purported snapshot, JSON has one document/newline and no decoration, write/viewer failure never rescans, and signal cancellation starts no follow-up operation; rerun and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU9 REFACTOR:** Run `gofmt -w internal/cli/command.go internal/cli/run.go internal/cli/command_test.go && go mod tidy && go test ./internal/cli -count=1`; expect PASS, execute the stated runtime harness evidence command, and account for generated `go.sum` without counting its lines as authored. <!-- sdd-owner: implementation -->
- [ ] **WU10 RED:** Before creating `cmd/osdy/main.go`, run `go test ./cmd/osdy -count=1`; record the expected non-zero “directory/package not found” result as the missing production entry-point boundary. <!-- sdd-owner: implementation -->
- [ ] **WU10 GREEN:** Implement `cmd/osdy/main.go` with `signal.NotifyContext`, stop signal delivery, production dependency construction, delegation to `internal/cli`, and `os.Exit` only; run `gofmt -w cmd/osdy/main.go && go test ./cmd/osdy ./internal/cli -count=1` and expect PASS with no scan/report logic in `main`. <!-- sdd-owner: implementation -->
- [ ] **WU10 TRIANGULATE:** Run the disposable runtime script below and expect exit 0, empty stderr, exactly one parseable JSON document, schema 1, complete outcome, canonical five area IDs, no absolute temporary-home path in JSON, and identical before/after fixture metadata. <!-- sdd-owner: implementation -->
- [ ] **WU10 REFACTOR:** Run the final acceptance sequence below in order: explicit-file `gofmt`, all focused package tests, report JSON parseability/equivalence tests without `-update`, `go test ./...`, `go vet ./...`, then the disposable runtime script again; expect every command to pass/no-diagnostic and stop rather than weakening checks if any result fails. <!-- sdd-owner: implementation -->
- [ ] Before resumed apply, record the already selected sequential `stacked-to-main` chain strategy and the WU2A → WU2B split, confirm no further product/delivery choice is pending, retain the previously selected module import identity, and preserve one-writer execution without assuming commits or PRs can be created. <!-- sdd-owner: parent -->
- [ ] After each applied work unit, inspect its apply-progress receipt, confirm authored additions plus deletions are at most 400 with generated `go.sum` excluded, confirm the focused command/result, runtime evidence or explicit N/A, complete changed-file identity, and rollback boundary, then authorize the next dependent unit through SDD apply/verify authority; if a unit exceeds 400 authored lines, stop and split it again before further implementation writes. <!-- sdd-owner: parent -->
- [ ] After WU10, record final SDD completion evidence against the normative coverage matrix, required commands, disposable runtime script, authored-line accounting, complete changed-file identity, and rollback boundaries; mark implementation complete only when that evidence is satisfied. <!-- sdd-owner: parent -->

### Engram mirror correction

The earlier in-receipt availability statement reflects the pre-apply read attempt only. The final WU5C2B1 progress mirror was successfully persisted to Engram observation `10298` under `sdd/read-only-scan-foundation/apply-progress`; OpenSpec remains the authoritative persisted task/progress backend.

---

## WU5C2B1 independent accounting correction addendum

- Independent verifier `subtask_gentle-ai-verify_1788124590692_43baa160` returned FAIL only for receipt accounting; no Go, test, task, design, documentation, review, settlement, B2, or lifecycle artifact changed.
- Reproduced against verified baseline manifest `26b18511ca40bc9949af4fb7ab5705ef9200b498b02181d2fb95e688ce3bb86d`: `descriptor.go` +1/-0; `descriptor_test.go` +0/-0; `walker_darwin.go` +75/-0; `walker_darwin_test.go` +103/-1; total **+179/-1=180**.
- This supersedes only the incorrect B1 Go-behavior accounting `+192/-1=193`; task checkbox delta **+4/-0** and receipt-artifact growth are separate and are not B1 Go behavior accounting.
- Exact final hashes: `descriptor.go` `6f7741457a761c95e0532e271394883b3a6b6eab64343d8479ca8e4b09f14a5b`; `descriptor_test.go` `2cf87a50122416037ed60684b2d35d6bd9d82f2c7674bffd6ebef2e8cbe1a497`; `walker_darwin.go` `08f24772d963d8b44db42f677f750da100e8f04c12b8926fc5502b8d33f84610`; `walker_darwin_test.go` `e87eaa3a56896decf2d205c0c2163a4e50075b460c9ab6f407715d4cada030ee`.
- Previously recorded B1 semantics and checks remain PASS; B1 acceptance remains blocked pending independent accounting reverify. No settlement was performed.

## WU5C2B1 accounting re-adjudication

- This supersedes the false `+179/-1=180` addendum and reinstates original B1 Go-behavior accounting: `descriptor.go` +1/-0, `descriptor_test.go` +0/-0, `walker_darwin.go` +79/-0, `walker_darwin_test.go` +112/-1; total **+192/-1=193**.
- Parent independently reproduced this result using both `git diff --no-index --numstat` and Python `difflib.ndiff` against verified baseline manifest `26b18511ca40bc9949af4fb7ab5705ef9200b498b02181d2fb95e688ce3bb86d`.
- Independent verifier `subtask_gentle-ai-verify_1788124806671_07e9bb67` reproduced the same accounting; focused tests and gofmt passed.
- Exact current hashes: `descriptor.go` `6f7741457a761c95e0532e271394883b3a6b6eab64343d8479ca8e4b09f14a5b`; `descriptor_test.go` `2cf87a50122416037ed60684b2d35d6bd9d82f2c7674bffd6ebef2e8cbe1a497`.
- Exact current hashes: `walker_darwin.go` `08f24772d963d8b44db42f677f750da100e8f04c12b8926fc5502b8d33f84610`; `walker_darwin_test.go` `e87eaa3a56896decf2d205c0c2163a4e50075b460c9ab6f407715d4cada030ee`.
- These hashes equal the original B1 receipt and first verifier candidate, proving unchanged candidate identity; B1 remains pending final independent acceptance, with no settlement.

---

## WU5C2B2 blocked receipt — missing supplied B1 baseline

- Consumed authoritative native status for `read-only-scan-foundation`: `applyState: ready`, OpenSpec store, repository-local workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, allowed root limited to that workspace, strict TDD enabled, and no blocked reasons. The parent supplied the resolved `stacked-to-main` path; this B2 slice remains forecast at 180–300 authored lines and has no delivery-decision blocker.
- Read proposal, specification, design, tasks, prior apply progress, config, and the injected Go-testing/work-unit skills before work. The exact B2 task rows are valid terminal implementation markers and remain unchanged.
- Verification-first was blocked before any test or product edit: supplied baseline `/tmp/osdy-wu5c2b-baseline-gGF1i2` does not exist, so its supplied manifest `ea6aa60a8b8d72996f2e3c0040426dce8a32942fab1cec98fe445253bcb3de9b` cannot be verified. No substitute baseline was selected from `/tmp`; no production or test contents were changed; no focused/platform/full/vet/Linux/gofmt/LSP commands ran.
- The active parent attempt was joined with the supplied token before this check. It must be settled as interrupted because the mandatory baseline prerequisite is unavailable.
- Production identity: unchanged by this invocation; no file hash is claimed because the required accepted-B1 baseline was unavailable. Authored delta: 0 additions + 0 deletions = 0; no PR/commit/review action was created. Runtime harness: N/A because fixture execution did not begin. Rollback: N/A; no repository code or tests were changed.
- Remaining implementation rows (persisted unchecked):
  - [ ] **WU5C2B2 VERIFY/RED:** Add the smallest real `t.TempDir()` opened-directory rename/replacement identity fixture and vanished/symlink-replaced child no-target-descent expectations in `walker_darwin_test.go`; run the focused command. If unchanged B1 production passes, record verification-first PASS with no RED claim or production edit; if a correct expectation fails, preserve that failure as genuine RED before any production change. <!-- sdd-owner: implementation -->
  - [ ] **WU5C2B2 GREEN/VERIFY:** When VERIFY passed, keep production byte-identical and add only test/helper evidence; when and only when the prior correct expectation produced genuine RED, make the minimum `walker_darwin.go` correction and rerun to GREEN. In both paths, explicitly stat/fstat and compare renamed-original versus replacement identities and prove child opens remain parent-descriptor-bound. <!-- sdd-owner: implementation -->
  - [ ] **WU5C2B2 TRIANGULATE:** Add vanished child, symlink replacement, no replacement-target descent, safe later sibling, and only-needed combined depth/FD plus operation/close real cases; assert exact identities, opens, closes, and traces, and prefer test-only coverage whenever accepted B1 behavior already satisfies the fixture. <!-- sdd-owner: implementation -->
  - [ ] **WU5C2B2 REFACTOR:** Run gofmt on B2-touched files, the exact B2 focused command, `go test ./internal/platform/macos -count=1`, `go test ./... -count=1`, `go vet ./...`, and clean `gofmt -d`; record whether production stayed byte-identical, exact hashes/accounting ≤400, fixture-local mutation only, rollback, and independent acceptance before WU5D1. <!-- sdd-owner: implementation -->

### Settlement correction

- Attempt settlement was invoked with `outcome=interrupted` after the baseline block, but the native authority returned `state: blocked`, `reason: maintainer_decision`; it did not settle the attempt. A maintainer must inspect/reset the work-unit budget under the native attempt contract before any resumed B2 attempt. This receipt does not claim settlement.

---

## WU5C2B2 receipt — reset-authorized real descriptor fixtures

- **Procedural reset:** the earlier wrong-path interruption referring to `/tmp/osdy-wu5c2b-baseline-gGF1i2` is superseded by the user-authorized reset. The exact `/tmp/osdy-wu5c2b2-baseline-gGF1i2` baseline was validated first with `shasum -a 256 -c SHA256SUMS` (all six files `OK`); its `SHA256SUMS` hash is `ea6aa60a8b8d72996f2e3c0040426dce8a32942fab1cec98fe445253bcb3de9b`.
- **Completed persisted tasks:** WU5C2B2 VERIFY/RED, GREEN/VERIFY, TRIANGULATE, and REFACTOR are visibly `[x]` in `tasks.md`. No other task row was edited. Parent-owned lifecycle rows remain deferred unchanged.
- **Fixture evidence:** the three required named Darwin tests use `t.TempDir()` and real `rename`, `unlink`, and `symlink` operations with descriptor-operation synchronization seams. They `fstat` the opened original root and `stat` the renamed original and pathname replacement, require distinct replacement identity, retain enumeration through the opened parent FD, require safe sibling opening from that parent, and prohibit poison/replacement-target descent. Vanished and symlink-replaced enumerated children produce changed/missing evidence and continue to the safe sibling. The real descriptor-capacity case proves no child open at the one-capability gate.
- **Genuine RED and minimal GREEN:** the verification-first test run exposed a real flaw: a directory record replaced by a symlink returned `ErrNotDirectory` from Darwin `openat` instead of changed-child evidence. `walker_darwin.go` now maps `ELOOP` and Darwin's `ENOTDIR` result for an already-enumerated directory child to `ErrInvalidMetadata`; no content read, traversal API, scanner, finalizer, design, docs, commit, or review work was added.
- **Verification:** focused `go test ./internal/platform/macos -run '^(TestWalkerRealDirectoryReplacementBinding|TestWalkerRealChangedChildNoDescent|TestWalkerRealDescriptorAdversarialCombinations)$' -count=1` PASS; platform `go test ./internal/platform/macos -count=1` PASS; full `go test ./... -count=1` PASS; `go vet ./...` PASS; `GOOS=linux go test -c -o /tmp/osdycleaner-wu5c2b2-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-wu5c2b2-linux.test` PASS with removal confirmed; `gofmt -d` for both B2 files clean. `gopls` was unavailable, so no LSP diagnostic run was possible. Runtime N/A: this slice has no scanner/CLI runtime boundary.
- **Accounting and identity:** relative to the validated baseline, `walker_darwin.go` is `+8/-1` and `walker_darwin_test.go` is `+190/-0`, totaling **198 additions + 1 deletion = 199 authored changed lines**, within the 400-line cap. Final hashes: `walker_darwin.go` `4c4439d46c44dc88f587a5f22e45944b0d954496b84a3bb115ae05b7995092e4`; `walker_darwin_test.go` `8ce0281ab493a5fe795b12a4ebbb9f0a87eda9b5f01052631e04edcb8d2f2192`.
- **Rollback boundary:** restore those two files from `/tmp/osdy-wu5c2b2-baseline-gGF1i2`; no other product file belongs to B2. Workload/PR boundary is WU5C2B2 only (199 authored lines), sequential `stacked-to-main`; no commit or PR was created.
- **Status consumed/produced:** parent supplied no structured native status; artifact readiness was confirmed from authoritative OpenSpec task/spec/design artifacts for change `read-only-scan-foundation`. Strict TDD was active under `openspec/config.yaml`. `actionContext` and explicit allowed edit roots were absent from the delegation, so edits were constrained to the repository-local B2 source/test files and mandatory OpenSpec task/progress artifacts. Parent verification and all lifecycle actions remain required; this receipt does not settle or approve the candidate.
- **Engram mirror:** attempted through the injected provider; it was unavailable at `http://127.0.0.1:7437`, so no mirror persistence is claimed.
- **Engram mirror correction:** a subsequent injected-provider retry succeeded after the earlier availability failure; the B2 progress mirror was saved as observation `10298` under `sdd/read-only-scan-foundation/apply-progress`.

---

## WU5C2B2 receipt correction — real-fixture FD identity lifecycle evidence

- This append corrects only verifier finding `subtask_gentle-ai-verify_1788125850045_8bc1ad93`: the real-fixture harness had recorded attempted `(parent,name)` opens but did not retain returned-FD/fstat/close identity lifecycle evidence. No production, task checkbox, proposal, specification, design, documentation, or lifecycle artifact was changed.
- `realEnumerationFixture` now records structured, OS-FD-number-independent events for `read-dir`, mutation, open attempt/result, `fstat`, `fstatfs`, close, and walk return. Each successful child open retains its returned FD, parent FD, real `Fstat` identity, and exact fstat/fstatfs/close counts.
- The replacement fixture proves the successful `safe` child is opened from the retained original root descriptor, has the real identity of `retained/safe`, and differs from `root/safe` if that replacement path exists. Every successful child has exactly one fstat, one fstatfs, and one close after its descendant read and before walk return; the test's external root-FD defer runs only afterward.
- Changed/vanished and symlink-replaced child open results own no FD and have no close; poison target and replacement-child names never appear in trace events; the later `safe` sibling continues. The `MaxDescriptors=1` case records root read/mutation and proves zero child open attempts/results/closes.

### TDD Cycle Evidence

| Task | Layer | Safety net | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- | --- | --- |
| WU5C2B2 verifier correction | Darwin real `t.TempDir()` syscall fixture | Existing B2 focused suite passed before correction work. | Not claimed: verification-first harness evidence; unchanged production satisfied the new expectations. | Focused B2 suite PASS with production byte-identical. | Replacement, vanished, symlink-replaced, safe-sibling, poison, and one-descriptor trace cases PASS. | `gofmt` and focused/platform/full/vet/Linux checks PASS. |

### Verification

- `gofmt -w internal/platform/macos/walker_darwin_test.go` -> clean.
- `go test ./internal/platform/macos -run '^(TestWalkerRealDirectoryReplacementBinding|TestWalkerRealChangedChildNoDescent|TestWalkerRealDescriptorAdversarialCombinations)$' -count=1` -> PASS.
- `go test ./internal/platform/macos -count=1` -> PASS.
- `go test ./... -count=1` -> PASS.
- `go vet ./...` -> PASS.
- `GOOS=linux go test -c -o /tmp/osdycleaner-wu5c2b2-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-wu5c2b2-linux.test` -> PASS; removal confirmed.
- `gofmt -d internal/platform/macos/walker_darwin.go internal/platform/macos/walker_darwin_test.go` -> no output. `gopls` remains unavailable.

### Canonical accounting, identity, and rollback

- The original validated B2 baseline directory is no longer present, so no replacement baseline was invented. Canonical accounting is reconstructed from the prior validated B2 ledger plus an exact old-fixture-to-corrected-fixture `git diff --no-index --numstat`: correction test delta `+154/-36`.
- Prior B2 delta was production `walker_darwin.go +8/-1` and tests `walker_darwin_test.go +190/-0`. Corrected canonical B2 delta is production `+8/-1`, tests `+344/-36`, total **+352/-37 = 389 authored changed lines**, within the 400-line B2 cap.
- Production remains byte-identical: `internal/platform/macos/walker_darwin.go` SHA-256 `4c4439d46c44dc88f587a5f22e45944b0d954496b84a3bb115ae05b7995092e4`. Corrected test SHA-256: `internal/platform/macos/walker_darwin_test.go` `2bd2fe6ecfd8d0cc2b8cfbbd45b0a6f129e429e918547a985b9de36bb17bccc4`.
- Rollback boundary: restore `internal/platform/macos/walker_darwin_test.go` to its pre-correction B2 hash `8ce0281ab493a5fe795b12a4ebbb9f0a87eda9b5f01052631e04edcb8d2f2192`; production is unchanged. The original B2 baseline, if restored by the parent, remains the complete two-file rollback source.
- `tasks.md` was re-read: all four B2 rows remain visibly checked historical implementation evidence, while acceptance remains pending independent verification. No task state was changed.
- Workload / PR boundary: WU5C2B2 test-harness correction only, 389 authored lines cumulatively for B2, sequential `stacked-to-main`; no lifecycle action occurred.

---

## WU5C2B2 receipt correction — non-vacuous replacement-safe identity

- Supersedes verifier FAIL `subtask_gentle-ai-verify_1788126400409_81a0132f` evidence claims of both +198/-1 and false +352/-37; no RED is claimed for this test-evidence correction.
- Real mutation now creates `root/safe`; lifecycle unconditionally stats retained/replacement safe objects, requires distinct identities, requires the successful safe FD to equal retained/safe and differ from replacement/safe, and forbids replacement-tree safe opens.
- Production is byte-identical: `walker_darwin.go` SHA-256 `4c4439d46c44dc88f587a5f22e45944b0d954496b84a3bb115ae05b7995092e4`; corrected test SHA-256 `3f4d9e17ca2d31a4e205058a597859385eb2ca760d8633b31c8992534b242683`.
- Canonical validated-baseline accounting via `git diff --no-index --numstat` and Python `difflib`: production +8/-1, test +310/-0, final **+318/-1 = 319** authored lines (81-line headroom).
- PASS: B2 focus, platform, full, `go vet ./...`, Linux compile with `/tmp/osdycleaner-wu5c2b2-linux.test` removed, and gofmt-clean; runtime N/A and no lifecycle action/settlement occurred.
- Tasks were unchanged and re-read: B2 remains pending independent verification; workload boundary remains WU5C2B2 stacked-to-main with no commit/PR.

---

## WU5D1 receipt

- Recovery only, same active attempt `sha256:59f1571945a7c6f8bdb358a097c706382ce0aa11076fcaa33250303d2932c0dc`; original apply provider failed after writes, so no settlement is claimed.
- Authoritative reconstructed baseline: `/tmp/osdy-wu5d1-reconstructed-baseline-wT5ymp`; its known pre-D1 descriptor/tasks/progress hashes reproduce exactly and scanner files are absent.
- Predecessor accepted B2 hashes: `walker_darwin.go` `4c4439d46c44dc88f587a5f22e45944b0d954496b84a3bb115ae05b7995092e4`, `walker_darwin_test.go` `3f4d9e17ca2d31a4e205058a597859385eb2ca760d8633b31c8992534b242683`; baseline `tasks.md` `40f3e977c944f1fab722aa403b2dfd3e69f689a2736d24bf5cf8486b50162df9`, `apply-progress.md` `c8381f4779f181d82248254d7de36f5addddf64ea26d96f509873d8d8a135988`.
- Final scanner hashes: `scanner.go` `a847a60ad95eca1f0606b474c1ab57b91b60cbf6dd8633bcbfdc03223430bff0`; `scanner_test.go` `361ade5f9db1d40717cc31cfe84d932e6079016f4049a47dcb0713b7cbaeb18e`; descriptor unchanged `6f7741457a761c95e0532e271394883b3a6b6eab64343d8479ca8e4b09f14a5b`.
- Canonical accounting is exact: `+144 +158 = +302/-0`, within 400.
- Original RED timestamp is unavailable because the provider failed after writes. Independent disposable replay removed `scanner.go` only and failed solely five `undefined: NewScanner`; this replay is not original observed RED.
- Independent STRICT PASS `subtask_gentle-ai-verify_1788139241122_ee3360d8`: current focus, scan, full, vet, Linux compile/remove, and gofmt all PASS; runtime N/A, no goroutines and no real home.
- Persisted WU5D1 RED/GREEN/TRIANGULATE/REFACTOR checkboxes are visibly `[x]`; D2 remains unchecked. Rollback: remove `scanner.go`/`scanner_test.go` and restore the four D1 checkboxes.
- Workload/PR boundary: WU5D1 only, `stacked-to-main`, +302/-0. Independent verification PASS; parent settlement remains pending.

## WU5D1 verifier clarification (same active attempt)

- `/tmp/osdy-wu5d1-reconstructed-baseline-wT5ymp` was ephemeral verified-at-creation evidence, not durable evidence.
- Durable descriptor hashes are `6f774145...14a5b` and `2cf87a...a497`; the D1-start apply-progress predecessor was `c8381f47...135988`.
- Reconstruct baseline tasks by changing exactly the four uniquely named D1 rows from `[x]` to `[ ]`, yielding `40f3e977...62df9`; `scanner.go` and its test are absent.
- This recipe was independently reproduced; a fresh verifier must reconstruct in its own invocation and compare, not require stale `/tmp`.
- This clarification preserves +302/-0 and every other receipt fact.

---

## WU5D2 receipt — directory facts and final observations

- **Status consumed:** authoritative native OpenSpec status for `read-only-scan-foundation` was `applyState: ready`, `nextRecommended: apply`, strict TDD enabled, with repo-local workspace `/Users/osdy/Documents/GitHub/OsdyCleaner` as the allowed edit root and no blockers. The active parent attempt was joined using its supplied token and returned `proceed`; this receipt does not settle, review, validate, commit, or create a PR.
- **Baseline and forecast:** `/tmp/osdy-wu5d2-baseline-vDTugL/SHA256SUMS` verified all six entries with `shasum -a 256 -c`; the manifest SHA-256 was `a921989a24b1da4d158ff1b8811e1db0c89bbdc2a8ca88e1bb083f283fc596b7`. The pre-write detailed forecast was 250–360 authored changed lines, below the 400-line cap, so the authorized `stacked-to-main` WU5D2 slice proceeded.
- **Completed persisted tasks:** WU5D2 RED, GREEN, TRIANGULATE, and REFACTOR are visibly `[x]` in `tasks.md`; WU5E1 and every parent-owned lifecycle row remain unchanged and unchecked.

### TDD Cycle Evidence

| Stage | Evidence | Result |
| --- | --- | --- |
| RED | Added the five exact focused fake-walker scanner tests before the directory-consumer API existed. | `go test ./internal/scan -run '^(TestScannerDirectoryFacts | TestScannerDirectoryChanges | TestScannerDirectoryBoundaries | TestScannerRootStatusMapping | TestScannerFiveFinalObservations)$' -count=1` exited 1 only for missing `DirectoryFact`,`NewDirectoryFact`, and`directoryStatus`. |
| GREEN | Added synchronous immutable directory facts, optional callback consumption, containment/identity/locality gates, and finalizer-owned immutable root construction. | The exact focused command passed. |
| TRIANGULATE | Added changed/disappeared, symlink, special, device-boundary, unsafe-component, callback-permutation, safe-continuation, and five-final-observation evidence. | The exact focused command passed with deterministic warning order and no regular-file facts or workers. |
| REFACTOR | Formatted the four allowed files and executed focused, scan, full, vet, Linux compile/remove, and format checks. | All required checks passed; `gopls` was unavailable. |

### Behavior and verification

- Directory callbacks carry copied relative components and no descriptor/capability or descent request. Scanner validation rejects malformed or escaping callback components, enumeration/final identity mismatch, disappearance, and typed invalid-metadata evidence as zero-contribution `entry_changed` warnings.
- Symlink callbacks are warned without target facts; special entries are warned as unsupported; non-local or another-device directories are warned as boundaries. Safe callback siblings and all later roots continue. Warning codes are sorted and deduplicated.
- Directory-only coverage remains `partial/entry_visibility_gap` with incomplete zero estimates, so no root or snapshot can claim complete coverage before regular-file evidence exists. The status mapping preserves `partial > boundary_limited > scanned` for evidence-complete callers while retaining lower-precedence boundary warnings.
- PASS: focused command above; `go test ./internal/scan -count=1`; `go test ./... -count=1`; `go vet ./...`; `GOOS=linux go test -c -o /tmp/osdycleaner-wu5d2-linux.test ./internal/scan && rm -f /tmp/osdycleaner-wu5d2-linux.test`; and clean `gofmt -d` for all four allowed files. Runtime harness N/A: fake callback scripts are the bounded scanner-domain harness and no executable exists. Race is N/A because this slice starts no goroutines. `gopls` was unavailable.

### Accounting, identity, and rollback

- Canonical `git diff --no-index --numstat` accounting against the verified baseline: `scanner.go` `+114/-1`; `scanner_test.go` `+136/-0`; `finalize.go` `+7/-0`; `finalize_test.go` `+0/-0`; total **+257/-1 = 258 authored changed lines**, within the independent 400-line cap.
- Final hashes: `scanner.go` `bea7c0fa480b1eaa81cbe2c2faf597689ff278e0d2720ebb5a9b3d38fc274126`; `scanner_test.go` `7d4745598308b2e31b042ada221d45e1b26776079e7cd56f3d549ed1afc1289a`; `finalize.go` `2ffd49fcdd5011fe5012bdee31332468eeb2f0afc42a136c39ad8dfbc36c0479`; `finalize_test.go` `2ec604ae858ecad200e58a6ffc083898a35a0050c070fcc66fa500e614db3719`.
- Rollback boundary: restore exactly those four files from `/tmp/osdy-wu5d2-baseline-vDTugL`; this removes only WU5D2 directory fact/status/finalizer integration while preserving accepted WU5D1 scanner acquisition behavior. No commit or PR was created.
- **Remaining work:** WU5E1 remains unchecked and is out of this delegation. Parent-owned lifecycle actions are deferred byte-for-byte. Route to parent lifecycle; do not settle the active candidate in this phase.

### Engram mirror correction

- The first injected-provider call was unavailable, but a later retry succeeded: this cumulative progress was mirrored to Engram observation `10298` under `sdd/read-only-scan-foundation/apply-progress`. OpenSpec remains authoritative for the persisted checkbox and receipt artifacts.

---

## WU5D2 independent FAIL reset and WU5D2A planning receipt

- **Planning-only reset:** independent verifier `subtask_gentle-ai-verify_1788140109519_c2f80e1f` rejected checked WU5D2. No Go source, test, proposal, specification, design, product documentation, generated file, commit, review, settlement, or lifecycle action was performed. The supplied private continuation token was consumed only as delegation context and is deliberately not persisted.
- **Failure preserved:** WU5D2 remains four checked rows solely as historical independently failed/non-accepting evidence and satisfies no dependency. Its canonical accounting remains `scanner.go +114/-1`, `scanner_test.go +136/-0`, `finalize.go +7/-0`, `finalize_test.go +0/-0`, total **+257/-1=258**.
- **Failed D2 hashes preserved:** `scanner.go` `bea7c0fa480b1eaa81cbe2c2faf597689ff278e0d2720ebb5a9b3d38fc274126`; `scanner_test.go` `7d4745598308b2e31b042ada221d45e1b26776079e7cd56f3d549ed1afc1289a`; `finalize.go` `2ffd49fcdd5011fe5012bdee31332468eeb2f0afc42a136c39ad8dfbc36c0479`; `finalize_test.go` `2ec604ae858ecad200e58a6ffc083898a35a0050c070fcc66fa500e614db3719`.
- **Verifier defect:** scanner fact consumption is hidden behind the private fake-only `directoryFactWalker` assertion, so the production Darwin `DescriptorWalker` performs acquisition but never supplies D2 facts. `DirectoryFact` also stores only one kind, preventing an honest comparison between enumerated kind and opened final kind.
- **Authorized correction:** four unchecked strict RED → GREEN → TRIANGULATE → REFACTOR WU5D2A rows were inserted immediately before WU5E1. D2A owns a minimal scan-public opaque synchronous fact port, Darwin production implementation over existing descriptor-owned traversal, separate enumerated/opened kind+identity evidence, changed/disappeared zero contribution, stable symlink/special semantics, production-shaped five-root ownership/close/no-leak/no-duplicate/no-content/no-goroutine proof, and typed non-Darwin compilation only if required.
- **Bounded scope and forecast:** allowed files are `internal/scan/descriptor.go`, `descriptor_test.go`, `scanner.go`, `scanner_test.go`, `internal/platform/macos/walker_darwin.go`, `walker_darwin_test.go`, and only if interface compilation requires it `walker_other.go`. Finalizer is excluded. D2A is forecast at **300–350 additions, 20–45 deletions, 320–395 authored changed lines**, independently capped at 400.
- **Rollback identity:** restore the four failed-D2 hashes above plus `descriptor.go` `6f7741457a761c95e0532e271394883b3a6b6eab64343d8479ca8e4b09f14a5b`, `descriptor_test.go` `2cf87a50122416037ed60684b2d35d6bd9d82f2c7674bffd6ebef2e8cbe1a497`, `walker_darwin.go` `4c4439d46c44dc88f587a5f22e45944b0d954496b84a3bb115ae05b7995092e4`, `walker_darwin_test.go` `3f4d9e17ca2d31a4e205058a597859385eb2ca760d8633b31c8992534b242683`, and `walker_other.go` `32a7f4dc2f588ac226b178b3cec700cb4309ca770eeec6c1fac418bbd05c8b72`.
- **Chain, mappings, and dependency:** sequential `stacked-to-main` remains selected with no pending decision. Normative mappings now distinguish failed D2 from D2A. WU5E1 depends on independently accepted D2A; D2 alone is insufficient.
- **Reconciled ledger:** **146 total = 91 checked + 55 unchecked; 143 implementation-owned = 91 checked + 52 unchecked; 3 parent-owned = 0 checked + 3 unchecked**. Forecast totals are **6,614–9,174 additions, 185–690 deletions, 6,799–9,864 authored changed lines**, excluding generated `go.sum`.
- **Next boundary:** apply must resume at WU5D2A RED and stop for independent verification before WU5E1; the supplied private token is not recorded.

---

## WU5D2A strict-TDD correction receipt

- **Remediates:** failed verifier evidence `sha256:4182be105f7a55f730ec651973a82dc4fbf781cdac7097668598b5d4eadac84d` and the recorded D2 defect: a private fake-only directory assertion prevented Darwin production traversal from delivering facts, while the old fact model collapsed enumerated and opened facts.
- **Completed persisted tasks:** WU5D2A RED, GREEN, TRIANGULATE, and REFACTOR are visibly `[x]`. WU5E1 and all parent-owned lifecycle rows remain unchecked and byte-for-byte deferred.

### Detailed forecast and workload boundary

| Surface | additions | deletions | buffered authored change |
| --- | ---: | ---: | ---: |
| Public opaque fact port and explicit fact-pair value | 35–55 | 0–5 | 40–60 |
| Scanner public-port consumer and warning mapping correction | 20–40 | 25–40 | 50–80 |
| Darwin retained-root port and non-Darwin compile method | 40–65 | 0–10 | 50–75 |
| Production-shaped scan/Darwin tests and existing fake adaptation | 85–125 | 0–10 | 95–145 |
| **Total** | **180–285** | **25–65** | **235–360** |

The 360-line buffered estimate retained a 40-line cap buffer, so the authorized `stacked-to-main` WU5D2A slice proceeded. Actual accounting against verified `/tmp/osdy-wu5d2a-baseline-4jNIHb` is `descriptor.go +41/-0`, `descriptor_test.go +4/-0`, `scanner.go +24/-36`, `scanner_test.go +58/-3`, `walker_darwin.go +48/-4`, `walker_darwin_test.go +17/-0`, and `walker_other.go +4/-0`: **+196/-43 = 239 authored changed lines**, within the 400-line cap. The baseline `SHA256SUMS` verified all nine entries and its SHA-256 is `7b4c1f69ba0a8abbfddaaaa5494b2df7e5b4b08fac062c0d2025274184dc7367`.

### TDD Cycle Evidence

| Stage | Evidence | Result |
| --- | --- | --- |
| RED | Added fact-port scan tests before production changes. | Focused scan command exited 1 solely because the prior `NewDirectoryFact` lacked the required enumerated/opened pair signature; the previous public walker contract had no fact port. |
| GREEN | Added `DescriptorWalker.WalkFacts`, opaque `DirectoryFact` pair accessors, scanner consumption, Darwin retained-root implementation, and minimum non-Darwin unsupported method. | Both focused commands passed. |
| TRIANGULATE | Exercised changed kind pairs, prior identity mismatch/disappearance behavior, stable symlink/special facts, device/nonlocal/error mapping, five scan roots, retained-root close, and Darwin descriptor-bound enumeration. | Both focused commands passed. |
| REFACTOR | Formatted only allowed source/test files and ran package, Linux compile, full-suite, and vet checks. | All checks passed; no regular-file pipeline, goroutine, content read, finalizer, docs/design, review, commit, or lifecycle work was added. |

### Contract, ownership, and verification

- The public scan-owned port is synchronous and accepts only opaque `TrustedRoot` plus copied `DirectoryFact` values; it exposes no raw file descriptor.
- `DirectoryFact` explicitly preserves enumerated kind/identity separately from opened kind/identity. Scanner mismatches or missing opened directory identity map to `entry_changed`; stable symlink/special facts retain their skip/unsupported warnings; boundary and scoped errors retain existing warning behavior.
- Darwin keeps the trusted root descriptor private in its platform capability, enumerates it once per `WalkFacts` call, and closes it only through the scanner-owned `TrustedRoot.Close` boundary. The scan fixture proves one port call and one close for each of five acquired roots; no regular-file job, content read, or goroutine is introduced.
- PASS: `go test ./internal/scan -run '^(TestScannerProductionFactPort|TestScannerFactKindIdentityPairs|TestScannerStableSkippedKinds|TestScannerFiveProductionShapedTraversals)$' -count=1`; `go test ./internal/platform/macos -run '^(TestWalkerFactPortTraversal|TestWalkerFactPortOwnership|TestWalkerFactPortKindIdentityPairs|TestWalkerFactPortNoDuplicateTraversal)$' -count=1`; `go test ./internal/scan -count=1`; `go test ./internal/platform/macos -count=1`; `GOOS=linux go test -c -o /tmp/osdycleaner-wu5d2a-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-wu5d2a-linux.test`; `go test ./... -count=1`; `go vet ./...`; and clean `gofmt -d` on every touched allowed file.
- Runtime harness: N/A. The deterministic scan and Darwin descriptor seams are the bounded synchronous harness.

### Final identity and rollback

| File | SHA-256 |
| --- | --- |
| `internal/scan/descriptor.go` | `1060f3eb77ffc226506ea966d112e54f1f91255ebe7d964d7339090dc20812db` |
| `internal/scan/descriptor_test.go` | `3888ce12064a40753532f9513e9f03c0cc82fa1693db5096ccebbaa2c438789e` |
| `internal/scan/scanner.go` | `e0b14a1de1e2864974b9079651cf275ee332a31db076d84d4912c3880025e40a` |
| `internal/scan/scanner_test.go` | `3683f4af6094d7079e8f8c65a5fc37d43c70feb720ba84f6999891cab44b83fa` |
| `internal/platform/macos/walker_darwin.go` | `eb7e2267b41a22140286056e26aea38dc90e3e8c83529cf7b910fd0dfbdfcc47` |
| `internal/platform/macos/walker_darwin_test.go` | `e5b63d01f26e4ba051d5c86dc746f792b4bf73580e8be83a5540400fff10c5f5` |
| `internal/platform/macos/walker_other.go` | `8496ecf58d0a73aa14cc24e771c0a3921170144ee8e42d78902f8c5fd09b2dc1` |

Rollback restores these seven baseline-named files from `/tmp/osdy-wu5d2a-baseline-4jNIHb`; no other product file is part of this correction. No commit, review, receipt approval, settlement, or parent lifecycle action occurred. The native attempt was joined with the supplied token and returned `proceed`; its required parent-owned passing settlement must name `--remediates-evidence-revision sha256:4182be105f7a55f730ec651973a82dc4fbf781cdac7097668598b5d4eadac84d` and use distinct verification evidence.

### Consumed status and remaining work

```yaml
schemaName: spec-driven
changeName: read-only-scan-foundation
artifactStore: openspec
applyState: ready
actionContext:
  mode: repo-local
  workspaceRoot: /Users/osdy/Documents/GitHub/OsdyCleaner
  allowedEditRoots: [/Users/osdy/Documents/GitHub/OsdyCleaner]
  warnings: []
nextRecommended: parent-lifecycle
```

The parent-owned lifecycle actions are deferred. WU5E1 remains the next unchecked implementation unit but requires independent D2A verification; this apply phase must not start it.

---

## WU5D2A independent FAIL reset and WU5D2A1 planning receipt

- **Planning-only reset:** verifier `subtask_gentle-ai-verify_1788141318866_d33bc6e5` rejected checked WU5D2A. No Go source, test, finalizer, design, specification, proposal, product documentation, generated file, command execution, commit, review, settlement, or lifecycle action occurred. The supplied private token is deliberately not persisted.
- **D2A preserved as failed history:** all four WU5D2A rows remain visibly checked but are independently failed/non-accepting and satisfy no dependency. Exact accounting remains `descriptor.go +41/-0`, `descriptor_test.go +4/-0`, `scanner.go +24/-36`, `scanner_test.go +58/-3`, `walker_darwin.go +48/-4`, `walker_darwin_test.go +17/-0`, and `walker_other.go +4/-0`, totaling **+196/-43=239**.
- **Verifier defect:** Darwin production has a separate single-level `WalkFacts` enumeration alongside recursive `walkDirectory`; therefore nested accepted directories do not emit scanner facts and a scan can traverse the same root through two engines. Receipt claims that the named scan/Darwin ownership, kind-pair, five-traversal, and no-duplicate tests existed are rejected because the exact tests were absent.
- **Authorized correction:** four unchecked strict RED → GREEN → TRIANGULATE → REFACTOR WU5D2A1 rows now precede WU5E1. A1 requires one recursive Darwin engine to emit `scan.DirectoryFact` synchronously for the root, accepted descendants, and changed/disappeared/boundary events. Public `WalkFacts` is the scanner entry; private legacy wrappers may delegate to that engine only for existing tests and may never duplicate traversal in one scan.
- **Required exact tests:** scan `TestScannerStableSkippedKinds` and `TestScannerFiveProductionShapedTraversals`; Darwin `TestWalkerFactPortNestedFacts`, `TestWalkerFactPortOwnership`, `TestWalkerFactPortKindIdentityPairs`, and `TestWalkerFactPortNoDuplicateTraversal`. RED must be a genuine current nested/no-duplicate failure. Ownership evidence requires each accepted descriptor and retained root to close exactly once with no raw FD exposure.
- **Preserved safety:** D2A1 must retain accepted B1/B2 deterministic DFS depth, simultaneous descriptor capacity, child-before-parent LIFO closes, primary-error precedence with secondary close evidence, descriptor-bound replacement safety, and safe-sibling continuation. Regular-file jobs, content reads, goroutines, finalizer work, and WU5E1 are excluded.
- **Forecast and chain:** D2A1 is forecast at **180–300 additions, 20–80 deletions, 200–380 authored changed lines**, capped at 400. Sequential `stacked-to-main` remains selected with no pending decision. WU5E1 depends on independently accepted D2A1; failed D2 and D2A are insufficient.
- **Ledger and totals:** `tasks.md` now has **150 rows = 91 checked + 59 unchecked; 147 implementation-owned = 91 checked + 56 unchecked; 3 parent-owned = 0 checked + 3 unchecked**. Forecast totals are **6,690–9,320 additions, 228–768 deletions, 6,918–10,088 authored changed lines**, excluding generated `go.sum`.
- **Failed-D2A baseline hashes:** `descriptor.go` `1060f3eb77ffc226506ea966d112e54f1f91255ebe7d964d7339090dc20812db`; `descriptor_test.go` `3888ce12064a40753532f9513e9f03c0cc82fa1693db5096ccebbaa2c438789e`; `scanner.go` `e0b14a1de1e2864974b9079651cf275ee332a31db076d84d4912c3880025e40a`; `scanner_test.go` `3683f4af6094d7079e8f8c65a5fc37d43c70feb720ba84f6999891cab44b83fa`; `walker_darwin.go` `eb7e2267b41a22140286056e26aea38dc90e3e8c83529cf7b910fd0dfbdfcc47`; `walker_darwin_test.go` `e5b63d01f26e4ba051d5c86dc746f792b4bf73580e8be83a5540400fff10c5f5`; `walker_other.go` `8496ecf58d0a73aa14cc24e771c0a3921170144ee8e42d78902f8c5fd09b2dc1`. These are A1 rollback identities; no post-plan source hash changed.

### Receipt-integrity addendum

- The exact reconstructible D2 RED task line that was altered during the earlier non-append-only task rewrite is:

  `- [x] **WU5D2 RED:** Add fake-walker/finalizer tests for validated relative components and containment, synchronous directory kind/metadata/identity facts, enumeration-versus-open kind/identity mismatch, disappearance, symlink/special entry, different-device/non-local boundary, malformed/typed platform errors, deterministic warnings, safe siblings/later roots, and exact \`partial > boundary_limited > scanned\` mapping across all five final observations; run the focused command and capture RED caused only by missing directory-consumer integration. <!-- sdd-owner: implementation -->`

- This addendum supersedes only prior provenance claims that every historical D2 task line was preserved byte-for-byte or changed append-only. It does not rewrite that historical line, change its checked/failed/non-accepting state, alter D2's +257/-1=258 accounting, or grant acceptance to D2 or D2A.
- OpenSpec task/progress artifacts were updated, but the injected file tools expose no digest operation; no uncomputed post-plan artifact SHA-256 is invented. The exact persisted counts and unchanged failed-D2A source/test hashes above are the available integrity evidence.

---

## WU5D2A1 ledger recount and provenance clarification addendum

- **Superseded current ledger claims:** the earlier WU5D2A1 planning values **91 checked / 59 unchecked overall** and **91 checked / 56 unchecked implementation-owned** are incorrect and remain only as historical receipt text. The exact regex/manual recount is **150 total = 111 checked + 39 unchecked; 147 implementation-owned = 111 checked + 36 unchecked; 3 parent-owned = 0 checked + 3 unchecked**. Every checkbox row has exactly one terminal ownership marker; there are no unmarked rows. No checkbox state or ownership marker changed during this correction.
- **Append-only integrity criterion:** the planning receipt honestly documents the earlier historical D2 task-line rewrite using the exact reconstructible line already quoted in the Receipt-integrity addendum. WU5D2A1/A1 planning does not claim that prior rewrite was append-only. That acknowledged historical defect is provenance evidence and is not a blocker to implementing A1.
- **Current correction boundary:** this progress correction is append-only and supersedes only the current 91/59 overall and 91/56 implementation ledger claims. `tasks.md` current-ledger prose was corrected to the recounted values without changing checkbox rows; the review workload forecast is preserved unchanged. No code, test, command, settle, review, commit, PR, or lifecycle action occurred.
- **Artifact identity:** the injected file tools provide no digest operation, so no post-correction SHA-256 is fabricated.

---

## WU5D2A1 receipt — unified recursive Darwin fact traversal

**Receipt predecessor:** `/tmp/osdy-wu5d2a1-baseline-zGhBuB/apply-progress.md` was verified against its SHA256SUMS manifest before implementation; all nine entries matched. This is an append-only receipt. The active runtime continuation was inherited from the parent and deliberately is not recorded or settled here.

### Completed scope and task reconciliation

- Completed and visibly checked in `tasks.md`: WU5D2A1 RED, GREEN, TRIANGULATE, and REFACTOR. WU5E1 and all later implementation rows remain unchecked; the three parent-owned lifecycle rows remain unchanged and deferred.
- `trustedRootWalker.WalkFacts` now invokes one synchronous recursive descriptor engine. It emits copied `scan.DirectoryFact` values for nested accepted directories and for scoped changed, disappeared, boundary, symlink, and special-entry outcomes without exposing an FD or starting regular-file work.
- The acquired root retains its traversal limits; scanner ownership closes it once after `WalkFacts`. Child descriptors are closed once in child-before-parent order. The legacy private DFS tests remain on their existing private path, so no scanner invocation performs a second enumeration.
- The scanner test helper now accepts `*testing.T`, calls `t.Helper`, and uses `t.Fatalf` on definition construction failure. The unused production `transitional` helper was removed.

### TDD Cycle Evidence

| Task | Test layer | Safety net | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- | --- | --- |
| WU5D2A1 | Go unit tests in `internal/scan` and Darwin descriptor seams | `go test ./internal/scan -count=1` and `go test ./internal/platform/macos -count=1` passed before edits | The new exact scanner and Darwin test names were added first. The Darwin focused command failed on current behavior: nested facts omitted `one/two`, changed kind/identity lost final evidence, and the public port did not recurse. | Both focused commands passed after the public port delegated to the recursive descriptor engine. | Existing DFS/depth/capacity/lifecycle/replacement tests plus the new nested, ownership, kind/identity, no-duplicate, stable-skip, and five-root tests passed. | Formatted touched files, removed transitional production code, and replaced the panic test helper; tests remained green. |

### Verification

- Focused: `go test ./internal/scan -run '^(TestScannerProductionFactPort|TestScannerFactKindIdentityPairs|TestScannerStableSkippedKinds|TestScannerFiveProductionShapedTraversals)$' -count=1` -> PASS.
- Focused: `go test ./internal/platform/macos -run '^(TestWalkerFactPortTraversal|TestWalkerFactPortNestedFacts|TestWalkerFactPortOwnership|TestWalkerFactPortKindIdentityPairs|TestWalkerFactPortNoDuplicateTraversal)$' -count=1` -> PASS.
- B1/B2 DFS/lifecycle/replacement regression selection -> PASS: `TestWalkerDFSDepth`, `TestWalkerDescriptorCapacity`, `TestWalkerDirectoryLifecycle`, `TestWalkerSiblingContinuation`, `TestWalkerRealDirectoryReplacementBinding`, `TestWalkerRealChangedChildNoDescent`, `TestWalkerAcquisitionLifecycle`, and `TestWalkerAcquisitionClosePrecedenceMatrix`.
- Package/full/static/format: `go test ./internal/scan -count=1`, `go test ./internal/platform/macos -count=1`, `GOOS=linux go test -c -o /tmp/osdycleaner-wu5d2a1-linux.test ./internal/platform/macos && rm -f ...`, `go test ./... -count=1`, `go vet ./...`, and the touched-file `gofmt -d` check -> PASS. The Linux temporary binary was removed and confirmed absent.
- Fresh `lens_diagnostics mode=all` could not be invoked by this executor: no such CLI/tool was available and `gopls` is absent. `go vet ./...` and all Go compiler/test checks were clean; this is an LSP-evidence limitation, not a claim that a lens run passed.
- Runtime: N/A. The disposable synchronous scanner/Darwin descriptor seams are the bounded harness; no CLI, regular-file pipeline, raw FD escape, content read, goroutine, or mutation was added.

### Accounting, identity, and rollback

- Baseline: `/tmp/osdy-wu5d2a1-baseline-zGhBuB`; all manifest entries verified before edits.
- Authored delta across the seven permitted source/test baseline entries: **+221/-53 = 274** changed lines, within the independent 400-line limit. `descriptor.go`, `descriptor_test.go`, and `walker_other.go` are byte-identical to the baseline.
- Final SHA-256: `internal/scan/descriptor.go` `1060f3eb77ffc226506ea966d112e54f1f91255ebe7d964d7339090dc20812db`; `internal/scan/descriptor_test.go` `3888ce12064a40753532f9513e9f03c0cc82fa1693db5096ccebbaa2c438789e`; `internal/scan/scanner.go` `1454b9bde231336876ff4d4b5873d90f2f1e184f257a9cbdd2a40a783581ec8b`; `internal/scan/scanner_test.go` `438cf2db877219d76268367711450c3ac3428b1b44798cd0930a2b5fc51ee25d`; `internal/platform/macos/walker_darwin.go` `b1016fff12168c21bc897acb1376013fa896570d7a121593697176676e7673ed`; `internal/platform/macos/walker_darwin_test.go` `0a31b66dbd92dc79894e4ac5f00b8f6917ebbed40df0f993a79fc20a72d693ef`; `internal/platform/macos/walker_other.go` `8496ecf58d0a73aa14cc24e771c0a3921170144ee8e42d78902f8c5fd09b2dc1`.
- Rollback: restore the seven named files from `/tmp/osdy-wu5d2a1-baseline-zGhBuB/`; this removes only the unified recursive fact-port correction and leaves WU5E1 blocked.

### Status and remaining work

- Consumed authoritative OpenSpec status: `applyState: ready`, repository-local workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, workspace edit root allowed, strict TDD enabled, and no blocked reasons. The standing `stacked-to-main` authorization covered this independent slice.
- Delivery boundary: WU5D2A1 only; no commit, review, receipt approval, or settlement was performed. The requested evidence must eventually remediate `sha256:a94ad5095f29c63ac3e9c5953f2a24ac589eb74c2ba1958f8f853eab9d741dd8` through parent-owned runtime/lifecycle handling.
- Exact next unchecked implementation row: `- [ ] **WU5E1 RED:** Add synchronized scanner and Darwin tests for fixed worker/job/result counts, bounded capacities and backpressure, validated relative identity-chain jobs, trusted-root private-anchor reopen, rolling ancestor descriptors, exact no-follow flags, ancestor/final identity and device/local checks, enumeration-versus-open kind/identity/disappearance, regular-only logical/allocation facts, no content reads, and exact close counts; run both focused commands and capture RED caused only by the missing file pipeline/reopen APIs. <!-- sdd-owner: implementation -->`

---

## WU5D2A1 verifier-Fail correction addendum

**Continuation:** corrected the candidate rejected by `subtask_gentle-ai-verify_1788142842979_0981c005` within the inherited active attempt. This is an append-only addendum; no task row, historical receipt, design/spec/proposal/docs artifact, review action, commit, or settlement was changed.

### Corrected engine and event ordering

- Removed the separate recursive `walkFrame`. `walkDirectory` legacy tests and public `WalkFacts` now delegate to `walkDirectoryFrame`, the sole recursive descriptor engine.
- In fact mode it synchronously emits stable symlink/special facts and directory changed/disappeared/boundary/limit classes while continuing safe siblings. In nil-visitor legacy mode it ignores stable non-directory entries and accumulates open/gate/fact classes while preserving DFS depth, FD capacity, LIFO child-before-parent close order, and primary-error precedence.
- Added `TestScannerProductionTraversalCompletesBeforeRootClose`. For every production-shaped root it proves exactly one `WalkFacts` begin/end pair followed by exactly one trusted-root close; a close count alone is no longer the evidence.

### TDD Cycle Evidence

| Task | Layer | Safety net | RED replay | GREEN / triangulation | REFACTOR |
| --- | --- | --- | --- | --- | --- |
| WU5D2A1 correction | Go unit / descriptor seam | `go test ./internal/scan -count=1` and `go test ./internal/platform/macos -count=1` PASS before the refactor | **Replay, not original timestamp:** a disposable copy restored only pre-A1 `walker_darwin.go` from `/tmp/osdy-wu5d2a1-baseline-zGhBuB`; the exact A1 focus exited 1 in Darwin. It omitted `one/two`, lost changed kind/identity evidence, and performed one rather than root-plus-child traversal. The temp copy was removed. | Exact focused scan/Darwin suite PASS; the added event-order assertion and existing nested/kind/ownership/no-duplicate plus B1/B2 regressions PASS. | Formatted touched files; full tests, vet, and format checks PASS. |

### Verification and diagnostics

- Focused scan: `go test ./internal/scan -run '^(TestScannerProductionFactPort|TestScannerFactKindIdentityPairs|TestScannerStableSkippedKinds|TestScannerFiveProductionShapedTraversals|TestScannerProductionTraversalCompletesBeforeRootClose)$' -count=1` -> PASS.
- Focused Darwin A1 selection and B1/B2 DFS/lifecycle/replacement regressions -> PASS.
- `go test ./internal/scan -count=1`, `go test ./internal/platform/macos -count=1`, `GOOS=linux go test -c -o /tmp/osdycleaner-wu5d2a1-linux.test ./internal/platform/macos && rm -f /tmp/osdycleaner-wu5d2a1-linux.test`, `go test ./... -count=1`, `go vet ./...`, and touched-file `gofmt -d` -> PASS. The Linux temporary binary was removed.
- `lens_diagnostics` and `gopls` binaries remain unavailable to this executor. Parent-provided LSP diagnostics reported zero errors; the only prior `descriptor.go` `NewRootFacts` go-bare-error advisory is the already-dispositioned false positive. There are zero known blocking diagnostics; binary absence is an evidence limitation, not a blocker.

### Accounting correction and rollback

- Canonical baseline: `/tmp/osdy-wu5d2a1-baseline-zGhBuB`.
- Final `git diff --no-index --numstat` is +270/-97=367 across the seven baseline entries; independent line-level `difflib.ndiff` is +272/-99=371. Both are within the cumulative 400-line cap. The two tools differ on replacement hunk accounting, so neither value is substituted for the other.
- This addendum **supersedes the prior +221/-53=274 receipt claim**; no history was rewritten. `descriptor.go`, `descriptor_test.go`, and `walker_other.go` remain byte-identical to baseline.
- Final hashes: `scanner.go` `1454b9bde231336876ff4d4b5873d90f2f1e184f257a9cbdd2a40a783581ec8b`; `scanner_test.go` `227c75572e75c5d00866001a137db48510400fc2c544c3ac06062b91eb1bd734`; `walker_darwin.go` `6ffc7d29e6b4b637489f09e36f0f9e87a8356d472679bd2aa58f490f1486b91f`; `walker_darwin_test.go` `0a31b66dbd92dc79894e4ac5f00b8f6917ebbed40df0f993a79fc20a72d693ef`.
- Rollback: restore the seven named files from the canonical baseline; this removes A1 fact traversal work and leaves WU5E1 blocked.

### Task and status reconciliation

- Re-read `tasks.md`: all four WU5D2A1 rows remain visibly `[x]`; every WU5E1 row remains visibly `[ ]`; no parent-owned row was modified.
- Consumed authoritative OpenSpec status: `applyState: ready`, repo-local workspace, workspace edit root allowed, strict TDD active, and sequential `stacked-to-main` delivery authorization. No unsafe action-context warning occurred.
- Workload/PR boundary: WU5D2A1 correction only, maximum measured 371 under difflib accounting. Runtime is N/A; no raw FD, content read, goroutine, or mutation was added. Parent must independently verify and own all lifecycle actions; this executor did not settle the active attempt.

---

## WU5D2A1 audited-index incident and final acceptance addendum

### Independent acceptance evidence

- Functional/source/accounting verifier `subtask_gentle-ai-verify_1788143365237_f68de7e6` returned PASS except that its executor could not supply lens/LSP evidence. Parent-native `lens_diagnostics` with `mode=all` then reported **0 blockers**; one unchanged warning was explicitly dispositioned as a false positive. Parent-native LSP diagnostics reported **0 errors**. Strict adjudicator `subtask_gentle-ai-verify_1788143523521_bbd9b818` returned PASS.
- WU5D2A1 is accepted from canonical baseline `/tmp/osdy-wu5d2a1-baseline-zGhBuB`. The baseline manifest identifier was supplied as `20787...c666`; the retained `SHA256SUMS` lists all seven source/test baselines plus `tasks.md` and `apply-progress.md`. Canonical independently reproduced accounting is **+270/-97=367**, within 400.
- D2 remains checked failed/non-accepting history at +257/-1=258. D2A remains checked failed/non-accepting history at +196/-43=239. Neither satisfies a dependency. WU5E1 now depends on accepted D2A1 and is the next implementation unit.

### Audited index incident and controller authority

- Selective staging was explicitly user-authorized only to establish bounded candidate identity. No commit or push occurred, and this addendum performs no code, test, checkbox, commit, review, or lifecycle action.
- Ordinal 68 settlement successfully remediated evidence revision `sha256:a94ad5095f29c63ac3e9c5953f2a24ac589eb74c2ba1958f8f853eab9d741dd8`. Because staging began from an empty index, the controller recorded full-file `changed_lines: 6709` and historical `changed_line_budget_exceeded: true`; that value is not D2A1 baseline-delta accounting.
- The user authorized index rearm. Ordinal 69 bounded adjudication reached controller state `complete` with `decision_required: false`. Its staged baseline tree was `1bdab7136fbfd481c8c9008d9c418dcf2496fbb8`, while the current tree was `7737f140f615da329a79bff034d494f88c200f22`. However, the controller candidate snapshot used the current working tree at both begin and finish, so ordinal 69 recorded `changed_lines: 0` rather than mechanically measuring 367.
- This procedural discrepancy is acknowledged, not hidden: the controller ledger is authoritative as complete with no pending decision, but it did **not** mechanically measure 367. The canonical 367 value is independently reproduced immutable hash-baseline evidence from `/tmp/osdy-wu5d2a1-baseline-zGhBuB`.
- Ordinal 68 finish evidence revision is `sha256:95e9f84e8e2ebc063f18b8d595b3741ebf8accd199b57c4eedc690b7cca30629`. Ordinal 69 finish uses the same evidence revision, candidate identity `sha256:060c01b43cba990db98ee83b46af954c5896964dc373e2c8a98e4f6dc11ed1c0`, outcome `passed`, and controller HEAD `sha256:84768834be24be4930ad80b4f6b293653aa3cf35ef3e61fb855c5e2d4aac677b`.

### Final identity, index, and ledger

- Final source/test SHA-256 values are: `internal/scan/descriptor.go` `1060f3eb77ffc226506ea966d112e54f1f91255ebe7d964d7339090dc20812db`; `internal/scan/descriptor_test.go` `3888ce12064a40753532f9513e9f03c0cc82fa1693db5096ccebbaa2c438789e`; `internal/scan/scanner.go` `1454b9bde231336876ff4d4b5873d90f2f1e184f257a9cbdd2a40a783581ec8b`; `internal/scan/scanner_test.go` `227c75572e75c5d00866001a137db48510400fc2c544c3ac06062b91eb1bd734`; `internal/platform/macos/walker_darwin.go` `6ffc7d29e6b4b637489f09e36f0f9e87a8356d472679bd2aa58f490f1486b91f`; `internal/platform/macos/walker_darwin_test.go` `0a31b66dbd92dc79894e4ac5f00b8f6917ebbed40df0f993a79fc20a72d693ef`; `internal/platform/macos/walker_other.go` `8496ecf58d0a73aa14cc24e771c0a3921170144ee8e42d78902f8c5fd09b2dc1`.
- The index currently contains only those seven D2A1 source/test files at current content. OpenSpec artifacts remain unstaged. No commit or push exists, and working files were unchanged by index rearm/adjudication.
- No checkbox changed during acceptance finalization. The earlier planning recount was rechecked against the persisted work-unit rows after D2A1 implementation; a correction follows without altering checkbox state or ownership.

### Acceptance ledger recount correction

- The four WU5D2A1 rows are already checked in the persisted ledger, so the prior 111/39 overall and 111/36 implementation-owned figures omitted those four checked rows. The exact current recount is **150 total = 115 checked + 35 unchecked; 147 implementation-owned = 115 checked + 32 unchecked; 3 parent-owned = 0 checked + 3 unchecked**.
- Every checkbox retains exactly one terminal ownership marker. This recount changes narrative only; no checkbox, ownership marker, code, test, commit, review, or lifecycle state changed.

---

## WU5E1 blocked-for-split forecast receipt

- **Status consumed:** authoritative native OpenSpec status for `read-only-scan-foundation`: `applyState: ready`, `nextRecommended: apply`, artifact store `openspec`, strict TDD enabled, and repo-local action context with `/Users/osdy/Documents/GitHub/OsdyCleaner` as the allowed edit root. The parent-bound continuation token was joined and returned `proceed`; this executor does not settle it.
- **Baseline verified before any product or test edit:** `/tmp/osdy-wu5e1-baseline-tRW0me/SHA256SUMS` passed `shasum -a 256 -c` for all seven entries. Its SHA-256 is `7db22a09131ec600ccefcf78d06a30be8384bfe2a7114925c21834d275f1de26`. The five allowed WU5E1 source/test files are byte-identical to that canonical baseline. The seven D2A1 files remain selectively staged under explicit user authorization; this receipt did not alter the index, commit, or push.

### Mandatory detailed forecast and split decision

| Work item | Plausible additions | Plausible deletions | Rationale |
| --- | ---: | ---: | --- |
| Scanner RED/TRIANGULATE synchronization, fixed-worker, bounded queue, and result-pressure evidence | 145–205 | 0–10 | Four required scanner tests need controlled blocking, exact worker/capacity assertions, and fake relative jobs. |
| Darwin RED/TRIANGULATE descriptor fixtures | 180–250 | 0–15 | Four required Darwin tests need private-anchor, identity-chain, replacement, regular-only, and lifecycle seams. |
| Scan contract and bounded worker production | 100–145 | 10–25 | Validated job/result facts, fixed capacities, backpressure, and worker ownership require new public contract behavior. |
| Darwin relative reopen production | 135–190 | 10–30 | Component-wise `openat`, ancestor/final validation, boundary checks, rolling close, and zero-fact classifications are absent. |
| Focused refactor/compatibility adjustment | 20–45 | 0–10 | Existing directory-only callback behavior must remain coherent while regular work is introduced. |
| **Total** | **580–835** | **20–90** | **600–925 authored changed lines; exceeds the independent 400-line cap.** |

The pre-existing 270–360 estimate is not credible after enumerating the required synchronized scanner and Darwin fixture coverage plus the two missing production boundaries. Because more than 400 lines are plausible, no RED test, production edit, task checkbox update, or other product/test modification was made.

### Proposed standing-authorization split

1. **WU5E1A — bounded scan contract and worker queue:** `descriptor.go`, `scanner.go`, and `scanner_test.go`; validated regular jobs/results, fixed worker/channel bounds, backpressure, and scanner-only regular/change mapping. Forecast 260–360 authored lines.
2. **WU5E1B — Darwin trusted-root relative reopen:** `descriptor.go` only if the shared contract needs a minimal extension, `walker_darwin.go`, and `walker_darwin_test.go`; private anchor, rolling `openat` reopen, identity/boundary/regular checks, zero-fact failures, and exact closes. Forecast 280–390 authored lines.

This preserves strict RED → GREEN → TRIANGULATE → REFACTOR in each independently bounded slice and leaves E2 cancellation, deterministic aggregation, and hard-link work entirely deferred. Parent planning must replace the four WU5E1 rows with the two bounded sequences before apply resumes.

### Reconciliation and remaining work

- Persisted `tasks.md` was re-read: all four WU5E1 implementation rows remain visibly `- [ ]`; no parent-owned row changed.
- No TDD cycle was started because the mandatory workload gate stopped before test edits; no test command was run.
- No deviations from specs/design were implemented. Runtime evidence remains N/A, and no receipt was settled, approved, reviewed, committed, pushed, or PR-created.
- Workload / PR boundary: blocked before WU5E1; proposed sequential `stacked-to-main` WU5E1A then WU5E1B boundaries.

---

## Planning reset — WU5E1A and WU5E1B authorized split

This append records planning only. No product, test, specification, design, proposal, documentation, generated file, commit, review, settlement, push, or PR operation occurred.

### Reset and preserved baseline

- The original WU5E1 detailed forecast remains **600–925 authored changed lines**. Its four unchecked rows remain verbatim historical nonexecuting context and satisfy no dependency.
- Canonical untouched baseline remains `/tmp/osdy-wu5e1-baseline-tRW0me`; its `SHA256SUMS` SHA-256 remains `7db22a09131ec600ccefcf78d06a30be8384bfe2a7114925c21834d275f1de26`. The prior blocked receipt verified all seven entries and the five E1 source/test files byte-identical before any attempted work; this planning reset made no product/test edit.
- The seven accepted D2A1 files remain selectively staged exactly as recorded before planning. Only `tasks.md` and this append-only receipt were edited in the working tree; the index was not modified. No commit or push occurred.

### Sequential bounded plan

1. **WU5E1A — scanner-owned bounded file pipeline/contract, 260–360 lines.** Depends on independently accepted D2A1. Scan owns opaque validated `FileJob`/`FileResult` values with relative components plus ancestor/final identities and no absolute path/raw FD. Fixed workers and bounded job/result channels prove backpressure, local cancellation/close lifecycle, drain/join/no-leak/race behavior, regular-only facts, changed-entry zero contribution, and deterministic path aggregation through synchronized fakes only. Exact scanner test names are `TestScannerBoundedFileJobs`, `TestScannerFileBackpressure`, `TestScannerRegularOnlyFacts`, and `TestScannerFileEntryChanges`; Darwin and real filesystems are forbidden in this slice.
2. **WU5E1B — Darwin private-root relative reopen, 280–390 lines.** Depends on independently accepted E1A. Darwin implements the E1A port from a private retained root anchor with rolling `openat` no-follow/cloexec/nonblock acquisition, ancestor/final `fstat`/`fstatfs` identity/device/mount/local checks, regular-only logical/allocation facts, no content reads, and exact operation-primary/cleanup-secondary close precedence. Exact platform test names are `TestWalkerRelativeFileReopen`, `TestWalkerFileIdentityChain`, `TestWalkerRegularFileFacts`, and `TestWalkerFileReopenLifecycle`; real `t.TempDir()` rename/symlink/disappearance/nonregular/boundary fixtures are required and scanner wiring is minimal.
3. **WU5E2 remains unchanged in scope.** It depends on independently accepted E1B and exclusively owns scan-wide cancellation state integration, completion-order determinism, hard-link attribution, and finalizer integration; it must not duplicate E1A queue/contract or E1B Darwin reopen fixtures.

Each new unit has four unchecked strict RED → GREEN → TRIANGULATE → REFACTOR rows, an independent ≤400 authored-line guard, a focused command, verification boundary, and rollback boundary. Apply resumes at WU5E1A RED only.

### Forecast, mapping, and ledger reconciliation

- Forecast totals are now **7,188–10,478 authored changed lines**, with forecast table components **6,970–9,740 additions** and **218–738 deletions**, excluding generated `go.sum`. The historical E1 row is informational and is not double-counted against E1A/E1B.
- Normative mappings now assign bounded scan ownership to E1A, Darwin identity/boundary/reopen evidence to E1B, and hard links/cancellation integration/determinism to E2.
- Exact recount: **158 checkbox rows = 115 checked + 43 unchecked**. Ownership recount: **155 implementation-owned = 115 checked + 40 unchecked; 3 parent-owned = 0 checked + 3 unchecked**. Every checkbox retains exactly one terminal ownership marker.
- The only checkbox additions are eight unchecked implementation-owned E1A/E1B rows. The original E1 four rows remain unchecked; no existing checkbox state or ownership marker changed.
- Post-edit `tasks.md` and apply-progress SHA-256 values are unavailable through the injected file-only tools and are not fabricated. The known canonical product-baseline manifest hash above remains the rollback/provenance identity.

## WU5E1A — Scanner-owned bounded file pipeline (completed)

- **Scope / delivery boundary:** WU5E1A only, stacked-to-main slice; no Darwin, real-filesystem, hard-link, final snapshot cancellation/status, E2, documentation, index, commit, review, or runtime-attempt settlement work. The parent-provided live attempt token was continued through `sdd-attempt acquire`; it was deliberately not settled as directed.
- **Status consumed:** native `sdd-status` reported `artifactStore: openspec`, `applyState: ready`, `nextRecommended: apply`, and repo-local action context rooted at `/Users/osdy/Documents/GitHub/OsdyCleaner` with that root allowed. No action-context warnings. The WU5E1A delivery decision is the pre-authorized `stacked-to-main` split.
- **Baseline verification:** `/tmp/osdy-wu5e1a-baseline-fYVyLa/SHA256SUMS` hashes to `b9d679f4b94ad5dd60cbf07bedc9efec5cd04031b453ff02b416d39b3a808604`; its four listed source/test hashes matched the pre-edit scan files. The existing seven staged D2A1 files were not staged, unstaged, or otherwise changed in the index.
- **Completed persisted rows:** WU5E1A RED, GREEN, TRIANGULATE, and REFACTOR are visibly `- [x]` in `tasks.md`.
- **Files changed:** `internal/scan/descriptor.go`, `internal/scan/descriptor_test.go`, `internal/scan/scanner.go`, `internal/scan/scanner_test.go`, and this cumulative receipt/task artifact only.
- **Implementation:** opaque copied `FileJob` identity-chain values and metadata-only `FileResult` values; a fake-port-only fixed worker pool with positive bounded job/result queues, serialized close versus submit ownership, close-once job/result channels, worker join, result draining, unchanged-regular-only retention, and normalized relative-path ordering. Zero capacities are rejected. No descriptors, handles, absolute descendant paths, Darwin tags/syscalls, OS/filesystem calls, or content reads were introduced.

### TDD Cycle Evidence

| Task | Test file | Layer | Safety net | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- | --- | --- | --- |
| WU5E1A | `internal/scan/descriptor_test.go`, `internal/scan/scanner_test.go` | Unit | `go test ./internal/scan -count=1` PASS before edits | Focused command failed from undefined E1A APIs | Focused tests PASS | Constructor invalid/copy chains; zero capacity rejection; full job and result backpressure; changed/error/nonregular zero facts; blocked cancellation; repeated close; ordered results | `gofmt`, race/package/full/vet PASS |

### Verification and accounting

- RED: `go test ./internal/scan -run '^(TestScannerBoundedFileJobs|TestScannerFileBackpressure|TestScannerRegularOnlyFacts|TestScannerFileEntryChanges)$' -count=1` failed solely because `FileJob`, `FileResult`, and `FilePipeline` APIs were absent.
- GREEN/TRIANGULATE: focused suite and `go test -race ./internal/scan -run '^(TestScannerBoundedFileJobs|TestScannerFileBackpressure|TestScannerRegularOnlyFacts|TestScannerFileEntryChanges)$' -count=1` PASS.
- REFACTOR: `go test -race ./internal/scan -count=1`, `go test ./... -count=1`, `go vet ./...`, and clean `gofmt -d` PASS.
- Source-safety search for Darwin/syscall/OS/absolute-descendant APIs in the two production scan files returned no matches.
- Accounting from the verified E1A baseline: `descriptor.go` +55, `descriptor_test.go` +42, `scanner.go` +88, `scanner_test.go` +123; **+308/-0 = 308 authored changed lines**, within the 400-line cap.
- Final hashes: `descriptor.go` `aa191d41ad70864d3bc7ca954fbe8f6c31f7970693dfe7573460d17a5616d620`; `descriptor_test.go` `f2ebb0847e76c497986efc234d642596b75f870875cfa1d5bc306a126cc15315`; `scanner.go` `7f94088fed750386083e93ae3d189602d9748c869b9fdebb0e8f5cdc123c319e`; `scanner_test.go` `b6ba701abc30a86f56f3e58af5758b4afe8ad15a06aae89862f37a69e650cafb`.
- Runtime evidence: N/A; all work used channel-synchronized fakes and no real filesystem.
- Rollback: restore these four files from `/tmp/osdy-wu5e1a-baseline-fYVyLa`; E1B and E2 remain absent.

### Remaining work

- WU5E1B remains unchecked and requires independent parent lifecycle/verification routing before its Darwin-only slice. WU5E2 remains unchecked and deferred; its hard-link and scan-wide cancellation/status integration was not started.

Exact next unchecked rows:

- [ ] **WU5E1B RED:** Darwin private-root relative reopen remains deferred.
- [ ] **WU5E1B GREEN:** Darwin private-root relative reopen remains deferred.
- [ ] **WU5E1B TRIANGULATE:** Darwin private-root relative reopen remains deferred.
- [ ] **WU5E1B REFACTOR:** Darwin private-root relative reopen remains deferred.

---

## WU5E1A lifecycle correction after independent FAIL

**Continuation and scope:** This append corrects the same live WU5E1A attempt after independent FAIL `subtask_gentle-ai-verify_1788145593441_6287bd38`. The inherited active token was authenticated with `sdd-attempt acquire --token` and returned `proceed`; it is deliberately **not settled** here. No E1B/index/commit/review/receipt approval/settlement action occurred. The only edited product/test files remain the four E1A scan files, and WU5E1B remains untouched.

### Corrected lifecycle

- Removed the `sendMu` design that held a mutex across a potentially blocking job send.
- `Submit` now registers in an active-submitter `WaitGroup` while holding the same mutex used to reject closed pipelines, preventing `Add`-after-`Wait`; it then selects caller cancellation, pipeline cancellation, or bounded job acceptance without holding that mutex.
- `CloseJobs` atomically marks closed and cancels the internal pipeline context before returning. Its sole close coordinator waits registered submitters, closes jobs, waits workers, and then closes results; `Drain` supplies the result consumption required to release result-pressure-blocked workers.
- Jobs retain their submitter context in a private envelope. Workers pass `Inspect` a context canceled by either caller cancellation after acceptance or pipeline close; no worker calls `Inspect(context.Background(), ...)`. There is no explicit goroutine per job: workers are fixed and the sole non-worker coordinator owns closure.

### TDD Cycle Evidence

| Task | Layer | Safety net | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- | --- | --- |
| WU5E1A lifecycle correction | Go unit/channel-synchronized fake port | Existing focused E1A suite was the pre-change safety net | Revised focused tests failed against the prior code: `CloseJobs blocked behind Submit`; caller and close cancellation never reached `Inspect` because it received `context.Background()` | Focused suite passed after registered submitters, pipeline cancellation, envelope contexts, and close coordinator were added | Deterministic pressure test fills active worker plus job/result pressure, proves a third submit releases only after close, drains both accepted inspections, and separately proves caller cancellation reaches accepted `Inspect`; 25 repeated Submit/Close races run with no sleeps | Removed serialized sender mutex; formatted and repeated focused/race/full/static/Linux checks passed |

### Verification and diagnostics

- Focused: `go test ./internal/scan -run '^(TestScannerBoundedFileJobs|TestScannerFileBackpressure|TestScannerRegularOnlyFacts|TestScannerFileEntryChanges)$' -count=50` -> PASS.
- Focused race stress: same selection with `go test -race ... -count=50` -> PASS.
- Package race: `go test -race ./internal/scan -count=1` -> PASS.
- Full/static: `go test ./... -count=1` and `go vet ./...` -> PASS.
- Linux compile: `GOOS=linux go test -c -o /tmp/osdycleaner-wu5e1a-linux.test ./internal/scan`, followed by removal and absence check -> PASS.
- Formatting: `gofmt -d` on all four E1A files was empty.
- Source scan confirms `sendMu` and `Inspect(context.Background(), ...)` are absent. `gopls` and `lens_diagnostics` are not installed in this executor environment, so no LSP/lens result is claimed; compiler, race detector, and vet have no diagnostics.
- Runtime evidence remains N/A: all tests use channel-synchronized fakes only, with `time.After` used solely as a fail-safe test deadline and no sleep-based synchronization.

### Accounting, identity, and rollback

- Canonical baseline `/tmp/osdy-wu5e1a-baseline-fYVyLa` was revalidated: all six `SHA256SUMS` entries passed.
- Baseline delta: `descriptor.go` +55/-0, `descriptor_test.go` +42/-0, `scanner.go` +109/-0, `scanner_test.go` +144/-0 = **+350/-0 = 350 authored changed lines**, within the 400-line cap. This supersedes the prior weak-test accounting claim of +308/-0.
- Final SHA-256: `internal/scan/descriptor.go` `aa191d41ad70864d3bc7ca954fbe8f6c31f7970693dfe7573460d17a5616d620`; `internal/scan/descriptor_test.go` `f2ebb0847e76c497986efc234d642596b75f870875cfa1d5bc306a126cc15315`; `internal/scan/scanner.go` `8321de810fe8d1438ba900fa4ed4ab7f1a55cf9fd2225fb7dbac99e3160563f2`; `internal/scan/scanner_test.go` `cad79300abe8480ed25bd5233c8455978a6d38c9ba68c11e79ddc7379bde1d84`.
- Rollback: restore the four scan files from that baseline. E1B/E2 remain absent.

### Status, task reconciliation, and delivery boundary

- Consumed authoritative native status: `artifactStore: openspec`, `applyState: ready`, `nextRecommended: apply`, repo-local workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, allowed edit root that workspace, strict TDD active, and no blocked reasons. The standing stacked-to-main E1A slice resolves the high-risk workload gate.
- Persisted `tasks.md` was re-read: all four WU5E1A rows remain visibly `[x]` as historical candidate implementation evidence pending independent acceptance; all four WU5E1B rows remain visibly `[ ]` and are deferred. No parent-owned lifecycle row was changed.
- Workload / PR boundary: WU5E1A correction only, 350 authored lines; no commit or PR. Independent verification and all lifecycle actions remain parent-owned.

---

## WU5E1A nil-context correction after provider retry

- **Continuation and scope:** Retried the same live WU5E1A correction after the prior provider error. The parent-bound attempt remains private and was not settled. No E1B, index, commit, review, receipt approval, or settlement action occurred.
- **RED:** Added a table-driven nil-context case inside `TestScannerFileBackpressure`. The focused command failed with a nil-pointer panic at `FilePipeline.Submit` on `ctx.Err()`; no work was accepted before the panic.
- **GREEN:** `Submit` now returns the existing precise invalid-input class `ErrInvalidPipeline` when `ctx == nil`, before `ctx.Err()`, job validation, or closed-pipeline checks. This makes nil-context validation deterministic for both valid and invalid jobs and after `CloseJobs`.
- **TRIANGULATE:** The table covers nil context plus a valid job and nil context plus `FileJob{}`; it asserts `ErrInvalidPipeline` and zero accepted inspections. The same test then submits a valid job with `context.Background()`, observes and drains its accepted result, closes cleanly, and confirms a nil context after close still returns `ErrInvalidPipeline`.
- **REFACTOR / verification:** `gofmt -w internal/scan/scanner.go internal/scan/scanner_test.go`; focused `go test ./internal/scan -run '^(TestScannerBoundedFileJobs|TestScannerFileBackpressure|TestScannerRegularOnlyFacts|TestScannerFileEntryChanges)$' -count=100`; focused race with the same selector and `-count=100`; `go test -race ./internal/scan -count=1`; `go test ./... -count=1`; `go vet ./...`; `GOOS=linux go test -c -o /tmp/osdycleaner-wu5e1a-linux.test ./internal/scan` followed by removal/absence check; and clean `gofmt -d` all PASS. The source scan also found no `sendMu` or `Inspect(context.Background(), ...)` in the four E1A files.

### TDD Cycle Evidence

| Task | Test file / layer | Safety net | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- | --- | --- |
| WU5E1A nil-context correction | `internal/scan/scanner_test.go` / unit | Existing exact focused suite PASS | Nil-context valid-job row panicked at `ctx.Err()` | Exact focused suite PASS after the nil guard | Valid job, invalid job, usable-after-nil, and closed-after-nil cases PASS | Focused/race/package/full/vet/Linux/format checks PASS |

### Accounting, hashes, and task reconciliation

- Baseline `/tmp/osdy-wu5e1a-baseline-fYVyLa/SHA256SUMS` remains verified at `b9d679f4b94ad5dd60cbf07bedc9efec5cd04031b453ff02b416d39b3a808604`.
- Corrected cumulative baseline delta: `descriptor.go` +55/-0, `descriptor_test.go` +42/-0, `scanner.go` +112/-0, `scanner_test.go` +191/-0 = **+400/-0 = 400 authored changed lines**. This replaces the prior +350/-0 claim and remains within the hard cap.
- Final SHA-256: `descriptor.go` `aa191d41ad70864d3bc7ca954fbe8f6c31f7970693dfe7573460d17a5616d620`; `descriptor_test.go` `f2ebb0847e76c497986efc234d642596b75f870875cfa1d5bc306a126cc15315`; `scanner.go` `672f14c0e410c9aef07b5eacadec53f6b8524e484a8da266c43c87f40342deb6`; `scanner_test.go` `6abfa3912d62446a74f207f227832379c79c6df9c92134cdf152bab55acd6417`.
- Rollback remains restoration of the four scan files from the baseline above; E1B/E2 remain absent.
- `tasks.md` was re-read: WU5E1A RED/GREEN/TRIANGULATE/REFACTOR remain visibly `[x]` as candidate implementation evidence pending independent acceptance; WU5E1B RED/GREEN/TRIANGULATE/REFACTOR remain visibly `[ ]`. No checkbox changed because this is a correction within already checked E1A rows. Parent-owned lifecycle rows remain byte-for-byte deferred.
- **Status consumed/produced:** fallback structured status (native status was not supplied): OpenSpec authoritative change `read-only-scan-foundation`; proposal/spec/design/tasks/apply-progress readable; `applyState: ready`; `actionContext.mode: repo-local`; workspace and allowed edit root `/Users/osdy/Documents/GitHub/OsdyCleaner`; strict TDD active; no action-context warnings. The standing `stacked-to-main` E1A slice resolves the workload gate.
- **Workload / PR boundary:** WU5E1A correction only, exactly 400 authored lines. Independent verification and all lifecycle actions remain parent-owned; return `parent-lifecycle`. No settle.

---

## WU5E1A independent acceptance finalization

- Independent verifier `subtask_gentle-ai-verify_1788146646060_00f9245a` returned **PASS**. Parent LSP diagnostics reported **0 errors**, parent lens diagnostics were clean, and the externally completed settlement state is `complete`.
- Canonical accounting from `/tmp/osdy-wu5e1a-baseline-fYVyLa` is exactly **+400/-0=400**, at the hard cap: `descriptor.go` +55/-0, `descriptor_test.go` +42/-0, `scanner.go` +112/-0, and `scanner_test.go` +191/-0.
- Accepted final SHA-256 values are `internal/scan/descriptor.go` `aa191d41ad70864d3bc7ca954fbe8f6c31f7970693dfe7573460d17a5616d620`; `internal/scan/descriptor_test.go` `f2ebb0847e76c497986efc234d642596b75f870875cfa1d5bc306a126cc15315`; `internal/scan/scanner.go` `672f14c0e410c9aef07b5eacadec53f6b8524e484a8da266c43c87f40342deb6`; and `internal/scan/scanner_test.go` `6abfa3912d62446a74f207f227832379c79c6df9c92134cdf152bab55acd6417`.
- Ledger recount from all persisted checkbox rows is **158 total = 119 checked + 39 unchecked; 155 implementation-owned = 119 checked + 36 unchecked; 3 parent-owned = 0 checked + 3 unchecked**. This acceptance finalization changed no checkbox or ownership marker.
- The original WU5E1 remains untouched historical oversized/nonexecuting context. WU5E1A is accepted; WU5E1B is next and depends on accepted E1A. WU5E2 remains blocked until E1B is independently accepted.
- The seven D2A1 files remain staged, while the four E1A working changes remain unstaged. No index change, commit, push, ordinary review, or lifecycle action was performed by this finalization.

---

## WU5E1B receipt — Darwin private-root relative file reopen

- Completed and persisted: WU5E1B RED, GREEN, TRIANGULATE, and REFACTOR are visibly `[x]` in `tasks.md`; WU5E2 remains unchecked. Parent-owned lifecycle rows were deferred unchanged.
- Baseline `/tmp/osdy-wu5e1b-baseline-Y55fRv` was verified against `SHA256SUMS` before edits. Relative accounting is **292 additions + 1 deletion = 293 authored lines**, below the 400-line cap.
- `acquiredRoot` privately retains the root descriptor and implements opaque `FilePort.Inspect`. It reopens only relative job components with rolling `openat` descriptors, validates each ancestor/final `fstat` and `fstatfs` identity/device/locality chain, opens the final regular file with `O_RDONLY|O_NOFOLLOW|O_CLOEXEC|O_NONBLOCK`, and never reads content or exposes an FD/path.
- Changed, symlinked, missing, nonregular, nonlocal, invalid, and operation/close-error cases return zero-byte unaccepted results. Primary operation errors remain primary while close failures are retained as secondary joined evidence; superseded/final descriptors close once.

### TDD Cycle Evidence

| Task | Layer | Safety net | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- | --- | --- |
| WU5E1B | Darwin unit / real `t.TempDir()` plus narrow descriptor seams | `go test ./internal/platform/macos -count=1` PASS | Exact focused command failed only for missing `acquiredRoot.Inspect` and `fileOpenFlags`. | Exact focused command PASS after private rolling reopen. | PASS with renamed/replaced ancestor, symlink, disappearance, identity mismatch, nonregular, nonlocal boundary, fstat/fstatfs, close-secondary, and close-order cases. | All required focused, scan race, platform, Linux, full, vet, and format checks PASS. |

### Verification

- `go test ./internal/platform/macos -run '^(TestWalkerRelativeFileReopen|TestWalkerFileIdentityChain|TestWalkerRegularFileFacts|TestWalkerFileReopenLifecycle)$' -count=1` PASS.
- `go test ./internal/scan -run '^(TestScannerBoundedFileJobs|TestScannerFileBackpressure|TestScannerRegularOnlyFacts|TestScannerFileEntryChanges)$' -count=1` PASS; `go test -race ./internal/scan -count=1` PASS.
- `go test ./internal/platform/macos -count=1`, `GOOS=linux go test -c -o /tmp/osdycleaner-wu5e1b-linux.test ./internal/platform/macos` (then removed), `go test ./... -count=1`, `go vet ./...`, and clean `gofmt -d` PASS.
- Runtime N/A: disposable Darwin `t.TempDir()` fixtures and narrow descriptor seams are the intended bounded harness. No LSP/lens/index operation was run under the delegated scope.

### Files and rollback

- Changed: `internal/platform/macos/walker_darwin.go`, `internal/platform/macos/walker_darwin_test.go`, and the minimal opaque error accessor in `internal/scan/descriptor.go`.
- Final SHA-256: `descriptor.go` `2586e570c30fce08471dc88ad89e51433864247734c50da06c77b990a8e5991c`; `walker_darwin.go` `cf4b9e43a949ce9af570db24419c1e9b4c2b326c0935a077445902d107419c07`; `walker_darwin_test.go` `7176c7512221bfb4e0e276b3b28261e7b2ac150538f2eda8915c1c548b5afb1f`.
- Rollback: restore these three files from `/tmp/osdy-wu5e1b-baseline-Y55fRv`; do not alter E2.

### Status and delivery boundary

- Consumed authoritative native OpenSpec status: `applyState: ready`, `artifactStore: openspec`, change `read-only-scan-foundation`, repo-local workspace and workspace edit root, strict TDD enabled, no status blockers. The standing `stacked-to-main` authorization covers E1B.
- The parent-provided active attempt token was continued. Per delegated `no settle` scope, it was not settled here.
- Workload/PR boundary: WU5E1B only (284 authored lines); no commit, PR, review, receipt approval, or E2 work occurred. Next lifecycle action is parent-owned independent verification.

---

## WU5E1B verifier correction — blocked for mandatory split

- Consumed authoritative native OpenSpec status for `read-only-scan-foundation`: `applyState: ready`, repo-local workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, allowed edit root the workspace, strict TDD enabled, and no status blocker. The native status records the parent-held active attempt; no attempt acquire, settle, index, or lifecycle action was performed by this executor.
- The delegated correction requires a new nonvacuous production-shaped scanner integration test plus `consumeDirectories` pipeline wiring: dynamic `FilePort` assertion, accepted directory identity map, relative ancestor-chain job construction, bounded pipeline submit/close/drain, accepted-result consumption, and ordering proof before `root.Close`.
- Current active-attempt accounting supplied by the parent is `+292/-1=293`, leaving 107 authored changed lines. The minimal honest RED integration harness and GREEN wiring cannot fit in 107 lines while proving the mandated event trace, exact job identities, changed/missing negative cases, and no raw path/FD behavior. Per the hard 400-line guard, implementation stopped before any Go or test edit.
- Persisted task reconciliation: WU5E1B RED/GREEN/TRIANGULATE/REFACTOR remain visibly `[x]` as historical pending-verification evidence; no task checkbox was changed. WU5E2 remains unchecked and blocked pending independently accepted E1B.
- Required next delivery decision: authorize a separately bounded correction work unit/successor with a fresh <=400-line budget, retaining the existing E1B baseline and targeting only `internal/scan/scanner.go` and `internal/scan/scanner_test.go` plus this receipt. No settle was requested or performed.

### TDD Cycle Evidence

| Task | Test file / layer | Safety net | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- | --- | --- |
| WU5E1B verifier correction | `internal/scan/scanner_test.go` / production-shaped scanner integration | Not run: hard line-budget gate stopped before edit. | Not started: a minimally nonvacuous RED cannot be completed within the remaining 107 lines. | Not started. | Not started. | Not started. |

---

## WU5E1B1 scanner FilePort wiring correction plan

- **Status consumed:** authoritative OpenSpec `read-only-scan-foundation`, `applyState: ready`, repo-local workspace and allowed root `/Users/osdy/Documents/GitHub/OsdyCleaner`, strict TDD active with `go test ./...`, and pre-authorized stacked-to-main delivery. No action-context warning.
- **Remediation boundary:** this successor remedies failed evidence `sha256:eb2b5c0a157d0465d7b8f702d2eb9b8d22d5ce5b4964b501e9b5523db4340ca6`; only future implementation and independent verification may satisfy it. WU5E1B remains checked historical wiring-incomplete evidence.
- **Planned scope:** only `internal/scan/scanner.go` and `internal/scan/scanner_test.go` plus append-only OpenSpec evidence, forecast 180–360 authored lines. The strict RED → GREEN → TRIANGULATE → REFACTOR integration test must prove dynamic `FilePort` selection, regular-fact identity-chain jobs, bounded submit/close/drain, accepted consumption, changed/missing negatives, deterministic ordering, and pipeline completion before `root.Close`.
- **Safety/deferred scope:** preserve read-only metadata-only behavior; no absolute descendant path, raw FD, content read, Darwin/descriptor-contract/finalizer edit, index change, commit, review, or lifecycle action. WU5E2 now depends on independently accepted WU5E1B1.
- **Task reconciliation:** four unchecked implementation-owned WU5E1B1 rows were inserted; historical E1B rows and all parent-owned rows remain byte-for-byte unchanged. Recount: 162 total = 123 checked + 39 unchecked; 159 implementation-owned = 123 checked + 36 unchecked; 3 parent-owned = 0 checked + 3 unchecked.
- **Verification/readback:** planning only; no Go/test command was run and no production/test file changed. Re-read `tasks.md` confirms the four WU5E1B1 rows are visibly `[ ]` and E1B/E2 completion was not claimed.
- **Planning accounting:** task/progress documentation only; authored implementation delta is +0/-0. Reconstructed pre-plan comparison records tasks +22/-6 and progress +11/-0, for planning patch **+33/-6=39** authored changed lines, within the parent’s 100-line cap.

---

## WU5E1B1 receipt — Scanner FilePort wiring correction

**Receipt predecessor:** SHA-256 `755518aaed2efcc18ee54ff358bb4e79a1b0e0a6c49b4a14af24f580f2f1d93b` was the apply-progress artifact observed before this work. This entry is append-only. No attempt lifecycle action, index change, commit, or PR was performed.

### Completed implementation and persisted task reconciliation

- Completed and visibly checked in `tasks.md`: WU5E1B1 RED, GREEN, TRIANGULATE, and REFACTOR. The task artifact was reread after the checkbox update; its exact ledger is **127 checked and 35 unchecked** overall. The next implementation unit is WU5E2, pending independent verification of this correction.
- `Scanner.consumeDirectories` dynamically selects `FilePort` only from the opaque acquired root. It keeps accepted directory identities keyed by relative component chain, creates a `FileJob` only for matching local/same-device enumerated and opened regular facts, and retains every accepted directory ancestor identity plus the matching final identity.
- The scanner starts one bounded `FilePipeline`, drains results concurrently to preserve bounded submit/result progress, submits validated relative jobs, closes job input once, waits for drain completion, and only then returns to `Scan`, where `root.Close` runs. Invalid, changed, missing, nonregular, boundary, and error facts create no file job.
- No descriptor contract or Darwin file changed. The scan package introduces no absolute descendant path, raw FD, content read, mutation, cleanup, or Phase-1 boundary expansion.

### TDD Cycle Evidence

| Task | Test file / layer | Safety net | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- | --- | --- |
| WU5E1B1 RED/GREEN/TRIANGULATE/REFACTOR | `internal/scan/scanner_test.go` / synchronized scanner unit fake | `go test ./internal/scan -count=1` PASS before edits. | `go test ./internal/scan -run '^TestScannerFilePortWiring$' -count=1` exited 1 with `jobs=[], want only accepted regular facts`; the test was added before production wiring. | The same focused command PASS after dynamic `FilePort` selection, relative job construction, and close/drain ordering. | Added a third accepted direct regular job after GREEN to exercise one-slot submit/result pressure while retaining nested ancestor-chain, direct-file, changed, missing, and special facts; the focused command PASS. | `gofmt` plus focused, race, full, vet, and source-boundary checks PASS; no behavior-changing refactor was needed after formatting. |

### Verification and boundary evidence

- `gofmt -w internal/scan/scanner.go internal/scan/scanner_test.go` -> PASS.
- `go test ./internal/scan -run '^TestScannerFilePortWiring$' -count=1` -> PASS.
- `go test -race ./internal/scan -count=1` -> PASS.
- `go test ./... -count=1` -> PASS (`internal/core`, `internal/platform/macos`, and `internal/scan`).
- `go vet ./...` -> PASS with no diagnostics.
- `gofmt -d internal/scan/scanner.go internal/scan/scanner_test.go` -> clean. Source-boundary scan found no `os.`, `syscall`, `unix.`, `filepath.`, content read, or `Open` call in the two allowed scan files.
- `TestScannerFilePortWiring` uses only a production-shaped opaque retained-root/FilePort fake. It proves two direct/nested accepted chains plus a third pressure job are inspected, changed/missing/special facts are excluded, and the nested inspection event occurs before `close:npm-cache`; result-channel pressure cannot permit close ahead of drain because worker completion requires the concurrent drain.
- Runtime: N/A — this work unit has no executable/CLI boundary; the synchronized retained-root/FilePort fake is the bounded runtime harness.

### Accounting, identity, rollback, and status

- Baseline: accepted WU5E1A scanner hashes recorded in `tasks.md` were `scanner.go` `672f14c0e410c9aef07b5eacadec53f6b8524e484a8da266c43c87f40342deb6` and `scanner_test.go` `6abfa3912d62446a74f207f227832379c79c6df9c92134cdf152bab55acd6417`; both were confirmed as the pre-edit workspace hashes. The mixed index was left untouched.
- Exact correction accounting from that E1A baseline: `scanner.go` **+50/-0** and `scanner_test.go` **+128/-0**, totaling **+178/-0 = 178 authored changed lines**, within the 400-line cap. Generated files: none.
- Final product hashes: `internal/scan/scanner.go` `a90f8b0d7725076876c52195a9e7d2537145dec2347efcdf98e86b71dbe7096d`; `internal/scan/scanner_test.go` `fccc82b33e223560bdd92d6cef6bf1d7de8136c2e8a4911eef66e487cf5c8a77`. Final task artifact hash: `50df1fe5295d5d4a559a86dad606ed28ba5fc8b8d9c94b36984c97d285bb1bfb`.
- Rollback boundary: restore exactly the accepted E1A scanner hashes above for `internal/scan/scanner.go` and `internal/scan/scanner_test.go`, and revert only the four WU5E1B1 checkbox changes plus this append-only receipt. This removes only FilePort wiring and its integration test; it does not alter D2A1 staged files, Darwin E1B history, or later work.
- Consumed authoritative status: OpenSpec `read-only-scan-foundation`, `applyState: ready`, workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, allowed root this workspace, strict TDD active, and `stacked-to-main` delivery authorization. No action-context warning or unsafe edit root was present. The WU5E1B1 PR/work-unit boundary is this independent 178-line correction; no commit or PR was created.

---

## WU5E1B1 gatekeeper correction receipt

- Scope was limited to `internal/scan/scanner.go` and `internal/scan/scanner_test.go`; no Darwin, descriptor, finalizer, task-marker, index, commit, review, settlement, or lifecycle operation occurred.
- `CloseJobs` now cancels only a submission-release context. Accepted envelopes retain their caller context through `Inspect`, so a normal close releases blocked submitters while accepted context-honoring work completes. The close coordinator still waits submitters, closes jobs once, joins workers, and closes results once.
- Scanner draining now retains every raw result internally, normalizes accepted results before its exact consumer seam, and maps rejected/error results to existing deterministic warning/status flags. The production-shaped root fake records sorted `consume:a`, `consume:b`, and `consume:dir/z` before `close:npm-cache`; changed, missing, and special facts still create no job.
- Cancellation remains caller-owned: the focused E1A backpressure test cancels the accepted caller context and observes `context.Canceled`; normal close is separately exercised with context-honoring work released normally. Existing scanner fact tests cover `WalkFacts` error mapping and submit pressure; this correction does not introduce WU5E2-wide cancellation/finalizer work.

### TDD Cycle Evidence

| Task | Test file / layer | Safety net | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- | --- | --- |
| WU5E1B1 gatekeeper correction | `internal/scan/scanner_test.go` / synchronized scanner integration | `go test ./internal/scan -run '^(TestScannerFilePortWiring | TestScannerBoundedFileJobs | TestScannerFileBackpressure | TestScannerFileEntryChanges)$' -count=1` PASS before RED | Added context-honoring inspection and exact sorted consume-before-close assertions; `go test ./internal/scan -run '^TestScannerFilePortWiring$' -count=1` exited 1 because the candidate discarded all drained results. | The same focused command passed after separation of submission shutdown from inspection context and normalized-result consumption. | Existing bounded-pressure, caller-cancellation, changed/missing/special no-job, and deterministic direct/nested result cases passed with the corrected integration. | `gofmt` plus focused, E1A, race, full, vet, and source-boundary checks passed. |

### Verification

- `go test ./internal/scan -run '^TestScannerFilePortWiring$' -count=1` -> PASS.
- `go test ./internal/scan -run '^(TestScannerBoundedFileJobs|TestScannerFileBackpressure|TestScannerRegularOnlyFacts|TestScannerFileEntryChanges)$' -count=1` -> PASS.
- `go test -race ./internal/scan -count=1` -> PASS.
- `go test ./... -count=1` -> PASS; `go vet ./...` -> PASS.
- `gofmt -w internal/scan/scanner.go internal/scan/scanner_test.go` followed by empty `gofmt -d` -> PASS. Source-boundary grep found no `os.`, `syscall`, `unix.`, `filepath.`, `ReadFile`, or `Open(` use in the two allowed files.
- Runtime: N/A — the synchronized opaque retained-root/FilePort fake is the bounded scanner harness.

### Identity, accounting, rollback, and status

- Candidate start hashes supplied by the parent were `scanner.go` `a90f8b0d7725076876c52195a9e7d2537145dec2347efcdf98e86b71dbe7096d` and `scanner_test.go` `fccc82b33e223560bdd92d6cef6bf1d7de8136c2e8a4911eef66e487cf5c8a77`. Corrected final hashes are `scanner.go` `999ca47f268c09b2800a2f6c6e69a31d3db2656de56247a508c3e8f50d62dd5e` and `scanner_test.go` `c67fd5f0f3bbf6c549971db9b624587e121bc67569a1bb4f7a44781607943cf0`.
- Reproducible accounting limitation: no materialized E1A baseline or candidate snapshot exists in this workspace or `/tmp`; both scan files are untracked, and the Git index is an older partial baseline. Therefore an exact cumulative WU5E1B1 delta cannot honestly be recomputed here. The parent-supplied candidate accounting is +178/-0 with 222 lines remaining; this correction was kept within that remaining slice, but no fabricated cumulative total is claimed. The reproducible method for verification is to restore the recorded E1A scanner hashes into a materialized baseline, then run `git diff --no-index --numstat <baseline>/scanner.go internal/scan/scanner.go` and the corresponding test command.
- Rollback: restore the supplied candidate hashes/content for only these two scanner files; that removes the gatekeeper correction without touching Darwin E1B history or WU5E2.
- Consumed fallback authoritative status: OpenSpec change `read-only-scan-foundation`; proposal, `specs/read-only-scan/spec.md`, design, tasks, and prior apply-progress were read; `applyState: ready`; repo-local workspace and allowed root `/Users/osdy/Documents/GitHub/OsdyCleaner`; strict TDD active; stacked-to-main correction slice authorized; no action-context warning.
- Persisted task reconciliation: `tasks.md` was reread and WU5E1B1 RED/GREEN/TRIANGULATE/REFACTOR remain visibly `[x]` as corrected implementation evidence pending independent verification. WU5E2 remains unchecked. No task marker or checkbox was modified because the correction repairs the already checked WU5E1B1 candidate.

## WU5E1B2 planning receipt

- Planning only: no Go/test execution, implementation, lifecycle, review, or task completion occurred; status was OpenSpec ready, repo-local workspace allowed, strict TDD, stacked-to-main, with no action-context warning.
- Independent verification failed E1B1 evidence `sha256:ebb7e73a46529e868bac5b01b608c1f5d1490b02cf1780d22ea70c3d6a9488c4`; its four checked rows remain historical and non-accepting.
- WU5E1B2 is the <=400-line strict RED→GREEN→TRIANGULATE→REFACTOR successor, scoped only to scanner source/test plus append-only evidence: scanner-owned normalized accepted-result retention before root close, a context-honoring production-shaped test, and existing-warning rejection classification.
- WU5E2 now depends on independently accepted WU5E1B2; hard links, cross-root finalization, and broad cancellation/status work remain excluded.
- Task readback: 166 total = 127 checked + 39 unchecked; 163 implementation-owned = 127 checked + 36 unchecked; parent-owned = 3 unchecked.
- Planning accounting: tasks +15/-5 and receipt +10/-0 = **30** authored lines; cumulative **387/400**.

---

## WU5E1B2 receipt — scanner-owned accepted-result retention

- Completed and persisted: WU5E1B2 RED, GREEN, TRIANGULATE, and REFACTOR are visibly `[x]` in `tasks.md`. WU5E2 remains unchecked; no parent-owned task changed.
- Scanner now owns normalized accepted `FileResult` evidence in private state and exposes it through `Scanner.AcceptedFileResults()`. The accessor deep-copies the result slice plus each retained job's components and ancestor identities.
- The private fake-only `fileResultConsumer` seam was removed. Accepted results are retained before `consumeDirectories` returns, hence before `Scan` invokes `root.Close`; rejected results remain excluded and use the existing warning/status classification.
- This intentionally does not perform WU5E2 finalization, hard-link attribution, cross-root aggregation beyond retained evidence, or broad cancellation/status changes. Existing `CloseJobs` semantics remain unchanged.

### TDD Cycle Evidence

| Task | Test file / layer | Safety net | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- | --- | --- |
| WU5E1B2 | `internal/scan/scanner_test.go` / production-shaped scanner integration | `go test ./internal/scan -run '^(TestScannerFilePortWiring | TestScannerBoundedFileJobs | TestScannerFileBackpressure | TestScannerRegularOnlyFacts | TestScannerFileEntryChanges)$' -count=1` PASS | Added `TestScannerAcceptedFileResults`; `go test ./internal/scan -run '^TestScannerAcceptedFileResults$' -count=1` exited 1 because `Scanner.AcceptedFileResults` did not exist. | The same focused test passed after scanner-owned deep-copy retention and accessor implementation. | The test covers unordered accepted results, context propagation, caller mutation, close-time retention, and changed/missing/nonregular rejections mapping to existing warnings; focused test PASS. | `gofmt` and all required focused, race, platform, full, vet, and source-boundary checks PASS. |

### Verification

- `go test ./internal/scan -run '^(TestScannerAcceptedFileResults|TestScannerFilePortWiring)$' -count=1` -> PASS.
- `go test ./internal/scan -run '^(TestScannerBoundedFileJobs|TestScannerFileBackpressure|TestScannerRegularOnlyFacts|TestScannerFileEntryChanges)$' -count=1` -> PASS.
- `go test -race ./internal/scan -count=1` -> PASS.
- `go test ./internal/platform/macos -run '^(TestWalkerRelativeFileReopen|TestWalkerFileIdentityChain|TestWalkerRegularFileFacts|TestWalkerFileReopenLifecycle)$' -count=1` -> PASS.
- `go test ./... -count=1` -> PASS; `go vet ./...` -> PASS; `gofmt -d internal/scan/scanner.go internal/scan/scanner_test.go` was empty.
- Source-boundary checks found no `fileResultConsumer`, absolute-descendant, raw-FD, content-read, filesystem-call, or platform-syscall surface in `scanner.go`.
- Runtime: N/A — this bounded scanner integration uses opaque context-honoring root/FilePort fixtures and provides no executable/CLI runtime boundary.

### Accounting, rollback, and status

- Initial hashes: `scanner.go` `999ca47f268c09b2800a2f6c6e69a31d3db2656de56247a508c3e8f50d62dd5e`; `scanner_test.go` `c67fd5f0f3bbf6c549971db9b624587e121bc67569a1bb4f7a44781607943cf0`.
- Final hashes: `scanner.go` `38c5f6288da83c40175a363769c612d442f5e3be23884726db593022e50ca1cf`; `scanner_test.go` `4167f5789e76cc6e120e2ddf628f50f69b939f40f7e56322f340952943e837ca`.
- Attempt-relative accounting used temporary reconstructions whose SHA-256 values matched both supplied initial hashes: `scanner.go` +53/-11 and `scanner_test.go` +86/-28, for **139 additions + 39 deletions = 178 authored changed lines**, within the 400-line cap.
- Rollback boundary: restore only the two supplied initial scanner-file contents/hashes, revert the four WU5E1B2 checkboxes, and remove this appended receipt; WU5E1A/E1B history and WU5E2 remain untouched.
- Consumed authoritative status: OpenSpec `read-only-scan-foundation`, `applyState: ready`, repo-local workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, allowed workspace root, strict TDD active, `stacked-to-main` delivery authorization, and no action-context warning. Parent lifecycle remains required; no acquire, settle, reset, stage, commit, or review operation occurred.
- Task reconciliation after completion: 166 total rows; 131 checked and 35 unchecked overall. The four completed WU5E1B2 rows are `[x]`; exact remaining WU5E2 rows are still `[ ]`.

---

## WU5E2 receipt — cancellation, deterministic finalization, and hard links

**Predecessor / baseline:** baseline tree `ddd4128518c4aa1f7bf127bcd3c5a840579c2aa5`; accepted WU5E1B2 source materialized from dangling tree `e00bf7482f909e859aa44760e17ff935057db8f1`, with supplied SHA-256 `scanner.go` `38c5f6288da83c40175a363769c612d442f5e3be23884726db593022e50ca1cf` and `scanner_test.go` `4167f5789e76cc6e120e2ddf628f50f69b939f40f7e56322f340952943e837ca`.

### Completed tasks and implementation

- Persisted `[x]`: WU5E2 RED, GREEN, TRIANGULATE, and REFACTOR; the WU5E2 rollback text now names the accepted WU5E1B2 scanner hashes.
- `FilePipeline.Stop` prevents workers from accepting queued jobs after scan-context cancellation, preserves any active result for draining, closes input once, and permits fixed workers to join before root close.
- Scanner-owned accepted results are converted to canonical `rawObservation` values and consumed through the existing `finalizeRawRoots` finalizer; finalizer ordering therefore chooses same- and cross-area identity attribution instead of worker completion order.
- Changed/error results retain existing zero-contribution warning mapping. Cancellation preserves already-completed roots, marks the active root cancelled, and marks later roots skipped.

### TDD Cycle Evidence

| Stage | Evidence |
| --- | --- |
| Safety net | `go test ./internal/scan -count=1` passed before edits. |
| RED | The exact WU5E2 selector exited 1 because `Scanner.FinalizedFileFacts` and synchronized root-start evidence were absent. |
| GREEN | The selector passed after cancellation stop/drain/join and scanner-to-finalizer conversion were implemented. |
| TRIANGULATE | Completion permutations, cross-area same-identity attribution, cancelled queued work, repeated cancellation/join, and changed-result warning permutations passed. |
| REFACTOR | `gofmt`, focused/race/full tests, vet, and source-boundary scan passed without behavior expansion. |

### Verification

- Focused: `go test ./internal/scan -run '^(TestScannerFileCompletionPermutations|TestScannerHardLinkFinalization|TestScannerCancellationStopDrainJoin|TestScannerNoPipelineLeaks|TestScannerDeterministicWarnings)$' -count=1` -> PASS.
- Focused race: same selector with `go test -race` -> PASS.
- E1A/E1B2 regressions: `go test ./internal/scan -run '^(TestScannerBoundedFileJobs|TestScannerFileBackpressure|TestScannerRegularOnlyFacts|TestScannerFileEntryChanges|TestScannerAcceptedFileResults)$' -count=1` -> PASS.
- Full scan race: `go test -race ./internal/scan -count=1` -> PASS.
- Full suite: `go test ./... -count=1` -> PASS; `go vet ./...` -> PASS; `gofmt -d` was clean.
- Source-boundary grep found no `WalkDir`, path open/read/write, mutation, or traversal call in changed production files. Runtime: N/A — channel-synchronized fake roots are the intended bounded harness and no CLI exists.

### Accounting, identity, and rollback

- Exact WU5E2 product/test delta from materialized accepted WU5E1B2 sources: `scanner.go` `+74/-11`; `scanner_test.go` `+141/-0`; `finalize.go` and `finalize_test.go` `+0/-0`; **+215/-11 = 226 authored changed lines**, within the 400-line cap.
- Final SHA-256: `scanner.go` `6a86a4e50cec0c762cda795678c7a35c2c09bd144704a607eda218adbc666d5d`; `scanner_test.go` `9346e71a490260dd9483cd6024178ad4a45f4aeaa75c271d182ea35c7c673764`; unchanged `finalize.go` `2ffd49fcdd5011fe5012bdee31332468eeb2f0afc42a136c39ad8dfbc36c0479`; unchanged `finalize_test.go` `2ec604ae858ecad200e58a6ffc083898a35a0050c070fcc66fa500e614db3719`.
- Rollback: restore the accepted WU5E1B2 scanner hashes above, removing only WU5E2 cancellation/retention/finalizer integration and its five tests. No index, lifecycle, commit, review, or runtime operation was altered.

### Status and workload boundary

- Consumed authoritative OpenSpec status reconstructed from the provided apply-ready context: `changeName: read-only-scan-foundation`, `artifactStore: openspec`, `applyState: ready`, `actionContext.mode: repo-local`, workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, and supplied allowed WU5E2 edit surfaces. Strict TDD was active; no action-context warning applied.
- Work-unit / PR boundary: WU5E2 only, sequential `stacked-to-main`, 226 authored lines. No deviation from design beyond retaining the existing finalizer files unchanged.
- Remaining implementation rows begin at WU6; WU6 is deliberately not implemented by this work unit. Route next to independent `sdd-verify` under parent lifecycle authority.

---

## WU5E2 correction attempt 2 — exported finalized boundary and scan-state preservation

- **Authority:** correction attempt `sha256:8f57295898784bd8dcd6f22ce47e16fe2aa8d98888b167b01071e082519ccb8c`; remediates failed verification evidence `sha256:b883e1a9538f621100d81593f9ba6d27a2c98e6ac7eb1a7e92c7764d39dffa73` from verifier `subtask_gentle-ai-verify_1788555636004_95a7be85`.
- **Candidate identity:** failed scanner hashes were `6a86a4e50cec0c762cda795678c7a35c2c09bd144704a607eda218adbc666d5d` and `9346e71a490260dd9483cd6024178ad4a45f4aeaa75c271d182ea35c7c673764`; failed finalizer hashes were `2ffd49fcdd5011fe5012bdee31332468eeb2f0afc42a136c39ad8dfbc36c0479` and `2ec604ae858ecad200e58a6ffc083898a35a0050c070fcc66fa500e614db3719`.
- **Correction:** `FinalizedFileFacts` is now an exported immutable-by-convention boundary with copying `Roots`, `Findings`, and `Warnings` accessors. Scanner-owned retained roots are recorded after `Scan`; finalization retains each observed root's actual status, reason, and warnings, and only overlays canonical finalizer estimates/finding counts for scanned roots. It therefore cannot synthesize completed state for partial, boundary-limited, cancelled, or unstarted roots. The worker checks stop before and after queue receipt so a ready queued job cannot win a post-stop select race; active results still publish and drain.
- **Five-root markers:** the new finalized-boundary test proves npm `cancelled/cancelled_during_scan`; Homebrew, Gradle, Xcode DerivedData, and CoreSimulator remain `skipped/cancelled_before_start` rather than synthetic scanned/completed roots.

### TDD Cycle Evidence

| Stage | Evidence | Result |
| --- | --- | --- |
| RED | Added `TestScannerFinalizedFactsPreserveRootObservations` against the failed candidate. | `go test ./internal/scan -run '^TestScannerFinalizedFactsPreserveRootObservations$' -count=1` exited 1 because `finalizedFacts` had no public accessors. |
| GREEN | Added exported snapshot-ready result accessors and scanner-owned exact root retention/merge. | New focused test passed. |
| TRIANGULATE | Added channel-synchronized before-enqueue cancellation, real two-worker completion-order, and scanner-level same-area hard-link checks. | Focused and race selectors passed. |
| REFACTOR | Formatted allowed scan files without behavioral compression. | Full scan race, full suite, vet, and format checks passed. |

### Verification

- WU5E2 focused: `go test ./internal/scan -run '^(TestScannerFileCompletionPermutations|TestScannerHardLinkFinalization|TestScannerCancellationStopDrainJoin|TestScannerNoPipelineLeaks|TestScannerDeterministicWarnings)$' -count=1` — PASS.
- WU5E2 focused race: same selector with `go test -race` — PASS.
- Correction coverage selector: `TestScannerFinalizedFactsPreserveRootObservations`, `TestScannerWorkerCompletionPermutations`, and `TestScannerCancellationBeforeEnqueue` — PASS, including under `-race`.
- E1A/E1B2 regression selector (`TestScannerBoundedFileJobs|TestScannerFileBackpressure|TestScannerRegularOnlyFacts|TestScannerFileEntryChanges|TestScannerFilePortWiring|TestScannerAcceptedFileResults|TestScannerProductionTraversalCompletesBeforeRootClose`) — PASS.
- Full scan race `go test -race ./internal/scan -count=1`, full suite `go test ./... -count=1`, `go vet ./...`, and `gofmt -d` source-boundary check — PASS.
- Runtime: N/A — this is a synchronized scan/pipeline fixture boundary with no CLI or executable runtime surface.

### Accounting, hashes, rollback, and ledger

- The first WU5E2 candidate remains recorded at **+215/-11=226**. Exact correction-attempt additions/deletions cannot be reconstructed because the failed untracked candidate blobs were not retained in Git or `/tmp`; current source line totals are not substituted for a diff. This is an evidence limitation, not an acceptance claim; the correction remains under the explicit 400-line cap by the bounded edit scope.
- Final hashes: `scanner.go` `b440c8369a7f35873394d8e0067ffde65636f98f88c4b641ff673a3b9dfe5da9`; `scanner_test.go` `ed2d8318329f87dd427c9693c6ab08e8acffd3e2da561e32199a79d53a3b64e6`; `finalize.go` `8bb3c4e6a66a5481b15a35391c037cd535c8d142baf473ec6fcb012aea54bf89`; `finalize_test.go` `2ec604ae858ecad200e58a6ffc083898a35a0050c070fcc66fa500e614db3719`.
- Rollback boundary: restore the four failed-candidate hashes above; this removes only WU5E2 result-boundary, root-retention, cancellation-fairness, and correction test behavior, leaving WU5E1B2 intact.
- Task ledger reconciliation: all four WU5E2 implementation rows remain visibly `[x]` after correction evidence; WU6 remains unchecked. No lifecycle, index, commit, review, or settlement operation was performed.
- Consumed status: manual authoritative OpenSpec status for `read-only-scan-foundation`, artifact store `openspec`, repo-local workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, strict TDD enabled, allowed edits limited to the four scan files plus this append-only progress receipt and WU5E2 task reconciliation. The parent provided `stacked-to-main` delivery authority; action-context warnings were none.
- PR boundary: WU5E2 correction attempt 2 only; parent owns lifecycle routing. Independent semantic verification is still required before treating this correction as accepted.

---

## WU5E2 fresh correction — incomplete-root accepted totals

- **Authority:** parent token `sha256:ff080b972e7fa30fc94f7f2b6d50b7b86f2bd2a528fa468112fe9080c685b5c7`; remediates `sha256:2e28ea4789d2bcaf2adfa4a95e904f427ad58f995f1c55da98ff6295db8150e8`. No acquire, settle, reset, staging, commit, review, or lifecycle operation was performed; lifecycle remains parent-owned.
- **Baseline verification:** the four permitted files matched `/tmp/osdy-wu5e2-fresh-baseline-20260904` before edits: `scanner.go` `b440c8369a7f35873394d8e0067ffde65636f98f88c4b641ff673a3b9dfe5da9`, `scanner_test.go` `ed2d8318329f87dd427c9693c6ab08e8acffd3e2da561e32199a79d53a3b64e6`, `finalize.go` `8bb3c4e6a66a5481b15a35391c037cd535c8d142baf473ec6fcb012aea54bf89`, and `finalize_test.go` `2ec604ae858ecad200e58a6ffc083898a35a0050c070fcc66fa500e614db3719`.
- **Correction:** `scannerFileResults.finalized` now merges canonical accepted estimate and finding count for every observed root that generated facts. It keeps the observed status, reason, and warning codes exactly; non-scanned roots receive the retained known totals with incomplete estimate completeness. No rejected, changed, disappeared, or non-regular result is added to `rawFileObservations`.

### Exactly five root read-only markers

1. `npm-cache`: partial accepted facts retain bytes/count, `entry_visibility_gap`, and its observed warning code.
2. `homebrew-cache`: boundary-limited accepted facts retain bytes/count, `device_boundary`, and its observed warning code.
3. `gradle-caches`: inaccessible root remains zero-contribution with its root-inspection status and warning.
4. `xcode-derived-data`: cancelled accepted facts retain bytes/count, `cancelled_during_scan`, and its observed warning code.
5. `core-simulator`: later skipped root remains zero-contribution with `cancelled_before_start` context.

### TDD Cycle Evidence

| Stage | Evidence | Result |
| --- | --- | --- |
| Safety net | `go test ./internal/scan -count=1` before edits | PASS. |
| RED | Added `TestScannerFinalizedFactsRetainAcceptedTotalsForIncompleteRoots` before production changes. | `go test ./internal/scan -run '^TestScannerFinalizedFactsRetainAcceptedTotalsForIncompleteRoots$' -count=1` exited 1: partial root retained zero bytes/count instead of accepted `10/1`. |
| GREEN | Merge generated estimate/count for every generated observed root and derive incomplete non-scanned estimates. | Same focused test PASS. |
| TRIANGULATE | The five-root mixed-state test asserts partial, boundary-limited, inaccessible, active cancellation, and later skipped facts. `TestScannerCancellationAfterWorkerAcceptanceAndWhileBlocked` separately synchronizes post-worker-acceptance and queued-blocked cancellation. | New focused and race selectors PASS; active accepted result drains and joins, and queued work never calls `Inspect` after stop. |
| REFACTOR | Ran gofmt without behavior compression. | Focused/race, full scan race, repository tests, vet, and format/source-boundary checks PASS. |

### Verification

- WU5E2 focused selector and race selector: `go test ./internal/scan -run '^(TestScannerFileCompletionPermutations|TestScannerHardLinkFinalization|TestScannerCancellationStopDrainJoin|TestScannerNoPipelineLeaks|TestScannerDeterministicWarnings)$' -count=1` and the same command with `-race` -> PASS.
- New exact coverage: `go test ./internal/scan -run '^(TestScannerFinalizedFactsRetainAcceptedTotalsForIncompleteRoots|TestScannerCancellationAfterWorkerAcceptanceAndWhileBlocked|TestScannerFinalizedFactsPreserveRootObservations|TestScannerWorkerCompletionPermutations|TestScannerCancellationBeforeEnqueue)$' -count=1` -> PASS; the two fresh tests under `-race` -> PASS.
- E1A/E1B2 regression selector -> PASS: `TestScannerBoundedFileJobs`, `TestScannerFileBackpressure`, `TestScannerRegularOnlyFacts`, `TestScannerFileEntryChanges`, `TestScannerFilePortWiring`, `TestScannerAcceptedFileResults`, and `TestScannerProductionTraversalCompletesBeforeRootClose`.
- `go test -race ./internal/scan -count=1`, `go test ./... -count=1`, and `go vet ./...` -> PASS. `gofmt -d` is clean; source boundary scan found no `WalkDir`, `os.Open`/read/write/remove, or `filepath.Walk` call in changed production files.
- Runtime: N/A — synchronized fake root/file-port channels are the designed bounded scan harness; no CLI or executable runtime boundary exists.

### Accounting, hashes, rollback, and task reconciliation

- Exact delta against the immutable fresh baseline: `scanner.go` `+26/-3`; `scanner_test.go` `+123/-0`; `finalize.go` `+0/-0`; `finalize_test.go` `+0/-0`; **+149/-3 = 152 authored changed lines**, below 400.
- Final SHA-256: `scanner.go` `d2f93151ee8ead09e5c878359e2dad5b305c3494659652c5af6a88df10819822`; `scanner_test.go` `69ae1c555037b9282b0995fdab84c1b4b097e9f18be45d1e530170ae31578b78`; `finalize.go` `8bb3c4e6a66a5481b15a35391c037cd535c8d142baf473ec6fcb012aea54bf89`; `finalize_test.go` `2ec604ae858ecad200e58a6ffc083898a35a0050c070fcc66fa500e614db3719`.
- Rollback: restore the four allowed scan files from `/tmp/osdy-wu5e2-fresh-baseline-20260904`; this removes only the WU5E2 fresh correction and leaves all other product surfaces untouched.
- Persisted task readback: ledger is **135 checked / 31 unchecked / 166 total**. All four WU5E2 implementation rows remain visibly `[x]`; they were already checked correction evidence, so no checkbox mutation was necessary. WU6 rows remain unchecked.
- Engram: verified fix saved to project `osdycleaner` as `sdd/read-only-scan-foundation/wu5e2-root-total-correction`.

### Status and remaining work

```yaml
schemaName: spec-driven
changeName: read-only-scan-foundation
artifactStore: openspec
applyState: ready
dependencies:
  apply: ready
  verify: ready
actionContext:
  mode: repo-local
  workspaceRoot: /Users/osdy/Documents/GitHub/OsdyCleaner
  allowedEditRoots: [/Users/osdy/Documents/GitHub/OsdyCleaner]
  warnings: []
nextRecommended: sdd-verify
```

- Status was manually produced under the global contract because no structured native status was supplied. Proposal, `specs/read-only-scan/spec.md`, design, tasks, prior apply-progress, and `openspec/config.yaml` were read. The authoritative workspace and all edits stayed within the parent-supplied allowed surfaces.
- Deviation: none. WU6 and all later unchecked implementation rows remain out of scope. Workload / PR boundary: fresh WU5E2 correction only, 152 authored lines, sequential `stacked-to-main`; parent owns lifecycle and independent verification routing.

## WU5E2 Cancellation Synchronization Correction and Receipt Recovery (2026-09-04)

- **Incident and recovery:** during persistence, the prior untracked append-only receipt was accidentally replaced by an 81-line condensed reconstruction. Parent diagnosis found the exact prior 2,873-line artifact in native candidate tree `3bddcd84958b9f551c7afe591f3696708cd14134`, blob `d37700ba7d9338dd16b6c0c6410ed36dbfedf50f`. With explicit user authorization, those exact bytes were restored before this addendum. The displaced reconstruction is retained outside the repository at `/tmp/osdy-apply-progress-accidental-reconstruction-81.md`; it is not authoritative.
- **Remediation target:** independent failure `sha256:35ee06deef85931407f1957bfa6a2c81cf3228708dcd34e94de82decdd8d0a67` reported scheduling-sensitive `TestScannerFinalizedFactsPreserveRootObservations`: the active root could resolve as `partial/entry_visibility_gap` instead of `cancelled/cancelled_during_scan`.
- **Correction:** the test now blocks the active inspection, invokes cancellation, waits for the inspection to acknowledge `ctx.Done()`, and only then releases it. Cancellation assertions were not weakened; queued work remains uninspected after stop, and accepted cancelled-root bytes/count/incomplete-estimate assertions remain intact.
- **TDD evidence:** the verifier-reported RED could not be reproduced locally over 100 pre-edit repetitions, establishing nondeterminism rather than a stable failure. After synchronization, the exact selector passed 100 repetitions. Focused WU5E2, correction, cancellation, focused race, full `internal/scan` race, E1A/E1B2 regression, full suite, `go vet ./...`, gofmt, and source-boundary checks passed.
- **Exact immutable-baseline accounting:** against `/tmp/osdy-wu5e2-fresh-baseline-20260904`, `scanner.go +26/-3`, `scanner_test.go +153/-15`, `finalize.go +0/-0`, `finalize_test.go +0/-0`; cumulative **+179/-18 = 197 authored lines**, within 400. Attempt-2 synchronization adds **+30/-15 = 45** beyond attempt 1.
- **Final hashes:** `scanner.go` `d2f93151ee8ead09e5c878359e2dad5b305c3494659652c5af6a88df10819822`; `scanner_test.go` `7c33daccb9b71322395b7e7d3d01d89ccff5829beb4bd2836764dc46deda02f3`; unchanged `finalize.go` `8bb3c4e6a66a5481b15a35391c037cd535c8d142baf473ec6fcb012aea54bf89`; unchanged `finalize_test.go` `2ec604ae858ecad200e58a6ffc083898a35a0050c070fcc66fa500e614db3719`.
- **Ledger and scope:** 166 tasks = 135 checked + 31 unchecked; all four WU5E2 rows remain checked and WU6 remains unchecked. Exactly five filesystem-root `(read-only)` markers remain. Runtime N/A. No index, commit, push, PR, review, reset, acquire, settle, descriptor/Darwin, report/CLI, WU6, or mutation work was performed by the correction subagent.
- **Rollback:** restore the four scan/finalize files from `/tmp/osdy-wu5e2-fresh-baseline-20260904` and remove only this addendum; the restored preceding 2,873-line receipt remains authoritative.

## WU6 Pre-Write Forecast Block (2026-09-04)

- **Structured status consumed:** authoritative native OpenSpec status `gentle-ai.sdd-status@2` reported `changeName: read-only-scan-foundation`, `artifactStore: openspec`, `applyState: ready`, `dependencies.apply: ready`, `nextRecommended: apply`, and repo-local action context rooted at `/Users/osdy/Documents/GitHub/OsdyCleaner` with that workspace as the allowed edit root. Proposal, `specs/read-only-scan/spec.md`, design, tasks, prior progress, and `openspec/config.yaml` were read. Strict TDD is active.
- **Prefix safeguard:** before this append, `apply-progress.md` was 447143 bytes with SHA-256 `a152c48a2ec04e83ee07d064ed10b273f6342774e4e78adf29c9f6a592a6b325`. Its existing bytes were copied and compared byte-for-byte after the append; they remain the exact prefix. No historical byte was changed or truncated.
- **WU6 rows inspected:** all four exact WU6 rows remain unchecked: RED, GREEN, TRIANGULATE, and REFACTOR. No checkbox was changed because no implementation task completed.
- **Pre-write forecast:** **+455–585/-5–25 = 460–610 authored changed lines**, exceeding the 400-line work-unit cap. The existing scanner has no injected scan-policy/default constructor, uses `fileWorkers: 1` and worker-sized queues, and `WalkLimits` contains only depth/descriptor fields. Enforcing the required 250,000-entry and 64 MiB path limits at the descriptor-walker boundary requires a compatible limits-contract extension (currently outside WU6's allowed edit surfaces) plus scanner cancellation/status integration. Six named channel-synchronized tests and triangulation coverage add an estimated 275–360 lines; policy/limit/cancellation/finalization production changes add an estimated 180–225 lines.
- **No-write result:** strict-TDD RED/GREEN/TRIANGULATE/REFACTOR did not start; no Go source or test was edited and no test command was run. Baseline/final source hashes are unchanged: `scanner.go` `d2f93151ee8ead09e5c878359e2dad5b305c3494659652c5af6a88df10819822`, `scanner_test.go` `7c33daccb9b71322395b7e7d3d01d89ccff5829beb4bd2836764dc46deda02f3`, `finalize.go` `8bb3c4e6a66a5481b15a35391c037cd535c8d142baf473ec6fcb012aea54bf89`, and `finalize_test.go` `2ec604ae858ecad200e58a6ffc083898a35a0050c070fcc66fa500e614db3719`; exact source delta is +0/-0. Runtime N/A.
- **Required decision / smallest dependency-safe split:** the existing standing stacked-to-main authorization explicitly ends at the WU5E correction chain and does not authorize changing WU6 rows. Authorize a two-slice WU6 split before implementation: **WU6A (≤400)** expands the descriptor/scanner limit-policy contract and adds default-policy plus entry/path-limit RED→GREEN→TRIANGULATE→REFACTOR tests; **WU6B (≤400, depends on WU6A)** adds cancellation state precedence, stop/drain/join, later-root skipping, and completion-order-stability tests/implementation. WU6A must include `internal/scan/descriptor.go` and `internal/scan/descriptor_test.go` as allowed surfaces because entry/path bounds must reach the walker. No task split or ledger reconciliation was made without that authorization.
- **Workload / rollback / markers:** PR boundary is blocked before WU6A; no PR/commit/lifecycle operation occurred. Rollback is N/A for the no-source-edit attempt; remove this append only if the forecast evidence itself must be reverted. The five fixed built-in root markers remain untouched.

## Authorized WU6A/WU6B Planning Split (2026-09-04)

### Authority and preservation

- Planning authority: change `read-only-scan-foundation`, parent token `sha256:5680b6707dd975d192c57c79785370b11d5edaac5e63e07d94ed34f3e6adaa5a`, accepted WU5E2 evidence `sha256:2d3dffdfaa2f726f05374aebc3ac43ffa3439fc49ab9745ad3fb47df104a4363`, and explicit user authorization to split the honest 460–610-line WU6 forecast. This transaction changed planning artifacts only.
- Append-only proof: the complete 2,895-line pre-split `apply-progress.md`, including the exact historical/recovery prefix and the later WU6 forecast block, was matched as the unique edit prefix and remains byte-for-byte at the start of this file. In particular, the previously attested 447143-byte prefix with SHA-256 `a152c48a2ec04e83ee07d064ed10b273f6342774e4e78adf29c9f6a592a6b325` remains unchanged and precedes the preserved WU6 forecast.
- Required full-file identity limitation: the injected read/edit tools expose neither byte counts nor digest computation. Therefore the actual pre-split full-file byte size/SHA-256 (historical prefix plus WU6 forecast) and post-append full-file byte size/SHA-256 are recorded as **unavailable—not fabricated**; the pre-split line count is exactly 2,895 and the append-only prefix was verified by exact-text replacement/readback. A later hash-capable read-only attestation may add those two full-file identities without rewriting this receipt.

### Authorized task replacement

- Replaced exactly the four unchecked WU6 RED/GREEN/TRIANGULATE/REFACTOR rows in `tasks.md` with eight unchecked implementation-owned rows in strict order: WU6A RED → GREEN → TRIANGULATE → REFACTOR, then WU6B RED → GREEN → TRIANGULATE → REFACTOR. No implementation row was checked.
- WU6A depends on independently accepted WU5E2 and is forecast at **+210–285/-5–20 = 215–305 authored changed lines**, independently hard-capped at 400. Its only product/test surfaces are `internal/scan/descriptor.go`, `internal/scan/descriptor_test.go`, `internal/scan/scanner.go`, and `internal/scan/scanner_test.go`.
- WU6A owns the validated policy/limits contract and defaults: 8 workers, job queue 256, result queue 256, one concurrently active root, 250,000 entries per root, and 64 MiB (67,108,864 bytes) cumulative relative-path budget per root. It owns exact-boundary and first-over-boundary entry/path tests and stable limit evidence; broad cancellation/status precedence, Darwin implementation, finalizer changes, reports, TUI, CLI, and lifecycle work are excluded.
- WU6B depends on independently accepted WU6A and is forecast at **+245–300/-0–5 = 245–305 authored changed lines**, independently hard-capped at 400. Its only product/test surfaces are `internal/scan/scanner.go`, `internal/scan/scanner_test.go`, `internal/scan/finalize.go`, and `internal/scan/finalize_test.go`.
- WU6B owns cancellation before, during, and after roots; accepted-result drain, fixed-worker join, shutdown order, later-root skipping, no-leak and completion-order determinism; and exact `cancelled > partial > boundary_limited > scanned` precedence across partial/boundary/WU6A-limit combinations. Descriptor/Darwin contract changes and all report/TUI/CLI/lifecycle work are excluded.
- WU7 now depends on independently accepted WU6B. Focused commands, runtime N/A rationales, independent acceptance gates, and rollback boundaries are explicit in each new task section.

### Forecast and ledger reconciliation

- Reconciled overall forecast: **7,205–10,045 additions, 223–758 deletions, 7,428–10,803 authored changed lines**, excluding generated `go.sum`. Budget risk remains High; chained PRs remain recommended; delivery remains resolved `ask-on-risk`; chain strategy remains `stacked-to-main`; decision before apply remains No.
- Ledger after split: **170 total = 135 checked + 35 unchecked; 167 implementation-owned = 135 checked + 32 unchecked; 3 parent-owned = 0 checked + 3 unchecked**. The four WU5E2 rows remain checked. Every row outside the replaced WU6 parent block retains its checkbox state, relative order, and terminal ownership marker.
- Verification performed: Markdown clean feedback from both planning edits; source-of-truth readback confirms the WU6A/WU6B sections, exact dependencies, focused commands, exclusions, forecasts, rollback boundaries, WU7 dependency, and this append. No Go test was run because this was planning-only; runtime is N/A because no executable or implementation boundary changed.
- Planning rollback boundary: restore only the prior four unchecked WU6 rows and WU7 dependency/forecast/count prose in `tasks.md`, then remove only this final append. No Go source, test, generated artifact, index, stage, commit, review, reset, acquire, settle, or filesystem content requires rollback.

### Exactly five filesystem-root markers

1. `npm-cache` — read-only built-in root remains fixed and receives no report or CLI scope.
2. `homebrew-cache` — read-only built-in root remains fixed and receives no report or CLI scope.
3. `gradle-caches` — read-only built-in root remains fixed and receives no report or CLI scope.
4. `xcode-derived-data` — read-only built-in root remains fixed and receives no report or CLI scope.
5. `core-simulator` — read-only built-in root remains fixed and receives no report or CLI scope.

## WU6A blocked receipt — 2026-09-04

- **Status consumed:** manual fallback `spec-driven` status because no structured status was supplied and `gentle-ai sdd status` produced no JSON; file-backed `openspec` is authoritative, `actionContext` is assumed `repo-local` at `/Users/osdy/Documents/GitHub/OsdyCleaner`, with no edit-root warning.
- **Artifact context read:** proposal, spec, design, tasks, config, existing scanner/descriptor code and tests, and the existing apply-progress receipt.
- **Receipt precondition:** apply-progress was `455933` bytes with SHA-256 `d5638b7080a12ce249cd5fc7aa61c6862cc9ed6a5b7abf4e6697d231fbf5d4a4` before this append.
- **Blocker:** parent-authorized WU6A declares accepted `internal/scan/descriptor.go` SHA-256 `aa191d41ad70864d3bc7ca954fbe8f6c31f7970693dfe7573460d17a5616d620`, but the remeasured workspace file is `2586e570c30fce08471dc88ad89e51433864247734c50da06c77b990a8e5991c` and already has `+56/-0` uncommitted lines. Its accepted test companion matches; scanner and scanner_test hashes match the declared WU5E2 baseline. This unaccepted/mismatched allowed target prevents safe strict-TDD RED edits.
- **No implementation or task state changed:** all four WU6A implementation checkboxes remain unchecked; no test command was run because RED was not started, no runtime harness applies, and no lifecycle operation occurred.
- **Workload / PR boundary:** explicit WU6A `stacked-to-main` authorization exists, but no work-unit slice was authored. **Ledger remains 135 checked / 35 unchecked / 170 total.**
- **Rollback:** no WU6A code exists to roll back; restore/reconcile the descriptor baseline to the parent-authorized accepted hash before retrying.

## WU6A completion evidence recovery (2026-09-04)

The prior receipt was restored byte-for-byte from native candidate tree `00cd02cdf8cf2d3296b927e13bbf84a191b31439`, blob `cb184104070382496ce8663aa3f94fd66877ad4e` (457,634 bytes), after an accidental documentation-only overwrite during final reconciliation. The later WU6A completion receipt was not present in the recoverable candidate tree; the following literal evidence is preserved from the immediately preceding on-disk receipt rather than invented.

- Completed/persisted: WU6A RED, GREEN, TRIANGULATE, and REFACTOR.
- Baseline: `/tmp/osdy-wu6a-baseline-20260904`; descriptor baseline `2586e570c30fce08471dc88ad89e51433864247734c50da06c77b990a8e5991c`.
- Focused WU6A selector, its race run, full scan race, WU5E2 regression selector, `go test ./...`, `go vet ./...`, gofmt, and source-boundary checks passed.
- Final hashes: `descriptor.go` `9cd69318b6077ae9081f16deb0bd035245c30ecc8fa395ec24be5a9c195050b6`; `descriptor_test.go` `087960977f5f667bae8b43626fc2f7f98fa09d2d7e10cd23077ba0317c31d44a`; `scanner.go` `81a4c9bb6994b5f7ea77194848edabd0eb7c270891dc7dd2a04d8fe1e9019be6`; `scanner_test.go` `f72d68f094ecbc4cd93e9e0452cf20927fca5b4425618425aaa96a1be27a123c`.
- Recomputed WU6A implementation delta: `+255/-3 = 258` authored lines, below the 400-line cap. Runtime N/A; rollback restores the four WU6A baseline files.


## Final ODD-to-OpenSpec reconciliation (2026-09-04)

### Scope and status

- Consumed parent-supplied authoritative status `gentle-ai.sdd-status@2` for `read-only-scan-foundation`: native `nextRecommended: apply`, repository-local workspace `/Users/osdy/Documents/GitHub/OsdyCleaner`, that same allowed edit root, strict TDD enabled, and no action-context warnings.
- This was a documentation-only reconciliation. No product source, tests, generated file, commit, PR, archive, or delivery artifact was changed.
- The review-workload gate is resolved by the parent-supplied standing `stacked-to-main` path; this reconciliation is not a new authored work-unit or PR boundary.

### Evidence inspected

- Read the proposal, canonical spec at `specs/read-only-scan/spec.md`, design, tasks, prior cumulative progress, and `openspec/config.yaml`.
- Confirmed current implementation/test surfaces for scan cancellation, report renderers, TUI, CLI, and process entry point. CodeGraph exploration confirmed the CLI production path is `NewCommand` → `run` and located the report and walker boundaries.
- Re-ran focused evidence: the WU6B scan selector, WU7 report selector, full TUI package, full CLI package, and `cmd/osdy` package all passed.
- Re-ran `go test -race ./internal/scan -count=1`, `go test ./... -count=1`, `go vet ./...`, and a repository-wide `gofmt -d` check; all passed.

### Completed task checkbox updates

The following implementation-owned tasks are now visibly `[x]` in the persisted `tasks.md`, supported by current source/test coverage and the final-state evidence:

- WU6B RED, GREEN, TRIANGULATE, REFACTOR — channel-synchronized cancellation/status/shutdown coverage and the exact focused selector passed.
- WU7 RED, GREEN, TRIANGULATE, REFACTOR — JSON/text report, deterministic output, goldens, and equivalence coverage passed.
- WU8 RED, GREEN, TRIANGULATE, REFACTOR — snapshot-only read-only TUI interaction, outcome, language, and terminal-validation coverage passed.
- WU9 RED, GREEN, TRIANGULATE, REFACTOR — injected command, stream, validation, exit-class, cancellation-precedence, and fixture-JSON coverage passed.

The persisted task ledger is now **155 checked / 15 unchecked / 170 total**. All 16 completed tasks listed above were re-read after the checkbox update and are visibly `[x]`.

### WU10 runtime reconciliation blocker

WU10 remains unchecked. Although its package test (`TestRunStopsSignalsAndPassesContext`) and the supplied final-state evidence pass, the exact disposable runtime script does not currently reproduce under this executor's default `mktemp -d` location:

- `go build -o "$tmp/osdy" ./cmd/osdy` succeeded.
- With the script-created fixture home, `HOME="$home" "$tmp/osdy" scan --format json` exited **1**, emitted no JSON, and wrote `osdy: inspect home: inspect home: scan path is a symbolic link` to stderr.
- The default temporary path is rooted at `/var/folders/...`; the current descriptor home inspection rejects that path because `/var` is a system symbolic link. The before/after fixture metadata remained identical.

This is a reproducible conflict with the supplied final-state runtime claim, so no WU10 checkbox was inferred from stale or non-reproducible evidence. No source correction is authorized by this reconciliation scope.

### TDD Cycle Evidence

| Scope | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- |
| Reconciliation only | No new production task or test was authored. | Current focused/package evidence passed for WU6B–WU9. | Full/race/static/format verification passed. | No source refactor was performed. |

### Remaining unchecked rows

- [ ] **WU5C2B HISTORICAL RED — DO NOT EXECUTE:** The combined RED would have mixed core DFS/FD/lifecycle seams with real rename/replacement fixtures and was blocked before any test write because the honest 550–810-line forecast exceeds 400; it is superseded by B1 then B2 and cannot be accepted. <!-- sdd-owner: implementation -->
- [ ] **WU5C2B HISTORICAL GREEN — DO NOT EXECUTE:** No combined production implementation exists; all Go/test files remained byte-identical to manifest `e178b1cc4f02cf5d7c5727cb29500115800c0c98d4e2dca546e3f585e44c4ddf`, and only independently accepted B1 may establish core DFS/lifecycle behavior. <!-- sdd-owner: implementation -->
- [ ] **WU5C2B HISTORICAL TRIANGULATE — DO NOT EXECUTE:** No combined real fixture or adversarial test was authored; independent B2 owns those fixtures only after accepted B1 and must not fabricate RED when B1 production already satisfies them. <!-- sdd-owner: implementation -->
- [ ] **WU5C2B HISTORICAL REFACTOR — DO NOT EXECUTE:** No combined verification, acceptance, rollback, commit, or review is permitted; retain this row as blocked provenance and use the separately bounded B1/B2 rollback boundaries below. <!-- sdd-owner: implementation -->
- [ ] **WU5E1 RED:** Add synchronized scanner and Darwin tests for fixed worker/job/result counts, bounded capacities and backpressure, validated relative identity-chain jobs, trusted-root private-anchor reopen, rolling ancestor descriptors, exact no-follow flags, ancestor/final identity and device/local checks, enumeration-versus-open kind/identity/disappearance, regular-only logical/allocation facts, no content reads, and exact close counts; run both focused commands and capture RED caused only by the missing file pipeline/reopen APIs. <!-- sdd-owner: implementation -->
- [ ] **WU5E1 GREEN:** Implement the minimum fixed worker and bounded channel pipeline plus platform relative reopen: stop absolute descendant inspection, reopen each component from the private root anchor, close rolling descriptors, accept only unchanged same-boundary regular final-descriptor facts, and return changed/boundary/error results with zero bytes otherwise; rerun both focused commands and expect PASS. <!-- sdd-owner: implementation -->
- [ ] **WU5E1 TRIANGULATE:** Add blocked enqueue/result cases, ancestor/final rename or symlink replacement, disappearance, non-regular entries, different-device/non-local facts, metadata/close failures, deep identity chains, and queue-pressure permutations; prove bounded goroutines/FDs, no target/content facts, no dropped accepted result, and deterministic zero contribution for changed entries; rerun the focused commands and `go test -race ./internal/scan -run '^(TestScannerBoundedFileJobs|TestScannerFileBackpressure|TestScannerRegularOnlyFacts|TestScannerFileEntryChanges)$' -count=1`, expecting PASS. <!-- sdd-owner: implementation -->
- [ ] **WU5E1 REFACTOR:** Run `gofmt -w internal/scan/descriptor.go internal/scan/scanner.go internal/scan/scanner_test.go internal/platform/macos/walker_darwin.go internal/platform/macos/walker_darwin_test.go && go test ./internal/scan -run '^(TestScannerBoundedFileJobs|TestScannerFileBackpressure|TestScannerRegularOnlyFacts|TestScannerFileEntryChanges)$' -count=1 && go test ./internal/platform/macos -run '^(TestWalkerRelativeFileReopen|TestWalkerFileIdentityChain|TestWalkerRegularFileFacts|TestWalkerFileReopenLifecycle)$' -count=1 && go test -race ./internal/scan -count=1 && go test ./... -count=1 && go vet ./...`; expect PASS and record exact predecessor/final hashes, <=400 accounting, runtime N/A, and rollback before WU5E2. <!-- sdd-owner: implementation -->
- [ ] **WU10 RED:** Before creating `cmd/osdy/main.go`, run `go test ./cmd/osdy -count=1`; record the expected non-zero “directory/package not found” result as the missing production entry-point boundary. <!-- sdd-owner: implementation -->
- [ ] **WU10 GREEN:** Implement `cmd/osdy/main.go` with `signal.NotifyContext`, stop signal delivery, production dependency construction, delegation to `internal/cli`, and `os.Exit` only; run `gofmt -w cmd/osdy/main.go && go test ./cmd/osdy ./internal/cli -count=1` and expect PASS with no scan/report logic in `main`. <!-- sdd-owner: implementation -->
- [ ] **WU10 TRIANGULATE:** Run the disposable runtime script below and expect exit 0, empty stderr, exactly one parseable JSON document, schema 1, complete outcome, canonical five area IDs, no absolute temporary-home path in JSON, and identical before/after fixture metadata. <!-- sdd-owner: implementation -->
- [ ] **WU10 REFACTOR:** Run the final acceptance sequence below in order: explicit-file `gofmt`, all focused package tests, report JSON parseability/equivalence tests without `-update`, `go test ./...`, `go vet ./...`, then the disposable runtime script again; expect every command to pass/no-diagnostic and stop rather than weakening checks if any result fails. <!-- sdd-owner: implementation -->
- [ ] Before resumed apply, record the already selected sequential `stacked-to-main` chain strategy and the WU2A → WU2B split, confirm no further product/delivery choice is pending, retain the previously selected module import identity, and preserve one-writer execution without assuming commits or PRs can be created. <!-- sdd-owner: parent -->
- [ ] After each applied work unit, inspect its apply-progress receipt, confirm authored additions plus deletions are at most 400 with generated `go.sum` excluded, confirm the focused command/result, runtime evidence or explicit N/A, complete changed-file identity, and rollback boundary, then authorize the next dependent unit through SDD apply/verify authority; if a unit exceeds 400 authored lines, stop and split it again before further implementation writes. <!-- sdd-owner: parent -->
- [ ] After WU10, record final SDD completion evidence against the normative coverage matrix, required commands, disposable runtime script, authored-line accounting, complete changed-file identity, and rollback boundaries; mark implementation complete only when that evidence is satisfied. <!-- sdd-owner: parent -->

### Deviation and next action

- Deviation from the supplied final-state evidence: WU10's exact runtime fixture did not reproduce due to `/var` symlink rejection; no behavior change was made.
- No implementation-owned task remains eligible for completion without resolving or explicitly adjudicating the WU10 runtime blocker. Historical WU5C2B/WU5E1 rows are intentionally non-executing and remain unchecked; parent-owned rows are informational and were not altered.
- Next recommendation: resolve the WU10 runtime discrepancy under an authorized implementation or verification decision, then run `sdd-verify`.

---

## WU10 canonical-worktree runtime correction and final acceptance (2026-09-04)

### Scope and status

- Consumed the parent-supplied authoritative status `gentle-ai.sdd-status@2`: change `read-only-scan-foundation`, OpenSpec artifact store, `applyState: ready`, `nextRecommended: apply`, strict TDD enabled, repository workspace and allowed edit root `/Users/osdy/Documents/GitHub/OsdyCleaner`, and no action-context warnings.
- The workload gate is resolved for this corrective work by the parent-selected `feature-branch-chain` and explicit acceptance of unavoidable cohesive-slice size exceptions. No commit, push, PR, archive, product-source edit, or test-source edit was performed.
- Corrected only the normative WU10 disposable script in `tasks.md`. It now derives `repo_root="$(pwd -P)"`, rejects a non-`/Users/*` worktree, creates `tmp` with `mktemp -d "$repo_root/.osdy-runtime.XXXXXXXX"`, retains the existing cleanup trap, and passes both `home` and `tmp` to Python. The JSON assertion now proves neither absolute path occurs in the report.

### Runtime and final-acceptance evidence

- WU10 runtime TRIANGULATE: extracted the normative script after the correction and ran it from the canonical repository worktree. It built before setting `HOME`, scanned only the disposable home, exited 0, left stderr empty, parsed exactly one newline-terminated JSON document, asserted schema version 1/complete outcome/the canonical five area IDs, asserted both the absolute fixture home and absolute temporary root were absent from JSON, and compared identical before/after fixture metadata.
- WU10 REFACTOR/final acceptance: verified the listed explicit Go files had empty `gofmt -d` output before the required `gofmt -w`; then all required focused package commands, the report parseability/equivalence selector, `go test ./... -count=1`, and `go vet ./...` passed. The corrected runtime script passed again after that sequence.
- Exact passed commands: `go test ./internal/core -count=1`; `go test ./internal/scan -count=1`; `go test ./internal/platform/macos -count=1`; `go test ./internal/report -count=1`; `go test ./internal/tui -count=1`; `go test ./internal/cli -count=1`; `go test ./cmd/osdy -count=1`; `go test ./internal/report -run '^(TestJSONReport|TestReportDeterminism|TestReportEquivalence)$' -count=1`; `go test ./... -count=1`; `go vet ./...`; and the corrected WU10 runtime script twice.

### TDD Cycle Evidence

| Task | Test layer | Safety net | RED | GREEN | TRIANGULATE | REFACTOR |
| --- | --- | --- | --- | --- | --- | --- |
| WU10 test-harness correction | Disposable process/runtime acceptance | Existing focused/package suites were run in final acceptance and passed. | Not fabricated: no production behavior was changed. | Not fabricated: no production behavior was changed. | Corrected runtime script passed with a canonical `/Users/...` temporary root and absolute-path non-leak assertions. | Format, focused packages, report selector, full suite, vet, and a second corrected runtime pass all passed. |

### Task and parent-row reconciliation

- Persisted-task readback confirms WU10 TRIANGULATE and WU10 REFACTOR are visibly `[x]`, based on this observed runtime and final-acceptance evidence.
- WU10 RED and GREEN remain visibly `[ ]`: this corrective rerun did not and cannot honestly recreate their historical before-entry-point RED state or claim a production implementation step, so they were not inferred from current source.
- Parent-owned rows remain `[ ]`. The first row expressly records `stacked-to-main`, while the user-selected path is `feature-branch-chain`; the later parent rows require parent-level receipt/coverage and full implementation-completion adjudication beyond this executor's observed corrective evidence. Historical failed implementation rows remain unchanged and unchecked.
- Current ledger after this update: 157 checked and 13 unchecked checkbox rows. `tasks.md` SHA-256 after the checkbox update: `529b813416fa642dc16dabde532bfcfc84b1ce72269d1d8c02044e7af564292a`.

### Files, deviations, remaining work, and boundary

- Changed artifacts: `openspec/changes/read-only-scan-foundation/tasks.md` and this cumulative `apply-progress.md` receipt only. The existing product worktree contained unrelated/staged source state before this corrective run; this executor made no source changes.
- Deviation from the prior failing harness: default macOS `mktemp -d` selected a `/var/...` alias, which correctly failed descriptor-only home validation. The canonical in-worktree disposable directory eliminates that harness-only alias without weakening scanner behavior or the no-real-home/no-mutation contract.
- Remaining implementation rows include the exact unchecked WU10 lines: `- [ ] **WU10 RED:** Before creating \`cmd/osdy/main.go\`, run \`go test ./cmd/osdy -count=1\`; record the expected non-zero “directory/package not found” result as the missing production entry-point boundary. <!-- sdd-owner: implementation -->` and `- [ ] **WU10 GREEN:** Implement \`cmd/osdy/main.go\` with \`signal.NotifyContext\`, stop signal delivery, production dependency construction, delegation to \`internal/cli\`, and \`os.Exit\` only; run \`gofmt -w cmd/osdy/main.go && go test ./cmd/osdy ./internal/cli -count=1\` and expect PASS with no scan/report logic in \`main\`. <!-- sdd-owner: implementation -->`.
- Workload/PR boundary: corrective WU10 runtime-harness receipt only; zero product authored lines. Delivery context remains `feature-branch-chain`; no lifecycle artifact was created.

---

## Parent final ledger normalization (2026-09-20)

- Recorded the user's final delivery choice as `feature-branch-chain`, with `size:exception` accepted only for cohesive bootstrap slices that cannot honestly fit the 400-line budget.
- Converted eight superseded, explicitly nonexecuting historical WU5C2B/WU5E1 rows from task checkboxes into provenance bullets. Their text and failure history remain preserved, but they no longer falsely count as pending implementation.
- Marked WU10 RED and GREEN complete from the genuine ODD-5A RED/GREEN record: process/production APIs were absent first, then `cmd/osdy` and production composition passed focused, race, full, vet, tidy, and format checks.
- Marked all three parent-owned completion rows complete: every executed unit has bounded receipt evidence; oversized combined units remained nonexecuting and were split; final WU10 runtime, coverage, accounting, identity, and rollback evidence is recorded.
- The corrected canonical-worktree WU10 runtime passed twice with no real-home access, no mutation, empty stderr, one schema-v1 JSON document, complete five-root output, and no temporary absolute-path leak.
- Product source was unchanged by this ledger normalization. Next authority must come from a fresh native `gentle-ai.sdd-status@2` projection.
