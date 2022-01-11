package conf

import (
	"path/filepath"

	"github.com/goofr-group/store-back-end/internal/app"
	"github.com/goofr-group/store-back-end/internal/utils/env"
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

// Path retrieves the configured base app path
func Path() string {
	return mainPath
}

// ConfPath retrieves the configuration files path
func ConfPath() string {
	return confPath
}

// StaticPath retrieves the configured static files path
func StaticPath() string {
	return staticPath
}
