# OsdyCleaner

OsdyCleaner is a macOS-only, Phase 1 read-only scanner for developer cache areas. It shows what may be reclaimable; it never deletes anything.

## Scan

```sh
osdy-cleaner scan [--format text|json|tui]
```

When both stdin and stdout are terminals, the default format is `tui`. Otherwise, the default is `text`.

Examples:

```sh
osdy-cleaner scan
osdy-cleaner scan --format text
osdy-cleaner scan --format json
```

OsdyCleaner scans only these fixed locations in the current user's home directory:

- `~/.npm` — npm cache
- `~/Library/Caches/Homebrew` — Homebrew cache
- `~/.gradle/caches` — Gradle caches
- `~/Library/Developer/Xcode/DerivedData` — Xcode DerivedData
- `~/Library/Developer/CoreSimulator` — CoreSimulator

There are no arbitrary path arguments. OsdyCleaner is read-only: it does not mutate or delete files, invoke `sudo` or a shell, send telemetry, access external volumes, or use the network.

## Install v0.2.0

Install with Homebrew:

```sh
brew install OsdyOrtiz/tap/osdy-cleaner
```

### Manual archive fallback

If you cannot use Homebrew, download the v0.2.0 archive for your Mac and verify it against `checksums.txt` from the same release:

| Mac | Archive |
| --- | --- |
| Apple silicon | `osdy-cleaner_0.2.0_darwin_arm64.tar.gz` |
| Intel | `osdy-cleaner_0.2.0_darwin_amd64.tar.gz` |

```sh
grep '  osdy-cleaner_0.2.0_darwin_arm64.tar.gz$' checksums.txt \
  | shasum -a 256 -c -
tar -xzf osdy-cleaner_0.2.0_darwin_arm64.tar.gz
mkdir -p ~/.local/bin
mv osdy-cleaner ~/.local/bin/osdy-cleaner
chmod 755 ~/.local/bin/osdy-cleaner
~/.local/bin/osdy-cleaner scan --format text
```

Replace `arm64` with `amd64` on an Intel Mac. Ensure `~/.local/bin` is on your `PATH` before using `osdy-cleaner` by name.

### Gatekeeper

v0.2.0 binaries are unsigned and not notarized. After verifying the release provenance and checksum, use macOS's per-app review flow: try opening the binary, then choose **Open Anyway** in **System Settings → Privacy & Security** if you trust it. Do not disable Gatekeeper globally or use broad quarantine-removal commands.

## Build from source

Go 1.24.2 or newer is required.

```sh
git clone https://github.com/osdy/OsdyCleaner.git
cd OsdyCleaner
go build -o osdy-cleaner ./cmd/osdy-cleaner
./osdy-cleaner scan --format text
```

## Automation and output

`--format json` emits deterministic schema-v1 JSON. It includes the scan outcome, policy, estimates, roots, findings, warnings, and caveats; it is intended for scripts and CI.

```sh
osdy-cleaner scan --format json > osdy-cleaner-report.json
```

Reported logical and allocated sizes are estimates, not promises of reclaimable disk space. APFS clones, snapshots, compression, and hard links can make actual recovery lower or different. CoreSimulator data also requires product-aware manual review.

## Exit codes

| Code | Meaning |
| --- | --- |
| `0` | Scan completed |
| `1` | Operational failure |
| `2` | Invalid command input or output format |
| `3` | Scan completed with partial visibility |
| `4` | Unsupported platform |
| `130` | Scan cancelled |

## Verify

After installation or a source build:

```sh
osdy-cleaner scan --format text
osdy-cleaner scan --format json
osdy-cleaner scan --format tui
```

The TUI command requires terminal input and output. For machine-readable verification, inspect `schema_version` in the JSON output; v0.2.0 emits `1`.

## License

[MIT](LICENSE)
