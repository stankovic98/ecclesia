package server

import (
	"fmt"
	"net/http"

	"github.com/stankovic98/ecclesia/repo"
)

type Server struct {
	Repo *repo.Repo
}

func (s *Server) GetRoutes() *http.ServeMux {
	routes := http.NewServeMux()
	routes.HandleFunc("/ping", s.ping)
	routes.HandleFunc("/all-parishes", s.getAllParishes)
	return routes
}

func (s *Server) ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}
