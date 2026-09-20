# Deliver a Safe Read-Only Scan Foundation

OsdyCleaner will gain its first useful product slice: a non-privileged scan of a fixed set of high-impact macOS developer/cache locations, with trustworthy text and JSON results and no mutation capability. The slice establishes the Go application and core scan domain while making partial filesystem visibility explicit instead of turning one unavailable location into a failed scan.

## Product decision at a glance

| Topic | First-slice decision |
| --- | --- |
| User outcome | A person can inspect where known developer caches consume space before deciding what, if anything, should happen later. |
| Coverage | Scan only built-in locations for npm, Homebrew, Gradle, Xcode DerivedData, and CoreSimulator. |
| Safety posture | Read-only, local, non-privileged, no symlink following, and no traversal outside the allowed root or its starting filesystem. |
| Results | Produce one explicitly versioned snapshot that can be rendered as deterministic text or as one structured JSON document. |
| Partial visibility | Missing, inaccessible, skipped, or boundary-limited roots are observations in the snapshot, not fatal global failures. |
| Risk communication | Preserve manual-review risk labels and estimate caveats; never describe a finding as safe to delete or as guaranteed reclaimable space. |
| TUI | Include a minimal read-only results viewer only if design confirms it can reuse the same snapshot without introducing a second workflow or action semantics. |

## User problem and current gap

The user currently has no product command that can answer a basic cleanup question: which known, high-impact developer/cache areas are present, how large do they appear, and where was the scanner unable to inspect? Broad disk coverage is a long-term goal, but beginning with whole-disk or arbitrary-path scanning would increase performance, privacy, and safety uncertainty before the core reporting contract exists.

The repository also has no Go module, executable, scan domain, built-in root definitions, or stable output schema. Consequently, later planning or cleanup work has no trustworthy snapshot foundation to build on and remains blocked.

## Goals

- Establish one Go module and the minimum core domain needed to represent built-in scan rules, findings, root observations, warnings, estimates, and complete or cancelled snapshots.
- Scan a bounded, built-in set of high-impact developer/cache roots without accepting arbitrary filesystem paths.
- Give users useful root summaries and matched findings with path, reason, risk/manual-review status, estimate basis, and relevant warnings.
- Make text and JSON views stable, equivalent, explicitly schema-versioned, and suitable for fixture-based regression checks.
- Continue across isolated missing, inaccessible, symlink, mount-boundary, or inspection failures and report each observed condition honestly.
- Leave a coherent read-only foundation for later specifications without introducing any cleanup mechanism.

## Non-goals

This change does not provide deletion, permanent mutation, Trash integration, cleanup plans or execution, selection for cleanup, confirmation flows, shell commands, tool adapters, `sudo`, privilege prompts, telemetry, a background daemon, arbitrary input paths, external-volume traversal, duplicate detection, or broad whole-disk discovery.

It also does not promise that reported bytes can be reclaimed, infer that age or size makes an item safe to remove, or interpret product-specific cleanup semantics for npm, Homebrew, Gradle, Xcode, or CoreSimulator.

## First-slice scope

### Required foundation

1. Bootstrap the repository as the selected single Go application only during implementation, after the remaining SDD artifacts approve that work.
2. Define the read-only domain vocabulary and immutable-by-convention snapshot contract needed by both human and machine-readable output.
3. Resolve only built-in roots under the current user's local macOS environment and scan within those boundaries.
4. Aggregate observations and findings into a stable order independent of filesystem enumeration or concurrent completion order.
5. Expose a scriptable scan result in text and JSON, with diagnostics kept separate from structured stdout.

### Built-in coverage

| Built-in area | First-slice treatment |
| --- | --- |
| npm cache | Measure and explain findings within the known npm cache root; do not invoke npm. |
| Homebrew cache | Measure and explain findings within Homebrew's known cache location; do not invoke `brew`. |
| Gradle caches | Measure and explain findings within the known Gradle cache root; do not invoke Gradle. |
| Xcode DerivedData | Report bounded DerivedData findings and estimates without declaring them cleanup-eligible. |
| CoreSimulator | Report the known root's footprint and warnings with explicit manual-review risk; do not infer simulator-device cleanup actions. |

