# kiro-go — Build Specification

> Hand this file to Claude Code. It contains everything needed to build the full project.
> Do not start coding until you have read this entire document.

---

## What This Project Is

This repo ships **two separate CLI binaries** from a single Go module:

### Binary 1: `kiro-specs`
Language-agnostic spec manager for any Kiro project regardless of tech stack.
Works with TypeScript, Python, Rust, Go — anything that uses `.kiro/specs/`.
This is the primary value driver of the repo — useful to every Kiro user.

### Binary 2: `kiro-go`
Opinionated Go project scaffolder. Generates `.kiro/` steering docs, hooks, and
spec templates tuned specifically for Go idioms. Includes `kiro-specs` functionality
built in (calls the same `specmanager` package internally).

Both binaries are distributed from the same repo. Users can install either or both.

---

## Splash Screens

### `kiro-specs` splash

When run with no arguments or `--help`:

```
 ██╗  ██╗██╗██████╗  ██████╗     ███████╗██████╗ ███████╗ ██████╗███████╗
 ██║ ██╔╝██║██╔══██╗██╔═══██╗    ██╔════╝██╔══██╗██╔════╝██╔════╝██╔════╝
 █████╔╝ ██║██████╔╝██║   ██║    ███████╗██████╔╝█████╗  ██║     ███████╗
 ██╔═██╗ ██║██╔══██╗██║   ██║    ╚════██║██╔═══╝ ██╔══╝  ██║     ╚════██║
 ██║  ██╗██║██║  ██║╚██████╔╝    ███████║██║     ███████╗╚██████╗███████║
 ╚═╝  ╚═╝╚═╝╚═╝  ╚═╝ ╚═════╝    ╚══════╝╚═╝     ╚══════╝ ╚═════╝╚══════╝
```

ASCII art colour: bright cyan (`\033[96m`)

Below the art:
```
v0.1.0  ·  works with any Kiro project  ·  language-agnostic

Spec manager for Kiro — track status, find what to work on next, archive completed work.
```

Then divider, then command menu:
```
  list               List all specs with status and progress  (default)
  show [name]        Show task list for a spec
  next               What should I work on right now?
  archive [name]     Move a completed spec out of the way
  stats              Summary stats across all specs

  Flags (all commands):
  --status [status]  Filter by: idea, draft, planned, in-progress, complete
  --sort [by]        Sort by: status (default), age, name
  --json             Machine-readable JSON output
  --path [dir]       Path to project root (default: current directory)
```

---

### `kiro-go` splash

When run with no arguments or `--help`:

```
 ██╗  ██╗██╗██████╗  ██████╗        ██████╗  ██████╗
 ██║ ██╔╝██║██╔══██╗██╔═══██╗      ██╔════╝ ██╔═══██╗
 █████╔╝ ██║██████╔╝██║   ██║█████╗██║  ███╗██║   ██║
 ██╔═██╗ ██║██╔══██╗██║   ██║╚════╝██║   ██║██║   ██║
 ██║  ██╗██║██║  ██║╚██████╔╝      ╚██████╔╝╚██████╔╝
 ╚═╝  ╚═╝╚═╝╚═╝  ╚═╝ ╚═════╝        ╚═════╝  ╚═════╝
```

ASCII art colour: bright cyan. Below the art print version pills and badge:
- `v0.1.0` blue pill
- `Go 1.23+` green pill
- `spec-driven` grey pill
- `[ GO ]` cyan badge

Then:
```
Opinionated Kiro template for Go — spec-driven, idiomatic, production-ready.
Steering docs · Hooks · REST API · CLI · Microservice presets
```

Then divider, then command menu:
```
  init [name]     Scaffold a new Go project
  add             Inject .kiro/ into an existing project
  preset [name]   Apply a preset (api, cli, svc)
  list            List available presets
  specs           Manage specs (powered by kiro-specs)
  version         Show version info
  help            Show this help screen
```

Note: `kiro-go specs` is identical to running `kiro-specs` — same code, same output.

---

## Project Structure

