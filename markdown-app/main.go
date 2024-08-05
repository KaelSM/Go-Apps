package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem
}

var cfg config

func main() {
	/* TODO:
	- Create a new fyne app
	- Create a new window
	- Get the user interface
	- Set the content of the window to the user interface
	- Show the window
	*/

	// Create a new fyne app
	a := app.New()

	// Create a new window
	win := a.NewWindow("Markdown Editor")

	// Get the user interface
	edit, preview := cfg.makeUI()

	// Set the content of the window to the user interface
	win.SetContent(container.NewHSplit(edit, preview))

	// Show the window
	win.Resize(fyne.Size{Width: 800, Height: 600})
	win.CenterOnScreen()
	win.ShowAndRun()
}

func (app *config) makeUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")
	app.EditWidget = edit
	app.PreviewWidget = preview

	edit.OnChanged = preview.ParseMarkdown

	return edit, preview
}
