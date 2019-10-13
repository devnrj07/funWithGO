package main

import (
	"fmt"
	"os"
)

func main() {

	defer foo()
	f, err := os.Open("xxx.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close() // even 'deffered' statements don't work

	fmt.Println("It should exit without reaching here in case of any error")
}

func foo() {
	fmt.Println("!!! When panic is called defer is allowed to run  !!!")
}

// Panicln is equivalent to Println() followed by a call to panic()
