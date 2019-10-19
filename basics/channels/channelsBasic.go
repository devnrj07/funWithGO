package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, welcome to channels")

	c := make(chan int)

	go foo(c)

	// go bar(c) don't launch two go routines, they won't be able to interlock.THINK LOGIC
	bar(c) // this will wait for it to finish!! 
	fmt.Println("End of the channels")

}

//send channel
func foo(c chan<- int) {
	c <- 42 // 42 onto c: send 42 to with c
}

// receive channel
func bar(c <-chan int) {
	fmt.Println(<-c)	//Block it here! CHANNELS BLOCK !!
}
