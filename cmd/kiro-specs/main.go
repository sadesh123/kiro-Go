package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sadesh123/kiro-go/internal/splash"
	"github.com/sadesh123/kiro-go/internal/specmanager"
	"github.com/sadesh123/kiro-go/specs"
	"golang.org/x/term"
)

func main() {
	var pathFlag string
	var jsonOut bool
	var statusFilter string
	var sortBy string

	root := &cobra.Command{
		Use:   "kiro-specs",
		Short: "Spec manager for Kiro",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Default command: list
			if jsonOut || statusFilter != "" || sortBy != "status" {
				return runList(pathFlag, statusFilter, sortBy, jsonOut)
			}
			splash.PrintSpecs()
			return runList(pathFlag, statusFilter, sortBy, jsonOut)
		},
	}

	root.PersistentFlags().StringVar(&pathFlag, "path", "", "Path to project root (default: current directory)")
	root.Flags().BoolVar(&jsonOut, "json", false, "Machine-readable JSON output")
	root.Flags().StringVar(&statusFilter, "status", "", "Filter by status")
	root.Flags().StringVar(&sortBy, "sort", "status", "Sort by: status, age, name")

	root.AddCommand(
		specs.NewListCmd(&pathFlag),
		specs.NewShowCmd(&pathFlag),
		specs.NewNextCmd(&pathFlag),
		specs.NewArchiveCmd(&pathFlag),
		specs.NewStatsCmd(&pathFlag),
	)

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

func runList(pathFlag, statusFilter, sortBy string, jsonOut bool) error {
	root := resolveRoot(pathFlag)
	sp, err := specmanager.ParseSpecDir(root)
	if err != nil {
		return fmt.Errorf("reading specs: %w", err)
	}

	if statusFilter != "" {
		var filtered []specmanager.Spec
		for _, s := range sp {
			if string(s.Status) == statusFilter {
				filtered = append(filtered, s)
			}
		}
		sp = filtered
	}

	if jsonOut {
		return specs.RunJSONOutput(sp)
	}

	noColor := os.Getenv("NO_COLOR") != "" || !term.IsTerminal(int(os.Stdout.Fd()))
	specmanager.PrintTable(os.Stdout, sp, noColor)
	return nil
}

func resolveRoot(pathFlag string) string {
	if pathFlag != "" {
		return pathFlag
	}
	cwd, err := os.Getwd()
	if err != nil {
		return "."
	}
	return cwd
}