```
kiro-go/                              # repo root
├── go.mod                            # module github.com/your-username/kiro-go (TODO: update)
├── go.sum
│
├── cmd/
│   ├── kiro-specs/
│   │   └── main.go                   # entry point for kiro-specs binary
│   └── kiro-go/
│       └── main.go                   # entry point for kiro-go binary
│
├── internal/
│   ├── specmanager/                  # shared — used by both binaries
│   │   ├── parser.go                 # reads spec folders, parses tasks.md checkboxes
│   │   ├── parser_test.go
│   │   ├── status.go                 # derives status from parsed Spec struct
│   │   ├── status_test.go
│   │   ├── table.go                  # renders ANSI table to io.Writer
│   │   └── types.go                  # Spec struct, Status enum, shared types
│   ├── splash/
│   │   ├── specs.go                  # kiro-specs splash screen
│   │   └── kirogo.go                 # kiro-go splash screen
│   ├── scaffold/
│   │   └── scaffold.go               # copies embedded templates to target dir
│   └── presets/
│       └── presets.go                # preset definitions and metadata
│
├── specs/                            # kiro-specs subcommands (used by both binaries)
│   ├── list.go
│   ├── show.go
│   ├── archive.go
│   ├── next.go
│   └── stats.go
│
├── kirogo/                           # kiro-go only commands
│   ├── init.go
│   ├── add.go
│   ├── preset.go
│   └── list.go
│
├── templates/                        # embedded via go:embed (kiro-go only)
│   └── .kiro/
│       ├── steering/
│       │   ├── 01-product.md
│       │   ├── 02-tech.md
│       │   ├── 03-structure.md
│       │   ├── 04-go-idioms.md
│       │   ├── 05-error-handling.md
│       │   ├── 06-testing.md
│       │   ├── 07-security.md
│       │   ├── 08-api-design.md
│       │   ├── 09-cli-design.md
│       │   └── 10-observability.md
│       ├── hooks/
│       │   ├── 01-lint-on-save.md
│       │   ├── 02-test-on-save.md
│       │   ├── 03-modernize.md
│       │   ├── 04-security-scan.md
│       │   └── 05-doc-check.md
│       └── specs/
│           ├── _FEATURE-TEMPLATE/
│           │   ├── requirements.md
│           │   ├── design.md
│           │   └── tasks.md
│           └── _API-ENDPOINT-TEMPLATE/
│               ├── requirements.md
│               ├── design.md
│               └── tasks.md
│
└── README.md
```

---

## Installation

### Install `kiro-specs` only (any Kiro project)
```bash
go install github.com/your-username/kiro-go/cmd/kiro-specs@latest
```

### Install `kiro-go` (Go projects — includes kiro-specs functionality)
```bash
go install github.com/your-username/kiro-go/cmd/kiro-go@latest
```

Both should be documented in the README prominently.

---

## `kiro-specs` — Full Command Specification

### How status is derived

The `specmanager` package scans `.kiro/specs/`, skips template folders (names starting with `_`) and the `_archive` folder, and derives status automatically.

Status derivation rules (priority order):

| Condition | Status |
|---|---|
| Folder exists, no files inside | `idea` |
| Has `requirements.md` only, no `tasks.md` | `draft` |
| Has `tasks.md`, zero checked boxes | `planned` |
| Has `tasks.md`, mix of checked and unchecked | `in-progress` |
| Has `tasks.md`, all boxes checked | `complete` |
| Folder name starts with `_` | skip |

Checkbox parsing: scan `tasks.md` line by line with `bufio.Scanner`. Count lines containing `- [x]` (checked) and `- [ ]` (unchecked) using `strings.Contains`. No regex.

Progress bar: 6 block characters. `filled = round(done/total * 6)`. Use `█` for filled, `░` for empty. If total is 0, return `──────`.

Last modified: `os.Stat(folder).ModTime()`. Format as relative time — implement `relativeTime(t time.Time) string` using only stdlib. Cases: same day → "today", 1 day ago → "yesterday", < 7 days → "N days ago", < 30 days → "N weeks ago", else → "N months ago".

`--path` flag: all commands accept `--path /some/project` to point at a project root other than cwd. Default is `os.Getwd()`. Spec dir is always `[path]/.kiro/specs/`.

---

### `kiro-specs list` (default when no subcommand given)

```
  SPEC                          STATUS        PROGRESS    AGE
  ──────────────────────────────────────────────────────────────────
  user-authentication           in-progress   ████░░ 4/6  2 days ago
  middleware-refactor           in-progress   ██░░░░ 2/7  yesterday
  rate-limiting                 planned       ░░░░░░ 0/8  1 week ago
  csv-export                    draft         ──────      5 days ago
  dark-mode                     idea          ──────      1 month ago
  payment-webhooks              complete      ██████ 5/5  3 weeks ago

  6 specs  |  1 complete  |  2 in-progress  |  1 planned  |  1 draft  |  1 idea
```

Colour (ANSI, disabled when `NO_COLOR` set or not a TTY):
- `complete` → green
- `in-progress` → yellow
- `planned` → blue
- `draft` → grey
- `idea` → dark grey

