package main

import (
	"fmt"
	"io"

	"fyne.io/fyne/v2"
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