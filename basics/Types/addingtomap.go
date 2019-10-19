package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`bond`:   []string{`James`, `Spy`, `Gadgets`},
		`walker`: []string{`alan`, `edm`, `djkit`},
		`parker`: []string{`peter`, `hero`, `web`},
	}
        //adding a new key value pair
	m[`mccleod`] = []string{`Tom`, `instructor`, `laptop`}
	for i, v := range m {

		fmt.Printf("key:=> %s\n", i)
		for j, val := range v {

			fmt.Printf("index: %v \t details: %v\n", j, val)
		}
	}

}
