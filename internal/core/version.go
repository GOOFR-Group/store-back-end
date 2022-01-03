package core

import "github.com/GOOFR-Group/store-back-end/internal/oapi"

func GetVersion() oapi.VersionSchema {
	return oapi.VersionSchema{
		Version: "0.0.1",
		Notes:   "Initial version",
	}
}
