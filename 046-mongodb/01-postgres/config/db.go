package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //this is a ...
)

//Db is a ...
var Db *sql.DB

func init() {

	var err error

	// open a connection
	connStr := "postgres://bond:password@localhost/bookstore?sslmode=disable"
	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// make sure we are connected
	err = Db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("You are connected")

}
