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
	Repo *repo.Repo
}

func (s *Server) GetRoutes() *http.ServeMux {
	routes := http.NewServeMux()
	routes.HandleFunc("/ping", s.ping)
	routes.HandleFunc("/all-parishes", s.getAllParishes)
	routes.HandleFunc("/", s.mainDispatcher)
	return routes
}

func (s *Server) ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func (s *Server) mainDispatcher(w http.ResponseWriter, r *http.Request) {
	urlPaths := strings.Split(r.URL.Path, "/")
	dioceseID := urlPaths[1]
	if len(urlPaths) < 3 {
		diocese, err := s.Repo.GetDioceseInfo(dioceseID)
		if err == sql.ErrNoRows {
			w.Write([]byte("diocese with id " + dioceseID + "doesn't exist\n"))
			return
		}
		json.NewEncoder(w).Encode(diocese)
		return
	}
	parishID := urlPaths[2]
	parish, err := s.Repo.GetParish(dioceseID, parishID)
	if err == sql.ErrNoRows {
		w.Write([]byte("parish with id " + parishID + " doesn't exist\n"))
		return
	}
	json.NewEncoder(w).Encode(parish)
	return
}
