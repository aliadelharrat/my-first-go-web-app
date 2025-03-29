package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const portNumber = ":8080"

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "<h1>Welcome body</h1>")
}

func user(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "<h1>Welcome %s</h1>", ps.ByName("name"))
}

func main() {
	router := httprouter.New()

	// My routes
	router.GET("/", home)
	router.GET("/user/:name", user)

	log.Fatal(http.ListenAndServe(portNumber, router))
}
