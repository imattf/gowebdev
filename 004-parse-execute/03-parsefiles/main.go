// parse multiple file templates
package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// build pointer to template; i.e. the bucket
	tmpl, err := template.ParseFiles("three.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// execute that contents of the bucket
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// append two more templates to the bucket
	tmpl, err = tmpl.ParseFiles("two.gohtml", "one.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// execute just one of appended
	err = tmpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// execute the other appended one
	err = tmpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// execute the orig template
	err = tmpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
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
