package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// recovers from panic
	mux.Use(middleware.Recoverer)

	mux.Get("/", app.Home)

	mux.Get("/movies", app.AllMovies)
	return mux
}
