package controllers

import (
	"log"
	"net/http"
)

// GetUser ...
func GetUser(resp http.ResponseWriter, req *http.Request) {
	userID := req.URL.Query().Get("user_id")
	log.Printf("About to process user_id %v", userID)
}
