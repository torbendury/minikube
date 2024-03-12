package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

var delayer <-chan time.Time
var seconds time.Duration

func oopsie() bool {
	return rand.Intn(10) < 5
}

func main() {
	srv := http.NewServeMux()

	seconds = 5
	delayer = time.Tick(seconds * time.Second)

	srv.HandleFunc("/service-a", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from service A\n"))
	})

	srv.HandleFunc("/random-fail", func(w http.ResponseWriter, r *http.Request) {
		if oopsie() {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed\n"))
			return
		}
		w.Write([]byte("Success\n"))
	})

	srv.HandleFunc("/random-delay", func(w http.ResponseWriter, r *http.Request) {
		if oopsie() {
			<-delayer
		}
		w.Write([]byte("Success\n"))
	})

	srv.HandleFunc("/retriable", func(w http.ResponseWriter, r *http.Request) {
		if oopsie() {
			w.Header().Set("x-try-again", "true")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed\n"))
			return
		}
		w.Write([]byte("Success\n"))
	})

	srv.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	srv.HandleFunc("/new-route", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("sorry we didnt tell you about the new route lol\n"))
	})

	log.Fatal(http.ListenAndServe(":8080", srv))
}
