package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type newInfo struct {
	Info string `json:"info"`
}

func (s Server) editInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("wrong method %s, use POST\n", r.Method)
		return
	}
	adminMail := r.Context().Value("email").(string)
	log.Printf("admin logged in as: %s\n", adminMail)
	var info newInfo
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		log.Printf("wrong format: %v\n", err)
		w.Write([]byte("can't decode request body"))
		return
	}
	err = s.Repo.UpdateInfo(info.Info, adminMail)
	if err != nil {
		log.Printf("server: can't update info: %v\n", err)
		w.Write([]byte("can't update info"))
		return
	}
	w.Write([]byte("success"))
}
