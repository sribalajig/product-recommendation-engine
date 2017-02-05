package persistence

import (
	"encoding/json"
)

/*ResultTransformer is an interface for specifying how result transformation happens*/
type ResultTransformer interface {
	Transform(raw *json.RawMessage) (interface{}, error)
}
