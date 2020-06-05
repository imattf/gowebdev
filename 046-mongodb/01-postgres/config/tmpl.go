package config

import "html/template"

//Tmpl is a...
var Tmpl *template.Template

func init() {

	Tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}
