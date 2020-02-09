// template variables using pipeline assignment
// trick is in the template
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

	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", `poof`)
	if err != nil {
		log.Fatalln(err)
	}

}
