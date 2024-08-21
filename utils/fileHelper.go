package utils

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func GetFileExtension(path string) string {
	var extension string

	for i := 0; i < len(path)-1; i++ {
		if string(path[len(path)-i-1]) == "." {
			extension = path[len(path)-i:]
			break
		}
	}

	return extension
}

func RetrievePlanetIcon(path string) string {
	err := godotenv.Load()

	if err != nil {
	}

	baseAssetPath := os.Getenv("ICON_PATH")
	extension := GetFileExtension(path)

	var iconPartialPath string

	switch extension {
	case "pdf":
		iconPartialPath = "\\pdf.png"
	case "txt":
		iconPartialPath = "\\txt.png"
	default:
		iconPartialPath = "\\default.png"
	}

	completePath := filepath.Join(baseAssetPath + iconPartialPath)

	return completePath
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

func retrieveIconFromSystem() {
	//Logic for retrieving icons from the user's theme goes here
}
