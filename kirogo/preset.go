package kirogo

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewPresetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "preset [name]",
		Short: "Apply a preset (api, cli, svc)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			kiroDir := ".kiro"
			if _, err := os.Stat(kiroDir); os.IsNotExist(err) {
				return fmt.Errorf(".kiro/ not found — run 'kiro-go add' first")
			}
			if err := applyPreset(kiroDir, name); err != nil {
				return err
			}
			fmt.Printf("Applied preset %q to .kiro/steering/\n", name)
			return nil
		},
	}
}
