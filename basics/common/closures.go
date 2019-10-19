package main

import (
	"fmt"
)

func main() {

	f := foo() // enclosed in this scope
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	g := foo() // new variable with new scope, this doesn't have access to f's scope, hence the variable gets reset
	fmt.Println(g())
	fmt.Println(g())
}

func foo() func() int {

	x := 0
	return func() int {
		x++ //closuring x to stack level
		return x
	}

}
