// Given n , print this series uptil n, for n = 7 --> 1 1 2 2 4 2 6

package main

import (
	"fmt"
)

func main() {

	for i := 1; i < 9; i++ {
		fmt.Println(euler(i), " ")
	}
}

func gcd(a int, b int) int {

	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func euler(n int) int {

	result := 1

	for i := 2; i < n; i++ {

		if gcd(i, n) == 1 {
			result += 1
		}
	}
	return result
}
