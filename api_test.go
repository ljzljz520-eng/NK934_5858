package waveboard

import (
	"net/http/httptest"
	"testing"
	"waveboard/internal/api"
	"waveboard/internal/service"
	"waveboard/internal/store"
)

func TestAPIBoard(t *testing.T) {
	s, e := store.Open(t.TempDir() + "/a.db")
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	r := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/board", nil)
	api.New(service.NewManager(s)).Routes().ServeHTTP(r, req)
	if r.Code != 200 {
		t.Fatal(r.Code)
	}
}
