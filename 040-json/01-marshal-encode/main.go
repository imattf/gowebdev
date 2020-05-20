//marshal and encode stuff..package main

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//go to be capilized to exported to JSON
type person struct {
	Fname string
	Lname string
	Stuff []string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/encode", encode)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	s := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>FOO</title>
	</head>
	<body>
	You are in Foo
	</body>
	</html>
	`
	w.Write([]byte(s))

}

func marshal(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"Bilbo",
		"Baggins",
		[]string{"Cape", "Pipe", "Lunch"},
	}
	// a variable needed for the json object w/ Marshal
	// writes to an intermediate variable that needs to be written to the stream w
	jsn, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	w.Write(jsn)

}

func encode(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"Bilbo",
		"Baggins",
		[]string{"Cape", "Pipe", "Lunch"},
	}
	// no variable needed for the json object w/ Encode
	// writes to the stream automatically w
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}

}
