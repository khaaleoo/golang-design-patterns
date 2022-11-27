package main

import (
	"fmt"

	composite "github.com/structural-patterns/composite"
)

func main() {
	/*
		Example Composite
	*/
	fmt.Println("*** Example Composite ***")
	root := &composite.Folder{}
	root.SetName("Root")

	folderA := &composite.Folder{}
	folderA.SetName("FolderA")

	fileA := &composite.File{}
	fileA.SetName("FileA")

	fileB := &composite.File{}
	fileB.SetName("FileB")

	folderX := &composite.Folder{}
	folderX.SetName("FolderX")

	fileY := &composite.File{}
	fileY.SetName("FileY")
	folderZ := &composite.Folder{}
	folderZ.SetName("FolderZ")

	fileW := &composite.File{}
	fileW.SetName("FileW")

	folderZ.Add(fileW)
	folderX.Add(fileY, folderZ)
	folderA.Add(fileA, fileB, folderX)

	folderB := &composite.Folder{}
	folderB.SetName("FolderB")

	fileC := &composite.File{}
	fileC.SetName("FileC")

	folderB.Add(fileC)

	root.Add(folderA, folderB)

	root.Print()

	fmt.Print("*** End of Composite ***\n\n\n")
}
