package main

import (
	"fmt"
)

func main() {

	f := loopfactorial(6)
	fmt.Println(f)
}

func loopfactorial(x int) int {
	fact := 1
	for; x>0 ; x--  {
		fact = fact * x
	}
	return fact
}
