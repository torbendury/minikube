package main

import (
	"log"
	"net/http"
)

func main() {
	srv := http.NewServeMux()

	srv.HandleFunc("/service-a", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from service A"))
	})

	srv.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Fatal(http.ListenAndServe(":8080", srv))
}
