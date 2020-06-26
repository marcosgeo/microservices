package app

import (
	"github.com/marcosgeo/microservices/src/api/controllers/polo"
	"github.com/marcosgeo/microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.POST("/repositories", repositories.CreateRepo)
}
