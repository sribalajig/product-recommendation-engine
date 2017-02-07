package persistence

import (
	"errors"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"home24/core/provider"
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

func (elasticProvider ElasticProvider) Get(
	predicates []provider.Predicate,
	index int,
	numItems int,
	typ reflect.Type) (interface{}, error) {
	baseQuery := elasticProvider.elasticClient.Search().Index("home24-test")

	boolQuery := elastic.NewBoolQuery()

	for _, predicate := range predicates {
		termQuery := elastic.NewTermQuery(predicate.Name, predicate.Value)

		if predicate.Weight != nil {
			termQuery.Boost(predicate.Weight.(float64))
		}

		boolQuery = boolQuery.Should(termQuery)
	}

	baseQuery = baseQuery.Query(boolQuery)

	searchResult, searchErr := baseQuery.Sort("_score", false).From(index).Size(numItems).Do(context.Background())

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
