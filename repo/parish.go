package repo

import (
	"database/sql"
	"log"

	"github.com/stankovic98/ecclesia/model"
)

func (r *Repo) GetAllParishes() []model.Parish {
	sqlStatement := "SELECT * FROM parishes"
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

// GetDioceseInfo returns the data for diocese, if diocese is not reigsterd
// it returns the ErrNoRows
func (r *Repo) GetDioceseInfo(id string) (model.Diocese, error) {
	sqlStatement := "SELECT * FROM dioceses WHERE UID=$1;"
	var diocese model.Diocese
	err := r.db.QueryRow(sqlStatement, id).Scan(&diocese.UID, &diocese.Name)
	if err == sql.ErrNoRows {
		log.Printf("diocese with id %s doesn't exist\n", id)
		return diocese, err
	} else if err != nil {
		log.Printf("Can't scan the database: %v\n", err)
		return diocese, err
	}
	return diocese, nil
}

func (r *Repo) GetParish(dioceseID, parishID string) (model.Parish, error) {
	sqlStatement := "SELECT * FROM parishes WHERE diocese_id=$1 AND UID=$2;"
	var parish model.Parish
	err := r.db.QueryRow(sqlStatement, dioceseID, parishID).Scan(&parish.UID, &parish.Name, &parish.Priest, &parish.DioceseID)
	if err == sql.ErrNoRows {
		log.Printf("parish with id %s doesn't exist\n", parishID)
		return parish, err
	} else if err != nil {
		log.Printf("Can't scan the database: %v\n", err)
		return parish, err
	}
	return parish, nil
}
