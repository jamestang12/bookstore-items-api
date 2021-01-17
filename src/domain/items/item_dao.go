package items

import (
	// "github.com/elastic/go-elasticsearch"
	"encoding/json"
	"errors"
	"fmt"

	"../../../../bookstore_utils_go/rest_errors"
	"../../clients/elasticsearch"
	"../queries"
)

const (
	indexItem = "items"
	typeItem  = "_doc"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItem, typeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(indexItem, typeItem, i.Id)
	if err != nil {

		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id), errors.New("database error"))
	}
	if !result.Found {
		return rest_errors.NewBadNotFoundError(fmt.Sprintf("no item found with id %s", i.Id))
	}
	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError("error when tyring to parse database response", errors.New("database error"))
	}
	fmt.Println(string(bytes))
	if err := json.Unmarshal(bytes, i); err != nil {
		return rest_errors.NewInternalServerError("error when tyring to parse database response", errors.New("database error"))
	}
	i.Id = itemId
	return nil
}

func (i *Item) Search(query queries.EsQuery) ([]Item, rest_errors.RestErr) {
	result, err := elasticsearch.Client.Search(indexItem, query.Build())
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to search documents", errors.New("database error"))
	}
	//fmt.Println(result)

	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_errors.NewInternalServerError("error when trying to parse response", errors.New("database error"))
		}
		item.Id = hit.Id
		items[index] = item
	}

	if len(items) == 0 {
		return nil, rest_errors.NewBadNotFoundError("no items found matching given criteria")
	}

	return items, nil
}
