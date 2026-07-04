package interpreter

type Expression interface {
	Interpret() float64
}
