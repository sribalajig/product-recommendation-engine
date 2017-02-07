package product

import (
	"home24/core/models"
	"home24/core/provider"
	"home24/recommendation/engine"
	"reflect"
)

func Build(product models.Product) (engine.Request, error) {
	var weightedAttributes []engine.WeightedAttribute

	for _, attr := range product.Attributes {
		wt, wtErr := GetWeight(attr.Name)

		if wtErr != nil {
			return engine.Request{}, wtErr
		}

		weightedAttributes = append(weightedAttributes, engine.WeightedAttribute{
			Attribute: attr,
			Weight:    engine.Weight{Value: wt},
		})
	}

	request := engine.Request{
		DesiredRecommendationType: reflect.TypeOf(models.Product{}),
		WeightedAttributes:        weightedAttributes,
		PaginationOptions: engine.Pagination{
			NumberOfItems: 10,
			StartIndex:    0,
		},
		ExclusionList: []provider.ExclusionItem{
			provider.ExclusionItem{Key: "name", Value: product.Name},
		},
	}

	return request, nil
}
