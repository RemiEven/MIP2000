package main

import (
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/remieven/mip2000/api"
	"github.com/remieven/mip2000/memeselector"
)

func main() {

	s := &http.Server{
		Addr:           ":8080",
		Handler:        http.HandlerFunc(api.Coucou),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
	fileSystem := os.DirFS(".")
	readDirFileSystem, ok := fileSystem.(fs.ReadDirFS)
	if !ok {
		log.Println("os.DirFS returned a value that doesn't implement fs.ReadDirFS")
		os.Exit(1)
	}
	winner, _ := memeselector.SelectMeme(readDirFileSystem)
	log.Println("And the winner is: ", winner)
}
