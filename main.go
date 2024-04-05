package main

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/remieven/mip2000/api"
)

func main() {

	fileSystem := os.DirFS(".")
	readDirFileSystem, ok := fileSystem.(fs.ReadDirFS)
	if !ok {
		fmt.Println("os.DirFS returned a value that doesn't implement fs.ReadDirFS")
		os.Exit(1)
	}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        api.NewRouter(readDirFileSystem),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
