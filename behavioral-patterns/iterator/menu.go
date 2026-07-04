package iterator

type MenuItem struct {
	Name  string
	Price float64
}

type Menu struct {
	items []MenuItem
}

func NewMenu() *Menu {
	return &Menu{}
}

func (m *Menu) AddItem(name string, price float64) {
	m.items = append(m.items, MenuItem{Name: name, Price: price})
}

func (m *Menu) CreateIterator() Iterator {
	return &MenuIterator{menu: m}
}
