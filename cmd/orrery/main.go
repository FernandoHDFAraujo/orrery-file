package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/FernandoHDFAraujo/orrery-file/pkg/ui/widgets"
)

func main() {
	a := app.New()
	w := a.NewWindow("Orrery")

	cards := widgets.ArrangePlanets()

	grid := container.NewGridWithColumns(3, cards...)

	w.SetContent(grid)
	w.ShowAndRun()
}
