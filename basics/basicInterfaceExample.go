package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64 // this is important to make connection with abover defined area() method of return type float64
}

func info(sh shape) {
	fmt.Println("The area is", sh.area())
}

func main() {
	box := square{4}
	ring := circle{2}

	info(box)
	info(ring)
}

//Remember an Interface says : hey if you have my method than you're my TYPE. so methods are basically the cohesion point for structs and interface to achieve polymorphism
