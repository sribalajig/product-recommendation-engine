package engine

import (
	"home24/core/models"
)

type WeightedAttribute struct {
	Attribute models.Attribute
	Weight    Weight
}
