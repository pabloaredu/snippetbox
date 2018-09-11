package main

import (
	"flag"
	"log"
	"net/http"
	"../../pkg/models"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	htmlDir := flag.String("html-dir", "./ui/html", "Path to HTML templates")
	staticDir := flag.String("static-dir", "./ui/static", "Path to static assets")	

	flag.Parse()

	app := &App {
		Database: &models.Database{},
		HTMLDir: *htmlDir,
		StaticDir: *staticDir,
	}

	log.Printf("Starting server on %s:", *addr)
	err := http.ListenAndServe(*addr, app.Routes())
	log.Fatal(err)
}


