package prototype

type BlankLayout struct {
	Layout
}

var BlankLayoutIns *Layout = &Layout{
	Header: nil,
	Body:   "Blank Body",
	Footer: nil,
}
