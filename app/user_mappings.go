package app

import (
	"GolangWorkspace/go-microservices/mvc/controllers"
)

func mapUrls() {
	router.GET("/users", controllers.GetUser)
}
