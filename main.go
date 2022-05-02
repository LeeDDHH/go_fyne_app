package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("Hello")
	// widgetSet := widget.NewVBox(
	widgetSet := widget.NewHBox(
		widget.NewLabel("Hello Fyne"),
		widget.NewLabel("second text"),
	)
	w.SetContent(widgetSet)

	w.ShowAndRun()
}
