package adapter

import "fmt"

type Fetch struct {
}

func (m *Fetch) Get(url string) (interface{}, error) {
	fmt.Printf("Http Get with Fetch: %s\n", url)
	return nil, nil
}
