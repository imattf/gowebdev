//This code is the same code from folder 044_postgres / 22_delete
// ... sections moved to books and packages

package main

import (
	"net/http"

	//"github.com/imattf/gowebdev/045-code-organization/03-multiple-packages/books"
	"github.com/imattf/gowebdev/046-mongodb/02-mongodb/books"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/books", books.Index)
	http.HandleFunc("/books/show", books.Show)
	http.HandleFunc("/books/shop", books.Shop)
	http.HandleFunc("/books/create", books.Create)
	http.HandleFunc("/books/create/process", books.CreateProcess)
	http.HandleFunc("/books/update", books.Update)
	http.HandleFunc("/books/update/process", books.UpdateProcess)
	http.HandleFunc("/books/delete/process", books.DeleteProcess)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/books", http.StatusSeeOther) //go see books
}
