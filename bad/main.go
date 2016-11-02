package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	// Not good practice

	paths1, err := filepath.Glob("../templates/404/*.html")
	check(err)

	paths2, err := filepath.Glob("../templates/includes/*.html")
	check(err)

	paths := append(paths1, paths2...)

	templates, err := template.New("main").ParseFiles(paths...)
	check(err)

	err = templates.ExecuteTemplate(os.Stdout, "404.html", nil)
	check(err)
}
