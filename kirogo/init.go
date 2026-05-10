package kirogo

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sadesh123/kiro-go/internal/presets"
	"github.com/sadesh123/kiro-go/internal/scaffold"
)

func NewInitCmd(templates fs.FS) *cobra.Command {
	var presetName string

	cmd := &cobra.Command{
		Use:   "init [name]",
		Short: "Scaffold a new Go project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			if err := os.MkdirAll(name, 0755); err != nil {
				return fmt.Errorf("creating directory: %w", err)
			}

			reader := bufio.NewReader(os.Stdin)
			fmt.Print("GitHub username or org: ")
			username, _ := reader.ReadString('\n')
			username = strings.TrimSpace(username)

			module := fmt.Sprintf("github.com/%s/%s", username, name)
			gomod := exec.Command("go", "mod", "init", module)
			gomod.Dir = name
			gomod.Stdout = os.Stdout
			gomod.Stderr = os.Stderr
			if err := gomod.Run(); err != nil {
				return fmt.Errorf("go mod init: %w", err)
			}

			mainContent := fmt.Sprintf("package main\n\nfunc main() {\n}\n")
			if err := os.WriteFile(fmt.Sprintf("%s/main.go", name), []byte(mainContent), 0644); err != nil {
				return fmt.Errorf("creating main.go: %w", err)
			}

			kiroDir := fmt.Sprintf("%s/.kiro", name)
			if err := scaffold.CopyDir(templates, "templates/templates/.kiro", kiroDir); err != nil {
				return fmt.Errorf("copying templates: %w", err)
			}

			if presetName == "" {
				fmt.Print("Which preset? (api / cli / svc / none): ")
				input, _ := reader.ReadString('\n')
				presetName = strings.TrimSpace(input)
			}

			if presetName != "" && presetName != "none" {
				if err := applyPreset(kiroDir, presetName); err != nil {
					fmt.Fprintf(os.Stderr, "Warning: could not apply preset %q: %v\n", presetName, err)
				}
			}

			fmt.Printf("\nScaffolded %s with module %s\n", name, module)
			fmt.Println("\nNext steps:")
			fmt.Printf("  cd %s\n", name)
			fmt.Println("  Fill in .kiro/steering/01-product.md")
			fmt.Println("  Uncomment packages in .kiro/steering/02-tech.md")
			fmt.Println("  Run: kiro-specs")
			return nil
		},
	}

	cmd.Flags().StringVar(&presetName, "preset", "", "Preset to apply: api, cli, svc")
	return cmd
}

func applyPreset(kiroDir, name string) error {
	p, ok := presets.Find(name)
	if !ok {
		return fmt.Errorf("unknown preset %q — available: api, cli, svc", name)
	}

	steeringDir := fmt.Sprintf("%s/steering", kiroDir)
	entries, err := os.ReadDir(steeringDir)
	if err != nil {
		return err
	}

	alwaysSet := map[string]bool{}
	for _, n := range p.AlwaysOn {
		alwaysSet[n] = true
	}

	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		fname := e.Name()
		num := fname[:2]
		path := fmt.Sprintf("%s/%s", steeringDir, fname)

		content, err := os.ReadFile(path)
		if err != nil {
			continue
		}

		var inclusion string
		if alwaysSet[num] {
			inclusion = "always"
		} else {
			inclusion = "manual"
		}

		updated := replaceInclusion(string(content), inclusion)
		os.WriteFile(path, []byte(updated), 0644)
	}
	return nil
}

func replaceInclusion(content, inclusion string) string {
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "inclusion:") {
			lines[i] = "inclusion: " + inclusion
		}
	}
	return strings.Join(lines, "\n")
}
