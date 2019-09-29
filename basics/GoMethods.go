package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hi my Name is %s and age %d\n", p.first, p.age)
}

func main() {

	p1 := person{"NRJ", "ROCKS", 23}
	p2 := person{"Riya","hotty", 21}
	p1.speak()
	p2.speak()
}
