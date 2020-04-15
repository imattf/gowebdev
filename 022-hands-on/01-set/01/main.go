// ListenAndServe on port ":8080" using the default ServeMux.
// Use HandleFunc to add the following routes to the default ServeMux:
// "/" "/dog/" "/me/
// Add a func for each of the routes.
// Have the "/me/" route print out your name.

package main

import (
	"io"
	"net/http"
)

func main() {

	// type conversion to a Handler
	http.HandleFunc("/", defawlt)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)

}

func defawlt(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "default stuf happens here")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "bark, bark, bark")
}

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hi Matthew Faulkner")
}
