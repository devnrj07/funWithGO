package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	
	anonymous_type :=struct{
	                 cars map[int]string
			 colors []string
	                       
	
	}{cars : map[int]string{1 : "dodge",2 : "NSW",}, 
						colors : []string{"red","black","neon-green"}}

        fmt.Println(anonymous_type)
           
for k,v :=range anonymous_type.cars{
    fmt.Println(k,v)
}

for i,v :=range anonymous_type.colors{
	fmt.Println(i,v)
}
}
