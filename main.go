package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello from yankbin"))
}

func yankView(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("display a specific yank"))
}

func yankCreate(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("display a form for creating an yank"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/yank/view", yankView)
	mux.HandleFunc("/yank/create", yankCreate)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
