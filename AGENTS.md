# OsdyCleaner Agent Contract

Build a safe, personal macOS terminal cleaner. Keep changes small, reviewable, and in English.

## Read first

Use this source-of-truth order:

1. User instructions
2. Active OpenSpec change
3. This file
4. [Architecture](docs/architecture.md)

SDD is automatic: maintain the applicable OpenSpec artifacts and save significant verified project decisions to Engram. Ask on risk. The session review preference is 800 changed lines, but **more than 400 authored changed lines requires review-load handling**: ask or split before proceeding. Generated technical artifacts are English unless explicitly requested otherwise.

## Phase 1: hard boundary

Phase 1 is read-only discovery and reporting only. Do **not** add deletion, Trash, cleanup backends, shell adapters, `sudo`, privilege prompts, background daemons, telemetry, arbitrary paths, or external-volume scanning.

## Design and Go rules

| Area | Rules |
| --- | --- |
| Layout | One Go module; `cmd/osdy`; narrow `internal/...` packages. |
| Selected stack | Go 1.24+, Bubble Tea v2, Bubbles v2, Lip Gloss v2, Cobra, and `encoding/json`. |
| Dependencies | Prefer the standard library. Add a dependency only for a concrete need; do not build generic frameworks early. |
| Domain | Use named ID/value types, validated typed constants, and unexported constructors where invariants matter. |
| State | Snapshots and plans are immutable by convention. Keep mutable state out of package globals. |
| Errors | Wrap context with `%w`; inspect with `errors.Is`/`errors.As`; honor cancellation; never panic for expected failure. |
| Data | Use deterministic, explicitly versioned `encoding/json` schemas. |
| Style | Run `gofmt` and `goimports` when configured/available. |

## Filesystem and concurrency

- On supported Darwin, discover through a narrow internal descriptor-relative walker rooted at trusted `/`: component-wise no-follow `openat`, enumeration from opened directory descriptors, child opens relative to their parent, and explicit device/local checks. Do not follow symlink targets or cross mounts. On non-Darwin, return a typed unsupported error.
- Use bounded worker pools, bounded channels, and backpressure—never one goroutine per file.
- Aggregate results deterministically.
- Reserve `os.Root`/`os.OpenInRoot` for later root-relative mutation; they do not replace mount checks.
- Never use `os.RemoveAll` on an arbitrary absolute path.

## TUI and CLI

| Surface | Contract |
| --- | --- |
| Bubble Tea v2 | Separate Model, Update, and View. View is pure; I/O/scanning run in commands/effects; messages carry results; render state stays small. Test Update directly first; use full interactive tests only when necessary. Clean up the terminal on exit/error. |
| Cobra | Commands are thin `RunE` adapters; application/domain packages own behavior. Keep stdout stable and stderr for progress/diagnostics. JSON is one structured stdout document. No command accepts a cleanup path. |

## Strict TDD

For behavior changes: **RED → GREEN → TRIANGULATE → REFACTOR**. Record observed evidence.

- Prefer table-driven tests with `t.Run`; use `t.TempDir` for filesystem cases.
- Cover failure paths, fake system boundaries, and deterministic golden files; rerun goldens without update.
- Mark external or slow tests with `testing.Short`.
- Test Bubble Tea state through direct `Update` calls before interactive harnesses.
- Never inspect the real home directory or mutate personal files.

## Safety checklist

Before approving any later mutation, verify the architecture's [Safety invariants](docs/architecture.md#safety-invariants):

- [ ] Protected-root containment, exclusions, no symlink target, and no mount crossing.
- [ ] Identity-bound snapshot/plan; no loose user path; immediate revalidation.
- [ ] Root-relative operation only, Trash first, accurate observed outcomes and audit.
- [ ] No shell, root, sudo, escalation, or arbitrary executable/path input.

## Verification

**Bootstrap era (before `go.mod`):** no Go package exists to test. Do not treat failing `go test ./...` or `go vet ./...` as a product failure; they become active after module bootstrap. Verify changed Markdown by rereading it. Do not add a formatter/linter tool unless the repository configures it.

**Active module, in order:**

```sh
# 1. Format changed files
gofmt -w <changed-go-files>
goimports -w <changed-go-files> # only when configured/available
# 2. Focused package first
go test ./internal/<package>
# 3. Full suite
go test ./...
# 4. Static checks
go vet ./...
```

If formatting changes files after a check, rerun the affected tests and `go vet`. Report exact commands and results.

## Efficient work units

1. Search identifiers and paths first; read only the relevant symbol/range before whole files.
2. Read this file, active OpenSpec artifacts, and only relevant architecture sections.
3. Cite paths and summarize decisions; do not paste full files, logs, or dependency docs into prompts/handoffs.
4. Run targeted tests before broader tests; savings never justify skipping tests, safety checks, or required artifact readback.
5. Keep one bounded work unit and one writer in a worktree; never parallelize writes.
6. Reuse existing types/helpers; avoid duplicate abstractions. Comment non-obvious invariants, not code narration.
7. Return concise evidence: files changed, commands/results, risks, and next action.

## Review and completion

- Keep tests with behavior and docs with user-visible changes.
- Do not commit unless explicitly requested.
- At over 400 authored changed lines, apply the review-load rule above despite the 800-line preference.

### Definition of Done

- [ ] Active OpenSpec/Engram obligations and source order were followed.
- [ ] Phase boundary and safety invariants remain intact.
- [ ] Focused tests and required checks pass, or gaps are explicitly reported.
- [ ] Diff is one bounded, reviewable work unit with no unrelated changes.
- [ ] Handoff names files, evidence, risks, and next action.
