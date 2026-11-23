# `strip-line-ending` (short)

A self-contained task: the whole workspace, the prompt and the hidden test are in this folder.

| | |
|---|---|
| Workspace | [`workspace/`](workspace/), a small Go module |
| Hidden test | [`hidden-test/`](hidden-test/), copied in after the agent exits |
| Hidden check | `go test -json -count=1 ./...` |
| Prompt SHA-256 | `a13ce5826f43706d5816b8b20962d56609bb38f0180a06d4c1d4223320e18ffb` |

## The prompt, exactly as every agent received it

```text
Small Go repair: strip-line-ending.

Implement StripLineEnding(text string) string. Remove exactly one trailing line ending, either "\r\n" or "\n", and nothing else. Text without a trailing line ending is returned unchanged, including leading or trailing spaces and tabs. A lone trailing "\r" is not a line ending and stays. Only one line ending is removed: "x\n\n" becomes "x\n".

Change only solution.go. Keep API names/signatures and package task. Do not edit go.mod or solution_test.go; do not add/delete/rename workspace files or use symlinks. Atomic save/format staging renamed directly to solution.go is allowed. Put any scratch programs/tests outside the workspace, for example under TMPDIR, and keep the final workspace file set unchanged. Use only the Go1.27 standard library. You may inspect/edit/format the file and run go test ./... Report checks actually run. Held-out tests only enforce the explicit contract above.
```

The task passes when the hidden check passes after the agent exits. Results: [the benchmark page](../../README.md).
