package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	//println("building template")
	tmpl = template.Must(template.ParseFiles("index.gohtml"))
}

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	//.Postform reads form data only
	//.Form reads URL & form data
	tmpl.ExecuteTemplate(w, "index.gohtml", req.Form)

	//fmt.Fprintln(w, "do what you want to do")
}

func main() {
	var d hotdog
	// can you print the value of d to the client later?
	d = 42
	http.ListenAndServe(":8080", d)

}
