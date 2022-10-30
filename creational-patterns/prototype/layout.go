package prototype

import "fmt"

type Layout struct {
	Header *string
	Body   string
	Footer *string
}

const (
	BLANK_LAYOUT = "BLANK"
	MAIN_LAYOUT  = "MAIN"
)

func (s *Layout) SetBody(body string) {
	s.Body = body
}

func (s *Layout) GetInfo() string {
	header, footer := "empty", "empty"
	if s.Header != nil {
		header = *s.Header
	}
	if s.Footer != nil {
		footer = *s.Footer
	}
	return fmt.Sprintf("Layout: Header: %s, Body: %s, Footer: %s", header, s.Body, footer)
}
