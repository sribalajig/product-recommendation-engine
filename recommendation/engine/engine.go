package engine

import (
	"fmt"
	"home24/core"
	"reflect"
)

var DataProvider core.DataProvider

/*Get will return recommendations based on the type requested for*/
func Get(request Request) (interface{}, error) {
	var predicates []core.Predicate

	for _, weightedAttribute := range request.WeightedAttributes {
		predicates = append(predicates, core.Predicate{
			Name:   weightedAttribute.Attribute.Name,
			Value:  weightedAttribute.Attribute.Value,
			Weight: weightedAttribute.Weight.Value,
		})
	}

	result, err := DataProvider.Get(predicates, 0, 11, reflect.TypeOf(core.Product{}))

	if err != nil {
		return nil, fmt.Errorf("There was an error retrieving data : %s", err.Error())
	}

	return result, nil
}
