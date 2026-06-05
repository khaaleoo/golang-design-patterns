package chain

import (
	"errors"
	"sync/atomic"
)

type RateLimitHandler struct {
	BaseHandler
	limit   int64
	current int64
}

func NewRateLimitHandler(limit int64) *RateLimitHandler {
	return &RateLimitHandler{limit: limit}
}

func (h *RateLimitHandler) Handle(url string) (interface{}, error) {
	if atomic.AddInt64(&h.current, 1) > h.limit {
		atomic.AddInt64(&h.current, -1)
		return nil, errors.New("rate limit exceeded")
	}
	result, err := h.nextHandle(url)
	if err != nil {
		atomic.AddInt64(&h.current, -1)
	}
	return result, err
}
