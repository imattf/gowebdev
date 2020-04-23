// passing data
// query the form for a key, print the value

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", foo)
	//this appears to only make a difference in behavior w/ Chrome based browsers
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {

	// putting a value into a url: localhost:8080/?q=dog
	// retrieve value from url with FormValue
	v := req.FormValue("q")
	y := req.FormValue("p")

	fmt.Fprintln(w, "Do my bidding: "+y)
	io.WriteString(w, "Do my search: "+v)
	// either way to resond on the writer if fine

	//visit localhost:8080/?q=dog

	//try localhost:8080/?q=dog&p=cat
}
