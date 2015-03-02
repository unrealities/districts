package main

import (
	"log"
	"net/http"
)

var (
	// for goxc builds
	VERSION = "0.0.1"
)

func main() {
	r := Routes()
	log.Fatal(http.ListenAndServe(":8080", r))
}
