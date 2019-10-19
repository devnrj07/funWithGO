package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 4, 6, 8, 16}
	fmt.Println(GCD(gcd, arr))
}

func gcd(a int, b int) int {

	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

// function to GCD of multiple numbers. Simply use GCD(gcd(gcd(a,b),c),...) use gcd(result, arr[0])
func GCD(f func(a, b int) int, arr []int) int {

	result := arr[0]
	for i := 1; i < len(arr); i++ {

		f(arr[i], result)
	}
	return result
}
