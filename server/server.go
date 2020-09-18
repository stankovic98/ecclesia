package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/stankovic98/ecclesia/repo"
)

type Server struct {
	Repo repo.DatabaseActioner
}

func (s *Server) GetRoutes() *http.ServeMux {
	routes := http.NewServeMux()
	routes.Handle("/api/admin/edit-info", middleware(http.HandlerFunc(s.editInfo)))
	routes.Handle("/api/admin/new-article", middleware(http.HandlerFunc(s.createArticle)))
	routes.HandleFunc("/api/ping", s.ping)
	routes.HandleFunc("/api/all-parishes", s.getAllParishes)
	routes.HandleFunc("/api/all-diocese", s.getAllDioceses)
	routes.HandleFunc("/api/login", s.login)
	routes.HandleFunc("/api/", s.mainDispatcher)
	return routes
}

func (s *Server) ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func (s *Server) mainDispatcher(w http.ResponseWriter, r *http.Request) {
	urlPaths := strings.Split(r.URL.Path, "/")
	dioceseID := urlPaths[2]
	if len(urlPaths) < 4 {
		diocese, err := s.Repo.GetDioceseInfo(dioceseID)
		if err == sql.ErrNoRows {
			http.Error(w, "diocese with id "+dioceseID+"doesn't exist\n", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(diocese)
		return
	}
	parishID := urlPaths[3]
	parish, err := s.Repo.GetParish(dioceseID, parishID)
	if err == sql.ErrNoRows {
		http.Error(w, "parish with id "+parishID+" doesn't exist\n", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(parish)
	return
}
