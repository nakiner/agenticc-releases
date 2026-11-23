# agenticc benchmarks

The tasks behind the numbers on [agenticc.hoolie.io](https://agenticc.hoolie.io/), and the
raw results.

Every agent got the same seven Go repair tasks, the same model, the same account and the
same harness. Each task is graded by a test the agent never sees, run after it exits; a task
passes only if that test passes. Cost is computed from the tokens each agent reported, at the
model's list price. Time is the agent process from start to exit.

Two tasks are small standalone functions. The other five are real bugs and one removed
feature from agenticc's own history: the agent works in agenticc's repository at the commit
before the fix (about 1,500 to 3,400 files), so they look like everyday work in a large codebase.
That repository is not public, so the code is not published here; the prompts are, exactly
as the agents received them.

## Tasks

### 1. Parse a port number from config text — `config-port` (short)

A small Go function: accept only ASCII digits in 1..65535, fall back to a default for empty text, reject everything else.

Graded by: a held-out Go test of the exact contract.

<details><summary>The prompt every agent received</summary>

```text
Small Go repair: config-port.

Implement ConfigPort(text string, fallback int) (int,error). Empty text uses fallback, which must be in 1..65535. A nonempty text ignores fallback and must consist only of ASCII decimal digits with a numeric value in 1..65535. Leading zeros are allowed. Reject whitespace, signs, non-ASCII digits, prefixes, fractions, overflow, zero and values above 65535. Every error returns 0 and non-nil error; error wording is unspecified. Do not add a digit-count cap.

Change only solution.go. Keep API names/signatures and package task. Do not edit go.mod or solution_test.go; do not add/delete/rename workspace files or use symlinks. Atomic save/format staging renamed directly to solution.go is allowed. Put any scratch programs/tests outside the workspace, for example under TMPDIR, and keep the final workspace file set unchanged. Use only the Go1.27 standard library. You may inspect/edit/format the file and run go test ./... Report checks actually run. Held-out tests only enforce the explicit contract above.
```

</details>

### 2. Remove exactly one trailing line ending — `strip-line-ending` (short)

A one-function Go fix: strip one `\r\n` or `\n` and nothing else.

Graded by: a held-out Go test of the exact contract.

<details><summary>The prompt every agent received</summary>

```text
Small Go repair: strip-line-ending.

Implement StripLineEnding(text string) string. Remove exactly one trailing line ending, either "\r\n" or "\n", and nothing else. Text without a trailing line ending is returned unchanged, including leading or trailing spaces and tabs. A lone trailing "\r" is not a line ending and stays. Only one line ending is removed: "x\n\n" becomes "x\n".

Change only solution.go. Keep API names/signatures and package task. Do not edit go.mod or solution_test.go; do not add/delete/rename workspace files or use symlinks. Atomic save/format staging renamed directly to solution.go is allowed. Put any scratch programs/tests outside the workspace, for example under TMPDIR, and keep the final workspace file set unchanged. Use only the Go1.27 standard library. You may inspect/edit/format the file and run go test ./... Report checks actually run. Held-out tests only enforce the explicit contract above.
```

</details>

### 3. A tool cache that never learns about new tools — `mcp-added-tools` (medium)

A real bug from agenticc's history: after an MCP server adds a tool, later runs keep offering the old tool list. The workspace is the repository at the commit before the fix.

Graded by: the test that the real fix was committed with.

<details><summary>The prompt every agent received</summary>

```text
With deferred warm refresh (DeferWarmRefresh), the MCP descriptor cache is only corrected when the called tool's own descriptor changed. If a server adds a tool and keeps its existing descriptors, every later run keeps the old inventory. After a successful call, the next run must offer the added tool. The current run's calls must keep working, and this must not open a second connection. Fix this and add a test. See internal/mcp/catalog.go.
```

</details>

### 4. A verifier that verifies nothing — `verifier-checks-default` (medium)

A real bug from agenticc's history: a verifier sub-agent skips the project's checks under the default mode, because it changed no files.

Graded by: the test that the real fix was committed with.

<details><summary>The prompt every agent received</summary>

```text
A verifier subagent runs no project checks under the default verification mode. The `changes` default skips checks when a turn changed nothing, and a verifier changes nothing, so it verifies nothing. It should run its configured checks exactly once under that default. An explicit `manual` mode must still stand. Fix this and add a test. Look in internal/subagent and internal/loop/verification.go.
```

</details>

### 5. A stale cache after a local module shadows pytest — `local-launcher-scope` (long)

A real bug from agenticc's history: a cached fact about `python3 -m pytest` survives creating or editing a local `pytest.py` that changes what runs.

Graded by: hidden behaviour tests written for the task; no exact patch is required.

<details><summary>The prompt every agent received</summary>

```text
Saved operational facts for `python3 -m pytest` are reused after local Python module shadowing changes. With the workspace root as the process working directory, track the presence and bounded contents of these root-relative files: pytest.py, pytest.pyc, pytest/__init__.py, pytest/__init__.pyc, pytest/__main__.py, pytest/__main__.pyc. Creating, editing (including same-length edits with restored timestamps), or deleting any of them must invalidate the prior launcher fact; unchanged supported files must still allow reuse. Preserve the existing aggregate read bound and existing handling of unsafe/special files. This issue concerns Python module/package resolution; console-script files such as bin/pytest are outside this change. Add regressions and run relevant Go tests/build checks. Explain the change and what you verified; do not commit.

Suggested package checks: go test -count=1 ./internal/operationalfacts; go build ./internal/operationalfacts.
```

</details>

### 6. Make a summary cheap by reusing the prompt cache — `summary-on-cached-conversation` (long)

A feature removed from agenticc for the task: ask for the compaction summary as one more round of the cached conversation instead of a fresh, uncached request, with a fallback.

Graded by: hidden behaviour tests written for the task; no exact patch is required.

<details><summary>The prompt every agent received</summary>

```text
Context compaction asks the model for a summary in a fresh request that holds the older conversation as flat text. Nothing of that request is in the provider's prompt cache, so the model is paid in full to read again what it was just sent, and on a cache-priced connection a summary costs more than the rounds it saves. Change it so the summary is asked as one more round of the conversation the provider already holds, and keep the existing separate request as the fallback.

Behaviour required:
- When a compaction of the current round needs a summary, first ask the main provider (Loop.Provider) with a request made of: the History of the last main request exactly as it was handed to the provider, then the messages the turn's history has gained since that request, then one final host note. The note is a user-role message marked Runtime that asks for a checkpoint summary of the whole conversation, says "do not call tools", and gives the output target (SummaryTargetTokens). The request must carry the same Tools, SessionID, Effort and MaxOutputTokens as that last main request, so the provider's cache key does not change. The provider call carries the provider purpose "summary".
- Use the reply's text as the summary only when the call succeeded, returned no tool calls, and its trimmed text is non-empty and within the existing summary size limit. In every other case, and when the request would not fit Loop.ContextBudget (estimate with historyEstimateTokens), fall back to the existing separate summary path. Ask on the conversation at most once per compaction; account for what a failed attempt cost.
- Add the owner setting efficiency.Policy.SummarySource (JSON key "summary_source"): "conversation" or "separate". efficiency.Defaults sets "conversation" when it is empty; Validate rejects any other value. A Loop whose policy does not say "conversation" (including a Loop built with a zero policy) never asks on the conversation.
- Everything else about compaction, the summary cache and its journal records stays as it is.

Names the regression tests rely on, all in package loop unless stated:
- type sentRequest struct { request Request; rawLen int } — the last main request handed to the provider and how many messages of the turn's history it covered; stored as l.prompt.last (*sentRequest) when a main request is sent.
- func withConversationSummary(ctx context.Context, l *Loop, history []Message) context.Context — returns a context that lets one compaction ask on the conversation; it returns ctx unchanged when the setting is not "conversation", when there is no last request or no provider, or when history is shorter than that request covered. The round passes this context to the code that prepares the wire history (l.compactWire runs under it).

Constraints of this repository apply: Loop may not gain methods (internal/loop/method_count_test.go), Loop.turn may not grow (scripts/funcsize), new code needs tests. Add regressions and run the relevant Go tests and build checks. Explain the change and what you verified; do not commit.

Suggested package checks: go test -count=1 ./internal/loop ./internal/efficiency; go build ./...; go run ./scripts/funcsize -base HEAD.
```

</details>

### 7. Commands that slip past a conflicting write — `workspace-resource-fencing` (long)

A real bug from agenticc's history: searches and tools with unknown effects could run alongside a conflicting file write.

Graded by: hidden behaviour tests written for the task; no exact patch is required.

<details><summary>The prompt every agent received</summary>

```text
Fix workspace scheduling isolation. A search covering the workspace root must wait for conflicting file writes. Extension tools and shell commands with unknown effects must not bypass a conflicting writer by supplying an unrelated path argument, including after recovery of an unknown outcome. Preserve concurrency for genuinely independent built-in file operations. Add regressions and run relevant Go tests and build checks. Explain what changed and what you verified; do not commit.

Suggested package checks: go test -count=1 ./internal/loop ./internal/runner; go build ./internal/loop ./internal/runner.
```

</details>

## Results

Medians of three runs per task, written as passed runs · time · cost. agenticc 0.8.0 was
measured on 23 Nov 2026. The other agents are their latest reference values: Claude Code and
Codex CLI on 21 Nov 2026 (three runs per task), Hermes v0.21.5 on 13 Nov 2026 (one run per
task). Times also reflect the provider's speed on the day; costs are steadier.

### Opus 5.5

| Task | agenticc 0.8.0 | Claude Code |
| --- | --- | --- |
| `config-port` | 3/3 · 9 s · $0.029 | 3/3 · 13 s · $0.080 |
| `strip-line-ending` | 3/3 · 8 s · $0.023 | 3/3 · 12 s · $0.071 |
| `mcp-added-tools` | 3/3 · 79 s · $0.342 | 3/3 · 299 s · $0.494 |
| `verifier-checks-default` | 3/3 · 66 s · $0.252 | 3/3 · 81 s · $0.317 |
| `local-launcher-scope` | 3/3 · 54 s · $0.189 | 3/3 · 193 s · $0.306 |
| `summary-on-cached-conversation` | 1/3 · 394 s · $1.198 | 3/3 · 763 s · $1.555 |
| `workspace-resource-fencing` | 3/3 · 238 s · $0.942 | 3/3 · 358 s · $1.035 |
| All seven | **$2.98** · 848 s · 86 requests | **$3.86** · 1719 s · 102 requests |

agenticc's two failed runs were both on `summary-on-cached-conversation`: the code the model
wrote left an extra user message ahead of the summary request, which the hidden test rejects.
A rerun of that task alone passed 3 of 3 (377 s, $1.220).

### GPT 6.1 Sol

| Task | agenticc 0.8.0 | Codex CLI | Hermes |
| --- | --- | --- | --- |
| `config-port` | 3/3 · 17 s · $0.018 | 3/3 · 27 s · $0.035 | 1/1 · 29 s · $0.048 |
| `strip-line-ending` | 3/3 · 12 s · $0.015 | 3/3 · 17 s · $0.027 | 1/1 · 51 s · $0.042 |
| `mcp-added-tools` | 3/3 · 106 s · $0.154 | 3/3 · 347 s · $0.335 | 1/1 · 449 s · $0.283 |
| `verifier-checks-default` | 3/3 · 101 s · $0.085 | 3/3 · 368 s · $0.347 | 1/1 · 469 s · $0.336 |
| `local-launcher-scope` | 3/3 · 92 s · $0.074 | 3/3 · 118 s · $0.133 | 1/1 · 200 s · $0.219 |
| `summary-on-cached-conversation` | 3/3 · 477 s · $0.382 | 3/3 · 546 s · $0.649 | 1/1 · 1297 s · $0.723 |
| `workspace-resource-fencing` | 3/3 · 294 s · $0.234 | 3/3 · 360 s · $0.390 | 1/1 · 457 s · $0.400 |
| All seven | **$0.96** · 1098 s · 64 requests | **$1.92** · 1784 s · 158 requests | **$2.05** · 2953 s · 147 requests |

Codex CLI ran with its user's memory and 25 skills loaded, which adds requests, so the gap
to Codex is overstated by an unknown amount.

## Raw results

[`results/`](results/) holds one JSON file per agent and model: per task the runs, passed
runs, median time, requests and median cost with its low and high.

## Limits

- Seven tasks, all Go, five from agenticc's own history, which agenticc may have been tuned
  against. Nothing here measures greenfield work, other languages or very long sessions.
- The other agents were measured on earlier days than agenticc 0.8.0, so time also reflects
  the provider's speed on the day. Costs are steadier.
- Cost is at list price. A subscription login is not billed this way.
