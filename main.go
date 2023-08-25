package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/flosch/pongo2/v6"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {

	appEnv := os.Getenv("APP_ENV")

	if appEnv == "" {
		appEnv = "dev"
	}

	fmt.Printf("appEnv=%v\n", appEnv)

	r := chi.NewRouter()
	fs := http.FileServer(http.Dir("./public"))

	r.Use(middleware.Logger)
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/hw", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]string{
			"message": "Hello, World!",
		})
	})

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {

		templ := pongo2.Must(pongo2.FromFile("templates/base.html"))

		err := templ.ExecuteWriter(pongo2.Context{
			"mode": appEnv,
		}, w)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	})
	http.ListenAndServe(":3000", r)
}

