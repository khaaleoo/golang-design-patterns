package observer

import "fmt"

type Cashier struct {
	Name string
}

func (c *Cashier) Update(orderID string, drinkName string) {
	fmt.Printf("Cashier %s: received order %s (%s)\n", c.Name, orderID, drinkName)
}
