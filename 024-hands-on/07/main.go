// # Serve the files in the "starting-files" folder
// ## To get your images to serve, use:
// 	func StripPrefix(prefix string, h Handler) Handler
// 	func FileServer(root FileSystem) Handler
// Constraint: you are not allowed to change the route being used for images in the template file

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

	http.HandleFunc("/", dogs)
	// http.Handle("/pics", fs)

	// my way...
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public", fs))

	// todd's way...
	// http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("public"))))

	http.ListenAndServe(":8080", nil)

}

func dogs(w http.ResponseWriter, req *http.Request) {

	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}

}
