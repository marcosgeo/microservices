package services

import "github.com/marcosgeo/microservices/mvc/domain"

// GetUser ...
func GetUser(userID int64) (*domain.User, error) {
	return domain.GetUser(userID)
}
