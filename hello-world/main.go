package main

import (
	//"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	output *widget.Label
}

var myApp App

func main() {
	a := app.New()
	w := a.NewWindow("Hello, World!")

	//w.SetContent(widget.NewLabel("Hello, Fyne!"))

	output, entry, btn := myApp.makeUI()
	w.SetContent(container.NewVBox(output, entry, btn))
	w.Resize(fyne.Size{Width: 500, Height: 500})
	w.ShowAndRun()
	//w.Show()
	//a.Run()

	//tydy()
}

// func tydy() {
// 	fmt.Print("would tidy up")
// }

func (app *App) makeUI() (*widget.Label, *widget.Entry, *widget.Button) {
	output := widget.NewLabel("Hello, Fyne!")
	entry := widget.NewEntry()
	btn := widget.NewButton("Enter", func() {
		app.output.SetText(entry.Text)
	})
	btn.Importance = widget.HighImportance
	app.output = output
	return output, entry, btn
}
