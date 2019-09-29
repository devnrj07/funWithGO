package main

import (
	"fmt"
)

func main() {
	x := [][]string{}
	y := []string{"1", "2", "3", "4", "5", "6"}
	z := []string{"one", "two", "third", "fourth", "five", "six"}
	x = [][]string{y, z}

	for i, v := range x {
		fmt.Println("record :", i)
		for j, r := range v {
			fmt.Printf("\tindex: %v ==> value : %v", j, r)
		}
	}
	fmt.Println(x)
}
