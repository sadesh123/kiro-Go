package specs

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/spf13/cobra"
	"github.com/sadesh123/kiro-go/internal/specmanager"
	"golang.org/x/term"
)

type jsonSpec struct {
	Name       string `json:"name"`
	Status     string `json:"status"`
	TasksTotal int    `json:"tasks_total"`
	TasksDone  int    `json:"tasks_done"`
	AgeDays    int    `json:"age_days"`
	Path       string `json:"path"`
}

func NewListCmd(pathFlag *string) *cobra.Command {
	var statusFilter string
	var sortBy string
	var jsonOut bool

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all specs with status and progress",
		RunE: func(cmd *cobra.Command, args []string) error {
			root := resolveRoot(*pathFlag)
			specs, err := specmanager.ParseSpecDir(root)
			if err != nil {
				return fmt.Errorf("reading specs: %w", err)
			}

			if statusFilter != "" {
				specs = filterByStatus(specs, specmanager.Status(statusFilter))
			}

			sortSpecs(specs, sortBy)

			if jsonOut {
				return printJSON(specs)
			}

			noColor := os.Getenv("NO_COLOR") != "" || !term.IsTerminal(int(os.Stdout.Fd()))
			specmanager.PrintTable(os.Stdout, specs, noColor)
			return nil
		},
	}

	cmd.Flags().StringVar(&statusFilter, "status", "", "Filter by status: idea, draft, planned, in-progress, complete")
	cmd.Flags().StringVar(&sortBy, "sort", "status", "Sort by: status, age, name")
	cmd.Flags().BoolVar(&jsonOut, "json", false, "Machine-readable JSON output")

	return cmd
}

func filterByStatus(specs []specmanager.Spec, status specmanager.Status) []specmanager.Spec {
	var out []specmanager.Spec
	for _, s := range specs {
		if s.Status == status {
			out = append(out, s)
		}
	}
	return out
}

func sortSpecs(specs []specmanager.Spec, by string) {
	statusOrder := map[specmanager.Status]int{
		specmanager.StatusInProgress: 0,
		specmanager.StatusPlanned:    1,
		specmanager.StatusDraft:      2,
		specmanager.StatusIdea:       3,
		specmanager.StatusComplete:   4,
	}

	sort.Slice(specs, func(i, j int) bool {
		switch by {
		case "age":
			return specs[i].ModTime.After(specs[j].ModTime)
		case "name":
			return specs[i].Name < specs[j].Name
		default:
			oi, oj := statusOrder[specs[i].Status], statusOrder[specs[j].Status]
			if oi != oj {
				return oi < oj
			}
			return specs[i].ModTime.After(specs[j].ModTime)
		}
	})
}

// RunJSONOutput is exported for use by the root command default handler.
func RunJSONOutput(specs []specmanager.Spec) error {
	return printJSON(specs)
}

func printJSON(specs []specmanager.Spec) error {
	out := make([]jsonSpec, len(specs))
	for i, s := range specs {
		out[i] = jsonSpec{
			Name:       s.Name,
			Status:     string(s.Status),
			TasksTotal: s.TasksTotal,
			TasksDone:  s.TasksDone,
			AgeDays:    int(time.Since(s.ModTime).Hours() / 24),
			Path:       s.Path,
		}
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(out)
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
