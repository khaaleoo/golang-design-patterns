package main

import (
	"fmt"

	adapter "github.com/structural-patterns/adapter"
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

	/*
		Example Adapter
	*/
	fmt.Println("*** Example Adapter ***")
	client := &adapter.Client{}

	fetchAdapter := &adapter.FetchAdapter{
		Instance: &adapter.Fetch{},
	}
	client.Get(fetchAdapter, "https://www.google.com")

	axiosAdapter := &adapter.AxiosAdapter{
		Instance: &adapter.Axios{},
	}
	client.Get(axiosAdapter, "https://www.bornhup.com")

	fmt.Print("*** End of Adapter ***\n\n\n")

}
