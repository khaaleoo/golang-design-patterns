package interpreter

type NumberExpression struct {
	value float64
}

func NewNumberExpression(value float64) NumberExpression {
	return NumberExpression{value: value}
}

func (e NumberExpression) Interpret() float64 {
	return e.value
}

type AddExpression struct {
	left  Expression
	right Expression
}

func NewAddExpression(left Expression, right Expression) AddExpression {
	return AddExpression{left: left, right: right}
}

func (e AddExpression) Interpret() float64 {
	return e.left.Interpret() + e.right.Interpret()
}

type SubtractExpression struct {
	left  Expression
	right Expression
}

func NewSubtractExpression(left Expression, right Expression) SubtractExpression {
	return SubtractExpression{left: left, right: right}
}

func (e SubtractExpression) Interpret() float64 {
	return e.left.Interpret() - e.right.Interpret()
}
