package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) getAllParishes(w http.ResponseWriter, r *http.Request) {
	parishes := s.Repo.GetAllParishes()
	json.NewEncoder(w).Encode(parishes)
}
