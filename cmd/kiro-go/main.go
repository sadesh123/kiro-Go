package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sadesh123/kiro-go/internal/scaffold"
	"github.com/sadesh123/kiro-go/internal/splash"
	"github.com/sadesh123/kiro-go/kirogo"
	"github.com/sadesh123/kiro-go/specs"
)

func main() {
	var pathFlag string

	root := &cobra.Command{
		Use:   "kiro-go",
		Short: "Opinionated Kiro template for Go",
		RunE: func(cmd *cobra.Command, args []string) error {
			splash.PrintKiroGo()
			return nil
		},
	}

	root.PersistentFlags().StringVar(&pathFlag, "path", "", "Path to project root (default: current directory)")

	root.AddCommand(
		kirogo.NewInitCmd(scaffold.Templates),
		kirogo.NewAddCmd(scaffold.Templates),
		kirogo.NewPresetCmd(),
		kirogo.NewListCmd(),
		newVersionCmd(),
	)

	specsCmd := &cobra.Command{
		Use:   "specs",
		Short: "Manage specs (powered by kiro-specs)",
	}
	specsCmd.AddCommand(
		specs.NewListCmd(&pathFlag),
		specs.NewShowCmd(&pathFlag),
		specs.NewNextCmd(&pathFlag),
		specs.NewArchiveCmd(&pathFlag),
		specs.NewStatsCmd(&pathFlag),
	)
	root.AddCommand(specsCmd)

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show version info",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("kiro-go v0.1.0")
			fmt.Println("Go 1.23+")
		},
	}
}
