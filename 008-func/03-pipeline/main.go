// date formatting

package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.New("").Funcs(fmap).ParseFiles("tmpl.gohtml"))
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

var fmap = template.FuncMap{
	"fdouble": double,
	"fsquare": square,
	"fsqRoot": sqRoot,
}

func main() {

	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", 3)
	if err != nil {
		log.Fatalln(err)
	}
}
