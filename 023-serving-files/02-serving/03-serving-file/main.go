// serve locally

package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", dog)
	// http.HandleFunc("/toby.jpg", dogPic)
	http.HandleFunc("/dogit", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// todd says image won't serve in previous sample, and it does not render there, but does so here
	// ...because the img tag makes an http callto get the file
	// io.WriteString(w, `<img src="/toby.jpg">`)
	io.WriteString(w, `<img src="/dogit">`)

}

func dogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "toby.jpg")
}
