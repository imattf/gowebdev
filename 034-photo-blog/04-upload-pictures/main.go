// photo-blog
// store some user data w/ cookie; session uuid which will serve as the session ID when we need it
// append user file names to the session cookie values
// from a form page allowing user to upload file names to the server

package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

	if req.Method == http.MethodPost {

		//dealing with multi-part file, file header and error
		mpf, fhdr, err := req.FormFile("newfile")
		if err != nil {
			fmt.Println(err)
		}
		defer mpf.Close()

		//create the sha file from the uploaded file...
		// parse uploaded File Header's Filename for the extension to use as a suffix on the new sha file
		ext := strings.Split(fhdr.Filename, ".")[1]
		// create a sha version for the uploaded file
		//  see https://pkg.go.dev/crypto/sha1?tab=doc#pkg-overview for the weirdness on this
		hash := sha1.New()
		io.Copy(hash, mpf)
		// build new file with the sha file name, using %x to get a hexidecimal value
		fname := fmt.Sprintf("%x", hash.Sum(nil)) + "." + ext

		//create new file
		// get your current working directory
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		// build the path name of current working directory
		// with the path to where we want to store the uploaded files
		// with the actual file name (sha version)
		path := filepath.Join(wd, "public", "pics", fname)
		// create the new file, which is empty to start
		newfile, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer newfile.Close()

		//Copy contents into newly created server file
		// reset our read/write head to the bginning of the file
		//   because our sha work put the pointer for this file at the end of it
		mpf.Seek(0, 0)
		// copy contents from uploaded file into new file
		io.Copy(newfile, mpf)

		// add filename to this users session cookie
		cookie = appendValue(w, cookie, fname)

	}

	//Display the contents of the session cookie on the page
	xs := strings.Split(cookie.Value, "|")
	// pass a slice of bytes to range over in the template
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
func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {

	//append
	cv := c.Value
	if !strings.Contains(cv, fname) {
		cv += "|" + fname
	} else {
		fmt.Println("that file is already in the cookie")
	}

	//update the cookie w/ new value on the response
	c.Value = cv
	http.SetCookie(w, c)
	return c

}
