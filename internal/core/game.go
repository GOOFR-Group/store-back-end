package core

import (
	"fmt"
	"net/smtp"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/gocraft/dbr/v2"
	"github.com/goofr-group/store-back-end/internal/conf"
	"github.com/goofr-group/store-back-end/internal/oapi"
	"github.com/goofr-group/store-back-end/internal/storage"
	"github.com/google/uuid"
)

// PostGame creates a new game
func PostGame(req oapi.PostGameJSONRequestBody) (oapi.GameSchema, error) {
	var id uuid.UUID
	var err error

	if id, err = uuid.NewRandom(); err != nil {
		return oapi.GameSchema{}, fmt.Errorf(ErrGeneratingUUID, err.Error())
	}

	var idPublisher uuid.UUID
	if idPublisher, err = uuid.Parse(req.IdPublisher); err != nil {
		return oapi.GameSchema{}, err
	}

	var object storage.Game

	if err = handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var ok bool

		if _, ok, err = storage.ReadPublisherByID(tx, idPublisher); err != nil {
			return err
		}
		if !ok {
			return ErrPublisherNotFound
		}

		if err = storage.CreateGame(tx, storage.Game{
			ID:           id,
			IDPublisher:  idPublisher,
			Name:         req.Name,
			Price:        req.Price,
			Discount:     req.Discount,
			State:        storage.StateGame(req.State),
			CoverImage:   req.CoverImage,
			ReleaseDate:  req.ReleaseDate.Time,
			Description:  req.Description,
			DownloadLink: req.DownloadLink,
		}); err != nil {
			return err
		}

		if object, ok, err = storage.ReadGameByID(tx, id); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		return nil
	}); err != nil {
		return oapi.GameSchema{}, err
	}

	return getGameFromModel(object), nil
}

// GetGame gets a game
func GetGame(params oapi.GetGameParams) ([]oapi.GameSchema, error) {
	var objects []storage.Game

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var err error

		if params.Id == nil {
			if objects, err = storage.ReadGames(tx); err != nil {
				return err
			}
		} else {
			var id uuid.UUID

			if id, err = uuid.Parse(*params.Id); err != nil {
				return err
			}

			var object storage.Game
			var ok bool

			if object, ok, err = storage.ReadGameByID(tx, id); err != nil {
				return err
			}
			if !ok {
				return ErrObjectNotFound
			}

			objects = append(objects, object)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return getGamesFromModel(objects), nil
}

// PutGame updates a game
func PutGame(params oapi.PutGameParams, req oapi.PutGameJSONRequestBody) (oapi.GameSchema, error) {
	var id uuid.UUID
	var idPublisher uuid.UUID
	var err error

	if id, err = uuid.Parse(params.Id); err != nil {
		return oapi.GameSchema{}, err
	}

	if idPublisher, err = uuid.Parse(req.IdPublisher); err != nil {
		return oapi.GameSchema{}, err
	}

	var prevObject storage.Game
	var object storage.Game

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var ok bool

		if _, ok, err = storage.ReadPublisherByID(tx, idPublisher); err != nil {
			return err
		}
		if !ok {
			return ErrPublisherNotFound
		}

		if prevObject, ok, err = storage.ReadGameByID(tx, id); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		if err = storage.UpdateGameByID(tx, storage.Game{
			ID:           id,
			IDPublisher:  idPublisher,
			Name:         req.Name,
			Price:        req.Price,
			Discount:     req.Discount,
			State:        storage.StateGame(req.State),
			CoverImage:   req.CoverImage,
			ReleaseDate:  req.ReleaseDate.Time,
			Description:  req.Description,
			DownloadLink: req.DownloadLink,
		}); err != nil {
			return err
		}

		if object, ok, err = storage.ReadGameByID(tx, id); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		return nil
	}); err != nil {
		return oapi.GameSchema{}, err
	}

	if prevObject.State == storage.StateGameUpcoming && object.State == storage.StateGameActive {
		var emails []string

		if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
			if emails, err = storage.ReadClientEmailsByGameInLibrary(tx, object.ID); err != nil {
				return err
			}

			return nil
		}); err != nil {
			return oapi.GameSchema{}, err
		}

		title := "GOOFR Store - Game Release"

		body := "You don't have to wait any longer!\n"
		body += "One of the games you had bought has just been released.\n"
		body += "Game:\n"
		body += fmt.Sprintf("\tName: %s\n", object.Name)
		body += fmt.Sprintf("\tDescription: %s\n", object.Description)
		body += fmt.Sprintf("\tRelease Date: %s\n", object.ReleaseDate.Format(timeLayout))
		body += fmt.Sprintf("\tPrice: %.2f", object.Price)

		message := []byte(fmt.Sprintf(smtpSubject, title) + body)
		if err = smtp.SendMail(conf.SMTPAddress(), conf.SMTPAuthentication(), conf.SMTPEmailAddress(), emails, message); err != nil {
			return oapi.GameSchema{}, err
		}
	}

	if prevObject.Discount != object.Discount && object.Discount > 0 {
		var emails []string

		if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
			if emails, err = storage.ReadClientEmailsByGameInWishlis(tx, object.ID); err != nil {
				return err
			}

			return nil
		}); err != nil {
			return oapi.GameSchema{}, err
		}

		title := "GOOFR Store - Game Discount"

		body := "We have got good news!\n"
		body += "One of the games on your wishlist just got a discount.\n"
		body += "Game:\n"
		body += fmt.Sprintf("\tName: %s\n", object.Name)
		body += fmt.Sprintf("\tDescription: %s\n", object.Description)
		body += fmt.Sprintf("\tPrice: %.2f\n", object.Price)
		body += fmt.Sprintf("\tDiscount: %.2f%%", object.Discount)

		message := []byte(fmt.Sprintf(smtpSubject, title) + body)
		if err = smtp.SendMail(conf.SMTPAddress(), conf.SMTPAuthentication(), conf.SMTPEmailAddress(), emails, message); err != nil {
			return oapi.GameSchema{}, err
		}
	}

	return getGameFromModel(object), nil
}

