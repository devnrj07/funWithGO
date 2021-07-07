/*Recover is a built-in function that regains control of a panicking goroutine.
//!!!!!!!Recover is only useful inside deferred functions. !!!!!!!!
During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking,
 a call to recover will capture the value given to panic and resume normal execution.
*/

package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() { // important to use recover as a defer statement
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r) // it will exit from here
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.") // this will never get executed as g(0) causes panic and control is taken over by recover
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i)) // panic is called so now control will goto recover
		// As a result all the defer statements will excute from the stack
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
