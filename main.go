package main

import (
	"strconv"

	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
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

	// エントリー
	e := widget.NewEntry()
	e.SetText("0")
	eCulculate := func() {
		n, _ := strconv.Atoi(e.Text)
		l.SetText("Total: " + strconv.Itoa(total(n)))
	}
	b := widget.NewButton("Click", eCulculate)

	// テーマ
	darkTheme := theme.DarkTheme()
	lightTheme := theme.LightTheme()
	currentTheme := darkTheme
	changeTheme := func() {
		changedTheme := darkTheme
		if currentTheme == changedTheme {
			currentTheme = lightTheme
			changedTheme = lightTheme
		} else {
			currentTheme = darkTheme
		}
		a.Settings().SetTheme(changedTheme)
	}
	changeThemeButton := widget.NewButton("Change Theme", changeTheme)

	// チェックボックス
	checkEvent := func(f bool){
		if f {
			l.SetText("checked!")
		} else {
			l.SetText("not checked")
		}
	}
	c := widget.NewCheck("checkbox", checkEvent)
	c.SetChecked(true)

	widgetSet := widget.NewVBox(
		l,
		e,
		b,
		changeThemeButton,
		c,
	)
	w.SetContent(widgetSet)

	w.ShowAndRun()
}
