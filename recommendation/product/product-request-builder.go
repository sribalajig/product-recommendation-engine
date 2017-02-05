package product

import (
	"home24/core"
	"home24/recommendation/engine"
	"reflect"
)

func Build(attributes []core.Attribute) (engine.Request, error) {
	var weightedAttributes []engine.WeightedAttribute

	for _, attr := range attributes {
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
		DesiredRecommendationType: reflect.TypeOf(core.Product{}),
		WeightedAttributes:        weightedAttributes,
		PaginationOptions: engine.Pagination{
			NumberOfItems: 10,
			StartIndex:    0,
		},
	}

	return request, nil
}
