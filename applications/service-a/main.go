package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

var delayer <-chan time.Time
var seconds time.Duration

func main() {
	srv := http.NewServeMux()

	seconds = 5
	delayer = time.Tick(seconds * time.Second)

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

	srv.HandleFunc("/random-delay", func(w http.ResponseWriter, r *http.Request) {
		num := rand.Intn(10)
		if num < 5 {
			<-delayer
		}
		w.Write([]byte("Success"))
	})

	srv.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Fatal(http.ListenAndServe(":8080", srv))
}
