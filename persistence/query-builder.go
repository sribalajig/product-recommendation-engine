package persistence

import (
	"gopkg.in/olivere/elastic.v5"
	"home24/core/provider"
)

type QueryBuilder struct {
}

func (queryBuilder QueryBuilder) Build(elasticClient *elastic.Client, request provider.Request) *elastic.SearchService {
	baseQuery := elasticClient.Search().Index("home24")

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

	return baseQuery
}
