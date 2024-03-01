package main

import (
	"log"
	"net/http"
	"time"

	"github.com/remieven/mip2000/api"
)

func main() {

	s := &http.Server{
		Addr:           ":8080",
		Handler:        api.NewRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
	// fileSystem := os.DirFS(".")
	// readDirFileSystem, ok := fileSystem.(fs.ReadDirFS)
	// if !ok {
	// 	log.Println("os.DirFS returned a value that doesn't implement fs.ReadDirFS")
	// 	os.Exit(1)
	// }
	// winner, _ := memeselector.SelectMeme(readDirFileSystem)
	// log.Println("And the winner is: ", winner)
}
