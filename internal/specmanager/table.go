package specmanager

import (
	"fmt"
	"io"
	"math"
	"strings"
	"time"
)

const (
	colorReset   = "\033[0m"
	colorGreen   = "\033[32m"
	colorYellow  = "\033[33m"
	colorBlue    = "\033[34m"
	colorGrey    = "\033[37m"
	colorDarkGrey = "\033[90m"
)

func PrintTable(w io.Writer, specs []Spec, noColor bool) {
	fmt.Fprintf(w, "  %-30s %-14s %-14s %s\n", "SPEC", "STATUS", "PROGRESS", "AGE")
	fmt.Fprintf(w, "  %s\n", strings.Repeat("─", 66))

	for _, s := range specs {
		bar := progressBar(s.TasksDone, s.TasksTotal)
		progress := bar
		if s.TasksTotal > 0 {
			progress = fmt.Sprintf("%s %d/%d", bar, s.TasksDone, s.TasksTotal)
		}

		statusStr := string(s.Status)
		if !noColor {
			statusStr = colorStatus(s.Status) + statusStr + colorReset
		}

		fmt.Fprintf(w, "  %-30s %-14s %-20s %s\n",
			s.Name,
			statusStr,
			progress,
			relativeTime(s.ModTime),
		)
	}

	fmt.Fprintln(w)
	printSummary(w, specs)
}

func colorStatus(s Status) string {
	switch s {
	case StatusComplete:
		return colorGreen
	case StatusInProgress:
		return colorYellow
	case StatusPlanned:
		return colorBlue
	case StatusDraft:
		return colorGrey
	case StatusIdea:
		return colorDarkGrey
	default:
		return ""
	}
}

func progressBar(done, total int) string {
	if total == 0 {
		return "──────"
	}
	filled := int(math.Round(float64(done) / float64(total) * 6))
	return strings.Repeat("█", filled) + strings.Repeat("░", 6-filled)
}

func relativeTime(t time.Time) string {
	days := int(time.Since(t).Hours() / 24)
	switch {
	case days == 0:
		return "today"
	case days == 1:
		return "yesterday"
	case days < 7:
		return fmt.Sprintf("%d days ago", days)
	case days < 30:
		return fmt.Sprintf("%d weeks ago", days/7)
	default:
		return fmt.Sprintf("%d months ago", days/30)
	}
}

func printSummary(w io.Writer, specs []Spec) {
	counts := map[Status]int{}
	for _, s := range specs {
		counts[s.Status]++
	}
	fmt.Fprintf(w, "  %d specs  |  %d complete  |  %d in-progress  |  %d planned  |  %d draft  |  %d idea\n",
		len(specs),
		counts[StatusComplete],
		counts[StatusInProgress],
		counts[StatusPlanned],
		counts[StatusDraft],
		counts[StatusIdea],
	)
}

// RelativeTime is exported for use in commands
func RelativeTime(t time.Time) string {
	return relativeTime(t)
}

// ProgressBar is exported for use in commands
func ProgressBar(done, total int) string {
	return progressBar(done, total)
}
