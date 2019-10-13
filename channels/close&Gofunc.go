package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

// returns an receiver channel of int so, let's range over it
func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(rc <-chan int) {

	for v := range rc {
		fmt.Println(v)
	}

}
