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

// makeUI initializes and returns a multi-line text entry widget and a rich text preview widget.
// It sets up the preview widget to update its content based on the markdown text entered in the text entry widget.
//
// Parameters:
// - app: a pointer to the config struct which holds the application's configuration and state.
//
// Returns:
// - *widget.Entry: a pointer to the multi-line text entry widget.
// - *widget.RichText: a pointer to the rich text preview widget.
//
// The function performs the following steps:
// 1. Creates a new multi-line text entry widget.
// 2. Creates a new rich text widget initialized with empty markdown content.
// 3. Assigns the created widgets to the EditWidget and PreviewWidget fields of the config struct.
// 4. Sets up an event handler for the text entry widget to update the preview widget's content
func (app *config) makeUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")
	app.EditWidget = edit
	app.PreviewWidget = preview

	edit.OnChanged = preview.ParseMarkdown

	return edit, preview
}

// createMenuItems sets up the main menu for the given window with "File" menu items: Open, Save, and Save As.
//
// Parameters:
// - app: a pointer to the config struct which holds the application's configuration and state.
// - win: the window to which the menu will be attached.
//
// The function performs the following steps:
// 1. Creates a new menu item "Open..." with an empty action function.
// 2. Creates a new menu item "Save" with an empty action function.
// 3. Creates a new menu item "Save As..." with an empty action function.
// 4. Creates a "File" menu containing the "Open...", "Save", and "Save As..." menu items.
// 5. Creates the main menu with the "File" menu.
// 6. Sets the created main menu to the provided window.
func (app *config) createMenuItems(win fyne.Window) {

	openMenuItem := fyne.NewMenuItem("Open...", func() {})

	saveMenuItem := fyne.NewMenuItem("Save", func() {})

	saveAsMenuItem := fyne.NewMenuItem("Save As...", func() {})

	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)

	menu := fyne.NewMainMenu(fileMenu)

	win.SetMainMenu(menu)

}
