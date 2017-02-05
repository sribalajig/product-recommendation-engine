package persistence

import (
	"errors"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"home24/core"
	"reflect"
)

type ElasticProvider struct {
	elasticClient *elastic.Client
}

/*NewElasticProvider is a factory methos*/
func NewElasticProvider() ElasticProvider {
	client, err := elastic.NewClient()

	if err != nil {
		panic(err)
	}

	return ElasticProvider{
		elasticClient: client,
	}
}

/*Get is a specific implementation for elastic search*/
func (elasticProvider ElasticProvider) Get(filters []core.Filter, typ reflect.Type) (interface{}, error) {
	baseQuery := elasticProvider.elasticClient.Search().Index("home24-test")

	for _, filter := range filters {
		termQuery := elastic.NewTermQuery(filter.Name, filter.Value)

		baseQuery = baseQuery.Query(termQuery)
	}

	searchResult, searchErr := baseQuery.Do(context.Background())

	if searchErr != nil {
		return nil, errors.New("There was an error while performing the search")
	}

	if searchResult.TotalHits() == 0 {
		return nil, errors.New("No matching results for given filters")
	}

	transformer, tranErr := Create(typ)

	if tranErr != nil {
		return nil, tranErr
	}

	return transformer.Transform(searchResult.Hits.Hits[0].Source)
}
