// hand build session

package main

import (
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"

	uuid "github.com/satori/go.uuid"
)

//user Record
type user struct {
	UserName string
	Password []byte
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

	//prep for login test
	bs, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	dbUsers["bob@aol.com"] = user{"bob@aol.com", bs, "Bilbo", "Baggins"}

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
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

	//see if already logged in
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	var u user

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

		//encrypt the password field
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		//store the user in dbUsers
		u := user{un, bs, fn, ln}
		dbUsers[un] = u

		//redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	//show the signup.html
	tmpl.ExecuteTemplate(w, "signup.gohtml", u)

}

func login(w http.ResponseWriter, req *http.Request) {

	//see if already logged in
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	var u user

	//process form submission
	if req.Method == http.MethodPost {

		//get form values
		un := req.FormValue("username")
		p := req.FormValue("password")

		//is there a matching user name?
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username/password do not match", http.StatusForbidden)
			return
		}

		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username/password do not match", http.StatusForbidden)
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

		//redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	//show the login.html
	tmpl.ExecuteTemplate(w, "login.gohtml", u)

}

func logout(w http.ResponseWriter, req *http.Request) {
	// if not logged in, redirect to /
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	//find the session cookie
	cookie, _ := req.Cookie("session")

	//delete the session from dbSession
	delete(dbSessions, cookie.Value)

	//remove the session cookie
	cookie = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	//set the cookie on the response
	http.SetCookie(w, cookie)

	//redirect
	http.Redirect(w, req, "/", http.StatusSeeOther)

}
