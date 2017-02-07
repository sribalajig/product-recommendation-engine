package provider

type Predicate struct {
	Name               string
	Type               string
	Value              interface{}
	Weight             interface{}
	ComparisonOperator interface{}
}
