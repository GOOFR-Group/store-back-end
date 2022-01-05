package core

import (
	"fmt"

	"github.com/GOOFR-Group/store-back-end/internal/oapi"
	"github.com/GOOFR-Group/store-back-end/internal/storage"
	"github.com/gocraft/dbr/v2"
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
		return storage.CreateNewTag(tx, storage.Tag{
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
			if objects, err = storage.GetAllTags(tx); err != nil {
				return err
			}
		} else {
			id, err := uuid.Parse(*params.Id)
			if err != nil {
				return err
			}

			var object storage.Tag
			var ok bool
			object, ok, err = storage.GetTagByID(tx, id)
			if err != nil {
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
	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var id uuid.UUID
		var err error
		var ok bool

		id, err = uuid.Parse(params.Id)
		if err != nil {
			return err
		}

		ok, err = storage.UpdateTagByID(tx, storage.Tag{
			ID:   id,
			Name: req.Name,
		})
		if err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		_, ok, err = storage.GetTagByID(tx, id)
		if err != nil {
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
	var object storage.Tag

	if err := handleTransaction(nil, func(tx dbr.SessionRunner) error {
		var err error
		var ok bool

		id, err := uuid.Parse(params.Id)
		if err != nil {
			return err
		}

		object, ok, err = storage.GetTagByID(tx, id)
		if err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
		}

		ok, err = storage.DeleteTagByID(tx, id)
		if err != nil {
			return err
		}
		if !ok {
			return ErrObjectNotFound
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
