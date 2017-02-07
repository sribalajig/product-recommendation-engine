package engine

import (
	"home24/core/provider"
	"reflect"
)

type Request struct {
	DesiredRecommendationType reflect.Type
	WeightedAttributes        []WeightedAttribute
	PaginationOptions         Pagination
	ExclusionList             []provider.ExclusionItem
}
