package services

import (
	"github.com/marcosgeo/microservices/mvc/domain"
	"github.com/marcosgeo/microservices/mvc/utils"
)

// GetUser ...
func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
