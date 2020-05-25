// Now reorganize and refactor my code into controllers and models folders

package main

import (
	"net/http"

	"github.com/imattf/gowebdev/042-mongodb/04-controllers/controllers"

	"github.com/julienschmidt/httprouter"
)

func main() {

	mux := httprouter.New()
	//create a new user controller, by using the
	uc := controllers.NewUserController()
	//...when i do, i now have these nifty methods available
	mux.GET("/user/:id", uc.GetUser)
	mux.POST("/user", uc.CreateUser)
	mux.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe(":8080", mux)
}
