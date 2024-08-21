package widgets

import (
	"log"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/FernandoHDFAraujo/orrery-file/pkg/models"
	"github.com/FernandoHDFAraujo/orrery-file/utils"
	"github.com/joho/godotenv"
)

// TODO: Clean this up this is awful and messy jesus christ
func ArrangePlanets() []models.Planet {
	err := godotenv.Load()

	if err != nil {
		log.Printf("Failed to read .env: %v", err)
		return nil
	}

	dirPath := os.Getenv("NOTE_PATH")

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

		button := widget.NewButton("", func() {
			utils.OpenFile(dirPath + "\\" + file.Name())
		})
		button.Importance = widget.LowImportance

		cardContainer := container.NewStack(cardContent, button)

		card := widget.NewCard("", "", cardContainer)

		cards = append(cards, models.Planet{
			ObjectInCanvas: card,
			Filepath:       file.Name(),
		})
	}

	return cards
}
