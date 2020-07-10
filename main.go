package main

import (
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/stankovic98/ecclesia/repo"
	"github.com/stankovic98/ecclesia/server"
)

func main() {
	time.Sleep(5 * time.Second) // Wait for docker to initalize
	s := server.Server{}
	s.Repo = repo.New()
	http.ListenAndServe(":5000", s.GetRoutes())
}
