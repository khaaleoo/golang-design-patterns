package main

import (
	"fmt"

	adapter "github.com/structural-patterns/adapter"
	bridge "github.com/structural-patterns/bridge"
	composite "github.com/structural-patterns/composite"
	proxy "github.com/structural-patterns/proxy"
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

	/*
		Example Bridge
	*/
	fmt.Println("*** Example Bridge ***")
	window := bridge.Window{}
	macOS := bridge.MacOS{}
	epson := bridge.Epson{}
	hp := bridge.HP{}

	window.SetPrinter(&epson)
	window.Print()

	macOS.SetPrinter(&hp)
	macOS.Print()

	fmt.Print("*** End of Bridge ***\n\n\n")

	/*
		Example Proxy
	*/
	fmt.Println("*** Example Proxy ***")

	mainDB := proxy.UsersDB{}

	user1 := proxy.User{ID: 1}
	user2 := proxy.User{ID: 2}
	user3 := proxy.User{ID: 3}

	mainDB.Add(user1).Add(user2).Add(user3)

	proxy := proxy.UserFinderProxy{
		MainDB:   mainDB,
		Stack:    proxy.UsersStack{},
		Capacity: 2,
	}

	proxy.Find(1)
	proxy.Find(2)
	proxy.Find(3)
	proxy.Find(2)
	proxy.Find(1)

	fmt.Print("*** End of Proxy ***\n\n\n")
}
