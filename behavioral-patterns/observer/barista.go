package observer

import "fmt"

type Barista struct {
	Name string
}

func (b *Barista) Update(orderID string, drinkName string) {
	fmt.Printf("Barista %s: start brewing %s for order %s\n", b.Name, drinkName, orderID)
}
