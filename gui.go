package main

import (
	"image/color"
	// "log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type gui struct {
	win fyne.Window
	//Project Label
	title binding.String
	// directory *widget.Label
}
	

func makeBanner() fyne.CanvasObject {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {}),
	)
	logo := canvas.NewImageFromResource(resourceApexLogoPng)
	logo.FillMode = canvas.ImageFillContain
	return container.NewStack(toolbar, container.NewPadded(logo))
}

func (g* gui) makeGUI() fyne.CanvasObject {
	top := makeBanner()
	left := widget.NewLabel("Left")
	right := widget.NewLabel("Right")

	// content := widget.NewLabel("CONTENT")
	directory := widget.NewLabelWithData(g.title)
	// directory := canvas.NewLabel("Directory")
	content := container.NewStack(canvas.NewRectangle(color.White), directory)
	
	dividers := [3]fyne.CanvasObject{
		widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(),
	}
	objs := []fyne.CanvasObject{top, left, right, content, dividers[0], dividers[1], dividers[2]}
	return container.New(newFysionLayout(top, left, right, content, dividers), objs...)
}

func (g* gui) openProjectDialog() {
	// Open a file dialog to select a project
	dialog.ShowFolderOpen(func(dir fyne.ListableURI, err error) {
		// log.Println("Selected: ", dir)
		if err != nil {
			dialog.ShowError(err, g.win)
			return
		}
		if dir == nil {
			g.openProjectDialog()
			return
		}

		g.openProject(dir)

	}, g.win)
}

func (g *gui) makeMenu() *fyne.MainMenu {
	file := fyne.NewMenu("File",
		fyne.NewMenuItem("Open Project", g.openProjectDialog),
	)
	return fyne.NewMainMenu(file)
}

func (g *gui) openProject(dir fyne.ListableURI) {
	name := dir.Path()

	// Load the project
	g.title.Set(name)
}