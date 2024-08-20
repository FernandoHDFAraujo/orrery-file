package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/FernandoHDFAraujo/orrery-file/pkg/models"
)

func main() {
	a := app.New()
	w := a.NewWindow("Orrery")

	cards := models.ReadPlanets()

	grid := container.NewGridWithColumns(3, cards...)

	w.SetContent(grid)
	w.ShowAndRun()
}
