package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static",fileServer))
	mux.HandleFunc("GET	/{$}", home)
	mux.HandleFunc("GET	/yank/view/{id}", yankView)
	mux.HandleFunc("GET	/yank/create", yankCreate)
	mux.HandleFunc("POST /yank/create", yankCreateItem)

	log.Print("starting server on %s", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
