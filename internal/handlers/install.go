package handlers

import (
	"fmt"

	"github.com/GOOFR-Group/store-back-end/internal/conf"
	"github.com/GOOFR-Group/store-back-end/internal/oapi"
)

type StoreImpl struct{}

func Register() error {
	router := conf.GetRouter()
	if router == nil {
		return fmt.Errorf("router is not initialized")
	}

	var storeApi StoreImpl
	router.Mount(conf.RouterPathPrefix, oapi.Handler(&storeApi))

	return nil
}
