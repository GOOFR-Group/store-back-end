package conf

import "github.com/go-chi/chi"

var router *chi.Mux

const RouterPathPrefix = "/api"

func InitRouter() {
	router = chi.NewRouter()
}

func GetRouter() *chi.Mux {
	return router
}
