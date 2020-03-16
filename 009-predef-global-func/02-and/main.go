package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

type user struct {
	Name  string
	Motto string
	Admin bool
}

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {

	u1 := user{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
		Admin: false,
	}

	u2 := user{
		Name:  "Gandhi",
		Motto: "Be the change",
		Admin: true,
	}

	u3 := user{
		Name:  "",
		Motto: "Nobody",
		Admin: true,
	}

	users := []user{u1, u2, u3}

	err := tmpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}