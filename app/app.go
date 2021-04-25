package app

import (
	"net/http"

	"GolangWorkspace/go-microservices/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
}