Default sort order: in-progress first, planned, draft, idea, complete last.

Flags:
- `--status [status]` filter to one status
- `--sort age` newest first by ModTime
- `--sort name` alphabetical
- `--json` JSON array to stdout

JSON output format:
```json
[
  {
    "name": "user-authentication",
    "status": "in-progress",
    "tasks_total": 6,
    "tasks_done": 4,
    "age_days": 2,
    "path": ".kiro/specs/user-authentication"
  }
]
```

---

### `kiro-specs show [name]`

```
  user-authentication  ·  in-progress  ·  4/6 tasks done  ·  2 days ago
  ──────────────────────────────────────────────────────────────────────

  [x] 1. Create user model and database schema
  [x] 2. Implement password hashing with bcrypt
  [x] 3. Build registration endpoint POST /api/v1/auth/register
  [x] 4. Build login endpoint POST /api/v1/auth/login
  [ ] 5. Implement JWT token refresh
  [ ] 6. Add rate limiting to auth endpoints

  Next up: Implement JWT token refresh
```

- "Next up" = first unchecked task
- If all complete: "All tasks complete."
- If no `tasks.md`: show `requirements.md` contents with note "(no tasks.md yet — showing requirements)"
- Fuzzy match: if exact name not found, find specs containing the query as substring. If one match, use it automatically. If multiple, list suggestions and exit.

---

### `kiro-specs next`

```
  Pick up where you left off:

  user-authentication  ·  in-progress  ·  2 tasks remaining

    [ ] 5. Implement JWT token refresh
    [ ] 6. Add rate limiting to auth endpoints

  Run: kiro-specs show user-authentication
```

Algorithm:
1. Find all `in-progress` specs, sort by fewest remaining tasks (closest to done first). Tie-break: most recently modified.
2. If none, find all `planned` specs, sort by oldest (started longest ago).
3. If none: print "All specs are complete or ideas. Start a new one in Kiro."

---

### `kiro-specs archive [name]`

Moves `.kiro/specs/[name]` to `.kiro/specs/_archive/[name]`.

Steps:
1. Check spec exists — error if not
2. If status is not `complete`, warn: "This spec is not complete (status: in-progress). Archive anyway? (y/N)"
3. Create `_archive/` if needed
4. `os.Rename` the folder
5. Print: "Archived user-authentication → .kiro/specs/_archive/user-authentication"

Flag: `--yes` / `-y` skips confirmation.

---

### `kiro-specs stats`

```
  Spec summary
  ─────────────────────────────
  Total specs       12
  Complete           4  (33%)
  In progress        3
  Planned            3
  Draft              1
  Ideas              1

  Tasks completed   31 / 67  (46%)
  Oldest active     middleware-refactor  (3 months ago)
  Most recent       dark-mode  (today)
```

"Oldest active" = oldest spec that is not `complete` or `idea`.

---

## `kiro-go` — Command Specification

### `kiro-go init [name]`

1. Create directory `[name]/`
2. Prompt for GitHub username/org → run `go mod init github.com/[username]/[name]`
3. Create minimal `main.go`
4. Copy full `.kiro/` template into `[name]/.kiro/`
5. Prompt: "Which preset? (api / cli / svc / none)"
6. Apply chosen preset
7. Print success + next steps

### `kiro-go add`

1. Check `go.mod` exists in cwd — warn and exit if not
2. Check if `.kiro/` exists — prompt "Overwrite? (y/N)" if so
3. Copy `.kiro/` template into cwd
4. Prompt for preset
5. Print success

### `kiro-go preset [name]`

Applies preset to existing `.kiro/steering/` by modifying `inclusion:` frontmatter.
Presets: `api`, `cli`, `svc`

### `kiro-go list`

```
PRESET   DESCRIPTION
api      REST API — net/http or chi router, middleware, handlers, JSON responses
cli      CLI tool — cobra or flag, subcommands, config file, stdout/stderr hygiene
svc      Microservice — gRPC or HTTP, health checks, graceful shutdown, telemetry
```

### `kiro-go specs [subcommand]`

Delegates directly to the same `specmanager` package and `specs/` commands used by `kiro-specs`. Identical behaviour. Just a convenience so Go developers don't need to install a second binary.

---

## Preset Definitions

### `api`
Steering files always active: 01, 02, 03, 04, 05, 06, 07, 08, 10
Steering files manual: 09
Hooks active: 01, 02, 03, 04, 05
Spec template: `_API-ENDPOINT-TEMPLATE/`

