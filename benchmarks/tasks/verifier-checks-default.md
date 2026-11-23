# `verifier-checks-default` (medium)

The workspace is agenticc's own repository at the commit before a real fix; the hidden check is the test that fix was committed with. agenticc's source is not public, so the workspace and the test are described here, not published.

| | |
|---|---|
| Workspace | agenticc at the parent of `f6d6ddba` (1491 files) |
| Hidden test | `TestVerifierRunsChecksUnderTheChangesDefault` in the fix commit's `internal/subagent/harness_test.go`, written into a copy of the workspace after the agent exits |
| Hidden check | `go test -count=1 -timeout=240s -run ^TestVerifierRunsChecksUnderTheChangesDefault$ ./internal/subagent` |
| Prompt SHA-256 | `227521e3d1e74de5b1c5bd40fb9735d47e9de08e1712ee71f7ae288674db6585` |

## The prompt, exactly as every agent received it

```text
A verifier subagent runs no project checks under the default verification mode. The `changes` default skips checks when a turn changed nothing, and a verifier changes nothing, so it verifies nothing. It should run its configured checks exactly once under that default. An explicit `manual` mode must still stand. Fix this and add a test. Look in internal/subagent and internal/loop/verification.go.
```

The task passes when the hidden check passes after the agent exits. Results: [the benchmark page](../README.md).
