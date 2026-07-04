package visitor

type Drink struct {
	Name  string
	Price float64
}

func NewDrink(name string, price float64) Drink {
	return Drink{Name: name, Price: price}
}

func (d Drink) Accept(visitor Visitor) string {
	return visitor.VisitDrink(d)
}

type Pastry struct {
	Name  string
	Price float64
}

func NewPastry(name string, price float64) Pastry {
	return Pastry{Name: name, Price: price}
}

func (p Pastry) Accept(visitor Visitor) string {
	return visitor.VisitPastry(p)
}
