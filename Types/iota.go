package main

import (
	"fmt"
)
const ( // iota is like a counter starts with 0 keeps incrementing until a new const() is defined. And it only works inside a const() :((
	x = iota
	y = 2017 + iota
	z = y + iota
	w = z + iota

)
func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(w)
}
