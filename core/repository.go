package core

import (
	"fmt"
	"reflect"
)

var Provider DataProvider

/*Get is a generic repository method to get an item of any given type*/
func Get(filters []Filter, typ reflect.Type) (interface{}, error) {
	result, err := Provider.Get(filters, typ)

	if err != nil {
		return nil, fmt.Errorf("There was an error while retrieving data : %s", err)
	}

	return result, nil
}
