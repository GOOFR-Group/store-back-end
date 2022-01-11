package conf

import "github.com/go-chi/chi/v5"

var router *chi.Mux

// RouterPathAPIPrefix represents the path prefix of this API
const RouterPathAPIPrefix = "/api"

// RouterPathAPIPrefix represents the path prefix of this API's documentation
const RouterPathDocsPrefix = "/docs"

// InitRouter creates a new router
func InitRouter() {
	router = chi.NewRouter()
}

// Router retrieves the router
func Router() *chi.Mux {
	return router
}
