package main

import (
	"fmt"
)

func main() {
	fmt.Println(lcm(4, 20))
	fmt.Println(lcm(4, 16))
	fmt.Println(lcm(14, 20))
	fmt.Println(lcm(15, 25))
	fmt.Println(lcm(0, 20))
	fmt.Println(lcm(4, 0))
}

//Lcm is the least number that is divisible by both a & b for ex. 75 is lcm of 15 and 25, similarly, 4,20 lcm is 40
//One way is to find prime factors, pair down the factors and multiple the singleton factors or
// simple use formula : a * b = lcm(a,b) * gcd(a,b)
func lcm(a int, b int) int {

	gcd := gcd(a, b)

	lcm := (a * b) / gcd
	return lcm
}

func gcd(a int, b int) int {

	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
