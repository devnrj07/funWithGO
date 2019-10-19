package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Create("log-file.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	read, err := os.Open("no-such-file.txt")

	if err != nil {
		//fmt.Println(err)
		log.Println("eh! Some error ocurred", err)
		//	log.Fatalln(err) // will print & exit the program
		//	panic()
	}

	defer read.Close()

	fmt.Println("New Error log file created")
}
