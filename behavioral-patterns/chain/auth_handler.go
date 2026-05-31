package chain

import "errors"

type AuthHandler struct {
	BaseHandler
	token string
}

func NewAuthHandler(token string) *AuthHandler {
	return &AuthHandler{token: token}
}

func (h *AuthHandler) Handle(url string) (interface{}, error) {
	if h.token == "" {
		return nil, errors.New("missing auth token")
	}
	return h.nextHandle(url)
}
