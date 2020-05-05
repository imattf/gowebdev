// hand build session

package main

import (
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

//userRecord
type userRecord struct {
	UserName string
	Password string
	First    string
	Last     string
}

// user storage
var dbUsers = map[string]userRecord{} //userID, user

// session storage
var dbSessions = map[string]string{} //sessionID, userID

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	//get the user
	u := getUser(req)

	//show the index.html
	tmpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {

	//get the user
	u := getUser(req)

	// if user doesn't exist, redirect to /
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	//show the bar.html
	tmpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	//process form submission
	if req.Method == http.MethodPost {

		//get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		fn := req.FormValue("firstname")
		ln := req.FormValue("lastname")

		//username taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		//create session
		sid, _ := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sid.String(),
		}
		//set the cookie on the response
		http.SetCookie(w, cookie)

		//store the session info from the cookie in dbSessions
		dbSessions[cookie.Value] = un

		//store the user in dbUsers
		u := userRecord{un, p, fn, ln}
		dbUsers[un] = u

		//redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	//show the signup.html
	tmpl.ExecuteTemplate(w, "signup.gohtml", nil)

}
