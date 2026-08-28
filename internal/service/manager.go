package service

import (
	"fmt"
	"sync"
	"waveboard/internal/domain"
	"waveboard/internal/store"
)

type Manager struct {
	Store *store.Store
	mu    sync.Mutex
}

func NewManager(s *store.Store) *Manager { return &Manager{Store: s} }
func (m *Manager) CreateWave(w domain.Wave) error {
	if e := domain.ValidateWave(w); e != nil {
		return e
	}
	return m.Store.SaveWave(w)
}
func (m *Manager) AddTask(t domain.PickTask) error {
	if e := domain.ValidateTask(t); e != nil {
		return e
	}
	return m.Store.SaveTask(t)
}
func (m *Manager) UpdateTask(id, status, picker string) error {
	if !domain.ValidStatus(status) {
		return fmt.Errorf("invalid status")
	}
	t, e := m.Store.Task(id)
	if e != nil {
		return e
	}
	t.Status = status
	t.Picker = picker
	return m.Store.SaveTask(t)
}
func (m *Manager) Snapshot(zone string) ([]domain.PickTask, []domain.Progress, error) {
	tasks, e := m.Store.AllTasks()
	if e != nil {
		return nil, nil, e
	}
	waves := map[string]domain.Wave{}
	for _, t := range tasks {
		if _, ok := waves[t.WaveID]; !ok {
			w, er := m.Store.Wave(t.WaveID)
			if er != nil {
				return nil, nil, er
			}
			waves[t.WaveID] = w
		}
	}
	filtered := store.FilterTasks(tasks, zone, waves)
	ps := []domain.Progress{}
	seen := map[string]bool{}
	for _, t := range filtered {
		if !seen[t.WaveID] {
			ps = append(ps, store.BuildProgress(tasks, t.WaveID))
			seen[t.WaveID] = true
		}
	}
	return filtered, ps, nil
}
