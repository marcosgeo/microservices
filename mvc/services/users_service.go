package services

import (
	"github.com/marcosgeo/microservices/mvc/domain"
	"github.com/marcosgeo/microservices/mvc/utils"
)

type usersService struct{}

var (
	// UsersService ...
	UsersService usersService
)

// GetUser ...
func (u *usersService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDao.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
