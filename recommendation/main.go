package main

import (
	"fmt"
	"home24/core"
	"home24/persistence"
	"reflect"
)

func main() {
	filter := core.Filter{
		Name:  "name",
		Value: "sku-35",
	}

	filters := []core.Filter{
		filter,
	}

	core.Provider = persistence.NewElasticProvider()

	results, err := core.Get(filters, reflect.TypeOf(core.Product{}))

	if err == nil {
		fmt.Println(results)
	} else {
		fmt.Println(err)
	}
}
