// # Starting with the code in the "starting-files" folder:
// ## wire this program up so that it works
// ParseGlob in an init function
// Use HandleFunc for each of the routes
// Combine apply & applyProcess into one func called "apply"
// Inside the func "apply", use this code to create the logic to respond differently to a POST method and a GET method
// if req.Method == http.MethodPost {
//     		// code here
//     		return

package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/about/", about)
	http.HandleFunc("/contact/", contact)
	http.HandleFunc("/apply", apply)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, req *http.Request) {
	err := tmpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, req *http.Request) {
	err := tmpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		err := tmpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
		HandleError(w, err)
		return
	}

	err := tmpl.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)

}

// HandleError is ...
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
