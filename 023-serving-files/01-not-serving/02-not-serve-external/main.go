// not serve because...
// ??serving from external server

package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// todd says image won't serve, and it does not render
	// ...because img tag doesn't have the actual file resource
	//    it has to be served up via http
	io.WriteString(w, `<img src="/toby.jpg">`)

}
