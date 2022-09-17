package factory_method

type CoffeeDrink struct {
	name string
}

func (me *CoffeeDrink) GetName() string {
	return me.name
}
