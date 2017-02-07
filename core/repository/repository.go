package repository

import (
	"fmt"
	"home24/core/provider"
	"reflect"
)

var DataProvider provider.DataProvider

/*Get is a generic repository method to get an item of any given type*/
func GetOne(predicates []provider.Predicate, typ reflect.Type) (interface{}, error) {
	request := provider.Request{
		Predicates: predicates,
		Typ:        typ,
		Index:      0,
		NumItems:   1,
	}

	results, err := DataProvider.Get(request)

	if err != nil {
		return nil, fmt.Errorf("There was an error while retrieving data : %s", err)
	}

	if results == nil || len(results.([]interface{})) == 0 {
		return nil, fmt.Errorf("The product was not found")
	}

	return results.([]interface{})[0], nil
}
