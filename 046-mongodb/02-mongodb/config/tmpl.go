package config

import "html/template"

var Tmpl *template.Template

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

	Tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}
