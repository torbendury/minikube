package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	betaFeature, err := strconv.ParseBool(os.Getenv("ENABLE_BETA_FEATURE"))
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	if betaFeature {
		fmt.Println("Beta features enabled")
		mux.HandleFunc("/service-b", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Beta feature requested")
			w.Write([]byte("Hello from service B beta\n"))
		})
	} else {
		mux.HandleFunc("/service-b", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Old feature requested")
			w.Write([]byte("Hello from service B\n"))
		})
	}

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
