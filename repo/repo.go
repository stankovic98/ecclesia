package repo

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/stankovic98/ecclesia/model"
)

type Repo struct {
	db *sql.DB
}

type DatabaseActioner interface {
	GetParish(dioceseID, parishID string) (model.Parish, error)
	GetDioceseInfo(id string) (model.Diocese, error)
	GetAllParishes(dioceseID string) []model.Parish
	GetAllDioceses() []model.Diocese
	ValidUser(email, password string) bool
	UpdateInfo(info, email string) error
	PublishArticle(article model.Aritcle) error
	GetInfo(email string) (string, error)
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
