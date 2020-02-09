// output to stdout
package main

import (
	"html/template"
	"log"
	"os"
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
