// # Serve the files in the "starting-files" folder
// ## To get your images to serve, use only this:
// 	fs := http.FileServer(http.Dir("public"))
// Hint: look to see what type FileServer returns, then think it through.

package main

import (
	"log"
	"net/http"
	"text/template"
)

var tmpl *template.Template

func init() {
	//println("building template")
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {

	fs := http.FileServer(http.Dir("public"))
	http.HandleFunc("/", dogs)
	http.Handle("/pics", fs)
	http.ListenAndServe(":8080", nil)

}

func dogs(w http.ResponseWriter, req *http.Request) {

	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}

}
