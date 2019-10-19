package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, Welcome to buffio example")

	const input = `So this is what you meant
			When you said that you were spent
			And now it's time to build from the bottom of the pit, right to the top
			Don't hold back
			Packing my bags and giving the Academy a rain check
			
			I don't ever want to let you down
			I don't ever want to leave this town
			'Cause after all
			This city never sleeps at night

			It's time to begin, isn't it?
			I get a little bit bigger but then I'll admit
			I'm just the same as I was
			Now don't you understand
			I'm never changing who I am`

	//buffio takes a reader and writer
	scanner := bufio.NewScanner(strings.NewReader(input))
	//scanner.Split(bufio.ScanRunes) // set up a splitFunc ex ScanRunes

	count := 0

	for scanner.Scan() {
		count++                // last \n will be added by go for termination
		line := scanner.Text() // this will read line by line as we don't want to split it into words UNCOMMENT LINE 31 and see the difference
		fmt.Println(line)
	}

	//Err() method throws an error: Go supports declaration and expression
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(count)

}
