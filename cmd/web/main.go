package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static",fileServer))
	mux.HandleFunc("GET	/{$}", home)
	mux.HandleFunc("GET	/yank/view/{id}", yankView)
	mux.HandleFunc("GET	/yank/create", yankCreate)
	mux.HandleFunc("POST /yank/create", yankCreateItem)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
