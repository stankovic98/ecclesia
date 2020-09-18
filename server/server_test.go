package server

import (
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stankovic98/ecclesia/model"
)

func TestPing(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "/api/ping", nil)
	if err != nil {
		t.Logf("server_test failed: %v\n", err)
		return
	}
	response := httptest.NewRecorder()
	server := Server{}
	server.GetRoutes().ServeHTTP(response, request)
	if bytes.Compare(bytes.TrimSpace(response.Body.Bytes()), []byte("pong")) != 0 {
		t.Fail()
	}
}

func TestMainDispatcher(t *testing.T) {
	existingUrls := [...]string{"/varazdinska-biskupija", "/varazdinska-biskupija/zupa-strigova", "/varazdinska-biskupija/zupa-nedelisce", "/zagrebacka-biskupija", "/zagrebacka-biskupija/marija-pomocnica"}
	nonExistingUrls := [...]string{"/", "/nepostojeca-zupa", "/kriva-zupa", "/zupa-strigova"}
	for i := 0; i < len(existingUrls); i++ {
		request, err := http.NewRequest(http.MethodGet, "/api"+existingUrls[i], nil)
		if err != nil {
			t.Logf("server_test failed: %v\n", err)
			return
		}
		response := httptest.NewRecorder()
		server := Server{testDB{}}
		server.GetRoutes().ServeHTTP(response, request)
		if response.Code != http.StatusOK {
			t.Logf("fail at url: %s, with message %v\n", existingUrls[i], response.Body)
			t.Fail()
		}
	}
	for i := 0; i < len(nonExistingUrls); i++ {
		request, err := http.NewRequest(http.MethodGet, "/api"+nonExistingUrls[i], nil)
		if err != nil {
			t.Logf("server_test failed: %v\n", err)
			return
		}
		response := httptest.NewRecorder()
		server := Server{testDB{}}
		server.GetRoutes().ServeHTTP(response, request)
		if response.Code != http.StatusNotFound {
			t.Logf("fail at url: %s, with message %v and code %v\n", existingUrls[i], response.Body, response.Code)
			t.Fail()
		}
	}
}

type testDB struct{}

func (db testDB) GetDioceseInfo(id string) (model.Diocese, error) {
	existingUrls := [...]string{"varazdinska-biskupija", "zagrebacka-biskupija"}
	for i := 0; i < len(existingUrls); i++ {
		if existingUrls[i] == id {
			return model.Diocese{
				UID:      id,
				Name:     "Test Success",
				Info:     "Success",
				Aritcles: nil,
			}, nil
		}
	}
	return model.Diocese{}, sql.ErrNoRows
}

func (db testDB) GetParish(dioceseID, parishID string) (model.Parish, error) {
	type mockData struct {
		DioceseID string
		ParishID  string
	}
	testData := []mockData{
		{"varazdinska-biskupija", "zupa-strigova"},
		{"varazdinska-biskupija", "zupa-nedelisce"},
		{"zagrebacka-biskupija", "marija-pomocnica"},
	}
	for i := 0; i < len(testData); i++ {
		if testData[i].DioceseID == dioceseID && testData[i].ParishID == parishID {
			return model.Parish{
				UID:       parishID,
				DioceseID: dioceseID,
				Name:      "test radi",
			}, nil
		}
	}
	return model.Parish{}, sql.ErrNoRows
}
func (db testDB) GetAllDioceses() []model.Diocese                { return nil }
func (db testDB) GetAllParishes(dioceseID string) []model.Parish { return nil }
func (db testDB) ValidUser(email, password string) bool          { return true }
func (db testDB) UpdateInfo(info, email string) error            { return nil }
func (db testDB) PublishArticle(article model.Aritcle) error     { return nil }
func (db testDB) GetInfo(email string) (string, error)           { return "", nil }
