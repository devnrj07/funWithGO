package main

//Go program to convert int to string and viceversa
//bar takes string and int as argument returns int and string 
import (
	"fmt"
	"strconv"
)

func main() {

	a := foo("lol")
	fmt.Println(a)
	a1,b := bar("lol", 100)
	fmt.Printf("a1 type %T \t b of type %T",a1,b)
}
func foo(x string) int {
	i, _ := strconv.Atoi(x)
	return i
}

func bar(x string, y int) (int, string) {
	i, _ := strconv.Atoi(x)
	str := strconv.Itoa(y)

	return i, str
}
