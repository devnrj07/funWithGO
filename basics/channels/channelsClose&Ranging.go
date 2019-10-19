package main

import (
	"fmt"
)

func main() {
	fmt.Println("Leran to close channels")

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {

			c <- i
		}
		close(c) // closing channel inorder for range to stop waiting to pull off values

	}()

	for v := range c {
		fmt.Println(v)
	} // ranging over channel to pull off the values
	//may cause deadlock if channel remains opened

}

//func foo(c chan<- int) {

//for i:= 0 ; i<10 ; i++ {

//	c <- i
//}
//close(c) // closing channel inorder for range to stop waiting to pull off values
//}
