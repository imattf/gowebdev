package controllers

import (
	"html/template"
	"net/http"

	"github.com/imattf/gowebdev/042-mongodb/10-hands-on/session"
)

//Controller is a ...
type Controller struct {
	tmpl *template.Template
}

//NewController is a...
func NewController(t *template.Template) *Controller {
	return &Controller{t}
}

//Index is a ...
func (c Controller) Index(w http.ResponseWriter, req *http.Request) {
	u := session.GetUser(w, req)
	session.ShowSessions() // for demonstration purposes
	c.tmpl.ExecuteTemplate(w, "index.gohtml", u)
}

//Bar is a ...
func (c Controller) Bar(w http.ResponseWriter, req *http.Request) {
	u := session.GetUser(w, req)
	if !session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	session.ShowSessions() // for demonstration purposes
	c.tmpl.ExecuteTemplate(w, "bar.gohtml", u)
}
