# agenticc benchmarks

Seven Go repair tasks of three sizes, run with the same model, the same account and the same
harness for every agent. Each task is graded by a hidden test the agent never sees: it is
written into a copy of the workspace after the agent exits, and the task passes only if that
test passes. Cost is computed from the tokens each agent reported, at the model's list price,
including the provider's split between cache reads and writes. Time is the agent process from
start to exit.

## Tasks

| Size | Task | What is wrong |
| --- | --- | --- |
| short | [`config-port`](tasks/config-port/) | Parse a port number from config text, with a fallback |
| short | [`strip-line-ending`](tasks/strip-line-ending/) | Remove exactly one trailing line ending |
| medium | [`mcp-added-tools`](tasks/mcp-added-tools.md) | MCP descriptor cache misses tools a server added |
| medium | [`verifier-checks-default`](tasks/verifier-checks-default.md) | Verifier sub-agent runs no project checks by default |
| long | [`local-launcher-scope`](tasks/local-launcher-scope.md) | Cached `python3 -m pytest` facts survive module shadowing |
| long | [`summary-on-cached-conversation`](tasks/summary-on-cached-conversation.md) | Ask for the compaction summary on the cached conversation |
| long | [`workspace-resource-fencing`](tasks/workspace-resource-fencing.md) | Searches and unknown-effect tools must wait for conflicting writes |

Each link shows the prompt exactly as every agent received it, the starting workspace, the
hidden check and SHA-256 sums of the prompt and workspace. The two short tasks are published
whole: workspace, prompt and hidden test. The medium and long tasks start from agenticc's
own repository at an earlier commit, which is not public, so their pages give the prompt,
the commit, the hidden test's name and the check command, but not the code.

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

The per-task medians with their low and high values are in [`results/`](results/), one JSON
file per agent and model.

## Limits

- Seven tasks, all Go repairs with a hidden test, five of them from agenticc's own history,
  which the agent may have been tuned against. Nothing here measures greenfield work, other
  languages or very long sessions.
- The other agents were measured on earlier days than agenticc 0.8.0; only a same-day run
  compares time fairly.
- Cost is computed at list price. A subscription login is not billed this way.
