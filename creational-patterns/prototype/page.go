package prototype

import (
	"errors"
	"fmt"
)

type Page struct {
	Route string
	Layout
}

func NewPage(route, layout string) *Page {
	layoutClone, err := cloneLayout(layout)
	if err != nil {
		panic(err)
	}
	return &Page{
		Route:  route,
		Layout: *layoutClone,
	}
}

func cloneLayout(layout string) (*Layout, error) {
	switch layout {
	case BLANK_LAYOUT:
		newLayout := *BlankLayoutIns
		return &newLayout, nil
	case MAIN_LAYOUT:
		newLayout := *MainLayoutIns
		return &newLayout, nil
	default:
		return nil, errors.New("Layout not found")
	}
}

func (s *Page) GetInfo() string {
	return fmt.Sprintf("Page route: %s, %s", s.Route, s.Layout.GetInfo())
}
