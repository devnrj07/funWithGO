//Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

//For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

//Bonus: Can you do this in one pass?

package main

import (
	"fmt"
)

func main() {
	x := []int{10, 15, 3, 6}
	fmt.Println("x :", sumPair(x, 16))

	y := []int{12, 3, 4, 7, 7, 9, 9, 0, 9, 7, 2, 8, 5}
	fmt.Println("y :", sumPair(y, 8))
}

// basically checking if sum - i , i pair exits. Suppose we have sum = a + b.
func sumPair(arr []int, sum int) bool {

	tempMap := map[int]int{}

	for i := 0; i < len(arr); i++ {

		if _, exists := tempMap[sum-arr[i]]; exists { //containskey() :If there's already sum - b in the set/map/list then we found the pair b, sum -b
			return true
		} else {

			tempMap[arr[i]] = i // If not simply store the value arr[i] with its index i into the set/map/list
		}

	}
	return false
}
