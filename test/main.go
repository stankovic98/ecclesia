package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	apiHost = "ecclesia:5000"
	apiURL  = "http://" + apiHost
)

func init() {
	for i := 0; i < 20; i++ {
		conn, err := net.DialTimeout("tcp", apiHost, time.Millisecond*100)
		if err == nil {
			log.Printf("API available at attempt %d", i+1)
			conn.Close()
			return
		}
		time.Sleep(time.Millisecond * 500)
	}
	log.Fatalf("Unable to connect to the API.")
}

type test struct {
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
	Want     string `json:"want"`
	Method   string `json:"method"`
	Body     string `json:"body"`
}

func main() {
	testFile, err := ioutil.ReadFile("tests.json")
	if err != nil {
		log.Fatalf("can't read file: %v\n", err)
	}
	var tests []test
	err = json.Unmarshal(testFile, &tests)
	if err != nil {
		log.Printf("can't unmarshal: %v\n", err)
		return
	}
	start := time.Now()
	for i := 0; i < len(tests); i++ {
		data := testingTemplate(&tests[i])
		if data != tests[i].Want {
			log.Printf("Test %s failed: want %s, got %s\n", tests[i].Name, tests[i].Want, data)
		}
	}
	log.Printf("tests complited in %s\n", time.Now().Sub(start))
}

func testingTemplate(t *test) string {
	resp := getResponse(t)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("can't read response from %s: %v\n", t.Endpoint, err)
	}
	data = bytes.TrimSpace(data)
	return string(data)
}

func getResponse(t *test) *http.Response {
	var resp *http.Response
	var err error
	if t.Method == http.MethodGet {
		resp, err = http.Get(apiURL + t.Endpoint)
	} else {
		resp, err = http.Post(apiURL+t.Endpoint, "application/json", strings.NewReader(t.Body))
	}
	if err != nil {
		log.Fatalf("can't ping server: %v\n", err)
	}
	return resp
}
