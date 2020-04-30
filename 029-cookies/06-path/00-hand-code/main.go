//01
package main

import (
	"net/http"
	"text/template"
)

//07
var tmpl *template.Template

//06
func init() {
	//08
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

//02
func main() {

	//04
	http.HandleFunc("/", index)

	//03
	http.ListenAndServe(":8080", nil)

}

//05
func index(w http.ResponseWriter, req *http.Request) {

	//09
	tmpl.ExecuteTemplate(w, "index.gohtml", nil)

}
