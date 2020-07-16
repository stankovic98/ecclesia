package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type creadentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type tokenResponse struct {
	jwt string `json:"jwt"`
}

func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	var cred creadentials
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		log.Printf("/login: can't decode body %v\n", err)
		w.Write([]byte("unvalid request body"))
		return
	}
	if s.Repo.ValidUser(cred.Email, cred.Password) == false {
		log.Printf("/login: user with username %s and password %s doesn't exists\n", cred.Email, cred.Password)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	claims := userClaims{
		Email: cred.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "ecclesia",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, err := token.SignedString([]byte("superTajnsLozinka")) // prebaci u env file
	json.NewEncoder(w).Encode(tokenResponse{jwt})
}
