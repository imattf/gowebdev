package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

type dblZero struct {
	person
	Ltk bool
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {

	p1 := dblZero{
		person{
			Name: "Bondo, Paint",
			Age:  42,
		},
		true,
	}

	err := tmpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