### `cli`
Steering files always active: 01, 02, 03, 04, 05, 06, 07, 09
Steering files manual: 08, 10
Hooks active: 01, 02, 03, 05
Spec template: `_FEATURE-TEMPLATE/`

### `svc`
Steering files always active: 01, 02, 03, 04, 05, 06, 07, 08, 10
Steering files manual: 09
Hooks active: 01, 02, 03, 04, 05
Spec template: `_API-ENDPOINT-TEMPLATE/`

---

## Steering File Contents

### `01-product.md`
```markdown
---
inclusion: always
---
# Product overview

<!-- TODO: Replace with your product description -->

## What this is
[Describe what your Go service/tool/API does in 2-3 sentences]

## Who uses it
[Describe the primary users or consumers]

## Key goals
1. [Primary goal]
2. [Secondary goal]
3. [Third goal]

## Out of scope
- [What this project intentionally does NOT do]
```

### `02-tech.md`
```markdown
---
inclusion: always
---
# Tech stack

## Language & runtime
- **Go 1.23+** — minimum version required
- **Modules** — `go.mod` / `go.sum`

## Web / transport layer
<!-- TODO: Uncomment what applies -->
- `net/http` — standard library HTTP (default)
<!-- - `github.com/go-chi/chi/v5` — lightweight idiomatic router -->
<!-- - `google.golang.org/grpc` — gRPC transport -->

## Database
<!-- TODO: Uncomment what applies -->
<!-- - `github.com/jackc/pgx/v5` — PostgreSQL (preferred) -->
<!-- - `database/sql` + `lib/pq` — PostgreSQL -->

## Configuration
- Environment variables via `os.Getenv` (default)
<!-- - `github.com/spf13/viper` — config file + env var support -->

## Logging
- `log/slog` — structured logging (Go 1.21+, stdlib)
- Never use `fmt.Println` for application logging
- Never use `log.Fatal` outside of `main()`

## Testing
- `testing` — stdlib (always)
- `github.com/stretchr/testify` — assertions and mocking
- Table-driven tests are the standard pattern

## Linting & quality
- `golangci-lint` — all-in-one linter runner
- `govulncheck` — vulnerability scanning

## Common commands
| Command | Description |
|---|---|
| `go run ./...` | Run |
| `go test ./...` | Test |
| `go test -race ./...` | Test with race detector |
| `golangci-lint run` | Lint |
| `govulncheck ./...` | Vulnerability scan |
| `go mod tidy` | Clean dependencies |
```

### `03-structure.md`
```markdown
---
inclusion: always
---
# Project structure

<!-- TODO: Update to match your actual layout -->

## Standard layout
```
cmd/[appname]/main.go   # Entry point only — no business logic
internal/[domain]/      # All application code
  [domain].go
  [domain]_test.go      # Tests alongside the code they test
api/                    # OpenAPI specs, proto files
configs/                # Config files
```

## Conventions
- `cmd/` contains only entry points — no business logic
- `internal/` for all code not meant for external use
- One package per directory
- Package names: lowercase, single words, no underscores
- File names: lowercase with underscores (`user_handler.go`)
- Interfaces: `-er` suffix where natural (`Reader`, `Storer`, `Handler`)
- Constructors: `New[Type](...)` — returns concrete type, not interface
- Sentinel errors: `Err[Description]` (`ErrNotFound`, `ErrTimeout`)
```

### `04-go-idioms.md`
```markdown
---
inclusion: always
---
# Go idioms — non-negotiable standards

## Errors
- Always return errors as the last return value
- Never discard errors with `_` without an explanatory comment
- Never use `panic` outside of `main()` startup failures
- Always wrap: `fmt.Errorf("doing X: %w", err)` — use `%w` not `%v`

## Interfaces
- Define at the point of use (consuming package), not the implementing package
- Keep small — one or two methods
- Never return interfaces from constructors — return concrete types
- Accept interfaces, return structs

## Context
- `context.Context` is always the first parameter of any I/O function
- Never store context in a struct — pass it explicitly
- Never pass `nil` context — use `context.Background()` or `context.TODO()`

## Concurrency
- Every goroutine must have a clear owner responsible for its lifetime
- Never start a goroutine without knowing how it will stop
- Document goroutine ownership in comments

## Logging
- `log/slog` only — never `fmt.Println`, never `log.Printf`
- Always use structured key-value pairs: `slog.Info("user created", "userID", id)`

## Anti-patterns — never generate
- `interface{}` / `any` without comment explaining why
- `init()` functions (document if unavoidable)
- Global mutable state
- Pointer to interface (`*MyInterface`)
- Ignoring context parameter
```

