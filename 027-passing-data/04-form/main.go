// passing data
// POSTing a form

package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {

	// retrieve values from url with FormValue
	f := req.FormValue("first")
	l := req.FormValue("last")
	s := req.FormValue("subscribed") == "on"

	err := tmpl.ExecuteTemplate(w, "index.gohtml", person{f, l, s})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
