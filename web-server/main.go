package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello, SRE's\n")
	})
	http.HandleFunc("/error", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "🔥\n")
	})
	log.Println("Started on port 8080")
	log.Println("To close connection CTRL+C :-)")
	http.ListenAndServe(":8080", nil)
}
