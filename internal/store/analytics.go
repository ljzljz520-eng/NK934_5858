package store

import "waveboard/internal/domain"

func GroupByWave(tasks []domain.PickTask) map[string][]domain.PickTask {
	out := map[string][]domain.PickTask{}
	for _, t := range tasks {
		out[t.WaveID] = append(out[t.WaveID], t)
	}
	return out
}
func CountByPicker(tasks []domain.PickTask) map[string]int {
	out := map[string]int{}
	for _, t := range tasks {
		p := t.Picker
		if p == "" {
			p = "unassigned"
		}
		out[p] += t.Quantity
	}
	return out
}
func Locations(tasks []domain.PickTask) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, t := range tasks {
		if !seen[t.Location] {
			seen[t.Location] = true
			out = append(out, t.Location)
		}
	}
	return out
}
