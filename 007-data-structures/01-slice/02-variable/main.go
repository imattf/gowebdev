package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {
	sages := []string{"Ghandi", "Buddha", "Jesus"}

	err := tmpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
