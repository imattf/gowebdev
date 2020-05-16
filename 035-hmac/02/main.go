// hash message authentication code (HMAC) stuff
// using a cookie to store

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/authN", authN)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		c = &http.Cookie{
			Name:  "session",
			Value: "",
		}
	}
	if req.Method == http.MethodPost {
		e := req.FormValue("email")
		c.Value = e + "|" + getCode(e)
	}

	http.SetCookie(w, c)

	io.WriteString(w,
		`<!DOCTYPE html>
		<html>
			<body>
				<form method="POST">
				<input type="email" name="email">
				<input type="submit">
				</form>
				<a href="/authN">Validate This: `+c.Value+`</a>
			</body>
		</html>`)

}

func authN(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if c.Value == "" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	xs := strings.Split(c.Value, "|")
	email := xs[0]
	codeRcvd := xs[1]
	//codeCheck := getCode(email + "s")	//force a mis-match
	codeCheck := getCode(email) //should be equal

	if codeRcvd != codeCheck {
		fmt.Println("HMAC codes do not match")
		fmt.Println(codeRcvd)
		fmt.Println(codeCheck)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	io.WriteString(w,
		`<!DOCTYPE html>
		<html>
			<body>
				<h1>`+codeRcvd+` - RECIEVED </h1>
				<h1>`+codeCheck+` - RECALCULATED </h1>
			</body>
		</html>`)

}

func getCode(s string) string {
	hash := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(hash, s)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
