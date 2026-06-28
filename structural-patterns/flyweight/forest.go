package flyweight

// Forest manages a large collection of Tree objects.
// Despite having thousands of trees, only a handful of TreeType objects exist.
type Forest struct {
	trees []*Tree
}

func (f *Forest) PlantTree(x, y int, name, color, texture string) {
	tt := GetTreeType(name, color, texture)
	f.trees = append(f.trees, &Tree{X: x, Y: y, TreeType: tt})
}

func (f *Forest) Render() {
	for _, t := range f.trees {
		t.Render()
	}
}

func (f *Forest) Count() int {
	return len(f.trees)
}
