package main

import (
	"fmt"
	"time"
)

func main() {
	for n := range gen() {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	time.Sleep(1 * time.Minute)
}

// gen is a broken generator that will leak a goroutine.
func gen() <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			ch <- n
			n++
			fmt.Println("        ...running goroutine:", n) //looks to me like it stops running here, ???
		}
	}()
	return ch
}
