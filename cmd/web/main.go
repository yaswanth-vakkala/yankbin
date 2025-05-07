package main

import (
	"flag"
	"log/slog"
	"os"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout,nil))

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static",fileServer))
	mux.HandleFunc("GET	/{$}", home)
	mux.HandleFunc("GET	/yank/view/{id}", yankView)
	mux.HandleFunc("GET	/yank/create", yankCreate)
	mux.HandleFunc("POST /yank/create", yankCreateItem)

	logger.Info("starting server", "addr", *addr)
	
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
