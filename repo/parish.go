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
		err := rows.Scan(&p.UID, &p.Name, &p.Priest, &p.Info, &p.DioceseID)
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
	err := r.db.QueryRow(sqlStatement, id).Scan(&diocese.UID, &diocese.Name, &diocese.Info)
	if err == sql.ErrNoRows {
		log.Printf("diocese with id %s doesn't exist\n", id)
		return diocese, err
	} else if err != nil {
		log.Printf("Can't scan the database: %v\n", err)
		return diocese, err
	}
	diocese.Aritcles = r.getArticles(id)
	return diocese, nil
}

func (r *Repo) GetParish(dioceseID, parishID string) (model.Parish, error) {
	sqlStatement := "SELECT * FROM parishes WHERE diocese_id=$1 AND UID=$2;"
	var parish model.Parish
	err := r.db.QueryRow(sqlStatement, dioceseID, parishID).Scan(&parish.UID, &parish.Name, &parish.Priest, &parish.Info, &parish.DioceseID)
	if err == sql.ErrNoRows {
		log.Printf("parish with id %s doesn't exist\n", parishID)
		return parish, err
	} else if err != nil {
		log.Printf("Can't scan the database: %v\n", err)
		return parish, err
	}
	parish.Aritcles = r.getArticles(parishID)
	return parish, nil
}

func (r *Repo) getArticles(id string) []model.Aritcle {
	sqlStatment := `
		SELECT a.title, a.content, a.create_at, a.author FROM articles AS a INNER JOIN published_articles as pub
			ON a.uid = pub.article_uid WHERE pub.published_under=$1; 
	`
	var articles []model.Aritcle
	rows, err := r.db.Query(sqlStatment, id)
	if err != nil {
		log.Printf("can't query the articles: %v\n", err)
		return nil
	}
	for rows.Next() {
		var a model.Aritcle
		err = rows.Scan(&a.Title, &a.Content, &a.CreatedAt, &a.Author)
		if err != nil {
			log.Printf("can't read row by row: %v\n", err)
			return nil
		}
		articles = append(articles, a)
	}
	return articles
}
