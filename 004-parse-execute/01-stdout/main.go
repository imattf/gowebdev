// output to stdout
package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tmpl, err := template.ParseFiles("tmpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
