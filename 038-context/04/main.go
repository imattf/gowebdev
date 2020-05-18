// Context makes it possible to manage a chain of calls within the same call path by signaling context’s Done channel.
// source: https://rakyll.org/leakingctx/
// this is the leaky verion because a context is not being set to interrupt the go routine

package main

import (
	"fmt"
	"time"
)

func main() {

	//start this mind bender...

	fmt.Println("starting main...")
	for n := range gen() {
		fmt.Println(" running n", n)
		if n == 5 {
			break
		}
	}
	fmt.Println(" going to sleep...")
	time.Sleep(1 * time.Minute)
	fmt.Println(" ...waking up")

	fmt.Println("...stoping main")

}

//gen is a broken generator that will leak a go routine
// the go routine will keep on running even after n == 5 above
func gen() <-chan int {

	fmt.Println("...inside gen")
	ch := make(chan int)
	go func() {
		var n int
		fmt.Println("   ...inside anonymous func")
		for {
			ch <- n
			n++
			fmt.Println("        ...running goroutine:", n) //looks to me like it stops running here, ???
		}
	}()
	fmt.Println("...leaving gen")
	return ch
}
