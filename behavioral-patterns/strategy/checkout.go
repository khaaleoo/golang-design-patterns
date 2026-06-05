package strategy

type Checkout struct {
	strategy PricingStrategy
}

func NewCheckout(strategy PricingStrategy) *Checkout {
	return &Checkout{strategy: strategy}
}

func (c *Checkout) SetStrategy(strategy PricingStrategy) {
	c.strategy = strategy
}

func (c *Checkout) Total(basePrice float64) float64 {
	return c.strategy.Calculate(basePrice)
}

func (c *Checkout) StrategyName() string {
	return c.strategy.Name()
}
