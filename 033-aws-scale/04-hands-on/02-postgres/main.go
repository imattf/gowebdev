//rdbms SQL w/ mysql

package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func main() {

	// see info: https://pkg.go.dev/github.com/lib/pq@v1.5.2?tab=doc
	connStr := "host=database-1.caqfkpxxkdnj.us-east-2.rds.amazonaws.com user=postgres password=mypostgres dbname=mypostgres sslmode=disable"
	db, err = sql.Open("postgres", connStr)

	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/instance", instance)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":80", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "Successfully completed.")
	check(err)
}

func ping(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "OK")
}

func instance(w http.ResponseWriter, req *http.Request) {

	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	if err != nil {
		fmt.Println(err)
		return
	}

	bs := make([]byte, resp.ContentLength)
	resp.Body.Read(bs)
	resp.Body.Close()
	io.WriteString(w, string(bs))
}

func amigos(w http.ResponseWriter, req *http.Request) {

	rows, err := db.Query(`SELECT aname FROM amigos;`)
	check(err)
	defer rows.Close()

	//data to be used in query
	var s, name string
	s = "RETRIEVED RECORDS:\n"

	//query
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, req *http.Request) {

	stmt, err := db.Prepare(`CREATE TABLE customer (name text);`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "CREATED TABLE customer", n)
}

func insert(w http.ResponseWriter, req *http.Request) {

	stmt, err := db.Prepare(`INSERT INTO customer("name") VALUES('James');`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "INSERTED RECORD", n)
}

func read(w http.ResponseWriter, req *http.Request) {

	rows, err := db.Query(`SELECT * FROM customer;`)
	check(err)
	defer rows.Close()

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Fprintln(w, "RETRIEVED RECORDS:", name)
	}
}

func update(w http.ResponseWriter, req *http.Request) {

	stmt, err := db.Prepare(`UPDATE customer SET name='Jimmy' WHERE name='James';`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "UPDATED RECORD", n)
}

func delete(w http.ResponseWriter, req *http.Request) {

	stmt, err := db.Prepare(`DELETE FROM customer WHERE name='Jimmy';`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "DELETED RECORD", n)
}

func drop(w http.ResponseWriter, req *http.Request) {

	stmt, err := db.Prepare(`DROP TABLE customer;`)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(w, "DROPPED TABLE customer")
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}

}