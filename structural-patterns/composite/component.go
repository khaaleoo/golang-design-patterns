package composite

type Component interface {
	GetName() string
	Print(args ...interface{})
}
