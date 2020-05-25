// Now reorganize and refactor my code into controllers and models folders

package main

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/imattf/gowebdev/042-mongodb/05-mongodb/01-update-user-controller/controllers"

	"github.com/julienschmidt/httprouter"
)

func main() {

	mux := httprouter.New()
	//create a new user controller, by using the
	uc := controllers.NewUserController(getSession())
	//...when i do, i now have these nifty methods available
	mux.GET("/user/:id", uc.GetUser)
	mux.POST("/user", uc.CreateUser)
	mux.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe(":8080", mux)
}

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		fmt.Println("¡ay chihauhua... mongo es no bueno!")
		panic(err)
	}
	fmt.Println("¡ariba... mongo es muy bueno!")
	return s
}
