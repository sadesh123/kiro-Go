# kiro-specs

```
 ██╗  ██╗██╗██████╗  ██████╗     ███████╗██████╗ ███████╗ ██████╗███████╗
 ██║ ██╔╝██║██╔══██╗██╔═══██╗    ██╔════╝██╔══██╗██╔════╝██╔════╝██╔════╝
 █████╔╝ ██║██████╔╝██║   ██║    ███████╗██████╔╝█████╗  ██║     ███████╗
 ██╔═██╗ ██║██╔══██╗██║   ██║    ╚════██║██╔═══╝ ██╔══╝  ██║     ╚════██║
 ██║  ██╗██║██║  ██║╚██████╔╝    ███████║██║     ███████╗╚██████╗███████║
 ╚═╝  ╚═╝╚═╝╚═╝  ╚═╝ ╚═════╝    ╚══════╝╚═╝     ╚══════╝ ╚═════╝╚══════╝
```

A terminal dashboard for your Kiro specs. See what's in progress, how many tasks are done, and what to work on next — without opening the IDE.

Works with **any Kiro project** regardless of language — TypeScript, Python, Go, Rust, anything that uses `.kiro/specs/`.

---

## Setup

### macOS

**Step 1 — Download the binary for your chip**

```bash
# Apple Silicon (M1/M2/M3/M4) — check with: uname -m
curl -L https://github.com/sadesh123/kiro-go/releases/latest/download/kiro-specs-darwin-arm64 -o kiro-specs

# Intel Mac
curl -L https://github.com/sadesh123/kiro-go/releases/latest/download/kiro-specs-darwin-amd64 -o kiro-specs
```

Not sure which chip? Run `uname -m` — `arm64` is Apple Silicon, `x86_64` is Intel.

**Step 2 — Make it executable and move to your PATH**

```bash
chmod +x kiro-specs
mv kiro-specs /usr/local/bin/kiro-specs
```

**Step 3 — Allow it in Security settings**

macOS will block it the first time since it isn't signed. Two ways to fix this:

- Open **System Settings → Privacy & Security**, scroll down, click **Allow Anyway**
- Or run: `xattr -d com.apple.quarantine /usr/local/bin/kiro-specs`

**Step 4 — Run it**

```bash
cd ~/my-kiro-project
kiro-specs

# or from anywhere
kiro-specs --path ~/my-kiro-project
```

---

### Windows WSL

**Step 1 — Download the Linux binary**

