package logging

import (
	"os"

	"github.com/GOOFR-Group/store-back-end/internal/app"
	"github.com/rs/zerolog"
)

// AppLogger logger for this app
var AppLogger zerolog.Logger

// Initialize configures the logging system
func Initialize() {
	// log setup
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: zerolog.TimeFormatUnix}
	AppLogger = zerolog.New(output).With().Timestamp().Logger()

	// default level is info, unless dev flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if app.Dev {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
