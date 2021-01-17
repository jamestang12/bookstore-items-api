package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"../../../bookstore_oauth_go/oauth"
	"../../../bookstore_utils_go/rest_errors"
	"../domain/items"
	"../services"
	"../utils/http_utils"
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
		http_utils.RespondJson(w, err.Status, err)
		// http_utils.RespondError(w, *err)

		return
	}

	sellerId := oauth.GetCallerId(r)

	if sellerId == 0 {
		// respErr := rest_errors.newUnauthorizedError("invalid request body")
		respErr := rest_errors.NewBadRequestError("unauthorized")
		http_utils.RespondError(w, respErr)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, respErr)
		return
	}

	itemRequest.Seller = sellerId

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		// http_utils.RespondJson(w, err.Status, err)
		http_utils.RespondError(w, createErr)

		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)

}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	item, err := services.ItemsService.Get(itemId)
	if err != nil {
		http_utils.RespondJson(w, http.StatusNotFound, err)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, item)
}
