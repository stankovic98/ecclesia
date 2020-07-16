package repo

import "log"

func (r *Repo) ValidUser(email, password string) bool {
	sql := `
		SELECT uid FROM admins
 		WHERE email = $1 
		   AND password = crypt($2, password);`
	var uid string
	err := r.db.QueryRow(sql, email, password).Scan(&uid)
	if err != nil {
		log.Printf("/login: user not valid: %v\n", err)
		return false
	}
	return true
}
