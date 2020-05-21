package domain

import (
	"errors"
	"fmt"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Antonio", LastName: "Marcos", Email: "marcosgeo@yahoo.com"},
	}
)

// GetUser ...
func GetUser(userID int64) (*User, error) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, errors.New(fmt.Sprintf("User %v was not found", userID))
}
