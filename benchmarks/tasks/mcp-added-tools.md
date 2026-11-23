# `mcp-added-tools` (medium)

The workspace is agenticc's own repository at the commit before a real fix; the hidden check is the test that fix was committed with. agenticc's source is not public, so the workspace and the test are described here, not published.

| | |
|---|---|
| Workspace | agenticc at the parent of `7ab5a2f6` (1491 files) |
| Hidden test | `TestDeferredWarmRefreshLearnsAboutAddedTools` in the fix commit's `internal/mcp/catalog_wait_test.go`, written into a copy of the workspace after the agent exits |
| Hidden check | `go test -count=1 -timeout=240s -run ^TestDeferredWarmRefreshLearnsAboutAddedTools$ ./internal/mcp` |
| Prompt SHA-256 | `411daa06bf8c5904386f5f52dc033d0f0d590a39eda7064c9c24a88b828bb8a7` |

## The prompt, exactly as every agent received it

```text
With deferred warm refresh (DeferWarmRefresh), the MCP descriptor cache is only corrected when the called tool's own descriptor changed. If a server adds a tool and keeps its existing descriptors, every later run keeps the old inventory. After a successful call, the next run must offer the added tool. The current run's calls must keep working, and this must not open a second connection. Fix this and add a test. See internal/mcp/catalog.go.
```

The task passes when the hidden check passes after the agent exits. Results: [the benchmark page](../README.md).