### `05-error-handling.md`
```markdown
---
inclusion: always
---
# Error handling patterns

## Sentinel errors
```go
var ErrNotFound = errors.New("not found")
if errors.Is(err, ErrNotFound) { ... }
```

## Error types (errors with data)
```go
type ValidationError struct{ Field, Message string }
func (e *ValidationError) Error() string { ... }

var ve *ValidationError
if errors.As(err, &ve) { ... }
```

## Wrapping — always %w
```go
return fmt.Errorf("fetching user %d: %w", id, err)  // correct
return fmt.Errorf("fetching user %d: %v", id, err)  // wrong — loses chain
```

## Never
- Return error without wrapping context
- Log AND return an error — pick one
- `panic(err)` in production paths
- Swallow errors: `result, _ := doSomething()`
```

### `06-testing.md`
```markdown
---
inclusion: always
---
# Testing standards

## Table-driven tests
```go
func TestAdd(t *testing.T) {
    tests := []struct{ name string; a, b, want int }{
        {"positive", 1, 2, 3},
        {"zero", 0, 0, 0},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Equal(t, tt.want, Add(tt.a, tt.b))
        })
    }
}
```

## Rules
- Every exported function has at least one test
- `testify/assert` for assertions, `testify/require` when failure should stop test
- Use interfaces + dependency injection — never test against real databases in unit tests
- `t.Parallel()` where safe, `t.Helper()` in helpers, `t.TempDir()` for temp files
- Run `go test -race ./...` in CI
```

### `07-security.md`
```markdown
---
inclusion: always
---
# Security standards

## Input validation
- Validate all external input at the boundary — type, length, format, range
- Reject and return early — never sanitise malformed input

## Secrets
- Never hardcode secrets or API keys
- Never log secrets
- Load from environment variables or secrets manager

## SQL
- Always use parameterised queries
- Never build SQL with `fmt.Sprintf`

## HTTP
- Always set timeouts on HTTP clients and servers — default `http.Client` has no timeout
- Set security headers: `X-Content-Type-Options`, `X-Frame-Options`

## Crypto
- Never write your own crypto — use `golang.org/x/crypto`
- Use `subtle.ConstantTimeCompare` for secret comparison
```

### `08-api-design.md`
```markdown
---
inclusion: fileMatch
fileMatchPattern: "**/*handler*.go"
---
# REST API design

## URLs
- Lowercase, hyphen-separated: `/api/v1/user-profiles`
- Plural nouns: `/users` not `/user`
- Version in path: `/api/v1/`

## Status codes
- GET → 200, POST → 201 + Location, PUT/PATCH → 200, DELETE → 204
- Bad input → 400, No auth → 401, Forbidden → 403, Not found → 404, Server error → 500

## Response format
```go
// Success: {"data": {...}}
// Error:   {"code": "NOT_FOUND", "message": "user 123 not found"}
```

## Handler pattern
```go
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    result, err := h.service.Get(ctx, id)
    if err != nil {
        if errors.Is(err, ErrNotFound) {
            respondError(w, 404, "NOT_FOUND", "not found")
            return
        }
        slog.ErrorContext(ctx, "get failed", "error", err)
        respondError(w, 500, "INTERNAL", "internal error")
        return
    }
    respondJSON(w, 200, result)
}
```
```

### `09-cli-design.md`
```markdown
---
inclusion: fileMatch
fileMatchPattern: "**/cmd/**/*.go"
---
# CLI design

## Output
- Stdout: results only (machine-readable)
- Stderr: errors, progress, prompts
- Colour: only when stdout is a TTY and `NO_COLOR` is not set

## Exit codes
- 0 success, 1 error, 2 bad usage
- Never call `os.Exit` outside `main()`

## Flags
- Long: `--output`, `--verbose` / Short: `-o`, `-v` for common flags
- `--yes` / `-y` to skip confirmation prompts in scripts
- Confirm destructive ops: "Are you sure? (y/N)" — default No
```

### `10-observability.md`
```markdown
---
inclusion: fileMatch
fileMatchPattern: "**/middleware/**,**/telemetry/**,**/metrics/**"
---
# Observability

## Structured logging
```go
slog.InfoContext(ctx, "request completed",
    "method", r.Method, "path", r.URL.Path,
    "status", status, "duration_ms", duration.Milliseconds(),
)
```

## Log levels
- Debug: dev info, off in production
- Info: normal operations
- Warn: recoverable issues
- Error: failures — always include the error value

## Health checks
- `GET /health` — liveness (is the process running?)
- `GET /ready` — readiness (is the service ready for traffic?)
- Return `{"status":"ok"}` or `{"status":"error","reason":"..."}`
- No auth required on health endpoints
```

