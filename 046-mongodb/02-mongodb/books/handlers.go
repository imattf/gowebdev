package books

import (
	"database/sql"
	"fmt"
	"net/http"

	//"github.com/imattf/gowebdev/045-code-organization/03-multiple-packages/config"
	"github.com/imattf/gowebdev/046-mongodb/02-mongodb/config"
)

//Index is a ...
func Index(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	//new stuff...
	bks, err := AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	//show the book
	config.Tmpl.ExecuteTemplate(w, "books.gohtml", bks)

}

//Show is a ...
func Show(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	//new stuff...
	bk, err := OneBook(r)

	switch {
	case err == sql.ErrNoRows: //no record found
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError) //unknown server error
		return
	}

	//show the book
	config.Tmpl.ExecuteTemplate(w, "show.gohtml", bk)

}

//Shop is a ...
func Shop(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "SHOPPING!")
}

//Create is a...
func Create(w http.ResponseWriter, r *http.Request) {
	config.Tmpl.ExecuteTemplate(w, "create.gohtml", nil)
}

//CreateProcess is a ...
func CreateProcess(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	//new stuff...
	bk, err := PutBook(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable) //bad form values input
		return
	}

	//confirm insertion
	config.Tmpl.ExecuteTemplate(w, "created.gohtml", bk)
}

//Update is a ...
func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	//new stuff...
	bk, err := OneBook(r)

	switch {
	case err == sql.ErrNoRows: //no record found
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError) //unknown server error
		return
	}

	//show the book
	config.Tmpl.ExecuteTemplate(w, "update.gohtml", bk)

}

//UpdateProcess is a...
func UpdateProcess(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	//new stuff...
	bk, err := UpdateBook(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable) //bad form values input
		return
	}

	//confirm insertion
	config.Tmpl.ExecuteTemplate(w, "updated.gohtml", bk)
}

//DeleteProcess is a ...
func DeleteProcess(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	//new stuff...
	err := DeleteBook(r)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
