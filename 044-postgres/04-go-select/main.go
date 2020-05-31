//rdbms connectw/ postgres
// and perform a select

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

func main() {

	// open a connection
	connStr := "postgres://bond:password@localhost/bookstore?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// make sure we are connected
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("You are connected")

	//prep a query...
	rows, err := db.Query("select * from books;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	bks := make([]Book, 0)
	//perform the cursor loop until Next = false
	for rows.Next() {
		bk := Book{} //just use composite literal
		// scan the query result into our Book struct
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // order matters
		if err != nil {
			panic(err)
		}
		bks = append(bks, bk)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	//printout the resulting slice of Books
	for _, bk := range bks {
		fmt.Printf("%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}

}
