package main

import (
	"io"
	"net/http"
)

type hotdog int

func (hd hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "bark, bark, bark")
}

type hotcat int

func (hc hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "meow, meow, meow")
}

func main() {
	var hd hotdog
	var hc hotcat
	// can you print the value of d to the client later?

	http.Handle("/dog/", hd)
	http.Handle("/cat", hc)
	// http.Handle("/", hd)

	http.ListenAndServe(":8080", nil)

}
