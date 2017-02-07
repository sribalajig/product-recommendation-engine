package engine

import (
	"fmt"
	"home24/core/models"
	"home24/core/provider"
	"reflect"
)

var DataProvider provider.DataProvider

/*Get will return recommendations based on the type requested for*/
func Get(request Request) (interface{}, error) {
	var predicates []provider.Predicate

	for _, weightedAttribute := range request.WeightedAttributes {
		predicates = append(predicates, provider.Predicate{
			Name:   weightedAttribute.Attribute.Name,
			Value:  weightedAttribute.Attribute.Value,
			Weight: weightedAttribute.Weight.Value,
		})
	}

	result, err := DataProvider.Get(predicates, 0, 11, reflect.TypeOf(models.Product{}))

	if err != nil {
		return nil, fmt.Errorf("There was an error retrieving data : %s", err.Error())
	}

	return result, nil
}
