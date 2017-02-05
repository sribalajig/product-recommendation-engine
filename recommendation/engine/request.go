package engine

import (
	"reflect"
)

type Request struct {
	RecommendationType reflect.Type
	WeightedAttributes []WeightedAttribute
	PaginationOptions  Pagination
}
