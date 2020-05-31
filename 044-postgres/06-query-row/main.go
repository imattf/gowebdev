//rdbms connectw/ postgres
// and perform a select
// and webify the main and batten down the hatches
// and query a row from the books table

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

	http.HandleFunc("/books", booksIndex)
	http.HandleFunc("/books/show", booksShow)
	http.HandleFunc("/books/shop", booksShop)
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
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	bks := make([]Book, 0)
	//execute the query cursor loop until Next = false
	for rows.Next() {
		bk := Book{} //just use composite literal
		// scan the query result into our Book struct
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // order matters
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		bks = append(bks, bk)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	//printout the resulting slice of Books
	for _, bk := range bks {
		//fmt.Printf("%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
		fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}

}

func booksShow(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	//get the isbn from query? parm off the URL
	// ...remember FormValue does both Form values OR URL query parms
	// ...see https://github.com/golang/go/blob/go1.14.3/src/net/http/request.go#L1325
	// ...curl -i localhost:8080/books/show?isbn=978-1505255607
	//     .... zsh has to have ? escaped like \?
	isbn := r.FormValue("isbn")
	if isbn == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	//prep a query...
	//pass in the isbn value from the first query parm
	row := db.QueryRow("select * from books where isbn = $1", isbn)

	bk := Book{}
	err := row.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // order matters
	switch {
	case err == sql.ErrNoRows: //no record found
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), 500) //unknown server error
		return
	}

	//printout the resulting slice of Books
	fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)

}

func booksShop(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "SHOPPING!")

}
