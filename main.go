package main

import (
  "fmt"
	"log"
  "net/http"
	"path/filepath"
	"github.com/docopt/docopt.go"
)

func main() {
	arguments, _ := docopt.Parse(usage(), nil, true, "0.1", false)

	port := arguments["--port"].(string)
	path, _ := filepath.Abs(arguments["<directory>"].(string))

	start(path, port)
}

func usage() string {
	return `Static Web Server

This tool serves static files in the given directory through http on localhost over the given port number (e.g 5000 by default)

Usage:
	staticws <directory> [--port=N]
	staticws -h | --help
	staticws --version

Options:
	-h --help    Show this screen.
	--version    Show version.
	--port=N     Web server port number [default: 5000].`
}


func start(path, port string) {
	log.Println("Serving files from", path)
	log.Println("Listening on port", port)
  panic(http.ListenAndServe(fmt.Sprintf(":%v", port), http.FileServer(http.Dir(path))))
}