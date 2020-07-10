package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

type server struct {
	db *sql.DB
}

func main() {
	fmt.Println("Hello world")
	time.Sleep(5 * time.Second) // Wait for docker to initalize
	s := server{}
	s.db = initDB()
	http.HandleFunc("/ping", s.ping)
	http.HandleFunc("/all-parishes", s.getAllParishes)
	http.ListenAndServe(":5000", nil)
}

func (s *server) ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func (s *server) getAllParishes(w http.ResponseWriter, r *http.Request) {
	sqlStatement := "SELECT * FROM parishes;"
	rows, err := s.db.Query(sqlStatement)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var parishes []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			log.Println(err)
		}
		parishes = append(parishes, name)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(parishes)
}

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "lozinka123"
)

func initDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println("successfull connection")
	return db
}
