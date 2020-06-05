package config

import (
	"fmt"

	// _ "github.com/lib/pq" //this is a ...
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2"
)

//Db is a ...
//var Db *sql.DB

//DB is a database
var DB *mgo.Database

//Books is a collection
var Books *mgo.Collection

func init() {

	var err error

	// Connect to our local mongo
	//s, err := mgo.Dial("mongodb://localhost/bookstore")
	s, err := mgo.Dial("mongodb://bond:moneypenny007@localhost/bookstore")

	if err != nil {
		fmt.Println("¡ay chihauhua... mongo es no bueno!")
		panic(err)
	}

	// make sure we are connected
	err = s.Ping()
	if err != nil {
		panic(err)
	}

	DB = s.DB("bookstore")
	Books = DB.C("books")

	fmt.Println("You are connected to mongo")

}
