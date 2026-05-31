package strategy

type PricingStrategy interface {
	Calculate(basePrice float64) float64
	Name() string
}
