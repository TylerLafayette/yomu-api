package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/imroc/req"

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
	return &Router{
		router,
	}
}

// CreateRouter creates and returns a Router struct with the default options/handlers.
func CreateRouter() *Router {
	return NewRouter(defaultChiRouter())
}

// newCORS returns a new CORS handler.
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

	// Attach the static folder handler.
	r.Mount("/public", http.StripPrefix("/public", staticHandler()))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.PlainText(w, r, "hello world :)")
	})

	r.Get("/dictionary", dictionary)

	return r
}

// staticHandler returns a FileServer for the public folder.
func staticHandler() http.Handler {
	return http.FileServer(http.Dir("public"))
}

// dictionary serves the /dictionary route.
func dictionary(w http.ResponseWriter, r *http.Request) {
	keyword, ok := r.URL.Query()["keyword"]
	if !ok {
		log.Println("Keyword failed")
		render.Status(r, 404)
		return
	}

	req := req.New()

	log.Println(fmt.Sprintf("https://jisho.org/api/v1/search/words?keyword=%v", keyword[0]))
	resp, err := req.Get(fmt.Sprintf("https://jisho.org/api/v1/search/words?keyword=%v", url.QueryEscape(keyword[0])))
	if err != nil {
		log.Println(err)
		render.Status(r, 404)
		return
	}

	bytes, err := resp.ToBytes()
	if err != nil {
		log.Println(err)
		render.Status(r, 404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
