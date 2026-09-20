# Prepare v0.1.0 release

## Goal

Publish a documented, continuously verified OsdyCleaner v0.1.0 with unsigned macOS arm64 and amd64 binaries.

## Constraints

- Phase 1 remains strictly read-only.
- License: MIT.
- Release assets: macOS arm64 and amd64 archives plus SHA-256 checksums.
- Binaries are not signed or notarized; documentation and release notes must say so.
- Keep each work unit reviewable and commit it with its verification evidence.

## Tasks

- [x] ODD-R1 — Add MIT license and a concise README covering installation, actual CLI usage, safety boundary, formats, exit codes, and unsigned-binary expectations.
- [x] ODD-R2 — Add macOS CI and tag-driven GitHub Release automation for tested arm64/amd64 archives and checksums.
- [x] ODD-R3 — Create evidence-backed follow-up issues for non-blocking native-review advisories, after duplicate and privacy checks.
- [ ] ODD-R4 — Verify the complete release candidate, obtain native review, merge the release PR, create annotated tag v0.1.0 from fetched origin/main, publish the GitHub Release, and verify assets.
  - [x] Merge verified release candidate PR #13 and push annotated tag `v0.1.0` at `bc0b30bafc234ee1a469d3a83ff428d7ae643a84`.
  - [ ] Recover failed release run #35512079963 by validating the annotated tag through the GitHub API and dispatching the trusted workflow from `main`.
  - [ ] Verify release assets and close issue #12.

## Evidence

| Task | Status | Commit / evidence |
| --- | --- | --- |
| ODD-R1 | complete | `e393d872ef052dddf9d81bc59b0b6c6282005e06`; Markdown readback and `git diff --check` passed. |
| ODD-R2 | complete | `f93ee0a34e089558434a7ec99ca7db15e8db21d9`; YAML/LSP validation, full Go checks, and local dual-architecture package rehearsal passed. |
| ODD-R3 | complete | Issues `#8`, `#9`, `#10`, and `#11`; duplicate searches, exact-body privacy scan, and target readback passed. |
| ODD-R4 | in progress | PR `#13` merged; tag `v0.1.0` pushed; run `#35512079963` failed because checkout exposed the peeled commit rather than a local annotated-tag object. |
