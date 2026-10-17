package main

//import

import (
	"net/http"
	"time"

	"github.com/Mildsauce9/1MRPS/internal/games"
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

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good!")) // This is the repsonse body that'll be returned to the request.
	})

	gameProcessor := games.NewProcessor()
	gameHandler := games.NewHandler(gameProcessor)

	r.Get("/game", gameHandler.SelectGames)
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
