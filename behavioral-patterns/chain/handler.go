package chain

type Handler interface {
	SetNext(handler Handler) Handler
	Handle(url string) (interface{}, error)
}

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *BaseHandler) nextHandle(url string) (interface{}, error) {
	if h.next == nil {
		return nil, nil
	}
	return h.next.Handle(url)
}
