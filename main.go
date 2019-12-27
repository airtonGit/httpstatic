package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
)

// FileSystem custom file system handler
type FileSystem struct {
	fs http.FileSystem
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	log.Printf("Serving %s", path)
	f, err := fs.fs.Open(path)
	if err != nil {
		log.Printf("Serving error %s", err.Error())
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		log.Printf("Serving error %s", err.Error())
	}
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.fs.Open(index); err != nil {
			return nil, err
		}
		log.Printf("Serving %s", index)
	}

	return f, nil
}

func main() {
	port := flag.String("p", "80", "port to serve on")
	directory := flag.String("d", "/pub", "the directory of static file to host")
	flag.Parse()

	fileServer := http.FileServer(FileSystem{http.Dir(*directory)})
	http.Handle("/", fileServer) // Original atende apenas /static/

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
