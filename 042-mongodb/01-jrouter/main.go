// this is the same as my main04-betterDefaultServeMux.go

package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	mux := httprouter.New()
	mux.GET("/", index)
	http.ListenAndServe(":8080", mux)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome here!")
}
