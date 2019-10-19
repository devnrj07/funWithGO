package main

import (
	"fmt"
)

type person struct{
	firstname string
	lastname string
	flavours []string
	
	}

func main() {
	fmt.Println("Hello, playground")
	
	p1 :=person{firstname : "nrj", lastname :"rocks"}
	p1.flavours =[]string{"strawberry","mango"}
	
	p2 :=person{lastname :"lmao",firstname : "noName", flavours : []string{"rose","butterscotch"}}
	p2.flavours = append(p2.flavours, "lemon", "orange")
	
	fmt.Println(p1)
	fmt.Println(p1.firstname)
	fmt.Println(p2)
	
	for i,v :=range p2.flavours{
	  fmt.Println(i,v)
	}
}
