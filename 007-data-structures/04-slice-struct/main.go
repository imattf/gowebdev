package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {

	sage1 := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	sage2 := sage{
		Name:  "MLK",
		Motto: "I have a dream...",
	}

	sage3 := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	sages := []sage{sage1, sage2, sage3}

	err := tmpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
