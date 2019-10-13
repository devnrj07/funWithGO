package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := sum(slice...)
	fmt.Println("sum :", s1)
	s2 := sumOfOdd(sum, slice...)
	fmt.Println("sumOfOdds : ", s2)
}

//function to return sum of given slice of int
func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

// a callback function: a function who calls 'sum' func for addition i.e taking a function as an argument
func sumOfOdd(f func(x ...int) int, slice ...int) int {

	//creating a temp slice to hold only odd values from a given slice
	var odd []int

	for _, v := range slice {
		if v%2 != 0 {
			odd = append(odd, v)
		}
	}
	//pass slice of odd integers to the function to perform sum operation
	return f(odd...)
}
