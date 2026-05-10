package specmanager

func DeriveStatus(s Spec, empty, hasRequirements, hasTasks bool) Status {
	switch {
	case empty:
		return StatusIdea
	case !hasTasks && hasRequirements:
		return StatusDraft
	case hasTasks && s.TasksDone == 0:
		return StatusPlanned
	case hasTasks && s.TasksDone > 0 && s.TasksDone < s.TasksTotal:
		return StatusInProgress
	case hasTasks && s.TasksTotal > 0 && s.TasksDone == s.TasksTotal:
		return StatusComplete
	default:
		return StatusIdea
	}
}
