// serve the AJAX enabled index.html file

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	tmpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	s := `Here is some text from foo`
	fmt.Fprintln(w, s)
}
