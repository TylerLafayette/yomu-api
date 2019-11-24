package main

import "net/http"

// App contains dependencies for the app.
type App struct {
	config *Config
	router *Router
}

// NewApp creates and returns a new App struct with the supplied Config and
// Router.
func NewApp(config *Config, router *Router) *App {
	return &App{
		config,
		router,
	}
}

// Listen listens on the app's port.
func (app *App) Listen() error {
	return http.ListenAndServe(app.config.port, app.router.router)
}
