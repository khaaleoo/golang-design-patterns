package main

import (
	"fmt"

	adapter "github.com/structural-patterns/adapter"
	bridge "github.com/structural-patterns/bridge"
	composite "github.com/structural-patterns/composite"
	decorator "github.com/structural-patterns/decorator"
	facade "github.com/structural-patterns/facade"
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
		Example Decorator
	*/
	fmt.Println("*** Example Decorator ***")
	fetchClient := &adapter.FetchAdapter{Instance: &adapter.Fetch{}}
	httpClient := &decorator.LoggingDecorator{
		Inner: &decorator.RetryDecorator{
			Inner:   fetchClient,
			Retries: 1,
		},
	}
	client.Get(httpClient, "https://www.example.com")
	fmt.Print("*** End of Decorator ***\n\n\n")

	/*
		Example Facade
	*/
	fmt.Println("*** Example Facade ***")
	storage := facade.NewStorageFacade("Root")
	storage.SaveFile("/docs", "README.md")
	storage.SaveFile("/docs", "guide.md")
	storage.SaveFile("/images", "logo.png")
	fmt.Println(storage.Summary())
	storage.ListAll()
	fmt.Print("*** End of Facade ***\n\n\n")

	/*
		Example Proxy (lazy HTTP client)
	*/
	fmt.Println("*** Example Proxy (Lazy HTTP) ***")
	lazyClient := proxy.NewLazyHttpProxy(func() adapter.Http {
		return &adapter.FetchAdapter{Instance: &adapter.Fetch{}}
	})
	client.Get(lazyClient, "https://www.example.com")
	client.Get(lazyClient, "https://www.example.com")
	fmt.Print("*** End of Proxy (Lazy HTTP) ***\n\n\n")

	/*
		Example Proxy (cache)
	*/
	fmt.Println("*** Example Proxy (Cache) ***")

	mainDB := proxy.UsersDB{}

	user1 := proxy.User{ID: 1}
	user2 := proxy.User{ID: 2}
	user3 := proxy.User{ID: 3}

	mainDB.Add(user1).Add(user2).Add(user3)

	userFinderProxy := proxy.UserFinderProxy{
		MainDB:   mainDB,
		Stack:    proxy.UsersStack{},
		Capacity: 2,
	}

	userFinderProxy.Find(1)
	userFinderProxy.Find(2)
	userFinderProxy.Find(3)
	userFinderProxy.Find(2)
	userFinderProxy.Find(1)

	fmt.Print("*** End of Proxy (Cache) ***\n\n\n")
}
