// Take the previous program in the previous folder and change it so that:
// - a template is parsed and served
// - you pass data into the template

package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("foo.gohtml"))
}

func main() {

	// type conversion to a Handler
	http.HandleFunc("/", defawlt)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.HandleFunc("/poof/", poof)

	http.ListenAndServe(":8080", nil)

}

func defawlt(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "default stuf happens here")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "bark, bark, bark")
}

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hi Matthew Faulkner")
}

func poof(res http.ResponseWriter, req *http.Request) {
	err := tmpl.ExecuteTemplate(res, "foo.gohtml", `poof`)
	if err != nil {
		log.Fatalln(err)
	}
}
