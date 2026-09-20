# Complete Read-Only Scan Foundation

## Objective

Finish the Phase 1 read-only product path from scanner cancellation semantics through deterministic reports, a read-only TUI, CLI orchestration, and final process/runtime acceptance.

## Problem and Why

The descriptor-relative scanner and bounded policy foundation exist, but the product is not yet usable end to end. Cancellation/shutdown semantics remain incomplete, and there is no stable report, terminal viewer, CLI orchestration, or process entry point.

## Migration Context

This feature continues the intent and constraints recorded in `openspec/changes/read-only-scan-foundation/` while moving ongoing execution tracking to ODD. Existing OpenSpec artifacts remain preserved as read-only requirements and provenance; they are not rewritten or treated as completed by this migration.

## Scope

- Finalize cancellation, root-status precedence, result draining, worker joining, and deterministic completion.
- Add deterministic schema-v1 JSON and equivalent text reports.
- Add a minimal Bubble Tea v2 read-only viewer.
- Add Cobra-based scan orchestration, stable stdout/stderr behavior, and exit classes.
- Wire the process entry point, signal cancellation, and disposable runtime acceptance.

## Non-Goals

- No cleanup, deletion, Trash, mutation, arbitrary paths, shell adapters, `sudo`, privilege prompts, telemetry, background daemons, or external-volume scanning.
- No commit, push, pull request, release, or OpenSpec archival without explicit user authorization.
- No unrelated refactor or dependency framework.

## Constraints

- Phase 1 remains strictly read-only.
- Darwin traversal remains descriptor-relative, no-follow, local-device bounded, and deterministic.
- Strict TDD is active from `openspec/config.yaml`: RED → GREEN → TRIANGULATE → REFACTOR.
- Exact project runner: `go test ./...`; format with `gofmt`; static checks with `go vet ./...`.
- Keep each implementation task at or below 400 authored changed lines; split before proceeding if a task exceeds that boundary.
- One writer at a time. Generated technical artifacts remain in English.
- Existing OpenSpec requirements and architecture safety invariants remain binding during migration.

## Delivery Forecast

Estimated total: approximately 1,250–1,570 authored changed lines across five sequential work units. Delivery must remain split into independently reviewable tasks of at most 400 authored changed lines. The user selected automatic execution and requested one delivery, but repository policy forbids a single 1,500-line work unit; no delivery action is authorized.

## Tasks

- [x] **ODD-1 — Cancellation, status precedence, and shutdown order**
  - Allowed implementation surfaces: `internal/scan/scanner.go`, `internal/scan/scanner_test.go`, `internal/scan/finalize.go`, `internal/scan/finalize_test.go`.
  - Add synchronized tests for cancellation before/during/after roots, accepted-result draining, later-root skipping, shutdown order, completion determinism, and no leaks.
  - Implement exact precedence: `cancelled > partial > boundary_limited > scanned`, retaining lower-precedence warnings.
  - Checks: focused scanner tests, `go test -race ./internal/scan -count=1`, `go test ./...`, `go vet ./...`, and gofmt readback.

- [x] **ODD-2A1 — Deterministic schema-v1 JSON projection**
  - Add the ordered JSON DTO/projection and non-golden behavioral tests for complete, partial, and cancelled snapshots.
  - Prove byte determinism, parseability, explicit unknowns, stable ordering, and honest caveats.
  - Checks: focused JSON behavioral tests, `go test ./...`, `go vet ./...`, and gofmt readback.

- [x] **ODD-2A2 — JSON golden fixtures**
  - Add and inspect complete/partial/cancelled JSON goldens against the accepted ODD-2A1 projection.
  - Prove exact schema bytes and rerun without golden updates.
  - Checks: focused golden tests, `go test ./...`, `go vet ./...`, and unchanged production readback.

- [x] **ODD-2B1 — Canonical text reports and goldens**
  - Add canonical text projection and complete/partial/cancelled text goldens on the accepted JSON fixtures.
  - Prove deterministic text, prominent incomplete states, all finalized facts, and no safety/reclaim promise.
  - Checks: focused text/golden tests without updates, `go test ./...`, `go vet ./...`, and gofmt readback.

- [x] **ODD-2B2 — Unambiguous cross-format equivalence proof**
  - Bind each text record and estimate to its JSON root/finding/warning with section membership and cardinality.
  - Add varied policy, per-record estimates/accounting, and non-null related-path evidence so duplicated substrings cannot mask omissions.
  - Prove exact caveat membership and every shared JSON fact without a second business model.
  - Checks: focused equivalence tests, report package, `go test ./...`, `go vet ./...`, and gofmt readback.

