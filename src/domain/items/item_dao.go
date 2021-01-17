package items

import (
	// "github.com/elastic/go-elasticsearch"
	"errors"

	"../../../../bookstore_utils_go/rest_errors"
	"../../clients/elasticsearch"
)

const (
	indexItem = "items"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}
