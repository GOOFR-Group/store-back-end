package core

import (
	"github.com/gocraft/dbr/v2"
	"github.com/goofr-group/store-back-end/internal/oapi"
	"github.com/goofr-group/store-back-end/internal/storage"
	"github.com/goofr-group/store-back-end/internal/utils/mathi"
)

const statisticGamesLimit = 100

// GetTopReviews gets the top reviewed games
func GetTopReviews() ([]oapi.TopReviewsSchema, error) {
	var games []storage.Game
	var averages []float64

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var err error

		if games, averages, err = storage.ReadGamesAndAverageOrderByAvgReviewDesc(tx, statisticGamesLimit); err != nil {
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

// GetBestSellers gets the bestselling games
func GetBestSellers() ([]oapi.BestSellersSchema, error) {
	var games []storage.Game
	var sales []int64

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var err error

		if games, sales, err = storage.ReadGamesAndSalesOrderBySalesDesc(tx, statisticGamesLimit); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	len := mathi.Min(len(games), len(sales))
	bestSellers := make([]oapi.BestSellersSchema, len)
	for i := 0; i < len; i++ {
		bestSellers[i] = oapi.BestSellersSchema{
			Game:  getGameFromModel(games[i]),
			Sales: sales[i],
		}
	}

	return bestSellers, nil
}
