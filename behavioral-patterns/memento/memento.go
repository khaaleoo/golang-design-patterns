package memento

type Memento struct {
	drinkName string
	size      string
}

func NewMemento(drinkName string, size string) Memento {
	return Memento{drinkName: drinkName, size: size}
}

func (m Memento) DrinkName() string {
	return m.drinkName
}

func (m Memento) Size() string {
	return m.size
}