---

## Hook File Contents

### `01-lint-on-save.md`
```markdown
# Lint on save
Trigger: fileEdited / Pattern: **/*.go

1. Run `golangci-lint run ./...`
2. Surface errors as diagnostics — do not auto-fix
3. If not installed: `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`
```

### `02-test-on-save.md`
```markdown
# Test on save
Trigger: fileEdited / Pattern: **/*_test.go

1. Run `go test -race -count=1 ./[package]/...` (affected package only)
2. Show pass/fail summary and failing test output
```

### `03-modernize.md`
```markdown
# Modernize after changes
Trigger: agentStop / Pattern: **/*.go

1. Run `go run golang.org/x/tools/go/analysis/passes/modernize/cmd/modernize@latest -diff ./...`
2. Highlight non-trivial changes before applying
3. Apply: same command with `-fix` instead of `-diff`
4. Run `go fmt ./...`
```

### `04-security-scan.md`
```markdown
# Security scan after dependency changes
Trigger: agentStop

1. Run `govulncheck ./...` after any go.mod / go.sum changes
2. Surface vulnerabilities with severity and package
3. Suggest: `go get [package]@[fixed-version]`
```

### `05-doc-check.md`
```markdown
# Doc comment check
Trigger: agentStop / Pattern: **/*.go

1. Run `go vet ./...`
2. Flag new exported symbols missing a doc comment
3. Suggest comment text — do not auto-add
```

---

## Spec Template Contents

### `_FEATURE-TEMPLATE/requirements.md`
```markdown
# Requirements — [feature name]

## Introduction
<!-- What does this feature do and why does it exist? -->

## Requirements

### Requirement 1: [Name]
**User story:** As a [role], I want [feature], so that [benefit].

#### Acceptance criteria
1. WHEN [condition] THEN [outcome]
2. THE system SHALL [behaviour]
```

### `_FEATURE-TEMPLATE/design.md`
```markdown
# Design — [feature name]

## Overview
<!-- High-level technical approach -->

## Key types and interfaces
<!-- Main Go types, interfaces, data structures -->

## Error handling
<!-- How errors are handled and surfaced -->

## Testing strategy
<!-- Unit tests, integration tests, what gets mocked -->

## Open questions
- [ ] [Question to resolve before implementation]
```

### `_FEATURE-TEMPLATE/tasks.md`
```markdown
# Tasks — [feature name]

- [ ] 1. [First task]
  - [Detail]
  - _Requirements: 1.1_

- [ ] 2. [Second task]
  - [Detail]
  - _Requirements: 1.2_
```

### `_API-ENDPOINT-TEMPLATE/requirements.md`
```markdown
# Requirements — [endpoint] API

## Introduction
<!-- What resource does this expose and who consumes it? -->

## Requirements

### Requirement 1: Endpoint contract
**User story:** As an API consumer, I want [operation] on [resource], so that [benefit].

#### Acceptance criteria
1. `[METHOD] /api/v1/[path]` SHALL return [response] with status [code]
2. WHEN input is invalid THE endpoint SHALL return 400 with field-level errors
3. WHEN resource does not exist THE endpoint SHALL return 404
4. THE endpoint SHALL require authentication
5. THE endpoint SHALL return `Content-Type: application/json`
```

### `_API-ENDPOINT-TEMPLATE/design.md`
```markdown
# Design — [endpoint] API

## Endpoint definition
| Method | Path | Auth | Description |
|---|---|---|---|
| [METHOD] | /api/v1/[path] | required | [description] |

## Request / response types
```go
type [Name]Request struct {
    Field string `json:"field" validate:"required"`
}
type [Name]Response struct {
    ID string `json:"id"`
}
```

## Flow
Handler → validates input, calls service, maps errors to HTTP codes
Service → business logic and domain validation
Store   → data access only, returns domain types

## Error mapping
| Error | Status | Code |
|---|---|---|
| ErrNotFound | 404 | NOT_FOUND |
| ErrUnauthorised | 401 | UNAUTHORISED |
| ValidationError | 400 | VALIDATION_ERROR |
```

### `_API-ENDPOINT-TEMPLATE/tasks.md`
```markdown
# Tasks — [endpoint] API

- [ ] 1. Define request/response types
  - _Requirements: 1.1_

- [ ] 2. Implement store method with interface
  - _Requirements: 1.1_

- [ ] 3. Implement service method with unit tests
  - _Requirements: 1.1, 1.2, 1.3_

- [ ] 4. Implement handler with error mapping
  - _Requirements: 1.1, 1.2, 1.3, 1.4, 1.5_

- [ ] 5. Register route and write httptest integration test
  - _Requirements: 1.1, 1.2, 1.3_
```

