package models

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/FernandoHDFAraujo/orrery-file/utils"
	"github.com/joho/godotenv"
)

type Planet struct {
	Name        string
	Path        string
	Type        string
	LastUpdated time.Time
}

func DrawPlanet() []*widget.Card {
	err := godotenv.Load()

	if err != nil {
		log.Printf("Failed to read .env: %v", err)
		return nil
	}

	filesInDrive, err := os.ReadDir(os.Getenv("NOTE_PATH"))

	if err != nil {
		log.Printf("Failed to read directory: %v", err)
		return nil
	}

	var cards []*widget.Card
	for _, file := range filesInDrive {
		cards = append(cards, &widget.Card{
			Title: filepath.Base(file.Name()),
			Image: canvas.NewImageFromFile(utils.RetrievePlanetIcon(file.Name())),
		})
	}

	return cards
}
