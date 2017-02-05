package main

import (
	"encoding/json"
	"fmt"
	"home24/core"
	"home24/persistence"
	"home24/recommendation/engine"
	"home24/recommendation/product"
)

func main() {
	core.Provider = persistence.NewElasticProvider()
	engine.DataProvider = persistence.NewElasticProvider()

	recommendedProducts, err := product.GetRecommendations("sku-12")

	if err != nil {
		fmt.Println(err)

		return
	}

	result, _ := json.MarshalIndent(recommendedProducts, "", "\t")

	fmt.Println(string(result))
}
