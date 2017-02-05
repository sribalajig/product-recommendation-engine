package product

import (
	"home24/core"
	"home24/recommendation/engine"
	"reflect"
)

func GetRecommendations(productCode string) ([]core.Product, error) {
	product, err := GetProduct(productCode)

	if err != nil {
		return nil, err
	}

	request, reqErr := Build(product.Attributes)

	if reqErr != nil {
		return nil, reqErr
	}

	recommendations, recErr := engine.Get(request)

	if recErr != nil {
		return nil, recErr
	}

	var products []core.Product

	for _, reco := range recommendations.([]interface{}) {
		products = append(products, reco.(core.Product))
	}

	return products, nil
}

func GetProduct(productCode string) (core.Product, error) {
	predicates := []core.Predicate{
		core.Predicate{
			Name:  "name",
			Value: productCode,
		},
	}

	result, err := core.GetOne(predicates, reflect.TypeOf(core.Product{}))

	if err != nil {
		return core.Product{}, err
	}

	return result.(core.Product), nil
}
