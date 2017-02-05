package engine

import (
	"reflect"
)

type Request struct {
	DesiredRecommendationType reflect.Type
	WeightedAttributes        []WeightedAttribute
	PaginationOptions         Pagination
}
