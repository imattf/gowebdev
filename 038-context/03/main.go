// simple context
//  with loading some user data
//  and adding timing a context out

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
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

	results, err := dbAccess(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}
	log.Println(results)
	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) (int, error) {

	//inject a time-out into the context
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	//make a channel that runs a long time to show the time-out behaviour
	ch := make(chan int)
	go func() {
		//ridiculous long running task
		uid := ctx.Value("userid").(int)
		time.Sleep(10 * time.Second)

		//check to make sure we're not running in vain
		// if ctx.Done() has
		if ctx.Err() != nil {
			return
		}
		ch <- uid

	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case i := <-ch:
		return i, nil
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println(ctx)
	fmt.Fprintln(w, ctx)

}
