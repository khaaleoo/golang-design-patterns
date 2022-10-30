package prototype

import "fmt"

type MainLayout struct {
	Layout
}

var MainLayoutIns *Layout = &Layout{
	Header: getHeader(),
	Body:   "Main Body",
	Footer: getFooter(),
}

func getHeader() *string {
	fmt.Println("Getting header data, it took 1 second...")
	header := "Main Header"
	return &header
}

func getFooter() *string {
	footer := "Main Footer"
	return &footer
}
