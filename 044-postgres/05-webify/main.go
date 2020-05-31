//rdbms connectw/ postgres
// and perform a select
// and webify the main and batten down the hatches

package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

//Book is all paper & ink...
type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

func init() {

	var err error

	// open a connection
	connStr := "postgres://bond:password@localhost/bookstore?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// make sure we are connected
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("You are connected")

}

func main() {

	//defer db.Close()

	http.HandleFunc("/books", booksIndex)
	http.ListenAndServe(":8080", nil)

}

func booksIndex(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	//prep a query...
	rows, err := db.Query("select * from books;")
	if err != nil {
		//panic(err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	bks := make([]Book, 0)
	//perform the cursor loop until Next = false
	for rows.Next() {
		bk := Book{} //just use composite literal
		// scan the query result into our Book struct
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // order matters
		if err != nil {
			//panic(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		bks = append(bks, bk)
	}

	if err = rows.Err(); err != nil {
		//panic(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	//printout the resulting slice of Books
	for _, bk := range bks {
		//fmt.Printf("%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
		fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}

}
