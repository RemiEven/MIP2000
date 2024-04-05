package memeselector

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"math/rand"
)

func GetRandomMeme(memeFS fs.ReadDirFS) (io.Reader, error) {
	memePath, err := selectMeme(memeFS)
	if err != nil {
		return nil, fmt.Errorf("failed to select meme: %w", err)
	}
	memeBytes, err := fs.ReadFile(memeFS, memePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read meme file: %w", err)
	}
	return bytes.NewReader(memeBytes), nil
}

func selectMeme(memeFS fs.ReadDirFS) (string, error) {
	dirEntries, err := memeFS.ReadDir("images")
	if err != nil {
		return "", fmt.Errorf("failed to read images folder: %w", err)
	}
	numberOfFiles := len(dirEntries)
	memeFiles := make([]string, 0, numberOfFiles)
	for _, file := range dirEntries {
		if file.IsDir() {
			continue
		}
		memeFiles = append(memeFiles, file.Name())
	}
	numberOfMemeFiles := len(memeFiles)
	if numberOfMemeFiles == 0 {
		return "", fmt.Errorf("images folder is empty")
	}
	return "images/" + dirEntries[rand.Intn(numberOfMemeFiles)].Name(), nil
}
