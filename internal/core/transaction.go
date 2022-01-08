package core

import (
	"github.com/goofr-group/store-back-end/internal/storage"

	"github.com/gocraft/dbr/v2"
)

// handleTransaction handles a new transaction to the database
func handleTransaction(log dbr.EventReceiver, requestHandler func(ctx dbr.SessionRunner) error) error {
	db := storage.GetDatabase()

	// Begin a new transaction
	tx, err := db.NewSession(log).Begin()
	if err != nil {
		return err
	}
	defer tx.RollbackUnlessCommitted()

	if err = requestHandler(tx); err != nil {
		return err
	}

	return tx.Commit()
}
