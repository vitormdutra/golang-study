package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")

	myWindow.SetContent(widget.NewLabel("Hello, World!"))
	myWindow.ShowAndRun()
}
