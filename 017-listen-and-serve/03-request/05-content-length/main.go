package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	data := struct {
		Method      string
		URL         *url.URL
		Submissions url.Values
		// Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		Method:        req.Method,
		URL:           req.URL,
		Submissions:   req.Form,
		Header:        req.Header,
		Host:          req.Host,
		ContentLength: req.ContentLength,
	}

	//pass my own data
	tmpl.ExecuteTemplate(w, "index.gohtml", data)

	//fmt.Fprintln(w, "do what you want to do")
}

var tmpl *template.Template

func init() {
	//println("building template")
	tmpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	// can you print the value of d to the client later?
	d = 42
	http.ListenAndServe(":8080", d)

}
