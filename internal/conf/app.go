package conf

import (
	"path/filepath"

	"github.com/GOOFR-Group/store-back-end/internal/app"
	"github.com/GOOFR-Group/store-back-end/internal/utils/env"
)

const (
	AppPath       = app.Name + "_PATH"
	AppConfPath   = app.Name + "_CONF_PATH"
	AppStaticPath = app.Name + "_STATIC_PATH"
)

var (
	mainPath   string
	confPath   string
	staticPath string
)

// InitApp starts the environment variables required for the application
func InitApp() {
	mainPath = env.GetEnvOrPanic(AppPath)
	confPath = env.GetEnvOrPanic(AppConfPath)
	staticPath = env.GetEnvOrPanic(AppStaticPath)
	mainPath = filepath.Clean(mainPath)

	// if the configuration files' path is relative to the base path
	if !filepath.IsAbs(confPath) {
		confPath = filepath.Join(mainPath, confPath)
	}

	// if the static files' path is relative to the base path
	if !filepath.IsAbs(staticPath) {
		staticPath = filepath.Join(mainPath, staticPath)
	}
}

// GetPath retrieves the configured base app path
func GetPath() string {
	return mainPath
}

// GetConfPath retrieves the configuration files path
func GetConfPath() string {
	return confPath
}

// GetStaticPath retrieves the configured static files path
func GetStaticPath() string {
	return staticPath
}
