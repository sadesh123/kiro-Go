package specs

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sadesh123/kiro-go/internal/specmanager"
)

func NewStatsCmd(pathFlag *string) *cobra.Command {
	return &cobra.Command{
		Use:   "stats",
		Short: "Summary stats across all specs",
		RunE: func(cmd *cobra.Command, args []string) error {
			root := resolveRoot(*pathFlag)
			specs, err := specmanager.ParseSpecDir(root)
			if err != nil {
				return fmt.Errorf("reading specs: %w", err)
			}

			printStats(specs)
			return nil
		},
	}
}

func printStats(specs []specmanager.Spec) {
	counts := map[specmanager.Status]int{}
	var totalDone, totalTasks int
	var oldestActive *specmanager.Spec
	var mostRecent *specmanager.Spec

	for i := range specs {
		s := &specs[i]
		counts[s.Status]++
		totalDone += s.TasksDone
		totalTasks += s.TasksTotal

		if s.Status != specmanager.StatusComplete && s.Status != specmanager.StatusIdea {
			if oldestActive == nil || s.ModTime.Before(oldestActive.ModTime) {
				oldestActive = s
			}
		}
		if mostRecent == nil || s.ModTime.After(mostRecent.ModTime) {
			mostRecent = s
		}
	}

	total := len(specs)
	complete := counts[specmanager.StatusComplete]

	pct := 0
	if total > 0 {
		pct = complete * 100 / total
	}
	taskPct := 0
	if totalTasks > 0 {
		taskPct = totalDone * 100 / totalTasks
	}

	fmt.Fprintln(os.Stdout, "  Spec summary")
	fmt.Fprintln(os.Stdout, "  ─────────────────────────────")
	fmt.Fprintf(os.Stdout, "  %-20s %d\n", "Total specs", total)
	fmt.Fprintf(os.Stdout, "  %-20s %d  (%d%%)\n", "Complete", complete, pct)
	fmt.Fprintf(os.Stdout, "  %-20s %d\n", "In progress", counts[specmanager.StatusInProgress])
	fmt.Fprintf(os.Stdout, "  %-20s %d\n", "Planned", counts[specmanager.StatusPlanned])
	fmt.Fprintf(os.Stdout, "  %-20s %d\n", "Draft", counts[specmanager.StatusDraft])
	fmt.Fprintf(os.Stdout, "  %-20s %d\n", "Ideas", counts[specmanager.StatusIdea])
	fmt.Fprintln(os.Stdout)
	fmt.Fprintf(os.Stdout, "  %-20s %d / %d  (%d%%)\n", "Tasks completed", totalDone, totalTasks, taskPct)

	if oldestActive != nil {
		fmt.Fprintf(os.Stdout, "  %-20s %s  (%s)\n", "Oldest active", oldestActive.Name, specmanager.RelativeTime(oldestActive.ModTime))
	}
	if mostRecent != nil {
		fmt.Fprintf(os.Stdout, "  %-20s %s  (%s)\n", "Most recent", mostRecent.Name, specmanager.RelativeTime(mostRecent.ModTime))
	}
	fmt.Fprintln(os.Stdout)
}
