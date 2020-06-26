package repositories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcosgeo/microservices/src/api/domain/repositories"
	"github.com/marcosgeo/microservices/src/api/services"
	"github.com/marcosgeo/microservices/src/api/utils/errors"
)

// CreateRepo ...
func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepo(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}
