package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

// Router contains a chi Router.
type Router struct {
	router *chi.Mux
}

// NewRouter creates and returns a new Router struct with the supplied chi Router.
func NewRouter(router *chi.Mux) *Router {
	return &Router{router}
}

// CreateRouter creates and returns a Router struct with the default options/handlers.
func CreateRouter() *Router {
	return NewRouter(defaultChiRouter())
}

func newCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-User-ID"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
}

// defaultChiRouter constructs the default router with handlers.
func defaultChiRouter() *chi.Mux {
	r := chi.NewRouter()

	// Create a new CORS handler and apply it to the Router.
	cors := newCORS()
	r.Use(cors.Handler)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.PlainText(w, r, "hello")
	})

	return r
}
