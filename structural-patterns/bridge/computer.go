package bridge

type Computer interface {
	Print() error
	SetPrinter(printer Printer)
}
