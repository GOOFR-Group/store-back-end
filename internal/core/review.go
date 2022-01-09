package core

import (
	"database/sql"
	"fmt"

	"github.com/gocraft/dbr/v2"
	"github.com/goofr-group/store-back-end/internal/oapi"
	"github.com/goofr-group/store-back-end/internal/storage"
	"github.com/google/uuid"
)

// PostReview creates a new review
func PostReview(req oapi.PostReviewJSONRequestBody) (oapi.ReviewSchema, error) {
	var id uuid.UUID
	var err error

	if id, err = uuid.NewRandom(); err != nil {
		return oapi.ReviewSchema{}, fmt.Errorf(ErrGeneratingUUID, err.Error())
	}

	var idGame uuid.UUID
	if idGame, err = uuid.Parse(req.IdGame); err != nil {
		return oapi.ReviewSchema{}, err
	}

	var idClient uuid.UUID
	if idClient, err = uuid.Parse(req.IdClient); err != nil {
		return oapi.ReviewSchema{}, err
	}

	var object storage.Review

	if err = handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var ok bool

		if _, ok, err = storage.ReadGameByID(tx, idGame); err != nil {
			return err
		}
		if !ok {
			return ErrGameNotFound
		}

		if _, ok, err = storage.ReadClientByID(tx, idClient); err != nil {
			return err
		}
		if !ok {
			return ErrClientNotFound
		}

		if _, ok, err = storage.ReadReviewByGameIDAndClientID(tx, idGame, idClient); err != nil {
			return err
		}
		if ok {
			return ErrObjectAlreadyCreated
		}

		var stars int64
		var starsOK bool
		if req.Stars != nil {
			stars = *req.Stars
			starsOK = true
		}

		var review string
		var reviewOK bool
		if req.Review != nil {
			review = *req.Review
			reviewOK = true
		}

		if err = storage.CreateReview(tx, storage.Review{
			ID:       id,
			IDGame:   idGame,
			IDClient: idClient,
			Stars: sql.NullInt64{
				Int64: stars,
				Valid: starsOK,
			},
			Review: sql.NullString{
				String: review,
				Valid:  reviewOK,
			},
		}); err != nil {
			return err
		}

		if object, ok, err = storage.ReadReviewByID(tx, id); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		return nil
	}); err != nil {
		return oapi.ReviewSchema{}, err
	}

	return getReviewFromModel(object), nil
}

// PutReview updates a review
func PutReview(params oapi.PutReviewParams, req oapi.PutReviewJSONRequestBody) (oapi.ReviewSchema, error) {
	var id uuid.UUID
	var err error

	if id, err = uuid.Parse(params.Id); err != nil {
		return oapi.ReviewSchema{}, err
	}

	var object storage.Review

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var stars int64
		var starsOK bool
		if req.Stars != nil {
			stars = *req.Stars
			starsOK = true
		}

		var review string
		var reviewOK bool
		if req.Review != nil {
			review = *req.Review
			reviewOK = true
		}

		if err = storage.UpdateReviewByID(tx, storage.Review{
			ID: id,
			Stars: sql.NullInt64{
				Int64: stars,
				Valid: starsOK,
			},
			Review: sql.NullString{
				String: review,
				Valid:  reviewOK,
			},
		}); err != nil {
			return err
		}

		var ok bool

		if object, ok, err = storage.ReadReviewByID(tx, id); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		return nil
	}); err != nil {
		return oapi.ReviewSchema{}, err
	}

	return getReviewFromModel(object), nil
}

// DeleteReview deletes a review
func DeleteReview(params oapi.DeleteReviewParams) (oapi.ReviewSchema, error) {
	var id uuid.UUID
	var err error

	if id, err = uuid.Parse(params.Id); err != nil {
		return oapi.ReviewSchema{}, err
	}

	var object storage.Review

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var ok bool

		if object, ok, err = storage.ReadReviewByID(tx, id); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		if err = storage.DeleteReviewByID(tx, id); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return oapi.ReviewSchema{}, err
	}

	return getReviewFromModel(object), nil
}

// GetGameReviews gets all the reviews for a game
func GetGameReviews(params oapi.GetGameReviewsParams) ([]oapi.ReviewSchema, error) {
	var idGame uuid.UUID
	var err error

	if idGame, err = uuid.Parse(params.Id); err != nil {
		return nil, err
	}

	var objects []storage.Review

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var ok bool

		if _, ok, err = storage.ReadGameByID(tx, idGame); err != nil {
			return err
		}
		if !ok {
			return ErrGameNotFound
		}

		if objects, err = storage.ReadReviewsByGameID(tx, idGame); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return getReviewsFromModel(objects), nil
}

func getReviewFromModel(model storage.Review) oapi.ReviewSchema {
	id := model.ID.String()
	return oapi.ReviewSchema{
		Id:       &id,
		IdGame:   model.IDGame.String(),
		IdClient: model.IDClient.String(),
		Stars:    &model.Stars.Int64,
		Review:   &model.Review.String,
	}
}

func getReviewsFromModel(model []storage.Review) []oapi.ReviewSchema {
	array := make([]oapi.ReviewSchema, len(model))
	for i, m := range model {
		array[i] = getReviewFromModel(m)
	}
	return array
}
