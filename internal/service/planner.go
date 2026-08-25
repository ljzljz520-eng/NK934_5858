package service

import (
	"fmt"
	"waveboard/internal/domain"
)

func (m *Manager) Assign(id, picker string) error {
	if picker == "" {
		return fmt.Errorf("picker required")
	}
	t, e := m.Store.Task(id)
	if e != nil {
		return e
	}
	if domain.IsTerminal(t.Status) {
		return fmt.Errorf("task terminal")
	}
	t.Picker = picker
	return m.Store.SaveTask(t)
}
func (m *Manager) Transition(id, next string) error {
	t, e := m.Store.Task(id)
	if e != nil {
		return e
	}
	if !domain.CanTransition(t.Status, next) {
		return fmt.Errorf("transition %s to %s denied", t.Status, next)
	}
	t.Status = next
	return m.Store.SaveTask(t)
}
func (m *Manager) WaveTasks(waveID string) ([]domain.PickTask, error) {
	all, e := m.Store.AllTasks()
	if e != nil {
		return nil, e
	}
	out := []domain.PickTask{}
	for _, t := range all {
		if t.WaveID == waveID {
			out = append(out, t)
		}
	}
	return domain.NormalizeTasks(out), nil
}
