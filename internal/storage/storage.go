//go:generate tgcon gen -s

package storage

import (
	"github.com/gocraft/dbr/v2"
	_ "github.com/lib/pq" //Register postgres driver

	"github.com/GOOFR-Group/store-back-end/internal/conf"
	"github.com/GOOFR-Group/store-back-end/internal/logging"
)

// Transaction is a more descriptive alias for the database record's SessionRunner interface
type Transaction dbr.SessionRunner

// InitStorage starts a new database connection
func InitStorage() {
	dbConn, err := newConnection(conf.DbConnectionString())
	if err != nil {
		logging.AppLogger.Fatal().Msgf("Could not initialize storage - failed to open connection to database - %v", err)

		return
	}

	if err = dbConn.Ping(); err != nil {
		logging.AppLogger.Fatal().Msgf("Could not initialize storage - unable to verify connection to database - %v", err)

		return
	}

	logging.AppLogger.Info().Msg("Connected to database")

	database = dbConn
}

var database *dbr.Connection

// GetDatabase retrieves a database
func GetDatabase() *dbr.Connection {
	return database
}

func newConnection(connStr string) (*dbr.Connection, error) {
	return dbr.Open("postgres", connStr, nil)
}
