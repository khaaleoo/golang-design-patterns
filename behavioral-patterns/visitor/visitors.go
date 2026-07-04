package visitor

import "fmt"

type PriceVisitor struct{}

func (v PriceVisitor) VisitDrink(drink Drink) string {
	return fmt.Sprintf("%s costs %.2f", drink.Name, drink.Price)
}

func (v PriceVisitor) VisitPastry(pastry Pastry) string {
	return fmt.Sprintf("%s costs %.2f", pastry.Name, pastry.Price)
}

type LabelVisitor struct{}

func (v LabelVisitor) VisitDrink(drink Drink) string {
	return "Drink: " + drink.Name
}

func (v LabelVisitor) VisitPastry(pastry Pastry) string {
	return "Pastry: " + pastry.Name
}
