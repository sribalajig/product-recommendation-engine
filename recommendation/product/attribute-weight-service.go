package product

import (
	"errors"
)

var weights map[string]float64

func init() {
	weights = make(map[string]float64)

	weights["att-a"] = 10
	weights["att-b"] = 9
	weights["att-c"] = 8
	weights["att-d"] = 7
	weights["att-e"] = 6
	weights["att-f"] = 5
	weights["att-g"] = 4
	weights["att-h"] = 3
	weights["att-i"] = 2
	weights["att-j"] = 1
}

func GetWeight(attributeName string) (float64, error) {
	if val, ok := weights[attributeName]; ok {
		return val, nil
	}

	return -1, errors.New("No weight specified for attribute")
}
