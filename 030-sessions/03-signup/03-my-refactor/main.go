// hand build session

package main

import (
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

//userRecord
type user struct {
	UserName string
	Password string
	First    string
	Last     string
}

// user storage
var dbUsers = map[string]user{} //userID, user

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

	// //get cookie
	// cookie, err := req.Cookie("session")

	// // if it doesn't exist, create it
	// if err != nil {
	// 	sid, _ := uuid.NewV4()
	// 	cookie = &http.Cookie{
	// 		Name:  "session",
	// 		Value: sid.String(),
	// 		// Secure: true,
	// 		// HttpOnly: true,
	// 		// Path:     "/",
	// 	}
	// 	//set the cookie on the response
	// 	http.SetCookie(w, cookie)
	// }

	//get the user info
	u := getUser(req)

	//process form submission; store the user & session info
	// if req.Method == http.MethodPost {
	// 	un := req.FormValue("username")
	// 	fn := req.FormValue("firstname")
	// 	ln := req.FormValue("lastname")
	// 	u = user{un, fn, ln}
	// 	dbSessions[cookie.Value] = un
	// 	dbUsers[un] = u
	// }

	//show the index.html
	tmpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {

	//get the user info
	user := getUser(req)

	//if not logged in, go get logged in
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	//show the bar.html
	tmpl.ExecuteTemplate(w, "bar.gohtml", user)
}

func getUser(req *http.Request) user {

	var u user

	//get cookie
	cookie, err := req.Cookie("session")
	if err != nil {
		return u
	}

	// retrieve userID info from session cookie
	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}
	return u

}

func signup(w http.ResponseWriter, req *http.Request) {

	// //get cookie
	// cookie, err := req.Cookie("session")

	// // if it doesn't exist, create it
	// if err != nil {
	// 	sid, _ := uuid.NewV4()
	// 	cookie = &http.Cookie{
	// 		Name:  "session",
	// 		Value: sid.String(),
	// 		// Secure: true,
	// 		// HttpOnly: true,
	// 		// Path:     "/",
	// 	}
	// 	//set the cookie on the response
	// 	http.SetCookie(w, cookie)
	// }

	//redirect
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	//process form submission
	if req.Method == http.MethodPost {

		// get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		fn := req.FormValue("firstname")
		ln := req.FormValue("lastname")

		//user name taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		//create a session cookie
		sid, _ := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sid.String(),
		}
		http.SetCookie(w, cookie)
		dbSessions[cookie.Value] = un

		// store user in dbUsers
		u := user{un, p, fn, ln}
		dbUsers[un] = u

		//redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	//show the bar.html
	tmpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func alreadyLoggedIn(req *http.Request) bool {

	//get the session cookie
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[cookie.Value]
	_, ok := dbUsers[un]
	return ok
}

//
