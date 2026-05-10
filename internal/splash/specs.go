package splash

import "fmt"

const specsArt = "\033[96m" + ` ██╗  ██╗██╗██████╗  ██████╗     ███████╗██████╗ ███████╗ ██████╗███████╗
 ██║ ██╔╝██║██╔══██╗██╔═══██╗    ██╔════╝██╔══██╗██╔════╝██╔════╝██╔════╝
 █████╔╝ ██║██████╔╝██║   ██║    ███████╗██████╔╝█████╗  ██║     ███████╗
 ██╔═██╗ ██║██╔══██╗██║   ██║    ╚════██║██╔═══╝ ██╔══╝  ██║     ╚════██║
 ██║  ██╗██║██║  ██║╚██████╔╝    ███████║██║     ███████╗╚██████╗███████║
 ╚═╝  ╚═╝╚═╝╚═╝  ╚═╝ ╚═════╝    ╚══════╝╚═╝     ╚══════╝ ╚═════╝╚══════╝` + "\033[0m"

func PrintSpecs() {
	fmt.Println(specsArt)
	fmt.Println()
	fmt.Println("v0.1.0  ·  works with any Kiro project  ·  language-agnostic")
	fmt.Println()
	fmt.Println("Spec manager for Kiro — track status, find what to work on next, archive completed work.")
	fmt.Println()
	fmt.Println("────────────────────────────────────────────────────────────────────")
	fmt.Println()
	fmt.Println("  list               List all specs with status and progress  (default)")
	fmt.Println("  show [name]        Show task list for a spec")
	fmt.Println("  next               What should I work on right now?")
	fmt.Println("  archive [name]     Move a completed spec out of the way")
	fmt.Println("  stats              Summary stats across all specs")
	fmt.Println()
	fmt.Println("  Flags (all commands):")
	fmt.Println("  --status [status]  Filter by: idea, draft, planned, in-progress, complete")
	fmt.Println("  --sort [by]        Sort by: status (default), age, name")
	fmt.Println("  --json             Machine-readable JSON output")
	fmt.Println("  --path [dir]       Path to project root (default: current directory)")
	fmt.Println()
}
