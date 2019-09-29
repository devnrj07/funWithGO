package main

import (
	"fmt"
)

func main() {

	values := []int{1, 3, 5, 9, 12, 15, 18, 22, 30, 42}
	fmt.Println(sumOfFirstLast(values)) //43

	x := sumFirstLastDivisibleBy15(sumOfFirstLast, values)
	fmt.Println(x)// new slice will have only {15,30}
}

//This will sum the first and last digit of a given slice

func sumOfFirstLast(x []int) int {

	return x[0] + x[len(x)-1]
}

//This will sum First and last only for values with lcm 15 using previous func as an argument
func sumFirstLastDivisibleBy15(f func(x []int) int, y []int) int {

	new_y := []int{}
	for _, v := range y {
		if v%15 == 0 {
			new_y = append(new_y, v)
		}
	}

	return f(new_y)
}
