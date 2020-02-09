// output to a file
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

	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("error creating file", err)
	}

	defer newFile.Close()

	err = tmpl.Execute(newFile, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
