package templatemethod

type Coffee struct{}

func (c Coffee) Brew() string {
	return "Dripped coffee through filter"
}

func (c Coffee) PourInCup() string {
	return "Poured coffee into cup"
}

func (c Coffee) AddCondiments() string {
	return "Added sugar and milk"
}
