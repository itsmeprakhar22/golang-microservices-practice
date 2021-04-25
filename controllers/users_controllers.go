package controllers

import (
	"GolangWorkspace/go-microservices/mvc/services"
	"encoding/json"
	"net/http"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	user, err := services.GetUser(123)

	jsonVal, err := json.Marshal(user)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("User ID must be number"))
	}
	resp.Write(jsonVal)
}
