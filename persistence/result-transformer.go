package persistence

import (
	"gopkg.in/olivere/elastic.v5"
)

/*ResultTransformer is an interface for specifying how result transformation happens*/
type ResultTransformer interface {
	Transform(searchHit *elastic.SearchHit) (interface{}, error)
}
