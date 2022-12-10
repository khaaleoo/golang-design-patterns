package adapter

type AxiosAdapter struct {
	Instance *Axios
}

func (m *AxiosAdapter) Get(url string) (interface{}, error) {
	return m.Instance.Get(url)
}
