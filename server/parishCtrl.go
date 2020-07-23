package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) getAllParishes(w http.ResponseWriter, r *http.Request) {
	dioceseID := r.URL.Query().Get("dioceseID")
	parishes := s.Repo.GetAllParishes(dioceseID)
	json.NewEncoder(w).Encode(parishes)
}

func (s *Server) getAllDioceses(w http.ResponseWriter, r *http.Request) {
	dioceses := s.Repo.GetAllDioceses()
	json.NewEncoder(w).Encode(dioceses)
}
