package main

import (
	"html/template"
	"os"
)

type User struct {
	Name        string
	Dog         string
	NumOfKisses int
	Hobbies     []string
}

func main() {
	t, err := template.ParseFiles("./hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name:        "Adel",
		Dog:         "Kaouter",
		NumOfKisses: 5,
		Hobbies:     []string{"Programming", "Watching animes", "Working in pharmacy"},
		// Hobbies: []string{},
	}

	if err = t.Execute(os.Stdout, data); err != nil {
		panic(err)
	}
}