// DeleteGame deletes a game
func DeleteGame(params oapi.DeleteGameParams) (oapi.GameSchema, error) {
	var id uuid.UUID
	var err error

	if id, err = uuid.Parse(params.Id); err != nil {
		return oapi.GameSchema{}, err
	}

	var object storage.Game

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var ok bool

		if object, ok, err = storage.ReadGameByID(tx, id); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		if err = storage.DeleteGameByID(tx, id); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return oapi.GameSchema{}, err
	}

	return getGameFromModel(object), nil
}

// PostGameTag adds a tag to a game
func PostGameTag(params oapi.PostGameTagParams) error {
	var idGame uuid.UUID
	var idTag uuid.UUID
	var err error

	if idGame, err = uuid.Parse(params.GameID); err != nil {
		return err
	}

	if idTag, err = uuid.Parse(params.TagID); err != nil {
		return err
	}

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var ok bool

		if _, ok, err = storage.ReadGameByID(tx, idGame); err != nil {
			return err
		}
		if !ok {
			return ErrGameNotFound
		}

		if _, ok, err = storage.ReadTagByID(tx, idTag); err != nil {
			return err
		}
		if !ok {
			return ErrTagNotFound
		}

		if _, ok, err = storage.ReadTagGameByID(tx, idTag, idGame); err != nil {
			return err
		}
		if ok {
			return ErrObjectAlreadyCreated
		}

		if err = storage.CreateTagGame(tx, idTag, idGame); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// GetGameTag gets all tags of a game
func GetGameTag(params oapi.GetGameTagParams) ([]oapi.TagSchema, error) {
	var idGame uuid.UUID
	var err error

	if idGame, err = uuid.Parse(params.Id); err != nil {
		return nil, err
	}

	var objects []storage.Tag

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var ok bool

		if _, ok, err = storage.ReadGameByID(tx, idGame); err != nil {
			return err
		}
		if !ok {
			return ErrGameNotFound
		}

		if objects, err = storage.ReadTagsByGameID(tx, idGame); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return getTagsFromModel(objects), nil
}

// DeleteGameTag removes a tag from a game
func DeleteGameTag(params oapi.DeleteGameTagParams) (oapi.TagSchema, error) {
	var idGame uuid.UUID
	var idTag uuid.UUID
	var err error

	if idGame, err = uuid.Parse(params.GameID); err != nil {
		return oapi.TagSchema{}, err
	}

	if idTag, err = uuid.Parse(params.TagID); err != nil {
		return oapi.TagSchema{}, err
	}

	var object storage.Tag

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var ok bool

		if _, ok, err = storage.ReadTagGameByID(tx, idTag, idGame); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		if _, ok, err = storage.ReadGameByID(tx, idGame); err != nil {
			return err
		}
		if !ok {
			return ErrGameNotFound
		}

		if object, ok, err = storage.ReadTagByID(tx, idTag); err != nil {
			return err
		}
		if !ok {
			return ErrTagNotFound
		}

		if err = storage.DeleteTagGameByID(tx, idTag, idGame); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return oapi.TagSchema{}, err
	}

	return getTagFromModel(object), nil
}

