package main

import (
	"errors"
	"fysion/internal/dialogs"
	"image/color"
	"os"

	// "log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
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

func (g *gui) makeGUI() fyne.CanvasObject {
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

func (g *gui) openProjectDialog() {
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

func (g *gui) makeCreateDetail(wizard dialogs.Wizard) fyne.CanvasObject {
	homeDir, _ := os.UserHomeDir()
	parent := storage.NewFileURI(homeDir)
	selectedDir, _ := storage.ListerForURI(parent)

	name := widget.NewEntry()
	name.Validator = func(s string) error {
		if s == "" {
			return errors.New("Name cannot be empty")
		}
		return nil
	}

	var location *widget.Button
	location = widget.NewButton(selectedDir.Name(), func() {
		d := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil || uri == nil {
				dialog.ShowError(err, g.win)
				return
			}
			selectedDir = uri
			location.SetText(selectedDir.Name())
		}, g.win)

		d.SetLocation(selectedDir)
		d.Show()
	})
	form := widget.NewForm(
		widget.NewFormItem("Name", name),
		widget.NewFormItem("Parent Directory", location),
	)
	form.OnSubmit = func() {
		project, err := createProject(name.Text, selectedDir)
		if err != nil {
			dialog.ShowError(err, g.win)
			return
		}
		wizard.Hide()
		g.openProject(project)
	}
	return form
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

func (g *gui) showCreate(w fyne.Window) {
	var wizard *dialogs.Wizard

	// Show a dialog to create a new project
	introText := widget.NewLabel(`Create a new project.
	
Or open an existing one.`)
	// End dialog in the intro text

	// Buttons for open and create
	open := widget.NewButton("Open Project", func() {
		wizard.Hide()
		g.openProjectDialog()
	})
	create := widget.NewButton("Create Project", func() {
		// step2 := widget.NewLabel("Step 2 Content")
		wizard.Push("Project Details", g.makeCreateDetail(*wizard))
	})
	create.Importance = widget.HighImportance

	// Need container for open and create buttons
	buttonContainer := container.NewGridWithColumns(2, open, create)

	home := container.NewVBox(introText, buttonContainer)

	// Show a dialog to create a new project
	wizard = dialogs.NewWizard("Create Project", home)
	// if home == "" {
	// 	return
	// }

	// Start the wizard
	wizard.Show(w)
	wizard.Resize(home.MinSize().AddWidthHeight(40, 80)) //fyne.NewSize(400, 300))
}
