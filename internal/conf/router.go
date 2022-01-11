package conf

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

var router *chi.Mux

// RouterPathAPIPrefix represents the path prefix of this API
const RouterPathAPIPrefix = "/api"

// RouterPathAPIPrefix represents the path prefix of this API's documentation
const RouterPathDocsPrefix = "/docs"

// CORSOptions represents our handler (CORS) options
var CORSOptions = cors.Options{
	AllowedOrigins: []string{"https://*", "http://*"},
	AllowedMethods: []string{http.MethodPost, http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodOptions},
}

// InitRouter creates a new router
func InitRouter() {
	router = chi.NewRouter()
	router.Use(cors.Handler(CORSOptions))
}

// Router retrieves the router
func Router() *chi.Mux {
	return router
}
