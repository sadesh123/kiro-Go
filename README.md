# kiro-go

Two terminal tools for Kiro developers — one for everyone, one for Go.

```
 ██╗  ██╗██╗██████╗  ██████╗     ███████╗██████╗ ███████╗ ██████╗███████╗
 ██║ ██╔╝██║██╔══██╗██╔═══██╗    ██╔════╝██╔══██╗██╔════╝██╔════╝██╔════╝
 █████╔╝ ██║██████╔╝██║   ██║    ███████╗██████╔╝█████╗  ██║     ███████╗
 ██╔═██╗ ██║██╔══██╗██║   ██║    ╚════██║██╔═══╝ ██╔══╝  ██║     ╚════██║
 ██║  ██╗██║██║  ██║╚██████╔╝    ███████║██║     ███████╗╚██████╗███████║
 ╚═╝  ╚═╝╚═╝╚═╝  ╚═╝ ╚═════╝    ╚══════╝╚═╝     ╚══════╝ ╚═════╝╚══════╝

 ██╗  ██╗██╗██████╗  ██████╗        ██████╗  ██████╗
 ██║ ██╔╝██║██╔══██╗██╔═══██╗      ██╔════╝ ██╔═══██╗
 █████╔╝ ██║██████╔╝██║   ██║█████╗██║  ███╗██║   ██║
 ██╔═██╗ ██║██╔══██╗██║   ██║╚════╝██║   ██║██║   ██║
 ██║  ██╗██║██║  ██║╚██████╔╝      ╚██████╔╝╚██████╔╝
 ╚═╝  ╚═╝╚═╝╚═╝  ╚═╝ ╚═════╝        ╚═════╝  ╚═════╝
```

---

## What is this?

Kiro is an AI-powered IDE that uses a spec-driven workflow. Before you build anything, you write a spec — a folder inside `.kiro/specs/` containing a `requirements.md`, a `design.md`, and a `tasks.md` full of checkboxes. Kiro works through those tasks as it builds your feature.

This repo ships two CLI tools that make working with Kiro better:

---

### `kiro-specs` — for every Kiro project

A terminal dashboard for your specs. See which features are in progress, how many tasks are done, and what to work on next — without opening the IDE.

Works with **any language** — TypeScript, Python, Rust, Go, anything that uses `.kiro/specs/`.

```bash
kiro-specs                        # dashboard view
kiro-specs next                   # what to work on right now
kiro-specs show my-feature        # full task list for one spec
kiro-specs stats                  # completion summary
```

→ [Full guide for kiro-specs](cmd/kiro-specs/README.md)

---

### `kiro-go` — for Go projects

An opinionated Go project scaffolder. Generates a full `.kiro/` folder with steering docs, hooks, and spec templates tuned for Go idioms — so Kiro generates idiomatic Go from day one. Includes `kiro-specs` built in.

```bash
kiro-go init my-api --preset api  # scaffold a new Go project
kiro-go add                       # inject .kiro/ into existing project
kiro-go specs next                # same as kiro-specs next
```

→ [Full guide for kiro-go](cmd/kiro-go/README.md)

---

## Install

### `kiro-specs` — any Kiro project
```bash
go install github.com/sadesh123/kiro-go/cmd/kiro-specs@latest
```

### `kiro-go` — Go projects (includes kiro-specs)
```bash
go install github.com/sadesh123/kiro-go/cmd/kiro-go@latest
```

---

## Build from source

Requires Go 1.23+.

```bash
git clone https://github.com/sadesh123/kiro-go
cd kiro-go
make build-all
```

Binaries are written to `bin/`.

---

## License

MIT
