package main

import (
	"fmt"
	"strconv"
	"net/http"
	"html/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request){
	w.Header().Add("server", "Go")
	
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
		"./ui/html/partials/nav.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w,r,err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w,r,err)
	}
	
}

func (app *application) yankView(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific yank with ID %d", id)

}

func (app *application) yankCreate(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("display a form for creating an yank"))
}

func (app *application) yankCreateItem(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("save a new yank"))
}

