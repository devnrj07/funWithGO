package main

import(
	"fmt"
	"github.com/letsgo/basics/Documentation/dog"
	
)

type details struct{
	name string
	age int
}
func main() {
	sampleDog := details{"tommy",dog.Years(5)}
	fmt.Println(sampleDog)
	
	//fmt.Println("Enter your age : ")
	//_, err := fmt.Scan(&x)
	//fmt.Println("Your age in Doggo years : ", dog.Years(x))

}
