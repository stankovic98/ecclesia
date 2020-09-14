package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/stankovic98/ecclesia/model"
)

type newInfo struct {
	Info string `json:"info"`
}

func (s Server) editInfo(w http.ResponseWriter, r *http.Request) {
	adminMail := r.Context().Value("email").(string)
	log.Printf("admin logged in as: %s\n", adminMail)
	if r.Method == http.MethodGet {
		info, err := s.Repo.GetInfo(adminMail)
		if err != nil {
			w.Write([]byte("can't get info"))
			return
		}
		err = json.NewEncoder(w).Encode(info)
		if err != nil {
			w.Write([]byte("can't encode info"))
		}
		return
	}
	if r.Method != http.MethodPost {
		log.Printf("wrong method %s, use POST\n", r.Method)
		return
	}
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

func (s Server) createArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("wrong method %s, use POST\n", r.Method)
		return
	}
	article := model.Aritcle{}
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		log.Printf("wrong format: %v\n", err)
		return
	}
	article.Author = r.Context().Value("email").(string)
	err = s.Repo.PublishArticle(article)
	if err != nil {
		log.Printf("can't store article: %v\n", err)
		http.Error(w, "can't store article", http.StatusBadRequest)
		return
	}
}
