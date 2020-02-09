// parse multiple file templates w/ ParseGlob
// ... and a subfolder for run
package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	println("building bucket")
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	// execute that contents of the bucket
	err := tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// execute just one of templates
	err = tmpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// execute another one
	err = tmpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// execute another
	err = tmpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// execute the whole template
	// it'll just execute the first one loaded
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
