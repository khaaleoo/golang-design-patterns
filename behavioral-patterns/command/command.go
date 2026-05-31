package command

type Command interface {
	Execute() string
	Undo() string
}
