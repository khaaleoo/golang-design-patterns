package state

type Order struct {
	ID    string
	Drink string
	state OrderState
}

func NewOrder(id string, drink string) *Order {
	order := &Order{
		ID:    id,
		Drink: drink,
	}
	order.state = &PendingState{order: order}
	return order
}

func (o *Order) Advance() string {
	return o.state.Next(o)
}

func (o *Order) Status() string {
	return o.state.Status()
}

func (o *Order) setState(state OrderState) {
	o.state = state
}
