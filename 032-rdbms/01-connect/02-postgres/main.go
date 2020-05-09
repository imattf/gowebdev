//rdbms connectw/ postgres

package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func main() {

	// see info: https://pkg.go.dev/github.com/lib/pq@v1.5.2?tab=doc
	// connStr := "host=database-1.caqfkpxxkdnj.us-east-2.rds.amazonaws.com user=postgres password=mypostgres dbname=postgres sslmode=disable"
	connStr := "host=database-1.caqfkpxxkdnj.us-east-2.rds.amazonaws.com user=postgres password=mypostgres dbname=mypostgres sslmode=disable"
	db, err = sql.Open("postgres", connStr)

	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "Successfully completed.")
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}

}
