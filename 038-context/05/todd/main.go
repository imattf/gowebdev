package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // make sure all paths cancel the context to avoid context leak

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel()
			break
		}
	}
	time.Sleep(1 * time.Minute)
}

func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				fmt.Println("        ...goroutine cancelled")
				return // avoid leaking of this goroutine when ctx is done.
			case ch <- n:
				fmt.Println("        ...running goroutine:", n) //looks to me like it stops running here, ???
				n++
			}
		}
	}()
	return ch
}
