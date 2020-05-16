// photo-blog
// store some user data w/ cookie; session uuid which will serve as the session ID when we need it
// append user file names to the session cookie value

package main

import (
	"html/template"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	//get session cookie
	cookie := getCookie(w, req)
	cookie = appendValue(w, cookie)
	xs := strings.Split(cookie.Value, "|")
	//pass a slice of bytes to range over in the template
	tmpl.ExecuteTemplate(w, "index.gohtml", xs)
}

//getCookie function out here
func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {

	cookie, err := req.Cookie("session")
	// if the session cookie doesn't exist, create it
	if err != nil {
		sid, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sid.String(),
		}
		//set the cookie on the response
		http.SetCookie(w, cookie)
	}
	return cookie
}

//append values to a past in cookie pointer
func appendValue(w http.ResponseWriter, c *http.Cookie) *http.Cookie {

	//values
	f1 := "one.jpg"
	f2 := "two.jpg"
	f3 := "three.jpg"

	//append
	cv := c.Value
	if !strings.Contains(cv, f1) {
		cv += "|" + f1
	}
	if !strings.Contains(cv, f2) {
		cv += "|" + f2
	}
	if !strings.Contains(cv, f3) {
		cv += "|" + f3
	}

	//update the cookie w/ new value on the response
	c.Value = cv
	http.SetCookie(w, c)
	return c

}
