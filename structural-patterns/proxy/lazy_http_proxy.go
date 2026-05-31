package proxy

import (
	"fmt"
	"sync"

	adapter "github.com/structural-patterns/adapter"
)

type LazyHttpProxy struct {
	createClient func() adapter.Http
	client       adapter.Http
	once         sync.Once
}

func NewLazyHttpProxy(createClient func() adapter.Http) *LazyHttpProxy {
	return &LazyHttpProxy{createClient: createClient}
}

func (p *LazyHttpProxy) Get(url string) (interface{}, error) {
	p.once.Do(func() {
		fmt.Println("Lazy proxy: creating HTTP client")
		p.client = p.createClient()
	})
	return p.client.Get(url)
}
