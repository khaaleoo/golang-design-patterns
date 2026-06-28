package flyweight

// Tree holds extrinsic (context-specific) state — position on the map.
// It delegates rendering to the shared TreeType flyweight.
type Tree struct {
	X        int
	Y        int
	TreeType *TreeType
}

func (t *Tree) Render() {
	t.TreeType.Render(t.X, t.Y)
}