// PostGameImage adds an image to a game
func PostGameImage(params oapi.PostGameImageParams) error {
	var idGame uuid.UUID
	var err error

	if idGame, err = uuid.Parse(params.GameID); err != nil {
		return err
	}

	var id uuid.UUID
	if id, err = uuid.NewRandom(); err != nil {
		return fmt.Errorf(ErrGeneratingUUID, err.Error())
	}

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var ok bool

		if _, ok, err = storage.ReadGameByID(tx, idGame); err != nil {
			return err
		}
		if !ok {
			return ErrGameNotFound
		}

		if err = storage.CreateImageGame(tx, storage.ImageGame{
			ID:     id,
			IDGame: idGame,
			Image:  params.Image,
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// GetGameImage gets all images of a game
func GetGameImage(params oapi.GetGameImageParams) ([]oapi.GameImageSchema, error) {
	var idGame uuid.UUID
	var err error

	if idGame, err = uuid.Parse(params.Id); err != nil {
		return nil, err
	}

	var objects []storage.ImageGame

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var ok bool

		if _, ok, err = storage.ReadGameByID(tx, idGame); err != nil {
			return err
		}
		if !ok {
			return ErrGameNotFound
		}

		if objects, err = storage.ReadImageGamesByGameID(tx, idGame); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return getGameImagesFromModel(objects), nil
}

// DeleteGameImage removes an image from a game
func DeleteGameImage(params oapi.DeleteGameImageParams) (oapi.GameImageSchema, error) {
	var id uuid.UUID
	var err error

	if id, err = uuid.Parse(params.Id); err != nil {
		return oapi.GameImageSchema{}, err
	}

	var object storage.ImageGame

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var ok bool

		if object, ok, err = storage.ReadImageGameByID(tx, id); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		if err = storage.DeleteImageGameByID(tx, id); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return oapi.GameImageSchema{}, err
	}

	return getGameImageFromModel(object), nil
}

func getGameFromModel(model storage.Game) oapi.GameSchema {
	id := model.ID.String()
	return oapi.GameSchema{
		Id:           &id,
		IdPublisher:  model.IDPublisher.String(),
		Name:         model.Name,
		Price:        model.Price,
		Discount:     model.Discount,
		State:        oapi.GameSchemaState(model.State),
		CoverImage:   model.CoverImage,
		ReleaseDate:  openapi_types.Date{Time: model.ReleaseDate},
		Description:  model.Description,
		DownloadLink: model.DownloadLink,
	}
}

func getGamesFromModel(model []storage.Game) []oapi.GameSchema {
	array := make([]oapi.GameSchema, len(model))
	for i, m := range model {
		array[i] = getGameFromModel(m)
	}
	return array
}

func getGameImageFromModel(model storage.ImageGame) oapi.GameImageSchema {
	return oapi.GameImageSchema{
		Id:     model.ID.String(),
		IdGame: model.IDGame.String(),
		Image:  model.Image,
	}
}

func getGameImagesFromModel(model []storage.ImageGame) []oapi.GameImageSchema {
	array := make([]oapi.GameImageSchema, len(model))
	for i, m := range model {
		array[i] = getGameImageFromModel(m)
	}
	return array
}
