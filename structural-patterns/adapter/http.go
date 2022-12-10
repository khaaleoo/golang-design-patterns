package adapter

type Http interface {
	Get(url string) (interface{}, error)
}
