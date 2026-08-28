package waveboard

import (
	"testing"
	"waveboard/internal/fixture"
	"waveboard/internal/service"
	"waveboard/internal/store"
)

func TestWorkflowZoneFilter(t *testing.T) {
	s, e := store.Open(t.TempDir() + "/z.db")
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	m := service.NewManager(s)
	fixture.Seed(m)
	tasks, _, e := m.Snapshot("A")
	if e != nil || len(tasks) != 2 {
		t.Fatalf("zone %v %d", e, len(tasks))
	}
}
