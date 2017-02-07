package persistence

import (
	"errors"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"home24/core/provider"
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

func (elasticProvider ElasticProvider) Get(request provider.Request) (interface{}, error) {
	baseQuery := elasticProvider.elasticClient.Search().Index("home24-test")

	boolQuery := elastic.NewBoolQuery()

	for _, predicate := range request.Predicates {
		termQuery := elastic.NewTermQuery(predicate.Name, predicate.Value)

		if predicate.Weight != nil {
			termQuery.Boost(predicate.Weight.(float64))
		}

		if predicate.ComparisonOperator == "EqualTo" {
			boolQuery = boolQuery.Should(termQuery)
		} else if predicate.ComparisonOperator == "NotEqualTo" {
			boolQuery = boolQuery.MustNot(termQuery)
		}
	}

	baseQuery = baseQuery.Query(boolQuery)

	searchResult, searchErr := baseQuery.Sort("_score", false).From(request.Index).Size(request.NumItems).Do(context.Background())

	if searchErr != nil {
		return nil, errors.New("There was an error while performing the search")
	}

	if searchResult.TotalHits() == 0 {
		return nil, errors.New("No matching results for given filters")
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
