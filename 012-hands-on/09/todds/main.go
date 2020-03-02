// Parse this CSV file, putting two fields from the contents of the CSV file into a data structure.
// Parse an html template, then pass the data from step 1 into the CSV template;
// have the html template display the CSV data in a web page.

package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Record is the file layout
type Record struct {
	Date     string
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   int
	AdjClose float64
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {

	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {

	// read the csv file
	records := parseFile("table.csv")

	// build the template
	tmpl, err := template.ParseFiles("tmpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// myway...
	// execute template
	// err = tmpl.Execute(os.Stdout, records)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// todd's wat...
	// execute template
	err = tmpl.Execute(res, records)
	if err != nil {
		log.Fatalln(err)
	}
}

// process the csv file
func parseFile(filePath string) []Record {
	src, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()

	//read file contents
	readr := csv.NewReader(src)
	rows, err := readr.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	// load records
	records := make([]Record, 0, len(rows))
	for i, row := range rows {
		if i == 0 {
			continue
		}
		// date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)
		date := row[0]

		records = append(records, Record{
			Date: date,
			Open: open,
		})
	}

	return records

}
