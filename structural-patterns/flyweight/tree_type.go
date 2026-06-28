package flyweight

import "fmt"

// TreeType holds intrinsic (shared) state — species, color, texture.
// Many Tree instances can share a single TreeType.
type TreeType struct {
	Name    string
	Color   string
	Texture string
}

func (t *TreeType) Render(x, y int) {
	fmt.Printf("Rendering tree [%s | color:%s | texture:%s] at (%d, %d)\n",
		t.Name, t.Color, t.Texture, x, y)
}
