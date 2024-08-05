package main

import (
	"io/ioutil"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

// config holds information that we can share with any function
// that has a receiver of type *config
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
	//a := app.New()

	// create a fyne app
	a := app.NewWithID("my markdown editor")
	a.Settings().SetTheme(&myTheme{})

	// create a window for the app
	win := a.NewWindow("Markdown")

	// get the user interface
	edit, preview := cfg.makeUI()
	cfg.createMenuItems(win)

	// set the content of the window
	win.SetContent(container.NewHSplit(edit, preview))

	// show window and run app
	win.Resize(fyne.Size{Width: 800, Height: 500})
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

	// create three menu items
	openMenuItem := fyne.NewMenuItem("Open...", app.openFunc(win))
	saveMenuItem := fyne.NewMenuItem("Save", app.saveFunc(win))
	app.SaveMenuItem = saveMenuItem
	app.SaveMenuItem.Disabled = true
	saveAsMenuItem := fyne.NewMenuItem("Save as...", app.saveAsFunc(win))

	// create a file menu, and add the three items to it
	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)

	// create a main menu, and add the file menu to it
	menu := fyne.NewMainMenu(fileMenu)

	// set the main menu for the application
	win.SetMainMenu(menu)
}

// filter so only md files are opened
var filter = storage.NewExtensionFileFilter([]string{".md", ".markdown", ".MD", ".MARKDOWN"})

func (app *config) saveFunc(win fyne.Window) func() {
	return func() {
		if app.CurrentFile != nil {
			write, err := storage.Writer(app.CurrentFile)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			write.Write([]byte(app.EditWidget.Text))
			defer write.Close()
		}
	}
}

// openFunc returns a function that opens a file dialog to load the content of a selected file into the EditWidget.
//
// Parameters:
// - app: a pointer to the config struct which holds the application's configuration and state.
// - win: the window to which the dialog will be attached.
//
// Returns:
// - func(): a function that opens the file dialog.
//
// The function performs the following steps:
//  1. Creates a new file open dialog with a callback function to handle the file reading process.
//  2. In the callback function:
//     a. If an error occurs, it shows an error dialog and returns.
//     b. If the user cancels the open operation (read is nil), it simply returns.
//     c. Otherwise, it reads the content of the selected file.
//     d. Closes the file reader after reading.
//     e. Sets the text of the EditWidget to the content of the file.
//     f. Updates the CurrentFile field in the config struct with the URI of the opened file.
//     g. Updates the window title to include the name of the opened file.
//     h. Enables the Save menu item.
//  3. Shows the open dialog.
func (app *config) openFunc(win fyne.Window) func() {
	return func() {
		openDialog := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			if read == nil {
				return
			}

			defer read.Close()

			data, err := ioutil.ReadAll(read)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			app.EditWidget.SetText(string(data))

			app.CurrentFile = read.URI()
			win.SetTitle(win.Title() + " - " + read.URI().Name())
			app.SaveMenuItem.Disabled = false

		}, win)

		openDialog.SetFilter(filter)
		openDialog.Show()
	}
}

// saveAsFunc returns a function that opens a "Save As" dialog to save the content of the EditWidget to a file.
//
// Parameters:
// - app: a pointer to the config struct which holds the application's configuration and state.
// - win: the window to which the dialog will be attached.
//
// Returns:
// - func(): a function that opens the "Save As" dialog.
//
// The function performs the following steps:
//  1. Creates a new file save dialog with a callback function to handle the file saving process.
//  2. In the callback function:
//     a. If an error occurs, it shows an error dialog and returns.
//     b. If the user cancels the save operation (write is nil), it simply returns.
//     c. If the file name does not end with ".md", it shows an information dialog and returns.
//     d. Otherwise, it writes the content of the EditWidget to the selected file.
//     e. Updates the CurrentFile field in the config struct with the URI of the saved file.
//     f. Closes the file writer after writing.
//     g. Updates the window title to include the name of the saved file.
//     h. Enables the Save menu item.
//  3. Sets the default file name to "untitled.md".
//  4. Sets a filter for the file dialog.
//  5. Shows the save dialog.
func (app *config) saveAsFunc(win fyne.Window) func() {
	return func() {
		saveDialog := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			if write == nil {
				// user cancelled
				return
			}

			if !strings.HasSuffix(strings.ToLower(write.URI().String()), ".md") {
				dialog.ShowInformation("Error", "Please name your file with .md extension", win)
				return
			}

			// save file
			write.Write([]byte(app.EditWidget.Text))
			app.CurrentFile = write.URI()

			defer write.Close()

			win.SetTitle(win.Title() + " - " + write.URI().Name())
			app.SaveMenuItem.Disabled = false

		}, win)

		saveDialog.SetFileName("untitled.md")
		saveDialog.SetFilter(filter)
		saveDialog.Show()
	}
}
