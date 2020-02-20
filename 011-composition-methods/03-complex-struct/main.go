package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {

	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-40", "Intro to Programming I", "4"},
				course{"CSCI-41", "Intro to Programming II", "4"},
				course{"CSCI-42", "Intro to Programming III", "4"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"CSCI-50", "Intermediate Programming I", "4"},
				course{"CSCI-51", "Intermediate  Programming II", "4"},
				course{"CSCI-52", "Intermediate  Programming III", "4"},
			},
		},
	}

	err := tmpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}
