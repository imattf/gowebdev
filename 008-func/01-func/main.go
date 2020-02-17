package main

import (
	"log"
	"os"
	"strings"
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

// create a FuncMap to register functions.
// "uc" is what the func will be called in the template
// "uc" is the ToUpper func from package strings
// "ft" is a func I declared
// "ft" slices a string, returning the first three characters
var fmap = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	//was, but doesn't wotk when passing functions
	//tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
	// template.New("").Funcs(fmap). is the expanded piece
	tmpl = template.Must(template.New("").Funcs(fmap).ParseFiles("tmpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	// if len(s) >= 3 {
	// 	s = s[:3]
	// }
	s = s[:3]
	return s
}

func main() {

	sage1 := sage{
		Name:  "gandhi",
		Motto: "Be the change",
	}

	sage2 := sage{
		Name:  "mlk",
		Motto: "I have a dream...",
	}

	sage3 := sage{
		Name:  "buddha",
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

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
