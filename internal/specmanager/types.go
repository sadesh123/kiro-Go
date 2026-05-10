package specmanager

import "time"

type Status string

const (
	StatusIdea       Status = "idea"
	StatusDraft      Status = "draft"
	StatusPlanned    Status = "planned"
	StatusInProgress Status = "in-progress"
	StatusComplete   Status = "complete"
)

type Spec struct {
	Name       string
	Path       string
	Status     Status
	TasksDone  int
	TasksTotal int
	ModTime    time.Time
}
