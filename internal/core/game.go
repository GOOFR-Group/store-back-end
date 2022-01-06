package core

import (
	"github.com/GOOFR-Group/store-back-end/internal/oapi"
	"github.com/GOOFR-Group/store-back-end/internal/storage"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

func getGameFromModel(model storage.Game) oapi.GameSchema {
	id := model.ID.String()
	return oapi.GameSchema{
		Id:          &id,
		IdPublisher: model.IDPublisher.String(),
		Name:        model.Name,
		Price:       model.Price,
		Discount:    model.Discount,
		State:       oapi.GameSchemaState(model.State),
		CoverImage:  model.CoverImage,
		ReleaseDate: openapi_types.Date{
			Time: model.ReleaseDate,
		},
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
