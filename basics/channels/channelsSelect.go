package main

import (
	"fmt"
)

func main() {
	fmt.Println("Select statement to receive values from channels")

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)

	receive(even, odd, quit)

}

func send(e, o, q chan<- int) {

	for i := 0; i < 50; i++ {
		if i&1 != 1 {
			e <- i
		} else {
			o <- i
		}

	}
	close(e)
	close(o)
	q <- 0 // important to close channels to prevent range clause from going crazy
}

func receive(e, o, q <-chan int) {

	for { // infinte loop

		select {
		case v := <-e:
			fmt.Println("from the even channel : ", v)
		case v := <-o:
			fmt.Println("from the odd channel : ", v)
		case v := <-q:
			fmt.Println("from the quit channel : ", v)
			return
		}

	}

}
