package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

// StartApp ...
func StartApp() {
	mapUrls()

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
