# Harden CLI cancellation and short writes

Issue: https://github.com/OsdyOrtiz/OsdyCleaner/issues/9

## Scope

Classify cancellation returned by either scan boundary as exit `130`, and detect incomplete report writes as exit `1`. Preserve finalized cancelled reports, existing writer errors, and the Phase 1 read-only boundary.

## Observed TDD evidence

Verified on 2026-11-26 with Go 1.24.2 on Linux/amd64:

1. **RED:** The unchanged baseline passed `go test ./...`. Five new CLI cases then failed: cancellation returned code `1` with a diagnostic in text, JSON, and TUI modes; short text/JSON writes returned code `0`.
2. **GREEN:** The same five cases passed after the two narrow checks in `internal/cli/run.go`; the test file was unchanged between these runs.
3. **TRIANGULATE:** Expanded to 32 passing cases covering wrapped cancellation, deadlines, zero/full writes, writer-error precedence, cancelled/partial snapshots, and renderer errors. The exact final test file was also run against the original source: 14 failures and 18 passing controls.
4. **REFACTOR:** Reviewed the branches and reused the existing `exitError`, `failure`, and `partialWriter` helpers. No further production refactoring was needed.

All original tests and module files remained unchanged. The two new scenarios in `openspec/specs/read-only-scan/spec.md` describe the corrected stream and exit behavior.

## Final verification

- `gofmt -l cmd internal`: no unformatted files.
- `go mod tidy -diff`: no differences.
- `go test ./internal/cli`: passed.
- `go test ./...`: passed for all seven Linux packages, using the normal Go test cache.
- `go test -race ./...`: passed for all seven Linux packages.
- `go vet ./...`: passed.
- `CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o <output> ./cmd/osdy-cleaner`: passed.
- `CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o <output> ./cmd/osdy-cleaner`: passed.

## Execution limits

Native macOS tests were not run in this Linux environment. The Darwin binaries were cross-built, not executed. CLI regression tests used injected scan and writer boundaries; no real home scan or interactive terminal session was run.

No Engram integration was available, so no Engram record was written. The verified decisions and TDD evidence are recorded here and in the canonical OpenSpec scenarios.
