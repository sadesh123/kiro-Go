package specmanager

import (
	"testing"
	"time"
)

func TestDeriveStatus(t *testing.T) {
	tests := []struct {
		name            string
		spec            Spec
		empty           bool
		hasRequirements bool
		hasTasks        bool
		want            Status
	}{
		{"empty folder → idea", Spec{}, true, false, false, StatusIdea},
		{"requirements only → draft", Spec{}, false, true, false, StatusDraft},
		{"tasks all unchecked → planned", Spec{TasksDone: 0, TasksTotal: 3}, false, false, true, StatusPlanned},
		{"tasks mixed → in-progress", Spec{TasksDone: 2, TasksTotal: 5}, false, false, true, StatusInProgress},
		{"tasks all done → complete", Spec{TasksDone: 4, TasksTotal: 4}, false, false, true, StatusComplete},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeriveStatus(tt.spec, tt.empty, tt.hasRequirements, tt.hasTasks)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestProgressBar(t *testing.T) {
	tests := []struct {
		done, total int
		want        string
	}{
		{0, 0, "──────"},
		{0, 6, "░░░░░░"},
		{6, 6, "██████"},
		{3, 6, "███░░░"},
		{1, 6, "█░░░░░"},
		{4, 6, "████░░"},
	}

	for _, tt := range tests {
		got := progressBar(tt.done, tt.total)
		if got != tt.want {
			t.Errorf("progressBar(%d,%d) = %q, want %q", tt.done, tt.total, got, tt.want)
		}
	}
}

func TestRelativeTime(t *testing.T) {
	now := time.Now()
	tests := []struct {
		t    time.Time
		want string
	}{
		{now, "today"},
		{now.Add(-25 * time.Hour), "yesterday"},
		{now.Add(-3 * 24 * time.Hour), "3 days ago"},
		{now.Add(-6 * 24 * time.Hour), "6 days ago"},
		{now.Add(-7 * 24 * time.Hour), "1 weeks ago"},
		{now.Add(-14 * 24 * time.Hour), "2 weeks ago"},
		{now.Add(-30 * 24 * time.Hour), "1 months ago"},
		{now.Add(-60 * 24 * time.Hour), "2 months ago"},
	}

	for _, tt := range tests {
		got := relativeTime(tt.t)
		if got != tt.want {
			t.Errorf("relativeTime(%v) = %q, want %q", tt.t, got, tt.want)
		}
	}
}
