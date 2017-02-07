package tests

import (
	"home24/core/models"
	"home24/core/repository"
	"home24/persistence"
	"home24/recommendation/engine"
	"home24/recommendation/product"
	"sort"
	"testing"
)

type Score struct {
	Value string
	Score float64
}

type ProductScore struct {
	Name  string
	Score float64
}

type ProductScores []ProductScore

func (slice ProductScores) Len() int {
	return len(slice)
}

func (slice ProductScores) Less(i, j int) bool {
	return slice[i].Score < slice[j].Score
}

func (slice ProductScores) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func TestRecommendedProducts(t *testing.T) {
	repository.DataProvider = persistence.NewElasticProvider()
	engine.DataProvider = persistence.NewElasticProvider()

	products, _ := product.GetRecommendations("sku-445")

	if len(products) != 10 {
		t.Errorf("Expected : %d, Got : %d", 10, len(products))
		t.Fail()
	}

	scores := map[string]Score{
		"att-a": Score{Value: "att-a-10", Score: 10},
		"att-b": Score{Value: "att-b-7", Score: 9},
		"att-c": Score{Value: "att-c-6", Score: 8},
		"att-d": Score{Value: "att-d-4", Score: 7},
		"att-e": Score{Value: "att-e-14", Score: 6},
		"att-f": Score{Value: "att-f-10", Score: 5},
		"att-g": Score{Value: "att-g-13", Score: 4},
		"att-h": Score{Value: "att-h-13", Score: 3},
		"att-i": Score{Value: "att-i-7", Score: 2},
		"att-j": Score{Value: "att-j-3", Score: 1},
	}

	var productScores ProductScores

	actualRanks := make(map[string]int)

	for index, product := range products {
		score := 0.0

		for _, attribute := range product.Item.(models.Product).Attributes {
			if scores[attribute.Name].Value == attribute.Value {
				score = score + scores[attribute.Name].Score
			}
		}

		productScores = append(productScores, ProductScore{Name: product.Item.(models.Product).Name, Score: score})

		actualRanks[product.Item.(models.Product).Name] = index
	}

	sort.Reverse(productScores)

	for index, productScore := range productScores {
		if actualRanks[productScore.Name] != index {
			t.Errorf("Expected rank : %d, Actual rank : %d for Product : %s", index, actualRanks[productScore.Name], productScore.Name)
			t.Errorf("Calculated score %f", productScore.Score)
			t.Fail()
		}
	}
}
