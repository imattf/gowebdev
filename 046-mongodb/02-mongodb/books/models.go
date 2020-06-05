package books

import (
	"errors"
	"net/http"
	"strconv"

	//"github.com/imattf/gowebdev/045-code-organization/03-multiple-packages/config"
	"github.com/imattf/gowebdev/046-mongodb/02-mongodb/config"
	"gopkg.in/mgo.v2/bson"
)

//Book is all paper & ink...
type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

//AllBooks is a ...
func AllBooks() ([]Book, error) {

	//prep a query...
	//rows, err := config.Db.Query("select * from books;") //postgres
	bks := []Book{}

	err := config.Books.Find(bson.M{}).All(&bks)
	if err != nil {
		return nil, err
	}
	//defer rows.Close() // postgres
	return bks, nil

	// bks := make([]Book, 0)
	// //execute the query cursor loop until Next = false
	// for rows.Next() {
	// 	bk := Book{} //just use composite literal
	// 	// scan the query result into our Book struct
	// 	err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price) // order matters
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	bks = append(bks, bk)
	// }

	// if err = rows.Err(); err != nil {
	// 	return nil, err
	// }
	// return bks, nil

}

//OneBook is a ...
func OneBook(r *http.Request) (Book, error) {

	bk := Book{}
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return bk, errors.New("400. Bad Request")
	}
	//row := config.Db.QueryRow("select * from books where isbn = $1", isbn)
	//err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price) // order matters
	err := config.Books.Find(bson.M{"isbn": isbn}).One(&bk)
	if err != nil {
		return bk, err
	}
	return bk, nil

}

//PutBook is a ...
func PutBook(r *http.Request) (Book, error) {

	//get form values
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	//validate form values
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. Bad Request. Fields can't be empty")
	}

	//convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. Please hit back and enter number for price")
	}
	bk.Price = float32(f64)

	//insert values
	//_, err = config.Db.Exec("insert into books (isbn, title, author, price) values ($1, $2, $3, $4);", bk.Isbn, bk.Title, bk.Author, bk.Price)
	err = config.Books.Insert(bk)
	if err != nil {
		return bk, errors.New("500. Internal Server Error" + err.Error())
	}
	return bk, nil
}

//UpdateBook is a ...
func UpdateBook(r *http.Request) (Book, error) {

	//get form values
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	//validate form values
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. Bad Request. Fields can't be empty")
	}

	//convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. Please hit back and enter number for price")
	}
	bk.Price = float32(f64)

	//update values
	// _, err = config.Db.Exec("update books set isbn = $1, title=$2, author=$3, price=$4 where isbn=$1;", bk.Isbn, bk.Title, bk.Author, bk.Price)
	err = config.Books.Update(bson.M{"isbn": bk.Isbn}, &bk)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

//DeleteBook is a ...
func DeleteBook(r *http.Request) error {

	//get the isbn from query? parm off the URL
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return errors.New("400. Bad Eequest")
	}

	//delete book
	//_, err := config.Db.Exec("delete from books where isbn=$1;", isbn)
	err := config.Books.Remove(bson.M{"isbn": isbn})
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}
