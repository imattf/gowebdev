// date formatting

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

	xs := []string{"zero", "one", "two", "three", "four", "five"}

	err := tmpl.Execute(os.Stdout, xs)
	if err != nil {
		log.Fatalln(err)
	}
}
