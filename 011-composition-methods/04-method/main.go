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

func (p person) SomeProcessing() int {
	return 7
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) TakesArg(x int) int {
	return x * 2
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
