package bridge

import "fmt"

type HP struct {
}

func (p *HP) Print() error {
	fmt.Println("HP printer printing...")
	return nil
}
