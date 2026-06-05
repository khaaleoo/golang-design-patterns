package observer

type Observer interface {
	Update(orderID string, drinkName string)
}
