package kirogo

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sadesh123/kiro-go/internal/scaffold"
)

func NewAddCmd(templates fs.FS) *cobra.Command {
	var presetName string

	return &cobra.Command{
		Use:   "add",
		Short: "Inject .kiro/ into an existing project",
		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			if err != nil {
				return err
			}

			if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
				fmt.Fprintln(os.Stderr, "Error: no go.mod found in current directory. Run 'go mod init' first.")
				os.Exit(1)
			}

			kiroDir := fmt.Sprintf("%s/.kiro", cwd)
			if _, err := os.Stat(kiroDir); err == nil {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print(".kiro/ already exists. Overwrite? (y/N): ")
				answer, _ := reader.ReadString('\n')
				answer = strings.TrimSpace(strings.ToLower(answer))
				if answer != "y" && answer != "yes" {
					fmt.Println("Aborted.")
					return nil
				}
			}

			if err := scaffold.CopyDir(templates, "templates/templates/.kiro", kiroDir); err != nil {
				return fmt.Errorf("copying templates: %w", err)
			}

			if presetName == "" {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("Which preset? (api / cli / svc / none): ")
				input, _ := reader.ReadString('\n')
				presetName = strings.TrimSpace(input)
			}

			if presetName != "" && presetName != "none" {
				if err := applyPreset(kiroDir, presetName); err != nil {
					fmt.Fprintf(os.Stderr, "Warning: could not apply preset %q: %v\n", presetName, err)
				}
			}

			fmt.Println("Injected .kiro/ into current project.")
			fmt.Println("\nNext steps:")
			fmt.Println("  Fill in .kiro/steering/01-product.md")
			fmt.Println("  Uncomment packages in .kiro/steering/02-tech.md")
			fmt.Println("  Run: kiro-specs")
			return nil
		},
	}
}
