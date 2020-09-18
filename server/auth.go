package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Creadentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type TokenResponse struct {
	Jwt string `json:"jwt"`
}

func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	var cred Creadentials
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
			ExpiresAt: int64(time.Now().Add(7 * 24 * time.Hour).Unix()),
			Issuer:    "ecclesia",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, err := token.SignedString([]byte("superTajnsLozinka")) // prebaci u env file
	json.NewEncoder(w).Encode(TokenResponse{jwt})
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			log.Println("Malformed Token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
			return
		}
		jwtFromHeader := authHeader[1]
		token, err := jwt.ParseWithClaims(
			jwtFromHeader,
			&userClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte("superTajnsLozinka"), nil
			},
		)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("parseWithClaims: token can't be parsed"))
			log.Printf("token can't be paresd: %v\n", err)
			return
		}
		claims, ok := token.Claims.(*userClaims)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("token can't be parsed"))
			log.Println("token can't be paresd")
			return
		}
		if claims.ExpiresAt < time.Now().UTC().Unix() {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("token expired"))
			log.Println("token expired")
			return
		}
		email := claims.Email
		ctx := context.WithValue(r.Context(), "email", email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
