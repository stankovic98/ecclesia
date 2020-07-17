package repo

import "log"

var sqlStatements = map[string]string{
	"getRegion":     "SELECT region FROM admins_of WHERE admin_email = $1",
	"checkDioceses": "UPDATE dioceses SET info = $1 WHERE uid = $2",
	"checkParishes": "UPDATE parishes SET info = $1 WHERE uid = $2",
}

func (r Repo) UpdateInfo(info, email string) error {
	var region string
	err := r.db.QueryRow(sqlStatements["getRegion"], email).Scan(&region)
	if err != nil {
		return err
	}
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
