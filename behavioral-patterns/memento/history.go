package memento

type History struct {
	states []Memento
}

func NewHistory() *History {
	return &History{}
}

func (h *History) Backup(editor *OrderEditor) {
	h.states = append(h.states, editor.Save())
}

func (h *History) Undo(editor *OrderEditor) string {
	if len(h.states) == 0 {
		return "No saved order state"
	}

	lastIndex := len(h.states) - 1
	editor.Restore(h.states[lastIndex])
	h.states = h.states[:lastIndex]

	return "Restored order to " + editor.Snapshot()
}
