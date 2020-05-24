package services

import (
	"net/http"
	"testing"

	"github.com/marcosgeo/microservices/mvc/domain"
	"github.com/marcosgeo/microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
)

var (
	getUserFunction func(userID int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct{}

// the mock version of get user signature
func (m *usersDaoMock) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userID)
}

func TestGetUser_NotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userID int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "User with id 100 doesn't exists",
		}
	}
	user, err := UsersService.GetUser(100)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "User with id 100 doesn't exists", err.Message)
}

func TestGetUser_NoError(t *testing.T) {
	getUserFunction = func(userID int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			ID: 123,
		}, nil
	}
	user, err := UsersService.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)
}
