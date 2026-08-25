package workflow

import "waveboard/internal/domain"

type Checklist struct{ Tasks []domain.PickTask }

func NewChecklist(tasks []domain.PickTask) Checklist {
	return Checklist{Tasks: domain.NormalizeTasks(tasks)}
}
func (c Checklist) PendingIDs() []string {
	out := []string{}
	for _, t := range c.Tasks {
		if t.Status == domain.StatusPending {
			out = append(out, t.ID)
		}
	}
	return out
}
func (c Checklist) Complete() bool {
	if len(c.Tasks) == 0 {
		return false
	}
	for _, t := range c.Tasks {
		if !domain.IsTerminal(t.Status) {
			return false
		}
	}
	return true
}
