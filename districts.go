package main

import (
	"net/http"
)

var (
	// for goxc builds
	VERSION = "0.0.1"
)

func init() {
	r := Routes()
	http.Handle("/", r)
}

func main() {
	r := Routes()
	http.ListenAndServe(":8080", r)
}
