package mediator

type Mediator interface {
	Notify(sender string, event string) string
}
