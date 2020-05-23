package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser_NoUserFound(t *testing.T) {
	// initialization
	id := int64(0)

	// Exectution
	user, err := GetUser(id)

	// Validadation
	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "User with id 0 doesn't exists", err.Message)
}

func TestGetUser_NoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "Antonio", user.FirstName)
	assert.EqualValues(t, "Marcos", user.LastName)
	assert.EqualValues(t, "marcosgeo@yahoo.com", user.Email)

}
