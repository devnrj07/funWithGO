package main

import(
	"fmt"
	"C:/Microservices/cdap-customer-order-history-service/docker/dog"
	
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
