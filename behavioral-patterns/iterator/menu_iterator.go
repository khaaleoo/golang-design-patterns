package iterator

type MenuIterator struct {
	menu  *Menu
	index int
}

func (i *MenuIterator) HasNext() bool {
	return i.index < len(i.menu.items)
}

func (i *MenuIterator) Next() MenuItem {
	item := i.menu.items[i.index]
	i.index++
	return item
}
