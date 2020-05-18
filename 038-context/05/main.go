// Context makes it possible to manage a chain of calls within the same call path by signaling context’s Done channel.
// source: https://rakyll.org/leakingctx/
// this is version uses a context to stop the go routine via the cancel()
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // make sure all paths cancel the context to avoid context leak

	//start this mind bender...

	fmt.Println("starting main...")
	for n := range gen(ctx) {
		fmt.Println(" running n", n)
		if n == 5 {
			cancel()
			break
		}
	}
	fmt.Println(" going to sleep...")
	time.Sleep(1 * time.Minute)
	fmt.Println(" ...waking up")

	fmt.Println("...stoping main")
}

//gen is a broken generator that will leak a go routine
func gen(ctx context.Context) <-chan int {

	fmt.Println("...inside gen")
	ch := make(chan int)
	go func() {
		var n int
		fmt.Println("   ...inside anonymous func")
		for {
			select {
			case <-ctx.Done():
				fmt.Println("        ...got a cancel, so stop")
				return // avoid leaking of this goroutine when ctx is done
			case ch <- n:
				n++
				fmt.Println("        ...running goroutine:", n)
			}
		}
	}()
	fmt.Println("...leaving gen")
	return ch
}
