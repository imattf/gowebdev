// Cross-site scripting (XSS) is a type of computer security vulnerability
// typically found in web applications.
// XSS enables attackers to inject client-side scripts
// into web pages viewed by other users.

// vulnerable version

package main

import (
	"html/template"
	"log"
	"os"
)

// Page is a typical DOM
type Page struct {
	Title   string
	Heading string
	Input   string
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {

	myPage := Page{
		Title:   "Not Vulnerable",
		Heading: "Not Vulnerable w/ html/template",
		Input:   `<script>Alert("XSS bad!");</script>`,
	}

	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", myPage)
	if err != nil {
		log.Fatalln(err)
	}

}
