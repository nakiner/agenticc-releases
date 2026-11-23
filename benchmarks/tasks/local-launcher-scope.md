# `local-launcher-scope` (long)

The workspace is agenticc's own repository at an earlier commit before a real fix; the hidden check is a set of `TestAudit` tests written for this task. agenticc's source is not public, so the workspace and the tests are described here, not published.

| | |
|---|---|
| Workspace | agenticc at `d13485d089c9` (1109 files) |
| Workspace SHA-256 | `3a2fd1a53fdf8f61309956095c4b7c8b7030ce94e8ebd450122bddc73745c82f` (source archive) |
| Hidden tests | `internal/operationalfacts/audit_module_scope_test.go` (copied in after the agent exits) |
| Hidden check | `go test -count=1 -timeout=20s -run ^TestAudit ./internal/operationalfacts` |
| Grading | behavior and host test outcomes; no exact patch, AST, command order or final text |
| Prompt SHA-256 | `c5705dfa09c519b42e4eba79d8c6573f0a0521f41ba784e85821a5932716b322` |

## The prompt, exactly as every agent received it

```text
Saved operational facts for `python3 -m pytest` are reused after local Python module shadowing changes. With the workspace root as the process working directory, track the presence and bounded contents of these root-relative files: pytest.py, pytest.pyc, pytest/__init__.py, pytest/__init__.pyc, pytest/__main__.py, pytest/__main__.pyc. Creating, editing (including same-length edits with restored timestamps), or deleting any of them must invalidate the prior launcher fact; unchanged supported files must still allow reuse. Preserve the existing aggregate read bound and existing handling of unsafe/special files. This issue concerns Python module/package resolution; console-script files such as bin/pytest are outside this change. Add regressions and run relevant Go tests/build checks. Explain the change and what you verified; do not commit.

Suggested package checks: go test -count=1 ./internal/operationalfacts; go build ./internal/operationalfacts.
```

The task passes when the hidden check passes after the agent exits. Results: [the benchmark page](../README.md).
