package services

import (
	"net/http"

	"github.com/marcosgeo/microservices/mvc/domain"
	"github.com/marcosgeo/microservices/mvc/utils"
)

type itemsService struct{}

var (
	// ItemsService ...
	ItemsService itemsService
)

// GetItem ...
func GetItem(itemID string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "Not implemented",
		StatusCode: http.StatusInternalServerError,
	}
}
