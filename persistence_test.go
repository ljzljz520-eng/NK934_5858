package waveboard

import (
	"path/filepath"
	"testing"
	"waveboard/internal/domain"
	"waveboard/internal/store"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "state.db")
	s, e := store.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	if e = s.SaveWave(domain.Wave{ID: "P-1", Zone: "A"}); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = store.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if _, e = s.Wave("P-1"); e != nil {
		t.Fatal(e)
	}
}
