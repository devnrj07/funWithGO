package main

import (
	"fmt"
)

type customErr struct {
	msg string
	id  string
}

//implementing Error interface by overrinding Error method
func (cE customErr) Error() string {
	return fmt.Sprintf("here's the detailed error%v : with errorId : %v ", cE.msg, cE.id)
}

func foo(e error) {
	//fmt.Println("error inside foo : ->", e.msg) //Now, this won't work bcause obiviously it's error Type not customErr Type
	fmt.Println("error inside foo : ->", e.(customErr).msg) //This is called ASSERTION implicitly telling complier to change the Type
}

func main() {
	c := customErr{
		"lol there's nothing to tell",
		"lolmsg404",
	}

	foo(c)
}
