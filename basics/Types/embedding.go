package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{vehicle : vehicle{ doors : 2, color : "black"}, fourWheel : false}
	t2 := truck{vehicle : vehicle{ doors : 3, color : "white"}, fourWheel : true}
	
	s1 := sedan{vehicle : vehicle{ doors : 4, color : "blue"}, luxury :true}
	s2 := sedan{vehicle : vehicle{ doors : 4, color : "red"}, luxury :true}

	  fmt.Println(t1,t2)
	  fmt.Println(s1,s2)
	
}
