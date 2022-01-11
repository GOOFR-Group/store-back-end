package core

import (
	"github.com/gocraft/dbr/v2"
	"github.com/goofr-group/store-back-end/internal/oapi"
	"github.com/goofr-group/store-back-end/internal/storage"
)

const topReviewedGamesLimit = 100

// GetTopReviews gets the top reviewed games
func GetTopReviews() ([]oapi.GameSchema, error) {
	var objects []storage.Game

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var err error

		if objects, err = storage.ReadGamesOrderByAvgReviewDesc(tx, topReviewedGamesLimit); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return getGamesFromModel(objects), nil
}
