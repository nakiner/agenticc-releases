# agenticc

A coding agent for your own repositories: a terminal chat, a one-shot CLI and a macOS app.
It asks before it writes or runs anything, runs commands in an operating-system sandbox,
and records every step so an interrupted turn can be resumed.

This repository holds the releases and the download page,
[agenticc.hoolie.io](https://agenticc.hoolie.io/).
The source is not published here.

## Install

```sh
curl -fsSL https://github.com/nakiner/agenticc-releases/releases/latest/download/install.sh | sh
```

The script picks the binary for your machine, checks its SHA-256 against the release's
`SHA256SUMS` and installs it as `~/.local/bin/agenticc`, without sudo. For the macOS app,
download the `.dmg` from the [latest release](https://github.com/nakiner/agenticc-releases/releases/latest).
Later releases: `agenticc update`; the app updates itself.

| Release file | For |
| --- | --- |
| `agenticc-darwin` | macOS 13 or later, one universal binary, signed and notarized |
| `agenticc-linux-amd64`, `agenticc-linux-arm64` | Linux (needs `bubblewrap` to run commands) |
| `agenticc-<version>.dmg` | The macOS desktop app, signed, notarized and stapled |
| `agenticc-<version>-macos.zip`, `agenticc-update.json` | The app's in-app update and its signed manifest |
| `SHA256SUMS`, `VERSION`, `install.sh` | Checksums, the version, the installer |

Benchmarks against Claude Code, Codex CLI and Hermes, the tasks and the raw results:
[benchmarks/](benchmarks/).

Released under the MIT license.
