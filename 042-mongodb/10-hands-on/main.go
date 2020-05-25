package main

import (
	"html/template"
	"net/http"

	"github.com/imattf/gowebdev/042-mongodb/10-hands-on/controllers"
	// "github.com/imattf/gowebdev/042-mongodb/11-solution/controllers"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	c := controllers.NewController(tmpl)
	http.HandleFunc("/", c.Index)
	http.HandleFunc("/bar", c.Bar)
	http.HandleFunc("/signup", c.SignUp)
	http.HandleFunc("/login", c.Login)
	http.HandleFunc("/logout", c.Logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
