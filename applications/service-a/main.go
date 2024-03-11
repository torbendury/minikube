package main

import (
	"log"
	"math/rand"
	"net/http"
)

func main() {
	srv := http.NewServeMux()

	srv.HandleFunc("/service-a", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from service A"))
	})

	srv.HandleFunc("/random-fail", func(w http.ResponseWriter, r *http.Request) {
		num := rand.Intn(10)
		if num < 5 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed"))
			return
		}
		w.Write([]byte("Success"))
	})

	srv.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Fatal(http.ListenAndServe(":8080", srv))
}
