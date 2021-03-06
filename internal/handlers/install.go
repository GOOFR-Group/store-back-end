package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-openapi/runtime/middleware"
	"github.com/goofr-group/store-back-end/internal/conf"
	"github.com/goofr-group/store-back-end/internal/oapi"
)

const docsFolder = "docs"
const specFile = "store.yaml"
const specURL = "/" + specFile

// StoreImpl represents the implementation of all the server handlers
type StoreImpl struct{}

// Register registers the server handlers
func Register() error {
	router := conf.Router()
	if router == nil {
		return fmt.Errorf("router is not initialized")
	}

	// register swagger spec
	specPath := filepath.Join(conf.StaticPath(), docsFolder, specFile)
	router.Get(specURL, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, specPath)
	})

	// register docs
	docOpts := middleware.SwaggerUIOpts{SpecURL: specURL, Path: ""}
	docHandler := middleware.SwaggerUI(docOpts, nil)
	router.Mount(conf.RouterPathDocsPrefix, docHandler)

	// register api
	var storeApi StoreImpl
	router.Mount(conf.RouterPathAPIPrefix, oapi.Handler(&storeApi))

	return nil
}
