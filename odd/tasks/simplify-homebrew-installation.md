# Simplify Homebrew installation

## Goal

Publish the next OsdyCleaner release with `osdy-cleaner` as the only executable name and provide a one-command Homebrew tap installation.

## Decisions

- Primary installation: `brew install OsdyOrtiz/tap/osdy-cleaner`.
- Executable and Cobra command: `osdy-cleaner`.
- Compatibility: immediate rename; do not ship an `osdy` alias.
- Existing v0.1.0 release and tag remain immutable.
- Version: v0.2.0, because changing the command name is a breaking pre-1.0 change.
- Phase 1 remains strictly read-only.

## Tasks

- [x] ODD-H1 — Rename the Go entry point and CLI identity to `osdy-cleaner` with strict RED/GREEN evidence and no `osdy` compatibility alias.
- [x] ODD-H2 — Update CI, release packaging, architecture, and user documentation for `osdy-cleaner` and Homebrew-first installation.
- [ ] ODD-H3 — Verify, review, deliver, tag, and publish immutable v0.2.0 macOS arm64/amd64 assets and checksums.
- [ ] ODD-H4 — Create the public `OsdyOrtiz/homebrew-tap` repository, publish a checksummed `osdy-cleaner` formula for v0.2.0, and verify install plus execution through Homebrew.

## Evidence

| Task | Status | Commit / evidence |
| --- | --- | --- |
| ODD-H1 | complete | Corrected RED reproduction: in a detached disposable worktree at exact base `2d5f63d`, applying only the `command_test.go` delta and running `go test ./internal/cli -run 'TestCommandIdentityAndHelp|TestCommandFailureDiagnostics' -count=1` failed with `Use="osdy"` and `osdy:` diagnostics (exit 1). GREEN/triangulation: the same selector plus `TestCommandExitClasses` passed on the candidate; `go test ./...`, `go test -race ./...`, and `go vet ./...` passed. Source commit `85d324860c4336337eaa7931e370383bf91e23c3`. |
| ODD-H2 | complete; source commit `576387ef5209436a18e2ba11a2d2923cdfea6854` | RED: scoped assertions found the old README/workflow/architecture command, assets, and path. GREEN: both workflow YAML files parsed; all 14 embedded shell blocks passed `bash -n`; migration assertions, full tests, race tests, vet, and both cross-builds passed. A disposable release rehearsal produced both v0.2.0 archives, verified checksums, exact `osdy-cleaner`/README/LICENSE contents, arm64 and x86_64 Mach-O identities, and native `--help`; all disposable output was removed. |
| ODD-H3 | pending | — |
| ODD-H4 | pending | — |
