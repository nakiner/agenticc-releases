# `summary-on-cached-conversation` (long)

The workspace is agenticc's own repository at an earlier commit with the feature taken out; the hidden check is a set of `TestAudit` tests written for this task. agenticc's source is not public, so the workspace and the tests are described here, not published.

| | |
|---|---|
| Workspace | agenticc at `cea77263aa9c` plus a patch that removes the feature (3386 files) |
| Workspace SHA-256 | `dbea6d59520f09144cb627d0d190970f7cc2b1d2a31fd6337923569b74261b7f` (removal patch) |
| Hidden tests | `internal/loop/audit_conversation_summary_test.go` (copied in after the agent exits) |
| Hidden check | `go test -count=1 -timeout=120s -run ^TestAudit ./internal/loop` |
| Grading | behavior and host test outcomes; no exact patch |
| Prompt SHA-256 | `212dc06161d3c064de6bd7a6c555aae54b1bf1a15cf32958f4a5355df88acfbb` |

## The prompt, exactly as every agent received it

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

The task passes when the hidden check passes after the agent exits. Results: [the benchmark page](../README.md).
