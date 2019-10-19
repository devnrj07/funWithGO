//Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

//For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

//Bonus: Can you do this in one pass?

package main

import (
	"fmt"
)

func main() {
	values := []int{10, 15, 3, 0, -6, 7}
	fmt.Println(arrayPairs(values, -4))

	values1 := []int{10, 15, 3, 0, -6, 7}
	fmt.Println(arrayPairs(values1, 4))

}

func arrayPairs(x []int, k int) bool {
	b := 0
	i := 0
	for i < len(x) {
		b = k - x[i] // find b = k-a
		//fmt.Println("b ", b)
		for j := i + 1; j < len(x); j++ {

			if x[j] == b { // use binary search to check if b exits
				fmt.Println(" x[] ", x[j], j)
				return true
			}
		}
		i += 1
	}
	return false

}
