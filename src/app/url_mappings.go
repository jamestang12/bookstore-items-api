package app

import (
	"net/http"

	"../controllers"
)

func mapUrls() {
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)

}
