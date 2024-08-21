package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/FernandoHDFAraujo/orrery-file/pkg/ui/widgets"
)

func main() {
	a := app.New()
	w := a.NewWindow("Orrery")

	cards := widgets.ArrangePlanets()

	var canvasCards []fyne.CanvasObject
	for _, card := range cards {
		canvasCards = append(canvasCards, card.ObjectInCanvas)
	}

	grid := container.NewGridWithColumns(3, canvasCards...)

	w.SetContent(grid)
	w.ShowAndRun()
}
