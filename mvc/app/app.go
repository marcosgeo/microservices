package app

import (
	"log"
	"net/http"

	"github.com/marcosgeo/microservices/mvc/controllers"
)

// StartApp ...
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}
