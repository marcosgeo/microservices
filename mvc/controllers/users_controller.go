package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/marcosgeo/microservices/mvc/services"
)

// GetUser ...
func GetUser(resp http.ResponseWriter, req *http.Request) {
	userID, err := strconv.Atoi(req.URL.Query().Get("user_id"))
	if err != nil {
		// handle the erro
		log.Printf("User not found")
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}
	user, err := services.GetUser(int64(userID))
	if err != nil {
		// handle the err and return
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
