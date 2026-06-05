package strategy

type HappyHourPricing struct{}

func (HappyHourPricing) Calculate(basePrice float64) float64 {
	return basePrice * 0.8
}

func (HappyHourPricing) Name() string {
	return "Happy Hour"
}
