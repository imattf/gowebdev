// simple context
//  with loading some user data

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ctx = context.WithValue(ctx, "userid", 777)
	ctx = context.WithValue(ctx, "fname", "james")

	results := dbAccess(ctx)

	log.Println(results)
	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) int {
	uid := ctx.Value("userid").(int)
	return uid
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println(ctx)
	fmt.Fprintln(w, ctx)

}
