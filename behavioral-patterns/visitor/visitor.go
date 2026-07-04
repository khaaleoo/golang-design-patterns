package visitor

type Visitor interface {
	VisitDrink(drink Drink) string
	VisitPastry(pastry Pastry) string
}

type MenuItem interface {
	Accept(visitor Visitor) string
}
