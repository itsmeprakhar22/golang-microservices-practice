package app

import (
	"net/http"

	"./controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
}
