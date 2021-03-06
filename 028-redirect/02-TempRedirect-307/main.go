// passing data
// query the form for a key, print the value

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

	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {

	// write to page...
	// fmt.Fprintln(w, "Do my bidding: foo")
	// write to page...
	// io.WriteString(w, "Do my search: foo")

	// write to console...
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")

}

func bar(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Your request method at bar: ", req.Method)
	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)

}

// redirects to /bar = bar-red
func barred(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Your request method at barred: ", req.Method)
	tmpl.ExecuteTemplate(w, "index.gohtml", nil)

}
