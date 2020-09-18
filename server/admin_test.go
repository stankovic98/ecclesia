package server

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stankovic98/ecclesia/model"
)

type TestData struct {
	Method           string
	Body             model.Aritcle
	wantedStatusCode int
}

func TestCreateArticle(t *testing.T) {
	testTable := []TestData{
		{http.MethodGet, model.Aritcle{Title: "helo", Content: "helo world and beyond"}, http.StatusMethodNotAllowed},
		{http.MethodPost, model.Aritcle{Title: "helo", Content: "helow world and beyond"}, http.StatusOK},
	}
	for i := 0; i < len(testTable); i++ {
		jwt := loginForTesting("kuhar@gmail.com", "lozinka123")
		structMarshaled, _ := json.Marshal(testTable[i].Body)
		request, err := http.NewRequest(testTable[i].Method, "/api/admin/new-article", bytes.NewReader(structMarshaled))
		if err != nil {
			t.Logf("server_test failed: %v\n", err)
			return
		}
		request.Header.Set("Authorization", "Bearer "+jwt)
		response := httptest.NewRecorder()
		server := Server{testDB{}}
		server.GetRoutes().ServeHTTP(response, request)
		if response.Code != testTable[i].wantedStatusCode {
			t.Errorf("want: %v, got: %v\n", testTable[i].wantedStatusCode, response.Code)
		}
	}
}

func loginForTesting(email, password string) string {
	structMarshaled, _ := json.Marshal(Creadentials{email, password})
	request, err := http.NewRequest(http.MethodPost, "/api/login", bytes.NewReader(structMarshaled))
	if err != nil {
		log.Printf("can't create login request: %v\n", err)
		return ""
	}
	response := httptest.NewRecorder()
	server := Server{testDB{}}
	server.GetRoutes().ServeHTTP(response, request)
	var jwt TokenResponse
	err = json.NewDecoder(response.Body).Decode(&jwt)
	if err != nil {
		log.Printf("can't decode response: %v\n", err)
		return ""
	}
	return jwt.Jwt
}
