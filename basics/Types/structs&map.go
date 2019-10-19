package main

import (
	"fmt"
)

type person struct{
	firstname string
	lastname string
	flavours []string
	
	}

func main() {
	
	
	p1 :=person{firstname : "nrj", lastname :"rocks"}
	p1.flavours =[]string{"strawberry","mango"}
	
	p2 :=person{lastname :"lmao",firstname : "noName", flavours : []string{"rose","butterscotch"}}
	p2.flavours = append(p2.flavours, "lemon", "orange")
	
	p3 :=person{lastname :"lola",firstname : "nona", flavours : []string{"fig&honey","vanilla"}}
	
	//var m map[string]person
         //m = make(map[string]person)
	//m[p1.lastname]= p1
	//m[p2.lastname]= p2
        //m[p3.lastname]= p3
    	
        new_map := map[string]person{
	   
	   p1.lastname : p1,
	   p2.lastname : p2,
	   p3.lastname : p3,
	
	}
	
	
	
	for _,v :=range new_map{
	  
	  fmt.Println(v.firstname)
	  fmt.Println(v.lastname)
	  for i,val := range v.flavours{
	       fmt.Println(i,val)
		} 
		fmt.Println("---------")
	}
}
