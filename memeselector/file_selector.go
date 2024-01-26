package memeselector

import (
	"fmt"
	"io/fs"
	"log"
)

func SelectMeme(memeFS fs.ReadDirFS) (string, error) {
	dirEntries, err := memeFS.ReadDir("images")
	if err != nil {
		return "", fmt.Errorf("failed to read images folder: %w", err)
	}
	for _, dirEntry := range dirEntries {
		log.Println(dirEntry.Name())
	}
	return "un vrai de vrai", nil
}
