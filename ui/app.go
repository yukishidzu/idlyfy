package ui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Run() {
	a := app.New()
	w := a.NewWindow("Idlyfy")

	label := widget.NewLabel("Welcome to Idlyfy!")
	w.SetContent(container.NewVBox(
		label,
		widget.NewButton("Settings", func() {
			label.SetText("Settings clicked")
		}),
	))

	w.ShowAndRun()
}
