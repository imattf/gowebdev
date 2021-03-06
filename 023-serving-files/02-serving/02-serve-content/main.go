// serve locally

package main

import (
	"io"
	"net/http"
	"os"
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
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "file NOT found", 404)
		return
	}
	defer f.Close()

	//need extra file meta before calling http.ServeContent
	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not FOUND", 404)
		return
	}

	//serve by writing to response writer w/ http.ServeContent()
	//more meta-data in play, e.g. mod-time
	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
