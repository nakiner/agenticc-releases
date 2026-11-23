# `config-port` (short)

A self-contained task: the whole workspace, the prompt and the hidden test are in this folder.

| | |
|---|---|
| Workspace | [`workspace/`](workspace/), a small Go module |
| Hidden test | [`hidden-test/`](hidden-test/), copied in after the agent exits |
| Hidden check | `go test -json -count=1 ./...` |
| Prompt SHA-256 | `05743281dc0c923e3f89552e13a86e6901cfa723489490cd699a150edc7ba04e` |

## The prompt, exactly as every agent received it

```text
Small Go repair: config-port.

Implement ConfigPort(text string, fallback int) (int,error). Empty text uses fallback, which must be in 1..65535. A nonempty text ignores fallback and must consist only of ASCII decimal digits with a numeric value in 1..65535. Leading zeros are allowed. Reject whitespace, signs, non-ASCII digits, prefixes, fractions, overflow, zero and values above 65535. Every error returns 0 and non-nil error; error wording is unspecified. Do not add a digit-count cap.

Change only solution.go. Keep API names/signatures and package task. Do not edit go.mod or solution_test.go; do not add/delete/rename workspace files or use symlinks. Atomic save/format staging renamed directly to solution.go is allowed. Put any scratch programs/tests outside the workspace, for example under TMPDIR, and keep the final workspace file set unchanged. Use only the Go1.27 standard library. You may inspect/edit/format the file and run go test ./... Report checks actually run. Held-out tests only enforce the explicit contract above.
```

The task passes when the hidden check passes after the agent exits. Results: [the benchmark page](../../README.md).
