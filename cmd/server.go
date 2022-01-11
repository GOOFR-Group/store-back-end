package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/goofr-group/store-back-end/internal/logging"
	"github.com/goofr-group/store-back-end/internal/server"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		select {
		case <-c:
		case <-ctx.Done():
		}
		cancel()
	}()

	if err := server.RunServer(ctx); err != nil {
		logging.AppLogger.Fatal().Err(err)
		os.Exit(1)
	}

	logging.AppLogger.Info().Msg("Server exited properly")
}
