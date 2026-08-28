package store

import (
	"sort"
	"waveboard/internal/domain"
)

func FilterTasks(tasks []domain.PickTask, zone string, waves map[string]domain.Wave) []domain.PickTask {
	out := make([]domain.PickTask, 0)
	for _, t := range tasks {
		w, ok := waves[t.WaveID]
		if !ok {
			continue
		}
		if zone != "" && w.Zone != zone {
			continue
		}
		out = append(out, t)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Location < out[j].Location })
	return out
}
func BuildProgress(tasks []domain.PickTask, waveID string) domain.Progress {
	p := domain.Progress{WaveID: waveID}
	for _, t := range tasks {
		if t.WaveID != waveID {
			continue
		}
		switch t.Status {
		case domain.StatusPending:
			p.Pending++
		case domain.StatusPicked:
			p.Picked++
		case domain.StatusReview:
			p.Review++
		case domain.StatusShortage:
			p.Shortage++
		case domain.StatusPacked:
			p.Packed++
		}
	}
	return p
}
