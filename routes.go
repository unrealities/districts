package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Districts.co: Domination without Legislation")
}

func Routes() http.Handler {
	router := httprouter.New()

	router.GET("/", Index)
	router.ServeFiles("/static/*filepath", http.Dir("static"))

	return router
}
