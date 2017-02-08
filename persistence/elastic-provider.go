package persistence

import (
	"errors"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"home24/core/provider"
)

type ElasticProvider struct {
	elasticClient *elastic.Client
	queryBuilder  QueryBuilder
}

/*NewElasticProvider is a factory method*/
func NewElasticProvider() ElasticProvider {
	client, err := elastic.NewClient()

	if err != nil {
		panic(err)
	}

	return ElasticProvider{
		elasticClient: client,
		queryBuilder:  QueryBuilder{},
	}
}

func (elasticProvider ElasticProvider) Get(request provider.Request) (interface{}, error) {
	query := elasticProvider.queryBuilder.Build(elasticProvider.elasticClient, request)

	searchResult, searchErr := query.
		Sort("_score", false).
		From(request.Index).
		Size(request.NumItems).
		Do(context.Background())

	if searchErr != nil {
		return nil, errors.New("There was an error while performing the search")
	}

	transformer, tranErr := Create(request.Typ)

	if tranErr != nil {
		return nil, tranErr
	}

	var items []interface{}

	for _, hit := range searchResult.Hits.Hits {
		result, transformErr := transformer.Transform(hit)

		if transformErr != nil {
			return nil, transformErr
		}

		items = append(items, result)
	}

	return items, nil
}
