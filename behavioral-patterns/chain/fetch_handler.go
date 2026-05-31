package chain

import "fmt"

type FetchHandler struct {
	BaseHandler
}

func NewFetchHandler() *FetchHandler {
	return &FetchHandler{}
}

func (h *FetchHandler) Handle(url string) (interface{}, error) {
	return fmt.Sprintf("response from %s", url), nil
}
