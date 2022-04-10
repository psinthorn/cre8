package cre8

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (c *Cre8) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	if c.Debug {
		mux.Use(middleware.Logger)
	}
	mux.Use(middleware.Recoverer)

	// mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Welcome to Cre8")
	// })

	return mux
}
