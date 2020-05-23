package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/marcosgeo/microservices/mvc/services"
	"github.com/marcosgeo/microservices/mvc/utils"
)

// GetUser ...
func GetUser(resp http.ResponseWriter, req *http.Request) {
	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		appErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(appErr)
		resp.WriteHeader(appErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	user, appErr := services.GetUser(userID)
	if err != nil {
		jsonValue, _ := json.Marshal(appErr)
		resp.WriteHeader(appErr.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
