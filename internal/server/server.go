package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/cors"

	"github.com/GOOFR-Group/store-back-end/internal/conf"
	"github.com/GOOFR-Group/store-back-end/internal/core"
	"github.com/GOOFR-Group/store-back-end/internal/handlers"
	"github.com/GOOFR-Group/store-back-end/internal/logging"
	"github.com/GOOFR-Group/store-back-end/internal/storage"
)

// CORSOptions represents our handler (CORS) options
var CORSOptions = cors.Options{
	AllowedOrigins: []string{"*"},
	AllowedMethods: []string{http.MethodPost, http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodOptions},
}

// RunServer entry point
func RunServer(ctx context.Context) error {
	// initialize the logging subsystem first
	logging.Initialize()

	// then do each of the other major subsystems
	conf.InitApp()
	conf.InitServer()
	conf.InitDB()
	conf.InitSMTP()
	conf.InitRouter()
	storage.InitStorage()

	// register the rest services
	if err := handlers.Register(); err != nil {
		return err
	}

	port := conf.Port()

	// log this service as a whole
	logging.AppLogger.Info().Str("version", core.Version().Version).Str("notes", core.Version().Notes).Msg("GOOFR Store API")
	logging.AppLogger.Info().Msgf("Listening on port %d", port)

	// add CORS options
	handler := cors.New(CORSOptions).Handler(conf.Router())

	s := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      handler,
		Addr:         fmt.Sprintf(":%d", port),
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logging.AppLogger.Fatal().Msgf("Listen: %+s\n", err)
		}
	}()

	logging.AppLogger.Info().Msg("Server started")

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.Shutdown(ctx)
}
