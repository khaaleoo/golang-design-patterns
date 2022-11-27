package composite

import (
	"fmt"
)

type File struct {
	name string
}

func (m *File) GetName() string {
	return m.name
}

func (m *File) SetName(name string) {
	m.name = name
}

func (m *File) Print(args ...interface{}) {
	fmt.Println(m.GetName())
}
