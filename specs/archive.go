package specs

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sadesh123/kiro-go/internal/specmanager"
)

func NewArchiveCmd(pathFlag *string) *cobra.Command {
	var yes bool

	cmd := &cobra.Command{
		Use:   "archive [name]",
		Short: "Move a completed spec out of the way",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			root := resolveRoot(*pathFlag)
			specs, err := specmanager.ParseSpecDir(root)
			if err != nil {
				return fmt.Errorf("reading specs: %w", err)
			}

			name := args[0]
			spec, err := findSpec(specs, name)
			if err != nil {
				return err
			}

			if spec.Status != specmanager.StatusComplete && !yes {
				fmt.Fprintf(os.Stderr, "This spec is not complete (status: %s). Archive anyway? (y/N) ", spec.Status)
				reader := bufio.NewReader(os.Stdin)
				answer, _ := reader.ReadString('\n')
				answer = strings.TrimSpace(strings.ToLower(answer))
				if answer != "y" && answer != "yes" {
					fmt.Println("Aborted.")
					return nil
				}
			}

			specsDir := filepath.Join(root, ".kiro", "specs")
			archiveDir := filepath.Join(specsDir, "_archive")
			if err := os.MkdirAll(archiveDir, 0755); err != nil {
				return fmt.Errorf("creating archive dir: %w", err)
			}

			src := filepath.Join(specsDir, spec.Name)
			dst := filepath.Join(archiveDir, spec.Name)
			if err := os.Rename(src, dst); err != nil {
				return fmt.Errorf("moving spec: %w", err)
			}

			fmt.Printf("Archived %s → .kiro/specs/_archive/%s\n", spec.Name, spec.Name)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&yes, "yes", "y", false, "Skip confirmation prompt")
	return cmd
}
