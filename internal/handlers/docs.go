package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/GOOFR-Group/store-back-end/internal/conf"
)

const docsFile = "docs.html"

// GetDocs handles the documentation endpoint for this API
func (*StoreImpl) GetDocs(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join(conf.GetStaticPath(), "swagger") + string(filepath.Separator) + docsFile
	http.ServeFile(w, r, filePath)
}
