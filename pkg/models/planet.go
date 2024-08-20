package models

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"fyne.io/fyne/v2"
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

func ReadPlanets() []fyne.CanvasObject {
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

	var c []*widget.Card
	for _, file := range filesInDrive {
		c = append(c, &widget.Card{
			Title: filepath.Base(file.Name()),
			Image: canvas.NewImageFromFile(utils.RetrievePlanetIcon(file.Name())),
		})
	}

	var cards []fyne.CanvasObject
	for _, card := range c {
		cards = append(cards, card)
	}

	return cards
}
