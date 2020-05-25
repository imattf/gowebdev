package main

import (
	"net/http"

	"github.com/imattf/gowebdev/042-mongodb/08-hands-on/controllers"
	"github.com/imattf/gowebdev/042-mongodb/08-hands-on/models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() map[string]models.User {
	return controllers.LoadUsers()
}
