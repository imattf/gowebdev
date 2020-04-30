// Experiment with the "path" attribute on cookies
// Try setting a cookie when at this path "/dog/browzer"
//   and see if you can access it at "/cat" and/or "/dog/bowzer"
// Try setting a cookie when at this path "/"
//   and see if you can access it at "/dog/bowzer" and/or "/"

package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tmpl *template.Template

func init() {

	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/dog/bowzer", bowzer)
	http.HandleFunc("/dog/pictures", pictures)
	http.HandleFunc("/cat", cat)

	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request) {

	//go get the cookie
	var c *http.Cookie // var is a type pointer to a Cookie
	c, err := req.Cookie("user-cookie")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%T\n", c)
	}

	tmpl.ExecuteTemplate(w, "index.gohtml", c)
}

func bowzer(w http.ResponseWriter, req *http.Request) {

	// create a cookie
	c := &http.Cookie{
		Name:  "user-cookie",
		Value: "this would be the value",
		Path:  "/",
		// Path:  "/dog/bowzer",
	}
	http.SetCookie(w, c)

	tmpl.ExecuteTemplate(w, "bowzer.gohtml", c)

}

func pictures(w http.ResponseWriter, req *http.Request) {

	//go get the cookie
	var c *http.Cookie
	c, err := req.Cookie("user-cookie")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%T\n", c)
	}

	tmpl.ExecuteTemplate(w, "pictures.gohtml", c)
}

func cat(w http.ResponseWriter, req *http.Request) {

	//go get the cookie
	var c *http.Cookie
	c, err := req.Cookie("user-cookie")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%T\n", c)
	}

	tmpl.ExecuteTemplate(w, "cat.gohtml", c)
}
