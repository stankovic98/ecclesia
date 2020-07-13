package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
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
		data := testingTemplate(tests[i].Endpoint)
		if data != tests[i].Want {
			log.Printf("Test %s failed: want %s, got %s\n", tests[i].Name, tests[i].Want, data)
		}
	}
	log.Printf("tests complited in %s\n", time.Now().Sub(start))
}

func testPing() {
	data := testingTemplate("/ping")
	if string(data) != "pong" {
		log.Printf("want %x, got %x\n", "pong", data)
		return
	}
	log.Println("Ping test succesfull")
}

func testGetAllParishes() {
	want := ``
	data := testingTemplate("/all-parishes")
	if data != want {
		log.Printf("want %x, got %x\n", want, data)
		return
	}
	log.Println("testGetAllParishes succesfull")
}

func testingTemplate(endpoint string) string {
	resp, err := http.Get(apiURL + endpoint)
	if err != nil {
		log.Fatalf("can't ping server: %v\n", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("can't read response from %s: %v\n", endpoint, err)
	}
	data = bytes.TrimSpace(data)
	return string(data)
}