- [x] **ODD-3A — TUI state invariants and terminal lifecycle**
  - Correct paging/resize so selection remains visible; validate invalid snapshots and nil files without panic.
  - Own Bubble Tea alternate-screen lifecycle and retain pure Init/Update/View with no commands after quit.
  - Checks: focused model/runner tests, `go mod tidy`, `go test ./...`, `go vet ./...`, and gofmt readback.

- [x] **ODD-3A2 — Paging detail-context reset correction**
  - Reset detail scrolling whenever Page Up/Down changes the selected category, including boundary/no-change behavior.
  - Add the missing focused regression without expanding presentation scope.
  - Checks: focused context-reset tests, TUI package, `go test ./...`, `go vet ./...`, and gofmt readback.

- [x] **ODD-3B — Complete read-only snapshot presentation**
  - Present aggregate/root estimates, findings/risk/accounting, full warnings, and caveats with summary/category/details or warnings state.
  - Use the Bubbles v2 viewport where scrolling is needed and prove category/detail controls, fact agreement, and forbidden-affordance absence.
  - Checks: focused direct Update/View tests, `go mod tidy`, `go test ./...`, `go vet ./...`, and gofmt readback.

- [x] **ODD-4A — CLI orchestration and core process contract**
  - Add Cobra adapters for one scan, text/JSON selection, terminal validation, render-before-write, and exit classes.
  - Keep behavior in application/domain packages, not Cobra handlers; accept no cleanup/arbitrary path.
  - Checks: focused CLI behavior tests, disposable JSON fixture, `go test ./...`, `go vet ./...`, tidy/format readback.

- [x] **ODD-4B — CLI failure and cancellation stream proof**
  - Prove viewer failures map to exit 1, cancelled text/JSON remains parseable with empty stderr and exit 130, and cancellation outranks partial.
  - Prove scan/render/view/write failures leave stdout clean as applicable and emit the exact stderr diagnostic contract.
  - Checks: focused failure/stream tests, CLI package, `go test ./...`, `go vet ./...`, and gofmt readback.

- [x] **ODD-5A — Production and process wiring**
  - Wire the executable, process signals, production scanner/report/TUI composition, terminal lifecycle, and safe home inspection.
  - Keep process dependency construction outside Cobra and reject non-Darwin before home resolution.
  - Checks: focused/full/race tests, vet, tidy/format readback, and independent descriptor-lifecycle review.

- [x] **ODD-5B — Complete-scan evidence semantics**
  - Correct the runtime-discovered scanner evidence-state defect that marks roots partial even when regular-file evidence is fully accepted.
  - Preserve directory-only partial semantics; every walker/file rejection becomes partial except explicit boundary-limited cases; retain cancellation precedence.
  - Checks: focused evidence/gap/status tests, scan package/race, full tests, vet, gofmt, and independent false-complete-path review.

- [x] **ODD-5C — Darwin regular-file dual-identity evidence**
  - Open enumerated regular files relative to the retained parent, verify no-follow type/identity/device/local invariants, and emit matching opened identity only after validation.
  - Preserve rejection classification and exact descriptor closure for missing/replaced/symlink/cross-device/nonlocal/stat/close failures.
  - Checks: focused walker/scanner tests, race/full tests, vet, tidy/format checks, independent safety/FD review, and successful disposable JSON/text plus SIGINT runtime evidence.

- [x] **ODD-5D — Exact empty/symlink/cancellation semantics and final acceptance**
  - Make successfully traversed empty roots `scanned/completed`; an allowed skipped entry symlink warns without making the root partial.
  - Check cancellation before every Darwin directory read while preserving cleanup and cancellation precedence.
  - Run focused/race/full tests, binary build, disposable empty/symlink/JSON/text/SIGINT mutation checks, vet, tidy/format checks, and final independent requirements review.

## Acceptance Criteria

- Cancellation never starts new roots/jobs after observation, accepted work drains, workers join, and root outcomes are deterministic.
- JSON and text expose equivalent finalized snapshot facts without recomputation or nondeterministic fields.
- The TUI is useful but strictly read-only.
- CLI stdout/stderr and exit codes are stable and test-covered.
- The executable handles signals and completes a disposable end-to-end scan without mutating fixtures.
- All applicable focused, race, full-suite, static, formatting, and runtime checks pass, or gaps are explicitly recorded.

## Progress and Evidence

