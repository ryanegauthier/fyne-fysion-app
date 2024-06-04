package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/storage"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(newFysionTheme())
	w := a.NewWindow("Fysion App")
	w.Resize(fyne.NewSize(1024, 768))
	ui := &gui{win: w, title: binding.NewString()}
	w.SetContent(ui.makeGUI())
	w.SetMainMenu(ui.makeMenu())
	// ui.openProject()
	ui.title.AddListener(binding.NewDataListener(func() {
		name, _ := ui.title.Get()
		w.SetTitle("Fysion App - " + name)
	}))
	flag.Usage = func() {
		fmt.Println("Usage: fysion [project directory]")
	}
	flag.Parse()
	if flag.NArg() > 0 {
		flagPath := flag.Arg(0)
		flagPath, err := filepath.Abs(flagPath)
		if err != nil {
			fmt.Println("Error resolving project path", err)
			return
		}

		dirURI := storage.NewFileURI(flagPath)
		dirPath, err := storage.ListerForURI(dirURI)
		if err != nil {
			fmt.Println("Error opening project", err)
			return
		}

		ui.openProject(dirPath)
	} else {
		ui.openProjectDialog()
	}

	w.ShowAndRun()	
}