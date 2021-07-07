package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	ExampleWriter()
	ExampleScanner_custom()
}

func ExampleWriter() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "world!")
	w.Flush() // Don't forget to flush!
	// Output: Hello, world!
}

// Use a Scanner with a custom split function (built by wrapping ScanWords) to validate
// 32-bit decimal input.
func ExampleScanner_custom() {
	// An artificial input source.
	const input = "1234 5678 1234567901234567890"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Create a custom split function by wrapping the existing ScanWords function.
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}
	// Set the split function for the scanning operation.
	scanner.Split(split)
	// Validate the input
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
	// Output:
	// 1234
	// 5678
	// Invalid input: strconv.ParseInt: parsing "1234567901234567890": value out of range
}
