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

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {

	p1 := person{
		Name: "Bondo, Paint",
		Age:  42,
	}

	err := tmpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
