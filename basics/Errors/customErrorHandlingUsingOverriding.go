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
	fmt.Println("error inside foo : ->", e)
}

func main() {
	c := customErr{
		"lol there's nothing to tell",
		"lolmsg404",
	}

	foo(c)
}
