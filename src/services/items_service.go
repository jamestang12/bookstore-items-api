package services

import (
	"../../../bookstore_utils_go/rest_errors"
	"../domain/items"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, *rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(itemRequest items.Item) (*items.Item, *rest_errors.RestErr) {
	//return nil, rest_errors.NewBadRequestError("testing")

	if err := itemRequest.Save(); err != nil {
		return nil, err
	}

	return &itemRequest, nil
}

func (s *itemsService) Get(string) (*items.Item, *rest_errors.RestErr) {
	return nil, rest_errors.NewBadRequestError("testing2")
}
