package main

import (
	"flag"
	"log/slog"
	"os"
	"net/http"
)

type application struct{
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout,nil))

	app := &application{
		logger: logger,
	}

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static",fileServer))
	mux.HandleFunc("GET	/{$}",app.home)
	mux.HandleFunc("GET	/yank/view/{id}", app.yankView)
	mux.HandleFunc("GET	/yank/create", app.yankCreate)
	mux.HandleFunc("POST /yank/create", app.yankCreateItem)

	logger.Info("starting server", "addr", *addr)
	
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
