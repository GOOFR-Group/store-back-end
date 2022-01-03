package conf

import "github.com/go-chi/chi/v5"

var router *chi.Mux

// RouterPathPrefix represents the path prefix of this API
const RouterPathPrefix = "/api"

// InitRouter creates a new router
func InitRouter() {
	router = chi.NewRouter()
}

// GetRouter retrieves the router
func GetRouter() *chi.Mux {
	return router
}
