# kiro-go

```
 ██╗  ██╗██╗██████╗  ██████╗        ██████╗  ██████╗
 ██║ ██╔╝██║██╔══██╗██╔═══██╗      ██╔════╝ ██╔═══██╗
 █████╔╝ ██║██████╔╝██║   ██║█████╗██║  ███╗██║   ██║
 ██╔═██╗ ██║██╔══██╗██║   ██║╚════╝██║   ██║██║   ██║
 ██║  ██╗██║██║  ██║╚██████╔╝      ╚██████╔╝╚██████╔╝
 ╚═╝  ╚═╝╚═╝╚═╝  ╚═╝ ╚═════╝        ╚═════╝  ╚═════╝
```

An opinionated Go project scaffolder for Kiro. Generates a full `.kiro/` folder with steering docs, hooks, and spec templates tuned for Go idioms — so Kiro writes idiomatic Go from day one.

Includes `kiro-specs` built in. You don't need both installed.

---

## Install

```bash
go install github.com/sadesh123/kiro-go/cmd/kiro-go@latest
```

---

## Commands

### `init [name]` — Scaffold a new Go project

Creates a new directory, runs `go mod init`, drops in a minimal `main.go`, and copies the full `.kiro/` template. Prompts for your GitHub username and which preset to apply.

```bash
kiro-go init my-api
```

```
GitHub username or org: sadesh123
Which preset? (api / cli / svc / none): api

Scaffolded my-api with module github.com/sadesh123/my-api

Next steps:
  cd my-api
  Fill in .kiro/steering/01-product.md
  Uncomment packages in .kiro/steering/02-tech.md
  Run: kiro-specs
```

Pass `--preset` to skip the prompt:

```bash
kiro-go init my-api --preset api
kiro-go init my-cli --preset cli
kiro-go init my-svc --preset svc
```

---

### `add` — Inject `.kiro/` into an existing Go project

Run inside an existing Go project to add the full `.kiro/` folder without touching your code. Requires a `go.mod` in the current directory.

```bash
cd my-existing-api
kiro-go add
```

If `.kiro/` already exists you'll be asked before anything is overwritten.

---

### `preset [name]` — Switch preset on an existing project

Updates which steering docs are always-on vs manual for the chosen project type.

```bash
kiro-go preset api
kiro-go preset cli
kiro-go preset svc
```

---

### `list` — See available presets

```bash
kiro-go list
```

```
PRESET     DESCRIPTION
api        REST API — net/http or chi router, middleware, handlers, JSON responses
cli        CLI tool — cobra or flag, subcommands, config file, stdout/stderr hygiene
svc        Microservice — gRPC or HTTP, health checks, graceful shutdown, telemetry
```

---

### `specs` — Full kiro-specs dashboard

Identical to running `kiro-specs`. Included so you don't need a second binary.

```bash
kiro-go specs               # dashboard
kiro-go specs next          # what to work on now
kiro-go specs show my-feat  # task list for one spec
kiro-go specs stats         # completion summary
```

See the [kiro-specs guide](../kiro-specs/README.md) for the full command reference.

---

### `version`

```bash
kiro-go version
```

---

## What gets scaffolded

### Steering docs (`.kiro/steering/`)

Ten guides that Kiro reads as context when generating code. Each file has an `inclusion` field that controls when Kiro loads it.

| File | Purpose | Active by default |
|---|---|---|
| `01-product.md` | What the project does and who uses it | All presets |
| `02-tech.md` | Stack, dependencies, common commands | All presets |
| `03-structure.md` | Layout, naming conventions, package rules | All presets |
| `04-go-idioms.md` | Errors, interfaces, context, concurrency rules | All presets |
| `05-error-handling.md` | Sentinel errors, wrapping, what never to do | All presets |
| `06-testing.md` | Table-driven tests, testify, race detector | All presets |
| `07-security.md` | Input validation, secrets, SQL, HTTP timeouts | All presets |
| `08-api-design.md` | REST conventions, status codes, handler pattern | `api`, `svc` |
| `09-cli-design.md` | stdout/stderr hygiene, exit codes, flag conventions | `cli` |
| `10-observability.md` | slog structured logging, health check endpoints | `api`, `svc` |

Fill in `01-product.md` first — it's the most important one. Uncomment the packages you're actually using in `02-tech.md`.

---

### Hooks (`.kiro/hooks/`)

Five hooks that Kiro runs automatically as you work:

| Hook | Trigger | What it does |
|---|---|---|
| `01-lint-on-save.md` | Any `.go` file saved | Runs `golangci-lint run ./...` |
| `02-test-on-save.md` | Any `_test.go` file saved | Runs `go test -race` on the affected package |
| `03-modernize.md` | Agent stops | Runs the Go modernize tool, then `go fmt` |
| `04-security-scan.md` | Agent stops | Runs `govulncheck` after `go.mod` changes |
| `05-doc-check.md` | Agent stops | Flags new exported symbols missing a doc comment |

---

### Spec templates (`.kiro/specs/`)

Two starter templates to copy when creating a new spec in Kiro:

**`_FEATURE-TEMPLATE/`** — for any new feature
- `requirements.md` — user story and acceptance criteria
- `design.md` — types, interfaces, error handling, testing strategy
- `tasks.md` — checkbox task list

**`_API-ENDPOINT-TEMPLATE/`** — for a new REST endpoint
- `requirements.md` — endpoint contract and acceptance criteria
- `design.md` — method/path table, request/response types, handler → service → store flow, error mapping
- `tasks.md` — five standard tasks: types, store, service, handler, route + integration test

---

## Presets in detail

Presets control which steering docs have `inclusion: always` (loaded by Kiro for every file) vs `inclusion: manual` (loaded only when you ask).

### `api` — REST API
Always on: product, tech, structure, Go idioms, errors, testing, security, API design, observability
Manual: CLI design

### `cli` — CLI tool
Always on: product, tech, structure, Go idioms, errors, testing, security, CLI design
Manual: API design, observability

### `svc` — Microservice
Always on: product, tech, structure, Go idioms, errors, testing, security, API design, observability
Manual: CLI design

---

## After scaffolding

1. Fill in `.kiro/steering/01-product.md` — describe what you're building
2. Uncomment the packages you'll use in `.kiro/steering/02-tech.md`
3. Update `.kiro/steering/03-structure.md` to match your actual layout
4. Run `kiro-go specs` to see your spec dashboard

---

## License

MIT
