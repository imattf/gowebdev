// passing data
// query the form for a key, print the value

package main

import (
	"fmt"
	"net/http"
)

// var tmpl *template.Template

// func init() {
// 	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
// }

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	// http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {

	// write to console...
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")

}

func bar(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Your request method at bar: ", req.Method)
	http.Redirect(w, req, "/", http.StatusMovedPermanently)

}

// // redirects to /bar = bar-red
// func barred(w http.ResponseWriter, req *http.Request) {

// 	fmt.Println("Your request method at barred: ", req.Method)
// 	tmpl.ExecuteTemplate(w, "index.gohtml", nil)

// }
