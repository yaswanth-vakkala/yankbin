package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET	/{$}", home)
	mux.HandleFunc("GET	/yank/view/{id}", yankView)
	mux.HandleFunc("GET	/yank/create", yankCreate)
	mux.HandleFunc("POST /yank/create", yankCreateItem)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
