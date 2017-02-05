package main

import (
	"fmt"
	"home24/core"
	"home24/persistence"
	"home24/recommendation/engine"
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

	result, err := core.Get(filters, reflect.TypeOf(core.Product{}))

	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	_ = result.(core.Product)

	engine.Get(engine.Request{})
}
