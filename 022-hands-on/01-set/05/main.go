// Take the previous program and change it so that:
// func main uses http.Handle instead of http.HandleFunc
// Contstraint: Do not change anything outside of func main

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
	http.Handle("/", http.HandlerFunc(defawlt))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))
	http.Handle("/poof/", http.HandlerFunc(poof))
	// http.HandleFunc("/", defawlt)
	// http.HandleFunc("/dog/", dog)
	// http.HandleFunc("/me/", me)
	// http.HandleFunc("/poof/", poof)

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
