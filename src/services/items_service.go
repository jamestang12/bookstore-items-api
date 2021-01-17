package services

import (
	"../../../bookstore_utils_go/rest_errors"
	"../domain/items"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(itemRequest items.Item) (*items.Item, rest_errors.RestErr) {
	//return nil, rest_errors.NewBadRequestError("testing")

	if err := itemRequest.Save(); err != nil {
		return nil, err
	}

	return &itemRequest, nil
}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	// return nil, rest_errors.NewBadRequestError("testing2")
	item := items.Item{Id: id}
	if err := item.Get(); err != nil {
		return nil, err
	}

	return &item, nil
	//return nil, rest_errors.NewRestErrpr("Not yet implement", http.StatusNotImplemented, "not_implemented", nil)
}
