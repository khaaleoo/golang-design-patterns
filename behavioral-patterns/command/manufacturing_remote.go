package command

type ManufacturingRemote struct {
	history []Command
}

func NewManufacturingRemote() *ManufacturingRemote {
	return &ManufacturingRemote{}
}

func (r *ManufacturingRemote) Run(command Command) string {
	result := command.Execute()
	r.history = append(r.history, command)
	return result
}

func (r *ManufacturingRemote) UndoLast() string {
	if len(r.history) == 0 {
		return "No commands to undo"
	}

	last := r.history[len(r.history)-1]
	r.history = r.history[:len(r.history)-1]
	return last.Undo()
}
