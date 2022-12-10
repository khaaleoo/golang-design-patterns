package adapter

type FetchAdapter struct {
	Instance *Fetch
}

func (m *FetchAdapter) Get(url string) (interface{}, error) {
	return m.Instance.Get(url)
}
