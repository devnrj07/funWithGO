package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M", // capital letters tells that they are to be Exported
		age:   54, // will throw error
	}

	users := []user{u1, u2, u3}

	//fmt.Println(users)

	// your code goes here
	j, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(j))
}
