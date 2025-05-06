package main

import (
	"fmt"
	"strconv"
	"net/http"
	"log"
	"html/template"
)

func home(w http.ResponseWriter, r *http.Request){
	w.Header().Add("server", "Go")
	
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	
}

func yankView(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific yank with ID %d", id)

}

func yankCreate(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("display a form for creating an yank"))
}

func yankCreateItem(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("save a new yank"))
}

