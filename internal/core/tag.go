package core

import (
	"fmt"

	"github.com/gocraft/dbr/v2"
	"github.com/goofr-group/store-back-end/internal/oapi"
	"github.com/goofr-group/store-back-end/internal/storage"
	"github.com/google/uuid"
)

// PostTag creates a new tag
func PostTag(req oapi.PostTagJSONRequestBody) error {
	var id uuid.UUID
	var err error

	if id, err = uuid.NewRandom(); err != nil {
		return fmt.Errorf(ErrGeneratingUUID, err.Error())
	}

	if err = handleTransaction(nil, func(tx dbr.SessionRunner) error {
		return storage.CreateTag(tx, storage.Tag{
			ID:   id,
			Name: req.Name,
		})
	}); err != nil {
		return err
	}

	return nil
}

// GetTag gets a tag
func GetTag(params oapi.GetTagParams) ([]oapi.TagSchema, error) {
	var objects []storage.Tag

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var err error

		if params.Id == nil {
			if objects, err = storage.ReadTags(tx); err != nil {
				return err
			}
		} else {
			var id uuid.UUID

			if id, err = uuid.Parse(*params.Id); err != nil {
				return err
			}

			var object storage.Tag
			var ok bool

			if object, ok, err = storage.ReadTagByID(tx, id); err != nil {
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

	return getTagsFromModel(objects), nil
}

// PutTag updates a tag
func PutTag(params oapi.PutTagParams, req oapi.PutTagJSONRequestBody) error {
	var id uuid.UUID
	var err error

	if id, err = uuid.Parse(params.Id); err != nil {
		return err
	}

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		if err = storage.UpdateTagByID(tx, storage.Tag{
			ID:   id,
			Name: req.Name,
		}); err != nil {
			return err
		}

		var ok bool

		if _, ok, err = storage.ReadTagByID(tx, id); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// DeleteTag deletes a tag
func DeleteTag(params oapi.DeleteTagParams) (oapi.TagSchema, error) {
	var id uuid.UUID
	var err error

	if id, err = uuid.Parse(params.Id); err != nil {
		return oapi.TagSchema{}, err
	}

	var object storage.Tag

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var ok bool

		if object, ok, err = storage.ReadTagByID(tx, id); err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		if err = storage.DeleteTagByID(tx, id); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return oapi.TagSchema{}, err
	}

	return getTagFromModel(object), nil
}

func getTagFromModel(model storage.Tag) oapi.TagSchema {
	id := model.ID.String()
	return oapi.TagSchema{
		Id:   &id,
		Name: model.Name,
	}
}

func getTagsFromModel(model []storage.Tag) []oapi.TagSchema {
	array := make([]oapi.TagSchema, len(model))
	for i, m := range model {
		array[i] = getTagFromModel(m)
	}
	return array
}
