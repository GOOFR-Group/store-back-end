package conf

import (
	"strconv"

	"github.com/GOOFR-Group/store-back-end/internal/app"
	"github.com/GOOFR-Group/store-back-end/internal/logging"
	"github.com/GOOFR-Group/store-back-end/internal/utils/env"
)

// AppPort represents the environment variable port of this API
const AppPort = app.Name + "_PORT"

// DefaultAppPort represents the default port of this API
const DefaultAppPort = "8080"

var port int

// InitServer starts the environment variables required for the Server
func InitServer() {
	portEnvValue := env.GetEnvOrDefault(AppPort, DefaultAppPort)

	portEnvValueInt, err := strconv.Atoi(portEnvValue)
	if err != nil {
		logging.AppLogger.Fatal().Msgf("%s variable has an invalid value. %v", AppPort, err)
	}

	port = portEnvValueInt
}

// GetPort retrieves the configured http port
func GetPort() int {
	return port
}
