package domain

import "sort"

type StatusCount struct {
	Status string
	Count  int
}

func StatusOrder() []string {
	return []string{StatusPending, StatusPicked, StatusReview, StatusShortage, StatusPacked}
}
func IsTerminal(s string) bool { return s == StatusShortage || s == StatusPacked }
func CanTransition(from, to string) bool {
	if !ValidStatus(from) || !ValidStatus(to) {
		return false
	}
	if from == to {
		return true
	}
	switch from {
	case StatusPending:
		return to == StatusPicked || to == StatusShortage
	case StatusPicked:
		return to == StatusReview || to == StatusShortage
	case StatusReview:
		return to == StatusPacked || to == StatusShortage
	case StatusShortage:
		return false
	case StatusPacked:
		return false
	}
	return false
}
func NormalizeTasks(in []PickTask) []PickTask {
	out := append([]PickTask(nil), in...)
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Location == out[j].Location {
			return out[i].SKU < out[j].SKU
		}
		return out[i].Location < out[j].Location
	})
	return out
}
func CountStatuses(tasks []PickTask) []StatusCount {
	counts := map[string]int{}
	for _, t := range tasks {
		counts[t.Status]++
	}
	out := make([]StatusCount, 0, 5)
	for _, s := range StatusOrder() {
		out = append(out, StatusCount{s, counts[s]})
	}
	return out
}
