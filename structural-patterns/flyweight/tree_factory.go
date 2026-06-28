package flyweight

import "fmt"

var treeTypes = map[string]*TreeType{}

// GetTreeType returns a cached TreeType or creates a new one if not found.
// This is the core of the Flyweight pattern — reuse shared instances.
func GetTreeType(name, color, texture string) *TreeType {
	key := fmt.Sprintf("%s-%s-%s", name, color, texture)
	if tt, ok := treeTypes[key]; ok {
		return tt
	}
	tt := &TreeType{Name: name, Color: color, Texture: texture}
	treeTypes[key] = tt
	fmt.Printf("Created new TreeType: %s\n", key)
	return tt
}

func TreeTypeCount() int {
	return len(treeTypes)
}
