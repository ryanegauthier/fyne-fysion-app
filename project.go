package main

import (
	"fmt"
	"io"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/storage"
)

func createProject(name string, parent fyne.ListableURI) (fyne.ListableURI, error) {
	// Create a new project
	dir, err := storage.Child(parent, name)
	if err != nil {
		return nil, err
	}
	err = storage.CreateListable(dir)
	if err != nil {
		return nil, err
	}
	mod, err := storage.Child(dir, "go.mod")
	if err != nil {
		return nil, err
	}
	w, err := storage.Writer(mod)
	if err != nil {
		return nil, err
	}
	defer w.Close()

	_, err = io.WriteString(w, fmt.Sprintf("module %s\n go 1.22\n require fyne.io/fyne/v2 v2.0.0\n", name))

	// We've created the directory as a listableUri - no error means success
	list, _ := storage.ListerForURI(dir)
	return list, err
}

func (g *gui) openProject(dir fyne.ListableURI) {
	name := dir.Name()

	// Load the project
	g.title.Set(name)

	// empty the data binding
	g.fileTree.Set(map[string][]string{}, map[string]fyne.URI{})

	addFilesToTree(dir, g.fileTree, binding.DataTreeRootID)
}

func addFilesToTree(dir fyne.ListableURI, tree binding.URITree, root string) {
	// Set the file tree to the new project
	items, _ := dir.List()
	for _, uri := range items {
		nodeID := uri.String()
		tree.Append(root, nodeID, uri)

		isDir, err := storage.CanList(uri)
		if err != nil {
			log.Println("Error checking if URI is listable:", err)
		}
		if isDir {
			child, _ := storage.ListerForURI(uri)
			addFilesToTree(child, tree, nodeID)
		}
	}
}

