// passing data
// POSTing a form

package main

import (
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

	// retrieve value from url with FormValue
	v := req.FormValue("q")
	// y := req.FormValue("p")

	// going to write out some html, se we need to set-up a header packet up some html
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w,
		`<form method="GET">
		   <input type="text" name="q">
		   <input type="submit">
		 </form>
		<br>`+v)
}

// study post vs get in the form method parm
// get puts tha value in the url and on the page
// post just puts it in the form
