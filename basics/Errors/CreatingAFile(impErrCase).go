package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	f, err := os.Create("newFile.txt") // creating a file using os return a *File and error
	if err != nil {
		fmt.Println(err)
		return // to exit maoin func
	}

	defer f.Close() // to any read/write operation. Due to 'defer' keyword

	r := strings.NewReader("Neeraj Mukta")
	io.Copy(f, r)

}
