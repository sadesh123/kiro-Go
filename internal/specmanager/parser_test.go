package specmanager

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseSpecDir(t *testing.T) {
	tests := []struct {
		name       string
		setup      func(root string)
		wantCount  int
		wantStatus Status
		specName   string
	}{
		{
			name: "empty folder is idea",
			setup: func(root string) {
				os.MkdirAll(filepath.Join(root, ".kiro", "specs", "my-feature"), 0755)
			},
			wantCount:  1,
			wantStatus: StatusIdea,
			specName:   "my-feature",
		},
		{
			name: "requirements only is draft",
			setup: func(root string) {
				dir := filepath.Join(root, ".kiro", "specs", "my-feature")
				os.MkdirAll(dir, 0755)
				os.WriteFile(filepath.Join(dir, "requirements.md"), []byte("# req"), 0644)
			},
			wantCount:  1,
			wantStatus: StatusDraft,
			specName:   "my-feature",
		},
		{
			name: "tasks all unchecked is planned",
			setup: func(root string) {
				dir := filepath.Join(root, ".kiro", "specs", "my-feature")
				os.MkdirAll(dir, 0755)
				os.WriteFile(filepath.Join(dir, "tasks.md"), []byte("- [ ] task 1\n- [ ] task 2\n"), 0644)
			},
			wantCount:  1,
			wantStatus: StatusPlanned,
			specName:   "my-feature",
		},
		{
			name: "tasks mixed is in-progress",
			setup: func(root string) {
				dir := filepath.Join(root, ".kiro", "specs", "my-feature")
				os.MkdirAll(dir, 0755)
				os.WriteFile(filepath.Join(dir, "tasks.md"), []byte("- [x] task 1\n- [ ] task 2\n"), 0644)
			},
			wantCount:  1,
			wantStatus: StatusInProgress,
			specName:   "my-feature",
		},
		{
			name: "tasks all checked is complete",
			setup: func(root string) {
				dir := filepath.Join(root, ".kiro", "specs", "my-feature")
				os.MkdirAll(dir, 0755)
				os.WriteFile(filepath.Join(dir, "tasks.md"), []byte("- [x] task 1\n- [x] task 2\n"), 0644)
			},
			wantCount:  1,
			wantStatus: StatusComplete,
			specName:   "my-feature",
		},
		{
			name: "underscore folders are skipped",
			setup: func(root string) {
				os.MkdirAll(filepath.Join(root, ".kiro", "specs", "_TEMPLATE"), 0755)
				os.MkdirAll(filepath.Join(root, ".kiro", "specs", "_archive"), 0755)
			},
			wantCount: 0,
		},
		{
			name: "no specs dir returns empty slice",
			setup: func(root string) {
				os.MkdirAll(filepath.Join(root, ".kiro"), 0755)
			},
			wantCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := t.TempDir()
			tt.setup(root)

			specs, err := ParseSpecDir(root)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(specs) != tt.wantCount {
				t.Fatalf("got %d specs, want %d", len(specs), tt.wantCount)
			}
			if tt.wantCount > 0 && specs[0].Status != tt.wantStatus {
				t.Errorf("got status %q, want %q", specs[0].Status, tt.wantStatus)
			}
		})
	}
}

func TestParseCheckboxes(t *testing.T) {
	tests := []struct {
		name      string
		content   string
		wantDone  int
		wantTotal int
	}{
		{"all checked", "- [x] a\n- [x] b\n", 2, 2},
		{"all unchecked", "- [ ] a\n- [ ] b\n", 0, 2},
		{"mixed", "- [x] a\n- [ ] b\n", 1, 2},
		{"uppercase X", "- [X] a\n- [ ] b\n", 1, 2},
		{"empty", "", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.CreateTemp(t.TempDir(), "tasks*.md")
			if err != nil {
				t.Fatal(err)
			}
			f.WriteString(tt.content)
			f.Close()

			done, total := parseCheckboxes(f.Name())
			if done != tt.wantDone {
				t.Errorf("done: got %d, want %d", done, tt.wantDone)
			}
			if total != tt.wantTotal {
				t.Errorf("total: got %d, want %d", total, tt.wantTotal)
			}
		})
	}
}
