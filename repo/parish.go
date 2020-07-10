package repo

import (
	"log"

	"github.com/stankovic98/ecclesia/model"
)

func (r *Repo) GetAllParishes() []model.Parish {
	sqlStatement := "SELECT * FROM parishes;"
	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var parishes []model.Parish
	for rows.Next() {
		var p model.Parish
		err := rows.Scan(&p.UID, &p.Name, &p.Priest, &p.DioceseID)
		if err != nil {
			log.Println(err)
		}
		parishes = append(parishes, p)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
	}
	return parishes
}
