package widgets

import (
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/FernandoHDFAraujo/orrery-file/pkg/models"
	"github.com/FernandoHDFAraujo/orrery-file/utils"
)

func ArrangePlanets() []models.Planet {
	filesInDrive := utils.ReadPlanets()

	var cards []models.Planet

	for _, file := range filesInDrive {
		img := canvas.NewImageFromFile(utils.RetrievePlanetIcon(file.Name()))
		img.FillMode = canvas.ImageFillContain
		img.SetMinSize(fyne.NewSize(100, 100))

		text := widget.NewLabel(filepath.Base(file.Name()))

		cardContent := container.NewVBox(
			container.NewCenter(img),
			container.NewCenter(text),
		)

		card := widget.NewCard("", "", cardContent)

		cards = append(cards, models.Planet{
			ObjectInCanvas: card,
			Filepath:       file.Name(),
		})
	}

	return cards
}