- 2026-09-19: User explicitly selected migration from the active SDD continuation to ODD.
- 2026-09-19: ODD tracker created before new source writes. Existing WU6A evidence remains accepted and unchanged.
- 2026-09-19: ODD-1 accepted at +119/-16 = 135 authored lines. Strict TDD evidence observed; focused cancellation tests, scan race suite, full suite, `go vet`, and formatting passed. Independent verification passed, and the parent reran the focused command successfully. LSP `waitgroup-done-scope` warnings were independently classified as false positives against safe mutex/wait ordering.
- 2026-09-19: ODD-2 was blocked before writes because six authored goldens plus both renderers exceed 400 lines. The user selected dependency-safe ODD-2A JSON first, followed by ODD-2B text/equivalence.
- 2026-09-19: ODD-2A was also blocked before writes because its three readable nested goldens consume nearly the full budget. The user selected ODD-2A1 projection/behavioral tests first and ODD-2A2 JSON goldens second.
- 2026-09-19: ODD-2A1 accepted at +235/-0 authored lines. Strict TDD, focused/package/full tests, `go vet`, formatting, independent verification, and the parent focused rerun all passed. Golden-byte comparison remains explicitly deferred to ODD-2A2.
- 2026-09-19: ODD-2A2 accepted at +47/-0 physical lines with three compact JSON goldens. RED on missing fixtures, update-only GREEN, non-update focused/package/full tests, `go vet`, and formatting passed. Independent functional verification passed but reported a provenance-only partial because Git could not prove an untracked production file unchanged; writer pre/post hashing and the parent current hash both matched `ba5656679770390bb7d75e98ddcb82d753a40bb9d4e9a297e352dbcb67791d27`, and edit confinement excluded `json.go`.
- 2026-09-19: ODD-2B initial implementation is +278/-0 with all commands passing, but independent verification found incomplete cross-format proof: policy, schema/complete, top-level estimate, root counts/warnings, finding warnings/accounting, deduplication completeness, and related paths were not all asserted.
- 2026-09-19: A +4/-0 correction added JSON-derived assertions, but re-verification rejected unanchored substring ownership, identical estimates, null-only related paths, repeated accounting, and loose caveat placement. The user approved a dependency-safe split: ODD-2B1 accepts the independently validated renderer/goldens at cumulative +282/-0; ODD-2B2 owns robust varied-fixture equivalence with a fresh ≤400-line budget.
- 2026-09-19: ODD-2B2 accepted at +178/-77 = 255 authored lines. Exact parsed section/order/cardinality comparisons, varied policy and per-record estimates/accounting, non-null related path, and exact caveats closed every prior gap. Focused/package/full tests, `go vet`, formatting, independent verification, and the parent focused rerun passed; renderers and all goldens remained unchanged.
- 2026-09-19: The initial ODD-3 candidate added 302 authored lines and all commands passed, but independent verification blocked acceptance: paging/resize can hide selection, shared snapshot facts are omitted, alternate-screen lifecycle is absent, navigation differs from the approved design, and invalid snapshot/nil file inputs can panic. The user approved fresh bounded ODD-3A state/lifecycle and ODD-3B full-presentation correction units; ODD-3 remains incomplete.
- 2026-09-19: ODD-3A initial correction added 109 physical lines and all commands passed. Re-verification confirmed selection/invalid-input/alternate-screen fixes but found `detailOffset` inert and unbounded; Bubble Tea v2 local source confirms `tea.View.AltScreen` owns entry and cleanup semantics.
- 2026-09-19: Visible bounded detail scrolling and integer-safe sizes brought cumulative ODD-3A to an estimated 392 physical lines. Final verification found one remaining defect: Page Up/Down category changes did not reset `detailOffset`, so the missing fix/test was isolated as fresh ODD-3A2.
- 2026-09-19: ODD-3A2 completed at +41/-0 lines. Paging now resets detail context only when the clamped category actually changes and preserves valid boundary state. Final independent verification accepted ODD-3A and ODD-3A2; focused/package/full tests, `go vet`, tidy/format diffs, and parent spot checks passed. Live PTY coverage remains intentionally deferred because Bubble Tea v2 lifecycle semantics and the deterministic program seam establish alternate-screen ownership.
- 2026-09-19: The first ODD-3B worker aborted after partial writes. Read-only incident diagnosis found tidy, formatted, valuable Bubbles v2/full-fact work but four regressions; restoration was unsafe because targets are untracked, so bounded in-place recovery was selected.
- 2026-09-19: Recovered ODD-3B accepted at an inferred 242 authored lines. Page Up/Down owns category paging; j/k/arrows own viewport scrolling; successful category/Tab changes reset viewport context; invalid snapshots stay out of detail states. Aggregate/root/finding/accounting/warning/caveat presentation and forbidden-affordance tests pass. Focused/package/full tests, `go vet`, tidy/format diffs, independent verification, and parent spot check passed.
- 2026-09-19: ODD-4 implementation/core tests completed at +347/-0 authored lines with Cobra v1.9.1. Focused/package/full tests, disposable JSON fixture, `go vet`, tidy and format checks passed; active LSP compile diagnostics are clean. Independent verification found behavior aligned but missing proof for viewer failure, cancelled stream purity/precedence, and failure stdout/stderr diagnostics. To preserve the 400-line boundary, accepted core work is ODD-4A and fresh test-only proof is ODD-4B.
- 2026-09-19: ODD-4B completed at +100/-0 test lines. Direct tests now prove viewer failure, cancelled text/JSON rendering before exit 130, one-document JSON, cancellation-over-partial precedence, exact scan/render/view/write/input/unsupported diagnostics, stdout partial-write policy, and silent valid partial/cancelled results. Focused/package/full tests, `go vet`, tidy/format checks, independent verification, and parent spot check passed. Complete ODD-4 is accepted.
- 2026-09-19: ODD-5A production/process wiring completed at +265/-0 authored lines. Signal cleanup, outside-Cobra composition, safe descriptor-relative home inspection, non-Darwin unsupported routing, focused/full/race tests, `go vet`, tidy/format checks, and independent FD-lifecycle review passed.
- 2026-09-19: First disposable-home run correctly rejected `/var` because macOS `/var` is a symlink; a fixture under the canonical `/Users` project path passed no-mutation and empty-stderr gates but exited 3 with `partial`. Investigation found `consumeDirectories` initializes `partial := true` and never clears it when all regular-file evidence is accepted. Split ODD-5B to correct and prove complete evidence semantics before final acceptance.
- 2026-09-19: ODD-5B accepted at 92/180 authored lines after two independent correction rounds. Accepted-only evidence is complete; absent evidence and every walker/file rejection are partial except explicit boundary-limited cases; cancellation remains highest precedence. Focused/package/race/full tests, vet, formatting, and final independent false-complete-path review passed.
- 2026-09-19: Disposable runtime still produced `entry_changed` with zero findings. Read-only incident diagnosis found Darwin walker regular-file facts contain enumerated identity but nil opened identity, so scanner correctly rejects every real regular file before FilePort inspection. Split ODD-5C for safe relative no-follow open/fstat/fstatfs dual-identity evidence and final runtime acceptance.
- 2026-09-19: ODD-5C accepted at +117/-16=133 authored lines. Independent review accepted relative no-follow dual-identity evidence and exact temporary-FD closure. Canonical disposable JSON/text scans exited 0 with five findings, one JSON document, empty stderr, and unchanged byte manifests. An lsof-handshaked SIGINT scan exited 130 with cancelled parseable JSON, empty stderr, and unchanged 200k-file metadata manifest.
- 2026-09-19: Final requirements review blocked completion on three exact OpenSpec mismatches: empty traversed roots incorrectly partial, allowed entry symlink warnings incorrectly partial, and Darwin `readDir` lacks a pre-read cancellation check. Split ODD-5D to restore source-of-truth semantics and final acceptance.
- 2026-09-19: ODD-5D completed and independently accepted. Empty traversed roots are complete, permitted entry symlinks warn without partial status, cancellation is checked before every Darwin directory read, and callback errors retain joined child-close evidence. Focused/package/race/full tests, vet, tidy/format checks, and active LSP error checks passed.
- 2026-09-19: Final disposable runtime acceptance passed after ODD-5D: canonical empty/symlink/regular JSON and text scans exited 0 with five scanned roots, one JSON document, empty stderr, and unchanged lstat/content/link manifest; lsof-handshaked SIGINT during a 100k-file scan exited 130 with cancelled parseable JSON, empty stderr, and unchanged stat manifest. Native Gentle review remained unavailable because the package-local v3.1.0 binary is missing; no lineage was created. Final independent rereview found no code/spec safety mismatch.
- Visible `todo` projection is unavailable in this runtime; this file and its Engram mirror are the durable progress sources.
- Commit evidence is intentionally absent because repository policy requires explicit commit authorization.

## Next Step

Phase 1 implementation and acceptance are complete. Await explicit user authorization before any commit, push, PR, OpenSpec archival, or later-phase mutation work.
