package main

import (
	"fmt"
)

type person struct {
	name    string
	age     int
	awesome bool
}

func (p *person) speak() {
	fmt.Println(p.name, "is awesome ?", p.awesome)
}

func saySomething(h human) {
	h.speak()
}

type human interface {
	speak()
}

func main() {
	fmt.Println("Understanding of methods")

	p := person{"nrj", 22, true}
//	saySomething(p)  // this will not work because receiver of speak() method is pointer to person
	saySomething(&p) //This works as we pass & reference to the memory
	
	p.speak() // this will always work
}
