package app

import (
	"net/http"

	"GolangWorkspace/go-microservices/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
