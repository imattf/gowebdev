//rdbms connectw/ postgres

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func main() {

	// test connection
	connStr := "postgres://bond:password@localhost/bookstore?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("You are connected")

}
