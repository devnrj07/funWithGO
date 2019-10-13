package main

import (
	"fmt"
)

func main() {

	defer open()
	fmt.Println("Inside main func")

}

func open() {
	defer func() {
		fmt.Println("This function closes the gates")
	}()

	fmt.Println("This function opens the gate")
}

//o/p without defer
/*
This function closes the gates
This function opens the gate
Inside main func

--> with defer at open() o/p :-
Inside main func
This function closes the gates
This function opens the gate

--> with defer inside open() o/p :-
Inside main func
This function opens the gate
This function closes the gates

--> defer: halts or stops execution until the outer fuction has either returned/exited
*/
