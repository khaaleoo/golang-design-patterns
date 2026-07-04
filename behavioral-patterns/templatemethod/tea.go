package templatemethod

type Tea struct{}

func (t Tea) Brew() string {
	return "Steeped tea leaves"
}

func (t Tea) PourInCup() string {
	return "Poured tea into cup"
}

func (t Tea) AddCondiments() string {
	return "Added lemon"
}
