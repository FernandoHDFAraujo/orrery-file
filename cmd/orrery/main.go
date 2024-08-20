package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/FernandoHDFAraujo/orrery-file/pkg/models"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	cards := models.DrawPlanet()

	cardObjects := make([]fyne.CanvasObject, len(cards))
	for i, card := range cards {
		cardObjects[i] = card
	}

	grid := container.NewGridWithColumns(3, cardObjects...)

	w.SetContent(grid)
	w.ShowAndRun()
}
