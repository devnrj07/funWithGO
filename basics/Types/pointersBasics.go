package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	p.first = "default" // shorthand for (*p).first
	(*p).last = "xxx"
	(*p).age = 22

}

func main() {

	p1 := person{"slash", "hudson", 55}
	changeMe(&p1)
	fmt.Println(p1)
	
	}

// * -> pointer of a type
// & -> address of the variable also used for deferencing a pointer or getting the value stored at that location	
