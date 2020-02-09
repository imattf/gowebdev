// parse multiple file templates w/ ParseGlob
// ... and a subfolder for run
package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// build pointer to template; i.e. the bucket
	tmpl, err := template.ParseGlob("templates/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// execute that contents of the bucket
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// execute just one of templates
	err = tmpl.ExecuteTemplate(os.Stdout, "01one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// execute another one
	err = tmpl.ExecuteTemplate(os.Stdout, "02two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// execute another
	err = tmpl.ExecuteTemplate(os.Stdout, "00three.gohtml", nil)
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
