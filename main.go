package main

import (
	"fmt"
	"strconv"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
	w.Header().Add("server", "Go")
	w.Write([]byte("Hello from yankbin"))
}

func yankView(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("display yank with id %d", id)
	w.Write([]byte(msg))
}

func yankCreate(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("display a form for creating an yank"))
}

func yankCreateItem(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("save a new yank"))
}

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
