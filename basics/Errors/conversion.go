package main

import (
	"fmt"
)

type hotdog int

func main() {
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	
	var y int
	y = int(x) // This is called conversion there's a difference
	fmt.Println(y)
	fmt.Printf("%T", y)
}

//	fmt.Println("foo ran -", e, "\n", e.(customErr).info)
