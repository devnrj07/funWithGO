package main

import (
	"fmt"
)

func main() {
	fmt.Println(gcd(14, 36))
}

//greatest common number that can divided both the numbers. use Euclidean Algorithm
func gcd(a int, b int) int {

	if a == 0 {
		return b
	}

	return gcd(b%a, a)
}
