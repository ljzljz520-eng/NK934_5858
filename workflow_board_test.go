package waveboard

import (
	"testing"
	"waveboard/internal/fixture"
	"waveboard/internal/service"
	"waveboard/internal/store"
)

func TestWorkflowBoardRefresh(t *testing.T) {
	s, e := store.Open(t.TempDir() + "/b.db")
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	m := service.NewManager(s)
	if e = fixture.Seed(m); e != nil {
		t.Fatal(e)
	}
	tasks, p, e := m.Snapshot("")
	if e != nil || len(tasks) != 4 || len(p) != 2 {
		t.Fatalf("snapshot %v %d %d", e, len(tasks), len(p))
	}
}
