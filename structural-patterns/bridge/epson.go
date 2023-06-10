package bridge

import "fmt"

type Epson struct {
}

func (p *Epson) Print() error {
	fmt.Println("Epson printer printing...")
	return nil
}
