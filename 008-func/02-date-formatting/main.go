// date formatting

package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.New("").Funcs(fmap).ParseFiles("tmpl.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

var fmap = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func main() {

	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
