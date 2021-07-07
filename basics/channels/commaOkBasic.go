package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, welcome to comma ok example")

	c := make(chan int)

	go func() {
		c <- 42
		close(c)

	}()

	v, ok := <-c
	fmt.Println("before", v, ok)

	v, ok = <-c
	fmt.Println("before", v, ok)

}
