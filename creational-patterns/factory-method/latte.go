package factory_method

type Latte struct {
	CoffeeDrink
}

func newLatte() *Latte {
	return &Latte{
		CoffeeDrink: CoffeeDrink{
			name: "Latte",
		},
	}
}
