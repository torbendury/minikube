package main

import (
	"log"
	"net/http"
)

func main() {
	srv := http.NewServeMux()

	srv.HandleFunc("/service-b", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from service B"))
	})

	srv.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Fatal(http.ListenAndServe(":8080", srv))
}
