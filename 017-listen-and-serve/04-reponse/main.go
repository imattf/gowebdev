package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("MeClould-Key", "this is mine")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "<h1>do what you want to do</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
