package main

import (
	"strconv"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func total(n int) int {
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	return t
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne")
	e := widget.NewEntry()

	e.SetText("0")
	eCulculate := func() {
		n, _ := strconv.Atoi(e.Text)
		l.SetText("Total: " + strconv.Itoa(total(n)))
	}
	b := widget.NewButton("Click", eCulculate)

	widgetSet := widget.NewVBox(
		l,
		e,
		b,
	)
	w.SetContent(widgetSet)

	w.ShowAndRun()
}
