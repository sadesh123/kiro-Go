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

func NewShowCmd(pathFlag *string) *cobra.Command {
	return &cobra.Command{
		Use:   "show [name]",
		Short: "Show task list for a spec",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			root := resolveRoot(*pathFlag)
			specs, err := specmanager.ParseSpecDir(root)
			if err != nil {
				return fmt.Errorf("reading specs: %w", err)
			}

			query := args[0]
			spec, err := findSpec(specs, query)
			if err != nil {
				return err
			}

			printSpec(root, spec)
			return nil
		},
	}
}

func findSpec(specs []specmanager.Spec, query string) (specmanager.Spec, error) {
	for _, s := range specs {
		if s.Name == query {
			return s, nil
		}
	}

	var matches []specmanager.Spec
	for _, s := range specs {
		if strings.Contains(s.Name, query) {
			matches = append(matches, s)
		}
	}

	switch len(matches) {
	case 0:
		return specmanager.Spec{}, fmt.Errorf("no spec found matching %q", query)
	case 1:
		fmt.Fprintf(os.Stderr, "(matched %s)\n", matches[0].Name)
		return matches[0], nil
	default:
		fmt.Fprintln(os.Stderr, "Multiple specs match:")
		for _, m := range matches {
			fmt.Fprintf(os.Stderr, "  %s\n", m.Name)
		}
		os.Exit(1)
		return specmanager.Spec{}, nil
	}
}

func printSpec(root string, s specmanager.Spec) {
	bar := specmanager.ProgressBar(s.TasksDone, s.TasksTotal)
	header := fmt.Sprintf("%s  ·  %s  ·  %d/%d tasks done  ·  %s",
		s.Name, s.Status, s.TasksDone, s.TasksTotal, specmanager.RelativeTime(s.ModTime))
	fmt.Println()
	fmt.Printf("  %s\n", header)
	fmt.Printf("  %s %s\n", bar, strings.Repeat("─", max(0, 66-len(bar))))
	fmt.Println()

	tasksPath := filepath.Join(root, ".kiro", "specs", s.Name, "tasks.md")
	f, err := os.Open(tasksPath)
	if err != nil {
		// No tasks.md — fall back to requirements.md
		reqPath := filepath.Join(root, ".kiro", "specs", s.Name, "requirements.md")
		content, err := os.ReadFile(reqPath)
		if err != nil {
			fmt.Println("  (no tasks.md or requirements.md found)")
			return
		}
		fmt.Println("  (no tasks.md yet — showing requirements)")
		fmt.Println()
		fmt.Println(string(content))
		return
	}
	defer f.Close()

	var nextUp string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "- [x]") || strings.Contains(line, "- [X]") || strings.Contains(line, "- [ ]") {
			fmt.Printf("  %s\n", line)
			if nextUp == "" && strings.Contains(line, "- [ ]") {
				nextUp = strings.TrimPrefix(strings.TrimPrefix(line, "  "), "- [ ] ")
			}
		}
	}

	fmt.Println()
	if nextUp != "" {
		fmt.Printf("  Next up: %s\n", nextUp)
	} else {
		fmt.Println("  All tasks complete.")
	}
	fmt.Println()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
