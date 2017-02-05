package core

import (
	"reflect"
)

/* Data provider defines an interface for any peristence provider (such as mongo, elastic search etc)*/
type DataProvider interface {
	Get(predicates []Predicate, typ reflect.Type) (interface{}, error)
}
