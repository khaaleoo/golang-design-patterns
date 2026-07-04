package mediator

type Cashier struct {
	Name     string
	mediator Mediator
}

func NewCashier(name string, mediator Mediator) *Cashier {
	return &Cashier{Name: name, mediator: mediator}
}

func (c *Cashier) TakeOrder(drinkName string) string {
	return c.mediator.Notify(drinkName, "order-received")
}

func (c *Cashier) CallCustomer(drinkName string) string {
	return c.Name + " called customer for " + drinkName
}

type Barista struct {
	Name     string
	mediator Mediator
}

func NewBarista(name string, mediator Mediator) *Barista {
	return &Barista{Name: name, mediator: mediator}
}

func (b *Barista) PrepareDrink(drinkName string) string {
	return b.Name + " prepared " + drinkName
}

func (b *Barista) FinishDrink(drinkName string) string {
	return b.mediator.Notify(drinkName, "drink-ready")
}
