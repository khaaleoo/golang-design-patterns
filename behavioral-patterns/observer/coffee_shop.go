package observer

type CoffeeShop struct {
	observers []Observer
}

func NewCoffeeShop() *CoffeeShop {
	return &CoffeeShop{}
}

func (s *CoffeeShop) Subscribe(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *CoffeeShop) PlaceOrder(orderID string, drinkName string) {
	for _, observer := range s.observers {
		observer.Update(orderID, drinkName)
	}
}
