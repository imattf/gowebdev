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

type car struct {
	Make  string
	Model string
	Doors int
}

// type items struct {
// 	Wisdom    []sage
// 	Transport []car
// }

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

	car1 := car{
		Make:  "Ford",
		Model: "F150",
		Doors: 2,
	}

	car2 := car{
		Make:  "Toyota",
		Model: "Corolla",
		Doors: 4,
	}

	sages := []sage{sage1, sage2, sage3}
	cars := []car{car1, car2}

	// was...
	// data := items{
	// 	Wisdom:    sages,
	// 	Transport: cars,
	// }

	// is...
	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
