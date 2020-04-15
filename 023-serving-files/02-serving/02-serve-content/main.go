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
	// http.HandleFunc("/chat.png", catPic)
	http.HandleFunc("/dogit", dogPic)
	http.HandleFunc("/catit", catPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// todd says image won't serve in previous sample, and it does not render there, but does so here
	// ...because the img tag makes an http callto get the file
	// io.WriteString(w, `<img src="/toby.jpg">`)
	// io.WriteString(w, `<img src="/chat.png">`)
	io.WriteString(w, `<img src="/dogit">`)
	io.WriteString(w, `<img src="/catit">`)

}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "file NOT found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}

func catPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("chat.png")
	if err != nil {
		http.Error(w, "file NOT found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}
