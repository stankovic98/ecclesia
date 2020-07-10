package repo

import (
	"database/sql"
	"fmt"
	"log"
)

type Repo struct {
	db *sql.DB
}

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "lozinka123"
)

func New() *Repo {
	r := Repo{}
	r.db = initDB()
	return &r
}

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
