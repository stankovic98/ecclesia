package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello world")
	http.HandleFunc("/ping", ping)
	http.ListenAndServe(":5000", nil)
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}
