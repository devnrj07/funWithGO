package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	defer foo()
	f, err := os.Open("xxx.txt")
	if err != nil {
		log.Panicln(err)
	}

	defer f.Close() // even 'deffered' statements don't work

	fmt.Println("It should exit without reaching here in case of any error")
}

func foo() {
	fmt.Println("!!! When os.Exit() is called nothing is allowed to run not even defer func !!!")
}

// Panicln is equivalent to Println() followed by a call to panic()
