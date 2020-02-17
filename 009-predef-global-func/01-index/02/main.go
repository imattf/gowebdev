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

	data := struct {
		Words []string
		Lname string
	}{
		xs,
		"Faulkner",
	}

	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
