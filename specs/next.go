package specs

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sadesh123/kiro-go/internal/specmanager"
)

func NewNextCmd(pathFlag *string) *cobra.Command {
	return &cobra.Command{
		Use:   "next",
		Short: "What should I work on right now?",
		RunE: func(cmd *cobra.Command, args []string) error {
			root := resolveRoot(*pathFlag)
			specs, err := specmanager.ParseSpecDir(root)
			if err != nil {
				return fmt.Errorf("reading specs: %w", err)
			}

			spec, found := pickNext(specs)
			if !found {
				fmt.Println("  All specs are complete or ideas. Start a new one in Kiro.")
				return nil
			}

			remaining := spec.TasksTotal - spec.TasksDone
			fmt.Println()
			fmt.Println("  Pick up where you left off:")
			fmt.Println()
			fmt.Printf("  %s  ·  %s  ·  %d tasks remaining\n", spec.Name, spec.Status, remaining)
			fmt.Println()

			printUnchecked(root, spec)

			fmt.Printf("  Run: kiro-specs show %s\n", spec.Name)
			fmt.Println()
			return nil
		},
	}
}

func pickNext(specs []specmanager.Spec) (specmanager.Spec, bool) {
	var inProgress []specmanager.Spec
	var planned []specmanager.Spec

	for _, s := range specs {
		switch s.Status {
		case specmanager.StatusInProgress:
			inProgress = append(inProgress, s)
		case specmanager.StatusPlanned:
			planned = append(planned, s)
		}
	}

	if len(inProgress) > 0 {
		sort.Slice(inProgress, func(i, j int) bool {
			ri := inProgress[i].TasksTotal - inProgress[i].TasksDone
			rj := inProgress[j].TasksTotal - inProgress[j].TasksDone
			if ri != rj {
				return ri < rj
			}
			return inProgress[i].ModTime.After(inProgress[j].ModTime)
		})
		return inProgress[0], true
	}

	if len(planned) > 0 {
		sort.Slice(planned, func(i, j int) bool {
			return planned[i].ModTime.Before(planned[j].ModTime)
		})
		return planned[0], true
	}

	return specmanager.Spec{}, false
}

func printUnchecked(root string, s specmanager.Spec) {
	tasksPath := filepath.Join(root, ".kiro", "specs", s.Name, "tasks.md")
	f, err := os.Open(tasksPath)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "- [ ]") {
			fmt.Printf("    %s\n", line)
		}
	}
	fmt.Println()
}