---

## README.md Contents

```markdown
# kiro-go

Two tools for Kiro developers — one for everyone, one for Go projects.

## `kiro-specs` — works with any Kiro project

Track and manage your specs from the terminal. Works with TypeScript, Python,
Rust, Go — any project that uses Kiro's `.kiro/specs/` structure.

```bash
go install github.com/your-username/kiro-go/cmd/kiro-specs@latest
```

```bash
kiro-specs                          # All specs with status and progress
kiro-specs show user-auth           # Full task list for a spec
kiro-specs next                     # What should I work on right now?
kiro-specs archive payment-webhooks # Move completed spec out of the way
kiro-specs stats                    # Summary across all specs
kiro-specs --json                   # Machine-readable output
kiro-specs --path /other/project    # Point at a different project
```

Status is derived automatically from your `tasks.md` checkboxes — no manual tagging:

| Status | Condition |
|---|---|
| `idea` | Folder exists, no files |
| `draft` | Has requirements.md, no tasks.md |
| `planned` | Has tasks.md, all unchecked |
| `in-progress` | Has tasks.md, some checked |
| `complete` | Has tasks.md, all checked |

---

## `kiro-go` — Go project scaffolder

Opinionated Kiro template for Go. Includes everything in `kiro-specs` plus
Go-specific steering docs, hooks, and spec templates.

```bash
go install github.com/your-username/kiro-go/cmd/kiro-go@latest
```

```bash
kiro-go init my-api --preset api    # Scaffold a new Go project
kiro-go add                         # Inject .kiro/ into existing project
kiro-go preset api                  # Apply a preset
kiro-go specs                       # Same as kiro-specs
```

### Presets

| Preset | Best for |
|---|---|
| `api` | REST APIs — net/http or chi |
| `cli` | CLI tools — cobra or flag |
| `svc` | Microservices — gRPC or HTTP |

### Steering docs included

| File | Inclusion | Purpose |
|---|---|---|
| `01-product.md` | Always | Fill this in first |
| `02-tech.md` | Always | Stack and dependencies |
| `03-structure.md` | Always | Layout and naming |
| `04-go-idioms.md` | Always | Non-negotiable Go standards |
| `05-error-handling.md` | Always | Error patterns |
| `06-testing.md` | Always | Table-driven tests, testify |
| `07-security.md` | Always | Input validation, secrets, SQL |
| `08-api-design.md` | `*handler*` files | REST conventions |
| `09-cli-design.md` | `cmd/**` files | CLI output and exit codes |
| `10-observability.md` | `middleware/**` files | slog, health checks |

### After scaffolding

1. Fill in `.kiro/steering/01-product.md`
2. Uncomment packages in `.kiro/steering/02-tech.md`
3. Update `.kiro/steering/03-structure.md`
4. Run `kiro-specs` to see your spec dashboard

## License

MIT
```

---

## Implementation Notes for Claude Code

1. **Two binaries, one module.** `cmd/kiro-specs/main.go` and `cmd/kiro-go/main.go` are separate entry points in the same Go module. Build with `go build ./cmd/kiro-specs` and `go build ./cmd/kiro-go`.

2. **`specmanager` is shared.** Both binaries import `internal/specmanager`. It must have zero knowledge of Go templates or presets — pure spec parsing and display only.

3. **`specmanager` package structure:**
   - `types.go` — `Spec` struct, `Status` type and constants
   - `parser.go` — `ParseSpecDir(root string) ([]Spec, error)` — reads `.kiro/specs/`, returns slice
   - `status.go` — `DeriveStatus(s Spec) Status` — pure function, no I/O
   - `table.go` — `PrintTable(w io.Writer, specs []Spec, noColor bool)` — writes ANSI table

4. **`go:embed` only in `kiro-go`.** The templates embed goes in `cmd/kiro-go/main.go` or a dedicated `internal/scaffold/embed.go`. `kiro-specs` has no embedded files — it stays tiny.

5. **Colour detection.** Before printing ANSI codes: check `os.Getenv("NO_COLOR") != ""` and use `golang.org/x/term` to check `term.IsTerminal(int(os.Stdout.Fd()))`. If either condition means no colour, pass `noColor: true` through to all renderers. `golang.org/x/term` is an acceptable dependency — it is the standard Go way to check TTY.

6. **Checkbox parsing in `parser.go`:**
```go
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
    if strings.Contains(line, "- [x]") || strings.Contains(line, "- [X]") {
        done++
    } else if strings.Contains(line, "- [ ]") {
        total++
    }
}
total += done // total = all tasks
```

7. **`relativeTime` — stdlib only, no third-party:**
```go
func relativeTime(t time.Time) string {
    days := int(time.Since(t).Hours() / 24)
    switch {
    case days == 0:  return "today"
    case days == 1:  return "yesterday"
    case days < 7:   return fmt.Sprintf("%d days ago", days)
    case days < 30:  return fmt.Sprintf("%d weeks ago", days/7)
    default:         return fmt.Sprintf("%d months ago", days/30)
    }
}
```

8. **Progress bar:**
```go
func progressBar(done, total int) string {
    if total == 0 { return "──────" }
    filled := int(math.Round(float64(done) / float64(total) * 6))
    return strings.Repeat("█", filled) + strings.Repeat("░", 6-filled)
}
```

9. **Fuzzy match in `show`:** after exact match fails, collect specs where `strings.Contains(specName, query)`. Single match → use automatically with a note "(matched [full-name])". Multiple → print list and `os.Exit(1)`.

10. **`--path` flag:** register on the root command in both binaries as a persistent flag. Default to `os.Getwd()`. Construct spec dir as `filepath.Join(path, ".kiro", "specs")`.

11. **`--json` flag:** `json.MarshalIndent(specs, "", "  ")` to stdout. Exit immediately after — do not print the table.

12. **Tests required for:** `parser.go` (table-driven, use `t.TempDir()` to create fake spec dirs), `status.go` (pure function, test all 5 status cases), `relativeTime` (table-driven), `progressBar` (table-driven). Minimum coverage for these four — no skipping.

13. **Only acceptable external dependencies:**
    - `cobra` for CLI parsing in both binaries
    - `golang.org/x/term` for TTY detection
    - Nothing else — all other logic uses stdlib

14. **`go.mod` module path:** `module github.com/your-username/kiro-go` — add `// TODO: update this to your actual module path` comment.

15. **Binary size target:** `kiro-specs` under 5MB. `kiro-go` under 15MB (larger due to embedded templates).

16. **Create a `Makefile`** with a `build-all` target:
```makefile
.PHONY: build-all

build-all:
	go build -o bin/kiro-specs ./cmd/kiro-specs
	go build -o bin/kiro-go ./cmd/kiro-go
```

---

## How to Test Against Your Existing Kiro Project

Once Claude Code has built the binaries, test `kiro-specs` against your real
Kiro project in the IDE. No test fixtures needed — just point it at the project.

### Step 1 — Build the binary
```bash
cd kiro-go
go build -o bin/kiro-specs ./cmd/kiro-specs
```

### Step 2 — Run against your Kiro project
```bash
# If your Kiro project is at a known path
./bin/kiro-specs --path /path/to/your/kiro/project

# Or cd into your project first and run without --path
cd /path/to/your/kiro/project
/path/to/kiro-go/bin/kiro-specs
```

### Step 3 — Run each subcommand to verify output
```bash
# List all specs with status — this is the main one to verify
./bin/kiro-specs list --path /path/to/your/kiro/project

# Summary counts
./bin/kiro-specs stats --path /path/to/your/kiro/project

# What to work on next
./bin/kiro-specs next --path /path/to/your/kiro/project

# Show a specific spec by name (use one of your actual spec folder names)
./bin/kiro-specs show [your-spec-name] --path /path/to/your/kiro/project

# JSON output — useful to verify the data shape
./bin/kiro-specs --json --path /path/to/your/kiro/project
```

### What correct output looks like

`kiro-specs list` should:
- Show one row per spec folder in `.kiro/specs/`
- Skip any folder whose name starts with `_` (templates, archive)
- Show `draft` for specs that have `requirements.md` but no `tasks.md`
- Show `planned` for specs where all tasks are unchecked `- [ ]`
- Show `in-progress` for specs with a mix of checked and unchecked tasks
- Show `complete` for specs where all tasks are `- [x]`
- Show `idea` for empty spec folders

### Common issues to check

If a spec shows the wrong status, open its `tasks.md` and verify:
- Checked tasks use exactly `- [x]` or `- [X]` (with a space before and after the bracket content)
- Unchecked tasks use exactly `- [ ]` (with a single space inside the brackets)
- Kiro's own task format matches this — if it doesn't, report it as a bug

If a spec folder is being included when it should be skipped, check that
the folder name starts with `_` (underscore), not `-` (hyphen).