All five areas belong to the initial coverage target. A location that does not exist is still represented as a `missing`-equivalent observation so that “not found” is distinguishable from “not scanned.” A location that cannot be read is represented as inaccessible with useful context while the remaining locations continue.

### Coherence-gated Bubble Tea viewer

A minimal Bubble Tea experience is in scope only if the design demonstrates that it:

- consumes the same completed or cancelled snapshot as text and JSON output;
- offers only summary, results, warnings/details, navigation, and quit behavior;
- exposes no cleanup selection, action, confirmation, or mutation affordance; and
- does not duplicate scan rules or create a separate result model.

If those conditions cannot be met in one bounded vertical slice, the viewer is explicitly deferred and the text/JSON scan remains the complete first product outcome.

## User-visible behavior

1. The user starts a scan without supplying a path or granting additional privileges.
2. The scanner checks the five built-in areas and stays within each allowed local filesystem boundary.
3. The result shows which roots were scanned, absent, inaccessible, skipped, or only partially observable.
4. Findings show a display path, why the built-in rule reported them, their manual-review risk, measured estimate basis, and warnings. APFS, hard-link, permission, or boundary limitations remain visible rather than being presented as exact reclaimable bytes.
5. Text output is concise and consistently ordered. JSON output is one schema-versioned document carrying the equivalent snapshot facts in a stable field and collection order.
6. Cancellation produces an explicitly incomplete/cancelled snapshot when a trustworthy partial result can be finalized; it never starts another operation.
7. A global failure is reserved for cases where scanning cannot initialize or cannot produce a trustworthy snapshot, not for an individual root-level observation.

For identical fixture contents, rule versions, scan options, and fixed run metadata, serialized output must be repeatable. Runtime identity or time metadata, if retained, must be explicit and must not allow filesystem or concurrency ordering to leak into the result.

## Acceptance outcomes

- [ ] A supported macOS user can run one read-only scan covering npm, Homebrew, Gradle, Xcode DerivedData, and CoreSimulator without network, shell, privilege escalation, or filesystem mutation.
- [ ] The scan accepts no arbitrary path and does not follow symlinks or cross the starting mount/device boundary by default.
- [ ] Every built-in root has a recorded observation, including missing and inaccessible roots, while isolated failures do not prevent other roots from being reported.
- [ ] Findings preserve manual-review risk labels, explanations, estimate bases, and warnings; no output claims guaranteed reclaimed space or deletion safety.
- [ ] Text and JSON represent the same explicitly versioned snapshot and are deterministic under fixed inputs and run metadata.
- [ ] Aggregate accounting avoids double-counting the same observed filesystem identity within a scan and labels APFS-related estimate uncertainty.
- [ ] Cancellation and partial visibility are explicit snapshot states and cannot be mistaken for a complete, warning-free scan.
- [ ] Fixture-based domain and scan evidence covers normal roots, empty roots, missing roots, inaccessible entries, symlinks, boundary skips, repeated identities, cancellation, and stable output without inspecting the developer's real home directory.
- [ ] If included, the Bubble Tea viewer is strictly read-only, presents the shared snapshot, and has no action or selection language.

## Affected areas and implications

| Area | Implication |
| --- | --- |
| User workflow | Users gain evidence before action, but must still manage any cleanup manually outside OsdyCleaner. |
| Product language | Findings, observations, warnings, estimates, and manual-review risk become user-facing concepts that later phases must preserve. |
| Data contract | The first JSON snapshot version becomes a compatibility surface; later changes require explicit schema evolution. |
| Filesystem behavior | macOS permissions, APFS accounting, hard links, symlinks, and mount boundaries must be represented honestly rather than hidden. |
| CLI and optional TUI | Every presentation must consume the same snapshot facts so scriptable and interactive views cannot disagree. |
| Support and privacy | Paths and local storage metadata stay on the user's machine; there is no telemetry or network transfer, but users remain responsible for redirected output files. |
| Future cleanup phases | Plans and mutation must depend on a later, separately approved identity-bound contract; this proposal does not make current findings actionable. |

