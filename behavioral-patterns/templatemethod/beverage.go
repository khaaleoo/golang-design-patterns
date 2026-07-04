package templatemethod

type Beverage interface {
	Brew() string
	PourInCup() string
	AddCondiments() string
}

func Prepare(beverage Beverage) []string {
	return []string{
		"Boiled water",
		beverage.Brew(),
		beverage.PourInCup(),
		beverage.AddCondiments(),
	}
}
