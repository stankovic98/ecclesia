package repo

import "log"

func (r *Repo) ValidUser(email, password string) bool {
	err := r.db.QueryRow("SELECT * FROM admins WHERE email = $1", email).Scan()
	if err != nil {
		log.Printf("/login: email not valid: %v\n", err)
		return false
	}
	return true
}
