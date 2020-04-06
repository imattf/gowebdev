package main

import (
	"io"
	"net/http"
)

func hd(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "bark, bark, bark")
}

func hc(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "meow, meow, meow")
}

func main() {

	http.HandleFunc("/dog/", hd)
	http.HandleFunc("/cat", hc)
	// http.Handle("/", hd)

	http.ListenAndServe(":8080", nil)

}
