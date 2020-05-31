//This code is the same code from folder 044_postgres / 22_delete
// ... sections moved to models package

package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	//"github.com/GoesToEleven/golang-web-dev/045-code-organization/02_two-packages/models"
	"github.com/imattf/gowebdev/045-code-organization/02-two-packages/models"
	_ "github.com/lib/pq"
)

// var db *sql.DB
var tmpl *template.Template

// //Book is all paper & ink...
// type Book struct {
// 	Isbn   string
// 	Title  string
// 	Author string
// 	Price  float32
// }

func init() {

	// var err error

	// // open a connection
	// connStr := "postgres://bond:password@localhost/bookstore?sslmode=disable"
	// db, err = sql.Open("postgres", connStr)
	// if err != nil {
	// 	panic(err)
	// }

	// // make sure we are connected
	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("You are connected")

	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/books", booksIndex)
	http.HandleFunc("/books/show", booksShow)
	http.HandleFunc("/books/shop", booksShop)
	http.HandleFunc("/books/create", booksCreateForm)
	http.HandleFunc("/books/create/process", booksCreateProcess)
	http.HandleFunc("/books/update", booksUpdateForm)
	http.HandleFunc("/books/update/process", booksUpdateProcess)
	http.HandleFunc("/books/delete/process", booksDeleteProcess)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/books", http.StatusSeeOther) //go see books
}

func booksIndex(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	//new stuff...
	bks, err := models.AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	// //prep a query...
	// rows, err := db.Query("select * from books;")
	// if err != nil {
	// 	http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	// 	return
	// }
	// defer rows.Close()

	// bks := make([]Book, 0)
	// //execute the query cursor loop until Next = false
	// for rows.Next() {
	// 	bk := Book{} //just use composite literal
	// 	// scan the query result into our Book struct
	// 	err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price) // order matters
	// 	if err != nil {
	// 		http.Error(w, http.StatusText(500), 500)
	// 		return
	// 	}
	// 	bks = append(bks, bk)
	// }

	// if err = rows.Err(); err != nil {
	// 	http.Error(w, http.StatusText(500), 500)
	// 	return
	// }

	//show the book
	tmpl.ExecuteTemplate(w, "books.gohtml", bks)

}

func booksShow(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	//new stuff...
	bk, err := models.OneBook(r)

	// //get the isbn from query? parm off the URL
	// isbn := r.FormValue("isbn")
	// if isbn == "" {
	// 	http.Error(w, http.StatusText(400), http.StatusBadRequest)
	// 	return
	// }
	// //prep a query...
	// //pass in the isbn value from the first query parm
	// row := db.QueryRow("select * from books where isbn = $1", isbn)

	// bk := Book{}
	// err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price) // order matters

	switch {
	case err == sql.ErrNoRows: //no record found
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError) //unknown server error
		return
	}

	//show the book
	tmpl.ExecuteTemplate(w, "show.gohtml", bk)

}

func booksShop(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "SHOPPING!")
}

func booksCreateForm(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "create.gohtml", nil)
}

func booksCreateProcess(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	//new stuff...
	bk, err := models.PutBook(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable) //bad form values input
		return
	}

	// //get form values
	// bk := Book{}
	// bk.Isbn = r.FormValue("isbn")
	// bk.Title = r.FormValue("title")
	// bk.Author = r.FormValue("author")
	// p := r.FormValue("price")

	// //validate form values
	// if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
	// 	http.Error(w, http.StatusText(400), http.StatusMethodNotAllowed) //bad form values input
	// 	return
	// }

	// //convert form values
	// f64, err := strconv.ParseFloat(p, 32)
	// if err != nil {
	// 	http.Error(w, http.StatusText(406)+"Please hit back and enter number for price.", http.StatusNotAcceptable) //bad form values input
	// 	return
	// }
	// bk.Price = float32(f64)

	// //insert values
	// _, err = db.Exec("insert into books (isbn, title,  author, price) values ($1, $2, $3, $4)", bk.Isbn, bk.Title, bk.Author, bk.Price)
	// if err != nil {
	// 	http.Error(w, http.StatusText(500), http.StatusInternalServerError) //unknown server error
	// 	return
	// }

	//confirm insertion
	tmpl.ExecuteTemplate(w, "created.gohtml", bk)
}

func booksUpdateForm(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	//new stuff...
	bk, err := models.OneBook(r)

	// //get the isbn from query? parm off the URL
	// isbn := r.FormValue("isbn")
	// if isbn == "" {
	// 	http.Error(w, http.StatusText(400), http.StatusBadRequest)
	// 	return
	// }
	// //prep a query...
	// //pass in the isbn value from the first query parm
	// row := db.QueryRow("select * from books where isbn = $1", isbn)

	// bk := Book{}
	// err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price) // order matters

	switch {
	case err == sql.ErrNoRows: //no record found
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError) //unknown server error
		return
	}

	//show the book
	tmpl.ExecuteTemplate(w, "update.gohtml", bk)

}

func booksUpdateProcess(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	//new stuff...
	bk, err := models.UpdateBook(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable) //bad form values input
		return
	}

	// //get form values
	// bk := Book{}
	// bk.Isbn = r.FormValue("isbn")
	// bk.Title = r.FormValue("title")
	// bk.Author = r.FormValue("author")
	// p := r.FormValue("price")

	// //validate form values
	// if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
	// 	http.Error(w, http.StatusText(400), http.StatusMethodNotAllowed) //bad form values input
	// 	return
	// }

	// //convert form values
	// f64, err := strconv.ParseFloat(p, 32)
	// if err != nil {
	// 	http.Error(w, http.StatusText(406)+"Please hit back and enter number for price.", http.StatusNotAcceptable) //bad form values input
	// 	return
	// }
	// bk.Price = float32(f64)

	// //insert values
	// _, err = db.Exec("update books set isbn = $1, title=$2, author=$3, price=$4 where isbn=$1;", bk.Isbn, bk.Title, bk.Author, bk.Price)
	// //_, err = db.Exec("insert into books (isbn, title,  author, price) values ($1, $2, $3, $4)", bk.Isbn, bk.Title, bk.Author, bk.Price)
	// if err != nil {
	// 	http.Error(w, http.StatusText(500), http.StatusInternalServerError) //unknown server error
	// 	return
	// }

	//confirm insertion
	tmpl.ExecuteTemplate(w, "updated.gohtml", bk)
}

func booksDeleteProcess(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	//new stuff...
	err := models.DeleteBook(r)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// //get the isbn from query? parm off the URL
	// isbn := r.FormValue("isbn")
	// if isbn == "" {
	// 	http.Error(w, http.StatusText(400), http.StatusBadRequest)
	// 	return
	// }

	// //delete book
	// _, err := db.Exec("delete from books where isbn=$1;", isbn)
	// if err != nil {
	// 	http.Error(w, http.StatusText(500), http.StatusInternalServerError) //unknown server error
	// 	return
	// }

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
