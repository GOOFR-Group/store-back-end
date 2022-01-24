package core

import (
	"github.com/gocraft/dbr/v2"
	"github.com/goofr-group/store-back-end/internal/oapi"
	"github.com/goofr-group/store-back-end/internal/storage"
	"github.com/goofr-group/store-back-end/internal/utils/mathi"
)

const topReviewedGamesLimit = 100

// GetTopReviews gets the top reviewed games
func GetTopReviews() ([]oapi.TopReviewsSchema, error) {
	var games []storage.Game
	var averages []float64

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var err error

		if games, averages, err = storage.ReadGamesAndAverageOrderByAvgReviewDesc(tx, topReviewedGamesLimit); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	len := mathi.Min(len(games), len(averages))
	topReviews := make([]oapi.TopReviewsSchema, len)
	for i := 0; i < len; i++ {
		topReviews[i] = oapi.TopReviewsSchema{
			Game:    getGameFromModel(games[i]),
			Average: averages[i],
		}
	}

	return topReviews, nil
}
