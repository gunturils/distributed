package api

import (
	"encoding/json"
	"net/http"

	"github.com/gunturils/distributed/internal/raft"
	"github.com/gunturils/distributed/internal/store"
)

type Server struct {
	Store *store.Store
	Node  *raft.Node
}

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func (s *Server) StatusHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	json.NewEncoder(w).Encode(map[string]any{
		"id":   s.Node.ID,
		"role": s.Node.Role,
		"term": s.Node.Term,
	})
}

func (s *Server) GetKeyHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	key := r.URL.Query().Get("key")
	val, ok := s.Store.Get(key)
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"key": key, "value": val})
}

func (s *Server) SetKeyHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	var body struct{ Key, Value string }
	json.NewDecoder(r.Body).Decode(&body)
	s.Store.Set(body.Key, body.Value)
	w.WriteHeader(http.StatusOK)
}