## Risks and tradeoffs

| Risk or tradeoff | Mitigation in this slice |
| --- | --- |
| Estimates can overstate reclaimable APFS space. | Label the measurement basis and uncertainty; deduplicate repeated filesystem identities where observed; never call an estimate “freed space.” |
| Permission differences can make scans look inconsistent. | Emit per-root and per-entry observations with context and continue safely. |
| CoreSimulator has domain-specific lifecycle semantics. | Treat it as read-only footprint evidence with manual-review risk and defer action interpretation. |
| Broad coverage pressure can expand the first slice into a whole-disk scanner. | Keep the root set built-in and fixed; defer arbitrary paths, external volumes, and additional categories. |
| A TUI can add a second state model and dilute the foundation. | Apply the coherence gate and defer it before weakening snapshot or CLI quality. |
| Stable output can be undermined by traversal order or runtime metadata. | Require canonical aggregation and make run-varying metadata explicit and controllable in verification. |
| Users may mistake a high-impact finding for a recommendation to delete. | Use manual-review labels and explanatory, non-actionable language throughout every view. |

## Rollback

Because this slice is read-only and creates no cleanup state, rollback requires no filesystem recovery or data migration. If the implemented scan proves unreliable, remove or disable the scan entry point and its optional viewer, revert the module/domain/output additions, and stop publishing the affected snapshot schema version. Previously redirected reports remain inert local documents and must not be accepted later as cleanup authorization.

## Explicit deferred work

- Cleanup plan creation, candidate selection for action, identity-bound execution, and all audit semantics for mutation.
- Trash, permanent deletion, recovery claims, confirmation phrases, and post-action results.
- Shell or tool adapters, including npm, Homebrew, Docker, Gradle, Xcode, and CoreSimulator commands.
- `sudo`, privilege prompts, privileged helpers, automatic/background operation, telemetry, and network services.
- Arbitrary or user-configured roots, whole-disk scanning, external volumes, and mount traversal.
- Duplicate detection, personal-file analysis, app leftovers, generalized cache discovery, and additional root coverage.
- Product-specific CoreSimulator device management and cleanup eligibility.
- Rich TUI features such as live progress dashboards, filtering, selection, plans, confirmations, and actions.

## Proposal question round

These questions are intended to improve the PRD by exposing product rules, implications, edge cases, and scope tradeoffs. In automatic execution mode, the proposal adopts the assumptions below for review; the user may answer, skip, correct the framing, or request a second question round.

1. **Must all five named areas ship together in the first slice?** Assumption: yes; if scope pressure appears, the optional TUI is deferred before any named root.
2. **What result granularity is most useful initially: root totals or bounded artifact findings?** Assumption: show root observations plus grouped rule findings, not an exhaustive file-by-file inventory in the primary view.
3. **What should happen when every built-in root is missing or inaccessible?** Assumption: produce a valid empty or warning-bearing snapshot rather than a global failure, provided the scanner itself initialized and finalized correctly.
4. **How much CoreSimulator interpretation belongs in this slice?** Assumption: footprint and boundary observations only, with manual-review risk and no device lifecycle or cleanup recommendation.
5. **Should the first slice require an interactive view?** Assumption: only under the coherence gate above; deterministic text and JSON are the non-negotiable product baseline.

## Success criteria

The change succeeds when a user can obtain a trustworthy, repeatable inventory of the five named high-impact areas, understand both findings and visibility gaps, and see no language or affordance that implies OsdyCleaner can act on those findings. The resulting versioned snapshot must be strong enough for later design work while the implemented behavior remains entirely local, bounded, non-privileged, and read-only.
