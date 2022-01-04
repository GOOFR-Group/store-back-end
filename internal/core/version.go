package core

import "github.com/GOOFR-Group/store-back-end/internal/oapi"

// Version returns the current API version of this Server
func Version() oapi.VersionSchema {
	return oapi.VersionSchema{
		Version: "1.0.0",
		Notes:   "Initial version",
	}
}
