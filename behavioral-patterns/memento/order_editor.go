package memento

type OrderEditor struct {
	drinkName string
	size      string
}

func NewOrderEditor(drinkName string, size string) *OrderEditor {
	return &OrderEditor{drinkName: drinkName, size: size}
}

func (e *OrderEditor) SetDrinkName(drinkName string) {
	e.drinkName = drinkName
}

func (e *OrderEditor) SetSize(size string) {
	e.size = size
}

func (e *OrderEditor) Save() Memento {
	return NewMemento(e.drinkName, e.size)
}

func (e *OrderEditor) Restore(memento Memento) {
	e.drinkName = memento.DrinkName()
	e.size = memento.Size()
}

func (e *OrderEditor) Snapshot() string {
	return e.size + " " + e.drinkName
}
