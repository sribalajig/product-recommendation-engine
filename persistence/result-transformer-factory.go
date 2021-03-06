package persistence

import (
	"fmt"
	"home24/core/models"
	"reflect"
)

/*Create is a factory method which returns a result transformer for a specific type*/
func Create(typ reflect.Type) (ResultTransformer, error) {
	if typ == reflect.TypeOf(models.Product{}) {
		return ProductResultTransformer{}, nil
	}

	return nil, fmt.Errorf(
		"Could not find a transformer for type %s", reflect.TypeOf(typ.Name()))
}
