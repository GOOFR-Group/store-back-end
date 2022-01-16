package core

import (
	"github.com/gocraft/dbr/v2"
	"github.com/goofr-group/store-back-end/internal/oapi"
	"github.com/goofr-group/store-back-end/internal/storage"
	"github.com/google/uuid"
)

const storeGamesLimit = 50

// GetYourStore client's main store
func GetYourStore(params oapi.GetYourStoreParams) (oapi.YourStoreSchema, error) {
	var idClient uuid.UUID
	var err error

	if idClient, err = uuid.Parse(params.Id); err != nil {
		return oapi.YourStoreSchema{}, err
	}

	var featuredGames []storage.Game
	var recommendedGames []storage.Game
	var specialOffersGames []storage.Game
	var discoverGames []storage.Game

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		if featuredGames, err = storage.ReadGamesMostPurchasedOrderByAvgReviewDesc(tx, storeGamesLimit); err != nil {
			return err
		}

		if recommendedGames, err = storage.ReadGamesRecommendedByClientID(tx, idClient, storeGamesLimit); err != nil {
			return err
		}

		if specialOffersGames, err = storage.ReadGamesWithDiscount(tx, storeGamesLimit); err != nil {
			return err
		}

		var ids []uuid.UUID
		for _, g := range featuredGames {
			ids = append(ids, g.ID)
		}
		for _, g := range recommendedGames {
			ids = append(ids, g.ID)
		}
		for _, g := range specialOffersGames {
			ids = append(ids, g.ID)
		}

		if discoverGames, err = storage.ReadGamesWithDifferentID(tx, ids, storeGamesLimit); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return oapi.YourStoreSchema{}, err
	}

	return oapi.YourStoreSchema{
		Featured:      getGamesFromModel(featuredGames),
		Recommended:   getGamesFromModel(recommendedGames),
		SpecialOffers: getGamesFromModel(specialOffersGames),
		Discover:      getGamesFromModel(discoverGames),
	}, nil
}

// GetNewStore returns the new games from the store
func GetNewStore(params oapi.GetNewStoreParams) (oapi.NewStoreSchema, error) {
	var featuredGames []storage.Game
	var newGames []storage.Game

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var err error

		if featuredGames, err = storage.ReadGamesOrderByReleaseDateDescAndByAvgReviewDesc(tx, storeGamesLimit); err != nil {
			return err
		}

		if params.Ids == nil {
			if newGames, err = storage.ReadGamesOrderByReleaseDateDesc(tx, storeGamesLimit); err != nil {
				return err
			}
		} else {
			var ids []uuid.UUID
			for _, id := range *params.Ids {
				var tempID uuid.UUID
				if tempID, err = uuid.Parse(id); err != nil {
					return err
				}
				ids = append(ids, tempID)
			}

			if newGames, err = storage.ReadGamesOrderByReleaseDateDescFilteredByTag(tx, ids, storeGamesLimit); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return oapi.NewStoreSchema{}, err
	}

	return oapi.NewStoreSchema{
		Featured: getGamesFromModel(featuredGames),
		New:      getGamesFromModel(newGames),
	}, nil
}

// GetNoteworthyStore returns featured games from the store
func GetNoteworthyStore(params oapi.GetNoteworthyStoreParams) (oapi.NoteworthyStoreSchema, error) {
	var featuredGames []storage.Game
	var noteworthyGames []storage.Game

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var err error

		if featuredGames, err = storage.ReadGamesMostPurchasedOrderByAvgReviewDesc(tx, storeGamesLimit); err != nil {
			return err
		}

		if params.Ids == nil {
			if noteworthyGames, err = storage.ReadGamesMostPurchased(tx, storeGamesLimit); err != nil {
				return err
			}
		} else {
			var ids []uuid.UUID
			for _, id := range *params.Ids {
				var tempID uuid.UUID
				if tempID, err = uuid.Parse(id); err != nil {
					return err
				}
				ids = append(ids, tempID)
			}

			if noteworthyGames, err = storage.ReadGamesMostPurchasedFilteredByTag(tx, ids, storeGamesLimit); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return oapi.NoteworthyStoreSchema{}, err
	}

	return oapi.NoteworthyStoreSchema{
		Featured:   getGamesFromModel(featuredGames),
		Noteworthy: getGamesFromModel(noteworthyGames),
	}, nil
}
