package main

import (
	"fmt"
)

func main() {
	var name, favfood, hobby string
	fmt.Println("Name")
	_, err := fmt.Scan(&name)
	if err != nil {
		panic(err)
	}
	fmt.Println("favourite food")
	_, err = fmt.Scan(&favfood)
	if err != nil {

		panic(err)
	}

	fmt.Println("Hobby")

	_, err = fmt.Scan(&hobby)
	if err != nil {

		panic(err)
	}
	fmt.Println(name, favfood, hobby)
}
