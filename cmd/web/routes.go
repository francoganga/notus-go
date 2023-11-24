package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {

	r := chi.NewRouter()
	fs := http.FileServer(http.Dir("./public"))

	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		a.templates.Render("views_home", w, nil)
	})

	return r
}

