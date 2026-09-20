# Osdy

Osdy is a macOS-only, Phase 1 read-only scanner for developer cache areas. It shows what may be reclaimable; it never deletes anything.

## Scan

```sh
osdy scan [--format text|json|tui]
```

When both stdin and stdout are terminals, the default format is `tui`. Otherwise, the default is `text`.

Examples:

```sh
osdy scan
osdy scan --format text
osdy scan --format json
```

Osdy scans only these fixed locations in the current user's home directory:

- `~/.npm` — npm cache
- `~/Library/Caches/Homebrew` — Homebrew cache
- `~/.gradle/caches` — Gradle caches
- `~/Library/Developer/Xcode/DerivedData` — Xcode DerivedData
- `~/Library/Developer/CoreSimulator` — CoreSimulator

There are no arbitrary path arguments. Osdy is read-only: it does not mutate or delete files, invoke `sudo` or a shell, send telemetry, access external volumes, or use the network.

## Install v0.1.0

Download the archive for your Mac from the v0.1.0 release:

| Mac | Archive |
| --- | --- |
| Apple silicon | `osdy_0.1.0_darwin_arm64.tar.gz` |
| Intel | `osdy_0.1.0_darwin_amd64.tar.gz` |

Verify the downloaded archive against the release checksum, then install the extracted `osdy` binary somewhere on your `PATH` (for example, `~/.local/bin`):

```sh
shasum -a 256 osdy_0.1.0_darwin_arm64.tar.gz
tar -xzf osdy_0.1.0_darwin_arm64.tar.gz
mkdir -p ~/.local/bin
mv osdy ~/.local/bin/osdy
chmod 755 ~/.local/bin/osdy
~/.local/bin/osdy scan --format text
```

Replace `arm64` with `amd64` on an Intel Mac. Ensure `~/.local/bin` is on your `PATH` before using `osdy` by name.

### Gatekeeper

v0.1.0 is unsigned and not notarized. After verifying its release provenance and checksum, use macOS's per-app review flow: try opening the binary, then choose **Open Anyway** in **System Settings → Privacy & Security** if you trust it. Do not disable Gatekeeper globally or use broad quarantine-removal commands.

## Build from source

Go 1.24.2 or newer is required.

```sh
git clone https://github.com/osdy/OsdyCleaner.git
cd OsdyCleaner
go build -o osdy ./cmd/osdy
./osdy scan --format text
```

## Automation and output

`--format json` emits deterministic schema-v1 JSON. It includes the scan outcome, policy, estimates, roots, findings, warnings, and caveats; it is intended for scripts and CI.

```sh
osdy scan --format json > osdy-report.json
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
osdy scan --format text
osdy scan --format json
osdy scan --format tui
```

The TUI command requires terminal input and output. For machine-readable verification, inspect `schema_version` in the JSON output; v0.1.0 emits `1`.

## License

[MIT](LICENSE)
