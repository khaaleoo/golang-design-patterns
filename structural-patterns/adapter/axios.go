package adapter

import "fmt"

type Axios struct {
}

func (m *Axios) Get(url string) (interface{}, error) {
	fmt.Printf("Http Get with Axios: %s\n", url)
	return nil, nil
}