Download `kiro-specs-linux` from the [releases page](https://github.com/sadesh123/kiro-go/releases) and place it somewhere on your Windows filesystem, e.g. `C:\Users\YourName\tools\`.

**Step 2 — Add it to your WSL PATH**

```bash
echo 'export PATH=$PATH:/mnt/c/Users/YourName/tools' >> ~/.bashrc
source ~/.bashrc
```

Replace `YourName` with your actual Windows username.

**Step 3 — Make it executable**

```bash
chmod +x /mnt/c/Users/YourName/tools/kiro-specs-linux
```

**Step 4 — Run it**

Your Kiro projects live under `/mnt/c/` in WSL. Use `--path` to point at them:

```bash
# From anywhere
kiro-specs-linux --path /mnt/c/Users/YourName/Desktop/my-kiro-project

# Or cd in first
cd /mnt/c/Users/YourName/Desktop/my-kiro-project
kiro-specs-linux
```

> The binary is named `kiro-specs-linux` on WSL. All commands and flags are identical — just swap `kiro-specs` for `kiro-specs-linux` in the examples below.

---

## Commands

### `list` — All specs at a glance

The default view. Run with no arguments or explicitly with `list`.

```bash
kiro-specs
kiro-specs list
```

```
  SPEC                           STATUS         PROGRESS       AGE
  ──────────────────────────────────────────────────────────────────
  kiro-workshop-resume           in-progress    ██░░░░ 19/52   2 days ago

  1 specs  |  0 complete  |  1 in-progress  |  0 planned  |  0 draft  |  0 idea
```

Status is derived automatically from your `tasks.md` checkboxes — no manual tagging needed:

| Status | Condition |
|---|---|
| `idea` | Folder exists, no files yet |
| `draft` | Has `requirements.md`, no `tasks.md` |
| `planned` | Has `tasks.md`, nothing checked off |
| `in-progress` | Has `tasks.md`, some checked off |
| `complete` | Has `tasks.md`, everything checked off |

Progress (`██░░░░ 19/52`) comes from counting `- [x]` and `- [ ]` lines in `tasks.md`.

---

### `show [name]` — Full task list for a spec

```bash
kiro-specs show kiro-workshop-resume
```

```
  kiro-workshop-resume  ·  in-progress  ·  19/52 tasks done  ·  2 days ago
  ──────────────────────────────────────────────────────────────────────

  [x] 1. Scaffold Astro project structure and configuration
  [x] 2. Define TypeScript data model and placeholder content
  [ ] 3. Create global CSS design tokens and layout shell
  [ ] 4. Implement SVG icon components
  ...

  Next up: Create global CSS design tokens and layout shell
```

Fuzzy match — you don't need the full name:

```bash
kiro-specs show workshop    # matches kiro-workshop-resume automatically
```

---

### `next` — What should I work on right now?

Picks the best spec to focus on. Prioritises the in-progress spec closest to completion. Falls back to the oldest planned spec if nothing is in progress.

```bash
kiro-specs next
```

```
  Pick up where you left off:

  kiro-workshop-resume  ·  in-progress  ·  33 tasks remaining

    - [ ] 3. Create global CSS design tokens and layout shell
      - [ ] 3.1 Create `src/styles/global.css` with all CSS custom properties
    - [ ] 4. Implement SVG icon components
    - [ ] 9. Implement smoke and integration tests
    - [ ] 10. Final checkpoint — Ensure all tests pass

  Run: kiro-specs show kiro-workshop-resume
```

---

### `stats` — Completion summary across all specs

```bash
kiro-specs stats
```

```
  Spec summary
  ─────────────────────────────
  Total specs          12
  Complete              4  (33%)
  In progress           3
  Planned               3
  Draft                 1
  Ideas                 1

  Tasks completed      31 / 67  (46%)
  Oldest active        middleware-refactor  (3 months ago)
  Most recent          kiro-workshop-resume  (today)
```

---

### `archive [name]` — Move a completed spec out of the way

Moves the spec to `.kiro/specs/_archive/` so it no longer appears in your dashboard.

```bash
kiro-specs archive kiro-workshop-resume
```

If the spec isn't complete, you'll get a confirmation prompt:

```
This spec is not complete (status: in-progress). Archive anyway? (y/N)
```

Skip the prompt with `-y`:

```bash
kiro-specs archive kiro-workshop-resume -y
```

---

## Flags

| Flag | Description |
|---|---|
| `--status [status]` | Filter by: `idea`, `draft`, `planned`, `in-progress`, `complete` |
| `--sort [by]` | Sort by: `status` (default — in-progress first), `age`, `name` |
| `--json` | Output JSON instead of the table |
| `--path [dir]` | Point at a project root other than the current directory |

All flags work with all commands.

### `--status` — Filter to one status

```bash
kiro-specs --status in-progress
kiro-specs --status planned
kiro-specs --status complete
```

Returns an empty table if nothing matches that status.

### `--sort age` — Sort by most recently modified

```bash
kiro-specs --sort age    # newest first
kiro-specs --sort name   # alphabetical
```

### `--json` — Machine-readable output

```bash
kiro-specs --json
```

```json
[
  {
    "name": "kiro-workshop-resume",
    "status": "in-progress",
    "tasks_total": 52,
    "tasks_done": 19,
    "age_days": 2,
    "path": ".kiro/specs/kiro-workshop-resume"
  }
]
```

Combine with `jq`:

```bash
# Names of all in-progress specs
kiro-specs --json | jq '.[] | select(.status == "in-progress") | .name'

# Overall completion percentage
kiro-specs --json | jq '[.[] | .tasks_done] | add / ([.[] | .tasks_total] | add) * 100'
```

---

## License

MIT
