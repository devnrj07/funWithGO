package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	ExampleScannerLines()
}

//ExampleScannerLines function reads a values from console using os.Stdin
func ExampleScannerLines() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() { //Scan() returns a bool
		fmt.Println(scanner.Text()) // Text() returns the values read
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
