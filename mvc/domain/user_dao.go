package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marcosgeo/microservices/mvc/utils"
)

var (
	// the "database"
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Antonio", LastName: "Marcos", Email: "marcosgeo@yahoo.com"},
	}

	// UserDao is a userDaoInterface variable and because of this has all of their methods
	UserDao userDaoInterface
)

// This is an interface, which facilitates mock implementations on tests cases
type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

// GetUser is an userDao method. This method met the requirements of userServiceInterface
func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	log.Println(" ### We're accessing the database ###")
	// accessing the "database 'users'"
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User with id %v doesn't exists", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
