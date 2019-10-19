// --- Directions
// Write a function that accepts a positive number N.
// The function should console log a step shape
// with N levels using the # character.  Make sure the
// step has spaces on the right hand side!
// --- Examples
//   steps(2)
//       '# '
//       '##'
//   steps(3)
//       '#  '
//       '## '
//       '###'
//   steps(4)
//       '#   '
//       '##  '
//       '### '
//       '####'

package main

import (
	"fmt"
)

func main() {

	steps(4)
	fmt.Println("-----------------")
	recursiveSteps(4, 0, "")
}

// a very easy to solve such pattern problems is to imagine them in a matrix form rows & columns

func steps(n int) {

	//step := "" //not here
	for row := 0; row < n; row++ {
		step := "" // because we new string for new line not an appended string
		for column := 0; column < n; column++ {

			if row >= column {
				step += "#"
			} else {
				step += "-"
			}

		}
		fmt.Println(step)
	}
}

func recursiveSteps(n int, row int, step string) {

	if n == row {
		return	// exit condition or base case
	}

	if n == len(step) {
		fmt.Println(step)
		recursiveSteps(n, row+1, "")	// reset condition , switching to next row after traversing a row completely
	}

	if row >= len(step) { // remember the matrix, if row >= column (here,len(step)) add "#" to string else add " "
		step += "#"
	} else {
		step += "-"
	}
	recursiveSteps(n, row, step)
}
