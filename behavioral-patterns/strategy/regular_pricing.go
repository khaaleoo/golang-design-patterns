package strategy

type RegularPricing struct{}

func (RegularPricing) Calculate(basePrice float64) float64 {
	return basePrice
}

func (RegularPricing) Name() string {
	return "Regular"
}
