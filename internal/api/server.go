package api

import (
	"encoding/json"
	"net/http"
	"waveboard/internal/service"
)

type Server struct{ Manager *service.Manager }

func New(m *service.Manager) *Server { return &Server{Manager: m} }
func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/board", s.board)
	mux.HandleFunc("/api/tasks", s.tasks)
	return mux
}
func (s *Server) board(w http.ResponseWriter, r *http.Request) {
	zone := r.URL.Query().Get("zone")
	tasks, progress, e := s.Manager.Snapshot(zone)
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(map[string]any{"tasks": tasks, "progress": progress, "zone": zone})
}
func (s *Server) tasks(w http.ResponseWriter, r *http.Request) {
	tasks, _, e := s.Manager.Snapshot("")
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}
