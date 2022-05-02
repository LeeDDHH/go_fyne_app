package main

import (
	"strconv"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	c := 0
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne")
	cCount := func(){
		c++
		l.SetText("count: " + strconv.Itoa(c))
	}
	b := widget.NewButton("Click", cCount)
	
	widgetSet := widget.NewVBox(
		l,
		b,
	)
	w.SetContent(widgetSet)

	w.ShowAndRun()
}
