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
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	err := tmpl.Execute(os.Stdout, sage1)
	if err != nil {
		log.Fatalln(err)
	}
}
