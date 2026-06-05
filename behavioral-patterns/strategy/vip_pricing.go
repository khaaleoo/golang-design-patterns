package strategy

type VIPPricing struct{}

func (VIPPricing) Calculate(basePrice float64) float64 {
	return basePrice * 0.7
}

func (VIPPricing) Name() string {
	return "VIP"
}
