package mediator

type CoffeeShopMediator struct {
	cashier *Cashier
	barista *Barista
}

func NewCoffeeShopMediator() *CoffeeShopMediator {
	return &CoffeeShopMediator{}
}

func (m *CoffeeShopMediator) SetCashier(cashier *Cashier) {
	m.cashier = cashier
}

func (m *CoffeeShopMediator) SetBarista(barista *Barista) {
	m.barista = barista
}

func (m *CoffeeShopMediator) Notify(sender string, event string) string {
	switch event {
	case "order-received":
		return m.barista.PrepareDrink(sender)
	case "drink-ready":
		return m.cashier.CallCustomer(sender)
	default:
		return "Unknown event"
	}
}
