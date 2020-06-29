package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type server struct {
	db *sql.DB
}

func main() {
	fmt.Println("Hello world")
	s := server{}
	s.db = initDB()
	http.HandleFunc("/ping", s.ping)
	http.ListenAndServe(":5000", nil)
}

func (s *server) ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "lozinka123"
	dbname   = "church"
)

func initDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("successfull connection")
	return db
}
