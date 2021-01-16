package controllers

import (
	"fmt"
	"net/http"

	"../../../bookstore_oauth_go/oauth"
	"../domain/items"
	"../services"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemsController struct {
}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {

	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte("ok"))
	// return

	if err := oauth.AuthenticatRequest(r); err != nil {
		// Return error json to the user
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		// Return error json to the user
		return
	}
	fmt.Println(result)

}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
