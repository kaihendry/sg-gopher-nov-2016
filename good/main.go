package main

import (
	"html/template"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	templates := template.Must(template.New("main").ParseGlob("../templates/404/*.html"))
	templates = template.Must(templates.ParseGlob("../templates/includes/*.html"))

	err := templates.ExecuteTemplate(os.Stdout, "404.html", struct {
		Title    string
		Synopsis string
	}{
		"Glorious 404 page",
		"Page was not found",
	})
	check(err)
}
