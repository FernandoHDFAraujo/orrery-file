package models

import (
	"io/fs"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Planet struct {
	Name        string
	Path        string
	Type        string
	LastUpdated time.Time
}

func ReadPlanets() []fs.DirEntry {
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

	return filesInDrive
}
