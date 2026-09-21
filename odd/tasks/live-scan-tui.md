# Build live scan TUI

## Goal

Launch the read-only scan inside a responsive cyber-neon Bubble Tea interface, show honest category progress while scanning, and present the finalized snapshot in a clearer dashboard-plus-detail layout.

## Decisions

- Visual direction: cyber-neon, with a responsive OsdyCleaner ASCII wordmark.
- Progress: one honest step per fixed built-in category, plus active-category activity; never invent a per-file percentage.
- Results: summary dashboard, category list, and detail/warning pane.
- `osdy-cleaner scan` starts the TUI immediately when attached to a terminal; explicit text and JSON modes remain non-interactive and stable.
- Scan I/O runs only in Bubble Tea commands/effects. `Update` remains state-only and `View` remains pure.
- Final results continue to consume exactly one finalized schema-v1 snapshot; progress events are transient UI state, not snapshot facts.
- Phase 1 remains strictly read-only with no cleanup, selection, deletion, arbitrary paths, privilege, network, or external-tool behavior.

## Tasks

- [x] ODD-V1 — Add a bounded scanner progress contract for the five canonical categories, with strict TDD for ordering, cancellation, and no effect on finalized snapshots.
- [x] ODD-V2A — Add the Bubble Tea async scan lifecycle with bounded progress delivery, state-only updates, cancellation-safe shutdown, and terminal restoration.
- [x] ODD-V2B — Wire the interactive lifecycle through CLI production dependencies while preserving one scan, text/JSON output, diagnostics, and exit classes.
- [x] ODD-V3A — Implement the responsive cyber-neon theme, ASCII logo, and honest live progress screen with wide/narrow View tests.
- [x] ODD-V3B — Implement the responsive summary dashboard, category list, and detail/warning pane with direct Update/View tests.
- [x] ODD-V4A — Reconcile the canonical specification, architecture, and README with live read-only scanning, honest progress, dashboard behavior, and controls.
- [x] ODD-V4B — Run full verification, disposable PTY runtime proof, and independent review, then prepare a bounded delivery decision.

## Evidence

| Task | Status | Delivery / evidence |
| --- | --- | --- |
| ODD-V1 | complete | Strict TDD RED/GREEN recorded; focused progress tests, full tests/race, and vet passed. A follow-up RED proved terminal cancellation counters incorrectly included unvisited categories (`5/5`); GREEN now reports only attempted categories (`0/5` before start, `1/5` during the first root). Included in the single squashed feature commit authorized in issue #21. |
| ODD-V2A | complete | Strict TDD RED/GREEN recorded for direct Init/Update lifecycle behavior. A follow-up RED exposed a duplicate listener and repeated cancellation command; GREEN keeps exactly one result waiter through cancellation. Focused TUI/race and full tests/race/vet passed. Included in the single squashed feature commit authorized in issue #21. |
| ODD-V2B | complete | Strict TDD RED proved the terminal CLI pre-scanned before opening its viewer; GREEN routes terminal TUI through one `InteractiveScan`/`tui.RunScan` lifecycle while text/JSON retain `Scan`. Focused CLI/production, CLI race, full tests/race, vet, and diff checks passed. Included in the single squashed feature commit authorized in issue #21. |
| ODD-V3A | complete | Strict TDD RED recorded for direct wide/narrow View cases; GREEN adds a bounded cyber-neon read-only panel, ASCII/compact marks, 0/5–5/5 category bars, cancellation/failure language, and ANSI-safe width coverage. Focused TUI/race and full tests/race/vet/diff passed, including the tiny-terminal correction. Included in the single squashed feature commit authorized in issue #21. |
| ODD-V3B | complete | Strict TDD RED captured direct finalized View width-bound/dashboard and Tab-mode failures; GREEN adds the finalized-snapshot-only cyber-neon dashboard, responsive side-by-side/stacked panes, IEC byte display, visible text indicators, and two-mode Findings/Warnings toggle. Focused TUI/race and full tests/race/vet/module/diff checks passed, including the compact-category regression correction. Included in the single squashed feature commit authorized in issue #21. |
| ODD-V4A | complete | Markdown readback confirmed the canonical live-scan requirement/scenarios, README quick path and controls, and architecture runtime-flow invariants; stale pre-scan viewer claims and old controls were checked by grep, and `git diff --check` passed. Included in the single squashed feature commit authorized in issue #21. |
| ODD-V4B | complete | Independent full tests/race/vet/format/module/diff checks passed. A controlled 100x30 PTY run showed the read-only finalized dashboard and all five scanned categories, exited cleanly on `q`, and preserved fixture digest `a0c000e7d41fded15479ba589b8a9a077ab36442d5e4792b95a7fb99b2245307`; controlled JSON was one schema-v1 document and text was byte-deterministic. Native review remained unavailable (`schema-incompatible`, no lineage). |
