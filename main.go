package main

import (
  "fmt"
	"log"
	"os"
  "net/http"
	"path/filepath"
	"github.com/docopt/docopt.go"
)

func main() {
	arguments, _ := docopt.Parse(usage(), nil, true, "0.2", false)

	port := arguments["--port"].(string)
	path, _ := filepath.Abs(arguments["<directory>"].(string))

	ok, _ := exists(path)
	if !ok {
		log.Println(path, "does not exists")
		os.Exit(1)
	}

	start(path, port)
}

func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return false, err
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

	http.Handle("/", http.FileServer(http.Dir(path)))
	panic(http.ListenAndServe(fmt.Sprintf(":%v", port), Log(http.DefaultServeMux)))
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

