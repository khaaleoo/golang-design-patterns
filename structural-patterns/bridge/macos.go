package bridge

import "fmt"

type MacOS struct {
	printer Printer
}

func (m *MacOS) Print() error {
	fmt.Println("MacOS printing...")
	return m.printer.Print()
}

func (m *MacOS) SetPrinter(printer Printer) {
	m.printer = printer
}
