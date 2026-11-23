# `workspace-resource-fencing` (long)

The workspace is agenticc's own repository at an earlier commit before a real fix; the hidden check is a set of `TestAudit` tests written for this task. agenticc's source is not public, so the workspace and the tests are described here, not published.

| | |
|---|---|
| Workspace | agenticc at `65caf52a9624` (1148 files) |
| Workspace SHA-256 | `238c6b38a45d929b5e83ea7f1fdb8d4986ebd332323e482eaf8a346bee62baf1` (source archive) |
| Hidden tests | `internal/loop/audit_loop_resource_claim_test.go` (copied in after the agent exits) |
| Hidden check | `go test -count=1 -timeout=20s -run ^TestAudit ./internal/loop ./internal/runner` |
| Grading | behavior and host test outcomes; no exact patch, AST, command order or final text |
| Prompt SHA-256 | `ec069932265f6f058094ba7b61b0f99ce019770112b28256ab85a868db441916` |

## The prompt, exactly as every agent received it

```text
Fix workspace scheduling isolation. A search covering the workspace root must wait for conflicting file writes. Extension tools and shell commands with unknown effects must not bypass a conflicting writer by supplying an unrelated path argument, including after recovery of an unknown outcome. Preserve concurrency for genuinely independent built-in file operations. Add regressions and run relevant Go tests and build checks. Explain what changed and what you verified; do not commit.

Suggested package checks: go test -count=1 ./internal/loop ./internal/runner; go build ./internal/loop ./internal/runner.
```

The task passes when the hidden check passes after the agent exits. Results: [the benchmark page](../README.md).
