package state

type OrderState interface {
	Next(order *Order) string
	Status() string
}
