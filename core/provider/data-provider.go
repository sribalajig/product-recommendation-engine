package provider

/* Data provider defines an interface for any peristence provider (such as mongo, elastic search etc)*/
type DataProvider interface {
	Get(request Request) (interface{}, error)
}
