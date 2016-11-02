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

	paths1, err := filepath.Glob("../templates/404/*.html")
	check(err)

	paths2, err := filepath.Glob("../templates/includes/*.html")
	check(err)

	paths := append(paths1, paths2...)

	templates := template.Must(template.New("main").ParseFiles(paths...))

	err = templates.ExecuteTemplate(os.Stdout, "404.html", nil)
	check(err)
}
