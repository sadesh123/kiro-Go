package kirogo

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sadesh123/kiro-go/internal/presets"
)

func NewListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List available presets",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("%-10s %s\n", "PRESET", "DESCRIPTION")
			for _, p := range presets.All {
				fmt.Printf("%-10s %s\n", p.Name, p.Description)
			}
			return nil
		},
	}
}
