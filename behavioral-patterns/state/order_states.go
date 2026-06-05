package state

type PendingState struct {
	order *Order
}

func (s *PendingState) Next(order *Order) string {
	order.setState(&BrewingState{order: order})
	return "Order accepted, brewing started"
}

func (s *PendingState) Status() string {
	return "Pending"
}

type BrewingState struct {
	order *Order
}

func (s *BrewingState) Next(order *Order) string {
	order.setState(&ServedState{order: order})
	return "Drink is ready"
}

func (s *BrewingState) Status() string {
	return "Brewing"
}

type ServedState struct {
	order *Order
}

func (s *ServedState) Next(order *Order) string {
	return "Order already served"
}

func (s *ServedState) Status() string {
	return "Served"
}
