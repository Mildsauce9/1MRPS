package main

//import

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//run

//mount

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	//middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	return r
}

//config

type application struct {
	config config
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
