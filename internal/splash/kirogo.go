package splash

import "fmt"

const kiroGoArt = "\033[96m" + ` ██╗  ██╗██╗██████╗  ██████╗        ██████╗  ██████╗
 ██║ ██╔╝██║██╔══██╗██╔═══██╗      ██╔════╝ ██╔═══██╗
 █████╔╝ ██║██████╔╝██║   ██║█████╗██║  ███╗██║   ██║
 ██╔═██╗ ██║██╔══██╗██║   ██║╚════╝██║   ██║██║   ██║
 ██║  ██╗██║██║  ██║╚██████╔╝      ╚██████╔╝╚██████╔╝
 ╚═╝  ╚═╝╚═╝╚═╝  ╚═╝ ╚═════╝        ╚═════╝  ╚═════╝` + "\033[0m"

func PrintKiroGo() {
	fmt.Println(kiroGoArt)
	fmt.Println()
	fmt.Printf("\033[34m v0.1.0\033[0m  \033[32m Go 1.23+\033[0m  \033[37m spec-driven\033[0m  \033[96m[ GO ]\033[0m")
	fmt.Println()
	fmt.Println()
	fmt.Println("Opinionated Kiro template for Go — spec-driven, idiomatic, production-ready.")
	fmt.Println("Steering docs · Hooks · REST API · CLI · Microservice presets")
	fmt.Println()
	fmt.Println("────────────────────────────────────────────────────────────────────")
	fmt.Println()
	fmt.Println("  init [name]     Scaffold a new Go project")
	fmt.Println("  add             Inject .kiro/ into an existing project")
	fmt.Println("  preset [name]   Apply a preset (api, cli, svc)")
	fmt.Println("  list            List available presets")
	fmt.Println("  specs           Manage specs (powered by kiro-specs)")
	fmt.Println("  version         Show version info")
	fmt.Println("  help            Show this help screen")
	fmt.Println()
}
