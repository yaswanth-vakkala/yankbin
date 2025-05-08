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

	logger.Info("starting server", "addr", *addr)
	
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
