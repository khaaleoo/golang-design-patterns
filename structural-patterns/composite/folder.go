package composite

import (
	"fmt"
	"log"
	"strings"
)

type Folder struct {
	name       string
	components []Component
}

func (m *Folder) GetName() string {
	return m.name
}

func (m *Folder) SetName(name string) {
	m.name = name
}

func (m *Folder) Print(args ...interface{}) {
	fmt.Println(m.name)
	nested := 0
	if len(args) > 0 {
		var ok bool
		nested, ok = args[0].(int)
		if !ok {
			log.Fatal("first argument must be a number")
		}
	}
	for _, s := range m.components {
		fmt.Printf("%s%s%s", strings.Repeat("  ", nested), strings.Repeat(" ", nested), "|--")
		s.Print(nested + 1)
	}
}

func (m *Folder) Add(c ...Component) {
	m.components = append(m.components, c...)
}
