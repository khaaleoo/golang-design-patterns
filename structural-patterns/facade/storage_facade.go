package facade

import (
	"fmt"
	"strings"

	composite "github.com/structural-patterns/composite"
)

type StorageFacade struct {
	root    *composite.Folder
	folders map[string]*composite.Folder
}

func NewStorageFacade(rootName string) *StorageFacade {
	root := &composite.Folder{}
	root.SetName(rootName)
	return &StorageFacade{
		root:    root,
		folders: map[string]*composite.Folder{"/": root},
	}
}

func (s *StorageFacade) SaveFile(folderPath string, fileName string) {
	folder := s.ensureFolder(folderPath)
	file := &composite.File{}
	file.SetName(fileName)
	folder.Add(file)
}

func (s *StorageFacade) ListAll() {
	s.root.Print()
}

func (s *StorageFacade) ensureFolder(folderPath string) *composite.Folder {
	normalized := normalizePath(folderPath)
	if folder, ok := s.folders[normalized]; ok {
		return folder
	}

	segments := strings.Split(strings.Trim(normalized, "/"), "/")
	currentPath := "/"
	current := s.root

	for _, segment := range segments {
		currentPath = currentPath + segment + "/"
		if folder, ok := s.folders[currentPath]; ok {
			current = folder
			continue
		}

		folder := &composite.Folder{}
		folder.SetName(segment)
		current.Add(folder)
		s.folders[currentPath] = folder
		current = folder
	}

	return current
}

func normalizePath(folderPath string) string {
	if folderPath == "" || folderPath == "/" {
		return "/"
	}

	path := folderPath
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	return path
}

func (s *StorageFacade) Summary() string {
	return fmt.Sprintf("Storage root: %s", s.root.GetName())
}
