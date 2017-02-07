package provider

import (
	"reflect"
)

type Request struct {
	Predicates []Predicate
	Index      int
	NumItems   int
	Typ        reflect.Type
}
