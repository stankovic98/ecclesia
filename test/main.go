package main

import (
	"bytes"
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

func main() {
	start := time.Now()
	testPing()
	testGetAllParishes()
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
	want := `[{"UID":"zupa-strigova","Name":"Župa Štrigova","Priest":"vlč. Kristijan Kuhar","DioceseID":"varazdinska-biskupija"},{"UID":"PL62ELIbTGUaaNTKIEZuFyns05crma","Name":"Župa Sveti Juraj na Bregu","Priest":"vlč Nikola Samodol","DioceseID":"varazdinska-biskupija"},{"UID":"PL62ELIbTGUaaNTKIEZuFyns05crmb","Name":"Župa Nedelišće","Priest":"Zvonimir Radoš","DioceseID":"varazdinska-biskupija"},{"UID":"PL62ELIbTGUaaNTKIEZuFyns05crmc","Name":"Župa Pribislavec","Priest":"Mladen Delić","DioceseID":"varazdinska-biskupija"},{"UID":"PL62ELIbTGUaaNTKIEZuFyns05crmd","Name":"Župa Blažene Djevice Marije Pomoćnice","Priest":"Tihomir Ladić","DioceseID":"zagrebacka-biskupija"}]`
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
