package repo

import (
	"log"

	"github.com/stankovic98/ecclesia/model"
)

var sqlStatements = map[string]string{
	"getRegion":      "SELECT region FROM admins_of WHERE admin_email = $1",
	"checkDioceses":  "UPDATE dioceses SET info = $1 WHERE uid = $2",
	"checkParishes":  "UPDATE parishes SET info = $1 WHERE uid = $2",
	"storeArticle":   "INSERT INTO articles (title, content, author) VALUES ($1, $2, $3) RETURNING uid;",
	"publishArticle": "INSERT INTO published_articles (article_uid, published_under) VALUES ($1, $2)",
	"getInfo":        "SELECT info FROM parishes WHERE uid = $1",
}

func (r Repo) UpdateInfo(info, email string) error {
	region := r.getRegionByEmail(email)
	rs, _ := r.db.Exec(sqlStatements["checkDioceses"], info, region)
	rowsAffected, err := rs.RowsAffected()
	if rowsAffected == 0 {
		rs, err = r.db.Exec(sqlStatements["checkParishes"], info, region)
		parishRows, err := rs.RowsAffected()
		if err != nil || parishRows == 0 {
			log.Printf("repo: can't update info in parishes %v\n", err)
		}
		return err
	}
	if err != nil {
		log.Printf("repo: can't update info in dioceses %v\n", err)
		return err
	}
	return nil
}

func (r Repo) PublishArticle(article model.Aritcle) error {
	var articleUID int
	err := r.db.QueryRow(sqlStatements["storeArticle"], article.Title, article.Content, article.Author).Scan(&articleUID)
	if err != nil {
		return err
	}
	region := r.getRegionByEmail(article.Author)
	_, err = r.db.Exec(sqlStatements["publishArticle"], articleUID, region)
	return err
}

func (r Repo) getRegionByEmail(email string) string {
	var region string
	err := r.db.QueryRow(sqlStatements["getRegion"], email).Scan(&region)
	if err != nil {
		log.Printf("couldn't get region: %v\n", err)
		return ""
	}
	return region
}

// GetInfo returns info only for the parishes so it can be
// displayed on the admin page
func (r Repo) GetInfo(email string) (string, error) {
	region := r.getRegionByEmail(email)
	var info string
	err := r.db.QueryRow(sqlStatements["getInfo"], region).Scan(&info)
	if err != nil {
		log.Printf("can't get info for: %s (info is only avialabe for parishes)", region)
		return "", err
	}
	return info, nil
}
