package waveboard

import (
	"testing"
	"waveboard/internal/domain"
	"waveboard/internal/store"
)

func TestQueryProgress(t *testing.T) {
	w := map[string]domain.Wave{"W": {ID: "W", Zone: "A"}}
	ts := []domain.PickTask{{ID: "2", WaveID: "W", Location: "B", Status: domain.StatusPicked}, {ID: "1", WaveID: "W", Location: "A", Status: domain.StatusPending}}
	if len(store.FilterTasks(ts, "A", w)) != 2 {
		t.Fatal("filter")
	}
	if store.BuildProgress(ts, "W").Pending != 1 {
		t.Fatal("progress")
	}
}
