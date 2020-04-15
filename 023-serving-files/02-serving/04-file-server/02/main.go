// serve locally

package main

import (
	"io"
	"net/http"
)

func main() {

	// serve up everything in this specific directory
	// ...now limit to a specific sub-directory /assets using StripPrefix()

	//http.HandleFunc("/dog", dog)
	http.HandleFunc("/", dog)

	// http.Handle("/resources", http.FileServer(http.Dir(".")))
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))

	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//io.WriteString(w, `<img src="toby.jpg">`)
	io.WriteString(w, `<img src="/resources/toby.jpg">`)

}
