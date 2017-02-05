package persistence

import (
	"encoding/json"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"home24/core"
)

/*ProductResultTransformer trnasforms a raw json message to a Product struct*/
type ProductResultTransformer struct {
}

/*Transform converts an elastic search hit into a Product struct*/
func (productResultTransformer ProductResultTransformer) Transform(searchHit *elastic.SearchHit) (interface{}, error) {
	productWithAttributes := make(map[string]string)

	err := json.Unmarshal(*searchHit.Source, &productWithAttributes)

	if err != nil {
		return nil, fmt.Errorf("Could not deserialize : %s", err.Error())
	}

	var product core.Product

	for key := range productWithAttributes {
		if key == "name" {
			product.Name = productWithAttributes[key]

			continue
		}

		product.Attributes = append(product.Attributes, core.Attribute{
			Name:  key,
			Value: productWithAttributes[key],
		})
	}

	return product, nil
}
