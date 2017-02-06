package core

import (
	"fmt"
	"reflect"
)

var Provider DataProvider

/*Get is a generic repository method to get an item of any given type*/
func GetOne(predicates []Predicate, typ reflect.Type) (interface{}, error) {
	results, err := Provider.Get(predicates, 0, 1, typ)

	if err != nil {
		return nil, fmt.Errorf("There was an error while retrieving data : %s", err)
	}

	if results == nil || len(results.([]interface{})) == 0 {
		return nil, fmt.Errorf("The product was not found")
	}

	return results.([]interface{})[0], nil
}
