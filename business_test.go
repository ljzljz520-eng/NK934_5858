package waveboard

import (
	"testing"
	"waveboard/internal/fixture"
	"waveboard/internal/service"
	"waveboard/internal/store"
	"waveboard/internal/workflow"
)

func TestBusinessChain29(t *testing.T) {
	s, e := store.Open(t.TempDir() + "/c.db")
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	m := service.NewManager(s)
	if e = fixture.Seed(m); e != nil {
		t.Fatal(e)
	}
	if e = workflow.NewProcessor(m).Process([]string{"T-1", "T-2"}, "picked"); e != nil {
		t.Fatalf("tasks should finish normally: %v", e)
	}
}
