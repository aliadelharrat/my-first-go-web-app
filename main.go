package main

import (
	"fmt"
	"go/my-web-application/views"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var homeView *views.View
var contactView *views.View

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := homeView.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := contactView.Execute(w, nil); err != nil {

		panic(err)
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Page not found</h1><p>The page you're looking for does not exist!</p>")
}

func main() {
	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")

	r := httprouter.New()
	r.GET("/", home)
	r.GET("/contact", contact)
	r.NotFound = http.HandlerFunc(notFound)

	log.Fatal(http.ListenAndServe(":8080", r))
}
