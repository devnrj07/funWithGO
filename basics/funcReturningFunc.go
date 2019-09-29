package main

import (
	"fmt"
)

func main() {

	//var trap func()
	//trap = friend()
	//trap()

	trap := friend()
	fmt.Println(trap())

}

func friend() func() int{

	return func() int { //fmt.Println("This is an anonymuous func used for returning a func() type!~")
		return 69
	}
}
