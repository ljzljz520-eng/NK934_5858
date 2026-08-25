package waveboard

import (
	"testing"
	"waveboard/internal/fixture"
	"waveboard/internal/service"
	"waveboard/internal/store"
	"waveboard/internal/workflow"
)

func TestWorkflowTaskProcessing(t *testing.T) {
	s, e := store.Open(t.TempDir() + "/p.db")
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	m := service.NewManager(s)
	fixture.Seed(m)
	p := workflow.NewProcessor(m)
	if e = p.Process([]string{"T-1"}, "picked"); e == nil {
		t.Fatal("known injected close defect should surface")
	}
}
