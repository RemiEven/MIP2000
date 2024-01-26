package main

import (
	"io/fs"
	"log"
	"os"

	"github.com/remieven/mip2000/memeselector"
)

func main() {
	fileSystem := os.DirFS(".")
	readDirFileSystem, ok := fileSystem.(fs.ReadDirFS)
	if !ok {
		log.Println("os.DirFS returned a value that doesn't implement fs.ReadDirFS")
		os.Exit(1)
	}
	winner, _ := memeselector.SelectMeme(readDirFileSystem)
	log.Println("And the winner is: ", winner)
}
