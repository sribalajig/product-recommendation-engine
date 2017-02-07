package product

import (
	"home24/core/models"
	"home24/core/provider"
	"home24/core/repository"
	"home24/recommendation/engine"
	"reflect"
)

func GetRecommendations(productCode string) ([]models.RecommendationResult, error) {
	product, err := GetProduct(productCode)

	if err != nil {
		return nil, err
	}

	request, reqErr := Build(product)

	if reqErr != nil {
		return nil, reqErr
	}

	recommendations, recErr := engine.Get(request)

	if recErr != nil {
		return nil, recErr
	}

	var products []models.RecommendationResult

	for _, reco := range recommendations.([]interface{}) {
		scoredProduct := reco.(models.RecommendationResult)

		products = append(products, scoredProduct)
	}

	return products, nil
}

func GetProduct(productCode string) (models.Product, error) {
	predicates := []provider.Predicate{
		provider.Predicate{
			Name:               "name",
			Value:              productCode,
			ComparisonOperator: "EqualTo",
		},
	}

	result, err := repository.GetOne(predicates, reflect.TypeOf(models.Product{}))

	if err != nil {
		return models.Product{}, err
	}

	scoredResult := result.(models.RecommendationResult)

	return scoredResult.Item.(models.Product), nil
}
