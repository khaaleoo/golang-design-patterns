package bridge

import "fmt"

type Window struct {
	printer Printer
}

func (w *Window) Print() error {
	fmt.Println("Window printing...")
	return w.printer.Print()
}

func (w *Window) SetPrinter(printer Printer) {
	w.printer = printer
}
