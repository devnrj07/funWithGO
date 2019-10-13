// varidiac parameter vs slice of a type

package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 4, 5, 7, 12, 46}
	s := foo(slice...)
	fmt.Println(s)

	s1 := bar(slice)
	fmt.Println(s1)
}

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

// so only variadic paramter need (unfurling ...) not the slices of TYPE
