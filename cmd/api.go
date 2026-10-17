package main

//import

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//mount

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	//middleware

	r.Use(middleware.RequestID) // helps with rate limiting
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//This is the route itself, it accepts a route and the http function that needs to be
	// compiled when you call the route
	// I think this can be written better, we could define the function in another package and
	// import the functions just for this file, ideally I would want each API or atleast each
	// module of the API to have the same file. So if we have a db, we need to split this by types.
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good!")) // This is the repsonse body that'll be returned to the request.
	})

	return r
}

// run
func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	return srv.ListenAndServe()
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
