// output to stdout w/ data replacing nil value
package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	println("building template")
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {

	err := tmpl.Execute(os.Stdout, "poof")
	if err != nil {
		log.Fatalln(err)
	}

}
