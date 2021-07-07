package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	f, err := os.Open("newFile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close() // always close the resource heavy operations

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}
