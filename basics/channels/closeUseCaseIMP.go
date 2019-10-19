package main

import (
	"fmt"
)

func main() {
	fmt.Println("Fan out In")

	fan := fanOut()
	fanIn(fan)

}

func fanOut() <-chan int {

	channel := make(chan int)

	for i := 0; i < 10; i++ {

		go func() {
			for j := 0; j < 10; j++ {
				channel <- j
			}
		}()

	}
	// close(channel) // don't close it 

	return channel
}

func fanIn(channel <-chan int) {

	for i := 0; i < 100; i++ {
		fmt.Println(<-channel)
	}
}
